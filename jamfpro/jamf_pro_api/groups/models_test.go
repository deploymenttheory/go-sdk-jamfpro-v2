package groups

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUnit_Groups_DivisionID(t *testing.T) {
	data, err := os.ReadFile("mocks/validate_division.json")
	require.NoError(t, err)
	var result ResourceGroup
	require.NoError(t, json.Unmarshal(data, &result))
	require.Nil(t, result.DivisionID)
	require.NoError(t, json.Unmarshal([]byte(`{"divisionId":"example"}`), &result))
	require.Equal(t, "example", *result.DivisionID)
}
