package tomcat_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/tomcat_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*TomcatSettings, *mocks.TomcatSettingsMock) {
	t.Helper()
	mock := mocks.NewTomcatSettingsMock()
	return NewTomcatSettings(mock), mock
}

func TestUnit_TomcatSettings_NewService(t *testing.T) {
	mock := mocks.NewTomcatSettingsMock()
	svc := NewTomcatSettings(mock)
	require.NotNil(t, svc)
}

func TestUnit_TomcatSettings_IssueTomcatSslCertificate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterIssueTomcatSslCertificateMock()

	resp, err := svc.IssueTomcatSslCertificate(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_TomcatSettings_IssueTomcatSslCertificate_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterIssueTomcatSslCertificateErrorMock()

	resp, err := svc.IssueTomcatSslCertificate(context.Background())
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "Jamf Pro API error")
	assert.Contains(t, err.Error(), "CERTIFICATE_ISSUE_FAILED")
}

func TestUnit_TomcatSettings_IssueTomcatSslCertificate_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.IssueTomcatSslCertificate(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_TomcatSettings_InterfaceCompliance(t *testing.T) {
	var _ TomcatSettingsServiceInterface = (*TomcatSettings)(nil)
}
