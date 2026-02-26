package account_preferences

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/account_preferences/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AccountPreferencesMock) {
	t.Helper()
	mock := mocks.NewAccountPreferencesMock()
	return NewService(mock), mock
}

func TestUnit_AccountPreferences_GetV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountPreferencesMock()

	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "en", result.Language)
	assert.Equal(t, "DARK", result.UserInterfaceDisplayTheme)
}

func TestUnit_AccountPreferences_UpdateV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountPreferencesMock()

	request := &ResourceAccountPreferencesV2{
		Language:                  "en",
		DateFormat:                "MM/dd/yyyy",
		Timezone:                  "America/Chicago",
		UserInterfaceDisplayTheme: "DARK",
		ResultsPerPage:            20,
	}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "en", result.Language)
}

func TestUnit_AccountPreferences_GetV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_AccountPreferences_UpdateV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_AccountPreferences_UpdateV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceAccountPreferencesV2{Language: "en"}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}
