package directory_bindings

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// DirectoryBindingsServiceInterface defines the interface for Classic API directory binding operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/directorybindings
	DirectoryBindingsServiceInterface interface {
		// List returns all directory bindings.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddirectorybindings
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyid
		GetByID(ctx context.Context, id int) (*ResourceDirectoryBinding, *resty.Response, error)

		// GetByName returns the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyname
		GetByName(ctx context.Context, name string) (*ResourceDirectoryBinding, *resty.Response, error)

		// Create creates a new directory binding.
		//
		// Returns the created directory binding with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdirectorybindingbyid
		Create(ctx context.Context, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error)

		// UpdateByID updates the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyid
		UpdateByID(ctx context.Context, id int, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error)

		// UpdateByName updates the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyname
		UpdateByName(ctx context.Context, name string, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error)

		// DeleteDirectoryBindingByID removes the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteDirectoryBindingByName removes the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the directory binding-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/directorybindings
	DirectoryBindings struct {
		client transport.HTTPClient
	}
)

var _ DirectoryBindingsServiceInterface = (*DirectoryBindings)(nil)

// NewService returns a new directory bindings Service backed by the provided HTTP client.
func NewDirectoryBindings(client transport.HTTPClient) *DirectoryBindings {
	return &DirectoryBindings{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Directory Bindings CRUD Operations
// -----------------------------------------------------------------------------

// List returns all directory bindings.
// URL: GET /JSSResource/directorybindings
// https://developer.jamf.com/jamf-pro/reference/findalldirectorybindings
func (s *DirectoryBindings) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicDirectoryBindings

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified directory binding by ID.
// URL: GET /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyid
func (s *DirectoryBindings) GetByID(ctx context.Context, id int) (*ResourceDirectoryBinding, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("directory binding ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDirectoryBindings, id)

	var result ResourceDirectoryBinding

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified directory binding by name.
// URL: GET /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyname
func (s *DirectoryBindings) GetByName(ctx context.Context, name string) (*ResourceDirectoryBinding, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("directory binding name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDirectoryBindings, url.PathEscape(name))

	var result ResourceDirectoryBinding

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new directory binding.
// URL: POST /JSSResource/directorybindings/id/0
// Returns the created directory binding with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createdirectorybindingbyid
func (s *DirectoryBindings) Create(ctx context.Context, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicDirectoryBindings)

	var result ResourceDirectoryBinding

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified directory binding by ID.
// URL: PUT /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyid
func (s *DirectoryBindings) UpdateByID(ctx context.Context, id int, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("directory binding ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDirectoryBindings, id)

	var result ResourceDirectoryBinding

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified directory binding by name.
// URL: PUT /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyname
func (s *DirectoryBindings) UpdateByName(ctx context.Context, name string, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("directory binding name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDirectoryBindings, url.PathEscape(name))

	var result ResourceDirectoryBinding

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified directory binding by ID.
// URL: DELETE /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyid
func (s *DirectoryBindings) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("directory binding ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDirectoryBindings, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified directory binding by name.
// URL: DELETE /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyname
func (s *DirectoryBindings) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("directory binding name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDirectoryBindings, url.PathEscape(name))

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
