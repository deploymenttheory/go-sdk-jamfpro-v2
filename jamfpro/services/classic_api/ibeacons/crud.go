package ibeacons

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// IBeaconsServiceInterface defines the interface for Classic API iBeacon operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ibeacons
	IBeaconsServiceInterface interface {
		// ListIBeacons returns all iBeacons.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallibeacons
		ListIBeacons(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetIBeaconByID returns the specified iBeacon by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findibeaconsbyid
		GetIBeaconByID(ctx context.Context, id int) (*ResourceIBeacon, *interfaces.Response, error)

		// GetIBeaconByName returns the specified iBeacon by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findibeaconsbyname
		GetIBeaconByName(ctx context.Context, name string) (*ResourceIBeacon, *interfaces.Response, error)

		// CreateIBeacon creates a new iBeacon.
		//
		// Returns the created iBeacon with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createibeaconbyid
		CreateIBeacon(ctx context.Context, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error)

		// UpdateIBeaconByID updates the specified iBeacon by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateibeaconbyid
		UpdateIBeaconByID(ctx context.Context, id int, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error)

		// UpdateIBeaconByName updates the specified iBeacon by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateibeaconbyname
		UpdateIBeaconByName(ctx context.Context, name string, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error)

		// DeleteIBeaconByID removes the specified iBeacon by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyid
		DeleteIBeaconByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteIBeaconByName removes the specified iBeacon by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyname
		DeleteIBeaconByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the iBeacon-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ibeacons
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ IBeaconsServiceInterface = (*Service)(nil)

// NewService returns a new iBeacons Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - iBeacons CRUD Operations
// -----------------------------------------------------------------------------

// ListIBeacons returns all iBeacons.
// URL: GET /JSSResource/ibeacons
// https://developer.jamf.com/jamf-pro/reference/findallibeacons
func (s *Service) ListIBeacons(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicIBeacons, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetIBeaconByID returns the specified iBeacon by ID.
// URL: GET /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findibeaconsbyid
func (s *Service) GetIBeaconByID(ctx context.Context, id int) (*ResourceIBeacon, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicIBeacons, id)

	var result ResourceIBeacon

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetIBeaconByName returns the specified iBeacon by name.
// URL: GET /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findibeaconsbyname
func (s *Service) GetIBeaconByName(ctx context.Context, name string) (*ResourceIBeacon, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("iBeacon name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicIBeacons, name)

	var result ResourceIBeacon

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateIBeacon creates a new iBeacon.
// URL: POST /JSSResource/ibeacons/id/0
// Returns the created iBeacon with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createibeaconbyid
func (s *Service) CreateIBeacon(ctx context.Context, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicIBeacons)

	var result ResourceIBeacon

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateIBeaconByID updates the specified iBeacon by ID.
// URL: PUT /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateibeaconbyid
func (s *Service) UpdateIBeaconByID(ctx context.Context, id int, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicIBeacons, id)

	var result ResourceIBeacon

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateIBeaconByName updates the specified iBeacon by name.
// URL: PUT /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateibeaconbyname
func (s *Service) UpdateIBeaconByName(ctx context.Context, name string, req *RequestIBeacon) (*ResourceIBeacon, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("iBeacon name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicIBeacons, name)

	var result ResourceIBeacon

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteIBeaconByID removes the specified iBeacon by ID.
// URL: DELETE /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyid
func (s *Service) DeleteIBeaconByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicIBeacons, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteIBeaconByName removes the specified iBeacon by name.
// URL: DELETE /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyname
func (s *Service) DeleteIBeaconByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("iBeacon name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicIBeacons, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
