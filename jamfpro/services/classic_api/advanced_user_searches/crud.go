package advanced_user_searches

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// AdvancedUserSearchesServiceInterface defines the interface for Classic API advanced user search operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
	AdvancedUserSearchesServiceInterface interface {
		// List returns all advanced user searches.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified advanced user search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyid
		GetByID(ctx context.Context, id int) (*ResourceAdvancedUserSearch, *resty.Response, error)

		// GetByName returns the specified advanced user search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersearchesbyname
		GetByName(ctx context.Context, name string) (*ResourceAdvancedUserSearch, *resty.Response, error)

		// Create creates a new advanced user search.
		//
		// Returns the created advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createadvancedusersearchgbyid
		Create(ctx context.Context, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified advanced user search by ID.
		//
		// Returns the updated advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyid
		UpdateByID(ctx context.Context, id int, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified advanced user search by name.
		//
		// Returns the updated advanced user search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyname
		UpdateByName(ctx context.Context, name string, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified advanced user search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified advanced user search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
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

// List returns all advanced user searches.
// URL: GET /JSSResource/advancedusersearches
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearches
func (s *Service) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
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

// GetByID returns the specified advanced user search by ID.
// URL: GET /JSSResource/advancedusersearches/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedusersearchesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceAdvancedUserSearch, *resty.Response, error) {
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

// GetByName returns the specified advanced user search by name.
// URL: GET /JSSResource/advancedusersearches/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersearchesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceAdvancedUserSearch, *resty.Response, error) {
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

// Create creates a new advanced user search.
// URL: POST /JSSResource/advancedusersearches/id/0
// Returns the created advanced user search ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createadvancedusersearchgbyid
func (s *Service) Create(ctx context.Context, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error) {
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

// UpdateByID updates the specified advanced user search by ID.
// URL: PUT /JSSResource/advancedusersearches/id/{id}
// Returns the updated advanced user search ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error) {
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

// UpdateByName updates the specified advanced user search by name.
// URL: PUT /JSSResource/advancedusersearches/name/{name}
// Returns the updated advanced user search ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedusersearchbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestAdvancedUserSearch) (*CreateUpdateResponse, *resty.Response, error) {
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

// DeleteByID removes the specified advanced user search by ID.
// URL: DELETE /JSSResource/advancedusersearches/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
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

// DeleteByName removes the specified advanced user search by name.
// URL: DELETE /JSSResource/advancedusersearches/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedusersearchbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
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
