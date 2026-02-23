package mobile_device_enrollment_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API mobile device enrollment profile operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceenrollmentprofiles
	ServiceInterface interface {
		// List returns all mobile device enrollment profiles.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofiles
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified mobile device enrollment profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyid
		GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error)

		// GetByName returns the specified mobile device enrollment profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyname
		GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error)

		// GetByInvitation returns the specified mobile device enrollment profile by invitation.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyinvitation
		GetByInvitation(ctx context.Context, invitation string) (*Resource, *interfaces.Response, error)

		// GetByIDWithSubset returns a specific subset of a mobile device enrollment profile by ID.
		// Subset values: General, Location, Purchasing, Attachments.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyiddatasubset
		GetByIDWithSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error)

		// GetByNameWithSubset returns a specific subset of a mobile device enrollment profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbynamedatasubset
		GetByNameWithSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error)

		// Create creates a new mobile device enrollment profile.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceenrollmentprofilebyid
		Create(ctx context.Context, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByID updates the specified mobile device enrollment profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyid
		UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByName updates the specified mobile device enrollment profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyname
		UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByInvitation updates the specified mobile device enrollment profile by invitation.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyinvitation
		UpdateByInvitation(ctx context.Context, invitation string, req *Resource) (*Resource, *interfaces.Response, error)

		// DeleteByID removes the specified mobile device enrollment profile by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified mobile device enrollment profile by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)

		// DeleteByInvitation removes the specified mobile device enrollment profile by invitation.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyinvitation
		DeleteByInvitation(ctx context.Context, invitation string) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device enrollment profiles-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceenrollmentprofiles
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new mobile device enrollment profiles Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Enrollment Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device enrollment profiles.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofiles
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointMobileDeviceEnrollmentProfiles

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

// GetByID returns the specified mobile device enrollment profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceEnrollmentProfiles, id)

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

// GetByName returns the specified mobile device enrollment profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceEnrollmentProfiles, name)

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

// GetByInvitation returns the specified mobile device enrollment profile by invitation.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyinvitation
func (s *Service) GetByInvitation(ctx context.Context, invitation string) (*Resource, *interfaces.Response, error) {
	if invitation == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", EndpointMobileDeviceEnrollmentProfiles, invitation)

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

// GetByIDWithSubset returns a specific subset of a mobile device enrollment profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbyiddatasubset
func (s *Service) GetByIDWithSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", EndpointMobileDeviceEnrollmentProfiles, id, subset)

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

// GetByNameWithSubset returns a specific subset of a mobile device enrollment profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceenrollmentprofilesbynamedatasubset
func (s *Service) GetByNameWithSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", EndpointMobileDeviceEnrollmentProfiles, name, subset)

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

// Create creates a new mobile device enrollment profile.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceenrollmentprofilebyid
func (s *Service) Create(ctx context.Context, req *Resource) (*Resource, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointMobileDeviceEnrollmentProfiles)

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

// UpdateByID updates the specified mobile device enrollment profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceEnrollmentProfiles, id)

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

// UpdateByName updates the specified mobile device enrollment profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceEnrollmentProfiles, name)

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

// UpdateByInvitation updates the specified mobile device enrollment profile by invitation.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceenrollmentprofilebyinvitation
func (s *Service) UpdateByInvitation(ctx context.Context, invitation string, req *Resource) (*Resource, *interfaces.Response, error) {
	if invitation == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device enrollment profile name is required in request")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", EndpointMobileDeviceEnrollmentProfiles, invitation)

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

// DeleteByID removes the specified mobile device enrollment profile by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device enrollment profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMobileDeviceEnrollmentProfiles, id)

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

// DeleteByName removes the specified mobile device enrollment profile by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device enrollment profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMobileDeviceEnrollmentProfiles, name)

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

// DeleteByInvitation removes the specified mobile device enrollment profile by invitation.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceenrollmentprofilebyinvitation
func (s *Service) DeleteByInvitation(ctx context.Context, invitation string) (*interfaces.Response, error) {
	if invitation == "" {
		return nil, fmt.Errorf("mobile device enrollment profile invitation cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", EndpointMobileDeviceEnrollmentProfiles, invitation)

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
