package computer_inventory

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerInventory_EnrollmentFields(t *testing.T) {
	for _, version := range []string{"3", "4"} {
		t.Run("v"+version, func(t *testing.T) {
			data, err := os.ReadFile("mocks/validate_enrollment_v" + version + ".json")
			require.NoError(t, err)
			check := func(data []byte, wantNull bool, wantBool bool) {
				t.Helper()
				var enrollment *string
				var awaiting, lockdown *bool
				if version == "3" {
					var result ResourceComputerInventory
					require.NoError(t, json.Unmarshal(data, &result))
					enrollment = result.General.AppleEnrollmentType
					awaiting = result.General.AwaitingConfiguration
					lockdown = result.Security.LockdownModeEnabled
				} else {
					var result ResourceComputerInventoryV4
					require.NoError(t, json.Unmarshal(data, &result))
					enrollment = result.General.AppleEnrollmentType
					awaiting = result.General.AwaitingConfiguration
					lockdown = result.Security.LockdownModeEnabled
				}
				if wantNull {
					require.Nil(t, enrollment)
					require.Nil(t, awaiting)
					require.Nil(t, lockdown)
					return
				}
				require.NotNil(t, enrollment)
				require.Equal(t, "supervised", *enrollment)
				require.NotNil(t, awaiting)
				require.Equal(t, wantBool, *awaiting)
				require.NotNil(t, lockdown)
				require.Equal(t, wantBool, *lockdown)
			}
			check(data, true, false)
			// Exercise populated variants separately from the captured null response.
			check([]byte(`{"general":{"appleEnrollmentType":"supervised","awaitingConfiguration":false},"security":{"lockdownModeEnabled":false}}`), false, false)
			check([]byte(`{"general":{"appleEnrollmentType":"supervised","awaitingConfiguration":true},"security":{"lockdownModeEnabled":true}}`), false, true)
		})
	}
}
