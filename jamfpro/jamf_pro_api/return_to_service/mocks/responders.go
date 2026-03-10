package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ReturnToServiceMock struct {
	*mocks.GenericMock
}

func NewReturnToServiceMock() *ReturnToServiceMock {
	return &ReturnToServiceMock{
		GenericMock: mocks.NewJSONMock("ReturnToServiceMock"),
	}
}

func (m *ReturnToServiceMock) RegisterListMock() {
	m.Register("GET", "/api/v1/return-to-service", 200, "validate_list.json")
}

func (m *ReturnToServiceMock) RegisterGetByIDMock() {
	m.Register("GET", "/api/v1/return-to-service/1", 200, "validate_get.json")
}

func (m *ReturnToServiceMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/return-to-service", 201, "validate_create.json")
}

func (m *ReturnToServiceMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/return-to-service/1", 200, "validate_update.json")
}

func (m *ReturnToServiceMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/return-to-service/1", 204, "")
}

func (m *ReturnToServiceMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/return-to-service", 500, "error_api.json", "")
}

func (m *ReturnToServiceMock) RegisterGetByIDErrorMock() {
	m.RegisterError("GET", "/api/v1/return-to-service/1", 500, "error_api.json", "")
}

func (m *ReturnToServiceMock) RegisterCreateErrorMock() {
	m.RegisterError("POST", "/api/v1/return-to-service", 500, "error_api.json", "")
}

func (m *ReturnToServiceMock) RegisterUpdateErrorMock() {
	m.RegisterError("PUT", "/api/v1/return-to-service/1", 500, "error_api.json", "")
}

func (m *ReturnToServiceMock) RegisterDeleteErrorMock() {
	m.RegisterError("DELETE", "/api/v1/return-to-service/1", 500, "error_api.json", "")
}

func (m *ReturnToServiceMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/return-to-service", 500, "error_internal.json", "")
}
