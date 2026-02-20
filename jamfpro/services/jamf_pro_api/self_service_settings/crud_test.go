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

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "USER", result.LoginSettings.UserLoginLevel)
	assert.True(t, result.ConfigurationSettings.NotificationsEnabled)
}

func TestUnitUpdate_Success(t *testing.T) {
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
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Update(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
