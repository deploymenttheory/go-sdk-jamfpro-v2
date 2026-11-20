package conditional_access

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_ConditionalAccess_RegistrationGracePeriod(t *testing.T) {
	// Synthetic shape: the live test computer has no compliance registrations.
	var result ComplianceVendorDeviceInformation
	require.NoError(t, json.Unmarshal([]byte(`{"deviceIds":[],"registrations":[{"entraDeviceId":"device","entraUserId":null,"sharedDevice":true,"gracePeriodUntil":null}],"gracePeriodSettingsState":{"enabled":false,"durationHours":null}}`), &result))
	require.Len(t, result.Registrations, 1)
	require.Nil(t, result.Registrations[0].EntraUserID)
	require.True(t, *result.Registrations[0].SharedDevice)
	require.Nil(t, result.Registrations[0].GracePeriodUntil)
	require.False(t, result.GracePeriodSettingsState.Enabled)
	require.Nil(t, result.GracePeriodSettingsState.DurationHours)
}
