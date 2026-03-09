package mobile_device_enrollment_profile

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_enrollment_profile/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*MobileDeviceEnrollmentProfile, *mocks.MobileDeviceEnrollmentProfileMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceEnrollmentProfileMock()
	return NewMobileDeviceEnrollmentProfile(mock), mock
}

func TestUnit_MobileDeviceEnrollmentProfile_NewService(t *testing.T) {
	mock := mocks.NewMobileDeviceEnrollmentProfileMock()
	svc := NewMobileDeviceEnrollmentProfile(mock)
	require.NotNil(t, svc)
}

func TestUnit_MobileDeviceEnrollmentProfile_GetDownloadProfileV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDownloadProfileMock("1")

	data, resp, err := svc.GetDownloadProfileV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, data)
	assert.Contains(t, string(data), "fake-profile-data-for-testing")
}

func TestUnit_MobileDeviceEnrollmentProfile_GetDownloadProfileV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.GetDownloadProfileV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDeviceEnrollmentProfile_GetDownloadProfileV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	data, resp, err := svc.GetDownloadProfileV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, data)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "failed to download profile")
}

func TestUnit_MobileDeviceEnrollmentProfile_GetDownloadProfileV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.GetDownloadProfileV1(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to download profile")
}

func TestUnit_MobileDeviceEnrollmentProfile_InterfaceCompliance(t *testing.T) {
	var _ MobileDeviceEnrollmentProfileServiceInterface = (*MobileDeviceEnrollmentProfile)(nil)
}
