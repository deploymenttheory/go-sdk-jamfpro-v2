package mobile_device_provisioning_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device provisioning profiles Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
	MobileDeviceProvisioningProfiles struct {
		client client.Client
	}
)

// NewService returns a new mobile device provisioning profiles Service backed by the provided HTTP client.
func NewMobileDeviceProvisioningProfiles(client client.Client) *MobileDeviceProvisioningProfiles {
	return &MobileDeviceProvisioningProfiles{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Provisioning Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device provisioning profiles.
//
// URL: GET /JSSResource/mobiledeviceprovisioningprofiles
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofiles
func (s *MobileDeviceProvisioningProfiles) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicMobileDeviceProvisioningProfiles

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

// GetByID returns the specified mobile device provisioning profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyid
func (s *MobileDeviceProvisioningProfiles) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id < 0 {
		return nil, nil, fmt.Errorf("mobile device provisioning profile ID must be a non-negative integer")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceProvisioningProfiles, id)

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

// GetByName returns the specified mobile device provisioning profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyname
func (s *MobileDeviceProvisioningProfiles) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, name)

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

// GetByUUID returns the specified mobile device provisioning profile by UUID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyuuid
func (s *MobileDeviceProvisioningProfiles) GetByUUID(ctx context.Context, uuid string) (*Resource, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/uuid/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, uuid)

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

// CreateByID creates a new mobile device provisioning profile by ID (use 0 for new).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyid
func (s *MobileDeviceProvisioningProfiles) CreateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if id < 0 {
		return nil, nil, fmt.Errorf("mobile device provisioning profile ID must be a non-negative integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceProvisioningProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// CreateByName creates a new mobile device provisioning profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyname
func (s *MobileDeviceProvisioningProfiles) CreateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required in request")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// CreateByUUID creates a new mobile device provisioning profile by UUID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyuuid
func (s *MobileDeviceProvisioningProfiles) CreateByUUID(ctx context.Context, uuid string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required in request")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/uuid/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, uuid)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByID updates the specified mobile device provisioning profile by ID.
//
// URL: PUT /JSSResource/mobiledeviceprovisioningprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyid
func (s *MobileDeviceProvisioningProfiles) UpdateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device provisioning profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceProvisioningProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByName updates the specified mobile device provisioning profile by name.
//
// URL: PUT /JSSResource/mobiledeviceprovisioningprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyname
func (s *MobileDeviceProvisioningProfiles) UpdateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required in request")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByUUID updates the specified mobile device provisioning profile by UUID.
//
// URL: PUT /JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyuuid
func (s *MobileDeviceProvisioningProfiles) UpdateByUUID(ctx context.Context, uuid string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name is required in request")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/uuid/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, uuid)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// DeleteByID removes the specified mobile device provisioning profile by ID.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyid
func (s *MobileDeviceProvisioningProfiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device provisioning profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceProvisioningProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified mobile device provisioning profile by name.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyname
func (s *MobileDeviceProvisioningProfiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByUUID removes the specified mobile device provisioning profile by UUID.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyuuid
func (s *MobileDeviceProvisioningProfiles) DeleteByUUID(ctx context.Context, uuid string) (*resty.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/uuid/%s", constants.EndpointClassicMobileDeviceProvisioningProfiles, uuid)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
