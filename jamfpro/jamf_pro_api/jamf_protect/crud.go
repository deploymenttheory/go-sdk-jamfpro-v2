package jamf_protect

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Protect-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
	JamfProtect struct {
		client client.Client
	}
)

// NewService creates a new Jamf Protect service.
func NewJamfProtect(client client.Client) *JamfProtect {
	return &JamfProtect{client: client}
}

// GetSettingsV1 retrieves the current Jamf Protect integration settings.
// URL: GET /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
func (s *JamfProtect) GetSettingsV1(ctx context.Context) (*ResourceJamfProtectSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProtectV1

	var result ResourceJamfProtectSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get Jamf Protect settings: %w", err)
	}

	return &result, resp, nil
}

// UpdateSettingsV1 updates Jamf Protect integration settings.
// URL: PUT /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-protect
func (s *JamfProtect) UpdateSettingsV1(ctx context.Context, request *RequestJamfProtectSettings) (*ResourceJamfProtectSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect settings request cannot be nil")
	}

	endpoint := constants.EndpointJamfProJamfProtectV1

	var result ResourceJamfProtectSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update Jamf Protect settings: %w", err)
	}

	return &result, resp, nil
}

// RegisterV1 registers a new Jamf Protect integration.
// URL: POST /api/v1/jamf-protect/register
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-register
func (s *JamfProtect) RegisterV1(ctx context.Context, request *RequestJamfProtectRegistration) (*ResourceJamfProtectSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect registration request cannot be nil")
	}

	endpoint := constants.EndpointJamfProJamfProtectRegisterV1

	var result ResourceJamfProtectSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to register Jamf Protect: %w", err)
	}

	return &result, resp, nil
}

// SyncPlansV1 synchronizes Jamf Protect plans from the Protect server.
// URL: POST /api/v1/jamf-protect/plans/sync
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-plans-sync
func (s *JamfProtect) SyncPlansV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProtectPlansV1 + "/sync"

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to sync Jamf Protect plans: %w", err)
	}

	return resp, nil
}

// CreateIntegrationV1 creates a complete Jamf Protect integration.
// This is a composite operation that performs registration, updates settings,
// and syncs plans in a single call.
func (s *JamfProtect) CreateIntegrationV1(ctx context.Context, registration *RequestJamfProtectRegistration, autoInstall bool) (*ResourceJamfProtectSettings, *resty.Response, error) {
	if registration == nil {
		return nil, nil, fmt.Errorf("Jamf Protect registration request cannot be nil")
	}

	// Step 1: Register
	result, resp, err := s.RegisterV1(ctx, registration)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to register during integration creation: %w", err)
	}

	// Step 2: Update settings with autoInstall preference
	settingsRequest := &RequestJamfProtectSettings{
		AutoInstall: autoInstall,
	}
	result, resp, err = s.UpdateSettingsV1(ctx, settingsRequest)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update settings during integration creation: %w", err)
	}

	// Step 3: Sync plans
	resp, err = s.SyncPlansV1(ctx)
	if err != nil {
		return result, resp, fmt.Errorf("failed to sync plans during integration creation: %w", err)
	}

	return result, resp, nil
}

// ListDeploymentTasksV1 retrieves deployment tasks for a specific deployment.
// URL: GET /api/v1/jamf-protect/deployments/{id}/tasks
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-deployments-id-tasks
func (s *JamfProtect) ListDeploymentTasksV1(ctx context.Context, deploymentID string, rsqlQuery map[string]string) (*ListResponseJamfProtectDeploymentTasks, *resty.Response, error) {
	if deploymentID == "" {
		return nil, nil, fmt.Errorf("deployment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/tasks", constants.EndpointJamfProJamfProtectDeploymentsV1, deploymentID)

	var result ListResponseJamfProtectDeploymentTasks

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceJamfProtectDeploymentTask
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect deployment tasks: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// RetryDeploymentTasksV1 retries failed deployment tasks for a deployment.
// URL: POST /api/v1/jamf-protect/deployments/{id}/tasks/retry
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-deployments-id-tasks-retry
func (s *JamfProtect) RetryDeploymentTasksV1(ctx context.Context, deploymentID string) (*resty.Response, error) {
	if deploymentID == "" {
		return nil, fmt.Errorf("deployment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/tasks/retry", constants.EndpointJamfProJamfProtectDeploymentsV1, deploymentID)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to retry Jamf Protect deployment tasks: %w", err)
	}

	return resp, nil
}

// ListHistoryV1 retrieves paginated Jamf Protect history entries.
// URL: GET /api/v1/jamf-protect/history
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-history
func (s *JamfProtect) ListHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectHistory, *resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProtectHistoryV1

	var result ListResponseJamfProtectHistory

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceJamfProtectHistory
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// CreateHistoryNoteV1 creates a new history note for Jamf Protect.
// URL: POST /api/v1/jamf-protect/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-history
func (s *JamfProtect) CreateHistoryNoteV1(ctx context.Context, request *RequestJamfProtectHistoryNote) (*ResourceJamfProtectHistoryCreate, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect history note request cannot be nil")
	}

	endpoint := constants.EndpointJamfProJamfProtectHistoryV1

	var result ResourceJamfProtectHistoryCreate

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create Jamf Protect history note: %w", err)
	}

	return &result, resp, nil
}

// ListPlansV1 retrieves paginated list of Jamf Protect plans.
// URL: GET /api/v1/jamf-protect/plans
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-plans
func (s *JamfProtect) ListPlansV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectPlans, *resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProtectPlansV1

	var result ListResponseJamfProtectPlans

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceJamfProtectPlan
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect plans: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// DeleteIntegrationV1 removes the Jamf Protect integration.
// URL: DELETE /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/delete_v1-jamf-protect
func (s *JamfProtect) DeleteIntegrationV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProtectV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to delete Jamf Protect integration: %w", err)
	}

	return resp, nil
}
