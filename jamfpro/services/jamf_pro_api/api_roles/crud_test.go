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

func TestUnitListAPIRolesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListAPIRolesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "Administrator", result.Results[0].DisplayName)
}

func TestUnitGetAPIRoleByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetAPIRoleByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Administrator", result.DisplayName)
}

func TestUnitGetAPIRoleByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetAPIRoleByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnitCreateAPIRoleV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestAPIRole{DisplayName: "Custom Role", Privileges: []string{"Read Computers"}}
	result, resp, err := svc.CreateAPIRoleV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "3", result.ID)
	require.Equal(t, "Custom Role", result.DisplayName)
}

func TestUnitCreateAPIRoleV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateAPIRoleV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnitDeleteAPIRoleByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteAPIRoleByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}
