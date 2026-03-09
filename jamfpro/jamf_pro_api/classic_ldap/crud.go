package classic_ldap

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the classic LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
	ClassicLdap struct {
		client transport.HTTPClient
	}
)

func NewClassicLdap(client transport.HTTPClient) *ClassicLdap {
	return &ClassicLdap{client: client}
}

// GetMappingsByIDV1 returns the LDAP attribute mappings for an OnPrem LDAP configuration by ID.
// URL: GET /api/v1/classic-ldap/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
func (s *ClassicLdap) GetMappingsByIDV1(ctx context.Context, id string) (*ResourceOnPremLdapMappingsV1, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProClassicLdapV1, id)
	var result ResourceOnPremLdapMappingsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
