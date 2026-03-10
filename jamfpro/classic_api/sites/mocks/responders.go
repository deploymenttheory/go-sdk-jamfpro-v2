package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SitesMock struct {
	*mocks.GenericMock
}

func NewSitesMock() *SitesMock {
	return &SitesMock{
		GenericMock: mocks.NewXMLMock("SitesMock"),
	}
}

func (m *SitesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *SitesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SitesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/sites", 200, "validate_list_sites.xml")
}

func (m *SitesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/sites/id/1", 200, "validate_get_site.xml")
}

func (m *SitesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/sites/name/Main Campus", 200, "validate_get_site.xml")
}

func (m *SitesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/sites/id/0", 201, "validate_create_site.xml")
}

func (m *SitesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/sites/id/1", 200, "validate_update_site.xml")
}

func (m *SitesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/sites/name/Main Campus", 200, "validate_update_site.xml")
}

func (m *SitesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/sites/id/1", 200, "")
}

func (m *SitesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/sites/name/Main Campus", 200, "")
}

func (m *SitesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/sites/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *SitesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/sites/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A site with that name already exists")
}

