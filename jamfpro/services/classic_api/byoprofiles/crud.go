package byoprofiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// BYOProfilesServiceInterface defines the interface for Classic API BYO profile operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/byoprofiles
	BYOProfilesServiceInterface interface {
		// List returns all BYO profiles.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofiles
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified BYO profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyid
		GetByID(ctx context.Context, id int) (*ResourceBYOProfile, *interfaces.Response, error)

		// GetByName returns the specified BYO profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyname
		GetByName(ctx context.Context, name string) (*ResourceBYOProfile, *interfaces.Response, error)

		// Create creates a new BYO profile.
		//
		// Returns the created BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createbyoprofilesbyid
		Create(ctx context.Context, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified BYO profile by ID.
		//
		// Returns the updated BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyid
		UpdateByID(ctx context.Context, id int, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified BYO profile by name.
		//
		// Returns the updated BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyname
		UpdateByName(ctx context.Context, name string, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified BYO profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified BYO profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the BYO profiles-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/byoprofiles
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ BYOProfilesServiceInterface = (*Service)(nil)

// NewService returns a new BYO profiles Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - BYO Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all BYO profiles.
// URL: GET /JSSResource/byoprofiles
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofiles
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicBYOProfiles

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

// GetByID returns the specified BYO profile by ID.
// URL: GET /JSSResource/byoprofiles/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceBYOProfile, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicBYOProfiles, id)

	var result ResourceBYOProfile

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

// GetByName returns the specified BYO profile by name.
// URL: GET /JSSResource/byoprofiles/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceBYOProfile, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("BYO profile name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicBYOProfiles, name)

	var result ResourceBYOProfile

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

// Create creates a new BYO profile.
// URL: POST /JSSResource/byoprofiles/id/0
// Returns the created BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createbyoprofilesbyid
func (s *Service) Create(ctx context.Context, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicBYOProfiles)

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

// UpdateByID updates the specified BYO profile by ID.
// URL: PUT /JSSResource/byoprofiles/id/{id}
// Returns the updated BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicBYOProfiles, id)

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

// UpdateByName updates the specified BYO profile by name.
// URL: PUT /JSSResource/byoprofiles/name/{name}
// Returns the updated BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("BYO profile name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicBYOProfiles, name)

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

// DeleteByID removes the specified BYO profile by ID.
// URL: DELETE /JSSResource/byoprofiles/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicBYOProfiles, id)

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

// DeleteByName removes the specified BYO profile by name.
// URL: DELETE /JSSResource/byoprofiles/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("BYO profile name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicBYOProfiles, name)

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
