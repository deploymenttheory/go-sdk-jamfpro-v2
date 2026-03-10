package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type FileShareDistributionPointsMock struct {
	*mocks.GenericMock
}

func NewFileShareDistributionPointsMock() *FileShareDistributionPointsMock {
	return &FileShareDistributionPointsMock{
		GenericMock: mocks.NewXMLMock("FileShareDistributionPointsMock"),
	}
}

func (m *FileShareDistributionPointsMock) RegisterMocks() {
	m.RegisterListFileShareDistributionPointsMock()
	m.RegisterGetFileShareDistributionPointByIDMock()
	m.RegisterGetFileShareDistributionPointByNameMock()
	m.RegisterCreateFileShareDistributionPointMock()
	m.RegisterUpdateFileShareDistributionPointByIDMock()
	m.RegisterUpdateFileShareDistributionPointByNameMock()
	m.RegisterDeleteFileShareDistributionPointByIDMock()
	m.RegisterDeleteFileShareDistributionPointByNameMock()
}

func (m *FileShareDistributionPointsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *FileShareDistributionPointsMock) RegisterListFileShareDistributionPointsMock() {
	m.Register("GET", "/JSSResource/distributionpoints", 200, "validate_list_file_share_distribution_points.xml")
}

func (m *FileShareDistributionPointsMock) RegisterGetFileShareDistributionPointByIDMock() {
	m.Register("GET", "/JSSResource/distributionpoints/id/1", 200, "validate_get_file_share_distribution_point.xml")
}

func (m *FileShareDistributionPointsMock) RegisterGetFileShareDistributionPointByNameMock() {
	m.Register("GET", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "validate_get_file_share_distribution_point.xml")
}

func (m *FileShareDistributionPointsMock) RegisterCreateFileShareDistributionPointMock() {
	m.Register("POST", "/JSSResource/distributionpoints/id/0", 201, "validate_create_file_share_distribution_point.xml")
}

func (m *FileShareDistributionPointsMock) RegisterUpdateFileShareDistributionPointByIDMock() {
	m.Register("PUT", "/JSSResource/distributionpoints/id/1", 200, "validate_update_file_share_distribution_point.xml")
}

func (m *FileShareDistributionPointsMock) RegisterUpdateFileShareDistributionPointByNameMock() {
	m.Register("PUT", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "validate_update_file_share_distribution_point.xml")
}

func (m *FileShareDistributionPointsMock) RegisterDeleteFileShareDistributionPointByIDMock() {
	m.Register("DELETE", "/JSSResource/distributionpoints/id/1", 200, "")
}

func (m *FileShareDistributionPointsMock) RegisterDeleteFileShareDistributionPointByNameMock() {
	m.Register("DELETE", "/JSSResource/distributionpoints/name/Main File Share DP", 200, "")
}

func (m *FileShareDistributionPointsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/distributionpoints/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *FileShareDistributionPointsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/distributionpoints/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A distribution point with that name already exists")
}

