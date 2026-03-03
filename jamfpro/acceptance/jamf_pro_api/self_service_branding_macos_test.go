package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_macos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Self Service Branding macOS
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • List(ctx, rsqlQuery) - Retrieves all self-service branding configurations
//   • GetByID(ctx, id) - Retrieves a branding configuration by ID
//   • GetByName(ctx, name) - Retrieves a branding configuration by name
//   • Create(ctx, request) - Creates a new branding configuration
//   • UpdateByID(ctx, id, request) - Updates an existing branding configuration
//   • UpdateByName(ctx, name, request) - Updates by name
//   • DeleteByID(ctx, id) - Deletes a branding configuration by ID
//   • DeleteByName(ctx, name) - Deletes by name
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Flow: Create → GetByID → Update → GetByID (verify) → Delete
//
//   ✓ Pattern 7: Validation Errors
//     -- Cases: Empty IDs, empty name, nil requests
//
// =============================================================================

func TestAcceptance_SelfServiceBrandingMacOS_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingMacOS
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test self-service branding macOS")

	createReq := &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "sdkv2_acc_Self Service",
		BrandingName:          acc.UniqueName("sdkv2_acc_acc-test-ssb-macos"),
		BrandingNameSecondary: "Acceptance Test",
		HomeHeading:           "Welcome",
		HomeSubheading:        "Choose an item below",
	}
	created, createResp, err := svc.Create(ctx, createReq)
	require.NoError(t, err, "Create should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	brandingID := created.ID
	acc.LogTestSuccess(t, "Self-service branding created with ID=%s", brandingID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, brandingID)
		acc.LogCleanupDeleteError(t, "self-service branding macOS", brandingID, delErr)
	})

	// 2. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting branding by ID=%s", brandingID)

	var fetched *self_service_branding_macos.ResourceSelfServiceBrandingMacOS
	var fetchResp *interfaces.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByID(ctx, brandingID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, brandingID, fetched.ID)
	assert.Equal(t, createReq.BrandingName, fetched.BrandingName)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.BrandingName)

	// 3. GetByName
	acc.LogTestStage(t, "GetByName", "Getting branding by name=%s", createReq.BrandingName)

	byName, _, err := svc.GetByName(ctx, createReq.BrandingName)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, brandingID, byName.ID)
	acc.LogTestSuccess(t, "GetByName: ID=%s", byName.ID)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating branding ID=%s", brandingID)

	updateReq := &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "sdkv2_acc_Self Service",
		BrandingName:          acc.UniqueName("sdkv2_acc_acc-test-ssb-macos-updated"),
		BrandingNameSecondary: "Acceptance Test Updated",
		HomeHeading:           "Welcome Back",
		HomeSubheading:        "Choose an item below",
	}
	updated, updateResp, err := svc.UpdateByID(ctx, brandingID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Branding updated: ID=%s", brandingID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetByID(ctx, brandingID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.BrandingName, fetched2.BrandingName)
	assert.Equal(t, updateReq.HomeHeading, fetched2.HomeHeading)
	acc.LogTestSuccess(t, "Update verified: name=%q heading=%q", fetched2.BrandingName, fetched2.HomeHeading)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting branding ID=%s", brandingID)

	deleteResp, err := svc.DeleteByID(ctx, brandingID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Branding ID=%s deleted", brandingID)
}

// =============================================================================
// TestAcceptance_SelfServiceBrandingMacOS_list_with_rsql_filter
// =============================================================================

func TestAcceptance_SelfServiceBrandingMacOS_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingMacOS
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-ssb-macos")
	createReq := &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "sdkv2_acc_Self Service",
		BrandingName:          name,
		BrandingNameSecondary: "RSQL Test",
		HomeHeading:           "Welcome",
		HomeSubheading:        "Choose an item below",
	}

	created, _, err := svc.Create(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	brandingID := created.ID
	acc.LogTestSuccess(t, "Created self-service branding macOS ID=%s name=%q", brandingID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, brandingID)
		acc.LogCleanupDeleteError(t, "self-service branding macOS", brandingID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`brandingName=="%s"`, name),
	}

	list, listResp, err := svc.List(ctx, rsqlQuery)
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
	assert.True(t, found, "self-service branding macOS should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target branding found=%v", list.TotalCount, found)
}

func TestAcceptance_SelfServiceBrandingMacOS_list(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingMacOS

	result, resp, err := svc.List(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
	if result.TotalCount > 0 {
		assert.NotNil(t, result.Results)
	}
}

func TestAcceptance_SelfServiceBrandingMacOS_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SelfServiceBrandingMacOS

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding configuration ID is required")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding configuration name is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
			BrandingName: "sdkv2_acc_x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "self-service branding configuration ID is required")
	})
}
