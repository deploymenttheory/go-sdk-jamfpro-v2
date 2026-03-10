package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type IconsMock struct {
	*mocks.GenericMock
}

func NewIconsMock() *IconsMock {
	return &IconsMock{
		GenericMock: mocks.NewJSONMock("IconsMock"),
	}
}

func (m *IconsMock) RegisterMocks() {
	m.Register("GET", "/api/v1/icon/1", 200, "validate_get.json")
	m.Register("POST", "/api/v1/icon", 201, "validate_upload.json")
	m.Register("GET", "/api/v1/icon/download/1", 200, "validate_download.bin")
}

func (m *IconsMock) RegisterGetByIDErrorMock(id int) {
	m.RegisterError("GET", "/api/v1/icon/999", 500, "error_internal.json", "")
}

func (m *IconsMock) RegisterUploadErrorMock() {
	m.RegisterError("POST", "/api/v1/icon", 500, "error_internal.json", "")
}
