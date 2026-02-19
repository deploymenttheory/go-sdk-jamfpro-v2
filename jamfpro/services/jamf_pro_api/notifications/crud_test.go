package notifications

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/notifications/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.NotificationsMock) {
	t.Helper()
	mock := mocks.NewNotificationsMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 1)
	require.Equal(t, "EXAMPLE", result[0].Type)
	require.Equal(t, "1", result[0].ID)
}
