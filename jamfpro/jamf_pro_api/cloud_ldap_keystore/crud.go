package cloud_ldap_keystore

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Cloud LDAP Keystore-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ldap-keystore-verify
	CloudLdapKeystore struct {
		client client.Client
	}
)

func NewCloudLdapKeystore(client client.Client) *CloudLdapKeystore {
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

	endpoint := constants.EndpointJamfProCloudLdapKeystoreV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
