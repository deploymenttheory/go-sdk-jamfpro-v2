package sites

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SitesServiceInterface defines the interface for Classic API site operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/sites
	SitesServiceInterface interface {
		// List returns all sites.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsites
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsitesbyid
		GetByID(ctx context.Context, id int) (*ResourceSite, *resty.Response, error)

		// GetByName returns the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsitesbyname
		GetByName(ctx context.Context, name string) (*ResourceSite, *resty.Response, error)

		// Create creates a new site.
		//
		// Returns the created site with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createsitebyid
		Create(ctx context.Context, req *RequestSite) (*ResourceSite, *resty.Response, error)

		// UpdateByID updates the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesitebyid
		UpdateByID(ctx context.Context, id int, req *RequestSite) (*ResourceSite, *resty.Response, error)

		// UpdateByName updates the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesitebyname
		UpdateByName(ctx context.Context, name string, req *RequestSite) (*ResourceSite, *resty.Response, error)

		// DeleteByID removes the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesitebyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesitebyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the sites-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/sites
	Sites struct {
		client interfaces.HTTPClient
	}
)

var _ SitesServiceInterface = (*Sites)(nil)

// NewService returns a new sites Service backed by the provided HTTP client.
func NewSites(client interfaces.HTTPClient) *Sites {
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

	endpoint := EndpointClassicSites

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

// GetByID returns the specified site by ID.
// URL: GET /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyid
func (s *Sites) GetByID(ctx context.Context, id int) (*ResourceSite, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("site ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

	var result ResourceSite

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

// GetByName returns the specified site by name.
// URL: GET /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyname
func (s *Sites) GetByName(ctx context.Context, name string) (*ResourceSite, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("site name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

	var result ResourceSite

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

// Create creates a new site.
// URL: POST /JSSResource/sites/id/0
// Returns the created site with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsitebyid
func (s *Sites) Create(ctx context.Context, req *RequestSite) (*ResourceSite, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicSites)

	var result ResourceSite

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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

	var result ResourceSite

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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

	var result ResourceSite

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

// DeleteByID removes the specified site by ID.
// URL: DELETE /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyid
func (s *Sites) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("site ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

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

// DeleteByName removes the specified site by name.
// URL: DELETE /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyname
func (s *Sites) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("site name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

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
