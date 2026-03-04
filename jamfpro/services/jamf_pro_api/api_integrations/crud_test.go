package api_integrations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_integrations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ApiIntegrationsMock) {
	t.Helper()
	mock := mocks.NewApiIntegrationsMock()
	return NewService(mock), mock
}

func TestUnit_ApiIntegrations_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Test Integration", result.Results[0].DisplayName)
	assert.Equal(t, 1, result.Results[0].ID)
}

func TestUnit_ApiIntegrations_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnit_ApiIntegrations_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_GetByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV1(context.Background(), "Test Integration")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnit_ApiIntegrations_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_CreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &RequestApiIntegration{DisplayName: "New Integration", Enabled: true}
	result, resp, err := svc.CreateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnit_ApiIntegrations_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", &RequestApiIntegration{})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_UpdateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock("1")

	request := &RequestApiIntegration{DisplayName: "Updated", Enabled: true}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ApiIntegrations_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_DeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ApiIntegrations_UpdateByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterUpdateByIDMock("1")

	request := &RequestApiIntegration{DisplayName: "Updated", Enabled: true}
	result, resp, err := svc.UpdateByNameV1(context.Background(), "Test Integration", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ApiIntegrations_UpdateByNameV1_NilRequest(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.UpdateByNameV1(context.Background(), "Test Integration", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_DeleteByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByNameV1(context.Background(), "Test Integration")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ApiIntegrations_DeleteByNameV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	resp, err := svc.DeleteByNameV1(context.Background(), "Nonexistent")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not found")
	// resp may be non-nil (list response) when name is not in results
	_ = resp
}

func TestUnit_ApiIntegrations_RefreshClientCredentialsByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RefreshClientCredentialsByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ApiIntegrations_RefreshClientCredentialsByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRefreshClientCredentialsMock("1")

	result, resp, err := svc.RefreshClientCredentialsByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "new-client-id", result.ClientID)
	assert.Equal(t, "new-client-secret", result.ClientSecret)
}

func TestUnit_ApiIntegrations_ListV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.ListV1(context.Background(), nil)
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_GetByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_GetByNameV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByNameV1(context.Background(), "test")
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_CreateV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreateV1(context.Background(), &RequestApiIntegration{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_UpdateByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByIDV1(context.Background(), "1", &RequestApiIntegration{DisplayName: "test"})
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_UpdateByNameV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByNameV1(context.Background(), "test", &RequestApiIntegration{DisplayName: "updated"})
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_DeleteByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
}

func TestUnit_ApiIntegrations_RefreshClientCredentialsByIDV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.RefreshClientCredentialsByIDV1(context.Background(), "1")
	require.Error(t, err)
}
