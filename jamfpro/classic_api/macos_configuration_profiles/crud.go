package macos_configuration_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/plist"
	"resty.dev/v3"
)

type (
	// Service handles communication with the macOS configuration profiles Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/osxconfigurationprofiles
	MacosConfigurationProfiles struct {
		client client.Client
	}
)

// NewService returns a new macOS configuration profiles Service backed by the provided HTTP client.
func NewMacosConfigurationProfiles(client client.Client) *MacosConfigurationProfiles {
	return &MacosConfigurationProfiles{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - macOS Configuration Profiles CRUD Operations
// -----------------------------------------------------------------------------

// List returns all macOS configuration profiles.
//
// URL: GET /JSSResource/osxconfigurationprofiles
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findosxconfigurationprofiles
func (s *MacosConfigurationProfiles) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicMacOSConfigurationProfiles

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

// GetByID returns the specified macOS configuration profile by ID.
//
// URL: GET /JSSResource/osxconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findosxconfigurationprofilesbyid
func (s *MacosConfigurationProfiles) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("macOS configuration profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacOSConfigurationProfiles, id)

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

// GetByName returns the specified macOS configuration profile by name.
//
// URL: GET /JSSResource/osxconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findosxconfigurationprofilesbyname
func (s *MacosConfigurationProfiles) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("macOS configuration profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacOSConfigurationProfiles, name)

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

// Create creates a new macOS configuration profile.
//
// URL: POST /JSSResource/osxconfigurationprofiles/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createosxconfigurationprofilebyid
func (s *MacosConfigurationProfiles) Create(ctx context.Context, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("macOS configuration profile name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMacOSConfigurationProfiles)

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

// UpdateByID updates the specified macOS configuration profile by ID with UUID preservation.
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
// URL: PUT /JSSResource/osxconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateosxconfigurationprofilebyid
func (s *MacosConfigurationProfiles) UpdateByID(ctx context.Context, id int, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("macOS configuration profile ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("macOS configuration profile name is required")
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

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacOSConfigurationProfiles, id)

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

// UpdateByName updates the specified macOS configuration profile by name with UUID preservation.
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
// URL: PUT /JSSResource/osxconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateosxconfigurationprofilebyname
func (s *MacosConfigurationProfiles) UpdateByName(ctx context.Context, name string, req *RequestResource) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("macOS configuration profile name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("macOS configuration profile name is required in request")
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

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacOSConfigurationProfiles, name)

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

// DeleteByID removes the specified macOS configuration profile by ID.
//
// URL: DELETE /JSSResource/osxconfigurationprofiles/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteosxconfigurationprofilebyid
func (s *MacosConfigurationProfiles) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("macOS configuration profile ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacOSConfigurationProfiles, id)

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

// DeleteByName removes the specified macOS configuration profile by name.
//
// URL: DELETE /JSSResource/osxconfigurationprofiles/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteosxconfigurationprofilebyname
func (s *MacosConfigurationProfiles) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("macOS configuration profile name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacOSConfigurationProfiles, name)

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
