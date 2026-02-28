package jamf_pro_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/app_request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: App Request
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListFormInputFieldsV1(ctx, rsqlQuery) - Lists form input fields
//   • ReplaceFormInputFieldsV1(ctx, request) - Replaces all form input fields
//   • CreateFormInputFieldV1(ctx, request) - Creates a form input field
//   • GetFormInputFieldByIDV1(ctx, id) - Gets a form input field by int ID
//   • UpdateFormInputFieldByIDV1(ctx, id, request) - Updates a form input field
//   • DeleteFormInputFieldByIDV1(ctx, id) - Deletes a form input field
//   • GetSettingsV1(ctx) - Gets app request settings
//   • UpdateSettingsV1(ctx, request) - Updates app request settings
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle (form input fields)
//     -- Flow: Create → GetByID → Update → Delete
//
//   ✓ Pattern 2: Settings/Configuration
//     -- Flow: Get → Update → Verify → Restore
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: nil requests
//
// Notes
// -----------------------------------------------------------------------------
//   • Form input field IDs are int (not string like most Pro API resources)
//   • No ID validation on form input fields (int IDs, no empty check)
//   • Settings are a singleton resource (Get/Update only)
//
// =============================================================================

// =============================================================================
// TestAcceptance_AppRequest_form_input_fields_lifecycle
// =============================================================================

func TestAcceptance_AppRequest_form_input_fields_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AppRequest
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test app request form input field")

	title := acc.UniqueName("sdkv2_acc_app-req-field")
	createReq := &app_request.RequestFormInputField{
		Title:    title,
		Priority: 10,
	}

	created, createResp, err := svc.CreateFormInputFieldV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.Positive(t, created.ID)

	fieldID := created.ID
	acc.LogTestSuccess(t, "Form input field created with ID=%d title=%q", fieldID, title)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteFormInputFieldByIDV1(cleanupCtx, fieldID)
		acc.LogCleanupDeleteError(t, "app request form input field", "", delErr)
	})

	// 2. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching form input field by ID=%d", fieldID)

	fetched, fetchResp, err := svc.GetFormInputFieldByIDV1(ctx, fieldID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, fieldID, fetched.ID)
	assert.Equal(t, title, fetched.Title)
	acc.LogTestSuccess(t, "GetFormInputFieldByIDV1: ID=%d title=%q", fetched.ID, fetched.Title)

	// 3. Update
	acc.LogTestStage(t, "Update", "Updating form input field ID=%d", fieldID)

	updatedTitle := acc.UniqueName("sdkv2_acc_app-req-field-updated")
	updateReq := &app_request.RequestFormInputField{
		Title:    updatedTitle,
		Priority: 20,
	}
	updated, updateResp, err := svc.UpdateFormInputFieldByIDV1(ctx, fieldID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updatedTitle, updated.Title)
	acc.LogTestSuccess(t, "Form input field updated: ID=%d title=%q", fieldID, updatedTitle)

	// 4. List
	acc.LogTestStage(t, "List", "Listing form input fields")

	list, listResp, err := svc.ListFormInputFieldsV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, f := range list.Results {
		if f.ID == fieldID {
			found = true
			break
		}
	}
	assert.True(t, found, "created form input field should appear in list")

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting form input field ID=%d", fieldID)

	deleteResp, err := svc.DeleteFormInputFieldByIDV1(ctx, fieldID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Form input field ID=%d deleted", fieldID)
}

// =============================================================================
// TestAcceptance_AppRequest_settings_get_and_update
// =============================================================================

func TestAcceptance_AppRequest_settings_get_and_update(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AppRequest
	ctx := context.Background()

	// 1. Get current settings
	acc.LogTestStage(t, "Get", "Fetching app request settings")

	original, getResp, err := svc.GetSettingsV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, original)
	assert.Equal(t, 200, getResp.StatusCode)
	acc.LogTestSuccess(t, "GetSettingsV1: isEnabled=%v appStoreLocale=%q", original.IsEnabled, original.AppStoreLocale)

	// 2. Update (toggle enabled)
	acc.LogTestStage(t, "Update", "Updating app request settings")

	updateReq := &app_request.ResourceAppRequestSettings{
		IsEnabled:            !original.IsEnabled,
		AppStoreLocale:       original.AppStoreLocale,
		RequesterUserGroupID: original.RequesterUserGroupID,
		ApproverEmails:       original.ApproverEmails,
	}
	updated, updateResp, err := svc.UpdateSettingsV1(ctx, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, !original.IsEnabled, updated.IsEnabled)
	acc.LogTestSuccess(t, "Settings updated: isEnabled=%v", updated.IsEnabled)

	// 3. Restore original
	acc.LogTestStage(t, "Restore", "Restoring original app request settings")

	restored, restoreResp, err := svc.UpdateSettingsV1(ctx, original)
	require.NoError(t, err)
	require.NotNil(t, restored)
	assert.Equal(t, 200, restoreResp.StatusCode)
	assert.Equal(t, original.IsEnabled, restored.IsEnabled)
	acc.LogTestSuccess(t, "Settings restored: isEnabled=%v", restored.IsEnabled)
}

// =============================================================================
// TestAcceptance_AppRequest_validation_errors
// =============================================================================

func TestAcceptance_AppRequest_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AppRequest

	t.Run("CreateFormInputFieldV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateFormInputFieldV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateFormInputFieldByIDV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateFormInputFieldByIDV1(context.Background(), 1, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("ReplaceFormInputFieldsV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.ReplaceFormInputFieldsV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateSettingsV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateSettingsV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})
}
