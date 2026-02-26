package local_admin_password

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/local_admin_password/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.LocalAdminPasswordMock) {
	t.Helper()
	mock := mocks.NewLocalAdminPasswordMock()
	return NewService(mock), mock
}

func TestUnit_LocalAdminPassword_GetPendingRotationsV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPendingRotationsMock()

	result, resp, err := svc.GetPendingRotationsV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "device-001", result.Results[0].LapsUser.ClientManagementID)
	assert.Equal(t, "admin", result.Results[0].LapsUser.Username)
}

func TestUnit_LocalAdminPassword_GetSettingsV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSettingsMock()

	result, resp, err := svc.GetSettingsV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.AutoDeployEnabled)
	assert.Equal(t, 90, result.PasswordRotationTime)
	assert.True(t, result.AutoRotateEnabled)
	assert.Equal(t, 7, result.AutoRotateExpirationTime)
}

func TestUnit_LocalAdminPassword_UpdateSettingsV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSettingsMock()

	settings := &SettingsResource{
		AutoDeployEnabled:        true,
		PasswordRotationTime:     60,
		AutoRotateEnabled:        true,
		AutoRotateExpirationTime: 14,
	}
	resp, err := svc.UpdateSettingsV2(context.Background(), settings)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_LocalAdminPassword_UpdateSettingsV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateSettingsV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "settings is required")
}

func TestUnit_LocalAdminPassword_GetPasswordHistoryByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPasswordHistoryMock()

	result, resp, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "device-001", "admin")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "encrypted-password-1", result.Results[0].Password)
	assert.Equal(t, 1, len(result.Results[0].Audits))
	assert.Equal(t, "admin@example.com", result.Results[0].Audits[0].ViewedBy)
}

func TestUnit_LocalAdminPassword_GetPasswordHistoryByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "", "admin")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetPasswordHistoryByClientManagementIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "device-001", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetCurrentPasswordByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCurrentPasswordMock()

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "device-001", "admin")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "SecureP@ssw0rd123!", result.Password)
}

func TestUnit_LocalAdminPassword_GetCurrentPasswordByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "", "admin")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetCurrentPasswordByClientManagementIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "device-001", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetFullHistoryByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetFullHistoryMock()

	result, resp, err := svc.GetFullHistoryByClientManagementIDV2(context.Background(), "device-001")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 4, result.TotalCount)
	require.Len(t, result.Results, 4)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "PASSWORD_VIEWED", result.Results[0].EventType)
	assert.Equal(t, "admin@example.com", result.Results[0].ViewedBy)
}

func TestUnit_LocalAdminPassword_GetFullHistoryByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetFullHistoryByClientManagementIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetCapableAccountsByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCapableAccountsMock()

	result, resp, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "device-001")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "device-001", result.Results[0].ClientManagementID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "LOCAL", result.Results[0].UserSource)
}

func TestUnit_LocalAdminPassword_GetCapableAccountsByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetCapableAccountsByClientManagementIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "device-999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_LocalAdminPassword_SetPasswordByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterSetPasswordMock()

	req := &SetPasswordRequest{
		LapsUserPasswordList: []LapsUserPassword{
			{Username: "admin", Password: "NewSecureP@ss123!"},
			{Username: "localadmin", Password: "AnotherP@ss456!"},
		},
	}
	result, resp, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "device-001", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.LapsUserPasswordList, 2)
	assert.Equal(t, "admin", result.LapsUserPasswordList[0].Username)
	assert.Equal(t, "localadmin", result.LapsUserPasswordList[1].Username)
}

func TestUnit_LocalAdminPassword_SetPasswordByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &SetPasswordRequest{
		LapsUserPasswordList: []LapsUserPassword{{Username: "admin", Password: "P@ss"}},
	}
	result, resp, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_SetPasswordByClientManagementIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "device-001", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "passwordList is required")
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryByUsernameMock()

	result, resp, err := svc.GetHistoryByUsernameV2(context.Background(), "device-001", "admin")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "2024-01-10T12:00:00Z", result.Results[0].CreatedDate)
	assert.Equal(t, "COMPLETED", result.Results[0].RotationStatus)
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByUsernameV2(context.Background(), "", "admin")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByUsernameV2(context.Background(), "device-001", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetAuditByUsernameAndGUIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAuditByUsernameAndGUIDMock()

	result, resp, err := svc.GetAuditByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "guid-123")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "encrypted-password-1", result.Results[0].Password)
	assert.Equal(t, "admin@example.com", result.Results[0].Audits[0].ViewedBy)
}

func TestUnit_LocalAdminPassword_GetAuditByUsernameAndGUIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAuditByUsernameAndGUIDV2(context.Background(), "", "admin", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetAuditByUsernameAndGUIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAuditByUsernameAndGUIDV2(context.Background(), "device-001", "", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetAuditByUsernameAndGUIDV2_EmptyGUID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAuditByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "guid is required")
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameAndGUIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryByUsernameAndGUIDMock()

	result, resp, err := svc.GetHistoryByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "guid-123")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "COMPLETED", result.Results[0].RotationStatus)
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameAndGUIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByUsernameAndGUIDV2(context.Background(), "", "admin", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameAndGUIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByUsernameAndGUIDV2(context.Background(), "device-001", "", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetHistoryByUsernameAndGUIDV2_EmptyGUID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "guid is required")
}

func TestUnit_LocalAdminPassword_GetPasswordByUsernameAndGUIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPasswordByUsernameAndGUIDMock()

	result, resp, err := svc.GetPasswordByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "guid-123")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "SecureP@ssw0rd123!", result.Password)
}

func TestUnit_LocalAdminPassword_GetPasswordByUsernameAndGUIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordByUsernameAndGUIDV2(context.Background(), "", "admin", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnit_LocalAdminPassword_GetPasswordByUsernameAndGUIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordByUsernameAndGUIDV2(context.Background(), "device-001", "", "guid-123")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnit_LocalAdminPassword_GetPasswordByUsernameAndGUIDV2_EmptyGUID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordByUsernameAndGUIDV2(context.Background(), "device-001", "admin", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "guid is required")
}
