package accounts_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the account groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/accounts
	AccountsGroups struct {
		client transport.HTTPClient
	}
)

// NewService returns a new account groups Service backed by the provided HTTP client.
func NewAccountsGroups(client transport.HTTPClient) *AccountsGroups {
	return &AccountsGroups{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Account Groups CRUD Operations
// -----------------------------------------------------------------------------

// GetByID returns the specified account group by ID.
// URL: GET /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
func (s *AccountsGroups) GetByID(ctx context.Context, id int) (*ResourceAccountGroup, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", constants.EndpointClassicAccounts, id)

	var result ResourceAccountGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified account group by name.
// URL: GET /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
func (s *AccountsGroups) GetByName(ctx context.Context, name string) (*ResourceAccountGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account group name is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", constants.EndpointClassicAccounts, name)

	var result ResourceAccountGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new account group.
// URL: POST /JSSResource/accounts/groupid/0
// Returns only the created account group's ID.
// https://developer.jamf.com/jamf-pro/reference/createaccountbyid
func (s *AccountsGroups) Create(ctx context.Context, req *RequestAccountGroup) (*CreateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupid/0", constants.EndpointClassicAccounts)

	var result CreateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified account group by ID.
// URL: PUT /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
func (s *AccountsGroups) UpdateByID(ctx context.Context, id int, req *RequestAccountGroup) (*UpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("account group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", constants.EndpointClassicAccounts, id)

	var result UpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified account group by name.
// URL: PUT /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
func (s *AccountsGroups) UpdateByName(ctx context.Context, name string, req *RequestAccountGroup) (*UpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("account group name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", constants.EndpointClassicAccounts, name)

	var result UpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified account group by ID.
// URL: DELETE /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
func (s *AccountsGroups) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("account group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/groupid/%d", constants.EndpointClassicAccounts, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified account group by name.
// URL: DELETE /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
func (s *AccountsGroups) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("account group name is required")
	}

	endpoint := fmt.Sprintf("%s/groupname/%s", constants.EndpointClassicAccounts, name)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
