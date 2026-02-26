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

func TestUnit_SmtpServer_GetV2_Success(t *testing.T) {
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

func TestUnit_SmtpServer_UpdateV2_Success(t *testing.T) {
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

func TestUnit_SmtpServer_UpdateV2_NilSettings(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV2(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_SmtpServer_GetHistoryV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, "1", result.Results[0].ID)
	require.Equal(t, "admin", result.Results[0].Username)
	require.Equal(t, "Sso settings update", result.Results[0].Note)
}

func TestUnit_SmtpServer_AddHistoryNoteV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &AddHistoryNoteRequest{Note: "Test note"}
	result, resp, err := svc.AddHistoryNoteV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "1", result.ID)
	require.NotEmpty(t, result.Href)
}

func TestUnit_SmtpServer_AddHistoryNoteV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.AddHistoryNoteV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_SmtpServer_AddHistoryNoteV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.AddHistoryNoteV1(context.Background(), &AddHistoryNoteRequest{Note: ""})
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}

func TestUnit_SmtpServer_TestV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &TestRequest{RecipientEmail: "test@example.com"}
	resp, err := svc.TestV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 202, resp.StatusCode)
}

func TestUnit_SmtpServer_TestV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.TestV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, resp)
}

func TestUnit_SmtpServer_TestV1_EmptyRecipientEmail(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.TestV1(context.Background(), &TestRequest{RecipientEmail: ""})
	require.Error(t, err)
	require.Nil(t, resp)
}
