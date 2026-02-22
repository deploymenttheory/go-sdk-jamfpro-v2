package managed_software_updates

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// ManagedSoftwareUpdatesServiceInterface defines the interface for managed software updates operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
	ManagedSoftwareUpdatesServiceInterface interface {
		// GetAvailableUpdates retrieves a list of all available managed software updates.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
		GetAvailableUpdates(ctx context.Context) (*ResponseAvailableUpdates, *interfaces.Response, error)

		// GetPlans retrieves a list of all managed software update plans.
		//
		// Query parameters can be used for filtering and pagination.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans
		GetPlans(ctx context.Context, params url.Values) (*ResponsePlanList, *interfaces.Response, error)

		// GetPlanByUUID retrieves a managed software update plan by its UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id
		GetPlanByUUID(ctx context.Context, uuid string) (*ResourcePlan, *interfaces.Response, error)

		// GetDeclarationsByPlanUUID retrieves all declarations associated with a managed software update plan by its UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id-declarations
		GetDeclarationsByPlanUUID(ctx context.Context, uuid string) (*ResponseDeclarationsList, *interfaces.Response, error)

		// CreatePlanByDeviceID creates a managed software update plan by device ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans
		CreatePlanByDeviceID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *interfaces.Response, error)

		// CreatePlanByGroupID creates a managed software update plan by group ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-group
		CreatePlanByGroupID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *interfaces.Response, error)

		// GetPlansByGroupID retrieves managed software update plans for a specific group ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-group-groupid
		GetPlansByGroupID(ctx context.Context, groupID string, groupType string) (*ResponsePlanList, *interfaces.Response, error)

		// GetFeatureToggle retrieves the current managed software update feature toggle settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle
		GetFeatureToggle(ctx context.Context) (*RequestFeatureToggle, *interfaces.Response, error)

		// UpdateFeatureToggle updates the feature toggle for managed software updates.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-managed-software-updates-plans-feature-toggle
		UpdateFeatureToggle(ctx context.Context, toggle *RequestFeatureToggle) (*ResponseFeatureToggle, *interfaces.Response, error)

		// GetFeatureToggleStatus retrieves the background status of the feature toggle.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle-status
		GetFeatureToggleStatus(ctx context.Context) (*ResponseFeatureToggleStatus, *interfaces.Response, error)

		// ForceStopFeatureToggleProcess forcefully stops any ongoing or stalled feature-toggle processes.
		// This "Break Glass" endpoint should not be used under nominal conditions.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-feature-toggle-abandon
		ForceStopFeatureToggleProcess(ctx context.Context) (*ResponseError, *interfaces.Response, error)
	}

	// Service handles communication with the managed software updates-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ManagedSoftwareUpdatesServiceInterface = (*Service)(nil)

// NewService returns a new managed software updates Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Managed Software Updates CRUD Operations (v1)
// -----------------------------------------------------------------------------

// GetAvailableUpdates retrieves a list of all available managed software updates.
// URL: GET /api/v1/managed-software-updates/available-updates
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
func (s *Service) GetAvailableUpdates(ctx context.Context) (*ResponseAvailableUpdates, *interfaces.Response, error) {
	endpoint := EndpointManagedSoftwareUpdates + "/available-updates"

	var result ResponseAvailableUpdates

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPlans retrieves a list of all managed software update plans.
// URL: GET /api/v1/managed-software-updates/plans
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans
func (s *Service) GetPlans(ctx context.Context, params url.Values) (*ResponsePlanList, *interfaces.Response, error) {
	endpoint := EndpointManagedSoftwareUpdates + "/plans"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	// Convert URL params to map for the client
	queryParams := make(map[string]string)
	for key, values := range params {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	// Use GetPaginated for the paginated endpoint
	var result ResponsePlanList

	mergePage := func(pageData []byte) error {
		var rawData map[string]interface{}
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		// Extract results array from the response
		if results, ok := rawData["results"].([]interface{}); ok {
			for _, item := range results {
				var plan ResourcePlan
				if err := mapstructure.Decode(item, &plan); err != nil {
					return fmt.Errorf("failed to decode plan: %w", err)
				}
				result.Results = append(result.Results, plan)
			}
		}
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, queryParams, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// GetPlanByUUID retrieves a managed software update plan by its UUID.
// URL: GET /api/v1/managed-software-updates/plans/{uuid}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id
func (s *Service) GetPlanByUUID(ctx context.Context, uuid string) (*ResourcePlan, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	endpoint := fmt.Sprintf("%s/plans/%s", EndpointManagedSoftwareUpdates, uuid)

	var result ResourcePlan

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeclarationsByPlanUUID retrieves all declarations associated with a managed software update plan by its UUID.
// URL: GET /api/v1/managed-software-updates/plans/{uuid}/declarations
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id-declarations
func (s *Service) GetDeclarationsByPlanUUID(ctx context.Context, uuid string) (*ResponseDeclarationsList, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	endpoint := fmt.Sprintf("%s/plans/%s/declarations", EndpointManagedSoftwareUpdates, uuid)

	var result ResponseDeclarationsList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreatePlanByDeviceID creates a managed software update plan by device ID.
// URL: POST /api/v1/managed-software-updates/plans
// https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans
func (s *Service) CreatePlanByDeviceID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *interfaces.Response, error) {
	if plan == nil {
		return nil, nil, fmt.Errorf("plan is required")
	}

	endpoint := EndpointManagedSoftwareUpdates + "/plans"

	var result ResponsePlanCreate

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, plan, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreatePlanByGroupID creates a managed software update plan by group ID.
// URL: POST /api/v1/managed-software-updates/plans/group
// https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-group
func (s *Service) CreatePlanByGroupID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *interfaces.Response, error) {
	if plan == nil {
		return nil, nil, fmt.Errorf("plan is required")
	}

	endpoint := EndpointManagedSoftwareUpdates + "/plans/group"

	var result ResponsePlanCreate

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, plan, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPlansByGroupID retrieves managed software update plans for a specific group ID.
// URL: GET /api/v1/managed-software-updates/plans/group/{groupId}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-group-groupid
func (s *Service) GetPlansByGroupID(ctx context.Context, groupID string, groupType string) (*ResponsePlanList, *interfaces.Response, error) {
	if groupID == "" {
		return nil, nil, fmt.Errorf("groupID is required")
	}
	if groupType == "" {
		return nil, nil, fmt.Errorf("groupType is required")
	}

	endpoint := fmt.Sprintf("%s/plans/group/%s?group-type=%s", EndpointManagedSoftwareUpdates, groupID, groupType)

	var result ResponsePlanList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetFeatureToggle retrieves the current managed software update feature toggle settings.
// URL: GET /api/v1/managed-software-updates/plans/feature-toggle
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle
func (s *Service) GetFeatureToggle(ctx context.Context) (*RequestFeatureToggle, *interfaces.Response, error) {
	endpoint := EndpointManagedSoftwareUpdates + "/plans/feature-toggle"

	var result RequestFeatureToggle

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateFeatureToggle updates the feature toggle for managed software updates.
// URL: PUT /api/v1/managed-software-updates/plans/feature-toggle
// https://developer.jamf.com/jamf-pro/reference/put_v1-managed-software-updates-plans-feature-toggle
func (s *Service) UpdateFeatureToggle(ctx context.Context, toggle *RequestFeatureToggle) (*ResponseFeatureToggle, *interfaces.Response, error) {
	if toggle == nil {
		return nil, nil, fmt.Errorf("toggle is required")
	}

	endpoint := EndpointManagedSoftwareUpdates + "/plans/feature-toggle"

	var result ResponseFeatureToggle

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, toggle, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetFeatureToggleStatus retrieves the background status of the feature toggle.
// URL: GET /api/v1/managed-software-updates/plans/feature-toggle/status
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle-status
func (s *Service) GetFeatureToggleStatus(ctx context.Context) (*ResponseFeatureToggleStatus, *interfaces.Response, error) {
	endpoint := EndpointManagedSoftwareUpdates + "/plans/feature-toggle/status"

	var result ResponseFeatureToggleStatus

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ForceStopFeatureToggleProcess forcefully stops any ongoing or stalled feature-toggle processes.
// This "Break Glass" endpoint should not be used under nominal conditions.
// URL: POST /api/v1/managed-software-updates/plans/feature-toggle/abandon
// https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-feature-toggle-abandon
func (s *Service) ForceStopFeatureToggleProcess(ctx context.Context) (*ResponseError, *interfaces.Response, error) {
	endpoint := EndpointManagedSoftwareUpdates + "/plans/feature-toggle/abandon"

	var result ResponseError

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
