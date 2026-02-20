package ldap_servers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// LDAPServersServiceInterface defines the interface for Classic API LDAP server operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ldapservers
	LDAPServersServiceInterface interface {
		// ListLDAPServers returns all LDAP servers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapservers
		ListLDAPServers(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetLDAPServerByID returns the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapserversbyid
		GetLDAPServerByID(ctx context.Context, id int) (*ResourceLDAPServer, *interfaces.Response, error)

		// GetLDAPServerByName returns the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findldapserversbyname
		GetLDAPServerByName(ctx context.Context, name string) (*ResourceLDAPServer, *interfaces.Response, error)

		// CreateLDAPServer creates a new LDAP server.
		//
		// Returns a ListItem with the created server's ID and name (not the full resource).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createldapserverbyid
		CreateLDAPServer(ctx context.Context, req *RequestLDAPServer) (*ListItem, *interfaces.Response, error)

		// UpdateLDAPServerByID updates the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateldapserverbyid
		UpdateLDAPServerByID(ctx context.Context, id int, req *RequestLDAPServer) (*ResourceLDAPServer, *interfaces.Response, error)

		// UpdateLDAPServerByName updates the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateldapserverbyname
		UpdateLDAPServerByName(ctx context.Context, name string, req *RequestLDAPServer) (*ResourceLDAPServer, *interfaces.Response, error)

		// DeleteLDAPServerByID removes the specified LDAP server by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyid
		DeleteLDAPServerByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteLDAPServerByName removes the specified LDAP server by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyname
		DeleteLDAPServerByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the LDAP servers-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ldapservers
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ LDAPServersServiceInterface = (*Service)(nil)

// NewService returns a new LDAP servers Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - LDAP Servers CRUD Operations
// -----------------------------------------------------------------------------

// ListLDAPServers returns all LDAP servers.
// URL: GET /JSSResource/ldapservers
// https://developer.jamf.com/jamf-pro/reference/findldapservers
func (s *Service) ListLDAPServers(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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

// GetLDAPServerByID returns the specified LDAP server by ID.
// URL: GET /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findldapserversbyid
func (s *Service) GetLDAPServerByID(ctx context.Context, id int) (*ResourceLDAPServer, *interfaces.Response, error) {
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

// GetLDAPServerByName returns the specified LDAP server by name.
// URL: GET /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findldapserversbyname
func (s *Service) GetLDAPServerByName(ctx context.Context, name string) (*ResourceLDAPServer, *interfaces.Response, error) {
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

// CreateLDAPServer creates a new LDAP server.
// URL: POST /JSSResource/ldapservers/id/0
// Returns a ListItem with the created server's ID and name (not the full resource).
// https://developer.jamf.com/jamf-pro/reference/createldapserverbyid
func (s *Service) CreateLDAPServer(ctx context.Context, req *RequestLDAPServer) (*ListItem, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicLDAPServers)

	var createResp CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &createResp)
	if err != nil {
		return nil, resp, err
	}

	// Convert CreateResponse to ListItem
	result := &ListItem{
		ID:   createResp.ID,
		Name: createResp.Name,
	}

	return result, resp, nil
}

// UpdateLDAPServerByID updates the specified LDAP server by ID.
// URL: PUT /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateldapserverbyid
func (s *Service) UpdateLDAPServerByID(ctx context.Context, id int, req *RequestLDAPServer) (*ResourceLDAPServer, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("LDAP server ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicLDAPServers, id)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateLDAPServerByName updates the specified LDAP server by name.
// URL: PUT /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateldapserverbyname
func (s *Service) UpdateLDAPServerByName(ctx context.Context, name string, req *RequestLDAPServer) (*ResourceLDAPServer, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("LDAP server name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicLDAPServers, name)

	var result ResourceLDAPServer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteLDAPServerByID removes the specified LDAP server by ID.
// URL: DELETE /JSSResource/ldapservers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyid
func (s *Service) DeleteLDAPServerByID(ctx context.Context, id int) (*interfaces.Response, error) {
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

// DeleteLDAPServerByName removes the specified LDAP server by name.
// URL: DELETE /JSSResource/ldapservers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteldapserverbyname
func (s *Service) DeleteLDAPServerByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
