package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/user_extension_attributes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_UserExtensionAttributes_lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_UserExtensionAttributes_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUserExtensionAttributes
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test user extension attribute")

	attrName := acc.UniqueName("sdkv2_acc_acc-test-userextattr")
	createReq := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        attrName,
		Description: "Acceptance test user extension attribute",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created user extension attribute ID should be a positive integer")

	attrID := created.ID
	acc.LogTestSuccess(t, "User extension attribute created with ID=%d name=%q", attrID, attrName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, attrID)
		acc.LogCleanupDeleteError(t, "user extension attribute", fmt.Sprintf("%d", attrID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new user extension attribute appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing user extension attributes to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, a := range list.UserExtensionAttributes {
		if a.ID == attrID {
			found = true
			assert.Equal(t, attrName, a.Name)
			break
		}
	}
	assert.True(t, found, "newly created user extension attribute should appear in list")
	acc.LogTestSuccess(t, "User extension attribute ID=%d found in list (%d total)", attrID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting user extension attribute by ID=%d", attrID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, attrID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, attrID, fetched.ID)
	assert.Equal(t, attrName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Getting user extension attribute by name=%q", attrName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx4, attrName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, attrID, fetchedByName.ID)
	assert.Equal(t, attrName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-userextattr-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating user extension attribute ID=%d to name=%q", attrID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        updatedName,
		Description: "Updated acceptance test user extension attribute",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx5, attrID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode())

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating user extension attribute name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        attrName,
		Description: "Acceptance test user extension attribute",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}
	reverted, revertResp, err := svc.UpdateByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode())

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetByID(ctx7, attrID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, attrName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting user extension attribute ID=%d", attrID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByID(ctx8, attrID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "User extension attribute ID=%d deleted", attrID)
}

// =============================================================================
// TestAcceptance_UserExtensionAttributes_delete_by_name creates a user extension
// attribute then deletes by name.
// =============================================================================

func TestAcceptance_UserExtensionAttributes_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUserExtensionAttributes
	ctx := context.Background()

	attrName := acc.UniqueName("sdkv2_acc_acc-test-userextattr-dbn")
	createReq := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        attrName,
		Description: "Delete by name test",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	attrID := created.ID
	acc.LogTestSuccess(t, "Created user extension attribute ID=%d name=%q for delete-by-name test", attrID, attrName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, attrID)
		acc.LogCleanupDeleteError(t, "user extension attribute", fmt.Sprintf("%d", attrID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, attrName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "User extension attribute %q deleted by name", attrName)
}

// =============================================================================
// TestAcceptance_UserExtensionAttributes_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_UserExtensionAttributes_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicUserExtensionAttributes

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user extension attribute name cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &user_extension_attributes.RequestUserExtensionAttribute{Name: "sdkv2_acc_test"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &user_extension_attributes.RequestUserExtensionAttribute{Name: "sdkv2_acc_x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user extension attribute name cannot be empty")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user extension attribute ID must be a positive integer")
	})
}
