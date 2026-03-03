package gsx_connection

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/gsx_connection/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.GSXConnectionMock) {
	t.Helper()
	mock := mocks.NewGSXConnectionMock()
	return NewService(mock), mock
}

func TestUnit_GsxConnection_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetGSXConnectionMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.Enabled)
	assert.Equal(t, "test@example.com", result.Username)
	assert.Equal(t, "12345", result.ServiceAccountNo)
	assert.Equal(t, "67890", result.ShipToNo)
	assert.Equal(t, "certificate.p12", result.GsxKeystore.Name)
	assert.Equal(t, int64(1691954900000), result.GsxKeystore.ExpirationEpoch)
}

func TestUnit_GsxConnection_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateGSXConnectionMock()

	request := &ResourceGSXConnection{
		Enabled:          false,
		Username:         "updated@example.com",
		ServiceAccountNo: "54321",
		ShipToNo:         "09876",
		GsxKeystore: GsxKeystore{
			Name:            "certificate.p12",
			ExpirationEpoch: 1691954900000,
			ErrorMessage:    "",
		},
	}

	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.Enabled)
	assert.Equal(t, "updated@example.com", result.Username)
	assert.Equal(t, "54321", result.ServiceAccountNo)
}

func TestUnit_GsxConnection_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_GsxConnection_GetHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "GSX connection enabled", result.Results[0].Note)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Updated service account", result.Results[1].Note)
}

func TestUnit_GsxConnection_GetHistory_WithRSQLQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	rsqlQuery := map[string]string{"sort": "date:desc"}
	result, resp, err := svc.GetHistoryV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_GsxConnection_Get_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_GsxConnection_Get_NotFoundError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "NOT-FOUND")
	assert.Contains(t, err.Error(), "GSX connection not found")
}

func TestUnit_GsxConnection_Update_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	request := &ResourceGSXConnection{Enabled: true, Username: "test@example.com"}

	result, resp, err := svc.UpdateV1(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_GsxConnection_GetHistory_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_GsxConnection_GetHistory_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryInvalidMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, err.Error(), "failed to get GSX connection history")
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_GsxConnection_NewService(t *testing.T) {
	mock := mocks.NewGSXConnectionMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}
