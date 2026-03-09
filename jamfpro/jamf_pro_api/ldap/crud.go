package ldap

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
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
		GetLdapGroupsV1(ctx context.Context, rsqlQuery map[string]string) (*ListGroupsResponseV1, *resty.Response, error)

		// GetLdapServersV1 retrieves every active LDAP or cloud identity provider server definition (Get LDAP servers).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-servers
		GetLdapServersV1(ctx context.Context) ([]ResourceLdapServerV1, *resty.Response, error)

		// GetLdapServersOnlyV1 retrieves LDAP servers only (not migrated to cloud).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-ldap-servers
		GetLdapServersOnlyV1(ctx context.Context) ([]ResourceLdapServerV1, *resty.Response, error)
	}

	// Service handles communication with the LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
	Ldap struct {
		client transport.HTTPClient
	}
)

var _ LdapServiceInterface = (*Ldap)(nil)

func NewLdap(client transport.HTTPClient) *Ldap {
	return &Ldap{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - LDAP CRUD Operations
// -----------------------------------------------------------------------------

// GetLdapGroupsV1 retrieves LDAP groups. Optional rsqlQuery: filter (RSQL), sort, page, page-size.
// URL: GET /api/v1/ldap/groups
// https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-groups
func (s *Ldap) GetLdapGroupsV1(ctx context.Context, rsqlQuery map[string]string) (*ListGroupsResponseV1, *resty.Response, error) {

	var result ListGroupsResponseV1

	endpoint := constants.EndpointJamfProLdapGroupsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Ldap) GetLdapServersV1(ctx context.Context) ([]ResourceLdapServerV1, *resty.Response, error) {
	var result []ResourceLdapServerV1

	endpoint := constants.EndpointJamfProLdapServersV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetLdapServersOnlyV1 retrieves LDAP servers only (not migrated to cloud).
// URL: GET /api/v1/ldap/ldap-servers
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ldap-ldap-servers
func (s *Ldap) GetLdapServersOnlyV1(ctx context.Context) ([]ResourceLdapServerV1, *resty.Response, error) {
	var result []ResourceLdapServerV1

	endpoint := constants.EndpointJamfProLdapServersOnlyV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
