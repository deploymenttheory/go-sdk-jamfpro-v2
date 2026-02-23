package jamf_pro_notifications

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_notifications/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetForUserAndSiteV1_Success(t *testing.T) {
	mock := mocks.NewNotificationsMock()
	mock.RegisterGetNotificationsMock()
	service := NewService(mock)

	result, resp, err := service.GetForUserAndSiteV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, "SYSTEM_ALERT", result[0].Type)
	assert.Equal(t, "notification-1", result[0].ID)
	assert.NotNil(t, result[0].Params)
}
