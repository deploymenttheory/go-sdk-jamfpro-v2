package mobile_device_prestages

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// MobileDevicePrestagesServiceInterface defines the interface for mobile device prestage operations.
	// CRUD uses v3 API; device scope uses v2 API. Supports optimistic locking via versionLock.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
	MobileDevicePrestagesServiceInterface interface {
		// ListV3 returns all mobile device prestages using pagination.
		//
		// This method automatically fetches all pages and returns the complete list.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
		ListV3(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByIDV3 returns the mobile device prestage by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages-id
		GetByIDV3(ctx context.Context, id string) (*ResourceMobileDevicePrestage, *interfaces.Response, error)

		// GetByNameV3 returns the mobile device prestage by display name.
		//
		// This is a convenience method that calls ListV3 and filters by DisplayName.
		GetByNameV3(ctx context.Context, name string) (*ResourceMobileDevicePrestage, *interfaces.Response, error)

		// CreateV3 creates a new mobile device prestage.
		// Returns CreateResponse (id, href).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-mobile-device-prestages
		CreateV3(ctx context.Context, prestage *ResourceMobileDevicePrestage) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV3 updates the mobile device prestage by ID.
		// Include versionLock from the current resource for optimistic locking.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-mobile-device-prestages-id
		UpdateByIDV3(ctx context.Context, id string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *interfaces.Response, error)

		// UpdateByNameV3 updates the mobile device prestage by display name.
		UpdateByNameV3(ctx context.Context, name string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *interfaces.Response, error)

		// DeleteByIDV3 deletes the mobile device prestage by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-mobile-device-prestages-id
		DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV3 deletes the mobile device prestage by display name.
		DeleteByNameV3(ctx context.Context, name string) (*interfaces.Response, error)

		// GetScopeByIDV2 returns the device scope for the mobile device prestage by ID (v2 API).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-prestages-id-scope
		GetScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *interfaces.Response, error)
	}

	// Service handles communication with the mobile device prestages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDevicePrestagesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV3 returns all mobile device prestages using pagination.
// URL: GET /api/v3/mobile-device-prestages
// This method automatically fetches all pages and returns the complete list.
// https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
func (s *Service) ListV3(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourceMobileDevicePrestage
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, EndpointMobileDevicePrestagesV3, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// GetByIDV3 returns the mobile device prestage by ID.
// URL: GET /api/v3/mobile-device-prestages/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages-id
func (s *Service) GetByIDV3(ctx context.Context, id string) (*ResourceMobileDevicePrestage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, id)

	var result ResourceMobileDevicePrestage

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

// GetByNameV3 returns the mobile device prestage by display name.
// This is a convenience method that calls ListV3 and filters by DisplayName.
func (s *Service) GetByNameV3(ctx context.Context, name string) (*ResourceMobileDevicePrestage, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	list, resp, err := s.ListV3(ctx)
	if err != nil {
		return nil, resp, err
	}

	for i := range list.Results {
		if list.Results[i].DisplayName == name {
			return &list.Results[i], resp, nil
		}
	}

	return nil, resp, fmt.Errorf("mobile device prestage with name %q not found", name)
}

// CreateV3 creates a new mobile device prestage.
// URL: POST /api/v3/mobile-device-prestages
// https://developer.jamf.com/jamf-pro/reference/post_v3-mobile-device-prestages
func (s *Service) CreateV3(ctx context.Context, prestage *ResourceMobileDevicePrestage) (*CreateResponse, *interfaces.Response, error) {
	if prestage == nil {
		return nil, nil, fmt.Errorf("prestage is required")
	}

	if prestage.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	var result CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointMobileDevicePrestagesV3, prestage, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the mobile device prestage by ID.
// URL: PUT /api/v3/mobile-device-prestages/{id}
// Include versionLock from the current resource for optimistic locking.
// https://developer.jamf.com/jamf-pro/reference/put_v3-mobile-device-prestages-id
func (s *Service) UpdateByIDV3(ctx context.Context, id string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if prestage == nil {
		return nil, nil, fmt.Errorf("prestage is required")
	}

	if prestage.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, id)

	var result ResourceMobileDevicePrestage

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, prestage, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByNameV3 updates the mobile device prestage by display name.
func (s *Service) UpdateByNameV3(ctx context.Context, name string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get mobile device prestage by name: %w", err)
	}

	return s.UpdateByIDV3(ctx, target.ID, prestage)
}

// DeleteByIDV3 deletes the mobile device prestage by ID.
// URL: DELETE /api/v3/mobile-device-prestages/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-mobile-device-prestages-id
func (s *Service) DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV3 deletes the mobile device prestage by display name.
func (s *Service) DeleteByNameV3(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return resp, fmt.Errorf("failed to get mobile device prestage by name: %w", err)
	}

	return s.DeleteByIDV3(ctx, target.ID)
}

// GetScopeByIDV2 returns the device scope for the mobile device prestage by ID.
// URL: GET /api/v2/mobile-device-prestages/{id}/scope
// https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-prestages-id-scope
func (s *Service) GetScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointMobileDevicePrestagesV2, id)

	var result ResourceDeviceScope

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
