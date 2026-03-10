package time_zones

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/time_zones/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*TimeZones, *mocks.TimeZonesMock) {
	t.Helper()
	mock := mocks.NewTimeZonesMock()
	mock.RegisterMocks()
	return NewTimeZones(mock), mock
}

func TestUnit_TimeZones_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	require.Equal(t, "America/Los_Angeles", result[0].ZoneId)
	require.Equal(t, "Pacific Time (US & Canada)", result[0].DisplayName)
}

func TestUnit_TimeZones_ListV1_EmptyList(t *testing.T) {
	mock := mocks.NewTimeZonesMock()
	mock.RegisterListV1EmptyMock()
	svc := NewTimeZones(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 0)
}

func TestUnit_TimeZones_ListV1_Error(t *testing.T) {
	mock := mocks.NewTimeZonesMock()
	mock.RegisterListV1ErrorMock()
	svc := NewTimeZones(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 500, resp.StatusCode())
	require.Contains(t, err.Error(), "mock client error")
}

func TestUnit_TimeZones_ListV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewTimeZonesMock()
	mock.RegisterListV1NoResponseErrorMock()
	svc := NewTimeZones(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_TimeZones_ListV1_InvalidJSON(t *testing.T) {
	mock := mocks.NewTimeZonesMock()
	mock.RegisterListV1InvalidJSONMock()
	svc := NewTimeZones(mock)

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
}

func TestUnit_TimeZones_NewService(t *testing.T) {
	mock := mocks.NewTimeZonesMock()
	svc := NewTimeZones(mock)
	require.NotNil(t, svc)
	require.NotNil(t, svc.client)
}
