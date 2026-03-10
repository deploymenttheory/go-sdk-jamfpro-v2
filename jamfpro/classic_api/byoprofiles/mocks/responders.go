package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type BYOProfilesMock struct {
	*mocks.GenericMock
}

func NewBYOProfilesMock() *BYOProfilesMock {
	return &BYOProfilesMock{
		GenericMock: mocks.NewXMLMock("BYOProfilesMock"),
	}
}

func (m *BYOProfilesMock) RegisterMocks() {
	m.RegisterListBYOProfilesMock()
	m.RegisterGetBYOProfileByIDMock()
	m.RegisterGetBYOProfileByNameMock()
	m.RegisterCreateBYOProfileMock()
	m.RegisterUpdateBYOProfileByIDMock()
	m.RegisterUpdateBYOProfileByNameMock()
	m.RegisterDeleteBYOProfileByIDMock()
	m.RegisterDeleteBYOProfileByNameMock()
}

func (m *BYOProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *BYOProfilesMock) RegisterListBYOProfilesMock() {
	m.Register("GET", "/JSSResource/byoprofiles", 200, "validate_list_byoprofiles.xml")
}

func (m *BYOProfilesMock) RegisterGetBYOProfileByIDMock() {
	m.Register("GET", "/JSSResource/byoprofiles/id/1", 200, "validate_get_byoprofile.xml")
}

func (m *BYOProfilesMock) RegisterGetBYOProfileByNameMock() {
	m.Register("GET", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "validate_get_byoprofile.xml")
}

func (m *BYOProfilesMock) RegisterCreateBYOProfileMock() {
	m.Register("POST", "/JSSResource/byoprofiles/id/0", 201, "validate_create_byoprofile.xml")
}

func (m *BYOProfilesMock) RegisterUpdateBYOProfileByIDMock() {
	m.Register("PUT", "/JSSResource/byoprofiles/id/1", 200, "validate_update_byoprofile.xml")
}

func (m *BYOProfilesMock) RegisterUpdateBYOProfileByNameMock() {
	m.Register("PUT", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "validate_update_byoprofile.xml")
}

func (m *BYOProfilesMock) RegisterDeleteBYOProfileByIDMock() {
	m.Register("DELETE", "/JSSResource/byoprofiles/id/1", 200, "")
}

func (m *BYOProfilesMock) RegisterDeleteBYOProfileByNameMock() {
	m.Register("DELETE", "/JSSResource/byoprofiles/name/Test BYO Profile 1", 200, "")
}

func (m *BYOProfilesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/byoprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *BYOProfilesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/byoprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A BYO profile with that name already exists")
}

