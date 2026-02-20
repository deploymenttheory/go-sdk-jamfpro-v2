package packages

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/crypto"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// PackagesServiceInterface defines the interface for package operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages
	PackagesServiceInterface interface {
		// ListPackagesV1 returns all package objects (Get Package objects).
		//
		// Returns a paged list of package objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages
		ListPackagesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetPackageByIDV1 returns the specified package by ID (Get specified Package object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id
		GetPackageByIDV1(ctx context.Context, id string) (*ResourcePackage, *interfaces.Response, error)

		// CreatePackageV1 creates a new package record (Create Package record).
		//
		// Creates metadata only; file upload is a separate step via UploadPackageV1.
		// Required: PackageName, FileName, CategoryID, Priority, and the seven *bool fields.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-packages
		CreatePackageV1(ctx context.Context, req *RequestPackage) (*CreateResponse, *interfaces.Response, error)

		// UploadPackageV1 uploads a package file to an existing package record.
		//
		// Call CreatePackageV1 first to create metadata, then upload the file.
		// filePath is the path to the package file on disk.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-upload
		UploadPackageV1(ctx context.Context, id string, filePath string) (*CreateResponse, *interfaces.Response, error)

		// UpdatePackageByIDV1 updates the specified package by ID (Update specified Package object).
		//
		// Sends full ResourcePackage (typically from GetPackageByIDV1, then modify and PUT).
		// Metadata only; no file upload.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-packages-id
		UpdatePackageByIDV1(ctx context.Context, id string, req *ResourcePackage) (*ResourcePackage, *interfaces.Response, error)

		// AssignManifestToPackageV1 assigns a manifest file to an existing package.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-manifest
		AssignManifestToPackageV1(ctx context.Context, id string, manifestPath string) (*CreateResponse, *interfaces.Response, error)

		// DeletePackageManifestV1 removes the manifest from a package.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id-manifest
		DeletePackageManifestV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeletePackageByIDV1 removes the specified package by ID (Remove specified Package record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id
		DeletePackageByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeletePackagesByIDV1 deletes multiple packages by their IDs (Delete multiple Packages by their IDs).
		//
		// Sends a POST to /api/v1/packages/delete-multiple with a body containing package IDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-packages-delete-multiple
		DeletePackagesByIDV1(ctx context.Context, req *DeletePackagesByIDRequest) (*interfaces.Response, error)

		// GetPackageHistoryV1 returns the history object for the specified package.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id-history
		GetPackageHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddPackageHistoryNotesV1 adds notes to the specified package history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-history
		AddPackageHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error)

		// DoPackageUpload creates package metadata, uploads the file, and verifies SHA3_512.
		//
		// Flow: 1) Calculate SHA3_512 and MD5 of local file; 2) Create metadata (FileName, MD5 set from file);
		// 3) Upload file; 4) Poll until HashType==SHA3_512 and HashValue populated; 5) Verify hash.
		DoPackageUpload(ctx context.Context, filePath string, req *RequestPackage) (*CreateResponse, *interfaces.Response, error)
	}

	// Service handles communication with the packages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-packages
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PackagesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Packages CRUD Operations
// -----------------------------------------------------------------------------

// ListPackagesV1 returns all package objects (Get Package objects).
// URL: GET /api/v1/packages
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages
func (s *Service) ListPackagesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointPackagesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPackageByIDV1 returns the specified package by ID (Get specified Package object).
// URL: GET /api/v1/packages/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id
func (s *Service) GetPackageByIDV1(ctx context.Context, id string) (*ResourcePackage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPackagesV1, id)

	var result ResourcePackage

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreatePackageV1 creates a new package record (Create Package record).
// URL: POST /api/v1/packages
// Body: JSON with metadata (name, category, info, notes, priority, etc.) - no file upload
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages
func (s *Service) CreatePackageV1(ctx context.Context, req *RequestPackage) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointPackagesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadPackageV1 uploads a package file to an existing package record.
// URL: POST /api/v1/packages/{id}/upload
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-upload
func (s *Service) UploadPackageV1(ctx context.Context, id string, filePath string) (*CreateResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}
	if filePath == "" {
		return nil, nil, fmt.Errorf("file path is required")
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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

	endpoint := fmt.Sprintf("%s/%s/upload", EndpointPackagesV1, id)
	fileName := info.Name()
	if fileName == "" {
		fileName = filePath
	}

	var result CreateResponse
	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, f, info.Size(), nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdatePackageByIDV1 updates the specified package by ID (Update specified Package object).
// URL: PUT /api/v1/packages/{id}
// Body: JSON with full ResourcePackage - no file upload
// https://developer.jamf.com/jamf-pro/reference/put_v1-packages-id
func (s *Service) UpdatePackageByIDV1(ctx context.Context, id string, req *ResourcePackage) (*ResourcePackage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPackagesV1, id)

	var result ResourcePackage

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AssignManifestToPackageV1 assigns a manifest file to an existing package.
// URL: POST /api/v1/packages/{id}/manifest
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-manifest
func (s *Service) AssignManifestToPackageV1(ctx context.Context, id string, manifestPath string) (*CreateResponse, *interfaces.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s/manifest", EndpointPackagesV1, id)
	fileName := info.Name()
	if fileName == "" {
		fileName = manifestPath
	}

	var result CreateResponse
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, f, info.Size(), nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeletePackageManifestV1 removes the manifest from a package.
// URL: DELETE /api/v1/packages/{id}/manifest
// https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id-manifest
func (s *Service) DeletePackageManifestV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/manifest", EndpointPackagesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeletePackageByIDV1 removes the specified package by ID (Remove specified Package record).
// URL: DELETE /api/v1/packages/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-packages-id
func (s *Service) DeletePackageByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPackagesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeletePackagesByIDV1 deletes multiple packages by their IDs (Delete multiple Packages by their IDs).
// URL: POST /api/v1/packages/delete-multiple
// Body: JSON with ids (array of package IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-delete-multiple
func (s *Service) DeletePackagesByIDV1(ctx context.Context, req *DeletePackagesByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointPackagesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetPackageHistoryV1 returns the history object for the specified package.
// URL: GET /api/v1/packages/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-packages-id-history
func (s *Service) GetPackageHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("package ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointPackagesV1, id)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddPackageHistoryNotesV1 adds notes to the specified package history.
// URL: POST /api/v1/packages/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-packages-id-history
func (s *Service) AddPackageHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("package ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointPackagesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DoPackageUpload creates package metadata, uploads the file, and verifies SHA3_512.
// Flow: 1) Calculate SHA3_512 and MD5 of local file; 2) Create metadata; 3) Upload file;
// 4) Poll until HashType==SHA3_512; 5) Verify hash.
func (s *Service) DoPackageUpload(ctx context.Context, filePath string, req *RequestPackage) (*CreateResponse, *interfaces.Response, error) {
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

	created, resp, err := s.CreatePackageV1(ctx, &createReq)
	if err != nil {
		return nil, resp, fmt.Errorf("create metadata: %w", err)
	}
	packageID := created.ID

	_, resp, err = s.UploadPackageV1(ctx, packageID, filePath)
	if err != nil {
		return nil, resp, fmt.Errorf("upload file: %w", err)
	}

	const maxAttempts = 10
	const sleepBetween = 2 * time.Second
	var uploaded *ResourcePackage
	for i := 1; i <= maxAttempts; i++ {
		uploaded, resp, err = s.GetPackageByIDV1(ctx, packageID)
		if err != nil {
			return nil, resp, fmt.Errorf("get package (attempt %d/%d): %w", i, maxAttempts, err)
		}
		if uploaded.HashType == "SHA3_512" && uploaded.HashValue != "" {
			break
		}
		if i < maxAttempts {
			time.Sleep(sleepBetween)
		}
	}

	if uploaded.HashType != "SHA3_512" || uploaded.HashValue == "" {
		return nil, resp, fmt.Errorf("timed out waiting for SHA3_512")
	}
	if uploaded.HashValue != initialHash {
		return nil, resp, fmt.Errorf("hash verification failed: initial=%s uploaded=%s", initialHash, uploaded.HashValue)
	}

	return created, resp, nil
}
