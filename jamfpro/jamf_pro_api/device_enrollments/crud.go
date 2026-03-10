package device_enrollments

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the device enrollments-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
	DeviceEnrollments struct {
		client client.Client
	}
)

// NewService returns a new device enrollments Service backed by the provided HTTP client.
func NewDeviceEnrollments(client client.Client) *DeviceEnrollments {
	return &DeviceEnrollments{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Device Enrollments Operations
// -----------------------------------------------------------------------------

// ListV1 returns a paginated list of device enrollment objects.
// URL: GET /api/v1/device-enrollments
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-device-enrollments
func (s *DeviceEnrollments) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProDeviceEnrollmentsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceDeviceEnrollment
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetByIDV1(ctx context.Context, id string) (*ResourceDeviceEnrollment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceDeviceEnrollment

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device enrollment by ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// GetByNameV1 returns the specified device enrollment by name.
// Note: This performs a client-side search through all device enrollments.
func (s *DeviceEnrollments) GetByNameV1(ctx context.Context, name string) (*ResourceDeviceEnrollment, *resty.Response, error) {
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
func (s *DeviceEnrollments) GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result HistoryResponse

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	mergePage := func(pageData []byte) error {
		var items []ResourceHistoryEntry
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetSyncStatesV1(ctx context.Context, id string) ([]ResourceSyncState, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result []ResourceSyncState

	endpoint := fmt.Sprintf("%s/%s/syncs", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetLatestSyncStateV1(ctx context.Context, id string) (*ResourceLatestSyncState, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceLatestSyncState

	endpoint := fmt.Sprintf("%s/%s/syncs/latest", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetAllSyncStatesV1(ctx context.Context) ([]ResourceSyncState, *resty.Response, error) {
	var result []ResourceSyncState

	endpoint := fmt.Sprintf("%s/syncs", constants.EndpointJamfProDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetPublicKeyV1(ctx context.Context) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/public-key", constants.EndpointJamfProDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept": constants.ApplicationXPEMFile,
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
func (s *DeviceEnrollments) CreateWithTokenV1(ctx context.Context, request *RequestTokenUpload) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.EncodedToken == "" {
		return nil, nil, fmt.Errorf("encodedToken is required")
	}

	var result CreateResponse

	endpoint := fmt.Sprintf("%s/upload-token", constants.EndpointJamfProDeviceEnrollmentsV1)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) UpdateByIDV1(ctx context.Context, id string, request *RequestUpdate) (*ResourceDeviceEnrollment, *resty.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) UpdateTokenByIDV1(ctx context.Context, id string, request *RequestTokenUpload) (*ResourceDeviceEnrollment, *resty.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s/upload-token", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) DisownDevicesByIDV1(ctx context.Context, id string, request *RequestDisown) (*ResponseDisown, *resty.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s/disown", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) AddHistoryNotesV1(ctx context.Context, id string, request *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeviceEnrollments) GetDevicesByIDV1(ctx context.Context, id string) (*DevicesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("device enrollment ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/devices", constants.EndpointJamfProDeviceEnrollmentsV1, id)

	var result DevicesResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get devices for device enrollment ID %s: %w", id, err)
	}

	return &result, resp, nil
}
