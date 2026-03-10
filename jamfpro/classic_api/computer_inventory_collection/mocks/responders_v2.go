package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// ComputerInventoryCollectionMockV2 demonstrates the new generic mock pattern for Classic API (XML).
// This replaces the old 157-line responders.go with just ~20 lines.
type ComputerInventoryCollectionMockV2 struct {
	*mocks.GenericMock
}

// NewComputerInventoryCollectionMockV2 creates a new mock for computer inventory collection testing.
func NewComputerInventoryCollectionMockV2() *ComputerInventoryCollectionMockV2 {
	return &ComputerInventoryCollectionMockV2{
		GenericMock: mocks.NewXMLMock("ComputerInventoryCollectionMock"),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerInventoryCollectionMockV2) RegisterMocks() {
	m.RegisterGetMock()
	m.RegisterUpdateMock()
}

func (m *ComputerInventoryCollectionMockV2) RegisterGetMock() {
	m.Register("GET", "/JSSResource/computerinventorycollection", 200, "validate_get.xml")
}

func (m *ComputerInventoryCollectionMockV2) RegisterUpdateMock() {
	m.Register("PUT", "/JSSResource/computerinventorycollection", 200, "validate_update.xml")
}
