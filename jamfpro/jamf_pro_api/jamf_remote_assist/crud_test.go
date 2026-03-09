package jamf_remote_assist

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_remote_assist/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*JamfRemoteAssist, *mocks.JamfRemoteAssistMock) {
	t.Helper()
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	return NewJamfRemoteAssist(mock), mock
}

func TestUnit_JamfRemoteAssist_ListSessionsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListSessionsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 1)
	require.Equal(t, "session-abc", result[0].SessionID)
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV1(context.Background(), "session-abc")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
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
	require.Nil(t, resp) // no mock registered for nonexistent path
}

func TestUnit_JamfRemoteAssist_ListSessionsV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListSessionsV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSessionByIDV2(context.Background(), "session-abc")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
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

func TestUnit_JamfRemoteAssist_ExportSessionsV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), &ExportSessionsRequest{}, "text/csv")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	_ = result
}

func TestUnit_JamfRemoteAssist_ExportSessionsV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), nil, "text/csv")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	_ = result
}

func TestUnit_JamfRemoteAssist_ExportSessionsV2_DefaultAcceptType(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ExportSessionsV2(context.Background(), nil, "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	_ = result
}

func TestUnit_JamfRemoteAssist_ListSessionsV1_Error(t *testing.T) {
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	mock.RegisterListSessionsV1ErrorMock()
	svc := NewJamfRemoteAssist(mock)

	result, resp, err := svc.ListSessionsV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "failed to list jamf remote assist sessions (v1)")
}

func TestUnit_JamfRemoteAssist_ListSessionsV2_Error(t *testing.T) {
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	mock.RegisterListSessionsV2ErrorMock()
	svc := NewJamfRemoteAssist(mock)

	result, resp, err := svc.ListSessionsV2(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "failed to list jamf remote assist sessions (v2)")
}

func TestUnit_JamfRemoteAssist_ListSessionsV2_MergePageError(t *testing.T) {
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	mock.RegisterListSessionsV2InvalidMock()
	svc := NewJamfRemoteAssist(mock)

	result, resp, err := svc.ListSessionsV2(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_JamfRemoteAssist_GetSessionByIDV2_NotFound(t *testing.T) {
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	mock.RegisterGetSessionByIDV2ErrorMock()
	svc := NewJamfRemoteAssist(mock)

	result, resp, err := svc.GetSessionByIDV2(context.Background(), "nonexistent")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "failed to get jamf remote assist session by ID (v2)")
}

func TestUnit_JamfRemoteAssist_ExportSessionsV2_Error(t *testing.T) {
	mock := mocks.NewJamfRemoteAssistMock()
	mock.RegisterMocks()
	mock.RegisterExportSessionsV2ErrorMock()
	svc := NewJamfRemoteAssist(mock)

	result, resp, err := svc.ExportSessionsV2(context.Background(), &ExportSessionsRequest{}, "text/csv")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Contains(t, err.Error(), "failed to export jamf remote assist sessions")
}
