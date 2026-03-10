package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DSSDeclarationsMock struct {
	*mocks.GenericMock
}

func NewDSSDeclarationsMock() *DSSDeclarationsMock {
	return &DSSDeclarationsMock{
		GenericMock: mocks.NewJSONMock("DSSDeclarationsMock"),
	}
}

func (m *DSSDeclarationsMock) RegisterGetByUUIDMock(uuid string) {
	m.Register("GET", "/api/v1/dss-declarations/"+uuid, 200, "validate_get.json")
}
