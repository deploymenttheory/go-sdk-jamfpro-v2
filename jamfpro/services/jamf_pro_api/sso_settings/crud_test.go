package sso_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SsoSettingsMock) {
	t.Helper()
	mock := mocks.NewSsoSettingsMock()
	return NewService(mock), mock
}

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.SsoEnabled)
	assert.Equal(t, "OIDC", result.ConfigurationType)
	require.NotNil(t, result.OidcSettings)
	assert.Equal(t, "USERNAME", result.OidcSettings.UserMapping)
}

func TestUnitUpdate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	request := &ResourceSsoSettings{
		SsoEnabled:       false,
		ConfigurationType: "OIDC",
		SsoBypassAllowed: true,
	}
	result, resp, err := svc.UpdateV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitGetEnrollmentCustomizationDependencies_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDependenciesMock()

	result, resp, err := svc.GetEnrollmentCustomizationDependenciesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.Dependencies, 1)
	assert.Equal(t, "Enrollment Customization A", result.Dependencies[0].Name)
	assert.Equal(t, "Enrollment Customization A", result.Dependencies[0].HumanReadableName)
}
