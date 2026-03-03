package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueNameMDEA(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_MobileDeviceExtensionAttributes_list_with_rsql_filter
// =============================================================================

func TestAcceptance_MobileDeviceExtensionAttributes_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceExtensionAttributes
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-mdea")
	createReq := &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
		Name:                 name,
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	eaID := created.ID
	acc.LogTestSuccess(t, "Created mobile device extension attribute ID=%s name=%q", eaID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "mobile device extension attribute", eaID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, ea := range list.Results {
		if ea.ID == eaID {
			found = true
			assert.Equal(t, name, ea.Name)
			break
		}
	}
	assert.True(t, found, "mobile device extension attribute should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target MDEA found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_MobileDeviceExtensionAttributes_delete_multiple
// =============================================================================

func TestAcceptance_MobileDeviceExtensionAttributes_delete_multiple(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceExtensionAttributes
	ctx := context.Background()

	c1, _, err := svc.CreateV1(ctx, &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
		Name:                 acc.UniqueName("sdkv2_acc_delmulti-mdea-1"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	})
	require.NoError(t, err)
	require.NotNil(t, c1)

	c2, _, err := svc.CreateV1(ctx, &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
		Name:                 acc.UniqueName("sdkv2_acc_delmulti-mdea-2"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	})
	require.NoError(t, err)
	require.NotNil(t, c2)

	acc.LogTestSuccess(t, "Created mobile device extension attributes ID=%s and ID=%s", c1.ID, c2.ID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, c1.ID)
		_, _ = svc.DeleteByIDV1(cleanupCtx, c2.ID)
	})

	resp, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(ctx, &mobile_device_extension_attributes.DeleteMobileDeviceExtensionAttributesByIDRequest{
		IDs: []string{c1.ID, c2.ID},
	})
	if err != nil && resp != nil && resp.StatusCode == 405 {
		t.Skip("Bulk delete endpoint not available (405 Method Not Allowed)")
	}
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
	acc.LogTestSuccess(t, "Bulk delete completed, status=%d", resp.StatusCode)
}

// =============================================================================
// TestAcceptance_MobileDeviceExtensionAttributes_validation_errors
// =============================================================================

func TestAcceptance_MobileDeviceExtensionAttributes_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceExtensionAttributes

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device extension attribute ID is required")
	})

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
			Name:                 "x",
			DataType:             "String",
			InventoryDisplayType: "General",
			InputType:            "TEXT",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByIDV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device extension attribute ID is required")
	})

	t.Run("DeleteMultiple_NilRequest", func(t *testing.T) {
		_, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ids are required")
	})

	t.Run("DeleteMultiple_EmptyIDs", func(t *testing.T) {
		_, err := svc.DeleteMobileDeviceExtensionAttributesByIDV1(context.Background(), &mobile_device_extension_attributes.DeleteMobileDeviceExtensionAttributesByIDRequest{
			IDs: []string{},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ids are required")
	})
}

func TestAcceptance_MobileDeviceExtensionAttributes_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceExtensionAttributes
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating test mobile device extension attribute")

	createReq := &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
		Name:                 uniqueNameMDEA("acc-test-mdea"),
		Description:          "Acceptance test MDEA",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	eaID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "mobile device extension attribute", eaID, delErr)
	})

	acc.LogTestStage(t, "List", "Listing mobile device extension attributes")
	list, listResp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, r := range list.Results {
		if r.ID == eaID {
			found = true
			assert.Equal(t, createReq.Name, r.Name)
			break
		}
	}
	assert.True(t, found)

	acc.LogTestStage(t, "GetByID", "Getting mobile device extension attribute by ID=%s", eaID)
	var fetched *mobile_device_extension_attributes.ResourceMobileDeviceExtensionAttribute
	var fetchResp *interfaces.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByIDV1(ctx, eaID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, eaID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)

	acc.LogTestStage(t, "Update", "Updating mobile device extension attribute ID=%s", eaID)
	updatedName := uniqueNameMDEA("acc-test-mdea-updated")
	updateReq := &mobile_device_extension_attributes.RequestMobileDeviceExtensionAttribute{
		Name:                 updatedName,
		Description:          "Updated description",
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, eaID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)
	assert.Equal(t, eaID, updated.ID)

	verified, _, err := svc.GetByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, updatedName, verified.Name)

	acc.LogTestStage(t, "Delete", "Deleting mobile device extension attribute ID=%s", eaID)
	deleteResp, err := svc.DeleteByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
}
