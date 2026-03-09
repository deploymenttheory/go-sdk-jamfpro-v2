package software_update_servers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the software update server-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
	SoftwareUpdateServers struct {
		client transport.HTTPClient
	}
)

// NewService returns a new software update servers Service backed by the provided HTTP client.
func NewSoftwareUpdateServers(client transport.HTTPClient) *SoftwareUpdateServers {
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

	endpoint := constants.EndpointClassicSoftwareUpdateServers

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

// GetByID returns the specified software update server by ID.
// URL: GET /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyid
func (s *SoftwareUpdateServers) GetByID(ctx context.Context, id int) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

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

// GetByName returns the specified software update server by name.
// URL: GET /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyname
func (s *SoftwareUpdateServers) GetByName(ctx context.Context, name string) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

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

// Create creates a new software update server.
// URL: POST /JSSResource/softwareupdateservers/id/0
// Returns the created software update server with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsoftwareupdateserverbyid
func (s *SoftwareUpdateServers) Create(ctx context.Context, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicSoftwareUpdateServers)

	var result ResourceSoftwareUpdateServer

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

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

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

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

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

// DeleteByID removes the specified software update server by ID.
// URL: DELETE /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyid
func (s *SoftwareUpdateServers) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicSoftwareUpdateServers, id)

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

// DeleteByName removes the specified software update server by name.
// URL: DELETE /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyname
func (s *SoftwareUpdateServers) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicSoftwareUpdateServers, name)

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
