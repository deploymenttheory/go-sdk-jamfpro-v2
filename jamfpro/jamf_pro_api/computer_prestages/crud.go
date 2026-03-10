package computer_prestages

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/version_locking"
	"resty.dev/v3"
)

type (
	// Service handles communication with the computer prestages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
	ComputerPrestages struct {
		client client.Client
	}
)

func NewComputerPrestages(client client.Client) *ComputerPrestages {
	return &ComputerPrestages{client: client}
}

// ListV3 returns a page of computer prestages.
// URL: GET /api/v3/computer-prestages
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
func (s *ComputerPrestages) ListV3(ctx context.Context, query map[string]string) (*ListResponse, *resty.Response, error) {
	result := &ListResponse{
		Results: []ResourceComputerPrestage{},
	}

	endpoint := constants.EndpointJamfProComputerPrestagesV3

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceComputerPrestage
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		result.TotalCount = len(result.Results)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// GetByIDV3 returns the computer prestage by ID.
// URL: GET /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages-id
func (s *ComputerPrestages) GetByIDV3(ctx context.Context, id string) (*ResourceComputerPrestage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerPrestagesV3, id)

	var result ResourceComputerPrestage

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByNameV3 returns the computer prestage by display name (searches first page).
func (s *ComputerPrestages) GetByNameV3(ctx context.Context, name string) (*ResourceComputerPrestage, *resty.Response, error) {
	list, resp, err := s.ListV3(ctx, nil)
	if err != nil {
		return nil, resp, err
	}
	for i := range list.Results {
		if list.Results[i].DisplayName == name {
			return &list.Results[i], resp, nil
		}
	}
	return nil, resp, fmt.Errorf("computer prestage with name %q not found", name)
}

// CreateV3 creates a new computer prestage.
// URL: POST /api/v3/computer-prestages
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computer-prestages
func (s *ComputerPrestages) CreateV3(ctx context.Context, request *ResourceComputerPrestage) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if err := validateRequest(request); err != nil {
		return nil, nil, fmt.Errorf("request validation failed: %w", err)
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProComputerPrestagesV3

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the computer prestage by ID.
// The current resource is fetched first so that all versionLock values
// (top-level, locationInformation, purchasingInformation, accountSettings)
// are injected transparently. Callers do not need to supply versionLock.
// URL: PUT /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-computer-prestages-id
func (s *ComputerPrestages) UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if err := validateRequest(request); err != nil {
		return nil, nil, fmt.Errorf("request validation failed: %w", err)
	}

	current, _, err := s.GetByIDV3(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current prestage for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(current, request)
	version_locking.EnsureVersionLock(&current.LocationInformation, &request.LocationInformation)
	version_locking.EnsureVersionLock(&current.PurchasingInformation, &request.PurchasingInformation)
	if current.AccountSettings != nil && request.AccountSettings != nil {
		version_locking.EnsureVersionLock(current.AccountSettings, request.AccountSettings)
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerPrestagesV3, id)
	var result ResourceComputerPrestage

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

// UpdateByNameV3 updates the computer prestage by display name.
// The resource fetched during the name lookup is reused directly for version
// lock injection, avoiding a second round-trip to the API.
func (s *ComputerPrestages) UpdateByNameV3(ctx context.Context, name string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	existing, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return nil, resp, err
	}

	if err := validateRequest(request); err != nil {
		return nil, nil, fmt.Errorf("request validation failed: %w", err)
	}

	version_locking.EnsureVersionLock(existing, request)
	version_locking.EnsureVersionLock(&existing.LocationInformation, &request.LocationInformation)
	version_locking.EnsureVersionLock(&existing.PurchasingInformation, &request.PurchasingInformation)
	if existing.AccountSettings != nil && request.AccountSettings != nil {
		version_locking.EnsureVersionLock(existing.AccountSettings, request.AccountSettings)
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerPrestagesV3, existing.ID)
	var result ResourceComputerPrestage

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err = s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteByIDV3 deletes the computer prestage by ID.
// URL: DELETE /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computer-prestages-id
func (s *ComputerPrestages) DeleteByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerPrestagesV3, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV3 deletes the computer prestage by display name.
func (s *ComputerPrestages) DeleteByNameV3(ctx context.Context, name string) (*resty.Response, error) {
	existing, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return resp, err
	}
	return s.DeleteByIDV3(ctx, existing.ID)
}

// GetDeviceScopeByIDV2 returns the device scope for the computer prestage by ID (v2 API).
// URL: GET /api/v2/computer-prestages/{id}/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-id-scope
func (s *ComputerPrestages) GetDeviceScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", constants.EndpointJamfProComputerPrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ReplaceDeviceScopeByIDV2 replaces the device scope for the computer prestage by ID (v2 API).
// It fetches the current scope first to obtain its versionLock and injects it
// transparently – callers only need to supply the desired serial numbers.
// URL: PUT /api/v2/computer-prestages/{id}/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-prestages-id-scope
func (s *ComputerPrestages) ReplaceDeviceScopeByIDV2(ctx context.Context, id string, request *ReplaceDeviceScopeRequest) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetDeviceScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope", constants.EndpointJamfProComputerPrestagesV2, id)

	var result ResourceDeviceScope

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

// GetAllDeviceScopeV2 returns device scope for all computer prestages.
// URL: GET /api/v2/computer-prestages/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-scope
func (s *ComputerPrestages) GetAllDeviceScopeV2(ctx context.Context) (*AllDeviceScopeResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProComputerPrestagesV2 + "/scope"

	var result AllDeviceScopeResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddDeviceScopeByIDV2 adds device scope (serial numbers) to the computer prestage by ID.
// URL: POST /api/v2/computer-prestages/{id}/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-prestages-id-scope
func (s *ComputerPrestages) AddDeviceScopeByIDV2(ctx context.Context, id string, request *AddDeviceScopeRequest) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetDeviceScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope", constants.EndpointJamfProComputerPrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RemoveDeviceScopeByIDV2 removes device scope (serial numbers) from the computer prestage by ID.
// URL: POST /api/v2/computer-prestages/{id}/scope/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-prestages-id-scope-delete-multiple
func (s *ComputerPrestages) RemoveDeviceScopeByIDV2(ctx context.Context, id string, request *RemoveDeviceScopeRequest) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetDeviceScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope/delete-multiple", constants.EndpointJamfProComputerPrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
