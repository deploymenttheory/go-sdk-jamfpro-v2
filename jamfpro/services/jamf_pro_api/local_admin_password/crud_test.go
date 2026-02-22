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

func TestUnitGetPendingRotationsV2_Success(t *testing.T) {
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

func TestUnitGetSettingsV2_Success(t *testing.T) {
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

func TestUnitUpdateSettingsV2_Success(t *testing.T) {
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

func TestUnitUpdateSettingsV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateSettingsV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "settings is required")
}

func TestUnitGetPasswordHistoryByClientManagementIDV2_Success(t *testing.T) {
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

func TestUnitGetPasswordHistoryByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "", "admin")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnitGetPasswordHistoryByClientManagementIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPasswordHistoryByClientManagementIDV2(context.Background(), "device-001", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnitGetCurrentPasswordByClientManagementIDV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCurrentPasswordMock()

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "device-001", "admin")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "SecureP@ssw0rd123!", result.Password)
}

func TestUnitGetCurrentPasswordByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "", "admin")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnitGetCurrentPasswordByClientManagementIDV2_EmptyUsername(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCurrentPasswordByClientManagementIDV2(context.Background(), "device-001", "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "username is required")
}

func TestUnitGetFullHistoryByClientManagementIDV2_Success(t *testing.T) {
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

func TestUnitGetFullHistoryByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetFullHistoryByClientManagementIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnitGetCapableAccountsByClientManagementIDV2_Success(t *testing.T) {
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

func TestUnitGetCapableAccountsByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementID is required")
}

func TestUnitGetCapableAccountsByClientManagementIDV2_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetCapableAccountsByClientManagementIDV2(context.Background(), "device-999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitSetPasswordByClientManagementIDV2_Success(t *testing.T) {
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

func TestUnitSetPasswordByClientManagementIDV2_EmptyClientManagementID(t *testing.T) {
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

func TestUnitSetPasswordByClientManagementIDV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.SetPasswordByClientManagementIDV2(context.Background(), "device-001", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "passwordList is required")
}
