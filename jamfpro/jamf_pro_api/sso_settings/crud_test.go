package sso_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sso_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*SsoSettings, *mocks.SsoSettingsMock) {
	t.Helper()
	mock := mocks.NewSsoSettingsMock()
	return NewSsoSettings(mock), mock
}

func TestUnit_SsoSettings_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.False(t, result.SsoEnabled)
	assert.Equal(t, "OIDC", result.ConfigurationType)
	require.NotNil(t, result.OidcSettings)
	assert.Equal(t, "USERNAME", result.OidcSettings.UserMapping)
}

func TestUnit_SsoSettings_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceSsoSettings{
		SsoEnabled:       false,
		ConfigurationType: "OIDC",
		SsoBypassAllowed: true,
	}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_SsoSettings_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SsoSettings_GetEnrollmentCustomizationDependencies_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDependenciesMock()

	result, resp, err := svc.GetEnrollmentCustomizationDependenciesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Dependencies, 1)
	assert.Equal(t, "Enrollment Customization A", result.Dependencies[0].Name)
	assert.Equal(t, "Enrollment Customization A", result.Dependencies[0].HumanReadableName)
}

func TestUnit_SsoSettings_DisableV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDisableMock()

	resp, err := svc.DisableV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_SsoSettings_GetHistoryV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV3(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_SsoSettings_AddHistoryNoteV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteMock()

	req := &AddHistoryNoteRequest{Note: "Test note"}
	result, resp, err := svc.AddHistoryNoteV3(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
}

func TestUnit_SsoSettings_AddHistoryNoteV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNoteV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SsoSettings_DownloadMetadataV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDownloadMetadataMock()

	data, resp, err := svc.DownloadMetadataV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	_ = data
}

func TestUnit_SsoSettings_GetV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetErrorMock()

	result, resp, err := svc.GetV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_UpdateV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateErrorMock()

	request := &ResourceSsoSettings{SsoEnabled: false, ConfigurationType: "OIDC"}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_GetEnrollmentCustomizationDependenciesV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDependenciesErrorMock()

	result, resp, err := svc.GetEnrollmentCustomizationDependenciesV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_DisableV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDisableErrorMock()

	resp, err := svc.DisableV3(context.Background())
	require.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_GetHistoryV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryErrorMock()

	result, resp, err := svc.GetHistoryV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_AddHistoryNoteV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteErrorMock()

	req := &AddHistoryNoteRequest{Note: "Test note"}
	result, resp, err := svc.AddHistoryNoteV3(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}

func TestUnit_SsoSettings_DownloadMetadataV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDownloadMetadataErrorMock()

	data, resp, err := svc.DownloadMetadataV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, data)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
}
