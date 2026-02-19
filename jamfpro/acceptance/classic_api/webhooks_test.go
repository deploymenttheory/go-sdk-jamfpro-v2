package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Webhooks_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Webhooks_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Webhooks
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test webhook")

	webhookName := uniqueName("acc-test-webhook")
	createReq := &webhooks.RequestWebhook{
		Name:               webhookName,
		Enabled:            true,
		URL:                "https://hooks.example.com/test",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		ConnectionTimeout:  5000,
		ReadTimeout:        5000,
		AuthenticationType: "NONE",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateWebhook(ctx1, createReq)
	require.NoError(t, err, "CreateWebhook should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created webhook ID should be a positive integer")

	webhookID := created.ID
	acc.LogTestSuccess(t, "Webhook created with ID=%d name=%q", webhookID, webhookName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteWebhookByID(cleanupCtx, webhookID)
		acc.LogCleanupDeleteError(t, "webhook", fmt.Sprintf("%d", webhookID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new webhook appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing webhooks to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListWebhooks(ctx2)
	require.NoError(t, err, "ListWebhooks should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, w := range list.Results {
		if w.ID == webhookID {
			found = true
			assert.Equal(t, webhookName, w.Name)
			break
		}
	}
	assert.True(t, found, "newly created webhook should appear in list")
	acc.LogTestSuccess(t, "Webhook ID=%d found in list (%d total)", webhookID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching webhook by ID=%d", webhookID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetWebhookByID(ctx3, webhookID)
	require.NoError(t, err, "GetWebhookByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, webhookID, fetched.ID)
	assert.Equal(t, webhookName, fetched.Name)
	assert.True(t, fetched.Enabled)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q enabled=%v", fetched.ID, fetched.Name, fetched.Enabled)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching webhook by name=%q", webhookName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetWebhookByName(ctx4, webhookName)
	require.NoError(t, err, "GetWebhookByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, webhookID, fetchedByName.ID)
	assert.Equal(t, webhookName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-webhook-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating webhook ID=%d to name=%q", webhookID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &webhooks.RequestWebhook{
		Name:               updatedName,
		Enabled:            true,
		URL:                "https://hooks.example.com/test",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		ConnectionTimeout:  5000,
		ReadTimeout:        5000,
		AuthenticationType: "NONE",
	}
	updated, updateResp, err := svc.UpdateWebhookByID(ctx5, webhookID, updateReq)
	require.NoError(t, err, "UpdateWebhookByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating webhook name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &webhooks.RequestWebhook{
		Name:               webhookName,
		Enabled:            true,
		URL:                "https://hooks.example.com/test",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		ConnectionTimeout:  5000,
		ReadTimeout:        5000,
		AuthenticationType: "NONE",
	}
	reverted, revertResp, err := svc.UpdateWebhookByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateWebhookByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetWebhookByID(ctx7, webhookID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, webhookName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting webhook ID=%d", webhookID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteWebhookByID(ctx8, webhookID)
	require.NoError(t, err, "DeleteWebhookByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Webhook ID=%d deleted", webhookID)
}

// =============================================================================
// TestAcceptance_Webhooks_DeleteByName creates a webhook then deletes by name.
// =============================================================================

func TestAcceptance_Webhooks_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Webhooks
	ctx := context.Background()

	webhookName := uniqueName("acc-test-webhook-dbn")
	createReq := &webhooks.RequestWebhook{
		Name:               webhookName,
		Enabled:            false,
		URL:                "https://hooks.example.com/dbn",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		AuthenticationType: "NONE",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateWebhook(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	webhookID := created.ID
	acc.LogTestSuccess(t, "Created webhook ID=%d name=%q for delete-by-name test", webhookID, webhookName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteWebhookByID(cleanupCtx, webhookID)
		acc.LogCleanupDeleteError(t, "webhook", fmt.Sprintf("%d", webhookID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteWebhookByName(ctx2, webhookName)
	require.NoError(t, err, "DeleteWebhookByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Webhook %q deleted by name", webhookName)
}

// =============================================================================
// TestAcceptance_Webhooks_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Webhooks_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Webhooks

	t.Run("GetWebhookByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetWebhookByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
	})

	t.Run("GetWebhookByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetWebhookByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook name is required")
	})

	t.Run("CreateWebhook_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateWebhook(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateWebhookByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateWebhookByID(context.Background(), 0, &webhooks.RequestWebhook{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
	})

	t.Run("UpdateWebhookByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateWebhookByName(context.Background(), "", &webhooks.RequestWebhook{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook name is required")
	})

	t.Run("DeleteWebhookByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteWebhookByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook ID must be a positive integer")
	})

	t.Run("DeleteWebhookByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteWebhookByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "webhook name is required")
	})
}
