package managed_software_updates

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ManagedSoftwareUpdatesServiceInterface defines the interface for managed software updates operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
	ManagedSoftwareUpdatesServiceInterface interface {
		// GetAvailableUpdates retrieves a list of all available managed software updates.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
		GetAvailableUpdates(ctx context.Context) (*ResponseAvailableUpdates, *resty.Response, error)

		// GetPlans retrieves a list of all managed software update plans.
		//
		// Query parameters can be used for filtering and pagination.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans
		GetPlans(ctx context.Context, params url.Values) (*ResponsePlanList, *resty.Response, error)

		// GetPlanByUUID retrieves a managed software update plan by its UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id
		GetPlanByUUID(ctx context.Context, uuid string) (*ResourcePlan, *resty.Response, error)

		// GetDeclarationsByPlanUUID retrieves all declarations associated with a managed software update plan by its UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id-declarations
		GetDeclarationsByPlanUUID(ctx context.Context, uuid string) (*ResponseDeclarationsList, *resty.Response, error)

		// CreatePlanByDeviceID creates a managed software update plan by device ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans
		CreatePlanByDeviceID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *resty.Response, error)

		// CreatePlanByGroupID creates a managed software update plan by group ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-group
		CreatePlanByGroupID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *resty.Response, error)

		// GetPlansByGroupID retrieves managed software update plans for a specific group ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-group-id
		GetPlansByGroupID(ctx context.Context, groupID string, groupType string) (*ResponsePlanList, *resty.Response, error)

		// GetFeatureToggle retrieves the current managed software update feature toggle settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle
		GetFeatureToggle(ctx context.Context) (*RequestFeatureToggle, *resty.Response, error)

		// UpdateFeatureToggle updates the feature toggle for managed software updates.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-managed-software-updates-plans-feature-toggle
		UpdateFeatureToggle(ctx context.Context, toggle *RequestFeatureToggle) (*ResponseFeatureToggle, *resty.Response, error)

		// GetFeatureToggleStatus retrieves the background status of the feature toggle.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-feature-toggle-status
		GetFeatureToggleStatus(ctx context.Context) (*ResponseFeatureToggleStatus, *resty.Response, error)

		// ForceStopFeatureToggleProcess forcefully stops any ongoing or stalled feature-toggle processes.
		// This "Break Glass" endpoint should not be used under nominal conditions.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-managed-software-updates-plans-feature-toggle-abandon
		ForceStopFeatureToggleProcess(ctx context.Context) (*ResponseError, *resty.Response, error)

		// GetPlanEventsByUUID retrieves the event store for a managed software update plan by its UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id-events
		GetPlanEventsByUUID(ctx context.Context, uuid string) (*ResponsePlanEvents, *resty.Response, error)

		// GetUpdateStatuses retrieves update statuses with RSQL filter and pagination support.
		//
		// Query parameters: filter (RSQL), page, page-size, sort.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses
		GetUpdateStatuses(ctx context.Context, params url.Values) (*ResponseUpdateStatusList, *resty.Response, error)

		// GetUpdateStatusesByComputerGroup retrieves update statuses for a computer group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-computer-groups-id
		GetUpdateStatusesByComputerGroup(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error)

		// GetUpdateStatusesByComputer retrieves update statuses for a computer.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-computers-id
		GetUpdateStatusesByComputer(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error)

		// GetUpdateStatusesByMobileDeviceGroup retrieves update statuses for a mobile device group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-mobile-device-groups-id
		GetUpdateStatusesByMobileDeviceGroup(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error)

		// GetUpdateStatusesByMobileDevice retrieves update statuses for a mobile device.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-mobile-devices-id
		GetUpdateStatusesByMobileDevice(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error)
	}

	// Service handles communication with the managed software updates-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
	ManagedSoftwareUpdates struct {
		client transport.HTTPClient
	}
)

var _ ManagedSoftwareUpdatesServiceInterface = (*ManagedSoftwareUpdates)(nil)

// NewService returns a new managed software updates Service backed by the provided HTTP client.
func NewManagedSoftwareUpdates(client transport.HTTPClient) *ManagedSoftwareUpdates {
	return &ManagedSoftwareUpdates{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Managed Software Updates CRUD Operations (v1)
// -----------------------------------------------------------------------------

// GetAvailableUpdates retrieves a list of all available managed software updates.
// URL: GET /api/v1/managed-software-updates/available-updates
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-available-updates
func (s *ManagedSoftwareUpdates) GetAvailableUpdates(ctx context.Context) (*ResponseAvailableUpdates, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/available-updates"

	var result ResponseAvailableUpdates

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) GetPlans(ctx context.Context, params url.Values) (*ResponsePlanList, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	queryParams := make(map[string]string)
	for key, values := range params {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	var result ResponsePlanList

	mergePage := func(pageData []byte) error {
		var pageItems []ResourcePlan
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
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
func (s *ManagedSoftwareUpdates) GetPlanByUUID(ctx context.Context, uuid string) (*ResourcePlan, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	endpoint := fmt.Sprintf("%s/plans/%s", constants.EndpointJamfProManagedSoftwareUpdates, uuid)

	var result ResourcePlan

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) GetDeclarationsByPlanUUID(ctx context.Context, uuid string) (*ResponseDeclarationsList, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	endpoint := fmt.Sprintf("%s/plans/%s/declarations", constants.EndpointJamfProManagedSoftwareUpdates, uuid)

	var result ResponseDeclarationsList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) CreatePlanByDeviceID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *resty.Response, error) {
	if plan == nil {
		return nil, nil, fmt.Errorf("plan is required")
	}

	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans"

	var result ResponsePlanCreate

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) CreatePlanByGroupID(ctx context.Context, plan *RequestPlanCreate) (*ResponsePlanCreate, *resty.Response, error) {
	if plan == nil {
		return nil, nil, fmt.Errorf("plan is required")
	}

	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans/group"

	var result ResponsePlanCreate

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, plan, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPlansByGroupID retrieves managed software update plans for a specific group ID.
// URL: GET /api/v1/managed-software-updates/plans/group/{groupId}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-group-id
func (s *ManagedSoftwareUpdates) GetPlansByGroupID(ctx context.Context, groupID string, groupType string) (*ResponsePlanList, *resty.Response, error) {
	if groupID == "" {
		return nil, nil, fmt.Errorf("groupID is required")
	}
	if groupType == "" {
		return nil, nil, fmt.Errorf("groupType is required")
	}

	endpoint := fmt.Sprintf("%s/plans/group/%s?group-type=%s", constants.EndpointJamfProManagedSoftwareUpdates, groupID, groupType)

	var result ResponsePlanList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) GetFeatureToggle(ctx context.Context) (*RequestFeatureToggle, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans/feature-toggle"

	var result RequestFeatureToggle

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) UpdateFeatureToggle(ctx context.Context, toggle *RequestFeatureToggle) (*ResponseFeatureToggle, *resty.Response, error) {
	if toggle == nil {
		return nil, nil, fmt.Errorf("toggle is required")
	}

	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans/feature-toggle"

	var result ResponseFeatureToggle

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) GetFeatureToggleStatus(ctx context.Context) (*ResponseFeatureToggleStatus, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans/feature-toggle/status"

	var result ResponseFeatureToggleStatus

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ManagedSoftwareUpdates) ForceStopFeatureToggleProcess(ctx context.Context) (*ResponseError, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/plans/feature-toggle/abandon"

	var result ResponseError

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPlanEventsByUUID retrieves the event store for a managed software update plan by its UUID.
// URL: GET /api/v1/managed-software-updates/plans/{id}/events
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-plans-id-events
func (s *ManagedSoftwareUpdates) GetPlanEventsByUUID(ctx context.Context, uuid string) (*ResponsePlanEvents, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	endpoint := fmt.Sprintf("%s/plans/%s/events", constants.EndpointJamfProManagedSoftwareUpdates, uuid)

	var result ResponsePlanEvents

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetUpdateStatuses retrieves update statuses with RSQL filter and pagination support.
// URL: GET /api/v1/managed-software-updates/update-statuses
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses
func (s *ManagedSoftwareUpdates) GetUpdateStatuses(ctx context.Context, params url.Values) (*ResponseUpdateStatusList, *resty.Response, error) {
	endpoint := constants.EndpointJamfProManagedSoftwareUpdates + "/update-statuses"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	queryParams := make(map[string]string)
	for key, values := range params {
		if len(values) > 0 {
			queryParams[key] = values[0]
		}
	}

	var result ResponseUpdateStatusList

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceUpdateStatus
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, queryParams, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetUpdateStatusesByComputerGroup retrieves update statuses for a computer group.
// URL: GET /api/v1/managed-software-updates/update-statuses/computer-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-computer-groups-id
func (s *ManagedSoftwareUpdates) GetUpdateStatusesByComputerGroup(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/update-statuses/computer-groups/%s", constants.EndpointJamfProManagedSoftwareUpdates, id)

	var result ResponseUpdateStatusList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetUpdateStatusesByComputer retrieves update statuses for a computer.
// URL: GET /api/v1/managed-software-updates/update-statuses/computers/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-computers-id
func (s *ManagedSoftwareUpdates) GetUpdateStatusesByComputer(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/update-statuses/computers/%s", constants.EndpointJamfProManagedSoftwareUpdates, id)

	var result ResponseUpdateStatusList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetUpdateStatusesByMobileDeviceGroup retrieves update statuses for a mobile device group.
// URL: GET /api/v1/managed-software-updates/update-statuses/mobile-device-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-mobile-device-groups-id
func (s *ManagedSoftwareUpdates) GetUpdateStatusesByMobileDeviceGroup(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/update-statuses/mobile-device-groups/%s", constants.EndpointJamfProManagedSoftwareUpdates, id)

	var result ResponseUpdateStatusList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetUpdateStatusesByMobileDevice retrieves update statuses for a mobile device.
// URL: GET /api/v1/managed-software-updates/update-statuses/mobile-devices/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-managed-software-updates-update-statuses-mobile-devices-id
func (s *ManagedSoftwareUpdates) GetUpdateStatusesByMobileDevice(ctx context.Context, id string) (*ResponseUpdateStatusList, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/update-statuses/mobile-devices/%s", constants.EndpointJamfProManagedSoftwareUpdates, id)

	var result ResponseUpdateStatusList

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
