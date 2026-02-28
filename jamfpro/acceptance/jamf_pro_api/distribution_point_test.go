package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/distribution_point"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_DistributionPoint_lifecycle
// =============================================================================

func TestAcceptance_DistributionPoint_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DistributionPoint
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test distribution point")

	name := acc.UniqueName("sdkv2_acc_distpoint")
	createReq := &distribution_point.RequestDistributionPoint{
		Name:                      name,
		ServerName:                "acc-test-server.example.com",
		FileSharingConnectionType: "NONE",
	}

	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	dpID := created.ID
	acc.LogTestSuccess(t, "Distribution point created with ID=%s name=%q", dpID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, dpID)
		acc.LogCleanupDeleteError(t, "distribution point", dpID, delErr)
	})

	// 2. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching distribution point by ID=%s", dpID)

	fetched, fetchResp, err := svc.GetByIDV1(ctx, dpID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, dpID, fetched.ID)
	assert.Equal(t, name, fetched.Name)
	acc.LogTestSuccess(t, "GetByIDV1: ID=%s name=%q", fetched.ID, fetched.Name)

	// 3. Update
	acc.LogTestStage(t, "Update", "Updating distribution point ID=%s", dpID)

	updatedName := acc.UniqueName("sdkv2_acc_distpoint-updated")
	updateReq := &distribution_point.RequestDistributionPoint{
		Name:                      updatedName,
		ServerName:                "acc-test-server-updated.example.com",
		FileSharingConnectionType: "NONE",
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, dpID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Distribution point updated: ID=%s", dpID)

	// 4. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV1(ctx, dpID)
	require.NoError(t, err)
	assert.Equal(t, updatedName, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting distribution point ID=%s", dpID)

	deleteResp, err := svc.DeleteByIDV1(ctx, dpID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Distribution point ID=%s deleted", dpID)
}

// =============================================================================
// TestAcceptance_DistributionPoint_list_with_rsql_filter
// =============================================================================

func TestAcceptance_DistributionPoint_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DistributionPoint
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-distpoint")
	createReq := &distribution_point.RequestDistributionPoint{
		Name:                      name,
		ServerName:                "acc-test-rsql-server.example.com",
		FileSharingConnectionType: "NONE",
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	dpID := created.ID
	acc.LogTestSuccess(t, "Created distribution point ID=%s name=%q", dpID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, dpID)
		acc.LogCleanupDeleteError(t, "distribution point", dpID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, dp := range list.Results {
		if dp.ID == dpID {
			found = true
			assert.Equal(t, name, dp.Name)
			break
		}
	}
	assert.True(t, found, "distribution point should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_DistributionPoint_bulk_delete
// =============================================================================

func TestAcceptance_DistributionPoint_bulk_delete(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DistributionPoint
	ctx := context.Background()

	dp1, _, err := svc.CreateV1(ctx, &distribution_point.RequestDistributionPoint{
		Name:                      acc.UniqueName("sdkv2_acc_bulk-del-dp-1"),
		ServerName:                "acc-bulk-del-dp-1.example.com",
		FileSharingConnectionType: "NONE",
	})
	require.NoError(t, err)
	require.NotNil(t, dp1)

	dp2, _, err := svc.CreateV1(ctx, &distribution_point.RequestDistributionPoint{
		Name:                      acc.UniqueName("sdkv2_acc_bulk-del-dp-2"),
		ServerName:                "acc-bulk-del-dp-2.example.com",
		FileSharingConnectionType: "NONE",
	})
	require.NoError(t, err)
	require.NotNil(t, dp2)

	acc.LogTestSuccess(t, "Created distribution points ID=%s and ID=%s", dp1.ID, dp2.ID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, dp1.ID)
		_, _ = svc.DeleteByIDV1(cleanupCtx, dp2.ID)
	})

	deleteResp, err := svc.DeleteMultipleV1(ctx, []string{dp1.ID, dp2.ID})
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Bulk delete completed, status=%d", deleteResp.StatusCode)
}

// =============================================================================
// TestAcceptance_DistributionPoint_validation_errors
// =============================================================================

func TestAcceptance_DistributionPoint_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DistributionPoint

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &distribution_point.RequestDistributionPoint{
			Name:                      "x",
			ServerName:                "x.example.com",
			FileSharingConnectionType: "NONE",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID is required")
	})

	t.Run("UpdateByIDV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "distribution point ID is required")
	})

	t.Run("DeleteMultipleV1_EmptyList", func(t *testing.T) {
		_, err := svc.DeleteMultipleV1(context.Background(), []string{})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "at least one ID is required")
	})
}
