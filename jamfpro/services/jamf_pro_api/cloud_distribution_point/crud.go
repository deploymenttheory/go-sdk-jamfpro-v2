package cloud_distribution_point

import (
	"context"
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
		CreateV1(ctx context.Context, req *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error)

		// UpdateV1 updates the cloud distribution point configuration (Update Cloud Distribution Point).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-cloud-distribution-point
		UpdateV1(ctx context.Context, req *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error)

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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) CreateV1(ctx context.Context, req *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceCloudDistributionPointV1

	endpoint := EndpointCloudDistributionPointV1

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

// UpdateV1 updates the cloud distribution point (PATCH).
// URL: PATCH /api/v1/cloud-distribution-point
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-cloud-distribution-point
func (s *Service) UpdateV1(ctx context.Context, req *RequestCloudDistributionPointV1) (*ResourceCloudDistributionPointV1, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceCloudDistributionPointV1

	endpoint := EndpointCloudDistributionPointV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, req, headers, &result)
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
