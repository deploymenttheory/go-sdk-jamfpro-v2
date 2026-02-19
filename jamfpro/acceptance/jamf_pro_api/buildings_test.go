package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueBuildingName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_Buildings_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → Update → GetByID (verify) → History → Delete.
// =============================================================================

func TestAcceptance_Buildings_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Buildings
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test building")

	createReq := &buildings.RequestBuilding{
		Name:           uniqueBuildingName("acc-test-building"),
		StreetAddress1: "123 Test St",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78701",
		Country:        "United States",
	}
	created, createResp, err := svc.CreateBuildingV1(ctx, createReq)
	require.NoError(t, err, "CreateBuildingV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	buildingID := created.ID
	acc.LogTestSuccess(t, "Building created with ID=%s", buildingID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteBuildingByIDV1(cleanupCtx, buildingID)
		acc.LogCleanupDeleteError(t, "building", buildingID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing buildings to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListBuildingsV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, b := range list.Results {
		if b.ID == buildingID {
			found = true
			assert.Equal(t, createReq.Name, b.Name)
			break
		}
	}
	assert.True(t, found, "newly created building should appear in list")
	acc.LogTestSuccess(t, "Building ID=%s found in list (%d total)", buildingID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching building by ID=%s", buildingID)

	fetched, fetchResp, err := svc.GetBuildingByIDV1(ctx, buildingID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, buildingID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating building ID=%s", buildingID)

	updateReq := &buildings.RequestBuilding{
		Name:           uniqueBuildingName("acc-test-building-updated"),
		StreetAddress1: "456 Updated Ave",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78702",
		Country:        "United States",
	}
	updated, updateResp, err := svc.UpdateBuildingByIDV1(ctx, buildingID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updateReq.Name, updated.Name)
	acc.LogTestSuccess(t, "Building updated: ID=%s", buildingID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetBuildingByIDV1(ctx, buildingID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 5a. Add history note and fetch history
	acc.LogTestStage(t, "History", "Adding history note and fetching history for ID=%s", buildingID)

	noteReq := &buildings.AddHistoryNotesRequest{
		Note: fmt.Sprintf("Acceptance test note at %s", time.Now().Format(time.RFC3339)),
	}
	noteResp, err := svc.AddBuildingHistoryNotesV1(ctx, buildingID, noteReq)
	require.NoError(t, err)
	require.NotNil(t, noteResp)
	assert.Equal(t, 201, noteResp.StatusCode)
	acc.LogTestSuccess(t, "History note added")

	history, histResp, err := svc.GetBuildingHistoryV1(ctx, buildingID, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 1)
	acc.LogTestSuccess(t, "History entries: %d", history.TotalCount)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting building ID=%s", buildingID)

	deleteResp, err := svc.DeleteBuildingByIDV1(ctx, buildingID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Building ID=%s deleted", buildingID)
}

// =============================================================================
// TestAcceptance_Buildings_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_Buildings_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Buildings
	ctx := context.Background()

	name := uniqueBuildingName("acc-rsql-building")
	createReq := &buildings.RequestBuilding{
		Name:  name,
		City:  "Austin",
		Country: "United States",
	}

	created, _, err := svc.CreateBuildingV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	buildingID := created.ID
	acc.LogTestSuccess(t, "Created building ID=%s name=%q", buildingID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteBuildingByIDV1(cleanupCtx, buildingID)
		acc.LogCleanupDeleteError(t, "building", buildingID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListBuildingsV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, b := range list.Results {
		if b.ID == buildingID {
			found = true
			assert.Equal(t, name, b.Name)
			break
		}
	}
	assert.True(t, found, "building should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target building found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_Buildings_ValidationErrors
// =============================================================================

func TestAcceptance_Buildings_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Buildings

	t.Run("GetBuildingByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetBuildingByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "building ID is required")
	})

	t.Run("CreateBuildingV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateBuildingV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateBuildingByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateBuildingByIDV1(context.Background(), "", &buildings.RequestBuilding{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteBuildingByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteBuildingByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "building ID is required")
	})
}
