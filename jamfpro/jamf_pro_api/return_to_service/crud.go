package return_to_service

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Return to Service-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
	ReturnToService struct {
		client client.Client
	}
)

func NewReturnToService(client client.Client) *ReturnToService {
	return &ReturnToService{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Return to Service Operations
// -----------------------------------------------------------------------------

// ListV1 returns all Return to Service configurations.
// URL: GET /api/v1/return-to-service
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service
func (s *ReturnToService) ListV1(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProReturnToServiceV1
	var result ListResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the specified Return to Service configuration by ID.
// URL: GET /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-return-to-service-id
func (s *ReturnToService) GetByIDV1(ctx context.Context, id string) (*ResourceReturnToServiceConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := constants.EndpointJamfProReturnToServiceV1 + "/" + id
	var result ResourceReturnToServiceConfiguration

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV1 creates a new Return to Service configuration.
// URL: POST /api/v1/return-to-service
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-return-to-service
func (s *ReturnToService) CreateV1(ctx context.Context, request *ResourceReturnToServiceConfiguration) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	endpoint := constants.EndpointJamfProReturnToServiceV1
	var result CreateResponse

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

// UpdateByIDV1 updates the specified Return to Service configuration by ID.
// URL: PUT /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-return-to-service-id
func (s *ReturnToService) UpdateByIDV1(ctx context.Context, id string, request *ResourceReturnToServiceConfiguration) (*ResourceReturnToServiceConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	endpoint := constants.EndpointJamfProReturnToServiceV1 + "/" + id
	var result ResourceReturnToServiceConfiguration

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

// DeleteByIDV1 deletes the specified Return to Service configuration by ID.
// URL: DELETE /api/v1/return-to-service/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-return-to-service-id
func (s *ReturnToService) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := constants.EndpointJamfProReturnToServiceV1 + "/" + id

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
