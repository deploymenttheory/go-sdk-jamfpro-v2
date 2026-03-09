package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/static_mobile_device_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// Acceptance Tests: Static Mobile Device Groups
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • List(ctx, rsqlQuery) - Lists all static mobile device groups
//   • GetByID(ctx, id) - Retrieves a static mobile device group by ID
//   • Create(ctx, request) - Creates a new static mobile device group
//   • UpdateByID(ctx, id, request) - Updates an existing static mobile device group
//   • DeleteByID(ctx, id) - Deletes a static mobile device group by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → Delete
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: Empty IDs, nil requests
//
// =============================================================================

func TestAcceptance_StaticMobileDeviceGroups_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.StaticMobileDeviceGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test static mobile device group")

	createReq := &static_mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        acc.UniqueName("sdkv2_acc_acc-test-static-md"),
		Description: "Acceptance test static mobile device group",
		SiteID:      "-1",
		Assignments: []static_mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	}
	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && (createResp.StatusCode() == 500 || createResp.StatusCode() == 404) {
		t.Skip("Static mobile device group endpoint not available (404/500)")
	}
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode())
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Static mobile device group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static mobile device group", groupID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing static mobile device groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.List(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, createReq.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "newly created static mobile device group should appear in list")
	acc.LogTestSuccess(t, "Static mobile device group ID=%s found in list (%d total)", groupID, list.TotalCount)

	// 3. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting static mobile device group by ID=%s", groupID)

	var fetched *static_mobile_device_groups.ResourceStaticMobileDeviceGroup
	var fetchResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByID(ctx, groupID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating static mobile device group ID=%s", groupID)

	updateReq := &static_mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        acc.UniqueName("sdkv2_acc_acc-test-static-md-updated"),
		Description: "Updated description",
		SiteID:      "-1",
		Assignments: []static_mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	}
	updated, updateResp, err := svc.UpdateByID(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode())
	acc.LogTestSuccess(t, "Static mobile device group updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetByID(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	assert.Equal(t, updateReq.Description, fetched2.Description)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting static mobile device group ID=%s", groupID)

	deleteResp, err := svc.DeleteByID(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Static mobile device group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_StaticMobileDeviceGroups_list_with_rsql_filter
// =============================================================================

func TestAcceptance_StaticMobileDeviceGroups_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.StaticMobileDeviceGroups
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-static-md")
	createReq := &static_mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        name,
		Description: "Acceptance test RSQL filter static mobile device group",
		SiteID:      "-1",
		Assignments: []static_mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	}

	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && (createResp.StatusCode() == 500 || createResp.StatusCode() == 404) {
		t.Skip("Static mobile device group endpoint not available (404/500)")
	}
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created static mobile device group ID=%s name=%q", groupID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static mobile device group", groupID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.List(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, name, g.Name)
			break
		}
	}
	assert.True(t, found, "static mobile device group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target group found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_StaticMobileDeviceGroups_validation_errors
// =============================================================================

func TestAcceptance_StaticMobileDeviceGroups_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.StaticMobileDeviceGroups

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static mobile device group ID is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &static_mobile_device_groups.RequestStaticMobileDeviceGroup{
			Name:   "x",
			SiteID: "-1",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByID_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static mobile device group ID is required")
	})
}
