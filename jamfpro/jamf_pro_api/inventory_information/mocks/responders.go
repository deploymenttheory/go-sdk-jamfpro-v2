package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type InventoryInformationMock struct {
	*mocks.GenericMock
}

func NewInventoryInformationMock() *InventoryInformationMock {
	return &InventoryInformationMock{
		GenericMock: mocks.NewJSONMock("InventoryInformationMock"),
	}
}

func (m *InventoryInformationMock) RegisterGetV1Mock() {
	m.Register("GET", "/api/v1/inventory-information", 200, "validate_get.json")
}

func (m *InventoryInformationMock) RegisterGetV1NotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/inventory-information", 404, "error_not_found.json", "")
}

func (m *InventoryInformationMock) RegisterInvalidJSONMock() {
	m.Register("GET", "/api/v1/inventory-information", 200, "validate_get_invalid.json")
}
