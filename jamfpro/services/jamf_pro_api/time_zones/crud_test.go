package time_zones

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/time_zones/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.TimeZonesMock) {
	t.Helper()
	mock := mocks.NewTimeZonesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitListTimeZonesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListTimeZonesV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 2)
	require.Equal(t, "America/Los_Angeles", result[0].ZoneId)
	require.Equal(t, "Pacific Time (US & Canada)", result[0].DisplayName)
}
