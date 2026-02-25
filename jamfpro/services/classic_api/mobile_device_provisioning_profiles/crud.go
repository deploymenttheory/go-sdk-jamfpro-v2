package mobile_device_provisioning_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API mobile device provisioning profile operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
	ServiceInterface interface {
		// List returns all mobile device provisioning profiles.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofiles
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified mobile device provisioning profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyid
		GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error)

		// GetByName returns the specified mobile device provisioning profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyname
		GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error)

		// GetByUUID returns the specified mobile device provisioning profile by UUID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofilesbyuuid
		GetByUUID(ctx context.Context, uuid string) (*Resource, *interfaces.Response, error)

		// CreateByID creates a new mobile device provisioning profile by ID (use 0 for new).
		// Returns the created profile.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyid
		CreateByID(ctx context.Context, id int, req *RequestResource) (*Resource, *interfaces.Response, error)

		// CreateByName creates a new mobile device provisioning profile by name.
		// Returns the created profile.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyname
		CreateByName(ctx context.Context, name string, req *RequestResource) (*Resource, *interfaces.Response, error)

		// CreateByUUID creates a new mobile device provisioning profile by UUID.
		// Returns the created profile.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceprovisioningprofilebyuuid
		CreateByUUID(ctx context.Context, uuid string, req *RequestResource) (*Resource, *interfaces.Response, error)

		// UpdateByID updates the specified mobile device provisioning profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyid
		UpdateByID(ctx context.Context, id int, req *RequestResource) (*Resource, *interfaces.Response, error)

		// UpdateByName updates the specified mobile device provisioning profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyname
		UpdateByName(ctx context.Context, name string, req *RequestResource) (*Resource, *interfaces.Response, error)

		// UpdateByUUID updates the specified mobile device provisioning profile by UUID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceprovisioningprofilebyuuid
		UpdateByUUID(ctx context.Context, uuid string, req *RequestResource) (*Resource, *interfaces.Response, error)

		// DeleteByID removes the specified mobile device provisioning profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified mobile device provisioning profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)

		// DeleteByUUID removes the specified mobile device provisioning profile by UUID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceprovisioningprofilebyuuid
		DeleteByUUID(ctx context.Context, uuid string) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device provisioning profiles Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new mobile device provisioning profiles Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Provisioning Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device provisioning profiles.
//
// URL: GET /JSSResource/mobiledeviceprovisioningprofiles
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceprovisioningprofiles
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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
func (s *Service) GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error) {
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
func (s *Service) GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error) {
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
func (s *Service) GetByUUID(ctx context.Context, uuid string) (*Resource, *interfaces.Response, error) {
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
func (s *Service) CreateByID(ctx context.Context, id int, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) CreateByName(ctx context.Context, name string, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) CreateByUUID(ctx context.Context, uuid string, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) UpdateByUUID(ctx context.Context, uuid string, req *RequestResource) (*Resource, *interfaces.Response, error) {
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

	var out Resource

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
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
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
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
func (s *Service) DeleteByUUID(ctx context.Context, uuid string) (*interfaces.Response, error) {
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
