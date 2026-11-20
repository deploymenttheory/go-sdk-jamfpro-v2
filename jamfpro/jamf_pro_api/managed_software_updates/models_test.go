package managed_software_updates

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUnit_ManagedSoftwareUpdates_TvOSVisionOS(t *testing.T) {
	data, err := os.ReadFile("mocks/validate_available_updates_1134.json")
	require.NoError(t, err)
	var result ResponseAvailableUpdates
	require.NoError(t, json.Unmarshal(data, &result))
	require.NotEmpty(t, result.AvailableUpdates.TvOS)
	require.NotEmpty(t, result.AvailableUpdates.VisionOS)
	require.Contains(t, result.AvailableUpdates.TvOS, "27.2.1")
	require.Contains(t, result.AvailableUpdates.VisionOS, "27.2.1")
	var device PlanDevice
	require.NoError(t, json.Unmarshal([]byte(`{"objectType":"VISION_PRO"}`), &device))
	require.Equal(t, PlanDeviceObjectTypeVisionPro, device.ObjectType)
}
