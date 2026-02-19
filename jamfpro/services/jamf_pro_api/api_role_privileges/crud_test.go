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

func TestUnitListPrivilegesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListPrivilegesV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Privileges, 3)
}

func TestUnitSearchPrivilegesByNameV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.SearchPrivilegesByNameV1(context.Background(), "Read", 10)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Privileges, 3)
}
