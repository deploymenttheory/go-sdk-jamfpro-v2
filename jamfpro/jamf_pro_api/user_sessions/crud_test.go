package user_sessions

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/user_sessions/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*UserSessions, *mocks.UserSessionsMock) {
	t.Helper()
	mock := mocks.NewUserSessionsMock()
	return NewUserSessions(mock), mock
}

// Test GetActiveV1 success
func TestUnit_UserSessions_GetActiveV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActiveSessionsMock()

	result, resp, err := svc.GetActiveV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "admin", result[0].Username)
}

// Test GetCountV1 success
func TestUnit_UserSessions_GetCountV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCountMock()

	result, resp, err := svc.GetCountV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Count)
}

// Error tests (no mock registered)
func TestUnit_UserSessions_GetActiveV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetActiveV1(context.Background())
	require.Error(t, err)
}

func TestUnit_UserSessions_GetCountV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetCountV1(context.Background())
	require.Error(t, err)
}

func TestUnit_UserSessions_V2(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.Register("GET", "/api/v2/user-sessions/active", 200, "validate_active_v2.json")
	mock.Register("GET", "/api/v2/user-sessions/count", 200, "validate_count_v2.json")
	active, resp, err := svc.GetActiveV2(context.Background())
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Empty(t, active)
	count, resp, err := svc.GetCountV2(context.Background())
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Zero(t, count.ActiveUserCount)
	// Nonzero synthetic count catches incorrect JSON tags hidden by an empty tenant.
	mock.RegisterRawBody("GET", "/api/v2/user-sessions/count", 200, []byte(`{"activeUserCount":3}`))
	count, _, err = svc.GetCountV2(context.Background())
	require.NoError(t, err)
	require.Equal(t, 3, count.ActiveUserCount)
	mock.RegisterError("GET", "/api/v2/user-sessions/active", 403, "", "Forbidden")
	active, resp, err = svc.GetActiveV2(context.Background())
	require.Error(t, err)
	require.Nil(t, active)
	require.Equal(t, 403, resp.StatusCode())
	mock.RegisterError("GET", "/api/v2/user-sessions/count", 500, "", "Server error")
	count, resp, err = svc.GetCountV2(context.Background())
	require.Error(t, err)
	require.Nil(t, count)
	require.Equal(t, 500, resp.StatusCode())
}
