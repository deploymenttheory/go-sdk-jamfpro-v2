package patch_external_sources

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// PatchExternalSourcesServiceInterface defines the interface for Classic API patch external source operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
	PatchExternalSourcesServiceInterface interface {
		// List returns all patch external sources.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		GetByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *resty.Response, error)

		// GetByName returns the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		GetByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *resty.Response, error)

		// Create creates a new patch external source.
		//
		// Returns the created patch external source with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		Create(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error)

		// UpdateByID updates the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		UpdateByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error)

		// UpdateByName updates the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		UpdateByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error)

		// DeleteByID removes the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)
	}

	// Service handles communication with the patch external source-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
	PatchExternalSources struct {
		client transport.HTTPClient
	}
)

var _ PatchExternalSourcesServiceInterface = (*PatchExternalSources)(nil)

// NewService returns a new patch external sources Service backed by the provided HTTP client.
func NewPatchExternalSources(client transport.HTTPClient) *PatchExternalSources {
	return &PatchExternalSources{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Patch External Sources CRUD Operations
// -----------------------------------------------------------------------------

// List returns all patch external sources.
// URL: GET /JSSResource/patchexternalsources
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicPatchExternalSources

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

// GetByID returns the specified patch external source by ID.
// URL: GET /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) GetByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("patch external source ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPatchExternalSources, id)

	var result ResourcePatchExternalSource

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

// GetByName returns the specified patch external source by name.
// URL: GET /JSSResource/patchexternalsources/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) GetByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("patch external source name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicPatchExternalSources, name)

	var result ResourcePatchExternalSource

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

// Create creates a new patch external source.
// URL: POST /JSSResource/patchexternalsources/id/0
// Returns the created patch external source with its assigned ID.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) Create(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicPatchExternalSources)

	var result ResourcePatchExternalSource

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

// UpdateByID updates the specified patch external source by ID.
// URL: PUT /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) UpdateByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("patch external source ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPatchExternalSources, id)

	var result ResourcePatchExternalSource

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

// UpdateByName updates the specified patch external source by name.
// URL: PUT /JSSResource/patchexternalsources/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) UpdateByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("patch external source name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicPatchExternalSources, name)

	var result ResourcePatchExternalSource

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

// DeleteByID removes the specified patch external source by ID.
// URL: DELETE /JSSResource/patchexternalsources/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/patchexternalsources
func (s *PatchExternalSources) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("patch external source ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPatchExternalSources, id)

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
