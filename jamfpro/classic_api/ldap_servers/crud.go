package ldap_servers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// LDAPServersServiceInterface defines the interface for Classic API LDAP server operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ldapservers
	LDAPServersServiceInterface interface {
		// List returns all LDAP servers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapservers
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapserversbyid
		GetByID(ctx context.Context, id int) (*ResourceLDAPServer, *resty.Response, error)

		// GetByName returns the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapserversbyname
		GetByName(ctx context.Context, name string) (*ResourceLDAPServer, *resty.Response, error)

		// Create creates a new LDAP server.
		//
		// Returns a ListItem with the created server's ID and name (not the full resource).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createldapserverbyid
		Create(ctx context.Context, request *RequestLDAPServer) (*ListItem, *resty.Response, error)

		// UpdateByID updates the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateldapserverbyid
		UpdateByID(ctx context.Context, id int, request *RequestLDAPServer) (*ResourceLDAPServer, *resty.Response, error)

		// UpdateByName updates the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateldapserverbyname
		UpdateByName(ctx context.Context, name string, request *RequestLDAPServer) (*ResourceLDAPServer, *resty.Response, error)

		// DeleteByID removes the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the LDAP servers-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ldapservers
	LdapServers struct {
		client transport.HTTPClient
	}
)

var _ LDAPServersServiceInterface = (*LdapServers)(nil)

// NewService returns a new LDAP servers Service backed by the provided HTTP client.
func NewLdapServers(client transport.HTTPClient) *LdapServers {
	return &LdapServers{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - LDAP Servers CRUD Operations
// -----------------------------------------------------------------------------

// List returns all LDAP servers.
// URL: GET /JSSResource/ldapservers
// https://developer.jamf.com/jamf-pro/reference/findldapservers
func (s *LdapServers) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicLDAPServers

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified LDAP server by ID.
// URL: GET /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findldapserversbyid
func (s *LdapServers) GetByID(ctx context.Context, id int) (*ResourceLDAPServer, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("LDAP server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicLDAPServers, id)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified LDAP server by name.
// URL: GET /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findldapserversbyname
func (s *LdapServers) GetByName(ctx context.Context, name string) (*ResourceLDAPServer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("LDAP server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicLDAPServers, name)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new LDAP server.
// URL: POST /JSSResource/ldapservers/id/0
// Returns a ListItem with the created server's ID and name (not the full resource).
// https://developer.jamf.com/jamf-pro/reference/createldapserverbyid
func (s *LdapServers) Create(ctx context.Context, request *RequestLDAPServer) (*ListItem, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicLDAPServers)

	var createResp CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &createResp)
	if err != nil {
		return nil, resp, err
	}

	// Convert CreateResponse to ListItem
	// Note: Classic API only returns ID in create response, name comes from request
	result := &ListItem{
		ID:   createResp.ID,
		Name: request.Connection.Name,
	}

	return result, resp, nil
}

// UpdateByID updates the specified LDAP server by ID.
// URL: PUT /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateldapserverbyid
func (s *LdapServers) UpdateByID(ctx context.Context, id int, request *RequestLDAPServer) (*ResourceLDAPServer, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("LDAP server ID must be a positive integer")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicLDAPServers, id)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified LDAP server by name.
// URL: PUT /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateldapserverbyname
func (s *LdapServers) UpdateByName(ctx context.Context, name string, request *RequestLDAPServer) (*ResourceLDAPServer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("LDAP server name is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicLDAPServers, name)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified LDAP server by ID.
// URL: DELETE /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyid
func (s *LdapServers) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("LDAP server ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicLDAPServers, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified LDAP server by name.
// URL: DELETE /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyname
func (s *LdapServers) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("LDAP server name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicLDAPServers, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
