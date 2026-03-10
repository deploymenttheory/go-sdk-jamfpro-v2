package advanced_user_content_searches

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the advanced user content searches-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-user-content-searches
	AdvancedUserContentSearches struct {
		client client.Client
	}
)

func NewAdvancedUserContentSearches(client client.Client) *AdvancedUserContentSearches {
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

	mergePage := func(pageData []byte) error {
		var items []ResourceAdvancedUserContentSearch
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

// DeleteByIDV1 removes the specified advanced user content search by ID.
// URL: DELETE /api/v1/advanced-user-content-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-user-content-searches-id
func (s *AdvancedUserContentSearches) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedUserContentSearchesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
