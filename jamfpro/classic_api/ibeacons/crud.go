package ibeacons

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the iBeacon-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ibeacons
	Ibeacons struct {
		client client.Client
	}
)

// NewService returns a new iBeacons Service backed by the provided HTTP client.
func NewIbeacons(client client.Client) *Ibeacons {
	return &Ibeacons{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - iBeacons CRUD Operations
// -----------------------------------------------------------------------------

// List returns all iBeacons.
// URL: GET /JSSResource/ibeacons
// https://developer.jamf.com/jamf-pro/reference/findibeacons
func (s *Ibeacons) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicIBeacons

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

// GetByID returns the specified iBeacon by ID.
// URL: GET /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findibeaconsbyid
func (s *Ibeacons) GetByID(ctx context.Context, id int) (*ResourceIBeacon, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicIBeacons, id)

	var result ResourceIBeacon

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

// GetByName returns the specified iBeacon by name.
// URL: GET /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findibeaconsbyname
func (s *Ibeacons) GetByName(ctx context.Context, name string) (*ResourceIBeacon, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("iBeacon name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicIBeacons, name)

	var result ResourceIBeacon

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

// Create creates a new iBeacon.
// URL: POST /JSSResource/ibeacons/id/0
// Returns the created iBeacon with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createibeaconbyid
func (s *Ibeacons) Create(ctx context.Context, req *RequestIBeacon) (*ResourceIBeacon, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicIBeacons)

	var result ResourceIBeacon

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

// UpdateByID updates the specified iBeacon by ID.
// URL: PUT /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateibeaconbyid
func (s *Ibeacons) UpdateByID(ctx context.Context, id int, req *RequestIBeacon) (*ResourceIBeacon, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicIBeacons, id)

	var result ResourceIBeacon

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

// UpdateByName updates the specified iBeacon by name.
// URL: PUT /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateibeaconbyname
func (s *Ibeacons) UpdateByName(ctx context.Context, name string, req *RequestIBeacon) (*ResourceIBeacon, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("iBeacon name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicIBeacons, name)

	var result ResourceIBeacon

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

// DeleteByID removes the specified iBeacon by ID.
// URL: DELETE /JSSResource/ibeacons/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyid
func (s *Ibeacons) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("iBeacon ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicIBeacons, id)

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

// DeleteByName removes the specified iBeacon by name.
// URL: DELETE /JSSResource/ibeacons/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteibeaconbyname
func (s *Ibeacons) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("iBeacon name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicIBeacons, name)

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
