package user

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/user/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.UserMock) {
	t.Helper()
	mock := mocks.NewUserMock()
	return NewService(mock), mock
}

func TestUnit_User_NewService(t *testing.T) {
	mock := mocks.NewUserMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_User_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "admin", result.Username)
	assert.Equal(t, "Administrator", result.RealName)
	assert.Equal(t, "admin@example.com", result.Email)
	assert.True(t, result.IsMultiSiteAdmin)
	assert.Equal(t, "FullAccess", result.AccessLevel)
	assert.Equal(t, "ADMINISTRATOR", result.PrivilegeSet)
	assert.Equal(t, 1, result.CurrentSiteID)
	assert.Len(t, result.GroupIDs, 2)
	assert.Equal(t, []int{1, 2}, result.GroupIDs)
}

func TestUnit_User_Get_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Get(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_User_Get_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetErrorMock()

	result, resp, err := svc.Get(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_User_Get_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInvalidJSONMock()

	result, resp, err := svc.Get(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, err.Error(), "unmarshal")
}

func TestUnit_User_ChangePassword_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterChangePasswordMock()

	req := &RequestChangePassword{
		CurrentPassword: "oldpass123",
		NewPassword:     "newpass456",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_User_ChangePassword_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.ChangePassword(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_User_ChangePassword_EmptyCurrentPassword(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestChangePassword{
		CurrentPassword: "",
		NewPassword:     "newpass456",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "currentPassword is required")
}

func TestUnit_User_ChangePassword_WhitespaceCurrentPassword(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestChangePassword{
		CurrentPassword: "   ",
		NewPassword:     "newpass456",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "currentPassword is required")
}

func TestUnit_User_ChangePassword_EmptyNewPassword(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestChangePassword{
		CurrentPassword: "oldpass123",
		NewPassword:     "",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "newPassword is required")
}

func TestUnit_User_ChangePassword_WhitespaceNewPassword(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestChangePassword{
		CurrentPassword: "oldpass123",
		NewPassword:     "\t\n",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "newPassword is required")
}

func TestUnit_User_ChangePassword_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterChangePasswordErrorMock()

	req := &RequestChangePassword{
		CurrentPassword: "oldpass123",
		NewPassword:     "newpass456",
	}
	resp, err := svc.ChangePassword(context.Background(), req)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 400, resp.StatusCode())
}

func TestUnit_User_UpdateSession_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSessionMock()

	req := &RequestUpdateSession{
		CurrentSiteID: 2,
	}
	resp, err := svc.UpdateSession(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_User_UpdateSession_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateSession(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_User_UpdateSession_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSessionErrorMock()

	req := &RequestUpdateSession{
		CurrentSiteID: 1,
	}
	resp, err := svc.UpdateSession(context.Background(), req)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}
