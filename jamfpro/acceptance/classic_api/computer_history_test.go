package classic_api

import (
	"context"
	"os"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_ComputerHistory_GetByID tests retrieving computer history by ID.
// Requires at least one computer in Jamf Pro. Set COMPUTER_ID to skip inventory lookup.
// =============================================================================

func TestAcceptance_ComputerHistory_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerHistory
	ctx := context.Background()

	var computerID string
	if raw := os.Getenv("COMPUTER_ID"); raw != "" {
		computerID = raw
	} else {
		invSvc := acc.Client.ComputerInventory
		list, _, err := invSvc.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			t.Skip("Set COMPUTER_ID or ensure at least one computer exists in inventory")
		}
		computerID = list.Results[0].ID
	}

	acc.LogTestStage(t, "GetByID", "Fetching computer history by ID=%s", computerID)

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	history, resp, err := svc.GetByID(ctx1, computerID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, history)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, history.General.Name, "computer name should be populated")
	acc.LogTestSuccess(t, "Computer history retrieved: ID=%s name=%q", computerID, history.General.Name)
}

// =============================================================================
// TestAcceptance_ComputerHistory_GetByIDAndSubset tests retrieving a subset of
// computer history by ID.
// =============================================================================

func TestAcceptance_ComputerHistory_GetByIDAndSubset(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerHistory
	ctx := context.Background()

	var computerID string
	if raw := os.Getenv("COMPUTER_ID"); raw != "" {
		computerID = raw
	} else {
		invSvc := acc.Client.ComputerInventory
		list, _, err := invSvc.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			t.Skip("Set COMPUTER_ID or ensure at least one computer exists in inventory")
		}
		computerID = list.Results[0].ID
	}

	acc.LogTestStage(t, "GetByIDAndSubset", "Fetching computer history subset General by ID=%s", computerID)

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	history, resp, err := svc.GetByIDAndSubset(ctx1, computerID, "General")
	require.NoError(t, err, "GetByIDAndSubset should not return an error")
	require.NotNil(t, history)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, history.General.Name)
	acc.LogTestSuccess(t, "Computer history subset retrieved: ID=%s", computerID)
}

// =============================================================================
// TestAcceptance_ComputerHistory_GetByName tests retrieving computer history by name.
// =============================================================================

func TestAcceptance_ComputerHistory_GetByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerHistory
	ctx := context.Background()

	computerName := os.Getenv("COMPUTER_NAME")
	if computerName == "" {
		invSvc := acc.Client.ComputerInventory
		list, _, err := invSvc.ListV3(ctx, nil)
		if err != nil || list == nil || len(list.Results) == 0 {
			t.Skip("Set COMPUTER_NAME or ensure at least one computer exists in inventory")
		}
		computerName = list.Results[0].General.Name
	}

	acc.LogTestStage(t, "GetByName", "Fetching computer history by name=%q", computerName)

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	history, resp, err := svc.GetByName(ctx1, computerName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, history)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, computerName, history.General.Name)
	acc.LogTestSuccess(t, "Computer history retrieved by name: %q", computerName)
}

// =============================================================================
// TestAcceptance_ComputerHistory_ValidationErrors tests validation error handling.
// =============================================================================

func TestAcceptance_ComputerHistory_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerHistory
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ID cannot be empty")
	})

	t.Run("GetByIDAndSubset_EmptySubset", func(t *testing.T) {
		_, _, err := svc.GetByIDAndSubset(ctx, "1", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "subset cannot be empty")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "name cannot be empty")
	})
}
