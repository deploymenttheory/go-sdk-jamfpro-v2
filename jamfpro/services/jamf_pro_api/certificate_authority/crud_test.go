package certificate_authority

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/certificate_authority/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.CertificateAuthorityMock) {
	t.Helper()
	mock := mocks.NewCertificateAuthorityMock()
	return NewService(mock), mock
}

func TestUnitGetActiveCertificateAuthorityV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActiveCertificateAuthorityMock()

	result, resp, err := svc.GetActiveCertificateAuthorityV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "CN=Jamf Pro CA", result.SubjectX500Principal)
	assert.Equal(t, "12345", result.SerialNumber)
	assert.Len(t, result.KeyUsage, 2)
}

func TestUnitGetActiveCertificateAuthorityDERV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActiveCertificateAuthorityDERMock()

	resp, data, err := svc.GetActiveCertificateAuthorityDERV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestUnitGetActiveCertificateAuthorityPEMV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetActiveCertificateAuthorityPEMMock()

	resp, data, err := svc.GetActiveCertificateAuthorityPEMV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, data)
	assert.Contains(t, string(data), "BEGIN CERTIFICATE")
}

func TestUnitGetCertificateAuthorityByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCertificateAuthorityByIDMock("1")

	result, resp, err := svc.GetCertificateAuthorityByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "CN=Jamf Pro CA", result.SubjectX500Principal)
}

func TestUnitGetCertificateAuthorityByIDDERV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCertificateAuthorityByIDMock("1")

	resp, data, err := svc.GetCertificateAuthorityByIDDERV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, data)
	assert.Greater(t, len(data), 0)
}

func TestUnitGetCertificateAuthorityByIDPEMV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCertificateAuthorityByIDMock("1")

	resp, data, err := svc.GetCertificateAuthorityByIDPEMV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, data)
	assert.Contains(t, string(data), "BEGIN CERTIFICATE")
}
