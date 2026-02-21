package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudLdap_GetDefaults(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.CloudLdap

	mappings, _, err := svc.GetDefaultMappingsV2(ctx, "GOOGLE")
	if err != nil {
		t.Skipf("Failed to get default mappings (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotNil(t, mappings)
	assert.NotEmpty(t, mappings.UserMappings.UserID)
	assert.NotEmpty(t, mappings.GroupMappings.GroupName)

	serverConfig, _, err := svc.GetDefaultServerConfigurationV2(ctx, "GOOGLE")
	if err != nil {
		t.Skipf("Failed to get default server configuration (may not be supported on this tenant): %v", err)
		return
	}

	assert.NotNil(t, serverConfig)
	assert.NotEmpty(t, serverConfig.ConnectionType)
}

func TestCloudLdap_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.CloudLdap

	displayName := acc.UniqueName("Test Google LDAP")

	createReq := &cloud_ldap.ResourceCloudLdap{
		CloudIdPCommon: &cloud_ldap.CloudIdPCommon{
			ProviderName: "GOOGLE",
			DisplayName:  displayName,
		},
		Server: &cloud_ldap.CloudLdapServer{
			Enabled:           true,
			UseWildcards:      false,
			ConnectionType:    "LDAPS",
			ServerUrl:         "ldaps://ldap.google.com",
			DomainName:        "example.com",
			Port:              636,
			ConnectionTimeout: 15,
			SearchTimeout:     60,
		},
		Mappings: &cloud_ldap.CloudLdapMappings{
			UserMappings: cloud_ldap.CloudLdapUserMappings{
				ObjectClassLimitation: "user",
				ObjectClasses:         "person",
				SearchBase:            "dc=example,dc=com",
				SearchScope:           "ALL_SUBTREES",
				UserID:                "uid",
				Username:              "uid",
				RealName:              "cn",
				EmailAddress:          "mail",
			},
			GroupMappings: cloud_ldap.CloudLdapGroupMappings{
				ObjectClassLimitation: "group",
				ObjectClasses:         "groupOfNames",
				SearchBase:            "dc=example,dc=com",
				SearchScope:           "ALL_SUBTREES",
				GroupID:               "cn",
				GroupName:             "cn",
			},
			MembershipMappings: cloud_ldap.CloudLdapMembershipMappings{
				GroupMembershipMapping: "member",
			},
		},
	}

	created, _, err := svc.CreateV2(ctx, createReq)
	if err != nil {
		t.Skipf("Failed to create Cloud LDAP (may not be supported on this tenant): %v", err)
		return
	}
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		svc.DeleteByIDV2(ctx, created.ID)
	})

	fetched, _, err := svc.GetByIDV2(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fetched.CloudIdPCommon.ID)
	assert.Equal(t, displayName, fetched.CloudIdPCommon.DisplayName)
	assert.Equal(t, "GOOGLE", fetched.CloudIdPCommon.ProviderName)

	updateReq := &cloud_ldap.ResourceCloudLdap{
		CloudIdPCommon: &cloud_ldap.CloudIdPCommon{
			DisplayName:  displayName + " Updated",
			ProviderName: "GOOGLE",
		},
		Server: fetched.Server,
	}

	updated, _, err := svc.UpdateByIDV2(ctx, created.ID, updateReq)
	require.NoError(t, err)
	assert.Equal(t, displayName+" Updated", updated.CloudIdPCommon.DisplayName)

	mappings, _, err := svc.GetMappingsByIDV2(ctx, created.ID)
	require.NoError(t, err)
	assert.NotNil(t, mappings)

	updatedMappings, _, err := svc.UpdateMappingsByIDV2(ctx, created.ID, mappings)
	require.NoError(t, err)
	assert.NotNil(t, updatedMappings)

	bindStats, _, err := svc.GetBindConnectionPoolStatsByIDV2(ctx, created.ID)
	if err != nil {
		t.Logf("Failed to get bind connection pool stats (may not be available): %v", err)
	} else {
		assert.NotNil(t, bindStats)
	}

	searchStats, _, err := svc.GetSearchConnectionPoolStatsByIDV2(ctx, created.ID)
	if err != nil {
		t.Logf("Failed to get search connection pool stats (may not be available): %v", err)
	} else {
		assert.NotNil(t, searchStats)
	}

	connStatus, _, err := svc.TestConnectionByIDV2(ctx, created.ID)
	if err != nil {
		t.Logf("Failed to test connection (may not be available): %v", err)
	} else {
		assert.NotNil(t, connStatus)
		assert.NotEmpty(t, connStatus.Status)
	}

	_, err = svc.DeleteByIDV2(ctx, created.ID)
	require.NoError(t, err)
}
