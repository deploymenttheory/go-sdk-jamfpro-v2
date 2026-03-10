package classic_ldap

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the classic LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
	ClassicLdap struct {
		client client.Client
	}
)

func NewClassicLdap(client client.Client) *ClassicLdap {
	return &ClassicLdap{client: client}
}

// GetMappingsByIDV1 returns the LDAP attribute mappings for an OnPrem LDAP configuration by ID.
// URL: GET /api/v1/classic-ldap/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
func (s *ClassicLdap) GetMappingsByIDV1(ctx context.Context, id string) (*ResourceOnPremLdapMappingsV1, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProClassicLdapV1, id)
	var result ResourceOnPremLdapMappingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
