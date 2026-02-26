package venafi

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/venafi/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Venafi_Create_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceVenafi{
		Name:           "Venafi Certificate Authority",
		ProxyAddress:   "localhost:9443",
		ClientID:       "jamf-pro",
		RefreshToken:   "qdkP4SrCFKd7tefAVM6N",
	}

	result, resp, err := svc.Create(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v1/pki/venafi/1", result.Href)
}

func TestUnit_Venafi_Create_NilRequest(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.Create(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Venafi_GetByID_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByID(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Venafi Certificate Authority", result.Name)
	assert.Equal(t, "localhost:9443", result.ProxyAddress)
	assert.True(t, result.RevocationEnabled)
	assert.True(t, result.RefreshTokenConfigured)
}

func TestUnit_Venafi_GetByID_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByID(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_UpdateByID_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceVenafi{
		Name: "Updated Venafi CA",
	}

	result, resp, err := svc.UpdateByID(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Venafi Certificate Authority", result.Name)
}

func TestUnit_Venafi_UpdateByID_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceVenafi{Name: "Updated"}

	result, resp, err := svc.UpdateByID(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_UpdateByID_NilRequest(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateByID(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Venafi_DeleteByID_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByID(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_Venafi_DeleteByID_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByID(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_GetConnectionStatus_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetConnectionStatusMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetConnectionStatus(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Successfully connected", result.Status)
}

func TestUnit_Venafi_GetConnectionStatus_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetConnectionStatus(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_GetDependentProfiles_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetDependentProfilesMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDependentProfiles(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "Test Configuration Profile", result.Results[0].Name)
	assert.Equal(t, "OSXConfigurationProfile.html?id=1", result.Results[0].URLPath)
}

func TestUnit_Venafi_GetDependentProfiles_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDependentProfiles(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_GetHistory_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetHistoryMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistory(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Test note", result.Results[0].Note)
}

func TestUnit_Venafi_GetHistory_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetHistory(ctx, "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_AddHistoryNote_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterAddHistoryNoteMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{
		Note: "Test history note",
	}

	result, resp, err := svc.AddHistoryNote(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, "/api/v1/pki/venafi/1/history/2", result.Href)
}

func TestUnit_Venafi_AddHistoryNote_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &HistoryNoteRequest{Note: "Test note"}

	result, resp, err := svc.AddHistoryNote(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_AddHistoryNote_NilRequest(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.AddHistoryNote(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Venafi_GetJamfPublicKey_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetJamfPublicKeyMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, data, err := svc.GetJamfPublicKey(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, data)
}

func TestUnit_Venafi_GetJamfPublicKey_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, data, err := svc.GetJamfPublicKey(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Venafi_GetProxyTrustStore_Success(t *testing.T) {
	mock := mocks.NewVenafiMock()
	mock.RegisterGetProxyTrustStoreMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, data, err := svc.GetProxyTrustStore(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, data)
}

func TestUnit_Venafi_GetProxyTrustStore_EmptyID(t *testing.T) {
	mock := mocks.NewVenafiMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, data, err := svc.GetProxyTrustStore(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, data)
	assert.Contains(t, err.Error(), "id is required")
}
