package bookmarks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the bookmarks-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: Undocumented
	Bookmarks struct {
		client transport.HTTPClient
	}
)

func NewBookmarks(client transport.HTTPClient) *Bookmarks {
	return &Bookmarks{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Bookmarks Operations
// -----------------------------------------------------------------------------

// ListV1 returns all bookmarks. Optional rsqlQuery: filter (RSQL), sort, page, page-size.
// URL: GET /api/v1/bookmarks
// Jamf Pro API docs: Undocumented
func (s *Bookmarks) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProBookmarksV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Bookmarks) GetByIDV1(ctx context.Context, id string) (*ResourceBookmark, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBookmarksV1, id)
	var result ResourceBookmark

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Bookmarks) CreateV1(ctx context.Context, request *ResourceBookmark) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProBookmarksV1

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

// UpdateByIDV1 updates the specified bookmark by ID.
// URL: PUT /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented
func (s *Bookmarks) UpdateByIDV1(ctx context.Context, id string, request *ResourceBookmark) (*ResourceBookmark, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBookmarksV1, id)

	var result ResourceBookmark

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

// DeleteByIDV1 removes the specified bookmark by ID.
// URL: DELETE /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented
func (s *Bookmarks) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBookmarksV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
