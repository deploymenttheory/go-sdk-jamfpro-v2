package smtp_server

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smtp_server/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SMTPServerMock) {
	t.Helper()
	mock := mocks.NewSMTPServerMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV2(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.False(t, result.Enabled)
	require.Equal(t, "NONE", result.AuthenticationType)
	require.NotNil(t, result.ConnectionSettings)
	require.Equal(t, "smtp.example.com", result.ConnectionSettings.Host)
}

func TestUnitUpdateV2_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	settings := &ResourceSMTPServer{
		Enabled:            false,
		AuthenticationType: "NONE",
		ConnectionSettings: &ResourceSMTPServerConnectionSettings{Host: "smtp.example.com", Port: 587},
		SenderSettings:     &ResourceSMTPServerSenderSettings{EmailAddress: "jamf@example.com"},
	}
	result, resp, err := svc.UpdateV2(context.Background(), settings)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateV2_NilSettings(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV2(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}
