package computer_prestages

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerPrestagesServiceInterface defines the interface for computer prestage operations.
	// CRUD uses v3 API; device scope uses v2 API. Supports optimistic locking via versionLock.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
	ComputerPrestagesServiceInterface interface {
		// ListV3 returns a page of computer prestages (Get Computer Prestages).
		//
		// Query params (optional, pass via query): page, page-size, sort (e.g. id:asc, displayName:desc).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
		ListV3(ctx context.Context, query map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV3 returns the computer prestage by ID (Get Computer Prestage by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages-id
		GetByIDV3(ctx context.Context, id string) (*ResourceComputerPrestage, *interfaces.Response, error)

		// GetByNameV3 returns the computer prestage by display name (searches first page of ListV3).
		GetByNameV3(ctx context.Context, name string) (*ResourceComputerPrestage, *interfaces.Response, error)

		// CreateV3 creates a new computer prestage (Create Computer Prestage).
		// Returns CreateResponse (id, href).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computer-prestages
		CreateV3(ctx context.Context, request *ResourceComputerPrestage) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV3 updates the computer prestage by ID (Update Computer Prestage by ID).
		// Include versionLock from the current resource for optimistic locking.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-computer-prestages-id
		UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *interfaces.Response, error)

		// UpdateByNameV3 updates the computer prestage by display name.
		UpdateByNameV3(ctx context.Context, name string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *interfaces.Response, error)

		// DeleteByIDV3 deletes the computer prestage by ID (Delete Computer Prestage by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computer-prestages-id
		DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV3 deletes the computer prestage by display name.
		DeleteByNameV3(ctx context.Context, name string) (*interfaces.Response, error)

		// GetDeviceScopeByIDV2 returns the device scope for the computer prestage by ID (Get scope; v2 API).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-id-scope
		GetDeviceScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *interfaces.Response, error)

		// ReplaceDeviceScopeByIDV2 replaces the device scope for the computer prestage by ID (Put scope; v2 API).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-prestages-id-scope
		ReplaceDeviceScopeByIDV2(ctx context.Context, id string, request *ReplaceDeviceScopeRequest) (*ResourceDeviceScope, *interfaces.Response, error)
	}

	// Service handles communication with the computer prestages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerPrestagesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV3 returns a page of computer prestages.
// URL: GET /api/v3/computer-prestages
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages
func (s *Service) ListV3(ctx context.Context, query map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointComputerPrestagesV3, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV3 returns the computer prestage by ID.
// URL: GET /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computer-prestages-id
func (s *Service) GetByIDV3(ctx context.Context, id string) (*ResourceComputerPrestage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerPrestagesV3, id)

	var result ResourceComputerPrestage

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

// GetByNameV3 returns the computer prestage by display name (searches first page).
func (s *Service) GetByNameV3(ctx context.Context, name string) (*ResourceComputerPrestage, *interfaces.Response, error) {
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
func (s *Service) CreateV3(ctx context.Context, request *ResourceComputerPrestage) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if err := validateRequest(request); err != nil {
		return nil, nil, fmt.Errorf("request validation failed: %w", err)
	}

	var result CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointComputerPrestagesV3, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the computer prestage by ID.
// URL: PUT /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-computer-prestages-id
func (s *Service) UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if err := validateRequest(request); err != nil {
		return nil, nil, fmt.Errorf("request validation failed: %w", err)
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerPrestagesV3, id)
	var result ResourceComputerPrestage

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

// UpdateByNameV3 updates the computer prestage by display name.
func (s *Service) UpdateByNameV3(ctx context.Context, name string, request *ResourceComputerPrestage) (*ResourceComputerPrestage, *interfaces.Response, error) {
	existing, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return nil, resp, err
	}
	return s.UpdateByIDV3(ctx, existing.ID, request)
}

// DeleteByIDV3 deletes the computer prestage by ID.
// URL: DELETE /api/v3/computer-prestages/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computer-prestages-id
func (s *Service) DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerPrestagesV3, id)

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

// DeleteByNameV3 deletes the computer prestage by display name.
func (s *Service) DeleteByNameV3(ctx context.Context, name string) (*interfaces.Response, error) {
	existing, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return resp, err
	}
	return s.DeleteByIDV3(ctx, existing.ID)
}

// GetDeviceScopeByIDV2 returns the device scope for the computer prestage by ID (v2 API).
// URL: GET /api/v2/computer-prestages/{id}/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-prestages-id-scope
func (s *Service) GetDeviceScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointComputerPrestagesV2, id)

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

// ReplaceDeviceScopeByIDV2 replaces the device scope for the computer prestage by ID (v2 API).
// URL: PUT /api/v2/computer-prestages/{id}/scope
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-prestages-id-scope
func (s *Service) ReplaceDeviceScopeByIDV2(ctx context.Context, id string, request *ReplaceDeviceScopeRequest) (*ResourceDeviceScope, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointComputerPrestagesV2, id)

	var result ResourceDeviceScope

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
