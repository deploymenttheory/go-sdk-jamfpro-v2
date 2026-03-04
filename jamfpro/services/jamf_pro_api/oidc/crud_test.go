package oidc

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/oidc/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Oidc_GetDirectIdPLoginURLV1_Success(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterGetDirectIdPLoginURLMock()
	service := NewService(mock)

	result, resp, err := service.GetDirectIdPLoginURLV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "https://idp.example.com/authorize?client_id=jamfpro", result.URL)
}

func TestUnit_Oidc_GetPublicKeyV1_Success(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterGetPublicKeyMock()
	service := NewService(mock)

	result, resp, err := service.GetPublicKeyV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Len(t, result.Keys, 1)
	assert.Equal(t, "RSA", result.Keys[0].Kty)
	assert.Equal(t, "AQAB", result.Keys[0].E)
	assert.Equal(t, "sig", result.Keys[0].Use)
	assert.Equal(t, "test-key-id", result.Keys[0].Kid)
	assert.Equal(t, "RS256", result.Keys[0].Alg)
	assert.Equal(t, int64(1609459200), result.Keys[0].Iat)
	assert.Equal(t, "xGOr-H7A...", result.Keys[0].N)
}

func TestUnit_Oidc_GenerateCertificateV1_Success(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterGenerateCertificateMock()
	service := NewService(mock)

	resp, err := service.GenerateCertificateV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Oidc_GetRedirectURLV1_Success(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterGetRedirectURLMock()
	service := NewService(mock)

	request := &RequestOIDCRedirectURL{
		OriginalURL:  "https://jamf.example.com/dashboard",
		EmailAddress: "user@example.com",
	}

	result, resp, err := service.GetRedirectURLV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "https://idp.example.com/login?redirect=https://jamf.example.com", result.RedirectURL)
}

func TestUnit_Oidc_GetRedirectURLV1_NilRequest(t *testing.T) {
	mock := mocks.NewOIDCMock()
	service := NewService(mock)

	result, resp, err := service.GetRedirectURLV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "OIDC redirect URL request cannot be nil")
}

func TestUnit_Oidc_GetPublicFeaturesV1_Success(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterGetPublicFeaturesMock()
	service := NewService(mock)

	result, resp, err := service.GetPublicFeaturesV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.JamfIdAuthenticationEnabled)
}

func TestUnit_Oidc_GetDirectIdPLoginURLV1_Error(t *testing.T) {
	mock := mocks.NewOIDCMock()
	// Do not register mock - triggers no mock registered error
	service := NewService(mock)

	result, resp, err := service.GetDirectIdPLoginURLV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get OIDC direct IdP login URL")
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Oidc_GetPublicKeyV1_Error(t *testing.T) {
	mock := mocks.NewOIDCMock()
	service := NewService(mock)

	result, resp, err := service.GetPublicKeyV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get OIDC public key")
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Oidc_GetPublicFeaturesV1_Error(t *testing.T) {
	mock := mocks.NewOIDCMock()
	service := NewService(mock)

	result, resp, err := service.GetPublicFeaturesV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get OIDC public features")
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Oidc_GenerateCertificateV1_Error(t *testing.T) {
	mock := mocks.NewOIDCMock()
	service := NewService(mock)

	resp, err := service.GenerateCertificateV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to generate OIDC certificate")
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Oidc_GetRedirectURLV1_Error(t *testing.T) {
	mock := mocks.NewOIDCMock()
	service := NewService(mock)

	request := &RequestOIDCRedirectURL{
		OriginalURL:  "https://jamf.example.com/dashboard",
		EmailAddress: "user@example.com",
	}

	result, resp, err := service.GetRedirectURLV1(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get OIDC redirect URL")
	assert.Contains(t, err.Error(), "no mock registered")
}

func TestUnit_Oidc_NewService(t *testing.T) {
	mock := mocks.NewOIDCMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_Oidc_RegisterMocks(t *testing.T) {
	mock := mocks.NewOIDCMock()
	mock.RegisterMocks()
	service := NewService(mock)

	// Verify all mocks work
	_, _, err := service.GetDirectIdPLoginURLV1(context.Background())
	require.NoError(t, err)

	_, _, err = service.GetPublicKeyV1(context.Background())
	require.NoError(t, err)

	_, _, err = service.GetPublicFeaturesV1(context.Background())
	require.NoError(t, err)

	_, err = service.GenerateCertificateV1(context.Background())
	require.NoError(t, err)

	_, _, err = service.GetRedirectURLV1(context.Background(), &RequestOIDCRedirectURL{
		OriginalURL: "https://example.com", EmailAddress: "u@ex.com",
	})
	require.NoError(t, err)
}
