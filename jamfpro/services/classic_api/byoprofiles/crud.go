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
		// ListBYOProfiles returns all BYO profiles.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofiles
		ListBYOProfiles(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetBYOProfileByID returns the specified BYO profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyid
		GetBYOProfileByID(ctx context.Context, id int) (*ResourceBYOProfile, *interfaces.Response, error)

		// GetBYOProfileByName returns the specified BYO profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyname
		GetBYOProfileByName(ctx context.Context, name string) (*ResourceBYOProfile, *interfaces.Response, error)

		// CreateBYOProfile creates a new BYO profile.
		//
		// Returns the created BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createbyoprofilebyid
		CreateBYOProfile(ctx context.Context, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateBYOProfileByID updates the specified BYO profile by ID.
		//
		// Returns the updated BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilebyid
		UpdateBYOProfileByID(ctx context.Context, id int, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateBYOProfileByName updates the specified BYO profile by name.
		//
		// Returns the updated BYO profile ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilebyname
		UpdateBYOProfileByName(ctx context.Context, name string, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteBYOProfileByID removes the specified BYO profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilebyid
		DeleteBYOProfileByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteBYOProfileByName removes the specified BYO profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilebyname
		DeleteBYOProfileByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// ListBYOProfiles returns all BYO profiles.
// URL: GET /JSSResource/byoprofiles
// https://developer.jamf.com/jamf-pro/reference/findbyoprofiles
func (s *Service) ListBYOProfiles(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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

// GetBYOProfileByID returns the specified BYO profile by ID.
// URL: GET /JSSResource/byoprofiles/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyid
func (s *Service) GetBYOProfileByID(ctx context.Context, id int) (*ResourceBYOProfile, *interfaces.Response, error) {
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

// GetBYOProfileByName returns the specified BYO profile by name.
// URL: GET /JSSResource/byoprofiles/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyname
func (s *Service) GetBYOProfileByName(ctx context.Context, name string) (*ResourceBYOProfile, *interfaces.Response, error) {
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

// CreateBYOProfile creates a new BYO profile.
// URL: POST /JSSResource/byoprofiles/id/0
// Returns the created BYO profile ID only.
// https://developer.jamf.com/jamf-pro/reference/createbyoprofilebyid
func (s *Service) CreateBYOProfile(ctx context.Context, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateBYOProfileByID updates the specified BYO profile by ID.
// URL: PUT /JSSResource/byoprofiles/id/{id}
// Returns the updated BYO profile ID only.
// https://developer.jamf.com/jamf-pro/reference/updatebyoprofilebyid
func (s *Service) UpdateBYOProfileByID(ctx context.Context, id int, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateBYOProfileByName updates the specified BYO profile by name.
// URL: PUT /JSSResource/byoprofiles/name/{name}
// Returns the updated BYO profile ID only.
// https://developer.jamf.com/jamf-pro/reference/updatebyoprofilebyname
func (s *Service) UpdateBYOProfileByName(ctx context.Context, name string, req *RequestBYOProfile) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// DeleteBYOProfileByID removes the specified BYO profile by ID.
// URL: DELETE /JSSResource/byoprofiles/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletebyoprofilebyid
func (s *Service) DeleteBYOProfileByID(ctx context.Context, id int) (*interfaces.Response, error) {
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

// DeleteBYOProfileByName removes the specified BYO profile by name.
// URL: DELETE /JSSResource/byoprofiles/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletebyoprofilebyname
func (s *Service) DeleteBYOProfileByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
