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

func TestUnitListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "Test Integration", result.Results[0].DisplayName)
	assert.Equal(t, 1, result.Results[0].ID)
}

func TestUnitGetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnitGetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitGetByNameV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.GetByNameV1(context.Background(), "Test Integration")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnitCreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitCreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	request := &ResourceApiIntegration{DisplayName: "New Integration", Enabled: true}
	result, resp, err := svc.CreateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Test Integration", result.DisplayName)
}

func TestUnitUpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", &ResourceApiIntegration{})
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock("1")

	request := &ResourceApiIntegration{DisplayName: "Updated", Enabled: true}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnitDeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitRefreshClientCredentialsByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RefreshClientCredentialsByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitRefreshClientCredentialsByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRefreshClientCredentialsMock("1")

	result, resp, err := svc.RefreshClientCredentialsByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "new-client-id", result.ClientID)
	assert.Equal(t, "new-client-secret", result.ClientSecret)
}
