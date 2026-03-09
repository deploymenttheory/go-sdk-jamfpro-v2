package jamf_pro_system_initialization

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_system_initialization/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*JamfProSystemInitialization, *mocks.JamfProSystemInitializationMock) {
	t.Helper()
	mock := mocks.NewJamfProSystemInitializationMock()
	return NewJamfProSystemInitialization(mock), mock
}

func TestUnit_Initialize_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInitializeMock()

	request := &ResourceSystemInitialize{
		ActivationCode:  "test-activation-code",
		InstitutionName: "Test Institution",
		EulaAccepted:    true,
		Username:        "admin",
		Password:        "secret",
		JssUrl:          "https://jamf.example.com",
	}

	resp, err := svc.Initialize(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Initialize_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.Initialize(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request cannot be nil")
}

func TestUnit_Initialize_UnregisteredPath(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceSystemInitialize{
		ActivationCode:  "test",
		InstitutionName: "Test",
		EulaAccepted:    true,
		Username:        "admin",
		Password:        "secret",
		JssUrl:          "https://jamf.example.com",
	}

	resp, err := svc.Initialize(context.Background(), request)
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "JamfProSystemInitializationMock")
}

func TestUnit_InitializeDatabaseConnection_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInitializeDatabaseConnectionMock()

	resp, err := svc.InitializeDatabaseConnection(context.Background(), "db-password")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InitializeDatabaseConnection_EmptyPassword(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterInitializeDatabaseConnectionMock()

	resp, err := svc.InitializeDatabaseConnection(context.Background(), "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InitializeDatabaseConnection_UnregisteredPath(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.InitializeDatabaseConnection(context.Background(), "password")
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "JamfProSystemInitializationMock")
}

func TestUnit_PlatformInitialize_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPlatformInitializeMock()

	request := &ResourcePlatformInitialize{
		ActivationCode:  "test-activation-code",
		InstitutionName: "Test Institution",
		EulaAccepted:    true,
		Username:        "admin",
		Email:           "admin@example.com",
		JssUrl:          "https://jamf.example.com",
	}

	resp, err := svc.PlatformInitialize(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_PlatformInitialize_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.PlatformInitialize(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request cannot be nil")
}

func TestUnit_PlatformInitialize_UnregisteredPath(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourcePlatformInitialize{
		ActivationCode:  "test",
		InstitutionName: "Test",
		EulaAccepted:    true,
		Username:        "admin",
		Email:           "admin@example.com",
		JssUrl:          "https://jamf.example.com",
	}

	resp, err := svc.PlatformInitialize(context.Background(), request)
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "JamfProSystemInitializationMock")
}
