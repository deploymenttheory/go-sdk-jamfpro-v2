package cloud_distribution_point

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the cloud distribution point methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
	CloudDistributionPoint struct {
		client client.Client
	}
)

func NewCloudDistributionPoint(client client.Client) *CloudDistributionPoint {
	return &CloudDistributionPoint{client: client}
}

// GetV1 returns the current cloud distribution point configuration.
// URL: GET /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
func (s *CloudDistributionPoint) GetV1(ctx context.Context) (*ResourceCloudDistributionPointV1, *resty.Response, error) {
	var result ResourceCloudDistributionPointV1

	endpoint := constants.EndpointJamfProCloudDistributionPointV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 provisions a cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point
func (s *CloudDistributionPoint) CreateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if _, ok := validCdnTypes[request.CdnType]; !ok {
		return nil, nil, fmt.Errorf("invalid cdnType %q: must be one of NONE, JAMF_CLOUD, RACKSPACE_CLOUD_FILES, AMAZON_S3, AKAMAI", request.CdnType)
	}

	var result ResourceCloudDistributionPointV1

	endpoint := constants.EndpointJamfProCloudDistributionPointV1

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

// UpdateV1 updates the cloud distribution point (PATCH).
// URL: PATCH /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-cloud-distribution-point
func (s *CloudDistributionPoint) UpdateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if _, ok := validCdnTypes[request.CdnType]; !ok {
		return nil, nil, fmt.Errorf("invalid cdnType %q: must be one of NONE, JAMF_CLOUD, RACKSPACE_CLOUD_FILES, AMAZON_S3, AKAMAI", request.CdnType)
	}

	var result ResourceCloudDistributionPointV1

	endpoint := constants.EndpointJamfProCloudDistributionPointV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Patch(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteV1 removes the cloud distribution point configuration.
// URL: DELETE /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-cloud-distribution-point
func (s *CloudDistributionPoint) DeleteV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProCloudDistributionPointV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetUploadCapabilityV1 returns upload capability for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/upload-capability
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-upload-capability
func (s *CloudDistributionPoint) GetUploadCapabilityV1(ctx context.Context) (*UploadCapabilityV1, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCloudDistributionPointV1 + "/upload-capability"
	var result UploadCapabilityV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTestConnectionV1 returns the test connection status.
// URL: GET /api/v1/cloud-distribution-point/test-connection
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-test-connection
func (s *CloudDistributionPoint) GetTestConnectionV1(ctx context.Context) (*TestConnectionV1, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCloudDistributionPointV1 + "/test-connection"
	var result TestConnectionV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV1 returns the history for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-history
func (s *CloudDistributionPoint) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProCloudDistributionPointHistoryV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

	if err != nil {
		return nil, resp, fmt.Errorf("failed to get cloud distribution point history: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetFilesV1 returns the inventory files for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/files
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-files
func (s *CloudDistributionPoint) GetFilesV1(ctx context.Context, rsqlQuery map[string]string) (*FilesResponse, *resty.Response, error) {
	var result FilesResponse

	mergePage := func(pageData []byte) error {
		var pageItems []FileItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProCloudDistributionPointFilesV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

	if err != nil {
		return nil, resp, fmt.Errorf("failed to get cloud distribution point files: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a history note for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-history
func (s *CloudDistributionPoint) AddHistoryNoteV1(ctx context.Context, request *RequestHistoryNoteV1) (*HistoryItem, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result HistoryItem

	endpoint := constants.EndpointJamfProCloudDistributionPointHistoryV1

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

// FailUploadV1 marks a specific file upload as failed for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/fail-upload/{id}
// Query params: file-name, type (PACKAGE, EBOOK, MOBILE_DEVICE_APP).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-fail-upload-id
func (s *CloudDistributionPoint) FailUploadV1(ctx context.Context, id string, fileName string, fileType string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if fileName == "" {
		return nil, fmt.Errorf("fileName is required")
	}
	if fileType == "" {
		return nil, fmt.Errorf("fileType is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCloudDistributionPointFailUploadV1, id)

	queryParams := map[string]string{
		"file-name": fileName,
		"type":      fileType,
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(queryParams).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RefreshInventoryV1 updates inventory data for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/refresh-inventory
// Optional query param: file-name (if provided, checks availability of that file; otherwise forces immediate inventory refresh).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-refresh-inventory
func (s *CloudDistributionPoint) RefreshInventoryV1(ctx context.Context, fileName string) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProCloudDistributionPointRefreshV1

	var queryParams map[string]string
	if fileName != "" {
		queryParams = map[string]string{"file-name": fileName}
	}

	reqBuilder := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON)

	if queryParams != nil {
		reqBuilder = reqBuilder.SetQueryParams(queryParams)
	}

	resp, err := reqBuilder.Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
