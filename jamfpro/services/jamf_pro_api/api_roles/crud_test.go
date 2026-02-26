package api_roles

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.APIRolesMock) {
	t.Helper()
	mock := mocks.NewAPIRolesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_ApiRoles_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "Administrator", result.Results[0].DisplayName)
}

func TestUnit_ApiRoles_GetByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Administrator", result.DisplayName)
}

func TestUnit_ApiRoles_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ApiRoles_CreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestAPIRole{DisplayName: "Custom Role", Privileges: []string{"Read Computers"}}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "3", result.ID)
	require.Equal(t, "Custom Role", result.DisplayName)
}

func TestUnit_ApiRoles_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_ApiRoles_DeleteByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnit_APIRoles_UpdateByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestAPIRole{DisplayName: "Updated Role", Privileges: []string{"Read Computers", "Update Computers"}}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Administrator", result.DisplayName)
}

func TestUnit_APIRoles_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestAPIRole{DisplayName: "Updated Role", Privileges: []string{"Read Computers"}}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_APIRoles_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_APIRoles_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "API role ID is required")
}

func TestUnit_ApiRoles_ListV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAPIRolesMock())
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_ApiRoles_GetByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAPIRolesMock())
	_, _, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_ApiRoles_CreateV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAPIRolesMock())
	_, _, err := svc.CreateV1(context.Background(), &RequestAPIRole{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_ApiRoles_UpdateByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAPIRolesMock())
	_, _, err := svc.UpdateByIDV1(context.Background(), "1", &RequestAPIRole{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_ApiRoles_DeleteByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAPIRolesMock())
	_, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
}
