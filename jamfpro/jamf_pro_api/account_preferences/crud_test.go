package account_preferences

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/account_preferences/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*AccountPreferences, *mocks.AccountPreferencesMock) {
	t.Helper()
	mock := mocks.NewAccountPreferencesMock()
	return NewAccountPreferences(mock), mock
}

func TestUnit_AccountPreferences_GetV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountPreferencesMock()

	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_AccountPreferences_UpdateV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountPreferencesMock()

	request := &ResourceAccountPreferencesV2{
		Language:         "en",
		DateFormat:       "MM/dd/yyyy",
		Timezone:         "America/New_York",
		DisableRelativeDates: false,
	}

	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_AccountPreferences_GetV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccountPreferencesErrorMock()

	result, resp, err := svc.GetV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
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
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccountPreferencesErrorMock()

	request := &ResourceAccountPreferencesV2{Language: "en"}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}
