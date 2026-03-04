package bookmarks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// BookmarksServiceInterface defines the interface for bookmark operations.
	//
	// Jamf Pro API docs: Undocumented
	BookmarksServiceInterface interface {
		// ListV1 returns all bookmarks. Optional rsqlQuery: filter (RSQL), sort, page, page-size (Get Bookmarks).
		//
		// Jamf Pro API docs: Undocumented
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified bookmark by ID (Get Bookmark by ID).
		//
		// Jamf Pro API docs: Undocumented-id
		GetByIDV1(ctx context.Context, id string) (*ResourceBookmark, *resty.Response, error)

		// CreateV1 creates a new bookmark (Create Bookmark).
		//
		// Jamf Pro API docs: Undocumented
		CreateV1(ctx context.Context, request *ResourceBookmark) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified bookmark by ID (Update Bookmark by ID).
		//
		// Jamf Pro API docs: Undocumented
		UpdateByIDV1(ctx context.Context, id string, request *ResourceBookmark) (*ResourceBookmark, *resty.Response, error)

		// DeleteByIDV1 removes the specified bookmark by ID (Delete Bookmark by ID).
		//
		// Jamf Pro API docs: Undocumented
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the bookmarks-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: Undocumented
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ BookmarksServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Bookmarks Operations
// -----------------------------------------------------------------------------

// ListV1 returns all bookmarks. Optional rsqlQuery: filter (RSQL), sort, page, page-size.
// URL: GET /api/v1/bookmarks
// Jamf Pro API docs: Undocumented
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointBookmarksV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified bookmark by ID.
// URL: GET /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceBookmark, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)
	var result ResourceBookmark

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new bookmark.
// URL: POST /api/v1/bookmarks
// Jamf Pro API docs: Undocumented
func (s *Service) CreateV1(ctx context.Context, request *ResourceBookmark) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointBookmarksV1

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

// UpdateByIDV1 updates the specified bookmark by ID.
// URL: PUT /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *ResourceBookmark) (*ResourceBookmark, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)

	var result ResourceBookmark

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

// DeleteByIDV1 removes the specified bookmark by ID.
// URL: DELETE /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
