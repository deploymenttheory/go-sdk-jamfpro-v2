package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_NetworkSegments_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_NetworkSegments_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.NetworkSegments
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test network segment")

	segmentName := uniqueName("acc-test-netseg")
	createReq := &network_segments.RequestNetworkSegment{
		Name:            segmentName,
		StartingAddress: "192.168.100.0",
		EndingAddress:   "192.168.100.255",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateNetworkSegment(ctx1, createReq)
	require.NoError(t, err, "CreateNetworkSegment should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created network segment ID should be a positive integer")

	segmentID := created.ID
	acc.LogTestSuccess(t, "Network segment created with ID=%d name=%q", segmentID, segmentName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteNetworkSegmentByID(cleanupCtx, segmentID)
		acc.LogCleanupDeleteError(t, "network segment", fmt.Sprintf("%d", segmentID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new segment appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing network segments to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListNetworkSegments(ctx2)
	require.NoError(t, err, "ListNetworkSegments should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == segmentID {
			found = true
			assert.Equal(t, segmentName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created network segment should appear in list")
	acc.LogTestSuccess(t, "Network segment ID=%d found in list (%d total)", segmentID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching network segment by ID=%d", segmentID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetNetworkSegmentByID(ctx3, segmentID)
	require.NoError(t, err, "GetNetworkSegmentByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, segmentID, fetched.ID)
	assert.Equal(t, segmentName, fetched.Name)
	assert.Equal(t, "192.168.100.0", fetched.StartingAddress)
	assert.Equal(t, "192.168.100.255", fetched.EndingAddress)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching network segment by name=%q", segmentName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetNetworkSegmentByName(ctx4, segmentName)
	require.NoError(t, err, "GetNetworkSegmentByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, segmentID, fetchedByName.ID)
	assert.Equal(t, segmentName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-netseg-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating network segment ID=%d to name=%q", segmentID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &network_segments.RequestNetworkSegment{
		Name:            updatedName,
		StartingAddress: "192.168.100.0",
		EndingAddress:   "192.168.100.255",
	}
	updated, updateResp, err := svc.UpdateNetworkSegmentByID(ctx5, segmentID, updateReq)
	require.NoError(t, err, "UpdateNetworkSegmentByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating network segment name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &network_segments.RequestNetworkSegment{
		Name:            segmentName,
		StartingAddress: "192.168.100.0",
		EndingAddress:   "192.168.100.255",
	}
	reverted, revertResp, err := svc.UpdateNetworkSegmentByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateNetworkSegmentByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetNetworkSegmentByID(ctx7, segmentID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, segmentName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting network segment ID=%d", segmentID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteNetworkSegmentByID(ctx8, segmentID)
	require.NoError(t, err, "DeleteNetworkSegmentByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Network segment ID=%d deleted", segmentID)
}

// =============================================================================
// TestAcceptance_NetworkSegments_DeleteByName creates a segment then deletes by name.
// =============================================================================

func TestAcceptance_NetworkSegments_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.NetworkSegments
	ctx := context.Background()

	segmentName := uniqueName("acc-test-netseg-dbn")
	createReq := &network_segments.RequestNetworkSegment{
		Name:            segmentName,
		StartingAddress: "172.16.50.0",
		EndingAddress:   "172.16.50.255",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateNetworkSegment(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	segmentID := created.ID
	acc.LogTestSuccess(t, "Created network segment ID=%d name=%q for delete-by-name test", segmentID, segmentName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteNetworkSegmentByID(cleanupCtx, segmentID)
		acc.LogCleanupDeleteError(t, "network segment", fmt.Sprintf("%d", segmentID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteNetworkSegmentByName(ctx2, segmentName)
	require.NoError(t, err, "DeleteNetworkSegmentByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Network segment %q deleted by name", segmentName)
}

// =============================================================================
// TestAcceptance_NetworkSegments_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_NetworkSegments_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.NetworkSegments

	t.Run("GetNetworkSegmentByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetNetworkSegmentByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
	})

	t.Run("GetNetworkSegmentByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetNetworkSegmentByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment name is required")
	})

	t.Run("CreateNetworkSegment_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateNetworkSegment(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateNetworkSegmentByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateNetworkSegmentByID(context.Background(), 0, &network_segments.RequestNetworkSegment{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
	})

	t.Run("UpdateNetworkSegmentByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateNetworkSegmentByName(context.Background(), "", &network_segments.RequestNetworkSegment{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment name is required")
	})

	t.Run("DeleteNetworkSegmentByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteNetworkSegmentByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment ID must be a positive integer")
	})

	t.Run("DeleteNetworkSegmentByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteNetworkSegmentByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "network segment name is required")
	})
}
