package mobile_device_provisioning_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ServiceInterface defines the interface for Classic API mobile device provisioning profile operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
	ServiceInterface interface {
		// List returns all mobile device provisioning profiles.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofiles
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified mobile device provisioning profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyid
		GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error)

		// GetByName returns the specified mobile device provisioning profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyname
		GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error)

		// GetByUUID returns the specified mobile device provisioning profile by UUID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyuuid
		GetByUUID(ctx context.Context, uuid string) (*Resource, *resty.Response, error)

		// CreateByID creates a new mobile device provisioning profile by ID (use 0 for new).
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyid
		CreateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// CreateByName creates a new mobile device provisioning profile by name.
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyname
		CreateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// CreateByUUID creates a new mobile device provisioning profile by UUID.
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyuuid
		CreateByUUID(ctx context.Context, uuid string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified mobile device provisioning profile by ID.
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyid
		UpdateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified mobile device provisioning profile by name.
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyname
		UpdateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByUUID updates the specified mobile device provisioning profile by UUID.
		// Returns the assigned ID from the Classic API response.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyuuid
		UpdateByUUID(ctx context.Context, uuid string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified mobile device provisioning profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified mobile device provisioning profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)

		// DeleteByUUID removes the specified mobile device provisioning profile by UUID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyuuid
		DeleteByUUID(ctx context.Context, uuid string) (*resty.Response, error)
	}

	// Service handles communication with the mobile device provisioning profiles Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
	MobileDeviceProvisioningProfiles struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*MobileDeviceProvisioningProfiles)(nil)

// NewService returns a new mobile device provisioning profiles Service backed by the provided HTTP client.
func NewMobileDeviceProvisioningProfiles(client interfaces.HTTPClient) *MobileDeviceProvisioningProfiles {
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
	endpoint := EndpointMobileDeviceProvisioningProfiles

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

// GetByID returns the specified mobile device provisioning profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyid
func (s *MobileDeviceProvisioningProfiles) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id < 0 {
		return nil, nil, fmt.Errorf("mobile device provisioning profile ID must be a non-negative integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceProvisioningProfiles, id)

	var out Resource

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

// GetByName returns the specified mobile device provisioning profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyname
func (s *MobileDeviceProvisioningProfiles) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceProvisioningProfiles, name)

	var out Resource

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

// GetByUUID returns the specified mobile device provisioning profile by UUID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyuuid
func (s *MobileDeviceProvisioningProfiles) GetByUUID(ctx context.Context, uuid string) (*Resource, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/uuid/%s", EndpointMobileDeviceProvisioningProfiles, uuid)

	var out Resource

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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceProvisioningProfiles, id)

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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceProvisioningProfiles, name)

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

	endpoint := fmt.Sprintf("%s/uuid/%s", EndpointMobileDeviceProvisioningProfiles, uuid)

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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceProvisioningProfiles, id)

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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceProvisioningProfiles, name)

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

	endpoint := fmt.Sprintf("%s/uuid/%s", EndpointMobileDeviceProvisioningProfiles, uuid)

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

// DeleteByID removes the specified mobile device provisioning profile by ID.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyid
func (s *MobileDeviceProvisioningProfiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device provisioning profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceProvisioningProfiles, id)

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

// DeleteByName removes the specified mobile device provisioning profile by name.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyname
func (s *MobileDeviceProvisioningProfiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device provisioning profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceProvisioningProfiles, name)

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

// DeleteByUUID removes the specified mobile device provisioning profile by UUID.
//
// URL: DELETE /JSSResource/mobiledeviceprovisioningprofiles/uuid/{uuid}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyuuid
func (s *MobileDeviceProvisioningProfiles) DeleteByUUID(ctx context.Context, uuid string) (*resty.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("mobile device provisioning profile UUID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/uuid/%s", EndpointMobileDeviceProvisioningProfiles, uuid)

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
