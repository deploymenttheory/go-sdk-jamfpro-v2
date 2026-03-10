package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerInventoryCollectionSettingsMock struct {
	*mocks.GenericMock
}

func NewComputerInventoryCollectionSettingsMock() *ComputerInventoryCollectionSettingsMock {
	return &ComputerInventoryCollectionSettingsMock{
		GenericMock: mocks.NewJSONMock("ComputerInventoryCollectionSettingsMock"),
	}
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v2/computer-inventory-collection-settings", 200, "validate_get.json")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterUpdateMock() {
	m.Register("PATCH", "/api/v2/computer-inventory-collection-settings", 204, "")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterCreateCustomPathMock() {
	m.Register("POST", "/api/v2/computer-inventory-collection-settings/custom-path", 201, "validate_create_custom_path.json")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterDeleteCustomPathMock(id string) {
	m.Register("DELETE", "/api/v2/computer-inventory-collection-settings/custom-path/"+id, 204, "")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-inventory-collection-settings", 500, "error_internal.json", "no mock registered")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterUpdateErrorMock() {
	m.RegisterError("PATCH", "/api/v2/computer-inventory-collection-settings", 500, "error_internal.json", "no mock registered")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterCreateCustomPathErrorMock() {
	m.RegisterError("POST", "/api/v2/computer-inventory-collection-settings/custom-path", 500, "error_internal.json", "no mock registered")
}

func (m *ComputerInventoryCollectionSettingsMock) RegisterDeleteCustomPathErrorMock(id string) {
	m.RegisterError("DELETE", "/api/v2/computer-inventory-collection-settings/custom-path/"+id, 500, "error_internal.json", "no mock registered")
}
