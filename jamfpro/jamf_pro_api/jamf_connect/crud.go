package jamf_connect

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
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
		GetSettingsV1(ctx context.Context) (*ResourceJamfConnect, *resty.Response, error)

		// ListConfigProfilesV1 lists all Jamf Connect config profiles with pagination support.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
		ListConfigProfilesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetConfigProfileByUUIDV1 retrieves a specific Jamf Connect config profile by UUID.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByUUIDV1(ctx context.Context, uuid string) (*ResourceJamfConnectConfigProfile, *resty.Response, error)

		// GetConfigProfileByIDV1 retrieves a specific Jamf Connect config profile by profile ID.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByIDV1(ctx context.Context, profileID int) (*ResourceJamfConnectConfigProfile, *resty.Response, error)

		// GetConfigProfileByNameV1 retrieves a specific Jamf Connect config profile by name.
		//
		// Helper method that searches through the list of profiles.
		GetConfigProfileByNameV1(ctx context.Context, name string) (*ResourceJamfConnectConfigProfile, *resty.Response, error)

		// UpdateConfigProfileByUUIDV1 updates a Jamf Connect config profile by UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-connect-config-profiles-uuid
		UpdateConfigProfileByUUIDV1(ctx context.Context, uuid string, request *ResourceJamfConnectConfigProfileUpdate) (*ResourceJamfConnectConfigProfile, *resty.Response, error)

		// GetDeploymentTasksByIDV1 retrieves deployment tasks for a specific Jamf Connect deployment.
		//
		// Supports optional RSQL filtering, pagination and sorting via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-deployments-id-tasks
		GetDeploymentTasksByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*DeploymentTasksResponse, *resty.Response, error)

		// GetHistoryV1 retrieves the history for Jamf Connect.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNoteV1 adds a note to the Jamf Connect history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-history
		AddHistoryNoteV1(ctx context.Context, req *RequestAddHistoryNote) (*resty.Response, error)

		// RetryDeploymentTasksByUUIDV1 retries Connect install tasks for specified computers.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-deployments-configprofileuuid-tasks-retry
		RetryDeploymentTasksByUUIDV1(ctx context.Context, uuid string, computerIDs []string) (*resty.Response, error)
	}

	// Service handles communication with the Jamf Connect-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
	JamfConnect struct {
		client transport.HTTPClient
	}
)

var _ JamfConnectServiceInterface = (*JamfConnect)(nil)

func NewJamfConnect(client transport.HTTPClient) *JamfConnect {
	return &JamfConnect{client: client}
}

// GetSettingsV1 retrieves the Jamf Connect settings.
// URL: GET /api/v1/jamf-connect
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect
func (s *JamfConnect) GetSettingsV1(ctx context.Context) (*ResourceJamfConnect, *resty.Response, error) {
	var result ResourceJamfConnect

	endpoint := constants.EndpointJamfProJamfConnectV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ListConfigProfilesV1 lists all Jamf Connect config profiles with pagination support.
// URL: GET /api/v1/jamf-connect/config-profiles
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *JamfConnect) ListConfigProfilesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/config-profiles", constants.EndpointJamfProJamfConnectV1)

	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceJamfConnectConfigProfile
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list jamf connect config profiles: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetConfigProfileByUUIDV1 retrieves a specific Jamf Connect config profile by UUID.
// URL: GET /api/v1/jamf-connect/config-profiles (searches through list)
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-config-profiles
func (s *JamfConnect) GetConfigProfileByUUIDV1(ctx context.Context, uuid string) (*ResourceJamfConnectConfigProfile, *resty.Response, error) {
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
func (s *JamfConnect) GetConfigProfileByIDV1(ctx context.Context, profileID int) (*ResourceJamfConnectConfigProfile, *resty.Response, error) {
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
func (s *JamfConnect) GetConfigProfileByNameV1(ctx context.Context, name string) (*ResourceJamfConnectConfigProfile, *resty.Response, error) {
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
func (s *JamfConnect) UpdateConfigProfileByUUIDV1(ctx context.Context, uuid string, request *ResourceJamfConnectConfigProfileUpdate) (*ResourceJamfConnectConfigProfile, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/config-profiles/%s", constants.EndpointJamfProJamfConnectV1, uuid)

	var result ResourceJamfConnectConfigProfile

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeploymentTasksByIDV1 retrieves deployment tasks for a specific Jamf Connect deployment.
// URL: GET /api/v1/jamf-connect/deployments/{id}/tasks
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-deployments-id-tasks
func (s *JamfConnect) GetDeploymentTasksByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*DeploymentTasksResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("deployment ID is required")
	}

	endpoint := fmt.Sprintf("%s/deployments/%s/tasks", constants.EndpointJamfProJamfConnectV1, id)

	var result DeploymentTasksResponse

	mergePage := func(pageData []byte) error {
		var pageItems []DeploymentTask
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get deployment tasks: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetHistoryV1 retrieves the history for Jamf Connect.
// URL: GET /api/v1/jamf-connect/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-connect-history
func (s *JamfConnect) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProJamfConnectV1)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get jamf connect history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a note to the Jamf Connect history.
// URL: POST /api/v1/jamf-connect/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-history
func (s *JamfConnect) AddHistoryNoteV1(ctx context.Context, req *RequestAddHistoryNote) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProJamfConnectV1)
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RetryDeploymentTasksByUUIDV1 retries Connect install tasks for specified computers.
// URL: POST /api/v1/jamf-connect/deployments/{uuid}/tasks/retry
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-connect-deployments-configprofileuuid-tasks-retry
func (s *JamfConnect) RetryDeploymentTasksByUUIDV1(ctx context.Context, uuid string, computerIDs []string) (*resty.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("uuid is required")
	}

	if len(computerIDs) == 0 {
		return nil, fmt.Errorf("at least one computer ID is required")
	}

	endpoint := fmt.Sprintf("%s/deployments/%s/tasks/retry", constants.EndpointJamfProJamfConnectV1, uuid)

	requestBody := &ResourceJamfConnectTaskRetry{
		IDs: computerIDs,
	}

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, requestBody, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
