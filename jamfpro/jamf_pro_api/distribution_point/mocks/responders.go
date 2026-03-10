package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DistributionPointMock struct {
	*mocks.GenericMock
}

func NewDistributionPointMock() *DistributionPointMock {
	return &DistributionPointMock{
		GenericMock: mocks.NewJSONMock("DistributionPointMock"),
	}
}

func (m *DistributionPointMock) RegisterMocks() {
	m.Register("GET", "/api/v1/distribution-points", 200, "validate_list.json")
	m.Register("POST", "/api/v1/distribution-points", 201, "validate_create.json")
	m.Register("POST", "/api/v1/distribution-points/delete-multiple", 204, "")
	m.Register("GET", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.Register("PUT", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/distribution-points/1", 204, "")
	m.Register("PATCH", "/api/v1/distribution-points/1", 200, "validate_get.json")
	m.Register("GET", "/api/v1/distribution-points/1/history", 200, "validate_history.json")
	m.Register("POST", "/api/v1/distribution-points/1/history", 201, "validate_history_note.json")
}

func (m *DistributionPointMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/distribution-points/999", 404, "error_not_found.json", "")
}

func (m *DistributionPointMock) RegisterListInvalidMock() {
	m.Register("GET", "/api/v1/distribution-points", 200, "validate_list_invalid.json")
}

func (m *DistributionPointMock) RegisterHistoryInvalidMock() {
	m.Register("GET", "/api/v1/distribution-points/1/history", 200, "validate_history_invalid.json")
}

