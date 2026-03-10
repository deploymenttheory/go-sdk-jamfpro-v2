package allowed_file_extensions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the allowed file extension-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextension
	AllowedFileExtensions struct {
		client client.Client
	}
)

// NewService returns a new allowed file extensions Service backed by the provided HTTP client.
func NewAllowedFileExtensions(client client.Client) *AllowedFileExtensions {
	return &AllowedFileExtensions{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Allowed File Extensions CRUD Operations
// -----------------------------------------------------------------------------

// List returns all allowed file extensions.
// URL: GET /JSSResource/allowedfileextensions
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextension
func (s *AllowedFileExtensions) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicAllowedFileExtensions

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

// GetByID returns the specified allowed file extension by ID.
// URL: GET /JSSResource/allowedfileextensions/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionbyid
func (s *AllowedFileExtensions) GetByID(ctx context.Context, id int) (*ResourceAllowedFileExtension, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("allowed file extension ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicAllowedFileExtensions, id)

	var result ResourceAllowedFileExtension

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

// GetByExtension returns the allowed file extension matching the given extension string.
// URL: GET /JSSResource/allowedfileextensions/extension/{extension}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallowedfileextensionbyname
func (s *AllowedFileExtensions) GetByExtension(ctx context.Context, extension string) (*ResourceAllowedFileExtension, *resty.Response, error) {
	if extension == "" {
		return nil, nil, fmt.Errorf("extension is required")
	}

	endpoint := fmt.Sprintf("%s/extension/%s", constants.EndpointClassicAllowedFileExtensions, extension)

	var result ResourceAllowedFileExtension

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

// Create creates a new allowed file extension.
// URL: POST /JSSResource/allowedfileextensions/id/0
// Returns the created resource with its assigned ID.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createallowedfileextensionbyid
func (s *AllowedFileExtensions) Create(ctx context.Context, req *RequestAllowedFileExtension) (*ResourceAllowedFileExtension, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicAllowedFileExtensions)

	var result ResourceAllowedFileExtension

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

// DeleteByID removes the specified allowed file extension by ID.
// URL: DELETE /JSSResource/allowedfileextensions/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteallowedfileextensionbyid
func (s *AllowedFileExtensions) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("allowed file extension ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicAllowedFileExtensions, id)

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
