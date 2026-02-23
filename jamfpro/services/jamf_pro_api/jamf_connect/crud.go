package jamf_connect

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfConnectServiceInterface defines the interface for Jamf Connect operations.
	//
	// Manages Jamf Connect settings and configuration profiles for the Jamf Connect app.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
	JamfConnectServiceInterface interface {
		// GetSettingsV1 retrieves the Jamf Connect settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
		GetSettingsV1(ctx context.Context) (*ResourceJamfConnect, *interfaces.Response, error)

		// ListConfigProfilesV1 lists all Jamf Connect config profiles with pagination support.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
		ListConfigProfilesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetConfigProfileByUUIDV1 retrieves a specific Jamf Connect config profile by UUID.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByUUIDV1(ctx context.Context, uuid string) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error)

		// GetConfigProfileByIDV1 retrieves a specific Jamf Connect config profile by profile ID.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByIDV1(ctx context.Context, profileID int) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error)

		// GetConfigProfileByNameV1 retrieves a specific Jamf Connect config profile by name.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByNameV1(ctx context.Context, name string) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error)

		// UpdateConfigProfileByUUIDV1 updates a Jamf Connect config profile by UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-connect-config-profiles-uuid
		UpdateConfigProfileByUUIDV1(ctx context.Context, uuid string, request *ResourceJamfConnectConfigProfileUpdate) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error)

		// RetryDeploymentTasksByUUIDV1 retries Connect install tasks for specified computers.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-deployments-configprofileuuid-tasks-retry
		RetryDeploymentTasksByUUIDV1(ctx context.Context, uuid string, computerIDs []string) (*interfaces.Response, error)
	}

	// Service handles communication with the Jamf Connect-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfConnectServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetSettingsV1 retrieves the Jamf Connect settings.
// URL: GET /api/v1/jamf-connect
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
func (s *Service) GetSettingsV1(ctx context.Context) (*ResourceJamfConnect, *interfaces.Response, error) {
	var result ResourceJamfConnect

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointJamfConnectV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ListConfigProfilesV1 lists all Jamf Connect config profiles with pagination support.
// URL: GET /api/v1/jamf-connect/config-profiles
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *Service) ListConfigProfilesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/config-profiles", EndpointJamfConnectV1)

	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetConfigProfileByUUIDV1 retrieves a specific Jamf Connect config profile by UUID.
// URL: GET /api/v1/jamf-connect/config-profiles (searches through list)
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *Service) GetConfigProfileByUUIDV1(ctx context.Context, uuid string) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	profiles, resp, err := s.ListConfigProfilesV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, profile := range profiles.Results {
		if profile.UUID == uuid {
			return &profile, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("no jamf connect config profile found with UUID: %s", uuid)
}

// GetConfigProfileByIDV1 retrieves a specific Jamf Connect config profile by profile ID.
// URL: GET /api/v1/jamf-connect/config-profiles (searches through list)
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *Service) GetConfigProfileByIDV1(ctx context.Context, profileID int) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error) {
	if profileID <= 0 {
		return nil, nil, fmt.Errorf("profile ID must be greater than 0")
	}

	profiles, resp, err := s.ListConfigProfilesV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, profile := range profiles.Results {
		if profile.ProfileID == profileID {
			return &profile, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("no jamf connect config profile found with ID: %d", profileID)
}

// GetConfigProfileByNameV1 retrieves a specific Jamf Connect config profile by name.
// URL: GET /api/v1/jamf-connect/config-profiles (searches through list)
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *Service) GetConfigProfileByNameV1(ctx context.Context, name string) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	profiles, resp, err := s.ListConfigProfilesV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, profile := range profiles.Results {
		if profile.ProfileName == name {
			return &profile, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("no jamf connect config profile found with name: %s", name)
}

// UpdateConfigProfileByUUIDV1 updates a Jamf Connect config profile by UUID.
// URL: PUT /api/v1/jamf-connect/config-profiles/{uuid}
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-connect-config-profiles-uuid
func (s *Service) UpdateConfigProfileByUUIDV1(ctx context.Context, uuid string, request *ResourceJamfConnectConfigProfileUpdate) (*ResourceJamfConnectConfigProfile, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/config-profiles/%s", EndpointJamfConnectV1, uuid)

	var result ResourceJamfConnectConfigProfile

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RetryDeploymentTasksByUUIDV1 retries Connect install tasks for specified computers.
// URL: POST /api/v1/jamf-connect/deployments/{uuid}/tasks/retry
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-deployments-configprofileuuid-tasks-retry
func (s *Service) RetryDeploymentTasksByUUIDV1(ctx context.Context, uuid string, computerIDs []string) (*interfaces.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("uuid is required")
	}

	if len(computerIDs) == 0 {
		return nil, fmt.Errorf("at least one computer ID is required")
	}

	endpoint := fmt.Sprintf("%s/deployments/%s/tasks/retry", EndpointJamfConnectV1, uuid)

	requestBody := &ResourceJamfConnectTaskRetry{
		IDs: computerIDs,
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, requestBody, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
