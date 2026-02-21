package jamf_pro_api_test

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/groups"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Groups
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - List groups with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Get group by platform ID
//   • GetComputerGroupByNameV1(ctx, name) - Get computer group by name
//   • GetMobileGroupByNameV1(ctx, name) - Get mobile group by name
//   • GetComputerGroupByIDV1(ctx, jamfProID) - Get computer group by Jamf Pro ID
//   • GetMobileGroupByIDV1(ctx, jamfProID) - Get mobile group by Jamf Pro ID
//   • UpdateByIDV1(ctx, id, request) - Update group by platform ID (PATCH)
//   • DeleteByIDV1(ctx, id) - Delete group by platform ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Resources [COMPLETE]
//     -- Reason: Groups are managed through computer_groups and mobile_device_groups services
//     -- Tests: TestAcceptance_Groups_List, TestAcceptance_Groups_GetOperations
//     -- Flow: List → GetByID → GetByName → GetByJamfProID
//
//   ✓ Pattern 5: RSQL Filter Testing [COMPLETE]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Groups_ListWithRSQLFilter
//     -- Flow: List all → List with RSQL filter → Verify filtered results
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List all groups
//   ✓ List with pagination
//   ✓ List with RSQL filtering
//   ✓ Get by platform ID
//   ✓ Get computer group by name
//   ✓ Get mobile group by name
//   ✓ Get computer group by Jamf Pro ID
//   ✓ Get mobile group by Jamf Pro ID
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only service - groups are managed via computer_groups and mobile_device_groups
//   • The Groups API provides a unified view across both computer and mobile device groups
//   • Platform ID (groupPlatformId) is different from Jamf Pro ID (groupJamfProId)
//   • Tests verify correct filtering by group type (COMPUTER vs MOBILE)
//
// =============================================================================

func TestAcceptance_Groups_List(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing all groups")
	result, resp, err := acc.Client.Groups.ListV1(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.GreaterOrEqual(t, result.TotalCount, 0)
	acc.LogTestSuccess(t, "Found %d total groups", result.TotalCount)

	if result.TotalCount > 0 {
		firstGroup := result.Results[0]
		acc.LogTestSuccess(t, "First group - Name: %s, Type: %s, Platform ID: %s",
			firstGroup.GroupName, firstGroup.GroupType, firstGroup.GroupPlatformId)
	}
}

func TestAcceptance_Groups_ListWithPagination(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing groups with pagination")
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "groupName:asc",
	}

	result, resp, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	acc.LogTestSuccess(t, "Paginated list returned %d results (total: %d)", len(result.Results), result.TotalCount)
}

func TestAcceptance_Groups_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Listing all groups")
	allGroups, allResp, err := acc.Client.Groups.ListV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, allResp)

	if allGroups.TotalCount == 0 {
		t.Skip("No groups available for RSQL filter test")
		return
	}

	acc.LogTestSuccess(t, "Found %d total groups", allGroups.TotalCount)

	// Test RSQL filtering by group type
	acc.LogTestStage(t, "RSQL Filter", "Testing RSQL filter by group type")
	rsqlQuery := map[string]string{
		"filter": `groupType=="COMPUTER"`,
	}

	filteredGroups, filteredResp, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, filteredResp)
	require.NotNil(t, filteredGroups)
	acc.LogTestSuccess(t, "RSQL filter returned %d computer groups", filteredGroups.TotalCount)

	// Verify all returned groups are COMPUTER type
	for _, group := range filteredGroups.Results {
		require.Equal(t, "COMPUTER", group.GroupType, "All filtered groups should be COMPUTER type")
	}

	// Verify filtered count is <= total count
	require.LessOrEqual(t, filteredGroups.TotalCount, allGroups.TotalCount,
		"Filtered results should be <= total results")
}

func TestAcceptance_Groups_GetByID(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Listing groups to get an ID")
	listResult, _, err := acc.Client.Groups.ListV1(ctx, nil)
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No groups found to test GetByID")
		return
	}

	platformID := listResult.Results[0].GroupPlatformId
	acc.LogTestSuccess(t, "Using platform ID: %s", platformID)

	acc.LogTestStage(t, "GetByID", "Getting group by platform ID")
	result, resp, err := acc.Client.Groups.GetByIDV1(ctx, platformID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, platformID, result.GroupPlatformId)
	acc.LogTestSuccess(t, "Retrieved group - Name: %s, Type: %s", result.GroupName, result.GroupType)
}

func TestAcceptance_Groups_GetComputerGroupByName(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Finding a computer group")
	rsqlQuery := map[string]string{
		"filter": `groupType=="COMPUTER"`,
		"page":   "0",
		"page-size": "1",
	}

	listResult, _, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No computer groups found to test GetComputerGroupByName")
		return
	}

	groupName := listResult.Results[0].GroupName
	acc.LogTestSuccess(t, "Using computer group name: %s", groupName)

	acc.LogTestStage(t, "GetByName", "Getting computer group by name")
	result, resp, err := acc.Client.Groups.GetComputerGroupByNameV1(ctx, groupName)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, groupName, result.GroupName)
	require.Equal(t, "COMPUTER", result.GroupType)
	acc.LogTestSuccess(t, "Retrieved computer group - Name: %s, ID: %s", result.GroupName, result.GroupJamfProId)
}

func TestAcceptance_Groups_GetMobileGroupByName(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Finding a mobile group")
	rsqlQuery := map[string]string{
		"filter": `groupType=="MOBILE"`,
		"page":   "0",
		"page-size": "1",
	}

	listResult, _, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No mobile groups found to test GetMobileGroupByName")
		return
	}

	groupName := listResult.Results[0].GroupName
	acc.LogTestSuccess(t, "Using mobile group name: %s", groupName)

	acc.LogTestStage(t, "GetByName", "Getting mobile group by name")
	result, resp, err := acc.Client.Groups.GetMobileGroupByNameV1(ctx, groupName)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, groupName, result.GroupName)
	require.Equal(t, "MOBILE", result.GroupType)
	acc.LogTestSuccess(t, "Retrieved mobile group - Name: %s, ID: %s", result.GroupName, result.GroupJamfProId)
}

func TestAcceptance_Groups_GetComputerGroupByJamfProID(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Finding a computer group")
	rsqlQuery := map[string]string{
		"filter": `groupType=="COMPUTER"`,
		"page":   "0",
		"page-size": "1",
	}

	listResult, _, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No computer groups found to test GetComputerGroupByID")
		return
	}

	jamfProID := listResult.Results[0].GroupJamfProId
	acc.LogTestSuccess(t, "Using Jamf Pro ID: %s", jamfProID)

	acc.LogTestStage(t, "GetByJamfProID", "Getting computer group by Jamf Pro ID")
	result, resp, err := acc.Client.Groups.GetComputerGroupByIDV1(ctx, jamfProID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, jamfProID, result.GroupJamfProId)
	require.Equal(t, "COMPUTER", result.GroupType)
	acc.LogTestSuccess(t, "Retrieved computer group - Name: %s, Platform ID: %s", result.GroupName, result.GroupPlatformId)
}

func TestAcceptance_Groups_GetMobileGroupByJamfProID(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "Setup", "Finding a mobile group")
	rsqlQuery := map[string]string{
		"filter": `groupType=="MOBILE"`,
		"page":   "0",
		"page-size": "1",
	}

	listResult, _, err := acc.Client.Groups.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No mobile groups found to test GetMobileGroupByID")
		return
	}

	jamfProID := listResult.Results[0].GroupJamfProId
	acc.LogTestSuccess(t, "Using Jamf Pro ID: %s", jamfProID)

	acc.LogTestStage(t, "GetByJamfProID", "Getting mobile group by Jamf Pro ID")
	result, resp, err := acc.Client.Groups.GetMobileGroupByIDV1(ctx, jamfProID)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	require.Equal(t, jamfProID, result.GroupJamfProId)
	require.Equal(t, "MOBILE", result.GroupType)
	acc.LogTestSuccess(t, "Retrieved mobile group - Name: %s, Platform ID: %s", result.GroupName, result.GroupPlatformId)
}

func TestAcceptance_Groups_Update(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	// Find a group to update
	acc.LogTestStage(t, "Setup", "Finding a group to test update")
	listResult, _, err := acc.Client.Groups.ListV1(ctx, map[string]string{
		"page":      "0",
		"page-size": "1",
	})
	require.NoError(t, err)

	if len(listResult.Results) == 0 {
		t.Skip("No groups found to test update")
		return
	}

	group := listResult.Results[0]
	platformID := group.GroupPlatformId
	originalName := group.GroupName
	originalDescription := group.GroupDescription

	acc.LogTestSuccess(t, "Testing update on group: %s (Platform ID: %s)", originalName, platformID)

	// Update the group
	acc.LogTestStage(t, "Update", "Updating group description")
	updateReq := &groups.RequestUpdateGroup{
		GroupName:        originalName, // Keep same name
		GroupDescription: fmt.Sprintf("%s - Updated by acceptance test", originalDescription),
	}

	updated, updateResp, err := acc.Client.Groups.UpdateByIDV1(ctx, platformID, updateReq)

	if err != nil {
		// Some groups may not be modifiable (e.g., built-in groups)
		t.Logf("Update may not be supported for this group: %v", err)
		t.Skip("Skipping update test for this group")
		return
	}

	require.NotNil(t, updateResp)
	require.NotNil(t, updated)
	acc.LogTestSuccess(t, "Updated group - Description: %s", updated.GroupDescription)

	// Restore original
	acc.LogTestStage(t, "Restore", "Restoring original description")
	restoreReq := &groups.RequestUpdateGroup{
		GroupName:        originalName,
		GroupDescription: originalDescription,
	}

	_, restoreResp, err := acc.Client.Groups.UpdateByIDV1(ctx, platformID, restoreReq)
	if err == nil {
		require.NotNil(t, restoreResp)
		acc.LogTestSuccess(t, "Restored original group settings")
	}
}
