package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ClassesMock struct {
	*mocks.GenericMock
}

func NewClassesMock() *ClassesMock {
	return &ClassesMock{
		GenericMock: mocks.NewXMLMock("ClassesMock"),
	}
}

func (m *ClassesMock) RegisterMocks() {
	m.RegisterListClassesMock()
	m.RegisterGetClassByIDMock()
	m.RegisterGetClassByNameMock()
	m.RegisterCreateClassMock()
	m.RegisterUpdateClassByIDMock()
	m.RegisterUpdateClassByNameMock()
	m.RegisterDeleteClassByIDMock()
	m.RegisterDeleteClassByNameMock()
}

func (m *ClassesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *ClassesMock) RegisterListClassesMock() {
	m.Register("GET", "/JSSResource/classes", 200, "validate_list_classes.xml")
}

func (m *ClassesMock) RegisterGetClassByIDMock() {
	m.Register("GET", "/JSSResource/classes/id/1", 200, "validate_get_class.xml")
}

func (m *ClassesMock) RegisterGetClassByNameMock() {
	m.Register("GET", "/JSSResource/classes/name/Test Class 1", 200, "validate_get_class.xml")
}

func (m *ClassesMock) RegisterCreateClassMock() {
	m.Register("POST", "/JSSResource/classes/id/0", 201, "validate_create_class.xml")
}

func (m *ClassesMock) RegisterUpdateClassByIDMock() {
	m.Register("PUT", "/JSSResource/classes/id/1", 200, "validate_update_class.xml")
}

func (m *ClassesMock) RegisterUpdateClassByNameMock() {
	m.Register("PUT", "/JSSResource/classes/name/Test Class 1", 200, "validate_update_class.xml")
}

func (m *ClassesMock) RegisterDeleteClassByIDMock() {
	m.Register("DELETE", "/JSSResource/classes/id/1", 200, "")
}

func (m *ClassesMock) RegisterDeleteClassByNameMock() {
	m.Register("DELETE", "/JSSResource/classes/name/Test Class 1", 200, "")
}

func (m *ClassesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/classes/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *ClassesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/classes/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A class with that name already exists")
}

