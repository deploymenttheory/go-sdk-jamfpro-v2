package mobile_devices

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDevices_EnrollmentFields(t *testing.T) {
	data, err := os.ReadFile("mocks/validate_enrollment_detail.json")
	require.NoError(t, err)
	var result ResourceMobileDeviceDetailsV2
	require.NoError(t, json.Unmarshal(data, &result))
	require.Nil(t, result.AppleEnrollmentType)
	require.Nil(t, result.AwaitingConfiguration)
	// Populated variants are synthetic; the temporary device reported null.
	require.NoError(t, json.Unmarshal([]byte(`{"appleEnrollmentType":"supervised","awaitingConfiguration":false,"ios":{"returnToServiceEnabled":false,"systemHealth":{"camera":"ok"},"security":{"lockdownModeEnabled":false}}}`), &result))
	require.Equal(t, "supervised", *result.AppleEnrollmentType)
	require.NotNil(t, result.AwaitingConfiguration)
	require.False(t, *result.AwaitingConfiguration)
	require.NotNil(t, result.Ios.ReturnToServiceEnabled)
	require.False(t, *result.Ios.ReturnToServiceEnabled)
	require.Equal(t, "ok", *result.Ios.SystemHealth.Camera)
	require.NotNil(t, result.Ios.Security.LockdownModeEnabled)
	require.False(t, *result.Ios.Security.LockdownModeEnabled)
}
