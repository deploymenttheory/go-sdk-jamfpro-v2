package accounts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the accounts-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	Accounts struct {
		client client.Client
	}
)

// NewService returns a new accounts Service backed by the provided HTTP client.
func NewAccounts(client client.Client) *Accounts {
	return &Accounts{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Accounts CRUD Operations
// -----------------------------------------------------------------------------

// List returns all accounts (both users and groups).
// URL: GET /JSSResource/accounts
// https://developer.jamf.com/jamf-pro/reference/findaccounts
func (s *Accounts) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicAccounts

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified account by ID.
// URL: GET /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
func (s *Accounts) GetByID(ctx context.Context, id int) (*ResourceAccount, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account ID must be a positive integer")
	}

	var result ResourceAccount

	endpoint := fmt.Sprintf("%s/userid/%d", constants.EndpointClassicAccounts, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified account by name.
// URL: GET /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
func (s *Accounts) GetByName(ctx context.Context, name string) (*ResourceAccount, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account name is required")
	}

	var result ResourceAccount

	endpoint := fmt.Sprintf("%s/username/%s", constants.EndpointClassicAccounts, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new account.
// URL: POST /JSSResource/accounts/userid/0
// Returns the created account with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createaccountbyid
func (s *Accounts) Create(ctx context.Context, req *RequestAccount) (*ResourceAccount, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccount

	endpoint := fmt.Sprintf("%s/userid/0", constants.EndpointClassicAccounts)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified account by ID.
// URL: PUT /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
func (s *Accounts) UpdateByID(ctx context.Context, id int, req *RequestAccount) (*ResourceAccount, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccount

	endpoint := fmt.Sprintf("%s/userid/%d", constants.EndpointClassicAccounts, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified account by name.
// URL: PUT /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
func (s *Accounts) UpdateByName(ctx context.Context, name string, req *RequestAccount) (*ResourceAccount, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAccount

	endpoint := fmt.Sprintf("%s/username/%s", constants.EndpointClassicAccounts, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified account by ID.
// URL: DELETE /JSSResource/accounts/userid/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
func (s *Accounts) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/userid/%d", constants.EndpointClassicAccounts, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified account by name.
// URL: DELETE /JSSResource/accounts/username/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
func (s *Accounts) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("account name is required")
	}

	endpoint := fmt.Sprintf("%s/username/%s", constants.EndpointClassicAccounts, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
