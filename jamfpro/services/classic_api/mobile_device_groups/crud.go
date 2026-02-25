package mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// MobileDeviceGroupsServiceInterface defines the interface for Classic API mobile device group operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledevicegroups
	MobileDeviceGroupsServiceInterface interface {
		// List returns all mobile device groups.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroups
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified mobile device group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroupsbyid
		GetByID(ctx context.Context, id int) (*ResourceMobileDeviceGroup, *interfaces.Response, error)

		// GetByName returns the specified mobile device group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroupsbyname
		GetByName(ctx context.Context, name string) (*ResourceMobileDeviceGroup, *interfaces.Response, error)

		// Create creates a new mobile device group.
		//
		// Returns the created mobile device group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledevicegroupbyid
		Create(ctx context.Context, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified mobile device group by ID.
		//
		// Returns the updated mobile device group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicegroupbyid
		UpdateByID(ctx context.Context, id int, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified mobile device group by name.
		//
		// Returns the updated mobile device group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicegroupbyname
		UpdateByName(ctx context.Context, name string, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified mobile device group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicegroupbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified mobile device group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicegroupbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile-device-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledevicegroups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDeviceGroupsServiceInterface = (*Service)(nil)

// NewService returns a new mobile device groups Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Groups CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device groups.
//
// URL: GET /JSSResource/mobiledevicegroups
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroups
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointMobileDeviceGroups

	var out ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByID returns the specified mobile device group by ID.
//
// URL: GET /JSSResource/mobiledevicegroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroupsbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceMobileDeviceGroup, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceGroups, id)

	var out ResourceMobileDeviceGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByName returns the specified mobile device group by name.
//
// URL: GET /JSSResource/mobiledevicegroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroupsbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceMobileDeviceGroup, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceGroups, name)

	var out ResourceMobileDeviceGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new mobile device group.
//
// Returns the created mobile device group ID only (Classic API behavior).
//
// URL: POST /JSSResource/mobiledevicegroups/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledevicegroupbyid
func (s *Service) Create(ctx context.Context, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointMobileDeviceGroups)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified mobile device group by ID.
//
// Returns the updated mobile device group ID only (Classic API behavior).
//
// URL: PUT /JSSResource/mobiledevicegroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicegroupbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceGroups, id)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified mobile device group by name.
//
// Returns the updated mobile device group ID only (Classic API behavior).
//
// URL: PUT /JSSResource/mobiledevicegroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicegroupbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device group name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceGroups, name)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified mobile device group by ID.
//
// URL: DELETE /JSSResource/mobiledevicegroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicegroupbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceGroups, id)

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

// DeleteByName removes the specified mobile device group by name.
//
// URL: DELETE /JSSResource/mobiledevicegroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicegroupbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceGroups, name)

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
