package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_ios"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// Acceptance Tests: Self-Service Branding Mobile (iOS)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists all self-service branding mobile configurations
//   • GetByIDV1(ctx, id) - Retrieves a branding configuration by ID
//   • GetByNameV1(ctx, name) - Retrieves a branding configuration by name
//   • CreateV1(ctx, request) - Creates a new branding configuration
//   • UpdateByIDV1(ctx, id, request) - Updates an existing branding configuration by ID
//   • UpdateByNameV1(ctx, name, request) - Updates a branding configuration by name
//   • DeleteByIDV1(ctx, id) - Deletes a branding configuration by ID
//   • DeleteByNameV1(ctx, name) - Deletes a branding configuration by name
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Flow: Create → GetByID → Update → Verify → Delete
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: Empty IDs, nil requests, empty name
//
// =============================================================================

func TestAcceptance_SelfServiceBrandingIOS_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingIOS
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test self-service branding mobile")

	createReq := &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
		BrandingName:              acc.UniqueName("sdkv2_acc_acc-test-ssb-mobile"),
		HeaderBackgroundColorCode: "FFFFFF",
		MenuIconColorCode:         "000000",
		BrandingNameColorCode:     "333333",
		StatusBarTextColor:        "light",
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	brandingID := created.ID
	acc.LogTestSuccess(t, "Self-service branding mobile created with ID=%s", brandingID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, brandingID)
		acc.LogCleanupDeleteError(t, "self-service branding mobile", brandingID, delErr)
	})

	// 2. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting branding by ID=%s", brandingID)

	var fetched *self_service_branding_ios.ResourceSelfServiceBrandingMobile
	var fetchResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByIDV1(ctx, brandingID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, brandingID, fetched.ID)
	assert.Equal(t, createReq.BrandingName, fetched.BrandingName)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.BrandingName)

	// 3. GetByName
	fetchedByName, _, err := svc.GetByNameV1(ctx, createReq.BrandingName)
	require.NoError(t, err)
	require.NotNil(t, fetchedByName)
	assert.Equal(t, brandingID, fetchedByName.ID)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating branding ID=%s", brandingID)

	updateReq := &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
		BrandingName:              acc.UniqueName("sdkv2_acc_acc-test-ssb-mobile-updated"),
		HeaderBackgroundColorCode: "F0F0F0",
		MenuIconColorCode:         "0066CC",
		BrandingNameColorCode:     "222222",
		StatusBarTextColor:        "dark",
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, brandingID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Branding updated: ID=%s", brandingID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV1(ctx, brandingID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.BrandingName, fetched2.BrandingName)
	assert.Equal(t, updateReq.HeaderBackgroundColorCode, fetched2.HeaderBackgroundColorCode)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.BrandingName)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting branding ID=%s", brandingID)

	deleteResp, err := svc.DeleteByIDV1(ctx, brandingID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Self-service branding mobile ID=%s deleted", brandingID)
}

// =============================================================================
// TestAcceptance_SelfServiceBrandingIOS_list_with_rsql_filter
// =============================================================================

func TestAcceptance_SelfServiceBrandingIOS_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingIOS
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-ssb-mobile")
	createReq := &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
		BrandingName:              name,
		HeaderBackgroundColorCode: "FFFFFF",
		MenuIconColorCode:         "000000",
		BrandingNameColorCode:     "333333",
		StatusBarTextColor:        "light",
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	brandingID := created.ID
	acc.LogTestSuccess(t, "Created self-service branding mobile ID=%s name=%q", brandingID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, brandingID)
		acc.LogCleanupDeleteError(t, "self-service branding mobile", brandingID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`brandingName=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, b := range list.Results {
		if b.ID == brandingID {
			found = true
			assert.Equal(t, name, b.BrandingName)
			break
		}
	}
	assert.True(t, found, "self-service branding mobile should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target branding found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_SelfServiceBrandingIOS_validation_errors
// =============================================================================

func TestAcceptance_SelfServiceBrandingIOS_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingIOS

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding mobile ID is required")
	})

	t.Run("GetByNameV1_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByNameV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding mobile name is required")
	})

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
			BrandingName:              "x",
			HeaderBackgroundColorCode: "#FFF",
			MenuIconColorCode:         "#000",
			BrandingNameColorCode:     "#333",
			StatusBarTextColor:        "light",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding mobile ID is required")
	})
}
