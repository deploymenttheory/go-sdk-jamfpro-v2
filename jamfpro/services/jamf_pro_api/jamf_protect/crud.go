package jamf_protect

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

// ServiceInterface defines the interface for Jamf Protect operations.
//
// Jamf Protect integration provides threat prevention and security for macOS devices.
// These endpoints manage Jamf Protect settings, plans, deployments, and integration.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
type ServiceInterface interface {
	// GetSettingsV1 retrieves the current Jamf Protect integration settings.
	//
	// Returns configuration including Protect URL, sync status, API client details,
	// and auto-install settings.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
	GetSettingsV1(ctx context.Context) (*ResourceJamfProtectSettings, *interfaces.Response, error)

	// UpdateSettingsV1 updates Jamf Protect integration settings.
	//
	// Allows modification of settings such as auto-install configuration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-protect
	UpdateSettingsV1(ctx context.Context, request *RequestJamfProtectSettings) (*ResourceJamfProtectSettings, *interfaces.Response, error)

	// RegisterV1 registers a new Jamf Protect integration.
	//
	// Establishes connection with Jamf Protect by providing Protect URL,
	// client ID, and password credentials.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-register
	RegisterV1(ctx context.Context, request *RequestJamfProtectRegistration) (*ResourceJamfProtectSettings, *interfaces.Response, error)

	// SyncPlansV1 synchronizes Jamf Protect plans from the Protect server.
	//
	// Triggers a sync operation to retrieve the latest plans from Jamf Protect.
	// Returns 204 No Content on success.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-plans-sync
	SyncPlansV1(ctx context.Context) (*interfaces.Response, error)

	// CreateIntegrationV1 creates a complete Jamf Protect integration.
	//
	// Composite operation that performs registration, updates settings with
	// auto-install preference, and syncs plans in a single call.
	//
	// This is a convenience method that combines RegisterV1, UpdateSettingsV1,
	// and SyncPlansV1 operations.
	CreateIntegrationV1(ctx context.Context, registration *RequestJamfProtectRegistration, autoInstall bool) (*ResourceJamfProtectSettings, *interfaces.Response, error)

	// ListDeploymentTasksV1 retrieves deployment tasks for a specific deployment.
	//
	// Returns paginated list of deployment tasks with their status, version,
	// and associated computer information. Supports filtering via rsqlQuery.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-deployments-id-tasks
	ListDeploymentTasksV1(ctx context.Context, deploymentID string, rsqlQuery map[string]string) (*ListResponseJamfProtectDeploymentTasks, *interfaces.Response, error)

	// RetryDeploymentTasksV1 retries failed deployment tasks for a deployment.
	//
	// Triggers retry of failed tasks for the specified deployment.
	// Returns 204 No Content on success.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-deployments-id-tasks-retry
	RetryDeploymentTasksV1(ctx context.Context, deploymentID string) (*interfaces.Response, error)

	// ListHistoryV1 retrieves paginated Jamf Protect history entries.
	//
	// Returns audit log entries for Jamf Protect operations including
	// user, date, notes, and details. Supports filtering via rsqlQuery.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-history
	ListHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectHistory, *interfaces.Response, error)

	// CreateHistoryNoteV1 creates a new history note for Jamf Protect.
	//
	// Adds a new audit log entry with note and details to the Jamf Protect history.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-history
	CreateHistoryNoteV1(ctx context.Context, request *RequestJamfProtectHistoryNote) (*ResourceJamfProtectHistory, *interfaces.Response, error)

	// ListPlansV1 retrieves paginated list of Jamf Protect plans.
	//
	// Returns available deployment plans synced from Jamf Protect server.
	// Supports filtering via rsqlQuery.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-plans
	ListPlansV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectPlans, *interfaces.Response, error)

	// DeleteIntegrationV1 removes the Jamf Protect integration.
	//
	// Deletes the configured Jamf Protect integration and all associated settings.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-jamf-protect
	DeleteIntegrationV1(ctx context.Context) (*interfaces.Response, error)
}

type (
	// Service handles communication with the Jamf Protect-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService creates a new Jamf Protect service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetSettingsV1 retrieves the current Jamf Protect integration settings.
// URL: GET /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect
func (s *Service) GetSettingsV1(ctx context.Context) (*ResourceJamfProtectSettings, *interfaces.Response, error) {
	endpoint := EndpointJamfProtectV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceJamfProtectSettings
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get Jamf Protect settings: %w", err)
	}

	return &result, resp, nil
}

// UpdateSettingsV1 updates Jamf Protect integration settings.
// URL: PUT /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-protect
func (s *Service) UpdateSettingsV1(ctx context.Context, request *RequestJamfProtectSettings) (*ResourceJamfProtectSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect settings request cannot be nil")
	}

	endpoint := EndpointJamfProtectV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result ResourceJamfProtectSettings
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update Jamf Protect settings: %w", err)
	}

	return &result, resp, nil
}

// RegisterV1 registers a new Jamf Protect integration.
// URL: POST /api/v1/jamf-protect/register
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-register
func (s *Service) RegisterV1(ctx context.Context, request *RequestJamfProtectRegistration) (*ResourceJamfProtectSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect registration request cannot be nil")
	}

	endpoint := EndpointJamfProtectRegisterV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result ResourceJamfProtectSettings
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to register Jamf Protect: %w", err)
	}

	return &result, resp, nil
}

// SyncPlansV1 synchronizes Jamf Protect plans from the Protect server.
// URL: POST /api/v1/jamf-protect/plans/sync
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-plans-sync
func (s *Service) SyncPlansV1(ctx context.Context) (*interfaces.Response, error) {
	endpoint := EndpointJamfProtectPlansV1 + "/sync"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to sync Jamf Protect plans: %w", err)
	}

	return resp, nil
}

// CreateIntegrationV1 creates a complete Jamf Protect integration.
// This is a composite operation that performs registration, updates settings,
// and syncs plans in a single call.
func (s *Service) CreateIntegrationV1(ctx context.Context, registration *RequestJamfProtectRegistration, autoInstall bool) (*ResourceJamfProtectSettings, *interfaces.Response, error) {
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
func (s *Service) ListDeploymentTasksV1(ctx context.Context, deploymentID string, rsqlQuery map[string]string) (*ListResponseJamfProtectDeploymentTasks, *interfaces.Response, error) {
	if deploymentID == "" {
		return nil, nil, fmt.Errorf("deployment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/tasks", EndpointJamfProtectDeploymentsV1, deploymentID)

	var result ListResponseJamfProtectDeploymentTasks

	mergePage := func(pageData []byte) error {
		var pageResponse ListResponseJamfProtectDeploymentTasks
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect deployment tasks: %w", err)
	}
	return &result, resp, nil
}

// RetryDeploymentTasksV1 retries failed deployment tasks for a deployment.
// URL: POST /api/v1/jamf-protect/deployments/{id}/tasks/retry
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-deployments-id-tasks-retry
func (s *Service) RetryDeploymentTasksV1(ctx context.Context, deploymentID string) (*interfaces.Response, error) {
	if deploymentID == "" {
		return nil, fmt.Errorf("deployment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/tasks/retry", EndpointJamfProtectDeploymentsV1, deploymentID)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to retry Jamf Protect deployment tasks: %w", err)
	}

	return resp, nil
}

// ListHistoryV1 retrieves paginated Jamf Protect history entries.
// URL: GET /api/v1/jamf-protect/history
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-history
func (s *Service) ListHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectHistory, *interfaces.Response, error) {
	endpoint := EndpointJamfProtectHistoryV1

	var result ListResponseJamfProtectHistory

	mergePage := func(pageData []byte) error {
		var pageResponse ListResponseJamfProtectHistory
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect history: %w", err)
	}
	return &result, resp, nil
}

// CreateHistoryNoteV1 creates a new history note for Jamf Protect.
// URL: POST /api/v1/jamf-protect/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-protect-history
func (s *Service) CreateHistoryNoteV1(ctx context.Context, request *RequestJamfProtectHistoryNote) (*ResourceJamfProtectHistory, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("Jamf Protect history note request cannot be nil")
	}

	endpoint := EndpointJamfProtectHistoryV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var result ResourceJamfProtectHistory
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create Jamf Protect history note: %w", err)
	}

	return &result, resp, nil
}

// ListPlansV1 retrieves paginated list of Jamf Protect plans.
// URL: GET /api/v1/jamf-protect/plans
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-protect-plans
func (s *Service) ListPlansV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponseJamfProtectPlans, *interfaces.Response, error) {
	endpoint := EndpointJamfProtectPlansV1

	var result ListResponseJamfProtectPlans

	mergePage := func(pageData []byte) error {
		var pageResponse ListResponseJamfProtectPlans
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list Jamf Protect plans: %w", err)
	}
	return &result, resp, nil
}

// DeleteIntegrationV1 removes the Jamf Protect integration.
// URL: DELETE /api/v1/jamf-protect
// https://developer.jamf.com/jamf-pro/reference/delete_v1-jamf-protect
func (s *Service) DeleteIntegrationV1(ctx context.Context) (*interfaces.Response, error) {
	endpoint := EndpointJamfProtectV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete Jamf Protect integration: %w", err)
	}

	return resp, nil
}
