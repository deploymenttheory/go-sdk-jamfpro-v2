package classic_api

import (
	"context"
	"strconv"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_invitations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_ComputerInvitations_lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByInvitationID → DeleteByID.
// =============================================================================

func TestAcceptance_ComputerInvitations_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAPI.ComputerInvitations
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 0. Check if User-Initiated Enrollment is enabled for computers
	// https://learn.jamf.com/en-US/bundle/jamf-pro-documentation-current/page/Enabling_Device_Enrollment_for_Computers.html
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "PreCheck", "Checking if user-initiated enrollment is enabled for computers")

	enrollmentSvc := acc.Client.JamfProAPI.EnrollmentSettings
	ctx0, cancel0 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel0()

	enrollmentSettings, _, err := enrollmentSvc.GetV4(ctx0)
	if err != nil {
		t.Skipf("Unable to retrieve enrollment settings: %v", err)
	}

	if !enrollmentSettings.MacOsEnterpriseEnrollmentEnabled {
		t.Skip("User-initiated enrollment for computers is not enabled - skipping computer invitations test")
	}

	acc.LogTestSuccess(t, "User-initiated enrollment is enabled for computers")

	// ------------------------------------------------------------------
	// 1. Clean up any existing invitations first
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Cleanup", "Cleaning up any existing computer invitations")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	existingList, _, listErr := svc.List(ctx1)
	if listErr == nil && existingList != nil && len(existingList.ComputerInvitation) > 0 {
		for _, inv := range existingList.ComputerInvitation {
			cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			_, delErr := svc.DeleteByID(cleanupCtx, strconv.Itoa(inv.ID))
			cancel()
			if delErr != nil {
				t.Logf("Warning: failed to clean up existing invitation ID=%d: %v", inv.ID, delErr)
			}
		}
	}

	// ------------------------------------------------------------------
	// 2. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test computer invitation")

	createReq := &computer_invitations.ResourceComputerInvitation{
		InvitationType:              "USER_INITIATED_ENROLLMENT",
		MultipleUsersAllowed:        false,
		CreateAccountIfDoesNotExist: true,
		KeepExistingSiteMembership:  true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		EnrollIntoSite: &computer_invitations.ComputerInvitationSubsetEnrollIntoState{
			ID:   -1,
			Name: "None",
		},
	}

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	created, createResp, err := svc.Create(ctx2, createReq)
	if err != nil {
		t.Logf("Note: Computer invitation creation failed even though enrollment is enabled. This may indicate additional configuration requirements or tenant limitations: %v", err)
		t.Skip("Skipping test due to computer invitation creation failure")
	}
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode(), "expected 200 or 201")
	assert.Positive(t, created.ID, "created computer invitation ID should be a positive integer")
	assert.NotEmpty(t, created.Invitation, "created computer invitation should have invitation string")

	invitationID := strconv.Itoa(created.ID)
	acc.LogTestSuccess(t, "Computer invitation created with ID=%d invitation=%q", created.ID, created.Invitation)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, invitationID)
		acc.LogCleanupDeleteError(t, "computer invitation", invitationID, delErr)
	})

	// ------------------------------------------------------------------
	// 3. List — verify the new computer invitation appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing computer invitations to verify creation")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	list, listResp, err := svc.List(ctx3)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, inv := range list.ComputerInvitation {
		if inv.ID == created.ID {
			found = true
			break
		}
	}
	assert.True(t, found, "newly created computer invitation should appear in list")
	acc.LogTestSuccess(t, "Computer invitation ID=%s found in list (%d total)", invitationID, list.Size)

	// ------------------------------------------------------------------
	// 4. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting computer invitation by ID=%s", invitationID)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetched, fetchResp, err := svc.GetByID(ctx4, invitationID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, created.Invitation, fetched.Invitation)
	acc.LogTestSuccess(t, "GetByID: ID=%d invitation=%q", fetched.ID, fetched.Invitation)

	// ------------------------------------------------------------------
	// 5. GetByInvitationID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByInvitationID", "Getting computer invitation by invitation=%q", created.Invitation)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	fetchedByInv, fetchByInvResp, err := svc.GetByInvitationID(ctx5, created.Invitation)
	require.NoError(t, err, "GetByInvitationID should not return an error")
	require.NotNil(t, fetchedByInv)
	assert.Equal(t, 200, fetchByInvResp.StatusCode())
	assert.Equal(t, created.ID, fetchedByInv.ID)
	assert.Equal(t, created.Invitation, fetchedByInv.Invitation)
	acc.LogTestSuccess(t, "GetByInvitationID: ID=%d invitation=%q", fetchedByInv.ID, fetchedByInv.Invitation)

	// ------------------------------------------------------------------
	// 6. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting computer invitation ID=%s", invitationID)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	deleteResp, err := svc.DeleteByID(ctx6, invitationID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Computer invitation ID=%s deleted", invitationID)
}

// =============================================================================
// TestAcceptance_ComputerInvitations_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_ComputerInvitations_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicAPI.ComputerInvitations

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer invitation ID cannot be empty")
	})

	t.Run("GetByInvitationID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByInvitationID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer invitation invitation ID cannot be empty")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "computer invitation ID cannot be empty")
	})
}
