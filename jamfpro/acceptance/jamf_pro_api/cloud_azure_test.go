package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudAzure_GetDefaultServerConfiguration(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudAzure
	ctx := context.Background()

	result, _, err := svc.GetDefaultServerConfigurationV1(ctx)
	if err != nil {
		t.Skipf("Failed to get default server configuration (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Mappings.UserId)
}

func TestCloudAzure_GetDefaultMappings(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudAzure
	ctx := context.Background()

	result, _, err := svc.GetDefaultMappingsV1(ctx)
	if err != nil {
		t.Skipf("Failed to get default mappings (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.NotEmpty(t, result.UserId)
}

func TestCloudAzure_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudAzure
	ctx := context.Background()

	displayName := acc.UniqueName("Test Azure IDP")

	createReq := &cloud_azure.ResourceCloudAzure{
		CloudIdPCommon: cloud_azure.CloudIdPCommon{
			DisplayName:  displayName,
			ProviderName: "AZURE",
		},
		Server: cloud_azure.CloudAzureServer{
			TenantId:                             "12345678-1234-1234-1234-123456789012",
			Enabled:                              true,
			Migrated:                             false,
			SearchTimeout:                        30,
			TransitiveMembershipEnabled:          true,
			TransitiveMembershipUserField:        "userPrincipalName",
			TransitiveDirectoryMembershipEnabled: false,
			Mappings: cloud_azure.CloudAzureServerMappings{
				UserId:    "objectGUID",
				UserName:  "userPrincipalName",
				RealName:  "displayName",
				Email:     "mail",
				GroupId:   "objectGUID",
				GroupName: "cn",
			},
		},
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	if err != nil {
		t.Skipf("Failed to create Azure Cloud IDP (may not be supported on this tenant): %v", err)
		return
	}
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		svc.DeleteByIDV1(ctx, created.ID)
	})

	fetched, _, err := svc.GetByIDV1(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fetched.Server.ID)
	assert.Equal(t, displayName, fetched.CloudIdPCommon.DisplayName)
	assert.Equal(t, "AZURE", fetched.CloudIdPCommon.ProviderName)

	fetchedByName, _, err := svc.GetByNameV1(ctx, displayName)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fetchedByName.Server.ID)

	updateReq := &cloud_azure.ResourceCloudAzure{
		CloudIdPCommon: cloud_azure.CloudIdPCommon{
			DisplayName:  displayName + " Updated",
			ProviderName: "AZURE",
		},
		Server: fetched.Server,
	}

	updated, _, err := svc.UpdateByIDV1(ctx, created.ID, updateReq)
	require.NoError(t, err)
	assert.Equal(t, displayName+" Updated", updated.CloudIdPCommon.DisplayName)

	_, err = svc.DeleteByIDV1(ctx, created.ID)
	require.NoError(t, err)
}
