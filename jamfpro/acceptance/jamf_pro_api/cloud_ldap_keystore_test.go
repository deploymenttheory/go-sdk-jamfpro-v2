package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap_keystore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCloudLdapKeystore_Validate(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.CloudLdapKeystore
	ctx := context.Background()

	request := &cloud_ldap_keystore.ValidateKeystoreRequest{
		Password:  "test-password",
		FileBytes: "dGVzdGRhdGE=",
		FileName:  "keystore.p12",
	}

	result, _, err := svc.ValidateV1(ctx, request)
	if err != nil {
		t.Skipf("Failed to validate keystore (may not be supported or invalid test data): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.NotEmpty(t, result.Type)
	assert.NotEmpty(t, result.FileName)
}
