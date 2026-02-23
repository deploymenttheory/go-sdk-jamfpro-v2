package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/vpp_assignments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_VPPAssignments_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → UpdateByID → DeleteByID.
// Note: VPP assignments require an existing VPP admin account. The test uses
// the first available VPP account from List; if none exist, the test is skipped.
// =============================================================================

func TestAcceptance_VPPAssignments_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicVPPAssignments
	vppSvc := acc.Client.ClassicVPPAccounts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 0. Get a VPP account ID (required for creating assignments)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Prereq", "Listing VPP accounts to find one for assignment")

	ctx0, cancel0 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel0()

	accounts, _, err := vppSvc.List(ctx0)
	require.NoError(t, err, "List VPP accounts should not return an error")
	if accounts == nil || len(accounts.Results) == 0 {
		t.Skip("No VPP accounts found; VPP assignments require an existing VPP account")
	}

	vppAccountID := accounts.Results[0].ID
	vppAccountName := accounts.Results[0].Name
	acc.LogTestSuccess(t, "Using VPP account ID=%d name=%q", vppAccountID, vppAccountName)

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test VPP assignment")

	assignmentName := acc.UniqueName("acc-test-vpp-assignment")
	createReq := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                assignmentName,
			VPPAdminAccountID:   vppAccountID,
			VPPAdminAccountName: vppAccountName,
		},
		Scope: vpp_assignments.SubsetScope{
			AllJSSUsers: false,
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 409 {
		t.Skip("VPP assignment create may require additional setup in this environment; skipping lifecycle")
	}
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	// Create may return ID 0 if API does not return it; we use List to find the created assignment
	acc.LogTestSuccess(t, "VPP assignment create response status=%d", createResp.StatusCode)

	// ------------------------------------------------------------------
	// 2. List — verify the new assignment appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing VPP assignments to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	var assignmentID int
	found := false
	for _, a := range list.VPPAssignments {
		if a.Name == assignmentName {
			found = true
			assignmentID = a.ID
			break
		}
	}
	require.True(t, found, "newly created VPP assignment should appear in list")
	assert.Positive(t, assignmentID, "assignment ID should be positive")
	acc.LogTestSuccess(t, "VPP assignment ID=%d found in list", assignmentID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, assignmentID)
		acc.LogCleanupDeleteError(t, "VPP assignment", fmt.Sprintf("%d", assignmentID), delErr)
	})

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching VPP assignment by ID=%d", assignmentID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByID(ctx3, assignmentID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, assignmentID, fetched.General.ID)
	assert.Equal(t, assignmentName, fetched.General.Name)
	assert.Equal(t, vppAccountID, fetched.General.VPPAdminAccountID)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)

	// ------------------------------------------------------------------
	// 4. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("acc-test-vpp-assignment-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating VPP assignment ID=%d to name=%q", assignmentID, updatedName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updateReq := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                updatedName,
			VPPAdminAccountID:   vppAccountID,
			VPPAdminAccountName: vppAccountName,
		},
		Scope: vpp_assignments.SubsetScope{
			AllJSSUsers: false,
		},
	}
	_, updateResp, err := svc.UpdateByID(ctx4, assignmentID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 5. GetByID — verify update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name update")

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	verified, verifyResp, err := svc.GetByID(ctx5, assignmentID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, updatedName, verified.General.Name, "name should reflect the update")
	acc.LogTestSuccess(t, "Name update verified: %q", verified.General.Name)

	// ------------------------------------------------------------------
	// 6. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting VPP assignment ID=%d", assignmentID)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	deleteResp, err := svc.DeleteByID(ctx6, assignmentID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "VPP assignment ID=%d deleted", assignmentID)
}

// =============================================================================
// TestAcceptance_VPPAssignments_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_VPPAssignments_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicVPPAssignments

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("Create_EmptyName", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), &vpp_assignments.Resource{
			General: vpp_assignments.SubsetGeneral{Name: ""},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP assignment name is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 0, &vpp_assignments.Resource{
			General: vpp_assignments.SubsetGeneral{Name: "x"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
	})

	t.Run("UpdateByID_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), 1, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "VPP assignment ID must be a positive integer")
	})
}
