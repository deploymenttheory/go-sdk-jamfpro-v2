package allowed_file_extensions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AllowedFileExtensionsServiceInterface defines the interface for Classic API allowed file extension operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/allowedfileextensions
	AllowedFileExtensionsServiceInterface interface {
		// ListAllowedFileExtensions returns all allowed file extensions.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextensions
		ListAllowedFileExtensions(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetAllowedFileExtensionByID returns the specified allowed file extension by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionsbyid
		GetAllowedFileExtensionByID(ctx context.Context, id int) (*ResourceAllowedFileExtension, *interfaces.Response, error)

		// GetAllowedFileExtensionByExtension returns the allowed file extension matching the given extension string.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionsbyextension
		GetAllowedFileExtensionByExtension(ctx context.Context, extension string) (*ResourceAllowedFileExtension, *interfaces.Response, error)

		// CreateAllowedFileExtension creates a new allowed file extension.
		//
		// Returns the created resource with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createallowedfileextensionbyid
		CreateAllowedFileExtension(ctx context.Context, req *RequestAllowedFileExtension) (*ResourceAllowedFileExtension, *interfaces.Response, error)

		// DeleteAllowedFileExtensionByID removes the specified allowed file extension by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteallowedfileextensionbyid
		DeleteAllowedFileExtensionByID(ctx context.Context, id int) (*interfaces.Response, error)
	}

	// Service handles communication with the allowed file extension-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/allowedfileextensions
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AllowedFileExtensionsServiceInterface = (*Service)(nil)

// NewService returns a new allowed file extensions Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Allowed File Extensions CRUD Operations
// -----------------------------------------------------------------------------

// ListAllowedFileExtensions returns all allowed file extensions.
// URL: GET /JSSResource/allowedfileextensions
// https://developer.jamf.com/jamf-pro/reference/findallowedfileextensions
func (s *Service) ListAllowedFileExtensions(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicAllowedFileExtensions

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

// GetAllowedFileExtensionByID returns the specified allowed file extension by ID.
// URL: GET /JSSResource/allowedfileextensions/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionsbyid
func (s *Service) GetAllowedFileExtensionByID(ctx context.Context, id int) (*ResourceAllowedFileExtension, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("allowed file extension ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAllowedFileExtensions, id)

	var result ResourceAllowedFileExtension

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

// GetAllowedFileExtensionByExtension returns the allowed file extension matching the given extension string.
// URL: GET /JSSResource/allowedfileextensions/extension/{extension}
// https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionsbyextension
func (s *Service) GetAllowedFileExtensionByExtension(ctx context.Context, extension string) (*ResourceAllowedFileExtension, *interfaces.Response, error) {
	if extension == "" {
		return nil, nil, fmt.Errorf("extension is required")
	}

	endpoint := fmt.Sprintf("%s/extension/%s", EndpointClassicAllowedFileExtensions, extension)

	var result ResourceAllowedFileExtension

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

// CreateAllowedFileExtension creates a new allowed file extension.
// URL: POST /JSSResource/allowedfileextensions/id/0
// Returns the created resource with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createallowedfileextensionbyid
func (s *Service) CreateAllowedFileExtension(ctx context.Context, req *RequestAllowedFileExtension) (*ResourceAllowedFileExtension, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicAllowedFileExtensions)

	var result ResourceAllowedFileExtension

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

// DeleteAllowedFileExtensionByID removes the specified allowed file extension by ID.
// URL: DELETE /JSSResource/allowedfileextensions/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteallowedfileextensionbyid
func (s *Service) DeleteAllowedFileExtensionByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("allowed file extension ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicAllowedFileExtensions, id)

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
