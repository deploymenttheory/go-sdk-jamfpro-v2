package computer_inventory

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Computer Inventory-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
	ComputerInventory struct {
		client client.Client
	}
)

func NewComputerInventory(client client.Client) *ComputerInventory {
	return &ComputerInventory{client: client}
}

// deprecatedV3Replacement is the migration hint logged for the deprecated V3
// (and legacy V1 action) computer-inventory methods.
const deprecatedV3Replacement = "use the v4 computers-inventory endpoints (\u2026V4 methods)"

// CreateV3 creates a new computer inventory record.
// URL: POST /api/v3/computers-inventory
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory
//
// Deprecated: deprecated in Jamf Pro 11.30; use CreateV4.
func (s *ComputerInventory) CreateV3(ctx context.Context, request *ResourceComputerInventory) (*CreateComputerResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.CreateV3", "11.30", deprecatedV3Replacement)

	endpoint := constants.EndpointJamfProComputerInventoryV3

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

// ListV3 returns all computer inventory records using automatic pagination.
// URL: GET /api/v3/computers-inventory
// rsqlQuery supports: filter (RSQL), sort, section (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
//
// Deprecated: deprecated in Jamf Pro 11.30; use ListV4.
func (s *ComputerInventory) ListV3(ctx context.Context, rsqlQuery map[string]string) (*ResponseComputerInventoryList, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.ListV3", "11.30", deprecatedV3Replacement)

	endpoint := constants.EndpointJamfProComputerInventoryV3

	var result ResponseComputerInventoryList

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceComputerInventory
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

// GetByIDV3 returns the specified computer inventory by ID.
// URL: GET /api/v3/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetByIDV4.
func (s *ComputerInventory) GetByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceComputerInventory

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDetailByIDV3 returns all sections of a computer.
// URL: GET /api/v3/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-detail-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetDetailByIDV4.
func (s *ComputerInventory) GetDetailByIDV3(ctx context.Context, id string) (*ResourceComputerInventory, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetDetailByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s-detail/%s", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	var result ResourceComputerInventory

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV3 updates the specified computer inventory by ID using merge-patch semantics.
// URL: PATCH /api/v3/computers-inventory-detail/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v3-computers-inventory-detail-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use UpdateByIDV4.
func (s *ComputerInventory) UpdateByIDV3(ctx context.Context, id string, request *ResourceComputerInventory) (*ResourceComputerInventory, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.UpdateByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s-detail/%s", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceComputerInventory

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

// DeleteByIDV3 removes the specified computer inventory by ID.
// URL: DELETE /api/v3/computers-inventory/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use DeleteByIDV4.
func (s *ComputerInventory) DeleteByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.DeleteByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use ListFileVaultV4.
func (s *ComputerInventory) ListFileVaultV3(ctx context.Context) (*FileVaultInventoryList, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.ListFileVaultV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/filevault", constants.EndpointJamfProComputerInventoryV3)

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

// GetFileVaultByIDV3 returns FileVault details for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/filevault
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-filevault
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetFileVaultByIDV4.
func (s *ComputerInventory) GetFileVaultByIDV3(ctx context.Context, id string) (*FileVaultInventory, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetFileVaultByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/filevault", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

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

// GetDeviceLockPinByIDV3 returns the device lock PIN for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/view-device-lock-pin
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-device-lock-pin
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetDeviceLockPinByIDV4.
func (s *ComputerInventory) GetDeviceLockPinByIDV3(ctx context.Context, id string) (*ResponseDeviceLockPin, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetDeviceLockPinByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/view-device-lock-pin", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

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

// GetRecoveryLockPasswordByIDV3 returns the recovery lock password for the specified computer by ID.
// URL: GET /api/v3/computers-inventory/{id}/view-recovery-lock-password
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-view-recovery-lock-password
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetRecoveryLockPasswordByIDV4.
func (s *ComputerInventory) GetRecoveryLockPasswordByIDV3(ctx context.Context, id string) (*ResponseRecoveryLockPassword, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetRecoveryLockPasswordByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/view-recovery-lock-password", constants.EndpointJamfProComputerInventoryV3, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

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

// UploadAttachmentByIDV3 uploads an attachment and assigns it to a computer.
// URL: POST /api/v3/computers-inventory/{id}/attachments
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-computers-inventory-id-attachments
//
// Deprecated: deprecated in Jamf Pro 11.30; use UploadAttachmentByIDV4.
func (s *ComputerInventory) UploadAttachmentByIDV3(ctx context.Context, computerID string, attachment []byte) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.UploadAttachmentByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/attachments", constants.EndpointJamfProComputerInventoryV3, computerID)

	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if len(attachment) == 0 {
		return nil, fmt.Errorf("attachment data is required")
	}

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

// GetAttachmentByIDV3 downloads a computer attachment by computer ID and attachment ID.
// URL: GET /api/v3/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory-id-attachments-attachmentid
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetAttachmentByIDV4.
func (s *ComputerInventory) GetAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) ([]byte, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.GetAttachmentByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/attachments/%s", constants.EndpointJamfProComputerInventoryV3, computerID, attachmentID)

	if computerID == "" {
		return nil, nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, nil, fmt.Errorf("attachmentID is required")
	}

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

// DeleteAttachmentByIDV3 deletes a computer attachment by computer ID and attachment ID.
// URL: DELETE /api/v3/computers-inventory/{id}/attachments/{attachmentId}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-computers-inventory-id-attachments-attachmentid
//
// Deprecated: deprecated in Jamf Pro 11.30; use DeleteAttachmentByIDV4.
func (s *ComputerInventory) DeleteAttachmentByIDV3(ctx context.Context, computerID, attachmentID string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.DeleteAttachmentByIDV3", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/attachments/%s", constants.EndpointJamfProComputerInventoryV3, computerID, attachmentID)

	if computerID == "" {
		return nil, fmt.Errorf("computerID is required")
	}
	if attachmentID == "" {
		return nil, fmt.Errorf("attachmentID is required")
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveMDMProfileByIDV1 removes the MDM profile from a computer by its ID.
// URL: POST /api/v1/computer-inventory/{id}/remove-mdm-profile
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-remove-mdm-profile
//
// Deprecated: deprecated in Jamf Pro 11.30; use RemoveMDMProfileByIDV4.
func (s *ComputerInventory) RemoveMDMProfileByIDV1(ctx context.Context, id string) (*ResponseRemoveMDMProfile, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.RemoveMDMProfileByIDV1", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/remove-mdm-profile", constants.EndpointJamfProComputerInventoryV1, id)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

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

// EraseByIDV1 erases a computer by its ID.
// URL: POST /api/v1/computer-inventory/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-computer-inventory-id-erase
//
// Deprecated: deprecated in Jamf Pro 11.30; use EraseByIDV4.
func (s *ComputerInventory) EraseByIDV1(ctx context.Context, id string, request *RequestEraseDeviceComputer) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/computer_inventory.ComputerInventory.EraseByIDV1", "11.30", deprecatedV3Replacement)

	endpoint := fmt.Sprintf("%s/%s/erase", constants.EndpointJamfProComputerInventoryV1, id)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

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
