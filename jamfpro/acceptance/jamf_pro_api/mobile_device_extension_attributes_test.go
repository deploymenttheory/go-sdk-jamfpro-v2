package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mobile_device_extension_attributes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueNameMDEA(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

func TestAcceptance_MobileDeviceExtensionAttributes_Lifecycle(t *testing.T) {
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
	created, createResp, err := svc.CreateMobileDeviceExtensionAttributeV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	eaID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteMobileDeviceExtensionAttributeByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "mobile device extension attribute", eaID, delErr)
	})

	acc.LogTestStage(t, "List", "Listing mobile device extension attributes")
	list, listResp, err := svc.ListMobileDeviceExtensionAttributesV1(ctx, map[string]string{"page": "0", "page-size": "200"})
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

	acc.LogTestStage(t, "GetByID", "Fetching mobile device extension attribute by ID=%s", eaID)
	fetched, fetchResp, err := svc.GetMobileDeviceExtensionAttributeByIDV1(ctx, eaID)
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
	updated, updateResp, err := svc.UpdateMobileDeviceExtensionAttributeByIDV1(ctx, eaID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)
	assert.Equal(t, eaID, updated.ID)

	verified, _, err := svc.GetMobileDeviceExtensionAttributeByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, updatedName, verified.Name)

	acc.LogTestStage(t, "Delete", "Deleting mobile device extension attribute ID=%s", eaID)
	deleteResp, err := svc.DeleteMobileDeviceExtensionAttributeByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
}
