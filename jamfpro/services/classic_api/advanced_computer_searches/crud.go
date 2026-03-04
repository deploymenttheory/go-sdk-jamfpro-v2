package advanced_computer_searches

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// AdvancedComputerSearchesServiceInterface defines the interface for Classic API advanced computer search operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
	AdvancedComputerSearchesServiceInterface interface {
		// List returns all advanced computer searches.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified advanced computer search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyid
		GetByID(ctx context.Context, id int) (*ResourceAdvancedComputerSearch, *resty.Response, error)

		// GetByName returns the specified advanced computer search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyname
		GetByName(ctx context.Context, name string) (*ResourceAdvancedComputerSearch, *resty.Response, error)

		// Create creates a new advanced computer search.
		//
		// Returns the created advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createadvancedcomputersearchgbyid
		Create(ctx context.Context, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified advanced computer search by ID.
		//
		// Returns the updated advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyid
		UpdateByID(ctx context.Context, id int, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified advanced computer search by name.
		//
		// Returns the updated advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyname
		UpdateByName(ctx context.Context, name string, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified advanced computer search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified advanced computer search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the advanced computer searches-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AdvancedComputerSearchesServiceInterface = (*Service)(nil)

// NewService returns a new advanced computer searches Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Advanced Computer Searches CRUD Operations
// -----------------------------------------------------------------------------

// List returns all advanced computer searches.
// URL: GET /JSSResource/advancedcomputersearches
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
func (s *Service) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicAdvancedComputerSearches

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

// GetByID returns the specified advanced computer search by ID.
// URL: GET /JSSResource/advancedcomputersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceAdvancedComputerSearch, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("advanced computer search ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedComputerSearches, id)

	var result ResourceAdvancedComputerSearch

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

// GetByName returns the specified advanced computer search by name.
// URL: GET /JSSResource/advancedcomputersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceAdvancedComputerSearch, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("advanced computer search name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedComputerSearches, name)

	var result ResourceAdvancedComputerSearch

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

// Create creates a new advanced computer search.
// URL: POST /JSSResource/advancedcomputersearches/id/0
// Returns the created advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/createadvancedcomputersearchgbyid
func (s *Service) Create(ctx context.Context, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicAdvancedComputerSearches)

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

// UpdateByID updates the specified advanced computer search by ID.
// URL: PUT /JSSResource/advancedcomputersearches/id/{id}
// Returns the updated advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("advanced computer search ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedComputerSearches, id)

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

// UpdateByName updates the specified advanced computer search by name.
// URL: PUT /JSSResource/advancedcomputersearches/name/{name}
// Returns the updated advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("advanced computer search name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedComputerSearches, name)

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

// DeleteByID removes the specified advanced computer search by ID.
// URL: DELETE /JSSResource/advancedcomputersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("advanced computer search ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAdvancedComputerSearches, id)

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

// DeleteByName removes the specified advanced computer search by name.
// URL: DELETE /JSSResource/advancedcomputersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("advanced computer search name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicAdvancedComputerSearches, name)

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
