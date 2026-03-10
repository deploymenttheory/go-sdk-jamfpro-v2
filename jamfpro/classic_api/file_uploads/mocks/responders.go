package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type FileUploadsMock struct {
	*mocks.GenericMock
}

func NewFileUploadsMock() *FileUploadsMock {
	return &FileUploadsMock{
		GenericMock: mocks.NewXMLMock("FileUploadsMock"),
	}
}

func (m *FileUploadsMock) RegisterMocks() {
	m.RegisterCreateAttachmentMock()
}

func (m *FileUploadsMock) RegisterErrorMocks() {
	m.RegisterInvalidResourceErrorMock()
	m.RegisterPeripheralsNameErrorMock()
}

func (m *FileUploadsMock) RegisterCreateAttachmentMock() {
	m.Register("POST", "/JSSResource/fileuploads/policies/id/1", 200, "")
}

func (m *FileUploadsMock) RegisterCreateAttachmentMockForPath(path string) {
	m.Register("POST", path, 200, "")
}

func (m *FileUploadsMock) RegisterInvalidResourceErrorMock() {
	m.RegisterError("POST", "/JSSResource/fileuploads/invalidresource/id/1", 400, "error_invalid_resource.xml", "Jamf Pro Classic API error (400): Invalid resource type")
}

func (m *FileUploadsMock) RegisterPeripheralsNameErrorMock() {
	m.RegisterError("POST", "/JSSResource/fileuploads/peripherals/name/somename", 400, "error_peripherals_name_type.xml", "Jamf Pro Classic API error (400): Peripherals only support ID type")
}

