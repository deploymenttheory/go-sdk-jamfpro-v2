package device_enrollments

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// DeviceEnrollmentsServiceInterface defines the interface for device enrollment operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
	DeviceEnrollmentsServiceInterface interface {
		// ListV1 returns a paginated list of device enrollment objects.
		//
		// Supports optional pagination and sorting via rsqlQuery
		// (keys: sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified device enrollment by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id
		GetByIDV1(ctx context.Context, id string) (*ResourceDeviceEnrollment, *resty.Response, error)

		// GetByNameV1 returns the specified device enrollment by name.
		//
		// Note: This performs a client-side search through all device enrollments.
		GetByNameV1(ctx context.Context, name string) (*ResourceDeviceEnrollment, *resty.Response, error)

		// GetHistoryV1 returns the history for the specified device enrollment.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-history
		GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// GetSyncStatesV1 retrieves all sync states for the specified device enrollment instance.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-syncs
		GetSyncStatesV1(ctx context.Context, id string) ([]ResourceSyncState, *resty.Response, error)

		// GetLatestSyncStateV1 retrieves the latest sync state for the specified device enrollment instance.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-syncs-latest
		GetLatestSyncStateV1(ctx context.Context, id string) (*ResourceLatestSyncState, *resty.Response, error)

		// GetAllSyncStatesV1 retrieves all sync states for all device enrollment instances.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-syncs
		GetAllSyncStatesV1(ctx context.Context) ([]ResourceSyncState, *resty.Response, error)

		// GetPublicKeyV1 retrieves the public key for device enrollments as a PEM file.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-public-key
		GetPublicKeyV1(ctx context.Context) ([]byte, *resty.Response, error)

		// CreateWithTokenV1 creates a new device enrollment instance using an MDM server token.
		//
		// Returns the created device enrollment's ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-upload-token
		CreateWithTokenV1(ctx context.Context, request *RequestTokenUpload) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the metadata for the specified device enrollment by ID.
		//
		// Returns the full updated device enrollment resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-enrollments-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestUpdate) (*ResourceDeviceEnrollment, *resty.Response, error)

		// UpdateTokenByIDV1 updates the token for the specified device enrollment by ID.
		//
		// Returns the full updated device enrollment resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-enrollments-id-upload-token
		UpdateTokenByIDV1(ctx context.Context, id string, request *RequestTokenUpload) (*ResourceDeviceEnrollment, *resty.Response, error)

		// DeleteByIDV1 removes the specified device enrollment by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-device-enrollments-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// DisownDevicesByIDV1 disowns devices from the specified device enrollment instance.
		//
		// Returns a map of device serial numbers to operation status (SUCCESS/FAILED).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-id-disown
		DisownDevicesByIDV1(ctx context.Context, id string, request *RequestDisown) (*ResponseDisown, *resty.Response, error)

		// AddHistoryNotesV1 adds notes to the specified device enrollment's history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-id-history
		AddHistoryNotesV1(ctx context.Context, id string, request *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error)

		// GetDevicesByIDV1 retrieves a list of devices assigned to the specified device enrollment instance.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-devices
		GetDevicesByIDV1(ctx context.Context, id string) (*DevicesResponse, *resty.Response, error)
	}

	// Service handles communication with the device enrollments-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DeviceEnrollmentsServiceInterface = (*Service)(nil)

// NewService returns a new device enrollments Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Device Enrollments Operations
// -----------------------------------------------------------------------------

// ListV1 returns a paginated list of device enrollment objects.
// URL: GET /api/v1/device-enrollments
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointDeviceEnrollmentsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceDeviceEnrollment
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list device enrollments: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified device enrollment by ID.
// URL: GET /api/v1/device-enrollments/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceDeviceEnrollment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceDeviceEnrollment

	endpoint := fmt.Sprintf("%s/%s", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device enrollment by ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// GetByNameV1 returns the specified device enrollment by name.
// Note: This performs a client-side search through all device enrollments.
func (s *Service) GetByNameV1(ctx context.Context, name string) (*ResourceDeviceEnrollment, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	list, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list device enrollments: %w", err)
	}

	for _, enrollment := range list.Results {
		if enrollment.Name == name {
			return &enrollment, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("device enrollment with name %q not found", name)
}

// GetHistoryV1 returns the history for the specified device enrollment.
// URL: GET /api/v1/device-enrollments/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-history
func (s *Service) GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result HistoryResponse

	endpoint := fmt.Sprintf("%s/%s/history", EndpointDeviceEnrollmentsV1, id)

	mergePage := func(pageData []byte) error {
		var items []ResourceHistoryEntry
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device enrollment history for ID %s: %w", id, err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetSyncStatesV1 retrieves all sync states for the specified device enrollment instance.
// URL: GET /api/v1/device-enrollments/{id}/syncs
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-syncs
func (s *Service) GetSyncStatesV1(ctx context.Context, id string) ([]ResourceSyncState, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result []ResourceSyncState

	endpoint := fmt.Sprintf("%s/%s/syncs", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device enrollment sync states for ID %s: %w", id, err)
	}

	return result, resp, nil
}

// GetLatestSyncStateV1 retrieves the latest sync state for the specified device enrollment instance.
// URL: GET /api/v1/device-enrollments/{id}/syncs/latest
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-syncs-latest
func (s *Service) GetLatestSyncStateV1(ctx context.Context, id string) (*ResourceLatestSyncState, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceLatestSyncState

	endpoint := fmt.Sprintf("%s/%s/syncs/latest", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get latest sync state for device enrollment ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// GetAllSyncStatesV1 retrieves all sync states for all device enrollment instances.
// URL: GET /api/v1/device-enrollments/syncs
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-syncs
func (s *Service) GetAllSyncStatesV1(ctx context.Context) ([]ResourceSyncState, *resty.Response, error) {
	var result []ResourceSyncState

	endpoint := fmt.Sprintf("%s/syncs", EndpointDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get all device enrollment sync states: %w", err)
	}

	return result, resp, nil
}

// GetPublicKeyV1 retrieves the public key for device enrollments as a PEM file.
// URL: GET /api/v1/device-enrollments/public-key
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-public-key
func (s *Service) GetPublicKeyV1(ctx context.Context) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/public-key", EndpointDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept": mime.ApplicationXPEMFile,
	}

	resp, result, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device enrollments public key: %w", err)
	}

	return result, resp, nil
}

// CreateWithTokenV1 creates a new device enrollment instance using an MDM server token.
// URL: POST /api/v1/device-enrollments/upload-token
// https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-upload-token
func (s *Service) CreateWithTokenV1(ctx context.Context, request *RequestTokenUpload) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.EncodedToken == "" {
		return nil, nil, fmt.Errorf("encodedToken is required")
	}

	var result CreateResponse

	endpoint := fmt.Sprintf("%s/upload-token", EndpointDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create device enrollment with token: %w", err)
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the metadata for the specified device enrollment by ID.
// URL: PUT /api/v1/device-enrollments/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-device-enrollments-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestUpdate) (*ResourceDeviceEnrollment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	var result ResourceDeviceEnrollment

	endpoint := fmt.Sprintf("%s/%s", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update device enrollment by ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// UpdateTokenByIDV1 updates the token for the specified device enrollment by ID.
// URL: PUT /api/v1/device-enrollments/{id}/upload-token
// https://developer.jamf.com/jamf-pro/reference/put_v1-device-enrollments-id-upload-token
func (s *Service) UpdateTokenByIDV1(ctx context.Context, id string, request *RequestTokenUpload) (*ResourceDeviceEnrollment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.EncodedToken == "" {
		return nil, nil, fmt.Errorf("encodedToken is required")
	}

	var result ResourceDeviceEnrollment

	endpoint := fmt.Sprintf("%s/%s/upload-token", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update device enrollment token by ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified device enrollment by ID.
// URL: DELETE /api/v1/device-enrollments/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-device-enrollments-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete device enrollment by ID %s: %w", id, err)
	}

	return resp, nil
}

// DisownDevicesByIDV1 disowns devices from the specified device enrollment instance.
// URL: POST /api/v1/device-enrollments/{id}/disown
// https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-id-disown
func (s *Service) DisownDevicesByIDV1(ctx context.Context, id string, request *RequestDisown) (*ResponseDisown, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if len(request.Devices) == 0 {
		return nil, nil, fmt.Errorf("devices list is required")
	}

	var result ResponseDisown

	endpoint := fmt.Sprintf("%s/%s/disown", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to disown devices for device enrollment ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// AddHistoryNotesV1 adds notes to the specified device enrollment's history.
// URL: POST /api/v1/device-enrollments/{id}/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-device-enrollments-id-history
func (s *Service) AddHistoryNotesV1(ctx context.Context, id string, request *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	var result ResponseAddHistoryNotes

	endpoint := fmt.Sprintf("%s/%s/history", EndpointDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to add history notes for device enrollment ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// GetDevicesByIDV1 retrieves a list of devices assigned to the specified device enrollment instance.
// URL: GET /api/v1/device-enrollments/{id}/devices
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments-id-devices
func (s *Service) GetDevicesByIDV1(ctx context.Context, id string) (*DevicesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("device enrollment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/devices", EndpointDeviceEnrollmentsV1, id)

	var result DevicesResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get devices for device enrollment ID %s: %w", id, err)
	}

	return &result, resp, nil
}
