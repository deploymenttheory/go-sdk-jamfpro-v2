package ebooks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// EbooksServiceInterface defines the interface for ebook operations (Jamf Pro API v1).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks
	EbooksServiceInterface interface {
		// ListV1 returns all ebook objects (Get Ebook objects).
		//
		// Returns a paged list of ebook objects. Optional query parameters support
		// pagination and sorting (page, page-size, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified ebook by ID (Get specified Ebook object).
		//
		// Returns a single ebook object for the given ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks-id
		GetByIDV1(ctx context.Context, id string) (*ResourceEbook, *resty.Response, error)

		// GetScopeByIDV1 returns the scope for the specified ebook by ID (Get specified scope of Ebook object).
		//
		// Returns scope with assignments, limitations, and exclusions.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks-id-scope
		GetScopeByIDV1(ctx context.Context, id string) (*ResourceScope, *resty.Response, error)
	}

	// Service handles communication with the ebooks-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks
	Ebooks struct {
		client interfaces.HTTPClient
	}
)

var _ EbooksServiceInterface = (*Ebooks)(nil)

func NewEbooks(client interfaces.HTTPClient) *Ebooks {
	return &Ebooks{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Ebooks CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all ebook objects (Get Ebook objects).
// URL: GET /api/v1/ebooks

// https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks
func (s *Ebooks) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointEbooksV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEbook
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list ebooks: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified ebook by ID (Get specified Ebook object).
// URL: GET /api/v1/ebooks/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks-id
func (s *Ebooks) GetByIDV1(ctx context.Context, id string) (*ResourceEbook, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("ebook ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointEbooksV1, id)

	var result ResourceEbook

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetScopeByIDV1 returns the scope for the specified ebook by ID (Get specified scope of Ebook object).
// URL: GET /api/v1/ebooks/{id}/scope
// https://developer.jamf.com/jamf-pro/reference/get_v1-ebooks-id-scope
func (s *Ebooks) GetScopeByIDV1(ctx context.Context, id string) (*ResourceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("ebook ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointEbooksV1, id)

	var result ResourceScope

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
