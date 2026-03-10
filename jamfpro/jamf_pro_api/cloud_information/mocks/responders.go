package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudInformationMock struct {
	*mocks.GenericMock
}

func NewCloudInformationMock() *CloudInformationMock {
	return &CloudInformationMock{
		GenericMock: mocks.NewJSONMock("CloudInformationMock"),
	}
}

func (m *CloudInformationMock) RegisterGetCloudInformationMock() {
	m.Register("GET", "/api/v1/cloud-information", 200, "validate_get.json")
}

func (m *CloudInformationMock) RegisterGetCloudInformationErrorMock() {
	m.RegisterError("GET", "/api/v1/cloud-information", 500, "error_internal.json", "mock client error")
}
