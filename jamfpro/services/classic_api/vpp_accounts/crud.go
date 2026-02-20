package vpp_accounts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// VPPAccountsServiceInterface defines the interface for Classic API VPP account operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppaccounts
	VPPAccountsServiceInterface interface {
		// ListVPPAccounts returns all VPP accounts.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallvppaccounts
		ListVPPAccounts(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetVPPAccountByID returns the specified VPP account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findvppaccountsbyid
		GetVPPAccountByID(ctx context.Context, id int) (*ResourceVPPAccount, *interfaces.Response, error)

		// CreateVPPAccount creates a new VPP account.
		//
		// Returns the created VPP account with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createvppaccount
		CreateVPPAccount(ctx context.Context, req *RequestVPPAccount) (*ResourceVPPAccount, *interfaces.Response, error)

		// UpdateVPPAccountByID updates the specified VPP account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatevppaccountbyid
		UpdateVPPAccountByID(ctx context.Context, id int, req *RequestVPPAccount) (*ResourceVPPAccount, *interfaces.Response, error)

		// DeleteVPPAccountByID removes the specified VPP account by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletevppaccountbyid
		DeleteVPPAccountByID(ctx context.Context, id int) (*interfaces.Response, error)
	}

	// Service handles communication with the VPP account-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppaccounts
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ VPPAccountsServiceInterface = (*Service)(nil)

// NewService returns a new VPP accounts Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - VPP Accounts CRUD Operations
// -----------------------------------------------------------------------------

// ListVPPAccounts returns all VPP accounts.
// URL: GET /JSSResource/vppaccounts
// https://developer.jamf.com/jamf-pro/reference/findallvppaccounts
func (s *Service) ListVPPAccounts(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicVPPAccounts

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

// GetVPPAccountByID returns the specified VPP account by ID.
// URL: GET /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findvppaccountsbyid
func (s *Service) GetVPPAccountByID(ctx context.Context, id int) (*ResourceVPPAccount, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicVPPAccounts, id)

	var result ResourceVPPAccount

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

// CreateVPPAccount creates a new VPP account.
// URL: POST /JSSResource/vppaccounts/id/0
// Returns the created VPP account with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createvppaccount
func (s *Service) CreateVPPAccount(ctx context.Context, req *RequestVPPAccount) (*ResourceVPPAccount, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicVPPAccounts)

	var result ResourceVPPAccount

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

// UpdateVPPAccountByID updates the specified VPP account by ID.
// URL: PUT /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatevppaccountbyid
func (s *Service) UpdateVPPAccountByID(ctx context.Context, id int, req *RequestVPPAccount) (*ResourceVPPAccount, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP account ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicVPPAccounts, id)

	var result ResourceVPPAccount

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

// DeleteVPPAccountByID removes the specified VPP account by ID.
// URL: DELETE /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletevppaccountbyid
func (s *Service) DeleteVPPAccountByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicVPPAccounts, id)

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
