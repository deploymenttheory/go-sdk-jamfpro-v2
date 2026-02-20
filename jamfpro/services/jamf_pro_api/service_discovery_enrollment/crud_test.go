package service_discovery_enrollment

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_discovery_enrollment/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ServiceDiscoveryEnrollmentMock) {
	t.Helper()
	mock := mocks.NewServiceDiscoveryEnrollmentMock()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetWellKnownSettingsMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.WellKnownSettings, 1)
	assert.Equal(t, "Example Org", result.WellKnownSettings[0].OrgName)
	assert.Equal(t, "A1B2C3D4-E5F6-7890-ABCD-EF1234567890", result.WellKnownSettings[0].ServerUUID)
	assert.Equal(t, "USER_ENROLLMENT", result.WellKnownSettings[0].EnrollmentType)
}

func TestUnitUpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateWellKnownSettingsMock()

	request := &WellKnownSettingsResponseV1{
		WellKnownSettings: []ResourceWellKnownSettingV1{
			{OrgName: "Example Org", ServerUUID: "A1B2C3D4-E5F6-7890-ABCD-EF1234567890", EnrollmentType: "USER_ENROLLMENT"},
		},
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
	assert.Nil(t, result)
}

func TestUnitUpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
