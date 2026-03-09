package advanced_user_content_searches

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// AdvancedUserContentSearchesServiceInterface defines the interface for advanced user content search operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches
	AdvancedUserContentSearchesServiceInterface interface {
		// ListV1 returns all advanced user content searches (Get Advanced User Content Searches).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified advanced user content search by ID (Get Advanced User Content Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches-id
		GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedUserContentSearch, *resty.Response, error)

		// CreateV1 creates a new advanced user content search (Create Advanced User Content Search).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-user-content-searches
		CreateV1(ctx context.Context, request *ResourceAdvancedUserContentSearch) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified advanced user content search by ID (Update Advanced User Content Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-user-content-searches-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceAdvancedUserContentSearch) (*ResourceAdvancedUserContentSearch, *resty.Response, error)

		// DeleteByIDV1 removes the specified advanced user content search by ID (Delete Advanced User Content Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-user-content-searches-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the advanced user content searches-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches
	AdvancedUserContentSearches struct {
		client transport.HTTPClient
	}
)

var _ AdvancedUserContentSearchesServiceInterface = (*AdvancedUserContentSearches)(nil)

func NewAdvancedUserContentSearches(client transport.HTTPClient) *AdvancedUserContentSearches {
	return &AdvancedUserContentSearches{client: client}
}

// ListV1 returns all advanced user content searches with automatic pagination.
// URL: GET /api/v1/advanced-user-content-searches
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches
func (s *AdvancedUserContentSearches) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProAdvancedUserContentSearchesV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var items []ResourceAdvancedUserContentSearch
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified advanced user content search by ID.
// URL: GET /api/v1/advanced-user-content-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches-id
func (s *AdvancedUserContentSearches) GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedUserContentSearch, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedUserContentSearchesV1, id)
	var result ResourceAdvancedUserContentSearch

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new advanced user content search.
// URL: POST /api/v1/advanced-user-content-searches
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-user-content-searches
func (s *AdvancedUserContentSearches) CreateV1(ctx context.Context, request *ResourceAdvancedUserContentSearch) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("search is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProAdvancedUserContentSearchesV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateByIDV1 updates the specified advanced user content search by ID.
// URL: PUT /api/v1/advanced-user-content-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-user-content-searches-id
func (s *AdvancedUserContentSearches) UpdateByIDV1(ctx context.Context, id string, request *ResourceAdvancedUserContentSearch) (*ResourceAdvancedUserContentSearch, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("search is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedUserContentSearchesV1, id)

	var result ResourceAdvancedUserContentSearch

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified advanced user content search by ID.
// URL: DELETE /api/v1/advanced-user-content-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-user-content-searches-id
func (s *AdvancedUserContentSearches) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedUserContentSearchesV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
