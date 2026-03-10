package byoprofiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the BYO profiles-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/byoprofiles
	Byoprofiles struct {
		client client.Client
	}
)

// NewService returns a new BYO profiles Service backed by the provided HTTP client.
func NewByoprofiles(client client.Client) *Byoprofiles {
	return &Byoprofiles{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - BYO Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all BYO profiles.
// URL: GET /JSSResource/byoprofiles
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofiles
func (s *Byoprofiles) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicBYOProfiles

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified BYO profile by ID.
// URL: GET /JSSResource/byoprofiles/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyid
func (s *Byoprofiles) GetByID(ctx context.Context, id int) (*ResourceBYOProfile, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}

	var result ResourceBYOProfile

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicBYOProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified BYO profile by name.
// URL: GET /JSSResource/byoprofiles/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findbyoprofilesbyname
func (s *Byoprofiles) GetByName(ctx context.Context, name string) (*ResourceBYOProfile, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("BYO profile name is required")
	}

	var result ResourceBYOProfile

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicBYOProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new BYO profile.
// URL: POST /JSSResource/byoprofiles/id/0
// Returns the created BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createbyoprofilesbyid
func (s *Byoprofiles) Create(ctx context.Context, req *RequestBYOProfile) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicBYOProfiles)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified BYO profile by ID.
// URL: PUT /JSSResource/byoprofiles/id/{id}
// Returns the updated BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyid
func (s *Byoprofiles) UpdateByID(ctx context.Context, id int, req *RequestBYOProfile) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicBYOProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified BYO profile by name.
// URL: PUT /JSSResource/byoprofiles/name/{name}
// Returns the updated BYO profile ID only.
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatebyoprofilesbyname
func (s *Byoprofiles) UpdateByName(ctx context.Context, name string, req *RequestBYOProfile) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("BYO profile name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicBYOProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified BYO profile by ID.
// URL: DELETE /JSSResource/byoprofiles/id/{id}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyid
func (s *Byoprofiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("BYO profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicBYOProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified BYO profile by name.
// URL: DELETE /JSSResource/byoprofiles/name/{name}
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletebyoprofilesbyname
func (s *Byoprofiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("BYO profile name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicBYOProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
