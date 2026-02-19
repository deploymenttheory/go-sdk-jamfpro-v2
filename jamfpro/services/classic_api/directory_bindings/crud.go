package directory_bindings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// DirectoryBindingsServiceInterface defines the interface for Classic API directory binding operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/directorybindings
	DirectoryBindingsServiceInterface interface {
		// ListDirectoryBindings returns all directory bindings.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findalldirectorybindings
		ListDirectoryBindings(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetDirectoryBindingByID returns the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyid
		GetDirectoryBindingByID(ctx context.Context, id int) (*ResourceDirectoryBinding, *interfaces.Response, error)

		// GetDirectoryBindingByName returns the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyname
		GetDirectoryBindingByName(ctx context.Context, name string) (*ResourceDirectoryBinding, *interfaces.Response, error)

		// CreateDirectoryBinding creates a new directory binding.
		//
		// Returns the created directory binding with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createdirectorybinding
		CreateDirectoryBinding(ctx context.Context, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error)

		// UpdateDirectoryBindingByID updates the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyid
		UpdateDirectoryBindingByID(ctx context.Context, id int, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error)

		// UpdateDirectoryBindingByName updates the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyname
		UpdateDirectoryBindingByName(ctx context.Context, name string, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error)

		// DeleteDirectoryBindingByID removes the specified directory binding by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyid
		DeleteDirectoryBindingByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteDirectoryBindingByName removes the specified directory binding by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyname
		DeleteDirectoryBindingByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the directory binding-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/directorybindings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DirectoryBindingsServiceInterface = (*Service)(nil)

// NewService returns a new directory bindings Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Directory Bindings CRUD Operations
// -----------------------------------------------------------------------------

// ListDirectoryBindings returns all directory bindings.
// URL: GET /JSSResource/directorybindings
// https://developer.jamf.com/jamf-pro/reference/findalldirectorybindings
func (s *Service) ListDirectoryBindings(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicDirectoryBindings, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDirectoryBindingByID returns the specified directory binding by ID.
// URL: GET /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyid
func (s *Service) GetDirectoryBindingByID(ctx context.Context, id int) (*ResourceDirectoryBinding, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("directory binding ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDirectoryBindings, id)

	var result ResourceDirectoryBinding

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDirectoryBindingByName returns the specified directory binding by name.
// URL: GET /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/finddirectorybindingsbyname
func (s *Service) GetDirectoryBindingByName(ctx context.Context, name string) (*ResourceDirectoryBinding, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("directory binding name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDirectoryBindings, name)

	var result ResourceDirectoryBinding

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateDirectoryBinding creates a new directory binding.
// URL: POST /JSSResource/directorybindings/id/0
// Returns the created directory binding with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createdirectorybinding
func (s *Service) CreateDirectoryBinding(ctx context.Context, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicDirectoryBindings)

	var result ResourceDirectoryBinding

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDirectoryBindingByID updates the specified directory binding by ID.
// URL: PUT /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyid
func (s *Service) UpdateDirectoryBindingByID(ctx context.Context, id int, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("directory binding ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDirectoryBindings, id)

	var result ResourceDirectoryBinding

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDirectoryBindingByName updates the specified directory binding by name.
// URL: PUT /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatedirectorybindingbyname
func (s *Service) UpdateDirectoryBindingByName(ctx context.Context, name string, req *RequestDirectoryBinding) (*ResourceDirectoryBinding, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("directory binding name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDirectoryBindings, name)

	var result ResourceDirectoryBinding

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteDirectoryBindingByID removes the specified directory binding by ID.
// URL: DELETE /JSSResource/directorybindings/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyid
func (s *Service) DeleteDirectoryBindingByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("directory binding ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDirectoryBindings, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteDirectoryBindingByName removes the specified directory binding by name.
// URL: DELETE /JSSResource/directorybindings/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletedirectorybindingbyname
func (s *Service) DeleteDirectoryBindingByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("directory binding name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDirectoryBindings, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
