package mobile_device_prestages

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/version_locking"
	"resty.dev/v3"
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
		ListV3(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByIDV3 returns the mobile device prestage by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages-id
		GetByIDV3(ctx context.Context, id string) (*ResourceMobileDevicePrestage, *resty.Response, error)

		// GetByNameV3 returns the mobile device prestage by display name.
		//
		// This is a convenience method that calls ListV3 and filters by DisplayName.
		GetByNameV3(ctx context.Context, name string) (*ResourceMobileDevicePrestage, *resty.Response, error)

		// CreateV3 creates a new mobile device prestage.
		// Returns CreateResponse (id, href).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-mobile-device-prestages
		CreateV3(ctx context.Context, prestage *ResourceMobileDevicePrestage) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV3 updates the mobile device prestage by ID.
		// The current resource is fetched first so that all versionLock values
		// are injected transparently. Callers do not need to supply versionLock.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-mobile-device-prestages-id
		UpdateByIDV3(ctx context.Context, id string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *resty.Response, error)

		// UpdateByNameV3 updates the mobile device prestage by display name.
		UpdateByNameV3(ctx context.Context, name string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *resty.Response, error)

		// DeleteByIDV3 deletes the mobile device prestage by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-mobile-device-prestages-id
		DeleteByIDV3(ctx context.Context, id string) (*resty.Response, error)

		// DeleteByNameV3 deletes the mobile device prestage by display name.
		DeleteByNameV3(ctx context.Context, name string) (*resty.Response, error)

		// GetScopeByIDV2 returns the device scope for the mobile device prestage by ID (v2 API).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-prestages-id-scope
		GetScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *resty.Response, error)

		// ReplaceScopeByIDV2 replaces the device scope for the mobile device prestage by ID (v2 API).
		ReplaceScopeByIDV2(ctx context.Context, id string, request *RequestReplaceScope) (*ResourceDeviceScope, *resty.Response, error)

		// AddScopeByIDV2 adds device scope (serial numbers) to the mobile device prestage by ID (v2 API).
		AddScopeByIDV2(ctx context.Context, id string, request *RequestAddScope) (*ResourceDeviceScope, *resty.Response, error)

		// RemoveScopeByIDV2 removes device scope (serial numbers) from the mobile device prestage by ID (v2 API).
		RemoveScopeByIDV2(ctx context.Context, id string, request *RequestRemoveScope) (*ResourceDeviceScope, *resty.Response, error)

		// GetAllSyncsV2 returns all prestage sync states for all mobile device prestages (v2 API).
		GetAllSyncsV2(ctx context.Context) ([]ResourcePrestageSync, *resty.Response, error)

		// GetSyncsByIDV2 returns sync states for a specific mobile device prestage by ID (v2 API).
		GetSyncsByIDV2(ctx context.Context, id string) ([]ResourcePrestageSync, *resty.Response, error)

		// GetLatestSyncByIDV2 returns the latest sync state for a mobile device prestage by ID (v2 API).
		GetLatestSyncByIDV2(ctx context.Context, id string) (*ResourcePrestageSync, *resty.Response, error)

		// GetAttachmentsByIDV3 returns attachments for a mobile device prestage by ID (v3 API).
		GetAttachmentsByIDV3(ctx context.Context, id string) ([]ResourceAttachment, *resty.Response, error)

		// UploadAttachmentV3 uploads an attachment to a mobile device prestage by ID (v3 API).
		UploadAttachmentV3(ctx context.Context, id string, fileReader io.Reader, fileSize int64, fileName string) (*ResourceAttachmentUpload, *resty.Response, error)

		// DeleteAttachmentsByIDV3 deletes attachments from a mobile device prestage by ID (v3 API).
		DeleteAttachmentsByIDV3(ctx context.Context, id string, request *RequestDeleteAttachments) (*resty.Response, error)

		// GetHistoryByIDV3 returns the history for a mobile device prestage by ID (v3 API).
		GetHistoryByIDV3(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNoteByIDV3 adds a history note to a mobile device prestage by ID (v3 API).
		AddHistoryNoteByIDV3(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error)
	}

	// Service handles communication with the mobile device prestages-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
	MobileDevicePrestages struct {
		client transport.HTTPClient
	}
)

var _ MobileDevicePrestagesServiceInterface = (*MobileDevicePrestages)(nil)

func NewMobileDevicePrestages(client transport.HTTPClient) *MobileDevicePrestages {
	return &MobileDevicePrestages{client: client}
}

// ListV3 returns all mobile device prestages using pagination.
// URL: GET /api/v3/mobile-device-prestages
// This method automatically fetches all pages and returns the complete list.
// https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages
func (s *MobileDevicePrestages) ListV3(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointMobileDevicePrestagesV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceMobileDevicePrestage
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV3 returns the mobile device prestage by ID.
// URL: GET /api/v3/mobile-device-prestages/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v3-mobile-device-prestages-id
func (s *MobileDevicePrestages) GetByIDV3(ctx context.Context, id string) (*ResourceMobileDevicePrestage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, id)

	var result ResourceMobileDevicePrestage

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV3 returns the mobile device prestage by display name.
// This is a convenience method that calls ListV3 and filters by DisplayName.
func (s *MobileDevicePrestages) GetByNameV3(ctx context.Context, name string) (*ResourceMobileDevicePrestage, *resty.Response, error) {
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
func (s *MobileDevicePrestages) CreateV3(ctx context.Context, prestage *ResourceMobileDevicePrestage) (*CreateResponse, *resty.Response, error) {
	if prestage == nil {
		return nil, nil, fmt.Errorf("prestage is required")
	}

	if prestage.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	var result CreateResponse

	endpoint := EndpointMobileDevicePrestagesV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, prestage, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the mobile device prestage by ID.
// The current resource is fetched first to obtain all versionLock values
// (top-level, locationInformation, purchasingInformation) and inject them
// transparently. Callers do not need to supply versionLock.
// URL: PUT /api/v3/mobile-device-prestages/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v3-mobile-device-prestages-id
func (s *MobileDevicePrestages) UpdateByIDV3(ctx context.Context, id string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if prestage == nil {
		return nil, nil, fmt.Errorf("prestage is required")
	}

	if prestage.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	current, _, err := s.GetByIDV3(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current prestage for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(current, prestage)
	version_locking.EnsureVersionLock(&current.LocationInformation, &prestage.LocationInformation)
	version_locking.EnsureVersionLock(&current.PurchasingInformation, &prestage.PurchasingInformation)

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
// The resource fetched during the name lookup is reused directly for version
// lock injection, avoiding a second round-trip to the API.
func (s *MobileDevicePrestages) UpdateByNameV3(ctx context.Context, name string, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	if prestage == nil {
		return nil, nil, fmt.Errorf("prestage is required")
	}

	if prestage.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	existing, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get mobile device prestage by name: %w", err)
	}

	version_locking.EnsureVersionLock(existing, prestage)
	version_locking.EnsureVersionLock(&existing.LocationInformation, &prestage.LocationInformation)
	version_locking.EnsureVersionLock(&existing.PurchasingInformation, &prestage.PurchasingInformation)

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, existing.ID)

	var result ResourceMobileDevicePrestage

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err = s.client.Put(ctx, endpoint, prestage, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV3 deletes the mobile device prestage by ID.
// URL: DELETE /api/v3/mobile-device-prestages/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-mobile-device-prestages-id
func (s *MobileDevicePrestages) DeleteByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDevicePrestagesV3, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV3 deletes the mobile device prestage by display name.
func (s *MobileDevicePrestages) DeleteByNameV3(ctx context.Context, name string) (*resty.Response, error) {
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
func (s *MobileDevicePrestages) GetScopeByIDV2(ctx context.Context, id string) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointMobileDevicePrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ReplaceScopeByIDV2 replaces the device scope for the mobile device prestage by ID.
// Fetches the current scope first to obtain versionLock and injects it transparently.
// URL: PUT /api/v2/mobile-device-prestages/{id}/scope
func (s *MobileDevicePrestages) ReplaceScopeByIDV2(ctx context.Context, id string, request *RequestReplaceScope) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointMobileDevicePrestagesV2, id)

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

// AddScopeByIDV2 adds device scope (serial numbers) to the mobile device prestage by ID.
// URL: POST /api/v2/mobile-device-prestages/{id}/scope
func (s *MobileDevicePrestages) AddScopeByIDV2(ctx context.Context, id string, request *RequestAddScope) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope", EndpointMobileDevicePrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RemoveScopeByIDV2 removes device scope (serial numbers) from the mobile device prestage by ID.
// URL: POST /api/v2/mobile-device-prestages/{id}/scope/delete-multiple
func (s *MobileDevicePrestages) RemoveScopeByIDV2(ctx context.Context, id string, request *RequestRemoveScope) (*ResourceDeviceScope, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	currentScope, _, err := s.GetScopeByIDV2(ctx, id)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch current device scope for version locking: %w", err)
	}

	version_locking.EnsureVersionLock(currentScope, request)

	endpoint := fmt.Sprintf("%s/%s/scope/delete-multiple", EndpointMobileDevicePrestagesV2, id)

	var result ResourceDeviceScope

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAllSyncsV2 returns all prestage sync states for all mobile device prestages.
// URL: GET /api/v2/mobile-device-prestages/syncs
func (s *MobileDevicePrestages) GetAllSyncsV2(ctx context.Context) ([]ResourcePrestageSync, *resty.Response, error) {
	endpoint := EndpointMobileDevicePrestagesV2 + "/syncs"

	var result []ResourcePrestageSync

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetSyncsByIDV2 returns sync states for a specific mobile device prestage by ID.
// URL: GET /api/v2/mobile-device-prestages/{id}/syncs
func (s *MobileDevicePrestages) GetSyncsByIDV2(ctx context.Context, id string) ([]ResourcePrestageSync, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/syncs", EndpointMobileDevicePrestagesV2, id)

	var result []ResourcePrestageSync

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetLatestSyncByIDV2 returns the latest sync state for a mobile device prestage by ID.
// URL: GET /api/v2/mobile-device-prestages/{id}/syncs/latest
func (s *MobileDevicePrestages) GetLatestSyncByIDV2(ctx context.Context, id string) (*ResourcePrestageSync, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/syncs/latest", EndpointMobileDevicePrestagesV2, id)

	var result ResourcePrestageSync

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetAttachmentsByIDV3 returns attachments for a mobile device prestage by ID.
// URL: GET /api/v3/mobile-device-prestages/{id}/attachments
func (s *MobileDevicePrestages) GetAttachmentsByIDV3(ctx context.Context, id string) ([]ResourceAttachment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments", EndpointMobileDevicePrestagesV3, id)

	var result []ResourceAttachment

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// UploadAttachmentV3 uploads an attachment to a mobile device prestage by ID.
// URL: POST /api/v3/mobile-device-prestages/{id}/attachments
func (s *MobileDevicePrestages) UploadAttachmentV3(ctx context.Context, id string, fileReader io.Reader, fileSize int64, fileName string) (*ResourceAttachmentUpload, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if fileReader == nil {
		return nil, nil, fmt.Errorf("file reader is required")
	}
	if fileName == "" {
		return nil, nil, fmt.Errorf("file name is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments", EndpointMobileDevicePrestagesV3, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceAttachmentUpload

	resp, err := s.client.PostMultipart(ctx, endpoint, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadAttachmentFromFileV3 opens the file at filePath and uploads it via UploadAttachmentV3.
func (s *MobileDevicePrestages) UploadAttachmentFromFileV3(ctx context.Context, id string, filePath string) (*ResourceAttachmentUpload, *resty.Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open attachment file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat attachment file: %w", err)
	}

	fileName := info.Name()
	if fileName == "" {
		fileName = filePath
	}

	return s.UploadAttachmentV3(ctx, id, f, info.Size(), fileName)
}

// DeleteAttachmentsByIDV3 deletes attachments from a mobile device prestage by ID.
// URL: POST /api/v3/mobile-device-prestages/{id}/attachments/delete-multiple
func (s *MobileDevicePrestages) DeleteAttachmentsByIDV3(ctx context.Context, id string, request *RequestDeleteAttachments) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments/delete-multiple", EndpointMobileDevicePrestagesV3, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryByIDV3 returns the history for a mobile device prestage by ID with pagination.
// URL: GET /api/v3/mobile-device-prestages/{id}/history
// Query params: page, page-size, sort, filter
func (s *MobileDevicePrestages) GetHistoryByIDV3(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointMobileDevicePrestagesV3, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get mobile device prestage history: %w", err)
	}
	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// AddHistoryNoteByIDV3 adds a history note to a mobile device prestage by ID.
// URL: POST /api/v3/mobile-device-prestages/{id}/history
func (s *MobileDevicePrestages) AddHistoryNoteByIDV3(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointMobileDevicePrestagesV3, id)

	var result ResponseAddHistoryNote

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
