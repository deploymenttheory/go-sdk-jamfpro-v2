package computer_inventory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// -----------------------------------------------------------------------------
// Computer Inventory CRUD (V4) — Jamf Pro 11.30, replaces the V3 surface.
//
// V4 also absorbs the two actions that previously lived on the singular
// /api/v1/computer-inventory path: erase and remove-mdm-profile.
// -----------------------------------------------------------------------------

// CreateV4 creates a new computer inventory record.
// URL: POST /api/v4/computers-inventory
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-computers-inventory
func (s *ComputerInventory) CreateV4(ctx context.Context, request *ResourceComputerInventoryV4) (*CreateComputerResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProComputerInventoryV4

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateComputerResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ListV4 returns all computer inventory records using automatic pagination.
// URL: GET /api/v4/computers-inventory
// rsqlQuery supports: filter (RSQL), sort, section (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory
func (s *ComputerInventory) ListV4(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryListV4, *resty.Response, error) {
	endpoint := constants.EndpointJamfProComputerInventoryV4

	var result ResponseComputerInventoryListV4

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceComputerInventoryV4
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV4 returns the specified computer inventory by ID.
// URL: GET /api/v4/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-id
func (s *ComputerInventory) GetByIDV4(ctx context.Context, id string) (*ResourceComputerInventoryV4, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResourceComputerInventoryV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDetailByIDV4 returns all sections of a computer.
// URL: GET /api/v4/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-detail-id
func (s *ComputerInventory) GetDetailByIDV4(ctx context.Context, id string) (*ResourceComputerInventoryV4, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s-detail/%s", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResourceComputerInventoryV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV4 updates the specified computer inventory by ID using merge-patch semantics.
// URL: PATCH /api/v4/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v4-computers-inventory-detail-id
func (s *ComputerInventory) UpdateByIDV4(ctx context.Context, id string, request *ResourceComputerInventoryV4) (*ResourceComputerInventoryV4, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s-detail/%s", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResourceComputerInventoryV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Patch(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV4 removes the specified computer inventory by ID.
// URL: DELETE /api/v4/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v4-computers-inventory-id
func (s *ComputerInventory) DeleteByIDV4(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerInventoryV4, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListFileVaultV4 returns all FileVault inventory records using automatic pagination.
// URL: GET /api/v4/computers-inventory/filevault
// Note: This endpoint only supports pagination (page, page-size), which is managed internally.
// No RSQL filtering or sorting is available for this endpoint.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-filevault
func (s *ComputerInventory) ListFileVaultV4(ctx context.Context) (*FileVaultInventoryList, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/filevault", constants.EndpointJamfProComputerInventoryV4)

	var result FileVaultInventoryList

	mergePage := func(pageData []byte) error {
		var pageResults []FileVaultInventory
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetFileVaultByIDV4 returns FileVault details for the specified computer by ID.
// URL: GET /api/v4/computers-inventory/{id}/filevault
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-id-filevault
func (s *ComputerInventory) GetFileVaultByIDV4(ctx context.Context, id string) (*FileVaultInventory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/filevault", constants.EndpointJamfProComputerInventoryV4, id)

	var result FileVaultInventory

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeviceLockPinByIDV4 returns the device lock PIN for the specified computer by ID.
// URL: GET /api/v4/computers-inventory/{id}/view-device-lock-pin
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-id-view-device-lock-pin
func (s *ComputerInventory) GetDeviceLockPinByIDV4(ctx context.Context, id string) (*ResponseDeviceLockPin, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/view-device-lock-pin", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResponseDeviceLockPin

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetRecoveryLockPasswordByIDV4 returns the recovery lock password for the specified computer by ID.
// URL: GET /api/v4/computers-inventory/{id}/view-recovery-lock-password
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-id-view-recovery-lock-password
func (s *ComputerInventory) GetRecoveryLockPasswordByIDV4(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/view-recovery-lock-password", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResponseRecoveryLockPassword

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadAttachmentByIDV4 uploads an attachment and assigns it to a computer.
// URL: POST /api/v4/computers-inventory/{id}/attachments
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-computers-inventory-id-attachments
func (s *ComputerInventory) UploadAttachmentByIDV4(ctx context.Context, computerID string, attachment []byte) (*resty.Response, error) {
	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if len(attachment) == 0 {
		return nil, fmt.Errorf("attachment data is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments", constants.EndpointJamfProComputerInventoryV4, computerID)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationOctetStream).
		SetBody(attachment).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetAttachmentByIDV4 downloads a computer attachment by computer ID and attachment ID.
// URL: GET /api/v4/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-computers-inventory-id-attachments-attachmentid
func (s *ComputerInventory) GetAttachmentByIDV4(ctx context.Context, computerID, attachmentID string) ([]byte, *resty.Response, error) {
	if computerID == "" {
		return nil, nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, nil, fmt.Errorf("attachmentID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments/%s", constants.EndpointJamfProComputerInventoryV4, computerID, attachmentID)

	var result []byte

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationOctetStream).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteAttachmentByIDV4 deletes a computer attachment by computer ID and attachment ID.
// URL: DELETE /api/v4/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v4-computers-inventory-id-attachments-attachmentid
func (s *ComputerInventory) DeleteAttachmentByIDV4(ctx context.Context, computerID, attachmentID string) (*resty.Response, error) {
	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, fmt.Errorf("attachmentID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/attachments/%s", constants.EndpointJamfProComputerInventoryV4, computerID, attachmentID)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveMDMProfileByIDV4 removes the MDM profile from a computer by its ID.
// URL: POST /api/v4/computers-inventory/{id}/remove-mdm-profile
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-computers-inventory-id-remove-mdm-profile
func (s *ComputerInventory) RemoveMDMProfileByIDV4(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/remove-mdm-profile", constants.EndpointJamfProComputerInventoryV4, id)

	var result ResponseRemoveMDMProfile

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// EraseByIDV4 erases a computer by its ID.
// URL: POST /api/v4/computers-inventory/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v4-computers-inventory-id-erase
func (s *ComputerInventory) EraseByIDV4(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/erase", constants.EndpointJamfProComputerInventoryV4, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
