package software_update_servers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SoftwareUpdateServersServiceInterface defines the interface for Classic API software update server operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
	SoftwareUpdateServersServiceInterface interface {
		// List returns all software update servers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateservers
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyid
		GetByID(ctx context.Context, id int) (*ResourceSoftwareUpdateServer, *resty.Response, error)

		// GetByName returns the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyname
		GetByName(ctx context.Context, name string) (*ResourceSoftwareUpdateServer, *resty.Response, error)

		// Create creates a new software update server.
		//
		// Returns the created software update server with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createsoftwareupdateserverbyid
		Create(ctx context.Context, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error)

		// UpdateByID updates the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyid
		UpdateByID(ctx context.Context, id int, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error)

		// UpdateByName updates the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyname
		UpdateByName(ctx context.Context, name string, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error)

		// DeleteByID removes the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the software update server-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
	SoftwareUpdateServers struct {
		client interfaces.HTTPClient
	}
)

var _ SoftwareUpdateServersServiceInterface = (*SoftwareUpdateServers)(nil)

// NewService returns a new software update servers Service backed by the provided HTTP client.
func NewSoftwareUpdateServers(client interfaces.HTTPClient) *SoftwareUpdateServers {
	return &SoftwareUpdateServers{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Software Update Servers CRUD Operations
// -----------------------------------------------------------------------------

// List returns all software update servers.
// URL: GET /JSSResource/softwareupdateservers
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateservers
func (s *SoftwareUpdateServers) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicSoftwareUpdateServers

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

// GetByID returns the specified software update server by ID.
// URL: GET /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyid
func (s *SoftwareUpdateServers) GetByID(ctx context.Context, id int) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

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

// GetByName returns the specified software update server by name.
// URL: GET /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyname
func (s *SoftwareUpdateServers) GetByName(ctx context.Context, name string) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

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

// Create creates a new software update server.
// URL: POST /JSSResource/softwareupdateservers/id/0
// Returns the created software update server with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsoftwareupdateserverbyid
func (s *SoftwareUpdateServers) Create(ctx context.Context, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicSoftwareUpdateServers)

	var result ResourceSoftwareUpdateServer

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

// UpdateByID updates the specified software update server by ID.
// URL: PUT /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyid
func (s *SoftwareUpdateServers) UpdateByID(ctx context.Context, id int, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("software update server ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

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

// UpdateByName updates the specified software update server by name.
// URL: PUT /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyname
func (s *SoftwareUpdateServers) UpdateByName(ctx context.Context, name string, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("software update server name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

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

// DeleteByID removes the specified software update server by ID.
// URL: DELETE /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyid
func (s *SoftwareUpdateServers) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

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

// DeleteByName removes the specified software update server by name.
// URL: DELETE /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyname
func (s *SoftwareUpdateServers) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

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
