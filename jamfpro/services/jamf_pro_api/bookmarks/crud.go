package bookmarks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// BookmarksServiceInterface defines the interface for bookmark operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks
	BookmarksServiceInterface interface {
		// ListV1 returns all bookmarks with optional query params (Get Bookmarks).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks
		ListV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified bookmark by ID (Get Bookmark by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks-id
		GetByIDV1(ctx context.Context, id string) (*ResourceBookmark, *interfaces.Response, error)

		// CreateV1 creates a new bookmark (Create Bookmark).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-bookmarks
		CreateV1(ctx context.Context, bookmark *ResourceBookmark) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified bookmark by ID (Update Bookmark by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-bookmarks-id
		UpdateByIDV1(ctx context.Context, id string, bookmark *ResourceBookmark) (*ResourceBookmark, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified bookmark by ID (Delete Bookmark by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-bookmarks-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the bookmarks endpoint.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ BookmarksServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all bookmarks.
// URL: GET /api/v1/bookmarks
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks
func (s *Service) ListV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	resp, err := s.client.Get(ctx, EndpointBookmarksV1, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the specified bookmark by ID.
// URL: GET /api/v1/bookmarks/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceBookmark, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)
	var result ResourceBookmark
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV1 creates a new bookmark.
// URL: POST /api/v1/bookmarks
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-bookmarks
func (s *Service) CreateV1(ctx context.Context, bookmark *ResourceBookmark) (*CreateResponse, *interfaces.Response, error) {
	if bookmark == nil {
		return nil, nil, fmt.Errorf("bookmark is required")
	}
	var result CreateResponse
	resp, err := s.client.Post(ctx, EndpointBookmarksV1, bookmark, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateByIDV1 updates the specified bookmark by ID.
// URL: PUT /api/v1/bookmarks/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-bookmarks-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, bookmark *ResourceBookmark) (*ResourceBookmark, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if bookmark == nil {
		return nil, nil, fmt.Errorf("bookmark is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)
	var result ResourceBookmark
	resp, err := s.client.Put(ctx, endpoint, bookmark, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteByIDV1 removes the specified bookmark by ID.
// URL: DELETE /api/v1/bookmarks/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-bookmarks-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointBookmarksV1, id)
	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
