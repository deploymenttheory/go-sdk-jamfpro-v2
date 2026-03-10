package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LocalAdminPasswordMock struct {
	*mocks.GenericMock
}

func NewLocalAdminPasswordMock() *LocalAdminPasswordMock {
	return &LocalAdminPasswordMock{
		GenericMock: mocks.NewJSONMock("LocalAdminPasswordMock"),
	}
}

func (m *LocalAdminPasswordMock) RegisterGetPendingRotationsMock() {
	m.Register("GET", "/api/v2/local-admin-password/pending-rotations", 200, "validate_pending_rotations.json")
}

func (m *LocalAdminPasswordMock) RegisterGetSettingsMock() {
	m.Register("GET", "/api/v2/local-admin-password/settings", 200, "validate_get_settings.json")
}

func (m *LocalAdminPasswordMock) RegisterUpdateSettingsMock() {
	m.Register("PUT", "/api/v2/local-admin-password/settings", 200, "")
}

func (m *LocalAdminPasswordMock) RegisterGetPasswordHistoryMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/audit", 200, "validate_password_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetCurrentPasswordMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/password", 200, "validate_current_password.json")
}

func (m *LocalAdminPasswordMock) RegisterGetHistoryByUsernameMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/history", 200, "validate_account_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetAuditByUsernameAndGUIDMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/guid-123/audit", 200, "validate_password_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetHistoryByUsernameAndGUIDMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/guid-123/history", 200, "validate_account_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetPasswordByUsernameAndGUIDMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/account/admin/guid-123/password", 200, "validate_current_password.json")
}

func (m *LocalAdminPasswordMock) RegisterGetFullHistoryMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/history", 200, "validate_full_history.json")
}

func (m *LocalAdminPasswordMock) RegisterGetCapableAccountsMock() {
	m.Register("GET", "/api/v2/local-admin-password/device-001/accounts", 200, "validate_capable_accounts.json")
}

func (m *LocalAdminPasswordMock) RegisterSetPasswordMock() {
	m.Register("PUT", "/api/v2/local-admin-password/device-001/set-password", 200, "validate_set_password.json")
}

func (m *LocalAdminPasswordMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/local-admin-password/device-999/accounts", 404, "error_not_found.json", "")
}
