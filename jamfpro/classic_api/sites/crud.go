package sites

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the sites-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/sites
	Sites struct {
		client client.Client
	}
)

// NewService returns a new sites Service backed by the provided HTTP client.
func NewSites(client client.Client) *Sites {
	return &Sites{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Sites CRUD Operations
// -----------------------------------------------------------------------------

// List returns all sites.
// URL: GET /JSSResource/sites
// https://developer.jamf.com/jamf-pro/reference/findsites
func (s *Sites) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicSites

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

// GetByID returns the specified site by ID.
// URL: GET /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyid
func (s *Sites) GetByID(ctx context.Context, id int) (*ResourceSite, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("site ID must be a positive integer")
	}

	var result ResourceSite

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSites, id)

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

// GetByName returns the specified site by name.
// URL: GET /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyname
func (s *Sites) GetByName(ctx context.Context, name string) (*ResourceSite, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("site name is required")
	}

	var result ResourceSite

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSites, name)

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

// Create creates a new site.
// URL: POST /JSSResource/sites/id/0
// Returns the created site with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsitebyid
func (s *Sites) Create(ctx context.Context, req *RequestSite) (*ResourceSite, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSite

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicSites)

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

// UpdateByID updates the specified site by ID.
// URL: PUT /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatesitebyid
func (s *Sites) UpdateByID(ctx context.Context, id int, req *RequestSite) (*ResourceSite, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("site ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSite

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSites, id)

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

// UpdateByName updates the specified site by name.
// URL: PUT /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatesitebyname
func (s *Sites) UpdateByName(ctx context.Context, name string, req *RequestSite) (*ResourceSite, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("site name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSite

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSites, name)

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

// DeleteByID removes the specified site by ID.
// URL: DELETE /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyid
func (s *Sites) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("site ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSites, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified site by name.
// URL: DELETE /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyname
func (s *Sites) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("site name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSites, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
