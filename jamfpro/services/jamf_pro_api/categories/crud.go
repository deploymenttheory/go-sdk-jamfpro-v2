package categories

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
)

type (
	// CategoriesServiceInterface defines the interface for category operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
	CategoriesServiceInterface interface {
		// ListCategories returns all category objects (Get Category objects).
		//
		// Returns a paged list of category objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
		ListCategories(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetCategoryByID returns the specified category by ID (Get specified Category object).
		//
		// Returns a single category object for the given ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
		GetCategoryByID(ctx context.Context, id string) (*ResourceCategory, *interfaces.Response, error)

		// CreateCategory creates a new category record (Create Category record).
		//
		// Creates a new category. The request body must include name; priority is optional.
		// Returns the created category ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories
		CreateCategory(ctx context.Context, req *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateCategoryByID updates the specified category by ID (Update specified Category object).
		//
		// Updates an existing category. All updatable fields (name, priority) may be sent.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-categories-id
		UpdateCategoryByID(ctx context.Context, id string, req *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteCategoryByID removes the specified category by ID (Remove specified Category record).
		//
		// Permanently deletes the category. This operation cannot be undone.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-categories-id
		DeleteCategoryByID(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteCategoriesByID deletes multiple categories by their IDs (Delete multiple Categories by their IDs).
		//
		// Sends a POST to /api/v1/categories/delete-multiple with a body containing category IDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
		DeleteCategoriesByID(ctx context.Context, req *DeleteCategoriesByIDRequest) (*interfaces.Response, error)

		// GetCategoryHistory returns the history object for the specified category (Get specified Category history object).
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id-history
		GetCategoryHistory(ctx context.Context, id string, rsqlQuery map[string]string) (*CategoryHistoryResponse, *interfaces.Response, error)

		// AddCategoryHistoryNotes adds notes to the specified category history (Add specified Category history object notes).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-categories-id-history
		AddCategoryHistoryNotes(ctx context.Context, id string, req *AddCategoryHistoryNotesRequest) (*interfaces.Response, error)
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

// ListCategories returns all category objects (Get Category objects).
// URL: GET /api/v1/categories
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories
func (s *Service) ListCategories(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointCategoriesV1

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ListResponse

	resp, err := s.client.Get(ctx, endpoint, queryParams, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCategoryByID returns the specified category by ID (Get specified Category object).
// URL: GET /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id
func (s *Service) GetCategoryByID(ctx context.Context, id string) (*ResourceCategory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ResourceCategory

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateCategory creates a new category record (Create Category record).
// URL: POST /api/v1/categories
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories
func (s *Service) CreateCategory(ctx context.Context, req *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointCategoriesV1

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result CreateUpdateResponse

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateCategoryByID updates the specified category by ID (Update specified Category object).
// URL: PUT /api/v1/categories/{id}
// Body: JSON with name, priority (optional)
// https://developer.jamf.com/jamf-pro/reference/put_v1-categories-id
func (s *Service) UpdateCategoryByID(ctx context.Context, id string, req *RequestCategory) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteCategoryByID removes the specified category by ID (Remove specified Category record).
// URL: DELETE /api/v1/categories/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-categories-id
func (s *Service) DeleteCategoryByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteCategoriesByID deletes multiple categories by their IDs (Delete multiple Categories by their IDs).
// URL: POST /api/v1/categories/delete-multiple
// Body: JSON with ids (array of category IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-delete-multiple
func (s *Service) DeleteCategoriesByID(ctx context.Context, req *DeleteCategoriesByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointCategoriesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetCategoryHistory returns the history object for the specified category (Get specified Category history object).
// URL: GET /api/v1/categories/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-categories-id-history
func (s *Service) GetCategoryHistory(ctx context.Context, id string, rsqlQuery map[string]string) (*CategoryHistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("category ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result CategoryHistoryResponse

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddCategoryHistoryNotes adds notes to the specified category history (Add specified Category history object notes).
// URL: POST /api/v1/categories/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-categories-id-history
func (s *Service) AddCategoryHistoryNotes(ctx context.Context, id string, req *AddCategoryHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("category ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCategoriesV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
