package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfProSystemInitializationMock struct {
	*mocks.GenericMock
}

func NewJamfProSystemInitializationMock() *JamfProSystemInitializationMock {
	return &JamfProSystemInitializationMock{
		GenericMock: mocks.NewJSONMock("JamfProSystemInitializationMock"),
	}
}

func (m *JamfProSystemInitializationMock) RegisterInitializeMock() {
	m.Register("POST", "/api/v1/system/initialize", 200, "validate_initialize.json")
}

func (m *JamfProSystemInitializationMock) RegisterInitializeDatabaseConnectionMock() {
	m.Register("POST", "/api/v1/system/initialize-database-connection", 200, "validate_initialize_database_connection.json")
}

func (m *JamfProSystemInitializationMock) RegisterPlatformInitializeMock() {
	m.Register("POST", "/api/v1/system/platform-initialize", 201, "validate_platform_initialize.json")
}

func (m *JamfProSystemInitializationMock) RegisterInitializeErrorMock() {
	m.RegisterError("POST", "/api/v1/system/initialize", 500, "error_internal.json", "")
}

func (m *JamfProSystemInitializationMock) RegisterInitializeDatabaseConnectionErrorMock() {
	m.RegisterError("POST", "/api/v1/system/initialize-database-connection", 500, "error_internal.json", "")
}

func (m *JamfProSystemInitializationMock) RegisterPlatformInitializeErrorMock() {
	m.RegisterError("POST", "/api/v1/system/platform-initialize", 500, "error_internal.json", "")
}
