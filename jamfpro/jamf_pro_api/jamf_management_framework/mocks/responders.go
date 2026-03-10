package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfManagementFrameworkMock struct {
	*mocks.GenericMock
}

func NewJamfManagementFrameworkMock() *JamfManagementFrameworkMock {
	return &JamfManagementFrameworkMock{
		GenericMock: mocks.NewJSONMock("JamfManagementFrameworkMock"),
	}
}

func (m *JamfManagementFrameworkMock) RegisterRedeployMock(computerID string) {
	m.Register("POST", "/api/v1/jamf-management-framework/redeploy/"+computerID, 201, "validate_redeploy.json")
}

func (m *JamfManagementFrameworkMock) RegisterNotFoundErrorMock(computerID string) {
	m.RegisterError("POST", "/api/v1/jamf-management-framework/redeploy/"+computerID, 404, "error_not_found.json", "")
}
