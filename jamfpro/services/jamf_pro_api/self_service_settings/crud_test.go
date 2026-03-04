package self_service_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SelfServiceSettingsMock) {
	t.Helper()
	mock := mocks.NewSelfServiceSettingsMock()
	return NewService(mock), mock
}

func TestUnit_SelfServiceSettings_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "USER", result.LoginSettings.UserLoginLevel)
	assert.True(t, result.ConfigurationSettings.NotificationsEnabled)
}

func TestUnit_SelfServiceSettings_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceSelfServiceSettings{
		InstallSettings: InstallSettings{InstallAutomatically: false, InstallLocation: "/Applications"},
		LoginSettings:   LoginSettings{UserLoginLevel: "USER", AllowRememberMe: true, UseFido2: false, AuthType: "JAMF"},
		ConfigurationSettings: ConfigurationSettings{
			NotificationsEnabled: true, AlertUserApprovedMdm: false, DefaultLandingPage: "HOME",
			DefaultHomeCategoryId: 0, BookmarksName: "Bookmarks",
		},
	}
	result, resp, err := svc.Update(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_SelfServiceSettings_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Update(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SelfServiceSettings_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Settings updated", result.Results[0].Note)
}

func TestUnit_SelfServiceSettings_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Updated login settings"}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 42, result.ID)
	assert.Equal(t, "/api/v1/self-service/settings/history/42", result.Href)
}

func TestUnit_SelfServiceSettings_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_SelfServiceSettings_AddHistoryNotesV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotesV1(context.Background(), &AddHistoryNotesRequest{Note: ""})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_SelfServiceSettings_Get_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered - dispatch returns nil, errNoMockRegistered

	result, resp, err := svc.Get(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_SelfServiceSettings_Update_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered for Update

	request := &ResourceSelfServiceSettings{
		InstallSettings: InstallSettings{InstallAutomatically: false, InstallLocation: "/Applications"},
		LoginSettings:   LoginSettings{UserLoginLevel: "USER", AllowRememberMe: true, UseFido2: false, AuthType: "JAMF"},
		ConfigurationSettings: ConfigurationSettings{
			NotificationsEnabled: true, AlertUserApprovedMdm: false, DefaultLandingPage: "HOME",
			DefaultHomeCategoryId: 0, BookmarksName: "Bookmarks",
		},
	}
	result, resp, err := svc.Update(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_SelfServiceSettings_GetHistoryV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered for GetHistoryV1

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get self service settings history")
}

func TestUnit_SelfServiceSettings_GetHistoryV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryInvalidMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to get self service settings history")
}

func TestUnit_SelfServiceSettings_AddHistoryNotesV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered for AddHistoryNotesV1

	req := &AddHistoryNotesRequest{Note: "test note"}
	result, resp, err := svc.AddHistoryNotesV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
