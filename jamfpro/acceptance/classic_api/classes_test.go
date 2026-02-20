package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/classes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Classes_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Classes_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Classes
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test class")

	className := uniqueName("acc-test-class")
	createReq := &classes.RequestClass{
		Name:        className,
		Description: "Acceptance test class",
		Students: []classes.Student{
			{Student: "student1@example.com"},
			{Student: "student2@example.com"},
		},
		Teachers: []classes.Teacher{
			{Teacher: "teacher1@example.com"},
		},
		MeetingTimes: &classes.MeetingTimesContainer{
			MeetingTime: classes.MeetingTime{
				Days:      "MWF",
				StartTime: 900,
				EndTime:   1000,
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateClass(ctx1, createReq)
	require.NoError(t, err, "CreateClass should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created class ID should be a positive integer")

	classID := created.ID
	acc.LogTestSuccess(t, "Class created with ID=%d name=%q", classID, className)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteClassByID(cleanupCtx, classID)
		acc.LogCleanupDeleteError(t, "class", fmt.Sprintf("%d", classID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new class appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing classes to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListClasses(ctx2)
	require.NoError(t, err, "ListClasses should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, c := range list.Results {
		if c.ID == classID {
			found = true
			assert.Equal(t, className, c.Name)
			break
		}
	}
	assert.True(t, found, "newly created class should appear in list")
	acc.LogTestSuccess(t, "Class ID=%d found in list (%d total)", classID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching class by ID=%d", classID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetClassByID(ctx3, classID)
	require.NoError(t, err, "GetClassByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, classID, fetched.ID)
	assert.Equal(t, className, fetched.Name)
	assert.Equal(t, "Acceptance test class", fetched.Description)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching class by name=%q", className)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetClassByName(ctx4, className)
	require.NoError(t, err, "GetClassByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, classID, fetchedByName.ID)
	assert.Equal(t, className, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-class-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating class ID=%d to name=%q", classID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &classes.RequestClass{
		Name:        updatedName,
		Description: "Updated description",
		Students: []classes.Student{
			{Student: "student3@example.com"},
		},
	}
	updated, updateResp, err := svc.UpdateClassByID(ctx5, classID, updateReq)
	require.NoError(t, err, "UpdateClassByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating class name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &classes.RequestClass{
		Name:        className,
		Description: "Reverted description",
	}
	reverted, revertResp, err := svc.UpdateClassByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateClassByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetClassByID(ctx7, classID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, className, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting class ID=%d", classID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteClassByID(ctx8, classID)
	require.NoError(t, err, "DeleteClassByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Class ID=%d deleted", classID)
}

// =============================================================================
// TestAcceptance_Classes_DeleteByName creates a class then deletes by name.
// =============================================================================

func TestAcceptance_Classes_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Classes
	ctx := context.Background()

	className := uniqueName("acc-test-class-del")
	createReq := &classes.RequestClass{
		Name:        className,
		Description: "Test class for delete by name",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateClass(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	classID := created.ID
	acc.LogTestSuccess(t, "Created class ID=%d name=%q for delete-by-name test", classID, className)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteClassByID(cleanupCtx, classID)
		acc.LogCleanupDeleteError(t, "class", fmt.Sprintf("%d", classID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteClassByName(ctx2, className)
	require.NoError(t, err, "DeleteClassByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Class %q deleted by name", className)
}

// =============================================================================
// TestAcceptance_Classes_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Classes_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Classes

	t.Run("GetClassByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetClassByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class ID must be a positive integer")
	})

	t.Run("GetClassByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetClassByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class name is required")
	})

	t.Run("CreateClass_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateClass(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateClassByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateClassByID(context.Background(), 0, &classes.RequestClass{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class ID must be a positive integer")
	})

	t.Run("UpdateClassByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateClassByName(context.Background(), "", &classes.RequestClass{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class name is required")
	})

	t.Run("DeleteClassByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteClassByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class ID must be a positive integer")
	})

	t.Run("DeleteClassByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteClassByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "class name is required")
	})
}
