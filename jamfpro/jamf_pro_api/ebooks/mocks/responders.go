package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EbooksMock struct {
	*mocks.GenericMock
}

func NewEbooksMock() *EbooksMock {
	return &EbooksMock{
		GenericMock: mocks.NewJSONMock("EbooksMock"),
	}
}

func (m *EbooksMock) RegisterListEbooksMock() {
	m.Register("GET", "/api/v1/ebooks", 200, "validate_list.json")
}

func (m *EbooksMock) RegisterGetEbookMock() {
	m.Register("GET", "/api/v1/ebooks/1", 200, "validate_get.json")
}

func (m *EbooksMock) RegisterGetEbookScopeMock() {
	m.Register("GET", "/api/v1/ebooks/1/scope", 200, "validate_get_scope.json")
}

func (m *EbooksMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/ebooks/999", 404, "error_not_found.json", "")
	m.RegisterError("GET", "/api/v1/ebooks/999/scope", 404, "error_not_found.json", "")
}

func (m *EbooksMock) RegisterInvalidJSONMock() {
	m.Register("GET", "/api/v1/ebooks", 200, "validate_list_invalid.json")
}
