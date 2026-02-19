package app_installers

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_installers/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AppInstallersMock) {
	t.Helper()
	mock := mocks.NewAppInstallersMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitListTitlesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListTitlesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
}

func TestUnitGetTitleByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetTitleByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Example App", result.TitleName)
}

func TestUnitListDeploymentsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListDeploymentsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
}

func TestUnitGetDeploymentByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetDeploymentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
}

func TestUnitCreateDeploymentV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDeployment{Name: "New", AppTitleId: "1"}
	result, resp, err := svc.CreateDeploymentV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "2", result.ID)
}

func TestUnitDeleteDeploymentByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteDeploymentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}
