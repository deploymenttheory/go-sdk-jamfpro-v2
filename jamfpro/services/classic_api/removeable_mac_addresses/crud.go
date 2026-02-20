package removeable_mac_addresses

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// RemoveableMacAddressesServiceInterface defines the interface for Classic API removeable MAC address operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
	RemoveableMacAddressesServiceInterface interface {
		// ListRemoveableMacAddresses returns all removeable MAC addresses.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
		ListRemoveableMacAddresses(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetRemoveableMacAddressByID returns the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyid
		GetRemoveableMacAddressByID(ctx context.Context, id int) (*ResourceRemoveableMacAddress, *interfaces.Response, error)

		// GetRemoveableMacAddressByName returns the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyname
		GetRemoveableMacAddressByName(ctx context.Context, name string) (*ResourceRemoveableMacAddress, *interfaces.Response, error)

		// CreateRemoveableMacAddress creates a new removeable MAC address.
		//
		// Returns the created removeable MAC address with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createremovablemacaddressbyid
		CreateRemoveableMacAddress(ctx context.Context, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error)

		// UpdateRemoveableMacAddressByID updates the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyid
		UpdateRemoveableMacAddressByID(ctx context.Context, id int, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error)

		// UpdateRemoveableMacAddressByName updates the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyname
		UpdateRemoveableMacAddressByName(ctx context.Context, name string, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error)

		// DeleteRemoveableMacAddressByID removes the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyid
		DeleteRemoveableMacAddressByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteRemoveableMacAddressByName removes the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyname
		DeleteRemoveableMacAddressByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the removeable MAC addresses-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ RemoveableMacAddressesServiceInterface = (*Service)(nil)

// NewService returns a new removeable MAC addresses Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Removeable MAC Addresses CRUD Operations
// -----------------------------------------------------------------------------

// ListRemoveableMacAddresses returns all removeable MAC addresses.
// URL: GET /JSSResource/removablemacaddresses
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
func (s *Service) ListRemoveableMacAddresses(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicRemoveableMacAddresses

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

// GetRemoveableMacAddressByID returns the specified removeable MAC address by ID.
// URL: GET /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyid
func (s *Service) GetRemoveableMacAddressByID(ctx context.Context, id int) (*ResourceRemoveableMacAddress, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRemoveableMacAddresses, id)

	var result ResourceRemoveableMacAddress

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

// GetRemoveableMacAddressByName returns the specified removeable MAC address by name.
// URL: GET /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyname
func (s *Service) GetRemoveableMacAddressByName(ctx context.Context, name string) (*ResourceRemoveableMacAddress, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("removeable MAC address name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRemoveableMacAddresses, name)

	var result ResourceRemoveableMacAddress

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

// CreateRemoveableMacAddress creates a new removeable MAC address.
// URL: POST /JSSResource/removablemacaddresses/id/0
// Returns the created removeable MAC address with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createremovablemacaddressbyid
func (s *Service) CreateRemoveableMacAddress(ctx context.Context, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicRemoveableMacAddresses)

	var result ResourceRemoveableMacAddress

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

// UpdateRemoveableMacAddressByID updates the specified removeable MAC address by ID.
// URL: PUT /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyid
func (s *Service) UpdateRemoveableMacAddressByID(ctx context.Context, id int, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRemoveableMacAddresses, id)

	var result ResourceRemoveableMacAddress

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

// UpdateRemoveableMacAddressByName updates the specified removeable MAC address by name.
// URL: PUT /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyname
func (s *Service) UpdateRemoveableMacAddressByName(ctx context.Context, name string, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("removeable MAC address name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRemoveableMacAddresses, name)

	var result ResourceRemoveableMacAddress

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

// DeleteRemoveableMacAddressByID removes the specified removeable MAC address by ID.
// URL: DELETE /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyid
func (s *Service) DeleteRemoveableMacAddressByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("removeable MAC address ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRemoveableMacAddresses, id)

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

// DeleteRemoveableMacAddressByName removes the specified removeable MAC address by name.
// URL: DELETE /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyname
func (s *Service) DeleteRemoveableMacAddressByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("removeable MAC address name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRemoveableMacAddresses, name)

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
