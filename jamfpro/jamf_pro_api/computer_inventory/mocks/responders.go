package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerInventoryMock struct {
	*mocks.GenericMock
}

func NewComputerInventoryMock() *ComputerInventoryMock {
	return &ComputerInventoryMock{
		GenericMock: mocks.NewJSONMock("ComputerInventoryMock"),
	}
}

func (m *ComputerInventoryMock) RegisterListMock() {
	m.Register("GET", "/api/v3/computers-inventory", 200, "validate_list.json")
}

func (m *ComputerInventoryMock) RegisterCreateMock() {
	m.Register("POST", "/api/v3/computers-inventory", 201, "validate_create.json")
}

func (m *ComputerInventoryMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v3/computers-inventory/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterGetDetailByIDMock(id string) {
	m.Register("GET", "/api/v3/computers-inventory-detail/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterUpdateByIDMock(id string) {
	m.Register("PATCH", "/api/v3/computers-inventory-detail/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v3/computers-inventory/"+id, 204, "")
}

func (m *ComputerInventoryMock) RegisterListFileVaultMock() {
	m.Register("GET", "/api/v3/computers-inventory/filevault", 200, "validate_filevault_list.json")
}

func (m *ComputerInventoryMock) RegisterGetFileVaultByIDMock(id string) {
	m.Register("GET", "/api/v3/computers-inventory/"+id+"/filevault", 200, "validate_filevault.json")
}

func (m *ComputerInventoryMock) RegisterGetRecoveryLockPasswordByIDMock(id string) {
	m.Register("GET", "/api/v3/computers-inventory/"+id+"/view-recovery-lock-password", 200, "validate_recovery_lock.json")
}

func (m *ComputerInventoryMock) RegisterUploadAttachmentMock(computerID string) {
	m.Register("POST", "/api/v3/computers-inventory/"+computerID+"/attachments", 201, "validate_attachment.json")
}

func (m *ComputerInventoryMock) RegisterGetAttachmentMock(computerID, attachmentID string) {
	path := "/api/v3/computers-inventory/" + computerID + "/attachments/" + attachmentID
	m.Register("GET", path, 200, "validate_attachment_get.json")
}

func (m *ComputerInventoryMock) RegisterDeleteAttachmentMock(computerID, attachmentID string) {
	m.Register("DELETE", "/api/v3/computers-inventory/"+computerID+"/attachments/"+attachmentID, 204, "")
}

func (m *ComputerInventoryMock) RegisterGetDeviceLockPinMock(id string) {
	m.Register("GET", "/api/v3/computers-inventory/"+id+"/view-device-lock-pin", 200, "validate_device_lock_pin.json")
}

func (m *ComputerInventoryMock) RegisterRemoveMDMProfileMock(id string) {
	m.Register("POST", "/api/v1/computer-inventory/"+id+"/remove-mdm-profile", 200, "validate_remove_mdm.json")
}

func (m *ComputerInventoryMock) RegisterEraseMock(id string) {
	m.Register("POST", "/api/v1/computer-inventory/"+id+"/erase", 204, "")
}

func (m *ComputerInventoryMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v3/computers-inventory", 500, "", "simulated ListV3 API error")
}

func (m *ComputerInventoryMock) RegisterListFileVaultErrorMock() {
	m.RegisterError("GET", "/api/v3/computers-inventory/filevault", 500, "", "simulated ListFileVault API error")
}

func (m *ComputerInventoryMock) RegisterListInvalidJSONMock() {
	m.Register("GET", "/api/v3/computers-inventory", 200, "validate_list_invalid.json")
}

func (m *ComputerInventoryMock) RegisterListFileVaultInvalidJSONMock() {
	m.Register("GET", "/api/v3/computers-inventory/filevault", 200, "validate_filevault_list_invalid.json")
}

func (m *ComputerInventoryMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/v3/computers-inventory", 500, "", "simulated CreateV3 API error")
}

func (m *ComputerInventoryMock) RegisterGetByIDErrorMock(id string) {
	m.RegisterError("GET", "/api/v3/computers-inventory/"+id, 500, "", "simulated GetByID API error")
}

func (m *ComputerInventoryMock) RegisterDeleteErrorMock(id string) {
	m.RegisterError("DELETE", "/api/v3/computers-inventory/"+id, 500, "", "simulated Delete API error")
}
