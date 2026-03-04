package removeable_mac_addresses

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// RemoveableMacAddressesServiceInterface defines the interface for Classic API removeable MAC address operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
	RemoveableMacAddressesServiceInterface interface {
		// List returns all removeable MAC addresses.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyid
		GetByID(ctx context.Context, id int) (*ResourceRemoveableMacAddress, *resty.Response, error)

		// GetByName returns the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyname
		GetByName(ctx context.Context, name string) (*ResourceRemoveableMacAddress, *resty.Response, error)

		// Create creates a new removeable MAC address.
		//
		// Returns the created removeable MAC address with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createremovablemacaddressbyid
		Create(ctx context.Context, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error)

		// UpdateByID updates the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyid
		UpdateByID(ctx context.Context, id int, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error)

		// UpdateByName updates the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyname
		UpdateByName(ctx context.Context, name string, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error)

		// DeleteByID removes the specified removeable MAC address by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified removeable MAC address by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
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

// List returns all removeable MAC addresses.
// URL: GET /JSSResource/removablemacaddresses
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddresses
func (s *Service) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
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

// GetByID returns the specified removeable MAC address by ID.
// URL: GET /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceRemoveableMacAddress, *resty.Response, error) {
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

// GetByName returns the specified removeable MAC address by name.
// URL: GET /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findremovablemacaddressesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceRemoveableMacAddress, *resty.Response, error) {
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

// Create creates a new removeable MAC address.
// URL: POST /JSSResource/removablemacaddresses/id/0
// Returns the created removeable MAC address with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createremovablemacaddressbyid
func (s *Service) Create(ctx context.Context, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
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

// UpdateByID updates the specified removeable MAC address by ID.
// URL: PUT /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
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

// UpdateByName updates the specified removeable MAC address by name.
// URL: PUT /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateremovablemacaddressbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestRemoveableMacAddress) (*ResourceRemoveableMacAddress, *resty.Response, error) {
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

// DeleteByID removes the specified removeable MAC address by ID.
// URL: DELETE /JSSResource/removablemacaddresses/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
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

// DeleteByName removes the specified removeable MAC address by name.
// URL: DELETE /JSSResource/removablemacaddresses/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteremovablemacaddressbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
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
