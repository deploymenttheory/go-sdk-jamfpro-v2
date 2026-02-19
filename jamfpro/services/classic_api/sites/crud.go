package sites

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// SitesServiceInterface defines the interface for Classic API site operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/sites
	SitesServiceInterface interface {
		// ListSites returns all sites.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallsites
		ListSites(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetSiteByID returns the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsitesbyid
		GetSiteByID(ctx context.Context, id int) (*ResourceSite, *interfaces.Response, error)

		// GetSiteByName returns the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsitesbyname
		GetSiteByName(ctx context.Context, name string) (*ResourceSite, *interfaces.Response, error)

		// CreateSite creates a new site.
		//
		// Returns the created site with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createsitebyid
		CreateSite(ctx context.Context, req *RequestSite) (*ResourceSite, *interfaces.Response, error)

		// UpdateSiteByID updates the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesitebyid
		UpdateSiteByID(ctx context.Context, id int, req *RequestSite) (*ResourceSite, *interfaces.Response, error)

		// UpdateSiteByName updates the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesitebyname
		UpdateSiteByName(ctx context.Context, name string, req *RequestSite) (*ResourceSite, *interfaces.Response, error)

		// DeleteSiteByID removes the specified site by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesitebyid
		DeleteSiteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteSiteByName removes the specified site by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesitebyname
		DeleteSiteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the sites-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/sites
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SitesServiceInterface = (*Service)(nil)

// NewService returns a new sites Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Sites CRUD Operations
// -----------------------------------------------------------------------------

// ListSites returns all sites.
// URL: GET /JSSResource/sites
// https://developer.jamf.com/jamf-pro/reference/findallsites
func (s *Service) ListSites(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicSites, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSiteByID returns the specified site by ID.
// URL: GET /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyid
func (s *Service) GetSiteByID(ctx context.Context, id int) (*ResourceSite, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("site ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

	var result ResourceSite

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSiteByName returns the specified site by name.
// URL: GET /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsitesbyname
func (s *Service) GetSiteByName(ctx context.Context, name string) (*ResourceSite, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("site name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

	var result ResourceSite

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateSite creates a new site.
// URL: POST /JSSResource/sites/id/0
// Returns the created site with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsitebyid
func (s *Service) CreateSite(ctx context.Context, req *RequestSite) (*ResourceSite, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicSites)

	var result ResourceSite

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSiteByID updates the specified site by ID.
// URL: PUT /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatesitebyid
func (s *Service) UpdateSiteByID(ctx context.Context, id int, req *RequestSite) (*ResourceSite, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("site ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

	var result ResourceSite

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSiteByName updates the specified site by name.
// URL: PUT /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatesitebyname
func (s *Service) UpdateSiteByName(ctx context.Context, name string, req *RequestSite) (*ResourceSite, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("site name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

	var result ResourceSite

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSiteByID removes the specified site by ID.
// URL: DELETE /JSSResource/sites/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyid
func (s *Service) DeleteSiteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("site ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSites, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSiteByName removes the specified site by name.
// URL: DELETE /JSSResource/sites/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesitebyname
func (s *Service) DeleteSiteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("site name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSites, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
