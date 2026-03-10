package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DiskEncryptionConfigurationsMock struct {
	*mocks.GenericMock
}

func NewDiskEncryptionConfigurationsMock() *DiskEncryptionConfigurationsMock {
	return &DiskEncryptionConfigurationsMock{
		GenericMock: mocks.NewXMLMock("DiskEncryptionConfigurationsMock"),
	}
}

func (m *DiskEncryptionConfigurationsMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *DiskEncryptionConfigurationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DiskEncryptionConfigurationsMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/diskencryptionconfigurations", 200, "validate_list_disk_encryption_configurations.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/diskencryptionconfigurations/id/1", 200, "validate_get_disk_encryption_configuration.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "validate_get_disk_encryption_configuration.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/diskencryptionconfigurations/id/0", 201, "validate_create_disk_encryption_configuration.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/diskencryptionconfigurations/id/1", 200, "validate_update_disk_encryption_configuration.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "validate_update_disk_encryption_configuration.xml")
}

func (m *DiskEncryptionConfigurationsMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/diskencryptionconfigurations/id/1", 200, "")
}

func (m *DiskEncryptionConfigurationsMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "")
}

func (m *DiskEncryptionConfigurationsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/diskencryptionconfigurations/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *DiskEncryptionConfigurationsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/diskencryptionconfigurations/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A disk encryption configuration with that name already exists")
}

