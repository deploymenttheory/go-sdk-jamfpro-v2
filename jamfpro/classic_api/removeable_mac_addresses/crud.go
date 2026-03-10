package removeable_mac_addresses

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the removeable MAC addresses-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
	RemoveableMacAddresses struct {
		client client.Client
	}
)

// NewService returns a new removeable MAC addresses Service backed by the provided HTTP client.
func NewRemoveableMacAddresses(client client.Client) *RemoveableMacAddresses {
	return &RemoveableMacAddresses{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Removeable MAC Addresses CRUD Operations
// -----------------------------------------------------------------------------

// List returns all removeable MAC addresses.
// URL: GET /JSSResource/removablemacaddresses
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
func (s *RemoveableMacAddresses) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicRemoveableMacAddresses

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

// GetByID returns the specified removeable MAC address by ID.
// URL: GET /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyid
func (s *RemoveableMacAddresses) GetByID(ctx context.Context, id int) (*ResourceRemoveableMacAddress, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicRemoveableMacAddresses, id)

	var result ResourceRemoveableMacAddress

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

// GetByName returns the specified removeable MAC address by name.
// URL: GET /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyname
func (s *RemoveableMacAddresses) GetByName(ctx context.Context, name string) (*ResourceRemoveableMacAddress, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("removeable MAC address name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicRemoveableMacAddresses, name)

	var result ResourceRemoveableMacAddress

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

// Create creates a new removeable MAC address.
// URL: POST /JSSResource/removablemacaddresses/id/0
// Returns the created removeable MAC address with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createremovablemacaddressbyid
func (s *RemoveableMacAddresses) Create(ctx context.Context, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicRemoveableMacAddresses)

	var result ResourceRemoveableMacAddress

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

// UpdateByID updates the specified removeable MAC address by ID.
// URL: PUT /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyid
func (s *RemoveableMacAddresses) UpdateByID(ctx context.Context, id int, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicRemoveableMacAddresses, id)

	var result ResourceRemoveableMacAddress

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

// UpdateByName updates the specified removeable MAC address by name.
// URL: PUT /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyname
func (s *RemoveableMacAddresses) UpdateByName(ctx context.Context, name string, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("removeable MAC address name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicRemoveableMacAddresses, name)

	var result ResourceRemoveableMacAddress

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

// DeleteByID removes the specified removeable MAC address by ID.
// URL: DELETE /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyid
func (s *RemoveableMacAddresses) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicRemoveableMacAddresses, id)

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

// DeleteByName removes the specified removeable MAC address by name.
// URL: DELETE /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyname
func (s *RemoveableMacAddresses) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("removeable MAC address name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicRemoveableMacAddresses, name)

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
