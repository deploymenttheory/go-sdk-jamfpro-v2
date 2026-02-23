package patch_external_sources

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// PatchExternalSourcesServiceInterface defines the interface for Classic API patch external source operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
	PatchExternalSourcesServiceInterface interface {
		// List returns all patch external sources.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		GetByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// GetByName returns the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		GetByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// Create creates a new patch external source.
		//
		// Returns the created patch external source with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		Create(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// UpdateByID updates the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		UpdateByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// UpdateByName updates the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		UpdateByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// DeleteByID removes the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)
	}

	// Service handles communication with the patch external source-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PatchExternalSourcesServiceInterface = (*Service)(nil)

// NewService returns a new patch external sources Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Patch External Sources CRUD Operations
// -----------------------------------------------------------------------------

// List returns all patch external sources.
// URL: GET /JSSResource/patchexternalsources
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicPatchExternalSources

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

// GetByID returns the specified patch external source by ID.
// URL: GET /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) GetByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("patch external source ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPatchExternalSources, id)

	var result ResourcePatchExternalSource

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

// GetByName returns the specified patch external source by name.
// URL: GET /JSSResource/patchexternalsources/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) GetByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("patch external source name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPatchExternalSources, name)

	var result ResourcePatchExternalSource

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

// Create creates a new patch external source.
// URL: POST /JSSResource/patchexternalsources/id/0
// Returns the created patch external source with its assigned ID.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) Create(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicPatchExternalSources)

	var result ResourcePatchExternalSource

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

// UpdateByID updates the specified patch external source by ID.
// URL: PUT /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("patch external source ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPatchExternalSources, id)

	var result ResourcePatchExternalSource

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

// UpdateByName updates the specified patch external source by name.
// URL: PUT /JSSResource/patchexternalsources/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("patch external source name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPatchExternalSources, name)

	var result ResourcePatchExternalSource

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

// DeleteByID removes the specified patch external source by ID.
// URL: DELETE /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("patch external source ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPatchExternalSources, id)

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
