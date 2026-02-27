package computer_inventory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInventoryServiceInterface defines the interface for Computer Inventory operations.
	// Uses v3 API for most operations; v1 API for device commands (erase, remove MDM profile).
	// Manages computer inventory records, FileVault, attachments, and device management.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
	ComputerInventoryServiceInterface interface {
		// ListV3 returns all computer inventory records using automatic pagination (Get Computer Inventory).
		//
		// Supports optional RSQL filtering via rsqlQuery (keys: filter, sort, section).
		// Note: page and page-size are managed internally; all pages are fetched automatically.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
		ListV3(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryList, *interfaces.Response, error)

		// CreateV3 creates a new computer inventory record (Create Computer Inventory record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory
		CreateV3(ctx context.Context, request *ResourceComputerInventory) (*CreateComputerResponse, *interfaces.Response, error)

		// GetByIDV3 returns the specified computer inventory by ID (Get Computer Inventory by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id
		GetByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error)

		// GetDetailByIDV3 returns all sections of a computer (Get all sections of a computer).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-detail-id
		GetDetailByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error)

		// UpdateByIDV3 updates the specified computer inventory by ID using merge-patch semantics (Update Computer Inventory).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-computers-inventory-detail-id
		UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerInventory) (*ResourceComputerInventory, *interfaces.Response, error)

		// DeleteByIDV3 removes the specified computer inventory by ID (Delete Computer Inventory).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id
		DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error)

		// ListFileVaultV3 returns all FileVault inventory records using automatic pagination (Get FileVault Inventory).
		//
		// Note: This endpoint only supports pagination (page, page-size), which is managed internally.
		// No RSQL filtering or sorting is available for this endpoint.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-filevault
		ListFileVaultV3(ctx context.Context) (*FileVaultInventoryList, *interfaces.Response, error)

		// GetFileVaultByIDV3 returns FileVault details for the specified computer by ID (Get FileVault Inventory by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-filevault
		GetFileVaultByIDV3(ctx context.Context, id string) (*FileVaultInventory, *interfaces.Response, error)

		// GetRecoveryLockPasswordByIDV3 returns the recovery lock password for the specified computer by ID (Get Recovery Lock Password).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-recovery-lock-password
		GetRecoveryLockPasswordByIDV3(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *interfaces.Response, error)

		// UploadAttachmentByIDV3 uploads an attachment and assigns it to a computer (Upload attachment).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory-id-attachments
		UploadAttachmentByIDV3(ctx context.Context, computerID string, attachment []byte) (*interfaces.Response, error)

		// GetAttachmentByIDV3 downloads a computer attachment by computer ID and attachment ID (Download attachment file).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-attachments-attachmentid
		GetAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) ([]byte, *interfaces.Response, error)

		// DeleteAttachmentByIDV3 deletes a computer attachment by computer ID and attachment ID (Delete Computer Attachment).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id-attachments-attachmentid
		DeleteAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) (*interfaces.Response, error)

		// GetDeviceLockPinByIDV3 returns the device lock PIN for the specified computer by ID (Get Device Lock PIN).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-device-lock-pin
		GetDeviceLockPinByIDV3(ctx context.Context, id string) (*ResponseDeviceLockPin, *interfaces.Response, error)

		// RemoveMDMProfileByIDV1 removes the MDM profile from a computer by its ID (Remove MDM Profile).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-remove-mdm-profile
		RemoveMDMProfileByIDV1(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *interfaces.Response, error)

		// EraseByIDV1 erases a computer by its ID (Erase Computer).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-erase
		EraseByIDV1(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*interfaces.Response, error)
	}

	// Service handles communication with the Computer Inventory-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerInventoryServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// CreateV3 creates a new computer inventory record.
// URL: POST /api/v3/computers-inventory
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory
func (s *Service) CreateV3(ctx context.Context, request *ResourceComputerInventory) (*CreateComputerResponse, *interfaces.Response, error) {
	endpoint := EndpointComputerInventoryV3

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateComputerResponse

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

// ListV3 returns all computer inventory records using automatic pagination.
// URL: GET /api/v3/computers-inventory
// rsqlQuery supports: filter (RSQL), sort, section (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
func (s *Service) ListV3(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryList, *interfaces.Response, error) {
	endpoint := EndpointComputerInventoryV3

	var result ResponseComputerInventoryList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageResponse ResponseComputerInventoryList
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV3 returns the specified computer inventory by ID.
// URL: GET /api/v3/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id
func (s *Service) GetByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	query := make(map[string]string)
	for _, section := range ComputerInventorySections {
		query["section"] = section
	}

	var result ResourceComputerInventory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDetailByIDV3 returns all sections of a computer.
// URL: GET /api/v3/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-detail-id
func (s *Service) GetDetailByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s-detail/%s", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceComputerInventory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the specified computer inventory by ID using merge-patch semantics.
// URL: PATCH /api/v3/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-computers-inventory-detail-id
func (s *Service) UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerInventory) (*ResourceComputerInventory, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s-detail/%s", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceComputerInventory

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV3 removes the specified computer inventory by ID.
// URL: DELETE /api/v3/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id
func (s *Service) DeleteByIDV3(ctx context.Context, id string) (*interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListFileVaultV3 returns all FileVault inventory records using automatic pagination.
// URL: GET /api/v3/computers-inventory/filevault
// Note: This endpoint only supports pagination (page, page-size), which is managed internally.
// No RSQL filtering or sorting is available for this endpoint.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-filevault
func (s *Service) ListFileVaultV3(ctx context.Context) (*FileVaultInventoryList, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/filevault", EndpointComputerInventoryV3)

	var result FileVaultInventoryList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageResponse FileVaultInventoryList
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetFileVaultByIDV3 returns FileVault details for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/filevault
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-filevault
func (s *Service) GetFileVaultByIDV3(ctx context.Context, id string) (*FileVaultInventory, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/filevault", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result FileVaultInventory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeviceLockPinByIDV3 returns the device lock PIN for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/view-device-lock-pin
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-device-lock-pin
func (s *Service) GetDeviceLockPinByIDV3(ctx context.Context, id string) (*ResponseDeviceLockPin, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/view-device-lock-pin", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResponseDeviceLockPin

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetRecoveryLockPasswordByIDV3 returns the recovery lock password for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/view-recovery-lock-password
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-recovery-lock-password
func (s *Service) GetRecoveryLockPasswordByIDV3(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/view-recovery-lock-password", EndpointComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResponseRecoveryLockPassword

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadAttachmentByIDV3 uploads an attachment and assigns it to a computer.
// URL: POST /api/v3/computers-inventory/{id}/attachments
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory-id-attachments
func (s *Service) UploadAttachmentByIDV3(ctx context.Context, computerID string, attachment []byte) (*interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/attachments", EndpointComputerInventoryV3, computerID)

	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if len(attachment) == 0 {
		return nil, fmt.Errorf("attachment data is required")
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationOctetStream,
	}

	resp, err := s.client.Post(ctx, endpoint, attachment, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetAttachmentByIDV3 downloads a computer attachment by computer ID and attachment ID.
// URL: GET /api/v3/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-attachments-attachmentid
func (s *Service) GetAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) ([]byte, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/attachments/%s", EndpointComputerInventoryV3, computerID, attachmentID)

	if computerID == "" {
		return nil, nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, nil, fmt.Errorf("attachmentID is required")
	}

	var result []byte

	headers := map[string]string{
		"Accept": mime.ApplicationOctetStream,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteAttachmentByIDV3 deletes a computer attachment by computer ID and attachment ID.
// URL: DELETE /api/v3/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id-attachments-attachmentid
func (s *Service) DeleteAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) (*interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/attachments/%s", EndpointComputerInventoryV3, computerID, attachmentID)

	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, fmt.Errorf("attachmentID is required")
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveMDMProfileByIDV1 removes the MDM profile from a computer by its ID.
// URL: POST /api/v1/computer-inventory/{id}/remove-mdm-profile
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-remove-mdm-profile
func (s *Service) RemoveMDMProfileByIDV1(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/remove-mdm-profile", EndpointComputerInventoryV1, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResponseRemoveMDMProfile

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

// EraseByIDV1 erases a computer by its ID.
// URL: POST /api/v1/computer-inventory/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-erase
func (s *Service) EraseByIDV1(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s/erase", EndpointComputerInventoryV1, id)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

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
