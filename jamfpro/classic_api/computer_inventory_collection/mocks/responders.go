package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerInventoryCollectionMock struct {
	*mocks.GenericMock
}

func NewComputerInventoryCollectionMock() *ComputerInventoryCollectionMock {
	return &ComputerInventoryCollectionMock{
		GenericMock: mocks.NewXMLMock("ComputerInventoryCollectionMock"),
	}
}

func (m *ComputerInventoryCollectionMock) RegisterMocks() {
	m.RegisterGetMock()
	m.RegisterUpdateMock()
}

func (m *ComputerInventoryCollectionMock) RegisterGetMock() {
	m.Register("GET", "/JSSResource/computerinventorycollection", 200, "validate_get.xml")
}

func (m *ComputerInventoryCollectionMock) RegisterUpdateMock() {
	m.Register("PUT", "/JSSResource/computerinventorycollection", 200, "validate_update.xml")
}

