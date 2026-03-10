package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudDistributionPointMock struct {
	*mocks.GenericMock
}

func NewCloudDistributionPointMock() *CloudDistributionPointMock {
	return &CloudDistributionPointMock{
		GenericMock: mocks.NewJSONMock("CloudDistributionPointMock"),
	}
}

func (m *CloudDistributionPointMock) RegisterMocks() {
	m.Register("GET", "/api/v1/cloud-distribution-point", 200, "validate_get.json")
	m.Register("POST", "/api/v1/cloud-distribution-point", 201, "validate_get.json")
	m.Register("PATCH", "/api/v1/cloud-distribution-point", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/cloud-distribution-point", 204, "")
	m.Register("GET", "/api/v1/cloud-distribution-point/upload-capability", 200, "validate_upload_capability.json")
	m.Register("GET", "/api/v1/cloud-distribution-point/test-connection", 200, "validate_test_connection.json")
	m.Register("GET", "/api/v1/cloud-distribution-point/history", 200, "validate_history.json")
	m.Register("GET", "/api/v1/cloud-distribution-point/files", 200, "validate_files.json")
	m.Register("POST", "/api/v1/cloud-distribution-point/history", 201, "validate_history_note.json")
	m.Register("POST", "/api/v1/cloud-distribution-point/fail-upload/test-id", 204, "")
	m.Register("POST", "/api/v1/cloud-distribution-point/refresh-inventory", 200, "")
}

func (m *CloudDistributionPointMock) RegisterHistoryInvalidMock() {
	m.Register("GET", "/api/v1/cloud-distribution-point/history", 200, "validate_history_invalid.json")
}

func (m *CloudDistributionPointMock) RegisterFilesInvalidMock() {
	m.Register("GET", "/api/v1/cloud-distribution-point/files", 200, "validate_files_invalid.json")
}
