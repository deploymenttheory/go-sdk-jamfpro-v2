package mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile-device-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledevicegroups
	MobileDeviceGroups struct {
		client transport.HTTPClient
	}
)

// NewService returns a new mobile device groups Service backed by the provided HTTP client.
func NewMobileDeviceGroups(client transport.HTTPClient) *MobileDeviceGroups {
	return &MobileDeviceGroups{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Groups CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device groups.
//
// URL: GET /JSSResource/mobiledevicegroups
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicegroups
func (s *MobileDeviceGroups) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicMobileDeviceGroups

	var out ListResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) GetByID(ctx context.Context, id int) (*ResourceMobileDeviceGroup, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceGroups, id)

	var out ResourceMobileDeviceGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) GetByName(ctx context.Context, name string) (*ResourceMobileDeviceGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceGroups, name)

	var out ResourceMobileDeviceGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) Create(ctx context.Context, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMobileDeviceGroups)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) UpdateByID(ctx context.Context, id int, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceGroups, id)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) UpdateByName(ctx context.Context, name string, req *RequestMobileDeviceGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device group name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("mobile device group name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceGroups, name)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *MobileDeviceGroups) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceGroups, id)

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

// DeleteByName removes the specified mobile device group by name.
//
// URL: DELETE /JSSResource/mobiledevicegroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicegroupbyname
func (s *MobileDeviceGroups) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceGroups, name)

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
