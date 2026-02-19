package software_update_servers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// SoftwareUpdateServersServiceInterface defines the interface for Classic API software update server operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
	SoftwareUpdateServersServiceInterface interface {
		// ListSoftwareUpdateServers returns all software update servers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallsoftwareupdateservers
		ListSoftwareUpdateServers(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetSoftwareUpdateServerByID returns the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyid
		GetSoftwareUpdateServerByID(ctx context.Context, id int) (*ResourceSoftwareUpdateServer, *interfaces.Response, error)

		// GetSoftwareUpdateServerByName returns the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyname
		GetSoftwareUpdateServerByName(ctx context.Context, name string) (*ResourceSoftwareUpdateServer, *interfaces.Response, error)

		// CreateSoftwareUpdateServer creates a new software update server.
		//
		// Returns the created software update server with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createsoftwareupdateserver
		CreateSoftwareUpdateServer(ctx context.Context, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error)

		// UpdateSoftwareUpdateServerByID updates the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyid
		UpdateSoftwareUpdateServerByID(ctx context.Context, id int, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error)

		// UpdateSoftwareUpdateServerByName updates the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyname
		UpdateSoftwareUpdateServerByName(ctx context.Context, name string, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error)

		// DeleteSoftwareUpdateServerByID removes the specified software update server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyid
		DeleteSoftwareUpdateServerByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteSoftwareUpdateServerByName removes the specified software update server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyname
		DeleteSoftwareUpdateServerByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the software update server-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/softwareupdateservers
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SoftwareUpdateServersServiceInterface = (*Service)(nil)

// NewService returns a new software update servers Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Software Update Servers CRUD Operations
// -----------------------------------------------------------------------------

// ListSoftwareUpdateServers returns all software update servers.
// URL: GET /JSSResource/softwareupdateservers
// https://developer.jamf.com/jamf-pro/reference/findallsoftwareupdateservers
func (s *Service) ListSoftwareUpdateServers(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicSoftwareUpdateServers, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSoftwareUpdateServerByID returns the specified software update server by ID.
// URL: GET /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyid
func (s *Service) GetSoftwareUpdateServerByID(ctx context.Context, id int) (*ResourceSoftwareUpdateServer, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSoftwareUpdateServerByName returns the specified software update server by name.
// URL: GET /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findsoftwareupdateserversbyname
func (s *Service) GetSoftwareUpdateServerByName(ctx context.Context, name string) (*ResourceSoftwareUpdateServer, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateSoftwareUpdateServer creates a new software update server.
// URL: POST /JSSResource/softwareupdateservers/id/0
// Returns the created software update server with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createsoftwareupdateserver
func (s *Service) CreateSoftwareUpdateServer(ctx context.Context, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicSoftwareUpdateServers)

	var result ResourceSoftwareUpdateServer

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSoftwareUpdateServerByID updates the specified software update server by ID.
// URL: PUT /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyid
func (s *Service) UpdateSoftwareUpdateServerByID(ctx context.Context, id int, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("software update server ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

	var result ResourceSoftwareUpdateServer

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSoftwareUpdateServerByName updates the specified software update server by name.
// URL: PUT /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatesoftwareupdateserverbyname
func (s *Service) UpdateSoftwareUpdateServerByName(ctx context.Context, name string, req *RequestSoftwareUpdateServer) (*ResourceSoftwareUpdateServer, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("software update server name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

	var result ResourceSoftwareUpdateServer

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSoftwareUpdateServerByID removes the specified software update server by ID.
// URL: DELETE /JSSResource/softwareupdateservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyid
func (s *Service) DeleteSoftwareUpdateServerByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("software update server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicSoftwareUpdateServers, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteSoftwareUpdateServerByName removes the specified software update server by name.
// URL: DELETE /JSSResource/softwareupdateservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletesoftwareupdateserverbyname
func (s *Service) DeleteSoftwareUpdateServerByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("software update server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicSoftwareUpdateServers, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
