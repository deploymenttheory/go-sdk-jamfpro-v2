package advanced_computer_searches

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AdvancedComputerSearchesServiceInterface defines the interface for Classic API advanced computer search operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
	AdvancedComputerSearchesServiceInterface interface {
		// ListAdvancedComputerSearches returns all advanced computer searches.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
		ListAdvancedComputerSearches(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetAdvancedComputerSearchByID returns the specified advanced computer search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyid
		GetAdvancedComputerSearchByID(ctx context.Context, id int) (*ResourceAdvancedComputerSearch, *interfaces.Response, error)

		// GetAdvancedComputerSearchByName returns the specified advanced computer search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyname
		GetAdvancedComputerSearchByName(ctx context.Context, name string) (*ResourceAdvancedComputerSearch, *interfaces.Response, error)

		// CreateAdvancedComputerSearch creates a new advanced computer search.
		//
		// Returns the created advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createadvancedcomputersearchgbyid
		CreateAdvancedComputerSearch(ctx context.Context, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateAdvancedComputerSearchByID updates the specified advanced computer search by ID.
		//
		// Returns the updated advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyid
		UpdateAdvancedComputerSearchByID(ctx context.Context, id int, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateAdvancedComputerSearchByName updates the specified advanced computer search by name.
		//
		// Returns the updated advanced computer search ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyname
		UpdateAdvancedComputerSearchByName(ctx context.Context, name string, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteAdvancedComputerSearchByID removes the specified advanced computer search by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyid
		DeleteAdvancedComputerSearchByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteAdvancedComputerSearchByName removes the specified advanced computer search by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyname
		DeleteAdvancedComputerSearchByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// ListAdvancedComputerSearches returns all advanced computer searches.
// URL: GET /JSSResource/advancedcomputersearches
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearches
func (s *Service) ListAdvancedComputerSearches(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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

// GetAdvancedComputerSearchByID returns the specified advanced computer search by ID.
// URL: GET /JSSResource/advancedcomputersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyid
func (s *Service) GetAdvancedComputerSearchByID(ctx context.Context, id int) (*ResourceAdvancedComputerSearch, *interfaces.Response, error) {
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

// GetAdvancedComputerSearchByName returns the specified advanced computer search by name.
// URL: GET /JSSResource/advancedcomputersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findadvancedcomputersearchesbyname
func (s *Service) GetAdvancedComputerSearchByName(ctx context.Context, name string) (*ResourceAdvancedComputerSearch, *interfaces.Response, error) {
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

// CreateAdvancedComputerSearch creates a new advanced computer search.
// URL: POST /JSSResource/advancedcomputersearches/id/0
// Returns the created advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/createadvancedcomputersearchgbyid
func (s *Service) CreateAdvancedComputerSearch(ctx context.Context, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateAdvancedComputerSearchByID updates the specified advanced computer search by ID.
// URL: PUT /JSSResource/advancedcomputersearches/id/{id}
// Returns the updated advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyid
func (s *Service) UpdateAdvancedComputerSearchByID(ctx context.Context, id int, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateAdvancedComputerSearchByName updates the specified advanced computer search by name.
// URL: PUT /JSSResource/advancedcomputersearches/name/{name}
// Returns the updated advanced computer search ID only.
// https://developer.jamf.com/jamf-pro/reference/updateadvancedcomputersearchbyname
func (s *Service) UpdateAdvancedComputerSearchByName(ctx context.Context, name string, req *RequestAdvancedComputerSearch) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// DeleteAdvancedComputerSearchByID removes the specified advanced computer search by ID.
// URL: DELETE /JSSResource/advancedcomputersearches/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyid
func (s *Service) DeleteAdvancedComputerSearchByID(ctx context.Context, id int) (*interfaces.Response, error) {
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

// DeleteAdvancedComputerSearchByName removes the specified advanced computer search by name.
// URL: DELETE /JSSResource/advancedcomputersearches/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteadvancedcomputersearchbyname
func (s *Service) DeleteAdvancedComputerSearchByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
