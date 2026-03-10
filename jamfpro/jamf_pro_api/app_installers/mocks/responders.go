package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AppInstallersMock struct {
	*mocks.GenericMock
}

func NewAppInstallersMock() *AppInstallersMock {
	return &AppInstallersMock{
		GenericMock: mocks.NewJSONMock("AppInstallersMock"),
	}
}

func (m *AppInstallersMock) RegisterMocks() {
	m.Register("GET", "/api/v1/app-installers/titles", 200, "validate_list_titles.json")
	m.Register("GET", "/api/v1/app-installers/titles/1", 200, "validate_get_title.json")
	m.Register("GET", "/api/v1/app-installers/deployments", 200, "validate_list_deployments.json")
	m.Register("GET", "/api/v1/app-installers/deployments/1", 200, "validate_get_deployment.json")
	m.Register("POST", "/api/v1/app-installers/deployments", 201, "validate_create_deployment.json")
	m.Register("PUT", "/api/v1/app-installers/deployments/1", 200, "validate_get_deployment.json")
	m.Register("DELETE", "/api/v1/app-installers/deployments/1", 204, "")
}

func (m *AppInstallersMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v1/app-installers/titles", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/app-installers/titles/1", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/app-installers/deployments", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/app-installers/deployments/1", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/app-installers/deployments", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/app-installers/deployments/1", 500, "error_internal.json", "")
	m.RegisterError("DELETE", "/api/v1/app-installers/deployments/1", 500, "error_internal.json", "")
}
