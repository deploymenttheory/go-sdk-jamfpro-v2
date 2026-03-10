package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EbooksMock struct {
	*mocks.GenericMock
}

func NewEbooksMock() *EbooksMock {
	return &EbooksMock{
		GenericMock: mocks.NewXMLMock("EbooksMock"),
	}
}

func (m *EbooksMock) RegisterMocks() {
	m.RegisterListEbooksMock()
	m.RegisterGetEbookByIDMock()
	m.RegisterGetEbookByNameMock()
	m.RegisterGetEbookByNameAndSubsetMock()
	m.RegisterCreateEbookMock()
	m.RegisterUpdateEbookByIDMock()
	m.RegisterUpdateEbookByNameMock()
	m.RegisterDeleteEbookByIDMock()
	m.RegisterDeleteEbookByNameMock()
}

func (m *EbooksMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *EbooksMock) RegisterListEbooksMock() {
	m.Register("GET", "/JSSResource/ebooks", 200, "validate_list_ebooks.xml")
}

func (m *EbooksMock) RegisterGetEbookByIDMock() {
	m.Register("GET", "/JSSResource/ebooks/id/1", 200, "validate_get_ebook.xml")
}

func (m *EbooksMock) RegisterGetEbookByNameMock() {
	m.Register("GET", "/JSSResource/ebooks/name/Sample Ebook", 200, "validate_get_ebook.xml")
}

func (m *EbooksMock) RegisterGetEbookByNameAndSubsetMock() {
	m.Register("GET", "/JSSResource/ebooks/name/Sample Ebook/subset/General", 200, "validate_get_ebook.xml")
}

func (m *EbooksMock) RegisterCreateEbookMock() {
	m.Register("POST", "/JSSResource/ebooks/id/0", 201, "validate_create_ebook.xml")
}

func (m *EbooksMock) RegisterUpdateEbookByIDMock() {
	m.Register("PUT", "/JSSResource/ebooks/id/1", 200, "validate_update_ebook.xml")
}

func (m *EbooksMock) RegisterUpdateEbookByNameMock() {
	m.Register("PUT", "/JSSResource/ebooks/name/Sample Ebook", 200, "validate_update_ebook.xml")
}

func (m *EbooksMock) RegisterDeleteEbookByIDMock() {
	m.Register("DELETE", "/JSSResource/ebooks/id/1", 200, "")
}

func (m *EbooksMock) RegisterDeleteEbookByNameMock() {
	m.Register("DELETE", "/JSSResource/ebooks/name/Sample Ebook", 200, "")
}

func (m *EbooksMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/ebooks/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *EbooksMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/ebooks/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An ebook with that name already exists")
}

