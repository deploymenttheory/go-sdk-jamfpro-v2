package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SLASAMock struct {
	*mocks.GenericMock
}

func NewSLASAMock() *SLASAMock {
	return &SLASAMock{
		GenericMock: mocks.NewJSONMock("SLASAMock"),
	}
}

func (m *SLASAMock) RegisterGetStatusAcceptedMock() {
	m.Register("GET", "/api/v1/slasa", 200, "validate_get_status_accepted.json")
}

func (m *SLASAMock) RegisterGetStatusNotAcceptedMock() {
	m.Register("GET", "/api/v1/slasa", 200, "validate_get_status_not_accepted.json")
}

func (m *SLASAMock) RegisterAcceptMock() {
	m.Register("POST", "/api/v1/slasa", 200, "")
}

func (m *SLASAMock) RegisterGetStatusErrorMock() {
	m.RegisterError("GET", "/api/v1/slasa", 500, "", "mock error")
}

func (m *SLASAMock) RegisterGetStatusNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/slasa", 500, "error_internal.json", "no response registered")
}

func (m *SLASAMock) RegisterAcceptNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v1/slasa", 500, "error_internal.json", "no response registered")
}
