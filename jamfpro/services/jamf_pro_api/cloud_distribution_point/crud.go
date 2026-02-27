package cloud_distribution_point

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CloudDistributionPointServiceInterface defines the interface for cloud distribution point operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
	CloudDistributionPointServiceInterface interface {
		// GetV1 returns the current cloud distribution point configuration (Get Cloud Distribution Point).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
		GetV1(ctx context.Context) (*ResourceCloudDistributionPointV1, *interfaces.Response, error)

		// CreateV1 provisions a cloud distribution point (Create Cloud Distribution Point).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point
		CreateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error)

		// UpdateV1 updates the cloud distribution point configuration (Update Cloud Distribution Point).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-cloud-distribution-point
		UpdateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error)

		// DeleteV1 removes the cloud distribution point configuration (Remove Cloud Distribution Point).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-cloud-distribution-point
		DeleteV1(ctx context.Context) (*interfaces.Response, error)

		// GetUploadCapabilityV1 returns the upload capability for the cloud distribution point.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-upload-capability
		GetUploadCapabilityV1(ctx context.Context) (*UploadCapabilityV1, *interfaces.Response, error)

		// GetTestConnectionV1 returns the test connection status for the cloud distribution point.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-test-connection
		GetTestConnectionV1(ctx context.Context) (*TestConnectionV1, *interfaces.Response, error)

		// GetHistoryV1 returns the history for the cloud distribution point.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// GetFilesV1 returns the inventory files for the cloud distribution point.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-files
		GetFilesV1(ctx context.Context, rsqlQuery map[string]string) (*FilesResponse, *interfaces.Response, error)

		// AddHistoryNoteV1 adds a history note for the cloud distribution point (Add Cloud Distribution Point History Note).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-history
		AddHistoryNoteV1(ctx context.Context, request *RequestHistoryNoteV1) (*HistoryItem, *interfaces.Response, error)

		// FailUploadV1 marks a specific file upload as failed for the cloud distribution point.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-fail-upload-id
		FailUploadV1(ctx context.Context, id string, fileName string, fileType string) (*interfaces.Response, error)

		// RefreshInventoryV1 updates inventory data for the cloud distribution point.
		//
		// Optional file-name query param: if provided, checks availability of that file; otherwise forces an immediate inventory refresh.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-refresh-inventory
		RefreshInventoryV1(ctx context.Context, fileName string) (*interfaces.Response, error)
	}

	// Service handles communication with the cloud distribution point methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CloudDistributionPointServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 returns the current cloud distribution point configuration.
// URL: GET /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point
func (s *Service) GetV1(ctx context.Context) (*ResourceCloudDistributionPointV1, *interfaces.Response, error) {
	var result ResourceCloudDistributionPointV1

	endpoint := EndpointCloudDistributionPointV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 provisions a cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point
func (s *Service) CreateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceCloudDistributionPointV1

	endpoint := EndpointCloudDistributionPointV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates the cloud distribution point (PATCH).
// URL: PATCH /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-cloud-distribution-point
func (s *Service) UpdateV1(ctx context.Context, request *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceCloudDistributionPointV1

	endpoint := EndpointCloudDistributionPointV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteV1 removes the cloud distribution point configuration.
// URL: DELETE /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-cloud-distribution-point
func (s *Service) DeleteV1(ctx context.Context) (*interfaces.Response, error) {
	endpoint := EndpointCloudDistributionPointV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetUploadCapabilityV1 returns upload capability for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/upload-capability
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-upload-capability
func (s *Service) GetUploadCapabilityV1(ctx context.Context) (*UploadCapabilityV1, *interfaces.Response, error) {
	endpoint := EndpointCloudDistributionPointV1 + "/upload-capability"
	var result UploadCapabilityV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTestConnectionV1 returns the test connection status.
// URL: GET /api/v1/cloud-distribution-point/test-connection
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-test-connection
func (s *Service) GetTestConnectionV1(ctx context.Context) (*TestConnectionV1, *interfaces.Response, error) {
	endpoint := EndpointCloudDistributionPointV1 + "/test-connection"
	var result TestConnectionV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV1 returns the history for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-history
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageResponse HistoryResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, EndpointCloudDistributionPointHistoryV1, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get cloud distribution point history: %w", err)
	}
	return &result, resp, nil
}

// GetFilesV1 returns the inventory files for the cloud distribution point.
// URL: GET /api/v1/cloud-distribution-point/files
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-distribution-point-files
func (s *Service) GetFilesV1(ctx context.Context, rsqlQuery map[string]string) (*FilesResponse, *interfaces.Response, error) {
	var result FilesResponse

	mergePage := func(pageData []byte) error {
		var pageResponse FilesResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, EndpointCloudDistributionPointFilesV1, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get cloud distribution point files: %w", err)
	}
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a history note for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-history
func (s *Service) AddHistoryNoteV1(ctx context.Context, request *RequestHistoryNoteV1) (*HistoryItem, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result HistoryItem

	endpoint := EndpointCloudDistributionPointHistoryV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// FailUploadV1 marks a specific file upload as failed for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/fail-upload/{id}
// Query params: file-name, type (PACKAGE, EBOOK, MOBILE_DEVICE_APP).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-fail-upload-id
func (s *Service) FailUploadV1(ctx context.Context, id string, fileName string, fileType string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if fileName == "" {
		return nil, fmt.Errorf("fileName is required")
	}
	if fileType == "" {
		return nil, fmt.Errorf("fileType is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudDistributionPointFailUploadV1, id)

	queryParams := map[string]string{
		"file-name": fileName,
		"type":      fileType,
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, queryParams, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RefreshInventoryV1 updates inventory data for the cloud distribution point.
// URL: POST /api/v1/cloud-distribution-point/refresh-inventory
// Optional query param: file-name (if provided, checks availability of that file; otherwise forces immediate inventory refresh).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point-refresh-inventory
func (s *Service) RefreshInventoryV1(ctx context.Context, fileName string) (*interfaces.Response, error) {
	endpoint := EndpointCloudDistributionPointRefreshV1

	var queryParams map[string]string
	if fileName != "" {
		queryParams = map[string]string{"file-name": fileName}
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, queryParams, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
