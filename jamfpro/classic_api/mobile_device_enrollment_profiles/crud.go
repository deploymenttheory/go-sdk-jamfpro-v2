package mobile_device_enrollment_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device enrollment profiles-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceenrollmentprofiles
	MobileDeviceEnrollmentProfiles struct {
		client client.Client
	}
)

// NewService returns a new mobile device enrollment profiles Service backed by the provided HTTP client.
func NewMobileDeviceEnrollmentProfiles(client client.Client) *MobileDeviceEnrollmentProfiles {
	return &MobileDeviceEnrollmentProfiles{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Enrollment Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device enrollment profiles.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofiles
func (s *MobileDeviceEnrollmentProfiles) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicMobileDeviceEnrollmentProfiles

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

// GetByID returns the specified mobile device enrollment profile by ID.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyid
func (s *MobileDeviceEnrollmentProfiles) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceEnrollmentProfiles, id)

	var out Resource

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

// GetByName returns the specified mobile device enrollment profile by name.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyname
func (s *MobileDeviceEnrollmentProfiles) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, name)

	var out Resource

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

// GetByInvitation returns the specified mobile device enrollment profile by invitation.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyinvitation
func (s *MobileDeviceEnrollmentProfiles) GetByInvitation(ctx context.Context, invitation string) (*Resource, *resty.Response, error) {
	if invitation == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, invitation)

	var out Resource

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

// GetByIDWithSubset returns a specific subset of a mobile device enrollment profile by ID.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles/id/{id}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyiddatasubset
func (s *MobileDeviceEnrollmentProfiles) GetByIDWithSubset(ctx context.Context, id int, subset string) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, id, subset)

	var out Resource

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

// GetByNameWithSubset returns a specific subset of a mobile device enrollment profile by name.
//
// URL: GET /JSSResource/mobiledeviceenrollmentprofiles/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbynamedatasubset
func (s *MobileDeviceEnrollmentProfiles) GetByNameWithSubset(ctx context.Context, name, subset string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, name, subset)

	var out Resource

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

// Create creates a new mobile device enrollment profile.
//
// URL: POST /JSSResource/mobiledeviceenrollmentprofiles/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceenrollmentprofilebyid
func (s *MobileDeviceEnrollmentProfiles) Create(ctx context.Context, req *Resource) (*Resource, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMobileDeviceEnrollmentProfiles)

	var out Resource

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

// UpdateByID updates the specified mobile device enrollment profile by ID.
//
// URL: PUT /JSSResource/mobiledeviceenrollmentprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyid
func (s *MobileDeviceEnrollmentProfiles) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceEnrollmentProfiles, id)

	var out Resource

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

// UpdateByName updates the specified mobile device enrollment profile by name.
//
// URL: PUT /JSSResource/mobiledeviceenrollmentprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyname
func (s *MobileDeviceEnrollmentProfiles) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, name)

	var out Resource

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

// UpdateByInvitation updates the specified mobile device enrollment profile by invitation.
//
// URL: PUT /JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyinvitation
func (s *MobileDeviceEnrollmentProfiles) UpdateByInvitation(ctx context.Context, invitation string, req *Resource) (*Resource, *resty.Response, error) {
	if invitation == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, invitation)

	var out Resource

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

// DeleteByID removes the specified mobile device enrollment profile by ID.
//
// URL: DELETE /JSSResource/mobiledeviceenrollmentprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyid
func (s *MobileDeviceEnrollmentProfiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceEnrollmentProfiles, id)

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

// DeleteByName removes the specified mobile device enrollment profile by name.
//
// URL: DELETE /JSSResource/mobiledeviceenrollmentprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyname
func (s *MobileDeviceEnrollmentProfiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, name)

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

// DeleteByInvitation removes the specified mobile device enrollment profile by invitation.
//
// URL: DELETE /JSSResource/mobiledeviceenrollmentprofiles/invitation/{invitation}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyinvitation
func (s *MobileDeviceEnrollmentProfiles) DeleteByInvitation(ctx context.Context, invitation string) (*resty.Response, error) {
	if invitation == "" {
		return nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", constants.EndpointClassicMobileDeviceEnrollmentProfiles, invitation)

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
