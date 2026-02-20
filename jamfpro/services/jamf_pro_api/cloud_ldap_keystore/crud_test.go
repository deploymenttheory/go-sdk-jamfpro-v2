package cloud_ldap_keystore

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_ldap_keystore/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateV1(t *testing.T) {
	mock := mocks.NewCloudLdapKeystoreMock()
	mock.RegisterValidateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ValidateKeystoreRequest{
		Password:  "test-password",
		FileBytes: "base64encodeddata",
		FileName:  "keystore.p12",
	}

	result, resp, err := svc.ValidateV1(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "PKCS12", result.Type)
	assert.Equal(t, "keystore.p12", result.FileName)
	assert.NotEmpty(t, result.Subject)
	assert.NotEmpty(t, result.ExpirationDate)
}

func TestValidateV1_NilRequest(t *testing.T) {
	mock := mocks.NewCloudLdapKeystoreMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ValidateV1(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}
