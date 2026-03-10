package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MDMMock struct {
	*mocks.GenericMock
}

func NewMDMMock() *MDMMock {
	return &MDMMock{
		GenericMock: mocks.NewJSONMock("MDMMock"),
	}
}

func (m *MDMMock) RegisterListCommandsMock() {
	m.Register("GET", "/api/v2/mdm/commands", 200, "validate_list_commands.json")
}

func (m *MDMMock) RegisterBlankPushMock() {
	m.Register("POST", "/api/v2/mdm/blank-push", 200, "validate_blank_push.json")
}

func (m *MDMMock) RegisterSendCommandMock() {
	m.Register("POST", "/api/v2/mdm/commands", 200, "validate_send_command.json")
}

func (m *MDMMock) RegisterDeployPackageMock() {
	m.Register("POST", "/api/v1/deploy-package?verbose=true", 200, "validate_deploy_package.json")
}

func (m *MDMMock) RegisterRenewProfileMock() {
	m.Register("POST", "/api/v1/mdm/renew-profile", 200, "validate_renew_profile.json")
}

func (m *MDMMock) RegisterNotFoundErrorMock() {
	m.RegisterError("POST", "/api/v2/mdm/commands", 404, "error_not_found.json", "")
}

func (m *MDMMock) RegisterBlankPushErrorMock() {
	m.RegisterError("POST", "/api/v2/mdm/blank-push", 500, "error_not_found.json", "")
}

func (m *MDMMock) RegisterDeployPackageErrorMock() {
	m.RegisterError("POST", "/api/v1/deploy-package?verbose=true", 500, "error_not_found.json", "")
}

func (m *MDMMock) RegisterRenewProfileErrorMock() {
	m.RegisterError("POST", "/api/v1/mdm/renew-profile", 500, "error_not_found.json", "")
}

func (m *MDMMock) RegisterListCommandsErrorMock() {
	m.RegisterError("GET", "/api/v2/mdm/commands", 500, "error_not_found.json", "")
}

func (m *MDMMock) RegisterListCommandsInvalidJSONMock() {
	m.Register("GET", "/api/v2/mdm/commands", 200, "validate_list_commands_invalid.json")
}

func (m *MDMMock) RegisterListCommandsNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/mdm/commands", 500, "error_internal.json", "")
}
