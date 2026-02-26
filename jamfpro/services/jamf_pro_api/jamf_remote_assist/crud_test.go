package jamf_remote_assist

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_remote_assist/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfRemoteAssistMock) {
	t.Helper()
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitListSessionsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListSessionsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 1)
	require.Equal(t, "session-abc", result[0].SessionID)
}

func TestUnitGetSessionByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV1(context.Background(), "session-abc")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "session-abc", result.SessionID)
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "session ID is required")
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV1_NotFound(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV1(context.Background(), "nonexistent")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnitListSessionsV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListSessionsV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
}

func TestUnitGetSessionByIDV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV2(context.Background(), "session-abc")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "session-abc", result.SessionID)
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV2_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV2(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "session ID is required")
}

func TestUnitExportSessionsV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), &ExportSessionsRequest{}, "text/csv")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	_ = result
}

func TestUnit_JamfRemoteAssist_ExportSessionsV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), nil, "text/csv")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	_ = result
}

func TestUnit_JamfRemoteAssist_ExportSessionsV2_DefaultAcceptType(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), nil, "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	_ = result
}
