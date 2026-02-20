package return_to_service

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ReturnToServiceServiceInterface defines the interface for Return to Service configuration operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
	ReturnToServiceServiceInterface interface {
		// ListV1 returns all Return to Service configurations (Get Return to Service configurations).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
		ListV1(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified Return to Service configuration by ID (Get specified Return to Service configuration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service-id
		GetByIDV1(ctx context.Context, id string) (*ResourceReturnToServiceConfiguration, *interfaces.Response, error)

		// CreateV1 creates a new Return to Service configuration (Create Return to Service configuration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-return-to-service
		CreateV1(ctx context.Context, request *ResourceReturnToServiceConfiguration) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified Return to Service configuration by ID (Update specified Return to Service configuration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-return-to-service-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceReturnToServiceConfiguration) (*ResourceReturnToServiceConfiguration, *interfaces.Response, error)

		// DeleteByIDV1 deletes the specified Return to Service configuration by ID (Delete specified Return to Service configuration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-return-to-service-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the Return to Service-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ReturnToServiceServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Return to Service Operations
// -----------------------------------------------------------------------------

// ListV1 returns all Return to Service configurations.
// URL: GET /api/v1/return-to-service
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
func (s *Service) ListV1(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointReturnToServiceV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the specified Return to Service configuration by ID.
// URL: GET /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceReturnToServiceConfiguration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := EndpointReturnToServiceV1 + "/" + id
	var result ResourceReturnToServiceConfiguration
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV1 creates a new Return to Service configuration.
// URL: POST /api/v1/return-to-service
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-return-to-service
func (s *Service) CreateV1(ctx context.Context, request *ResourceReturnToServiceConfiguration) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result CreateResponse
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, EndpointReturnToServiceV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateByIDV1 updates the specified Return to Service configuration by ID.
// URL: PUT /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-return-to-service-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *ResourceReturnToServiceConfiguration) (*ResourceReturnToServiceConfiguration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	endpoint := EndpointReturnToServiceV1 + "/" + id
	var result ResourceReturnToServiceConfiguration
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteByIDV1 deletes the specified Return to Service configuration by ID.
// URL: DELETE /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-return-to-service-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := EndpointReturnToServiceV1 + "/" + id
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
