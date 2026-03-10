package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type PatchManagementMock struct {
	*mocks.GenericMock
}

func NewPatchManagementMock() *PatchManagementMock {
	return &PatchManagementMock{
		GenericMock: mocks.NewJSONMock("PatchManagementMock"),
	}
}

func (m *PatchManagementMock) RegisterAcceptDisclaimerMock() {
	m.Register("POST", "/api/v2/patch-management-accept-disclaimer", 200, "validate_accept_disclaimer.json")
}

func (m *PatchManagementMock) RegisterAcceptDisclaimerErrorMock() {
	m.RegisterError("POST", "/api/v2/patch-management-accept-disclaimer", 500, "", "mock error")
}

func (m *PatchManagementMock) RegisterAcceptDisclaimerNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v2/patch-management-accept-disclaimer", 500, "error_internal.json", "no response registered")
}
