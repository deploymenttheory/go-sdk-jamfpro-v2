package notifications

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/notifications/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Notifications, *mocks.NotificationsMock) {
	t.Helper()
	mock := mocks.NewNotificationsMock()
	mock.RegisterMocks()
	return NewNotifications(mock), mock
}

func TestUnit_Notifications_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 1)
	require.Equal(t, "EXAMPLE", result[0].Type)
	require.Equal(t, "1", result[0].ID)
}

func TestUnit_Notifications_ListV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListErrorMock()
	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Notifications_DeleteByTypeAndIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "APNS_CERT_REVOKED", "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Notifications_DeleteByTypeAndIDV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteErrorMock()
	resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "APNS_CERT_REVOKED", "1")
	require.Error(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Notifications_DeleteByTypeAndIDV1_ValidationErrors(t *testing.T) {
	svc, _ := setupMockService(t)

	t.Run("empty notificationType", func(t *testing.T) {
		resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "", "1")
		require.Error(t, err)
		require.Nil(t, resp)
		require.Contains(t, err.Error(), "notificationType is required")
	})

	t.Run("empty id", func(t *testing.T) {
		resp, err := svc.DeleteByTypeAndIDV1(context.Background(), "APNS_CERT_REVOKED", "")
		require.Error(t, err)
		require.Nil(t, resp)
		require.Contains(t, err.Error(), "id is required")
	})
}
