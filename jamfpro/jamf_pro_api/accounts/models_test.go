package accounts

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnit_Accounts_RopfMigrationScope(t *testing.T) {
	for _, scope := range []string{"AUTO", "EXCLUDED"} {
		t.Run(scope, func(t *testing.T) {
			body, err := json.Marshal(RequestAccount{RopfMigrationScope: scope})
			require.NoError(t, err)
			var result ResourceAccount
			require.NoError(t, json.Unmarshal(body, &result))
			require.Equal(t, scope, result.RopfMigrationScope)
		})
	}
}
