package bookmarks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the bookmarks-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: Undocumented
	Bookmarks struct {
		client client.Client
	}
)

func NewBookmarks(client client.Client) *Bookmarks {
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

	reqBuilder := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result)

	if rsqlQuery != nil {
		reqBuilder = reqBuilder.SetQueryParams(rsqlQuery)
	}

	resp, err := reqBuilder.Get(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

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

// DeleteByIDV1 removes the specified bookmark by ID.
// URL: DELETE /api/v1/bookmarks/{id}
// Jamf Pro API docs: Undocumented
func (s *Bookmarks) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBookmarksV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
