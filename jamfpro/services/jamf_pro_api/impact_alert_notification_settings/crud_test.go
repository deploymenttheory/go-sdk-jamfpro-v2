package impact_alert_notification_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/impact_alert_notification_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ImpactAlertNotificationSettingsMock) {
	t.Helper()
	mock := mocks.NewImpactAlertNotificationSettingsMock()
	return NewService(mock), mock
}

func TestUnit_ImpactAlertNotificationSettings_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.ScopeableObjectsAlertEnabled)
	assert.False(t, result.ScopeableObjectsConfirmationCodeEnabled)
	assert.True(t, result.DeployableObjectsAlertEnabled)
	assert.False(t, result.DeployableObjectsConfirmationCodeEnabled)
}

func TestUnit_ImpactAlertNotificationSettings_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceImpactAlertNotificationSettings{
		ScopeableObjectsAlertEnabled:             false,
		ScopeableObjectsConfirmationCodeEnabled:  true,
		DeployableObjectsAlertEnabled:            false,
		DeployableObjectsConfirmationCodeEnabled: true,
	}

	resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Update returns 204 No Content
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_ImpactAlertNotificationSettings_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
