package cloud_ldap_keystore

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
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
		ValidateV1(ctx context.Context, request *ValidateKeystoreRequest) (*ResponseValidateKeystore, *resty.Response, error)
	}

	// Service handles communication with the Cloud LDAP Keystore-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
	CloudLdapKeystore struct {
		client transport.HTTPClient
	}
)

var _ CloudLdapKeystoreServiceInterface = (*CloudLdapKeystore)(nil)

func NewCloudLdapKeystore(client transport.HTTPClient) *CloudLdapKeystore {
	return &CloudLdapKeystore{client: client}
}

// ValidateV1 validates a Cloud LDAP keystore.
// URL: POST /api/v1/ldap-keystore/verify
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
func (s *CloudLdapKeystore) ValidateV1(ctx context.Context, request *ValidateKeystoreRequest) (*ResponseValidateKeystore, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if err := validateKeystoreRequest(request); err != nil {
		return nil, nil, err
	}

	var result ResponseValidateKeystore

	Endpoint := EndpointCloudLdapKeystoreV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, Endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
