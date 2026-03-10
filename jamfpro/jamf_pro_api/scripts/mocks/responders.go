package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ScriptsMock struct {
	*mocks.GenericMock
}

func NewScriptsMock() *ScriptsMock {
	return &ScriptsMock{
		GenericMock: mocks.NewJSONMock("ScriptsMock"),
	}
}

func (m *ScriptsMock) RegisterListScriptsMock() {
	m.Register("GET", "/api/v1/scripts", 200, "validate_list_scripts.json")
}

func (m *ScriptsMock) RegisterListScriptsRSQLMock() {
	m.Register("GET", "/api/v1/scripts", 200, "validate_list_scripts_rsql.json")
}

func (m *ScriptsMock) RegisterGetScriptMock() {
	m.Register("GET", "/api/v1/scripts/1", 200, "validate_get_script.json")
}

func (m *ScriptsMock) RegisterDownloadScriptMock() {
	m.Register("GET", "/api/v1/scripts/1/download", 200, "validate_download_script.txt")
}

func (m *ScriptsMock) RegisterCreateScriptMock() {
	m.Register("POST", "/api/v1/scripts", 201, "validate_create_script.json")
}

func (m *ScriptsMock) RegisterUpdateScriptMock() {
	m.Register("PUT", "/api/v1/scripts/1", 200, "validate_update_script.json")
}

func (m *ScriptsMock) RegisterDeleteScriptMock() {
	m.Register("DELETE", "/api/v1/scripts/1", 204, "")
}

func (m *ScriptsMock) RegisterGetScriptHistoryMock() {
	m.Register("GET", "/api/v1/scripts/1/history", 200, "validate_get_history.json")
}

func (m *ScriptsMock) RegisterAddScriptHistoryNotesMock() {
	m.Register("POST", "/api/v1/scripts/1/history", 201, "")
}

func (m *ScriptsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/scripts/999", 404, "error_not_found.json", "")
}

func (m *ScriptsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/scripts", 409, "error_conflict.json", "")
}
