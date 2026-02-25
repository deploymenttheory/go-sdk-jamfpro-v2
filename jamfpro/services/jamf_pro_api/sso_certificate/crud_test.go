package sso_certificate

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_certificate/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SsoCertificateMock) {
	t.Helper()
	mock := mocks.NewSsoCertificateMock()
	return NewService(mock), mock
}

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "SELF_SERVICE", result.Keystore.Type)
}

func TestUnitCreate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	result, resp, err := svc.CreateV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDelete_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteMock()

	resp, err := svc.DeleteV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_SsoCertificate_UpdateV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	req := &UpdateKeystoreRequest{
		Type:              "PKCS12",
		KeystoreSetupType: "UPLOADED",
		KeystorePassword:  "password",
	}
	result, resp, err := svc.UpdateV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "SELF_SERVICE", result.Keystore.Type)
}

func TestUnit_SsoCertificate_DownloadV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDownloadMock()

	data, resp, err := svc.DownloadV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	_ = data
}

func TestUnit_SsoCertificate_ParseV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterParseMock()

	req := &ParseKeystoreRequest{
		KeystoreFile:     "base64encodedfile",
		KeystorePassword: "password",
		KeystoreFileName: "keystore.p12",
	}
	result, resp, err := svc.ParseV2(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "PKCS12", result.Type)
	assert.Equal(t, "UPLOADED", result.KeystoreSetupType)
	require.Len(t, result.Keys, 1)
	assert.Equal(t, "key-1", result.Keys[0].ID)
	assert.True(t, result.Keys[0].Valid)
}
