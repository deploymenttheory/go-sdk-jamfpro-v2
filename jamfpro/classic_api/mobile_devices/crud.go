package mobile_devices

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile-devices-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledevices
	MobileDevices struct {
		client client.Client
	}
)

// NewService returns a new mobile devices Service backed by the provided HTTP client.
func NewMobileDevices(client client.Client) *MobileDevices {
	return &MobileDevices{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Devices CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile devices.
//
// URL: GET /JSSResource/mobiledevices
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevices
func (s *MobileDevices) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicMobileDevices

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByID returns the specified mobile device by ID.
//
// URL: GET /JSSResource/mobiledevices/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicesbyid
func (s *MobileDevices) GetByID(ctx context.Context, id string) (*ResponseMobileDevice, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device ID cannot be empty")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicMobileDevices, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByName returns the specified mobile device by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicesbyname
func (s *MobileDevices) GetByName(ctx context.Context, name string) (*ResponseMobileDevice, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device name cannot be empty")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDevices, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByIDAndDataSubset returns a specific subset of data for the mobile device by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicesbyid
func (s *MobileDevices) GetByIDAndDataSubset(ctx context.Context, id, subset string) (*ResponseMobileDevice, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device ID cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("data subset cannot be empty")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/id/%s/subset/%s", constants.EndpointClassicMobileDevices, id, subset)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByNameAndDataSubset returns a specific subset of data for the mobile device by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledevicesbyname
func (s *MobileDevices) GetByNameAndDataSubset(ctx context.Context, name, subset string) (*ResponseMobileDevice, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("data subset cannot be empty")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", constants.EndpointClassicMobileDevices, name, subset)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// Create creates a new mobile device.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledevicebyid
func (s *MobileDevices) Create(ctx context.Context, device *ResponseMobileDevice) (*ResponseMobileDevice, *resty.Response, error) {
	if device == nil {
		return nil, nil, fmt.Errorf("mobile device is required")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMobileDevices)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(device).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByID updates the specified mobile device by ID.
//
// URL: PUT /JSSResource/mobiledevices/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicebyid
func (s *MobileDevices) UpdateByID(ctx context.Context, id string, device *ResponseMobileDevice) (*ResponseMobileDevice, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device ID cannot be empty")
	}
	if device == nil {
		return nil, nil, fmt.Errorf("mobile device is required")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicMobileDevices, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(device).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByName updates the specified mobile device by name.
//
// URL: PUT /JSSResource/mobiledevices/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledevicebyname
func (s *MobileDevices) UpdateByName(ctx context.Context, name string, device *ResponseMobileDevice) (*ResponseMobileDevice, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device name cannot be empty")
	}
	if device == nil {
		return nil, nil, fmt.Errorf("mobile device is required")
	}

	var out ResponseMobileDevice

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDevices, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(device).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// DeleteByID removes the specified mobile device by ID.
//
// URL: DELETE /JSSResource/mobiledevices/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicebyid
func (s *MobileDevices) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("mobile device ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicMobileDevices, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified mobile device by name.
//
// URL: DELETE /JSSResource/mobiledevices/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledevicebyname
func (s *MobileDevices) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDevices, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
