package digicert

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/digicert/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Digicert, *mocks.DigicertMock) {
	t.Helper()
	mock := mocks.NewDigicertMock()
	return NewDigicert(mock), mock
}

func TestUnit_Digicert_NewService(t *testing.T) {
	mock := mocks.NewDigicertMock()
	svc := NewDigicert(mock)
	require.NotNil(t, svc)
	assert.NotNil(t, svc.client)
}

func TestUnit_Digicert_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &ResourceDigicertTrustLifecycleManager{
		CAName:   "DigiCert CA",
		FQDN:     "digicert.example.com",
		ClientCert: &ResourceDigicertClientCert{
			Filename: "client.p12",
			Data:     []string{"base64data"},
			Password: "secret",
		},
	}

	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v1/pki/digicert/trust-lifecycle-manager/1")
}

func TestUnit_Digicert_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Digicert_Create_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateErrorMock()

	req := &ResourceDigicertTrustLifecycleManager{CAName: "CA", FQDN: "example.com"}
	result, resp, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock("1")

	result, resp, err := svc.GetByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "DigiCert CA", result.CAName)
	assert.Equal(t, "digicert.example.com", result.FQDN)
	assert.True(t, result.RevocationEnabled)
	require.NotNil(t, result.ClientCert)
	assert.Equal(t, "client.p12", result.ClientCert.Filename)
	assert.Equal(t, "ABC123", result.ClientCert.SerialNumber)
	assert.Equal(t, "CN=DigiCert Client", result.ClientCert.Subject)
}

func TestUnit_Digicert_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Digicert_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	result, resp, err := svc.GetByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Digicert_GetByID_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDErrorMock("1")

	result, resp, err := svc.GetByID(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock("1")

	req := &ResourceDigicertTrustLifecycleManager{
		CAName: "DigiCert CA Updated",
		FQDN:   "updated.example.com",
	}

	resp, err := svc.UpdateByID(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Digicert_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceDigicertTrustLifecycleManager{CAName: "CA", FQDN: "example.com"}
	resp, err := svc.UpdateByID(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Digicert_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.UpdateByID(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Digicert_UpdateByID_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDErrorMock("1")

	req := &ResourceDigicertTrustLifecycleManager{CAName: "CA", FQDN: "example.com"}
	resp, err := svc.UpdateByID(context.Background(), "1", req)
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock("1")

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Digicert_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Digicert_DeleteByID_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDErrorMock("1")

	resp, err := svc.DeleteByID(context.Background(), "1")
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_ValidateClientCertificate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterValidateClientCertificateMock()

	req := &ValidateClientCertificateRequest{
		Filename: "client.p12",
		Data:     []string{"base64data"},
		Password: strPtr("secret"),
	}

	resp, err := svc.ValidateClientCertificate(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Digicert_ValidateClientCertificate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.ValidateClientCertificate(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Digicert_ValidateClientCertificate_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterValidateClientCertificateErrorMock()

	req := &ValidateClientCertificateRequest{Filename: "client.p12", Data: []string{"data"}}
	resp, err := svc.ValidateClientCertificate(context.Background(), req)
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_GetConnectionStatusByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetConnectionStatusMock("1")

	result, resp, err := svc.GetConnectionStatusByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Connected", result.Status)
}

func TestUnit_Digicert_GetConnectionStatusByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetConnectionStatusByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Digicert_GetConnectionStatusByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConnectionStatusNotFoundErrorMock("999")

	result, resp, err := svc.GetConnectionStatusByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Digicert_GetConnectionStatusByID_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetConnectionStatusErrorMock("1")

	result, resp, err := svc.GetConnectionStatusByID(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_Digicert_GetDependenciesByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDependenciesMock("1")

	result, resp, err := svc.GetDependenciesByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, 1, result.Results[0].ConfigProfileId)
	assert.Equal(t, "Test Config Profile", result.Results[0].ConfigProfileName)
	assert.Equal(t, "Certificate", result.Results[0].ConfigProfileType)
}

func TestUnit_Digicert_GetDependenciesByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDependenciesByID(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Digicert_GetDependenciesByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDependenciesNotFoundErrorMock("999")

	result, resp, err := svc.GetDependenciesByID(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Digicert_GetDependenciesByID_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDependenciesErrorMock("1")

	result, resp, err := svc.GetDependenciesByID(context.Background(), "1")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func strPtr(s string) *string {
	return &s
}
