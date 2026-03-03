package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/activation_code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// We should not atttempt to update the activation code in this test suite as it risks
// breaking the jamf instance.

// TestAcceptance_ActivationCode_get_history_v1 tests retrieving activation code history with various RSQL queries.
func TestAcceptance_ActivationCode_get_history_v1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	// First, add a history note to ensure we have data to test with
	t.Run("AddHistoryNote", func(t *testing.T) {
		noteReq := &activation_code.HistoryNoteRequest{
			Note: "Acceptance test history note for activation code",
		}
		result, resp, err := client.ActivationCode.AddHistoryNoteV1(context.Background(), noteReq)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 201, resp.StatusCode)
		assert.NotEmpty(t, result.ID)
		t.Logf("Added history note with ID: %d", result.ID)
	})

	// Test 1: Get all history (no filter)
	t.Run("GetAll", func(t *testing.T) {
		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), nil)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)
		assert.GreaterOrEqual(t, result.TotalCount, 1, "Should have at least the note we just added")
		t.Logf("Found %d total activation code history entries", result.TotalCount)
	})

	// Test 2: Get history with RSQL filter for specific username
	t.Run("FilterByUsername", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `username==admin`,
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries match the filter
		for _, entry := range result.Results {
			assert.Equal(t, "admin", entry.Username, "Expected all entries to have username 'admin'")
		}
		t.Logf("Found %d history entries for username 'admin'", len(result.Results))
	})

	// Test 3: Get history with RSQL filter for specific details
	t.Run("FilterByDetails", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `details==enabled`,
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries match the filter
		for _, entry := range result.Results {
			assert.Equal(t, "enabled", entry.Details, "Expected all entries to have details 'enabled'")
		}
		t.Logf("Found %d history entries with details 'enabled'", len(result.Results))
	})

	// Test 4: Get history with complex RSQL filter
	t.Run("ComplexRSQLFilter", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"filter": `username!=admin and details==enabled`,
			"sort":   "date:desc",
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)

		// Verify all returned entries match the filter criteria
		for _, entry := range result.Results {
			assert.NotEqual(t, "admin", entry.Username)
			assert.Equal(t, "enabled", entry.Details)
		}
		t.Logf("Found %d history entries matching complex filter", len(result.Results))
	})

	// Test 5: Get history with pagination parameters
	t.Run("WithPagination", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"page":      "0",
			"page-size": "5",
			"sort":      "date:desc",
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		// Note: GetPaginated fetches all pages, so we should get all results
		t.Logf("Retrieved %d history entries (pagination handled automatically)", len(result.Results))
	})

	// Test 6: Get history with sorting
	t.Run("WithSorting", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "date:asc",
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Retrieved %d history entries sorted by date ascending", len(result.Results))
	})

	// Test 7: Get history with multiple sort fields
	t.Run("WithMultipleSortFields", func(t *testing.T) {
		rsqlQuery := map[string]string{
			"sort": "username:asc,date:desc",
		}

		result, resp, err := client.ActivationCode.GetHistoryV1(context.Background(), rsqlQuery)
		require.NoError(t, err)
		require.NotNil(t, result)
		assert.Equal(t, 200, resp.StatusCode)
		t.Logf("Retrieved %d history entries with multiple sort fields", len(result.Results))
	})
}
