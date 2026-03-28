package packages

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_distribution_point"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/crypto"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/tools/upload_counter"
	"resty.dev/v3"
)

type (
	// Service handles communication with the packages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages
	Packages struct {
		client client.Client
		// Added for convenience helpers to refresh cloud distribution point inventory.
		cloudDistributionPoint *cloud_distribution_point.CloudDistributionPoint
	}
)

func NewPackages(client client.Client) *Packages {
	return &Packages{
		client:                 client,
		cloudDistributionPoint: cloud_distribution_point.NewCloudDistributionPoint(client),
	}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Packages CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all package objects (Get Package objects).
// URL: GET /api/v1/packages
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages
func (s *Packages) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProPackagesV1

	mergePage := func(pageData []byte) error {
		var items []ResourcePackage
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list packages: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified package by ID (Get specified Package object).
// URL: GET /api/v1/packages/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id
func (s *Packages) GetByIDV1(ctx context.Context, id string) (*ResourcePackage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPackagesV1, id)

	var result ResourcePackage

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new package record (Create Package record).
// URL: POST /api/v1/packages
// Body: JSON with metadata (name, category, info, notes, priority, etc.) - no file upload
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages
func (s *Packages) CreateV1(ctx context.Context, request *RequestPackage) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	cdp, _, err := s.cloudDistributionPoint.GetV1(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("preflight: failed to check cloud distribution point: %w", err)
	}

	if cdp.CdnType == cloud_distribution_point.CdnTypeNone {
		return nil, nil, fmt.Errorf("The content delivery network to use for the distribution point must be enabled first before creating a package")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProPackagesV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadV1 uploads a package file to an existing package record.
// URL: POST /api/v1/packages/{id}/upload
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-upload
func (s *Packages) UploadV1(ctx context.Context, id string, filePath string) (*CreateResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}
	if filePath == "" {
		return nil, nil, fmt.Errorf("file path is required")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open package file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat package file: %w", err)
	}

	endpoint := fmt.Sprintf("%s/%s/upload", constants.EndpointJamfProPackagesV1, id)
	fileName := info.Name()
	if fileName == "" {
		fileName = filePath
	}

	var result CreateResponse

	bar := upload_counter.New(nil)
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetMultipartFile("file", fileName, f, info.Size(), client.MultipartProgressCallback(bar.Callback)).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified package by ID (Update specified Package object).
// URL: PUT /api/v1/packages/{id}
// Body: JSON with full ResourcePackage - no file upload
// https://developer.jamf.com/jamf-pro/reference/put_v1-packages-id
func (s *Packages) UpdateByIDV1(ctx context.Context, id string, request *ResourcePackage) (*ResourcePackage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPackagesV1, id)

	var result ResourcePackage

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AssignManifestToPackageV1 assigns a manifest file to an existing package.
// URL: POST /api/v1/packages/{id}/manifest
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-manifest
func (s *Packages) AssignManifestToPackageV1(ctx context.Context, id string, manifestPath string) (*CreateResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}
	if manifestPath == "" {
		return nil, nil, fmt.Errorf("manifest path is required")
	}

	f, err := os.Open(manifestPath)
	if err != nil {
		return nil, nil, fmt.Errorf("open manifest file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat manifest file: %w", err)
	}

	endpoint := fmt.Sprintf("%s/%s/manifest", constants.EndpointJamfProPackagesV1, id)
	fileName := info.Name()
	if fileName == "" {
		fileName = manifestPath
	}

	var result CreateResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetMultipartFile("file", fileName, f, info.Size(), nil).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeletePackageManifestV1 removes the manifest from a package.
// URL: DELETE /api/v1/packages/{id}/manifest
// https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id-manifest
func (s *Packages) DeletePackageManifestV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/manifest", constants.EndpointJamfProPackagesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByIDV1 removes the specified package by ID (Remove specified Package record).
// URL: DELETE /api/v1/packages/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id
func (s *Packages) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPackagesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeletePackagesByIDV1 deletes multiple packages by their IDs (Delete multiple Packages by their IDs).
// URL: POST /api/v1/packages/delete-multiple
// Body: JSON with ids (array of package IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-delete-multiple
func (s *Packages) DeletePackagesByIDV1(ctx context.Context, req *DeletePackagesByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProPackagesV1 + "/delete-multiple"

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV1 returns the history object for the specified package.
// URL: GET /api/v1/packages/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id-history
func (s *Packages) GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPackagesV1, id)

	var result HistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV1 adds notes to the specified package history.
// URL: POST /api/v1/packages/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-history
func (s *Packages) AddHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPackagesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ExportV1 exports the packages collection as CSV or JSON.
// URL: POST /api/v1/packages/export
// Query params: export-fields, export-labels, page, page-size, sort, filter
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-export
func (s *Packages) ExportV1(ctx context.Context, rsqlQuery map[string]string, body *ExportRequest, acceptType string) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProPackagesExport
	if acceptType == "" {
		acceptType = constants.ApplicationJSON
	}

	var reqBody any
	if body != nil {
		reqBody = body
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", acceptType).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetBody(reqBody).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export packages: %w", err)
	}
	return resp.Bytes(), resp, nil
}

// ExportHistoryV1 exports the package history for a specified package as CSV or JSON.
// URL: POST /api/v1/packages/{id}/history/export
// Query params: export-fields, export-labels, page, page-size, sort, filter
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-history-export
func (s *Packages) ExportHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string, body *ExportRequest, acceptType string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s%s", constants.EndpointJamfProPackagesV1, id, constants.EndpointJamfProPackagesHistoryExport)
	if acceptType == "" {
		acceptType = constants.ApplicationJSON
	}

	var reqBody any
	if body != nil {
		reqBody = body
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", acceptType).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetBody(reqBody).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export package history: %w", err)
	}
	return resp.Bytes(), resp, nil
}

// CreateAndUpload creates package metadata, uploads the file, and verifies SHA3-512.
// Flow: 1) Calculate SHA3-512 and MD5 of local file; 2) Create metadata (CDP preflight);
// 3) Upload file; 4) Refresh cloud distribution point inventory; 5) Poll until hashValue
// (SHA3-512) is populated; 6) Verify hash matches local file.
func (s *Packages) CreateAndUpload(ctx context.Context, filePath string, req *RequestPackage) (*CreateResponse, *resty.Response, error) {
	if filePath == "" {
		return nil, nil, fmt.Errorf("file path is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	initialHash, err := crypto.CalculateSHA3_512(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("SHA3_512: %w", err)
	}
	md5Hash, err := crypto.CalculateMD5(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("MD5: %w", err)
	}

	createReq := *req
	createReq.FileName = filepath.Base(filePath)
	createReq.MD5 = md5Hash

	created, resp, err := s.CreateV1(ctx, &createReq)
	if err != nil {
		return nil, resp, fmt.Errorf("create metadata: %w", err)
	}
	packageID := created.ID

	_, resp, err = s.UploadV1(ctx, packageID, filePath)
	if err != nil {
		return nil, resp, fmt.Errorf("upload file: %w", err)
	}

	const maxAttempts = 60
	const sleepBetween = 3 * time.Second
	var uploaded *ResourcePackage
	for i := 1; i <= maxAttempts; i++ {
		_, refreshErr := s.cloudDistributionPoint.RefreshInventoryV1(ctx, createReq.FileName)
		if refreshErr != nil {
			return created, resp, fmt.Errorf("refresh cloud distribution point inventory: %w", refreshErr)
		}
		uploaded, resp, err = s.GetByIDV1(ctx, packageID)
		if err != nil {
			return created, resp, fmt.Errorf("get package (attempt %d/%d): %w", i, maxAttempts, err)
		}
		if uploaded.HashValue != "" {
			break
		}
		if i < maxAttempts {
			time.Sleep(sleepBetween)
		}
	}

	if uploaded.HashValue == "" {
		return created, resp, fmt.Errorf("timed out waiting for SHA3_512 hash after %d attempts", maxAttempts)
	}
	if uploaded.HashValue != initialHash {
		return created, resp, fmt.Errorf("hash verification failed: initial=%s uploaded=%s", initialHash, uploaded.HashValue)
	}

	return created, resp, nil
}

// UpdateAndUpload updates package metadata, uploads a new file, and verifies SHA3-512.
// Flow: 1) Calculate SHA3-512 and MD5 of local file; 2) Update metadata; 3) Upload file;
// 4) Refresh cloud distribution point inventory (forces immediate hash recalculation);
// 5) Poll until hashValue (SHA3-512) is populated; 6) Verify hash matches local file.
func (s *Packages) UpdateAndUpload(ctx context.Context, id string, filePath string, req *ResourcePackage) (*ResourcePackage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}
	if filePath == "" {
		return nil, nil, fmt.Errorf("file path is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	initialHash, err := crypto.CalculateSHA3_512(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("SHA3_512: %w", err)
	}
	md5Hash, err := crypto.CalculateMD5(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("MD5: %w", err)
	}

	updateReq := *req
	updateReq.FileName = filepath.Base(filePath)
	updateReq.MD5 = md5Hash

	updated, resp, err := s.UpdateByIDV1(ctx, id, &updateReq)
	if err != nil {
		return nil, resp, fmt.Errorf("update metadata: %w", err)
	}

	_, resp, err = s.UploadV1(ctx, id, filePath)
	if err != nil {
		return nil, resp, fmt.Errorf("upload file: %w", err)
	}

	const maxAttempts = 60
	const sleepBetween = 3 * time.Second
	var uploaded *ResourcePackage
	for i := 1; i <= maxAttempts; i++ {

		_, refreshErr := s.cloudDistributionPoint.RefreshInventoryV1(ctx, updateReq.FileName)
		if refreshErr != nil {
			return updated, resp, fmt.Errorf("refresh cloud distribution point inventory: %w", refreshErr)
		}

		uploaded, resp, err = s.GetByIDV1(ctx, id)
		if err != nil {
			return updated, resp, fmt.Errorf("get package (attempt %d/%d): %w", i, maxAttempts, err)
		}
		if uploaded.HashValue != "" {
			break
		}
		if i < maxAttempts {
			time.Sleep(sleepBetween)
		}
	}

	if uploaded.HashValue == "" {
		return updated, resp, fmt.Errorf("timed out waiting for SHA3_512 hash after %d attempts", maxAttempts)
	}
	if uploaded.HashValue != initialHash {
		return updated, resp, fmt.Errorf("hash verification failed: initial=%s uploaded=%s", initialHash, uploaded.HashValue)
	}

	return updated, resp, nil
}
