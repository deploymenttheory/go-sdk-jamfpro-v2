package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LicensedSoftwareMock struct {
	*mocks.GenericMock
}

func NewLicensedSoftwareMock() *LicensedSoftwareMock {
	return &LicensedSoftwareMock{
		GenericMock: mocks.NewXMLMock("LicensedSoftwareMock"),
	}
}

func (m *LicensedSoftwareMock) RegisterMocks() {
	m.RegisterListLicensedSoftwareMock()
	m.RegisterGetLicensedSoftwareByIDMock()
	m.RegisterGetLicensedSoftwareByNameMock()
	m.RegisterCreateLicensedSoftwareMock()
	m.RegisterUpdateLicensedSoftwareByIDMock()
	m.RegisterUpdateLicensedSoftwareByNameMock()
	m.RegisterDeleteLicensedSoftwareByIDMock()
	m.RegisterDeleteLicensedSoftwareByNameMock()
}

func (m *LicensedSoftwareMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *LicensedSoftwareMock) RegisterListLicensedSoftwareMock() {
	m.Register("GET", "/JSSResource/licensedsoftware", 200, "validate_list_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterGetLicensedSoftwareByIDMock() {
	m.Register("GET", "/JSSResource/licensedsoftware/id/1", 200, "validate_get_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterGetLicensedSoftwareByNameMock() {
	m.Register("GET", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "validate_get_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterCreateLicensedSoftwareMock() {
	m.Register("POST", "/JSSResource/licensedsoftware/id/0", 201, "validate_create_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterUpdateLicensedSoftwareByIDMock() {
	m.Register("PUT", "/JSSResource/licensedsoftware/id/1", 200, "validate_update_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterUpdateLicensedSoftwareByNameMock() {
	m.Register("PUT", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "validate_update_licensed_software.xml")
}

func (m *LicensedSoftwareMock) RegisterDeleteLicensedSoftwareByIDMock() {
	m.Register("DELETE", "/JSSResource/licensedsoftware/id/1", 200, "")
}

func (m *LicensedSoftwareMock) RegisterDeleteLicensedSoftwareByNameMock() {
	m.Register("DELETE", "/JSSResource/licensedsoftware/name/Sample Licensed Software", 200, "")
}

func (m *LicensedSoftwareMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/licensedsoftware/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *LicensedSoftwareMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/licensedsoftware/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): Licensed software with that name already exists")
}

