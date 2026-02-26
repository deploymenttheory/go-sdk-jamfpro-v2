package jamf_protect

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_protect/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfProtectMock) {
	t.Helper()
	mock := mocks.NewJamfProtectMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_JamfProtect_GetSettingsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetSettingsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.Equal(t, "https://protect.example.com", result.ProtectURL)
}

func TestUnit_JamfProtect_UpdateSettingsV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	request := &RequestJamfProtectSettings{AutoInstall: true}
	result, resp, err := svc.UpdateSettingsV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUnit_JamfProtect_UpdateSettingsV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateSettingsV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "cannot be nil")
}

func TestUnit_JamfProtect_RegisterV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	request := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.RegisterV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
}

func TestUnit_JamfProtect_RegisterV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.RegisterV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "cannot be nil")
}

func TestUnit_JamfProtect_SyncPlansV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.SyncPlansV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnit_JamfProtect_ListDeploymentTasksV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListDeploymentTasksV1(context.Background(), "deploy-123", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
}

func TestUnit_JamfProtect_ListDeploymentTasksV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListDeploymentTasksV1(context.Background(), "", nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "deployment ID is required")
}

func TestUnit_JamfProtect_RetryDeploymentTasksV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.RetryDeploymentTasksV1(context.Background(), "deploy-123")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnit_JamfProtect_RetryDeploymentTasksV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.RetryDeploymentTasksV1(context.Background(), "")
	require.Error(t, err)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "deployment ID is required")
}

func TestUnit_JamfProtect_ListHistoryV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
}

func TestUnit_JamfProtect_CreateHistoryNoteV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	request := &RequestJamfProtectHistoryNote{
		Note:    "Test note",
		Details: "Test details",
	}
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Contains(t, []int{200, 201}, resp.StatusCode)
	require.Equal(t, 3, result.ID)
}

func TestUnit_JamfProtect_CreateHistoryNoteV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "cannot be nil")
}

func TestUnit_JamfProtect_ListPlansV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.ListPlansV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
}

func TestUnitDeleteIntegrationV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteIntegrationV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}

func TestUnitCreateIntegrationV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	registration := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.CreateIntegrationV1(context.Background(), registration, true)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
}

func TestUnit_JamfProtect_CreateIntegrationV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateIntegrationV1(context.Background(), nil, true)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "cannot be nil")
}

func TestUnit_JamfProtect_GetSettingsV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	result, resp, err := svc.GetSettingsV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_UpdateSettingsV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	request := &RequestJamfProtectSettings{AutoInstall: true}
	result, resp, err := svc.UpdateSettingsV1(context.Background(), request)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_RegisterV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	request := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.RegisterV1(context.Background(), request)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_SyncPlansV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	resp, err := svc.SyncPlansV1(context.Background())
	require.Error(t, err)
	_ = resp
}

func TestUnit_JamfProtect_CreateIntegrationV1_RegisterError(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	// No register mock → RegisterV1 fails
	svc := NewService(mock)

	registration := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.CreateIntegrationV1(context.Background(), registration, true)
	require.Error(t, err)
	require.Nil(t, result)
	require.Contains(t, err.Error(), "failed to register during integration creation")
	_ = resp
}

func TestUnit_JamfProtect_CreateIntegrationV1_UpdateError(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterRegisterMock() // Register succeeds
	// No UpdateSettings mock → UpdateSettingsV1 fails
	svc := NewService(mock)

	registration := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.CreateIntegrationV1(context.Background(), registration, true)
	require.Error(t, err)
	require.Nil(t, result)
	require.Contains(t, err.Error(), "failed to update settings during integration creation")
	_ = resp
}

func TestUnit_JamfProtect_CreateIntegrationV1_SyncError(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterRegisterMock()
	mock.RegisterUpdateSettingsMock()
	// No SyncPlans mock → SyncPlansV1 fails
	svc := NewService(mock)

	registration := &RequestJamfProtectRegistration{
		ProtectURL: "https://protect.example.com",
		ClientID:   "client-123",
		Password:   "secret",
	}
	result, resp, err := svc.CreateIntegrationV1(context.Background(), registration, true)
	require.Error(t, err)
	require.NotNil(t, result)
	require.Contains(t, err.Error(), "failed to sync plans during integration creation")
	_ = resp
}

func TestUnit_JamfProtect_ListDeploymentTasksV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	result, resp, err := svc.ListDeploymentTasksV1(context.Background(), "deploy-123", nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListDeploymentTasksV1_BadJSON(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListDeploymentTasksBadJSONMock()
	svc := NewService(mock)

	result, resp, err := svc.ListDeploymentTasksV1(context.Background(), "deploy-123", nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListDeploymentTasksV1_BadResults(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListDeploymentTasksBadResultsMock()
	svc := NewService(mock)

	result, resp, err := svc.ListDeploymentTasksV1(context.Background(), "deploy-123", nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_RetryDeploymentTasksV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	resp, err := svc.RetryDeploymentTasksV1(context.Background(), "deploy-123")
	require.Error(t, err)
	_ = resp
}

func TestUnit_JamfProtect_ListHistoryV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	result, resp, err := svc.ListHistoryV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListHistoryV1_BadJSON(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListHistoryBadJSONMock()
	svc := NewService(mock)

	result, resp, err := svc.ListHistoryV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListHistoryV1_BadResults(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListHistoryBadResultsMock()
	svc := NewService(mock)

	result, resp, err := svc.ListHistoryV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_CreateHistoryNoteV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	request := &RequestJamfProtectHistoryNote{Note: "Test"}
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), request)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListPlansV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	result, resp, err := svc.ListPlansV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListPlansV1_BadJSON(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListPlansBadJSONMock()
	svc := NewService(mock)

	result, resp, err := svc.ListPlansV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_ListPlansV1_BadResults(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	mock.RegisterListPlansBadResultsMock()
	svc := NewService(mock)

	result, resp, err := svc.ListPlansV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	_ = resp
}

func TestUnit_JamfProtect_DeleteIntegrationV1_Error(t *testing.T) {
	mock := mocks.NewJamfProtectMock()
	svc := NewService(mock)

	resp, err := svc.DeleteIntegrationV1(context.Background())
	require.Error(t, err)
	_ = resp
}
