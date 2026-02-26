package jamf_connect

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_connect/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfConnectMock) {
	t.Helper()
	mock := mocks.NewJamfConnectMock()
	return NewService(mock), mock
}

func TestUnit_JamfConnect_GetSettings_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSettingsMock()

	result, resp, err := svc.GetSettingsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Jamf Connect Settings", result.DisplayName)
	assert.True(t, result.Enabled)
	assert.Equal(t, "2.0.0", result.Version)
}

func TestUnit_JamfConnect_ListConfigProfiles_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.ListConfigProfilesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", result.Results[0].UUID)
	assert.Equal(t, "Test Profile 1", result.Results[0].ProfileName)
	assert.Equal(t, 1, result.Results[0].ProfileID)
}

func TestUnit_JamfConnect_GetConfigProfileByUUID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", result.UUID)
	assert.Equal(t, "Test Profile 1", result.ProfileName)
}

func TestUnit_JamfConnect_GetConfigProfileByUUID_EmptyUUID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByUUIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "uuid is required")
}

func TestUnit_JamfConnect_GetConfigProfileByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByIDV1(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ProfileID)
	assert.Equal(t, "Test Profile 1", result.ProfileName)
}

func TestUnit_JamfConnect_GetConfigProfileByID_InvalidID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByIDV1(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "profile ID must be greater than 0")
}

func TestUnit_JamfConnect_GetConfigProfileByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByNameV1(context.Background(), "Test Profile 2")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Test Profile 2", result.ProfileName)
	assert.Equal(t, 2, result.ProfileID)
}

func TestUnit_JamfConnect_GetConfigProfileByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByNameV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

func TestUnit_JamfConnect_UpdateConfigProfileByUUID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateConfigProfileMock()

	request := &ResourceJamfConnectConfigProfileUpdate{
		JamfConnectVersion: "2.1.0",
		AutoDeploymentType: "INSTALL_AUTOMATICALLY",
	}

	result, resp, err := svc.UpdateConfigProfileByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", result.UUID)
	assert.Equal(t, "2.1.0", result.Version)
}

func TestUnit_JamfConnect_UpdateConfigProfileByUUID_EmptyUUID(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceJamfConnectConfigProfileUpdate{
		JamfConnectVersion: "2.1.0",
	}

	result, resp, err := svc.UpdateConfigProfileByUUIDV1(context.Background(), "", request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "uuid is required")
}

func TestUnit_JamfConnect_UpdateConfigProfileByUUID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateConfigProfileByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_JamfConnect_RetryDeploymentTasksByUUID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRetryDeploymentTasksMock()

	computerIDs := []string{"1", "2", "3"}
	resp, err := svc.RetryDeploymentTasksByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", computerIDs)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_JamfConnect_RetryDeploymentTasksByUUID_EmptyUUID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.RetryDeploymentTasksByUUIDV1(context.Background(), "", []string{"1"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "uuid is required")
}

func TestUnit_JamfConnect_RetryDeploymentTasksByUUID_EmptyComputerIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.RetryDeploymentTasksByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", []string{})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "at least one computer ID is required")
}

func TestUnit_JamfConnect_GetDeploymentTasksByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeploymentTasksMock()

	result, resp, err := svc.GetDeploymentTasksByIDV1(context.Background(), "dep-123", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "COMPLETED", result.Results[0].Status)
	assert.Equal(t, "2.0.0", result.Results[0].Version)
}

func TestUnit_JamfConnect_GetDeploymentTasksByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDeploymentTasksByIDV1(context.Background(), "", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "deployment ID is required")
}

func TestUnit_JamfConnect_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_JamfConnect_GetSettingsV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSettingsV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_ListConfigProfilesV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListConfigProfilesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_ListConfigProfilesV1_BadJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesBadJSONMock()

	result, resp, err := svc.ListConfigProfilesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_ListConfigProfilesV1_BadResults(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesBadResultsMock()

	result, resp, err := svc.ListConfigProfilesV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetConfigProfileByUUIDV1_ListError(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByUUIDV1(context.Background(), "unknown-uuid")
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetConfigProfileByUUIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByUUIDV1(context.Background(), "non-existent-uuid")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "no jamf connect config profile found with UUID")
}

func TestUnit_JamfConnect_GetConfigProfileByIDV1_ListError(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByIDV1(context.Background(), 999)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetConfigProfileByIDV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByIDV1(context.Background(), 9999)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "no jamf connect config profile found with ID")
}

func TestUnit_JamfConnect_GetConfigProfileByNameV1_ListError(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConfigProfileByNameV1(context.Background(), "Unknown")
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetConfigProfileByNameV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListConfigProfilesMock()

	result, resp, err := svc.GetConfigProfileByNameV1(context.Background(), "Non Existent Profile")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "no jamf connect config profile found with name")
}

func TestUnit_JamfConnect_UpdateConfigProfileByUUIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceJamfConnectConfigProfileUpdate{JamfConnectVersion: "2.0.0"}
	result, resp, err := svc.UpdateConfigProfileByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetDeploymentTasksByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDeploymentTasksByIDV1(context.Background(), "dep-123", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetDeploymentTasksByIDV1_BadJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeploymentTasksBadJSONMock()

	result, resp, err := svc.GetDeploymentTasksByIDV1(context.Background(), "dep-123", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetDeploymentTasksByIDV1_BadResults(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDeploymentTasksBadResultsMock()

	result, resp, err := svc.GetDeploymentTasksByIDV1(context.Background(), "dep-123", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetHistoryV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetHistoryV1_BadJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryBadJSONMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_GetHistoryV1_BadResults(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryBadResultsMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_JamfConnect_RetryDeploymentTasksByUUIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.RetryDeploymentTasksByUUIDV1(context.Background(), "123e4567-e89b-12d3-a456-426614174000", []string{"1"})
	require.Error(t, err)
	_ = resp
}
