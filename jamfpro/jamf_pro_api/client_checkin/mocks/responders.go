package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ClientCheckinMock struct {
	*mocks.GenericMock
}

func NewClientCheckinMock() *ClientCheckinMock {
	return &ClientCheckinMock{
		GenericMock: mocks.NewJSONMock("ClientCheckinMock"),
	}
}

func (m *ClientCheckinMock) RegisterMocks() {
	m.Register("GET", "/api/v3/check-in", 200, "validate_get.json")
	m.Register("PUT", "/api/v3/check-in", 200, "validate_get.json")
	m.Register("GET", "/api/v3/check-in/history", 200, "validate_history_get.json")
	m.Register("POST", "/api/v3/check-in/history", 201, "validate_history_post.json")
}

func (m *ClientCheckinMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v3/check-in", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v3/check-in", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v3/check-in/history", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v3/check-in/history", 500, "error_internal.json", "")
}
