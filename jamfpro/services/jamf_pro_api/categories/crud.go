package categories

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CategoriesServiceInterface defines the interface for category operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
	CategoriesServiceInterface interface {
		// ListV1 returns all category objects (Get Category objects).
		//
		// Returns a paged list of category objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified category by ID (Get specified Category object).
		//
		// Returns a single category object for the given ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
		GetByIDV1(ctx context.Context, id string) (*ResourceCategory, *interfaces.Response, error)

		// CreateV1 creates a new category record (Create Category record).
		//
		// Creates a new category. The request body must include name; priority is optional.
		// Returns the created category ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories
		CreateV1(ctx context.Context, request *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified category by ID (Update specified Category object).
		//
		// Updates an existing category. All updatable fields (name, priority) may be sent.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-categories-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified category by ID (Remove specified Category record).
		//
		// Permanently deletes the category. This operation cannot be undone.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-categories-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteCategoriesByIDV1 deletes multiple categories by their IDs (Delete multiple Categories by their IDs).
		//
		// Sends a POST to /api/v1/categories/delete-multiple with a body containing category IDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
		DeleteCategoriesByIDV1(ctx context.Context, req *DeleteCategoriesByIDRequest) (*interfaces.Response, error)

		// GetCategoryHistoryV1 returns the history object for the specified category (Get specified Category history object).
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id-history
		GetCategoryHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*CategoryHistoryResponse, *interfaces.Response, error)

		// AddCategoryHistoryNotesV1 adds notes to the specified category history (Add specified Category history object notes).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories-id-history
		AddCategoryHistoryNotesV1(ctx context.Context, id string, req *AddCategoryHistoryNotesRequest) (*interfaces.Response, error)
	}

	// Service handles communication with the categories
	// related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CategoriesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Categories CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all category objects (Get Category objects).
// URL: GET /api/v1/categories
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointCategoriesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified category by ID (Get specified Category object).
// URL: GET /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceCategory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

	var result ResourceCategory

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

// CreateV1 creates a new category record (Create Category record).
// URL: POST /api/v1/categories
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories
func (s *Service) CreateV1(ctx context.Context, request *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := EndpointCategoriesV1

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

// UpdateByIDV1 updates the specified category by ID (Update specified Category object).
// URL: PUT /api/v1/categories/{id}
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/put_v1-categories-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

	var result CreateUpdateResponse

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

// DeleteByIDV1 removes the specified category by ID (Remove specified Category record).
// URL: DELETE /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-categories-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

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

// DeleteCategoriesByIDV1 deletes multiple categories by their IDs (Delete multiple Categories by their IDs).
// URL: POST /api/v1/categories/delete-multiple
// Body: JSON with ids (array of category IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
func (s *Service) DeleteCategoriesByIDV1(ctx context.Context, req *DeleteCategoriesByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointCategoriesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetCategoryHistoryV1 returns the history object for the specified category (Get specified Category history object).
// URL: GET /api/v1/categories/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id-history
func (s *Service) GetCategoryHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*CategoryHistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCategoriesV1, id)

	var result CategoryHistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddCategoryHistoryNotesV1 adds notes to the specified category history (Add specified Category history object notes).
// URL: POST /api/v1/categories/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-id-history
func (s *Service) AddCategoryHistoryNotesV1(ctx context.Context, id string, req *AddCategoryHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
