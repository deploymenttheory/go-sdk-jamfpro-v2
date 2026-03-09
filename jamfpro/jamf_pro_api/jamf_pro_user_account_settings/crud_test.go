package jamf_pro_user_account_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_user_account_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testKeyID = "testkey"

func setupMockService(t *testing.T) (*JamfProUserAccountSettings, *mocks.UserAccountSettingsMock) {
	t.Helper()
	mock := mocks.NewUserAccountSettingsMock()
	return NewJamfProUserAccountSettings(mock), mock
}

func TestUnit_JamfProUserAccountSettings_NewService(t *testing.T) {
	mock := mocks.NewUserAccountSettingsMock()
	svc := NewJamfProUserAccountSettings(mock)
	require.NotNil(t, svc)
}

func TestUnit_JamfProUserAccountSettings_GetSettingsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSettingsV1Mock(testKeyID)

	result, resp, err := svc.GetSettingsV1(context.Background(), testKeyID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "admin", result.Username)
	assert.Equal(t, "testkey", result.Key)
	assert.Equal(t, []string{"value1", "value2"}, result.Values)
}

func TestUnit_JamfProUserAccountSettings_GetSettingsV1_EmptyKeyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSettingsV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "keyId is required")
}

func TestUnit_JamfProUserAccountSettings_GetSettingsV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock(testKeyID)

	result, resp, err := svc.GetSettingsV1(context.Background(), testKeyID)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_JamfProUserAccountSettings_GetSettingsV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSettingsV1(context.Background(), testKeyID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfProUserAccountSettings_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV1Mock(testKeyID)

	result, resp, err := svc.GetV1(context.Background(), testKeyID)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "myvalue", result)
}

func TestUnit_JamfProUserAccountSettings_GetV1_PlainTextResponse(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV1PlainTextMock(testKeyID)

	result, resp, err := svc.GetV1(context.Background(), testKeyID)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "plaintext-value", result)
}

func TestUnit_JamfProUserAccountSettings_GetV1_EmptyKeyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV1(context.Background(), "")
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "keyId is required")
}

func TestUnit_JamfProUserAccountSettings_GetV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock(testKeyID)

	result, resp, err := svc.GetV1(context.Background(), testKeyID)
	assert.Error(t, err)
	assert.Empty(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_JamfProUserAccountSettings_GetV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV1(context.Background(), testKeyID)
	assert.Error(t, err)
	assert.Empty(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfProUserAccountSettings_PutV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPutV1Mock(testKeyID)

	values := RequestUserPreferences{"setting1": "value1", "setting2": "value2"}
	resp, err := svc.PutV1(context.Background(), testKeyID, values)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_JamfProUserAccountSettings_PutV1_EmptyKeyID(t *testing.T) {
	svc, _ := setupMockService(t)

	values := RequestUserPreferences{"key": "value"}
	resp, err := svc.PutV1(context.Background(), "", values)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "keyId is required")
}

func TestUnit_JamfProUserAccountSettings_PutV1_NilValues(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.PutV1(context.Background(), testKeyID, nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "values is required")
}

func TestUnit_JamfProUserAccountSettings_PutV1_EmptyValues(t *testing.T) {
	svc, _ := setupMockService(t)

	values := RequestUserPreferences{}
	resp, err := svc.PutV1(context.Background(), testKeyID, values)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "values is required")
}

func TestUnit_JamfProUserAccountSettings_PutV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock(testKeyID)

	values := RequestUserPreferences{"key": "value"}
	resp, err := svc.PutV1(context.Background(), testKeyID, values)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_JamfProUserAccountSettings_PutV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	values := RequestUserPreferences{"key": "value"}
	resp, err := svc.PutV1(context.Background(), testKeyID, values)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfProUserAccountSettings_DeleteV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteV1Mock(testKeyID)

	resp, err := svc.DeleteV1(context.Background(), testKeyID)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_JamfProUserAccountSettings_DeleteV1_EmptyKeyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "keyId is required")
}

func TestUnit_JamfProUserAccountSettings_DeleteV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock(testKeyID)

	resp, err := svc.DeleteV1(context.Background(), testKeyID)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_JamfProUserAccountSettings_DeleteV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteV1(context.Background(), testKeyID)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}
