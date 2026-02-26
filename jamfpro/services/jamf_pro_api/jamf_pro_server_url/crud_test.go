package jamf_pro_server_url

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_server_url/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfProServerURLMock) {
	t.Helper()
	mock := mocks.NewJamfProServerURLMock()
	return NewService(mock), mock
}

func TestUnit_JamfProServerUrl_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "https://jamf.example.com", result.URL)
	assert.Equal(t, "http://jamf.example.com:8080", result.UnsecuredEnrollmentUrl)
}

func TestUnit_JamfProServerUrl_UpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceJamfProServerURL{
		URL:                    "https://jamf-updated.example.com",
		UnsecuredEnrollmentUrl: "http://jamf-updated.example.com:8080",
	}

	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "https://jamf-updated.example.com", result.URL)
	assert.Equal(t, "http://jamf-updated.example.com:8080", result.UnsecuredEnrollmentUrl)
}

func TestUnit_JamfProServerUrl_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_JamfProServerUrl_GetV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetV1(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get Jamf Pro server URL settings")
}

func TestUnit_JamfProServerUrl_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Server URL updated", result.Results[0].Note)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Enrollment URL modified", result.Results[1].Note)
}

func TestUnit_JamfProServerUrl_GetHistoryV1_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryMock()

	query := map[string]string{"page": "1", "page-size": "10", "sort": "date:desc"}
	result, resp, err := svc.GetHistoryV1(context.Background(), query)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
}

func TestUnit_JamfProServerUrl_GetHistoryV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get Jamf Pro server URL history")
}

func TestUnit_JamfProServerUrl_CreateHistoryNoteV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateHistoryNoteMock()

	req := &CreateHistoryNoteRequest{Note: "Added note via API"}
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
	assert.Equal(t, "admin", result.Username)
	assert.Equal(t, "Added note via API", result.Note)
}

func TestUnit_JamfProServerUrl_CreateHistoryNoteV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_JamfProServerUrl_CreateHistoryNoteV1_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &CreateHistoryNoteRequest{Note: ""}
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_JamfProServerUrl_CreateHistoryNoteV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &CreateHistoryNoteRequest{Note: "test note"}
	result, resp, err := svc.CreateHistoryNoteV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to create Jamf Pro server URL history note")
}

func TestUnit_JamfProServerUrl_UpdateV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceJamfProServerURL{
		URL:                    "https://jamf.example.com",
		UnsecuredEnrollmentUrl: "http://jamf.example.com:8080",
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to update Jamf Pro server URL settings")
}

func TestUnit_JamfProServerUrl_NewService(t *testing.T) {
	mock := mocks.NewJamfProServerURLMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}
