package api_role_privileges

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_role_privileges/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.APIRolePrivilegesMock) {
	t.Helper()
	mock := mocks.NewAPIRolePrivilegesMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_ApiRolePrivileges_ListV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Privileges, 3)
}

func TestUnit_ApiRolePrivileges_SearchPrivilegesByNameV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.SearchPrivilegesByNameV1(context.Background(), "Read", 10)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Privileges, 3)
}

func TestUnit_APIRolePrivileges_SearchPrivilegesByNameV1_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.SearchPrivilegesByNameV1(context.Background(), "", 10)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "name parameter is required")
}

func TestUnit_APIRolePrivileges_ListV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListErrorMock()
	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_APIRolePrivileges_SearchPrivilegesByNameV1_DefaultLimit(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterSearchDefaultLimitMock()
	result, resp, err := svc.SearchPrivilegesByNameV1(context.Background(), "Read", 0)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
}
