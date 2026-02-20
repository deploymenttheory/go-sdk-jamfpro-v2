package api_integrations

import (
	"context"
	"fmt"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ApiIntegrationsServiceInterface defines the interface for API integrations operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
	ApiIntegrationsServiceInterface interface {
		// ListV1 returns a page of API integrations (Get API Integrations).
		//
		// Optional rsqlQuery keys: page, page-size, sort.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the API integration by ID (Get API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getapiintegrationbyid
		GetByIDV1(ctx context.Context, id string) (*ResourceApiIntegration, *interfaces.Response, error)

		// GetByNameV1 returns the API integration by display name (searches first page of ListV1).
		GetByNameV1(ctx context.Context, name string) (*ResourceApiIntegration, *interfaces.Response, error)

		// CreateV1 creates a new API integration (Create API Integration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/createapiintegration
		CreateV1(ctx context.Context, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// UpdateByIDV1 updates the API integration by ID (Update API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/updateapiintegrationbyid
		UpdateByIDV1(ctx context.Context, id string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// UpdateByNameV1 updates the API integration by display name.
		UpdateByNameV1(ctx context.Context, name string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// DeleteByIDV1 deletes the API integration by ID (Delete API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/deleteapiintegrationbyid
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV1 deletes the API integration by display name.
		DeleteByNameV1(ctx context.Context, name string) (*interfaces.Response, error)

		// RefreshClientCredentialsByIDV1 creates new client credentials for the API integration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/refreshclientcredentials
		RefreshClientCredentialsByIDV1(ctx context.Context, id string) (*ResourceClientCredentials, *interfaces.Response, error)
	}

	// Service handles communication with the API integrations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ApiIntegrationsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns a page of API integrations.
// URL: GET /api/v1/api-integrations
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getallapiintegrations
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointApiIntegrationsV1, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the API integration by ID.
// URL: GET /api/v1/api-integrations/{id}
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceApiIntegration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointApiIntegrationsV1, id)

	var result ResourceApiIntegration

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

// GetByNameV1 returns the API integration by display name (searches first page).
func (s *Service) GetByNameV1(ctx context.Context, name string) (*ResourceApiIntegration, *interfaces.Response, error) {
	list, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}
	for i := range list.Results {
		if list.Results[i].DisplayName == name {
			return &list.Results[i], resp, nil
		}
	}
	return nil, resp, fmt.Errorf("api integration with name %q not found", name)
}

// CreateV1 creates a new API integration.
// URL: POST /api/v1/api-integrations
func (s *Service) CreateV1(ctx context.Context, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceApiIntegration

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointApiIntegrationsV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the API integration by ID.
// URL: PUT /api/v1/api-integrations/{id}
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointApiIntegrationsV1, id)
	var result ResourceApiIntegration

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateByNameV1 updates the API integration by display name.
func (s *Service) UpdateByNameV1(ctx context.Context, name string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error) {
	existing, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return nil, resp, err
	}
	idStr := strconv.Itoa(existing.ID)
	return s.UpdateByIDV1(ctx, idStr, request)
}

// DeleteByIDV1 deletes the API integration by ID.
// URL: DELETE /api/v1/api-integrations/{id}
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointApiIntegrationsV1, id)

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

// DeleteByNameV1 deletes the API integration by display name.
func (s *Service) DeleteByNameV1(ctx context.Context, name string) (*interfaces.Response, error) {
	existing, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return resp, err
	}
	idStr := strconv.Itoa(existing.ID)
	return s.DeleteByIDV1(ctx, idStr)
}

// RefreshClientCredentialsByIDV1 creates new client credentials for the API integration by ID.
// URL: POST /api/v1/api-integrations/{id}/client-credentials
func (s *Service) RefreshClientCredentialsByIDV1(ctx context.Context, id string) (*ResourceClientCredentials, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s/client-credentials", EndpointApiIntegrationsV1, id)

	var result ResourceClientCredentials

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
