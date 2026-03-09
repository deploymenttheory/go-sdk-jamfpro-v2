package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_mobile_device_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// TestAcceptance_SmartMobileDeviceGroups_lifecycle exercises Create → GetByID →
// Update → GetByID (verify) → Delete.
// =============================================================================

func TestAcceptance_SmartMobileDeviceGroups_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test smart mobile device group")

	groupName := acc.UniqueName("sdkv2_acc_acc-test-smart-md")
	createReq := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        groupName,
		GroupDescription: "Acceptance test smart mobile device group",
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{Name: "Last Inventory Update", Priority: 1, AndOr: "and", SearchType: "more than x days ago", Value: "365"},
		},
	}

	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && (createResp.StatusCode() == 500 || createResp.StatusCode() == 404) {
		t.Skip("Smart mobile device group endpoint not available (404/500)")
	}
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode())
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Smart mobile device group created with ID=%s name=%q", groupID, groupName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart mobile device group", groupID, delErr)
	})

	// 2. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting smart mobile device group by ID=%s", groupID)

	var fetched *smart_mobile_device_groups.ResourceSmartMobileDeviceGroup
	var fetchResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByID(ctx, groupID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, groupID, fetched.GroupID)
	assert.Equal(t, groupName, fetched.GroupName)
	acc.LogTestSuccess(t, "GetByID: GroupID=%s GroupName=%q", fetched.GroupID, fetched.GroupName)

	// 3. Update
	acc.LogTestStage(t, "Update", "Updating smart mobile device group ID=%s", groupID)

	updatedName := acc.UniqueName("sdkv2_acc_acc-test-smart-md-updated")
	updateReq := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        updatedName,
		GroupDescription: "Updated description",
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{Name: "Last Inventory Update", Priority: 1, AndOr: "and", SearchType: "more than x days ago", Value: "365"},
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode())
	acc.LogTestSuccess(t, "Smart mobile device group updated: ID=%s", groupID)

	// 4. Re-fetch to verify
	fetched2, _, err := svc.GetByID(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updatedName, fetched2.GroupName)
	assert.Equal(t, updateReq.GroupDescription, fetched2.GroupDescription)
	acc.LogTestSuccess(t, "Update verified: GroupName=%q", fetched2.GroupName)

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting smart mobile device group ID=%s", groupID)

	deleteResp, err := svc.DeleteByID(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Smart mobile device group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_SmartMobileDeviceGroups_list_with_rsql_filter
// =============================================================================

func TestAcceptance_SmartMobileDeviceGroups_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups
	ctx := context.Background()

	groupName := acc.UniqueName("sdkv2_acc_rsql-smart-md")
	createReq := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        groupName,
		GroupDescription: "Acceptance test RSQL filter smart mobile device group",
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{Name: "Last Inventory Update", Priority: 1, AndOr: "and", SearchType: "more than x days ago", Value: "365"},
		},
	}

	created, createResp, err := svc.Create(ctx, createReq)
	if err != nil && createResp != nil && (createResp.StatusCode() == 500 || createResp.StatusCode() == 404) {
		t.Skip("Smart mobile device group endpoint not available (404/500)")
	}
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created smart mobile device group ID=%s name=%q", groupID, groupName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart mobile device group", groupID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`groupName=="%s"`, groupName),
	}

	list, listResp, err := svc.List(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, g := range list.Results {
		if g.GroupID == groupID {
			found = true
			assert.Equal(t, groupName, g.GroupName)
			break
		}
	}
	assert.True(t, found, "smart mobile device group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target group found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_SmartMobileDeviceGroups_validation_errors
// =============================================================================

func TestAcceptance_SmartMobileDeviceGroups_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart mobile device group ID is required")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart mobile device group name is required")
	})

	t.Run("GetMembership_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetMembership(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart mobile device group ID is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
			GroupName: "x",
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
		assert.Contains(t, err.Error(), "smart mobile device group ID is required")
	})
}

// TestAcceptance_SmartMobileDeviceGroups_list verifies listing smart mobile device groups.
func TestAcceptance_SmartMobileDeviceGroups_list(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "100"})
	if err != nil && resp != nil && resp.StatusCode() == 404 {
		t.Skip("Smart mobile device group endpoint not available (404)")
	}
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_SmartMobileDeviceGroups_get_by_id fetches a smart group by ID when at least one exists.
func TestAcceptance_SmartMobileDeviceGroups_get_by_id(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "1"})
	if err != nil && resp != nil && resp.StatusCode() == 404 {
		t.Skip("Smart mobile device group endpoint not available (404)")
	}
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping GetByID")
	}

	got, resp, err := svc.GetByID(ctx, list.Results[0].GroupID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, list.Results[0].GroupID, got.GroupID)
}

// TestAcceptance_SmartMobileDeviceGroups_get_membership fetches membership when at least one group exists.
func TestAcceptance_SmartMobileDeviceGroups_get_membership(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.SmartMobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "1"})
	if err != nil && resp != nil && resp.StatusCode() == 404 {
		t.Skip("Smart mobile device group endpoint not available (404)")
	}
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping GetMembership")
	}

	membership, resp, err := svc.GetMembership(ctx, list.Results[0].GroupID, map[string]string{"page": "0", "page-size": "10"})
	require.NoError(t, err)
	require.NotNil(t, membership)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, membership.TotalCount, 0)
}
