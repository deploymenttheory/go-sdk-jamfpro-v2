package api_integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// ApiIntegrationsServiceInterface defines the interface for API integrations operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-integrations
	ApiIntegrationsServiceInterface interface {
		// ListV1 returns a page of API integrations (Get API Integrations / Get with Search Criteria).
		//
		// Query params (optional, pass via rsqlQuery): page (default 0), page-size (default 100),
		// sort (e.g. "id:asc", "displayName:desc"), filter (RSQL, e.g. displayName=="IntegrationName").
		// Allowed sort/filter fields: id, displayName.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-integrations
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the API integration by ID (Get API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getoneapiintegration
		GetByIDV1(ctx context.Context, id string) (*ResourceApiIntegration, *interfaces.Response, error)

		// GetByNameV1 returns the API integration by display name (searches first page of ListV1).
		GetByNameV1(ctx context.Context, name string) (*ResourceApiIntegration, *interfaces.Response, error)

		// CreateV1 creates a new API integration (Create API Integration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/postcreateapiintegration
		CreateV1(ctx context.Context, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// UpdateByIDV1 updates the API integration by ID (Update API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putupdateapiintegration
		UpdateByIDV1(ctx context.Context, id string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// UpdateByNameV1 updates the API integration by display name.
		UpdateByNameV1(ctx context.Context, name string, request *ResourceApiIntegration) (*ResourceApiIntegration, *interfaces.Response, error)

		// DeleteByIDV1 deletes the API integration by ID (Delete API Integration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/deleteapiintegration
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV1 deletes the API integration by display name.
		DeleteByNameV1(ctx context.Context, name string) (*interfaces.Response, error)

		// RefreshClientCredentialsByIDV1 creates client credentials for the API integration by ID (Create client credentials).
		// POST /api/v1/api-integrations/{id}/client-credentials
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/postcreateclientcredentials
		RefreshClientCredentialsByIDV1(ctx context.Context, id string) (*ResourceClientCredentials, *interfaces.Response, error)
	}

	// Service handles communication with the API integrations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-integrations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ApiIntegrationsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all API integrations with automatic pagination.
// URL: GET /api/v1/api-integrations
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/get_v1-api-integrations
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var integration ResourceApiIntegration
				if err := mapstructure.Decode(item, &integration); err != nil {
					return fmt.Errorf("failed to decode api integration: %w", err)
				}
				result.Results = append(result.Results, integration)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, EndpointApiIntegrationsV1, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the API integration by ID.
// URL: GET /api/v1/api-integrations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getoneapiintegration
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
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/postcreateapiintegration
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
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putupdateapiintegration
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
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/deleteapiintegration
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
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/postcreateclientcredentials
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
