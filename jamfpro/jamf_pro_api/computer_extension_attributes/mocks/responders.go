package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerExtensionAttributesMock struct {
	*mocks.GenericMock
}

func NewComputerExtensionAttributesMock() *ComputerExtensionAttributesMock {
	return &ComputerExtensionAttributesMock{
		GenericMock: mocks.NewJSONMock("ComputerExtensionAttributesMock"),
	}
}

func (m *ComputerExtensionAttributesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateMock()
	m.RegisterDeleteMock()
	m.RegisterDeleteMultipleMock()
}

func (m *ComputerExtensionAttributesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *ComputerExtensionAttributesMock) RegisterListMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes", 200, "validate_list.json")
}

func (m *ComputerExtensionAttributesMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/1", 200, "validate_get.json")
}

func (m *ComputerExtensionAttributesMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/computer-extension-attributes", 201, "validate_create.json")
}

func (m *ComputerExtensionAttributesMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/computer-extension-attributes/1", 200, "validate_update.json")
}

func (m *ComputerExtensionAttributesMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/computer-extension-attributes/1", 204, "")
}

func (m *ComputerExtensionAttributesMock) RegisterDeleteMultipleMock() {
	m.Register("POST", "/api/v1/computer-extension-attributes/delete-multiple", 204, "")
}

func (m *ComputerExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/computer-extension-attributes/999", 404, "error_not_found.json", "")
}

func (m *ComputerExtensionAttributesMock) RegisterHistoryMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/1/history", 200, "validate_history.json")
}

func (m *ComputerExtensionAttributesMock) RegisterAddHistoryNoteMock() {
	m.Register("POST", "/api/v1/computer-extension-attributes/1/history", 201, "")
}

func (m *ComputerExtensionAttributesMock) RegisterListTemplatesMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/templates", 200, "validate_templates_list.json")
}

func (m *ComputerExtensionAttributesMock) RegisterListTemplatesInvalidMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/templates", 200, "validate_templates_list_invalid.json")
}

func (m *ComputerExtensionAttributesMock) RegisterGetTemplateMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/templates/1", 200, "validate_template_get.json")
}

func (m *ComputerExtensionAttributesMock) RegisterUploadMock() {
	m.Register("POST", "/api/v1/computer-extension-attributes/upload", 200, "validate_get.json")
}

func (m *ComputerExtensionAttributesMock) RegisterDataDependencyMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/1/data-dependency", 200, "validate_data_dependency.json")
}

func (m *ComputerExtensionAttributesMock) RegisterDownloadMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/1/download", 200, "validate_download.xml")
}

func (m *ComputerExtensionAttributesMock) RegisterListInvalidMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes", 200, "validate_list_invalid.json")
}

func (m *ComputerExtensionAttributesMock) RegisterHistoryInvalidMock() {
	m.Register("GET", "/api/v1/computer-extension-attributes/1/history", 200, "validate_history_invalid.json")
}
