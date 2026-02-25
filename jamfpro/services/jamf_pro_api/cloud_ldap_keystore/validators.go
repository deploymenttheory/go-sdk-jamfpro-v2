package cloud_ldap_keystore

import (
	"fmt"
	"regexp"
)

var (
	base64Regex = regexp.MustCompile(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)
)

// validateKeystoreRequest validates the keystore request fields.
func validateKeystoreRequest(req *ValidateKeystoreRequest) error {
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}

	if req.FileBytes == "" {
		return fmt.Errorf("fileBytes is required")
	}

	if !base64Regex.MatchString(req.FileBytes) {
		return fmt.Errorf("fileBytes must be valid base64 encoded data")
	}

	if req.FileName == "" {
		return fmt.Errorf("fileName is required")
	}

	return nil
}
