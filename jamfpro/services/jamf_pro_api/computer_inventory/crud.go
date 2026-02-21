package computer_inventory

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInventoryServiceInterface defines the interface for Computer Inventory operations.
	// Uses v1 API for all operations. Manages computer inventory records, FileVault, attachments, and device management.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory
	ComputerInventoryServiceInterface interface {
		// ListV1 returns a paginated list of computer inventory records (Get Computer Inventory).
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery (keys: filter, sort, page, page-size, section).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryList, *interfaces.Response, error)

		// GetByIDV1 returns the specified computer inventory by ID (Get Computer Inventory by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id
		GetByIDV1(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified computer inventory by ID using merge-patch semantics (Update Computer Inventory).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-computers-inventory-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceComputerInventory) (*ResourceComputerInventory, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified computer inventory by ID (Delete Computer Inventory).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computers-inventory-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// ListFileVaultV1 returns a paginated list of FileVault inventory records (Get FileVault Inventory).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-filevault
		ListFileVaultV1(ctx context.Context, rsqlQuery map[string]string) (*FileVaultInventoryList, *interfaces.Response, error)

		// GetFileVaultByIDV1 returns FileVault details for the specified computer by ID (Get FileVault Inventory by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id-filevault
		GetFileVaultByIDV1(ctx context.Context, id string) (*FileVaultInventory, *interfaces.Response, error)

		// GetRecoveryLockPasswordByIDV1 returns the recovery lock password for the specified computer by ID (Get Recovery Lock Password).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id-view-recovery-lock-password
		GetRecoveryLockPasswordByIDV1(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *interfaces.Response, error)

		// DeleteAttachmentByIDV1 deletes a computer attachment by computer ID and attachment ID (Delete Computer Attachment).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computers-inventory-id-attachments-attachmentid
		DeleteAttachmentByIDV1(ctx context.Context, computerID, attachmentID string) (*interfaces.Response, error)

		// RemoveMDMProfileByIDV1 removes the MDM profile from a computer by its ID (Remove MDM Profile).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computers-inventory-id-remove-mdm-profile
		RemoveMDMProfileByIDV1(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *interfaces.Response, error)

		// EraseByIDV1 erases a computer by its ID (Erase Computer).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computers-inventory-id-erase
		EraseByIDV1(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*interfaces.Response, error)
	}

	// Service handles communication with the Computer Inventory-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerInventoryServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns a paginated list of computer inventory records.
// URL: GET /api/v1/computers-inventory
// rsqlQuery supports: filter (RSQL), sort, page, page-size, section (all optional).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryList, *interfaces.Response, error) {
	var result ResponseComputerInventoryList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointComputerInventoryV1, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified computer inventory by ID.
// URL: GET /api/v1/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceComputerInventory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerInventoryV1, id)

	query := make(map[string]string)
	for _, section := range ComputerInventorySections {
		query["section"] = section
	}

	var result ResourceComputerInventory

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified computer inventory by ID using merge-patch semantics.
// URL: PATCH /api/v1/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-computers-inventory-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *ResourceComputerInventory) (*ResourceComputerInventory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerInventoryV1, id)

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

// DeleteByIDV1 removes the specified computer inventory by ID.
// URL: DELETE /api/v1/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computers-inventory-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointComputerInventoryV1, id)

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

// ListFileVaultV1 returns a paginated list of FileVault inventory records.
// URL: GET /api/v1/computers-inventory/filevault
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-filevault
func (s *Service) ListFileVaultV1(ctx context.Context, rsqlQuery map[string]string) (*FileVaultInventoryList, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/filevault", EndpointComputerInventoryV1)

	var result FileVaultInventoryList

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetFileVaultByIDV1 returns FileVault details for the specified computer by ID.
// URL: GET /api/v1/computers-inventory/{id}/filevault
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id-filevault
func (s *Service) GetFileVaultByIDV1(ctx context.Context, id string) (*FileVaultInventory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/filevault", EndpointComputerInventoryV1, id)

	var result FileVaultInventory

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

// GetRecoveryLockPasswordByIDV1 returns the recovery lock password for the specified computer by ID.
// URL: GET /api/v1/computers-inventory/{id}/view-recovery-lock-password
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computers-inventory-id-view-recovery-lock-password
func (s *Service) GetRecoveryLockPasswordByIDV1(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/view-recovery-lock-password", EndpointComputerInventoryV1, id)

	var result ResponseRecoveryLockPassword

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

// DeleteAttachmentByIDV1 deletes a computer attachment by computer ID and attachment ID.
// URL: DELETE /api/v1/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-computers-inventory-id-attachments-attachmentid
func (s *Service) DeleteAttachmentByIDV1(ctx context.Context, computerID, attachmentID string) (*interfaces.Response, error) {
	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, fmt.Errorf("attachmentID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments/%s", EndpointComputerInventoryV1, computerID, attachmentID)

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

// RemoveMDMProfileByIDV1 removes the MDM profile from a computer by its ID.
// URL: POST /api/v1/computers-inventory/{id}/remove-mdm-profile
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computers-inventory-id-remove-mdm-profile
func (s *Service) RemoveMDMProfileByIDV1(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/remove-mdm-profile", EndpointComputerInventoryV1, id)

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
// URL: POST /api/v1/computers-inventory/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computers-inventory-id-erase
func (s *Service) EraseByIDV1(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/erase", EndpointComputerInventoryV1, id)

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
