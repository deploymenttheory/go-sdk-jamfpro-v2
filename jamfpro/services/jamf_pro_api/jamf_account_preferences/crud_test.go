package jamf_account_preferences

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_account_preferences/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.JamfAccountPreferencesMock) {
	t.Helper()
	mock := mocks.NewJamfAccountPreferencesMock()
	return NewService(mock), mock
}

func TestUnit_JamfAccountPreferences_NewService(t *testing.T) {
	mock := mocks.NewJamfAccountPreferencesMock()
	svc := NewService(mock)
	require.NotNil(t, svc)
}

func TestUnit_JamfAccountPreferences_GetV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV3Mock()

	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "en", result.Language)
	assert.Equal(t, "MM/dd/yyyy", result.DateFormat)
	assert.Equal(t, "America/Chicago", result.Timezone)
	assert.Equal(t, 25, result.ResultsPerPage)
	assert.Equal(t, "MATCH_SYSTEM", result.UserInterfaceDisplayTheme)
	assert.False(t, result.DisableRelativeDates)
	assert.Equal(t, "EXACT_MATCH", result.ComputerSearchMethod)
	assert.Equal(t, "STARTS_WITH", result.ComputerApplicationSearchMethod)
	assert.Equal(t, "CONTAINS", result.UserEbookSearchMethod)
}

func TestUnit_JamfAccountPreferences_GetV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV3(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfAccountPreferences_GetV3_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV3ErrorMock()

	result, resp, err := svc.GetV3(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
	assert.Contains(t, err.Error(), "NOT-FOUND")
}

func TestUnit_JamfAccountPreferences_GetV3_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInvalidJSONMock()

	result, resp, err := svc.GetV3(context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Contains(t, err.Error(), "unmarshal")
}

func TestUnit_JamfAccountPreferences_UpdateV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateV3Mock()

	req := &ResourceAccountPreferences{
		Language:                  "en",
		DateFormat:                "MM/dd/yyyy",
		Timezone:                  "America/New_York",
		ResultsPerPage:            50,
		UserInterfaceDisplayTheme: "DARK",
		DisableRelativeDates:      true,
	}

	result, resp, err := svc.UpdateV3(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "en", result.Language)
	assert.Equal(t, "America/New_York", result.Timezone)
	assert.Equal(t, 50, result.ResultsPerPage)
	assert.Equal(t, "DARK", result.UserInterfaceDisplayTheme)
	assert.True(t, result.DisableRelativeDates)
}

func TestUnit_JamfAccountPreferences_UpdateV3_204NoContent(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateV3_204NoContentMock()

	req := &ResourceAccountPreferences{
		Language:                  "en",
		DateFormat:                "dd/MM/yyyy",
		Timezone:                  "Europe/London",
		ResultsPerPage:            100,
		UserInterfaceDisplayTheme: "LIGHT",
	}

	result, resp, err := svc.UpdateV3(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode)
	assert.Equal(t, req, result)
	assert.Equal(t, "en", result.Language)
	assert.Equal(t, "Europe/London", result.Timezone)
	assert.Equal(t, 100, result.ResultsPerPage)
	assert.Equal(t, "LIGHT", result.UserInterfaceDisplayTheme)
}

func TestUnit_JamfAccountPreferences_UpdateV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV3(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_JamfAccountPreferences_UpdateV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceAccountPreferences{Language: "en"}
	result, resp, err := svc.UpdateV3(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response registered")
}

func TestUnit_JamfAccountPreferences_UpdateV3_ServerError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateV3ErrorMock()

	req := &ResourceAccountPreferences{Language: "en"}
	result, resp, err := svc.UpdateV3(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}
