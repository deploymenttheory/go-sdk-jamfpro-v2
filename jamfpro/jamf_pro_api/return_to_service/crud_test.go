package return_to_service

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/return_to_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*ReturnToService, *mocks.ReturnToServiceMock) {
	t.Helper()
	mock := mocks.NewReturnToServiceMock()
	return NewReturnToService(mock), mock
}

func TestUnit_ReturnToService_ListV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListNoResponseErrorMock()

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_ReturnToService_ListV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()
	result, resp, err := svc.ListV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Default", result.Results[0].DisplayName)
}

func TestUnit_ReturnToService_GetByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ReturnToService_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_ReturnToService_DeleteByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteByIDV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_ReturnToService_GetByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Default", result.DisplayName)
}

func TestUnit_ReturnToService_CreateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &ResourceReturnToServiceConfiguration{
		DisplayName:   "New Config",
		WifiProfileID: "wifi-1",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_ReturnToService_UpdateByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &ResourceReturnToServiceConfiguration{
		DisplayName:   "Updated Config",
		WifiProfileID: "wifi-1",
	}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Updated Config", result.DisplayName)
}

func TestUnit_ReturnToService_UpdateByIDV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &ResourceReturnToServiceConfiguration{DisplayName: "Updated Config"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ReturnToService_UpdateByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ReturnToService_DeleteByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ReturnToService_ListV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListErrorMock()

	result, resp, err := svc.ListV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_ReturnToService_GetByIDV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_ReturnToService_CreateV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateErrorMock()

	req := &ResourceReturnToServiceConfiguration{DisplayName: "New Config"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_ReturnToService_UpdateByIDV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateErrorMock()

	req := &ResourceReturnToServiceConfiguration{DisplayName: "Updated Config"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_ReturnToService_DeleteByIDV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteErrorMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, resp)
}
