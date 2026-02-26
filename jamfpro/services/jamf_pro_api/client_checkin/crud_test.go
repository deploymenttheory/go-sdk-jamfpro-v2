package client_checkin

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ClientCheckinMock) {
	t.Helper()
	mock := mocks.NewClientCheckinMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_ClientCheckin_GetV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 15, result.CheckInFrequency)
	require.True(t, result.CreateHooks)
}

func TestUnit_ClientCheckin_UpdateV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	settings := &ResourceClientCheckinSettings{CheckInFrequency: 30, CreateHooks: true}
	result, resp, err := svc.UpdateV3(context.Background(), settings)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUnit_ClientCheckin_UpdateV3_NilSettings(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV3(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ClientCheckin_GetHistoryV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryV3(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "admin", result.Results[0].Username)
	require.Equal(t, "Initial config", result.Results[0].Note)
}

func TestUnit_ClientCheckin_AddHistoryNoteV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestClientCheckinHistoryNote{Note: "Test note"}
	result, resp, err := svc.AddHistoryNoteV3(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.NotEmpty(t, result.HREF)
}

func TestUnit_ClientCheckin_AddHistoryNoteV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.AddHistoryNoteV3(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ClientCheckin_GetV3_Error(t *testing.T) {
	svc := NewService(mocks.NewClientCheckinMock())
	result, resp, err := svc.GetV3(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ClientCheckin_GetHistoryV3_Error(t *testing.T) {
	svc := NewService(mocks.NewClientCheckinMock())
	result, resp, err := svc.GetHistoryV3(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ClientCheckin_UpdateV3_Error(t *testing.T) {
	svc := NewService(mocks.NewClientCheckinMock())
	settings := &ResourceClientCheckinSettings{CheckInFrequency: 30, CreateHooks: true}
	result, resp, err := svc.UpdateV3(context.Background(), settings)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ClientCheckin_AddHistoryNoteV3_Error(t *testing.T) {
	svc := NewService(mocks.NewClientCheckinMock())
	req := &RequestClientCheckinHistoryNote{Note: "Test note"}
	result, resp, err := svc.AddHistoryNoteV3(context.Background(), req)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}
