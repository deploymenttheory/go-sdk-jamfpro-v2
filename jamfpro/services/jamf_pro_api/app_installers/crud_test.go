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

func TestUnit_AppInstallers_ListTitlesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListTitlesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
}

func TestUnit_AppInstallers_GetTitleByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetTitleByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Example App", result.TitleName)
}

func TestUnit_AppInstallers_ListDeploymentsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListDeploymentsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, 1, result.TotalCount)
}

func TestUnit_AppInstallers_GetDeploymentByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetDeploymentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "1", result.ID)
}

func TestUnit_AppInstallers_CreateDeploymentV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDeployment{Name: "New", AppTitleId: "1"}
	result, resp, err := svc.CreateDeploymentV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 201, resp.StatusCode())
	require.Equal(t, "2", result.ID)
}

func TestUnit_AppInstallers_DeleteDeploymentByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteDeploymentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode())
}

func TestUnit_AppInstallers_UpdateDeploymentByIDV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	enabled := true
	req := &RequestDeployment{Name: "Updated Deployment", AppTitleId: "1", Enabled: &enabled}
	result, resp, err := svc.UpdateDeploymentByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "1", result.ID)
	require.Equal(t, "Example Deployment", result.Name)
}

func TestUnit_AppInstallers_UpdateDeploymentByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestDeployment{Name: "Updated Deployment", AppTitleId: "1"}
	result, resp, err := svc.UpdateDeploymentByIDV1(context.Background(), "", req)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "id is required")
}

func TestUnit_AppInstallers_UpdateDeploymentByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateDeploymentByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_AppInstallers_GetTitleByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetTitleByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "title ID is required")
}

func TestUnit_AppInstallers_GetDeploymentByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetDeploymentByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "deployment ID is required")
}

func TestUnit_AppInstallers_CreateDeploymentV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateDeploymentV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_AppInstallers_DeleteDeploymentByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteDeploymentByIDV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "deployment ID is required")
}

func TestUnit_AppInstallers_ListTitlesV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	result, resp, err := svc.ListTitlesV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_GetTitleByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	result, resp, err := svc.GetTitleByIDV1(context.Background(), "1")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_ListDeploymentsV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	result, resp, err := svc.ListDeploymentsV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_GetDeploymentByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	result, resp, err := svc.GetDeploymentByIDV1(context.Background(), "1")
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_CreateDeploymentV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	req := &RequestDeployment{Name: "New", AppTitleId: "1"}
	result, resp, err := svc.CreateDeploymentV1(context.Background(), req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_UpdateDeploymentByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	enabled := true
	req := &RequestDeployment{Name: "Updated", AppTitleId: "1", Enabled: &enabled}
	result, resp, err := svc.UpdateDeploymentByIDV1(context.Background(), "1", req)
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_AppInstallers_DeleteDeploymentByIDV1_Error(t *testing.T) {
	svc := NewService(mocks.NewAppInstallersMock())
	resp, err := svc.DeleteDeploymentByIDV1(context.Background(), "1")
	require.Error(t, err)
	require.NotNil(t, resp)
}
