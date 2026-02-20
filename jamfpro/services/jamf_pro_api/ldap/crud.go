package ldap

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// LdapServiceInterface defines the interface for LDAP operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
	LdapServiceInterface interface {
		// GetLdapGroupsV1 retrieves LDAP groups whose names contain the supplied search text (Get LDAP groups).
		//
		// Optional rsqlQuery keys: "filter" (RSQL), "sort", "page", "page-size". Omit or nil for no filter.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
		GetLdapGroupsV1(ctx context.Context, rsqlQuery map[string]string) (*ListGroupsResponseV1, *interfaces.Response, error)

		// GetLdapServersV1 retrieves every active LDAP or cloud identity provider server definition (Get LDAP servers).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-servers
		GetLdapServersV1(ctx context.Context) ([]ResourceLdapServerV1, *interfaces.Response, error)
	}

	// Service handles communication with the LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ LdapServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - LDAP CRUD Operations
// -----------------------------------------------------------------------------

// GetLdapGroupsV1 retrieves LDAP groups. Optional rsqlQuery: filter (RSQL), sort, page, page-size.
// URL: GET /api/v1/ldap/groups
// https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
func (s *Service) GetLdapGroupsV1(ctx context.Context, rsqlQuery map[string]string) (*ListGroupsResponseV1, *interfaces.Response, error) {

	var result ListGroupsResponseV1

	endpoint := EndpointLdapGroupsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetLdapServersV1 retrieves every active LDAP or cloud identity provider server definition.
// URL: GET /api/v1/ldap/servers
// https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-servers
func (s *Service) GetLdapServersV1(ctx context.Context) ([]ResourceLdapServerV1, *interfaces.Response, error) {
	var result []ResourceLdapServerV1

	endpoint := EndpointLdapServersV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
