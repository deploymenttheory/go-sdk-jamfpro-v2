package access_management_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/access_management_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.AccessManagementSettingsMock) {
	t.Helper()
	mock := mocks.NewAccessManagementSettingsMock()
	return NewService(mock), mock
}

func TestUnitGetV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV4(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitCreateV4_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV4(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitCreateV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPostMock()

	request := &ResourceAccessManagementSettings{AutomatedDeviceEnrollmentServerUuid: ""}
	result, resp, err := svc.CreateV4(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}
