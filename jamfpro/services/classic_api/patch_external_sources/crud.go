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
		// ListPatchExternalSources returns all patch external sources.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallpatchexternalsources
		ListPatchExternalSources(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetPatchExternalSourceByID returns the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpatchexternalsourcesbyid
		GetPatchExternalSourceByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// GetPatchExternalSourceByName returns the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpatchexternalsourcesbyname
		GetPatchExternalSourceByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// CreatePatchExternalSource creates a new patch external source.
		//
		// Returns the created patch external source with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createpatchexternalsource
		CreatePatchExternalSource(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// UpdatePatchExternalSourceByID updates the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatepatchexternalsourcebyid
		UpdatePatchExternalSourceByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// UpdatePatchExternalSourceByName updates the specified patch external source by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatepatchexternalsourcebyname
		UpdatePatchExternalSourceByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error)

		// DeletePatchExternalSourceByID removes the specified patch external source by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletepatchexternalsourcebyid
		DeletePatchExternalSourceByID(ctx context.Context, id int) (*interfaces.Response, error)
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

// ListPatchExternalSources returns all patch external sources.
// URL: GET /JSSResource/patchexternalsources
// https://developer.jamf.com/jamf-pro/reference/findallpatchexternalsources
func (s *Service) ListPatchExternalSources(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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

// GetPatchExternalSourceByID returns the specified patch external source by ID.
// URL: GET /JSSResource/patchexternalsources/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findpatchexternalsourcesbyid
func (s *Service) GetPatchExternalSourceByID(ctx context.Context, id int) (*ResourcePatchExternalSource, *interfaces.Response, error) {
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

// GetPatchExternalSourceByName returns the specified patch external source by name.
// URL: GET /JSSResource/patchexternalsources/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findpatchexternalsourcesbyname
func (s *Service) GetPatchExternalSourceByName(ctx context.Context, name string) (*ResourcePatchExternalSource, *interfaces.Response, error) {
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

// CreatePatchExternalSource creates a new patch external source.
// URL: POST /JSSResource/patchexternalsources/id/0
// Returns the created patch external source with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createpatchexternalsource
func (s *Service) CreatePatchExternalSource(ctx context.Context, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
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

// UpdatePatchExternalSourceByID updates the specified patch external source by ID.
// URL: PUT /JSSResource/patchexternalsources/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatepatchexternalsourcebyid
func (s *Service) UpdatePatchExternalSourceByID(ctx context.Context, id int, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
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

// UpdatePatchExternalSourceByName updates the specified patch external source by name.
// URL: PUT /JSSResource/patchexternalsources/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatepatchexternalsourcebyname
func (s *Service) UpdatePatchExternalSourceByName(ctx context.Context, name string, req *RequestPatchExternalSource) (*ResourcePatchExternalSource, *interfaces.Response, error) {
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

// DeletePatchExternalSourceByID removes the specified patch external source by ID.
// URL: DELETE /JSSResource/patchexternalsources/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletepatchexternalsourcebyid
func (s *Service) DeletePatchExternalSourceByID(ctx context.Context, id int) (*interfaces.Response, error) {
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
