package accounts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AccountsServiceInterface defines the interface for Classic API account operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	AccountsServiceInterface interface {
		// ListAccounts returns all accounts (both users and groups).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccounts
		ListAccounts(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetAccountByID returns the specified account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
		GetAccountByID(ctx context.Context, id int) (*ResourceAccount, *interfaces.Response, error)

		// GetAccountByName returns the specified account by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
		GetAccountByName(ctx context.Context, name string) (*ResourceAccount, *interfaces.Response, error)

		// CreateAccount creates a new account.
		//
		// Returns the created account with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createaccountbyid
		CreateAccount(ctx context.Context, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error)

		// UpdateAccountByID updates the specified account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
		UpdateAccountByID(ctx context.Context, id int, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error)

		// UpdateAccountByName updates the specified account by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
		UpdateAccountByName(ctx context.Context, name string, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error)

		// DeleteAccountByID removes the specified account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
		DeleteAccountByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteAccountByName removes the specified account by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
		DeleteAccountByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the accounts-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AccountsServiceInterface = (*Service)(nil)

// NewService returns a new accounts Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Accounts CRUD Operations
// -----------------------------------------------------------------------------

// ListAccounts returns all accounts (both users and groups).
// URL: GET /JSSResource/accounts
// https://developer.jamf.com/jamf-pro/reference/findaccounts
func (s *Service) ListAccounts(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicAccounts

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

// GetAccountByID returns the specified account by ID.
// URL: GET /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
func (s *Service) GetAccountByID(ctx context.Context, id int) (*ResourceAccount, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/userid/%d", EndpointClassicAccounts, id)

	var result ResourceAccount

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

// GetAccountByName returns the specified account by name.
// URL: GET /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
func (s *Service) GetAccountByName(ctx context.Context, name string) (*ResourceAccount, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account name is required")
	}

	endpoint := fmt.Sprintf("%s/username/%s", EndpointClassicAccounts, name)

	var result ResourceAccount

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

// CreateAccount creates a new account.
// URL: POST /JSSResource/accounts/userid/0
// Returns the created account with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createaccountbyid
func (s *Service) CreateAccount(ctx context.Context, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/userid/0", EndpointClassicAccounts)

	var result ResourceAccount

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

// UpdateAccountByID updates the specified account by ID.
// URL: PUT /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
func (s *Service) UpdateAccountByID(ctx context.Context, id int, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/userid/%d", EndpointClassicAccounts, id)

	var result ResourceAccount

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

// UpdateAccountByName updates the specified account by name.
// URL: PUT /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
func (s *Service) UpdateAccountByName(ctx context.Context, name string, req *RequestAccount) (*ResourceAccount, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/username/%s", EndpointClassicAccounts, name)

	var result ResourceAccount

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

// DeleteAccountByID removes the specified account by ID.
// URL: DELETE /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
func (s *Service) DeleteAccountByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/userid/%d", EndpointClassicAccounts, id)

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

// DeleteAccountByName removes the specified account by name.
// URL: DELETE /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
func (s *Service) DeleteAccountByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("account name is required")
	}

	endpoint := fmt.Sprintf("%s/username/%s", EndpointClassicAccounts, name)

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
