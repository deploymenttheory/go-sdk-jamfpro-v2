package cloud_ldap_keystore

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CloudLdapKeystoreServiceInterface defines the interface for Cloud LDAP Keystore operations.
	// Uses v1 API for validation. Supports keystore verification for Cloud LDAP configurations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
	CloudLdapKeystoreServiceInterface interface {
		// ValidateV1 validates a Cloud LDAP keystore (Validate LDAP Keystore).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
		ValidateV1(ctx context.Context, request *ValidateKeystoreRequest) (*ResponseValidateKeystore, *interfaces.Response, error)
	}

	// Service handles communication with the Cloud LDAP Keystore-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CloudLdapKeystoreServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ValidateV1 validates a Cloud LDAP keystore.
// URL: POST /api/v1/ldap-keystore/verify
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
func (s *Service) ValidateV1(ctx context.Context, request *ValidateKeystoreRequest) (*ResponseValidateKeystore, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResponseValidateKeystore

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointCloudLdapKeystoreV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
