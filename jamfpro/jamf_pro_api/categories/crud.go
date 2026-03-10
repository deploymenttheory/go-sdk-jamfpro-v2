package categories

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the categories
	// related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
	Categories struct {
		client client.Client
	}
)

func NewCategories(client client.Client) *Categories {
	return &Categories{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Categories CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all category objects (Get Category objects).
// URL: GET /api/v1/categories
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories
func (s *Categories) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProCategoriesV1

	mergePage := func(pageData []byte) error {
		var items []ResourceCategory
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
		return nil, resp, fmt.Errorf("failed to list categories: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified category by ID (Get specified Category object).
// URL: GET /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
func (s *Categories) GetByIDV1(ctx context.Context, id string) (*ResourceCategory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCategoriesV1, id)

	var result ResourceCategory

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new category record (Create Category record).
// URL: POST /api/v1/categories
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories
func (s *Categories) CreateV1(ctx context.Context, request *RequestCategory) (*CreateUpdateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := constants.EndpointJamfProCategoriesV1

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

// UpdateByIDV1 updates the specified category by ID (Update specified Category object).
// URL: PUT /api/v1/categories/{id}
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/put_v1-categories-id
func (s *Categories) UpdateByIDV1(ctx context.Context, id string, request *RequestCategory) (*CreateUpdateResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCategoriesV1, id)

	var result CreateUpdateResponse

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

// DeleteByIDV1 removes the specified category by ID (Remove specified Category record).
// URL: DELETE /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-categories-id
func (s *Categories) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCategoriesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteCategoriesByIDV1 deletes multiple categories by their IDs (Delete multiple Categories by their IDs).
// URL: POST /api/v1/categories/delete-multiple
// Body: JSON with ids (array of category IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
func (s *Categories) DeleteCategoriesByIDV1(ctx context.Context, req *DeleteCategoriesByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProCategoriesV1 + "/delete-multiple"

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetCategoryHistoryV1 returns the history object for the specified category (Get specified Category history object).
// URL: GET /api/v1/categories/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id-history
func (s *Categories) GetCategoryHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*CategoryHistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProCategoriesV1, id)

	var result CategoryHistoryResponse

	mergePage := func(pageData []byte) error {
		var items []HistoryObject
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
		return nil, resp, fmt.Errorf("failed to get category history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddCategoryHistoryNotesV1 adds notes to the specified category history (Add specified Category history object notes).
// URL: POST /api/v1/categories/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-id-history
func (s *Categories) AddCategoryHistoryNotesV1(ctx context.Context, id string, req *AddCategoryHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProCategoriesV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
