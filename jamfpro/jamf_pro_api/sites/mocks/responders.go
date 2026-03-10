package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SitesMock struct {
	*mocks.GenericMock
}

func NewSitesMock() *SitesMock {
	return &SitesMock{
		GenericMock: mocks.NewJSONMock("SitesMock"),
	}
}

func (m *SitesMock) RegisterListV1Mock() {
	m.Register("GET", "/api/v1/sites", 200, "validate_list.json")
}

func (m *SitesMock) RegisterGetObjectsByIDV1Mock() {
	m.Register("GET", "/api/v1/sites/1/objects", 200, "validate_objects.json")
}

func (m *SitesMock) RegisterListV1ErrorMock() {
	m.RegisterError("GET", "/api/v1/sites", 404, "error_not_found.json", "")
}

func (m *SitesMock) RegisterGetObjectsByIDV1ErrorMock() {
	m.RegisterError("GET", "/api/v1/sites/999/objects", 404, "error_not_found.json", "")
}
