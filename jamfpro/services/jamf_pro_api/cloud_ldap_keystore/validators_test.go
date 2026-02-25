package cloud_ldap_keystore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_Success(t *testing.T) {
	req := &ValidateKeystoreRequest{
		Password:  "test-password",
		FileBytes: "SGVsbG8gV29ybGQ=",
		FileName:  "keystore.jks",
	}

	err := validateKeystoreRequest(req)
	assert.NoError(t, err)
}

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_ValidBase64Patterns(t *testing.T) {
	validBase64Strings := []string{
		"SGVsbG8gV29ybGQ=",
		"dGVzdA==",
		"YQ==",
		"YWJj",
		"MTIzNDU2Nzg5MA==",
		"",
	}

	for _, b64 := range validBase64Strings {
		req := &ValidateKeystoreRequest{
			Password:  "test-password",
			FileBytes: b64,
			FileName:  "keystore.jks",
		}

		if b64 == "" {
			err := validateKeystoreRequest(req)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "fileBytes is required")
		} else {
			err := validateKeystoreRequest(req)
			assert.NoError(t, err, "Expected %q to be valid base64", b64)
		}
	}
}

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_InvalidBase64(t *testing.T) {
	invalidBase64Strings := []string{
		"Not base64!",
		"Invalid@Characters#Here",
		"SGVsbG8=====",
	}

	for _, b64 := range invalidBase64Strings {
		req := &ValidateKeystoreRequest{
			Password:  "test-password",
			FileBytes: b64,
			FileName:  "keystore.jks",
		}

		err := validateKeystoreRequest(req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "fileBytes must be valid base64 encoded data")
	}
}

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_MissingPassword(t *testing.T) {
	req := &ValidateKeystoreRequest{
		Password:  "",
		FileBytes: "SGVsbG8gV29ybGQ=",
		FileName:  "keystore.jks",
	}

	err := validateKeystoreRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "password is required")
}

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_MissingFileBytes(t *testing.T) {
	req := &ValidateKeystoreRequest{
		Password:  "test-password",
		FileBytes: "",
		FileName:  "keystore.jks",
	}

	err := validateKeystoreRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fileBytes is required")
}

func TestUnit_CloudLdapKeystore_ValidateKeystoreRequest_MissingFileName(t *testing.T) {
	req := &ValidateKeystoreRequest{
		Password:  "test-password",
		FileBytes: "SGVsbG8gV29ybGQ=",
		FileName:  "",
	}

	err := validateKeystoreRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "fileName is required")
}
