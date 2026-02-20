package accounts_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AccountGroupsServiceInterface defines the interface for Classic API account group operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	AccountGroupsServiceInterface interface {
		// GetAccountGroupByID returns the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
		GetAccountGroupByID(ctx context.Context, id int) (*ResourceAccountGroup, *interfaces.Response, error)

		// GetAccountGroupByName returns the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
		GetAccountGroupByName(ctx context.Context, name string) (*ResourceAccountGroup, *interfaces.Response, error)

		// CreateAccountGroup creates a new account group.
		//
		// Returns only the created account group's ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createaccountbyid
		CreateAccountGroup(ctx context.Context, req *RequestAccountGroup) (*CreateResponse, *interfaces.Response, error)

		// UpdateAccountGroupByID updates the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
		UpdateAccountGroupByID(ctx context.Context, id int, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error)

		// UpdateAccountGroupByName updates the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
		UpdateAccountGroupByName(ctx context.Context, name string, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error)

		// DeleteAccountGroupByID removes the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
		DeleteAccountGroupByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteAccountGroupByName removes the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
		DeleteAccountGroupByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the account groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AccountGroupsServiceInterface = (*Service)(nil)

// NewService returns a new account groups Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Account Groups CRUD Operations
// -----------------------------------------------------------------------------

// GetAccountGroupByID returns the specified account group by ID.
// URL: GET /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
func (s *Service) GetAccountGroupByID(ctx context.Context, id int) (*ResourceAccountGroup, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", EndpointClassicAccounts, id)

	var result ResourceAccountGroup

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

// GetAccountGroupByName returns the specified account group by name.
// URL: GET /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
func (s *Service) GetAccountGroupByName(ctx context.Context, name string) (*ResourceAccountGroup, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account group name is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", EndpointClassicAccounts, name)

	var result ResourceAccountGroup

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

// CreateAccountGroup creates a new account group.
// URL: POST /JSSResource/accounts/groupid/0
// Returns only the created account group's ID.
// https://developer.jamf.com/jamf-pro/reference/createaccountbyid
func (s *Service) CreateAccountGroup(ctx context.Context, req *RequestAccountGroup) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupid/0", EndpointClassicAccounts)

	var result CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateAccountGroupByID updates the specified account group by ID.
// URL: PUT /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
func (s *Service) UpdateAccountGroupByID(ctx context.Context, id int, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", EndpointClassicAccounts, id)

	var result ResourceAccountGroup

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

// UpdateAccountGroupByName updates the specified account group by name.
// URL: PUT /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
func (s *Service) UpdateAccountGroupByName(ctx context.Context, name string, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account group name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", EndpointClassicAccounts, name)

	var result ResourceAccountGroup

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

// DeleteAccountGroupByID removes the specified account group by ID.
// URL: DELETE /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
func (s *Service) DeleteAccountGroupByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("account group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", EndpointClassicAccounts, id)

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

// DeleteAccountGroupByName removes the specified account group by name.
// URL: DELETE /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
func (s *Service) DeleteAccountGroupByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("account group name is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", EndpointClassicAccounts, name)

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
