package jamf_pro_notifications

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_notifications/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.NotificationsMock) {
	t.Helper()
	mock := mocks.NewNotificationsMock()
	return NewService(mock), mock
}

func TestUnit_JamfProNotifications_NewService(t *testing.T) {
	mock := mocks.NewNotificationsMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
	assert.NotNil(t, svc.client)
}

func TestUnit_JamfProNotifications_GetForUserAndSiteV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetNotificationsMock()

	result, resp, err := svc.GetForUserAndSiteV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, "SYSTEM_ALERT", result[0].Type)
	assert.Equal(t, "notification-1", result[0].ID)
	assert.NotNil(t, result[0].Params)
	assert.Equal(t, "Test notification", result[0].Params["message"])
	assert.Equal(t, "info", result[0].Params["severity"])
}

func TestUnit_JamfProNotifications_GetForUserAndSiteV1_EmptyList(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetNotificationsEmptyMock()

	result, resp, err := svc.GetForUserAndSiteV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 0)
}

func TestUnit_JamfProNotifications_GetForUserAndSiteV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetNotificationsErrorMock()

	result, resp, err := svc.GetForUserAndSiteV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
	assert.Contains(t, err.Error(), "failed to get notifications")
}

func TestUnit_JamfProNotifications_GetForUserAndSiteV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// Do not register any mock

	result, resp, err := svc.GetForUserAndSiteV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfProNotifications_DeleteByTypeAndIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteNotificationMock()

	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "SYSTEM_ALERT", "notification-1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_JamfProNotifications_DeleteByTypeAndIDV1_EmptyType(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "", "notification-1")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "notification type is required")
}

func TestUnit_JamfProNotifications_DeleteByTypeAndIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "SYSTEM_ALERT", "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "notification id is required")
}

func TestUnit_JamfProNotifications_DeleteByTypeAndIDV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// Register only Get mock, not Delete
	mock := mocks.NewNotificationsMock()
	svc = NewService(mock)
	mock.RegisterGetNotificationsMock()

	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "SYSTEM_ALERT", "notification-1")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}
