package advanced_user_searches

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AdvancedUserSearchesServiceInterface defines the interface for Classic API advanced user search operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
	AdvancedUserSearchesServiceInterface interface {
		// ListAdvancedUserSearches returns all advanced user searches.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
		ListAdvancedUserSearches(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetAdvancedUserSearchByID returns the specified advanced user search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyid
		GetAdvancedUserSearchByID(ctx context.Context, id int) (*ResourceAdvancedUserSearch, *interfaces.Response, error)

		// GetAdvancedUserSearchByName returns the specified advanced user search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyname
		GetAdvancedUserSearchByName(ctx context.Context, name string) (*ResourceAdvancedUserSearch, *interfaces.Response, error)

		// CreateAdvancedUserSearch creates a new advanced user search.
		//
		// Returns the created advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createadvancedusersearchbyid
		CreateAdvancedUserSearch(ctx context.Context, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateAdvancedUserSearchByID updates the specified advanced user search by ID.
		//
		// Returns the updated advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyid
		UpdateAdvancedUserSearchByID(ctx context.Context, id int, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateAdvancedUserSearchByName updates the specified advanced user search by name.
		//
		// Returns the updated advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyname
		UpdateAdvancedUserSearchByName(ctx context.Context, name string, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteAdvancedUserSearchByID removes the specified advanced user search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyid
		DeleteAdvancedUserSearchByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteAdvancedUserSearchByName removes the specified advanced user search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyname
		DeleteAdvancedUserSearchByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the advanced user searches-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AdvancedUserSearchesServiceInterface = (*Service)(nil)

// NewService returns a new advanced user searches Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Advanced User Searches CRUD Operations
// -----------------------------------------------------------------------------

// ListAdvancedUserSearches returns all advanced user searches.
// URL: GET /JSSResource/advancedusersearches
// https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
func (s *Service) ListAdvancedUserSearches(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicAdvancedUserSearches

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAdvancedUserSearchByID returns the specified advanced user search by ID.
// URL: GET /JSSResource/advancedusersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyid
func (s *Service) GetAdvancedUserSearchByID(ctx context.Context, id int) (*ResourceAdvancedUserSearch, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("advanced user search ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedUserSearches, id)

	var result ResourceAdvancedUserSearch

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAdvancedUserSearchByName returns the specified advanced user search by name.
// URL: GET /JSSResource/advancedusersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyname
func (s *Service) GetAdvancedUserSearchByName(ctx context.Context, name string) (*ResourceAdvancedUserSearch, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("advanced user search name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedUserSearches, name)

	var result ResourceAdvancedUserSearch

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateAdvancedUserSearch creates a new advanced user search.
// URL: POST /JSSResource/advancedusersearches/id/0
// Returns the created advanced user search ID only.
// https://developer.jamf.com/jamf-pro/reference/createadvancedusersearchbyid
func (s *Service) CreateAdvancedUserSearch(ctx context.Context, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicAdvancedUserSearches)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateAdvancedUserSearchByID updates the specified advanced user search by ID.
// URL: PUT /JSSResource/advancedusersearches/id/{id}
// Returns the updated advanced user search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyid
func (s *Service) UpdateAdvancedUserSearchByID(ctx context.Context, id int, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("advanced user search ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedUserSearches, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateAdvancedUserSearchByName updates the specified advanced user search by name.
// URL: PUT /JSSResource/advancedusersearches/name/{name}
// Returns the updated advanced user search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyname
func (s *Service) UpdateAdvancedUserSearchByName(ctx context.Context, name string, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("advanced user search name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedUserSearches, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteAdvancedUserSearchByID removes the specified advanced user search by ID.
// URL: DELETE /JSSResource/advancedusersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyid
func (s *Service) DeleteAdvancedUserSearchByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("advanced user search ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedUserSearches, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteAdvancedUserSearchByName removes the specified advanced user search by name.
// URL: DELETE /JSSResource/advancedusersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyname
func (s *Service) DeleteAdvancedUserSearchByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("advanced user search name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedUserSearches, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
