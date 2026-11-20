package account_preferences

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnit_AccountPreferences_DirectoryGroupUuidColumn(t *testing.T) {
	data, err := os.ReadFile("mocks/validate_directory_group_preferences.json")
	require.NoError(t, err)
	var result ResourceAccountPreferencesV2
	require.NoError(t, json.Unmarshal(data, &result))
	require.False(t, result.ShowDirectoryGroupUuidColumn)
	for _, value := range []bool{false, true} {
		result.ShowDirectoryGroupUuidColumn = value
		body, err := json.Marshal(result)
		require.NoError(t, err)
		var fields map[string]any
		require.NoError(t, json.Unmarshal(body, &fields))
		require.Contains(t, fields, "showDirectoryGroupUuidColumn")
		require.Equal(t, value, fields["showDirectoryGroupUuidColumn"])
	}
}
