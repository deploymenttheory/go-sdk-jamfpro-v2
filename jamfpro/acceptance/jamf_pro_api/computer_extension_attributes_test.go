package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func uniqueCategoryName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

func TestAcceptance_ComputerExtensionAttributes_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerExtensionAttributes
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating test computer extension attribute")

	enabled := true
	createReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 uniqueCategoryName("acc-test-ea"),
		Description:          "Acceptance test EA",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}
	created, createResp, err := svc.CreateComputerExtensionAttributeV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	eaID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteComputerExtensionAttributeByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "computer extension attribute", eaID, delErr)
	})

	acc.LogTestStage(t, "List", "Listing computer extension attributes")
	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListComputerExtensionAttributesV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
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

	acc.LogTestStage(t, "GetByID", "Fetching computer extension attribute by ID=%s", eaID)
	fetched, fetchResp, err := svc.GetComputerExtensionAttributeByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, eaID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)

	acc.LogTestStage(t, "Update", "Updating computer extension attribute ID=%s", eaID)
	updatedName := uniqueCategoryName("acc-test-ea-updated")
	updateReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 updatedName,
		Description:          "Updated description",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}
	updated, updateResp, err := svc.UpdateComputerExtensionAttributeByIDV1(ctx, eaID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)
	assert.Equal(t, eaID, updated.ID)

	verified, _, err := svc.GetComputerExtensionAttributeByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, updatedName, verified.Name)

	acc.LogTestStage(t, "Delete", "Deleting computer extension attribute ID=%s", eaID)
	deleteResp, err := svc.DeleteComputerExtensionAttributeByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
}

func TestAcceptance_ComputerExtensionAttributes_DeleteMultiple(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerExtensionAttributes
	ctx := context.Background()

	enabled := true
	createReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 uniqueCategoryName("acc-delmulti-ea-1"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
		Enabled:              &enabled,
	}
	c1, _, err := svc.CreateComputerExtensionAttributeV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, c1)

	createReq2 := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 uniqueCategoryName("acc-delmulti-ea-2"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
		Enabled:              &enabled,
	}
	c2, _, err := svc.CreateComputerExtensionAttributeV1(ctx, createReq2)
	require.NoError(t, err)
	require.NotNil(t, c2)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteComputerExtensionAttributeByIDV1(cleanupCtx, c1.ID)
		_, _ = svc.DeleteComputerExtensionAttributeByIDV1(cleanupCtx, c2.ID)
	})

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(ctx, &computer_extension_attributes.DeleteComputerExtensionAttributesByIDRequest{
		IDs: []string{c1.ID, c2.ID},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}
