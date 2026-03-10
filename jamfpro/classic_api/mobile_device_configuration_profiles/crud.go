package mobile_device_configuration_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/plist"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device configuration profiles Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/mobiledeviceconfigurationprofiles
	MobileDeviceConfigurationProfiles struct {
		client client.Client
	}
)

// NewService returns a new mobile device configuration profiles Service backed by the provided HTTP client.
func NewMobileDeviceConfigurationProfiles(client client.Client) *MobileDeviceConfigurationProfiles {
	return &MobileDeviceConfigurationProfiles{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Configuration Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all mobile device configuration profiles.
//
// URL: GET /JSSResource/mobiledeviceconfigurationprofiles
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofiles
func (s *MobileDeviceConfigurationProfiles) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicMobileDeviceConfigurationProfiles

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

// GetByID returns the specified mobile device configuration profile by ID.
//
// URL: GET /JSSResource/mobiledeviceconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofilesbyid
func (s *MobileDeviceConfigurationProfiles) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device configuration profile ID must be a positive integer")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceConfigurationProfiles, id)

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

// GetByName returns the specified mobile device configuration profile by name.
//
// URL: GET /JSSResource/mobiledeviceconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofilesbyname
func (s *MobileDeviceConfigurationProfiles) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name cannot be empty")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceConfigurationProfiles, name)

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

// GetByIDWithSubset returns the specified mobile device configuration profile by ID with a data subset.
//
// URL: GET /JSSResource/mobiledeviceconfigurationprofiles/id/{id}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofilesbyiddatasubset
func (s *MobileDeviceConfigurationProfiles) GetByIDWithSubset(ctx context.Context, id int, subset string) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device configuration profile ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", constants.EndpointClassicMobileDeviceConfigurationProfiles, id, subset)

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

// GetByNameWithSubset returns the specified mobile device configuration profile by name with a data subset.
//
// URL: GET /JSSResource/mobiledeviceconfigurationprofiles/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmobiledeviceconfigurationprofilesbynamedatasubset
func (s *MobileDeviceConfigurationProfiles) GetByNameWithSubset(ctx context.Context, name, subset string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset cannot be empty")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", constants.EndpointClassicMobileDeviceConfigurationProfiles, name, subset)

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

// Create creates a new mobile device configuration profile.
//
// URL: POST /JSSResource/mobiledeviceconfigurationprofiles/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledeviceconfigurationprofilebyid
func (s *MobileDeviceConfigurationProfiles) Create(ctx context.Context, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name is required")
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMobileDeviceConfigurationProfiles)

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

// UpdateByID updates the specified mobile device configuration profile by ID with UUID preservation.
//
// This method automatically:
// 1. Fetches the existing profile from Jamf Pro
// 2. Extracts PayloadUUID and PayloadIdentifier from the existing plist
// 3. Injects them into the new plist to maintain UUID continuity
// 4. Sends the update request with preserved UUIDs
//
// Jamf Pro modifies the top-level PayloadUUID and PayloadIdentifier upon profile creation.
// If these UUIDs are not preserved during updates, Jamf Pro treats the update as a brand
// new plist, which can cause configuration issues.
//
// URL: PUT /JSSResource/mobiledeviceconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceconfigurationprofilebyid
func (s *MobileDeviceConfigurationProfiles) UpdateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mobile device configuration profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name is required")
	}

	if req.General.Payloads != "" {
		existingProfile, _, err := s.GetByID(ctx, id)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get existing profile for UUID preservation: %w", err)
		}

		if existingProfile.General.Payloads != "" {
			updatedPayloads, err := plist.PreservePlistUUIDs(existingProfile.General.Payloads, req.General.Payloads)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to preserve plist UUIDs: %w", err)
			}
			req.General.Payloads = updatedPayloads
		}
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceConfigurationProfiles, id)

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

// UpdateByName updates the specified mobile device configuration profile by name with UUID preservation.
//
// This method automatically:
// 1. Fetches the existing profile from Jamf Pro
// 2. Extracts PayloadUUID and PayloadIdentifier from the existing plist
// 3. Injects them into the new plist to maintain UUID continuity
// 4. Sends the update request with preserved UUIDs
//
// Jamf Pro modifies the top-level PayloadUUID and PayloadIdentifier upon profile creation.
// If these UUIDs are not preserved during updates, Jamf Pro treats the update as a brand
// new plist, which can cause configuration issues.
//
// URL: PUT /JSSResource/mobiledeviceconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemobiledeviceconfigurationprofilebyname
func (s *MobileDeviceConfigurationProfiles) UpdateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mobile device configuration profile name is required in request")
	}

	if req.General.Payloads != "" {
		existingProfile, _, err := s.GetByName(ctx, name)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get existing profile for UUID preservation: %w", err)
		}

		if existingProfile.General.Payloads != "" {
			updatedPayloads, err := plist.PreservePlistUUIDs(existingProfile.General.Payloads, req.General.Payloads)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to preserve plist UUIDs: %w", err)
			}
			req.General.Payloads = updatedPayloads
		}
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceConfigurationProfiles, name)

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

// DeleteByID removes the specified mobile device configuration profile by ID.
//
// URL: DELETE /JSSResource/mobiledeviceconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceconfigurationprofilebyid
func (s *MobileDeviceConfigurationProfiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mobile device configuration profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMobileDeviceConfigurationProfiles, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified mobile device configuration profile by name.
//
// URL: DELETE /JSSResource/mobiledeviceconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemobiledeviceconfigurationprofilebyname
func (s *MobileDeviceConfigurationProfiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mobile device configuration profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMobileDeviceConfigurationProfiles, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
