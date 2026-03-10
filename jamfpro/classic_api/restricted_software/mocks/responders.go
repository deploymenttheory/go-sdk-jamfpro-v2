package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type RestrictedSoftwareMock struct {
	*mocks.GenericMock
}

func NewRestrictedSoftwareMock() *RestrictedSoftwareMock {
	return &RestrictedSoftwareMock{
		GenericMock: mocks.NewXMLMock("RestrictedSoftwareMock"),
	}
}

func (m *RestrictedSoftwareMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *RestrictedSoftwareMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *RestrictedSoftwareMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/restrictedsoftware", 200, "validate_list_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/restrictedsoftware/id/1", 200, "validate_get_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/restrictedsoftware/name/Calculator", 200, "validate_get_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/restrictedsoftware/id/0", 201, "validate_create_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/restrictedsoftware/id/1", 200, "validate_update_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/restrictedsoftware/name/Calculator", 200, "validate_update_restricted_software.xml")
}

func (m *RestrictedSoftwareMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/restrictedsoftware/id/1", 200, "")
}

func (m *RestrictedSoftwareMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/restrictedsoftware/name/Calculator", 200, "")
}

func (m *RestrictedSoftwareMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/restrictedsoftware/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *RestrictedSoftwareMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/restrictedsoftware/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A restricted software with that name already exists")
}
