package classic_ldap

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ClassicLdapServiceInterface defines the interface for classic LDAP operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
	ClassicLdapServiceInterface interface {
		// GetMappingsByIDV1 returns the LDAP attribute mappings for an OnPrem LDAP configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
		GetMappingsByIDV1(ctx context.Context, id string) (*ResourceOnPremLdapMappingsV1, *resty.Response, error)
	}

	// Service handles communication with the classic LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ClassicLdapServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetMappingsByIDV1 returns the LDAP attribute mappings for an OnPrem LDAP configuration by ID.
// URL: GET /api/v1/classic-ldap/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
func (s *Service) GetMappingsByIDV1(ctx context.Context, id string) (*ResourceOnPremLdapMappingsV1, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", EndpointClassicLdapV1, id)
	var result ResourceOnPremLdapMappingsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
