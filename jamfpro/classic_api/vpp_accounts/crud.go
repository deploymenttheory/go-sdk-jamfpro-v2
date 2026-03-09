package vpp_accounts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the VPP account-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppaccounts
	VppAccounts struct {
		client transport.HTTPClient
	}
)

// NewService returns a new VPP accounts Service backed by the provided HTTP client.
func NewVppAccounts(client transport.HTTPClient) *VppAccounts {
	return &VppAccounts{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - VPP Accounts CRUD Operations
// -----------------------------------------------------------------------------

// List returns all VPP accounts.
// URL: GET /JSSResource/vppaccounts
// https://developer.jamf.com/jamf-pro/reference/findvppadminaccount
func (s *VppAccounts) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicVPPAccounts

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

// GetByID returns the specified VPP account by ID.
// URL: GET /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findvppadminaccountbyid
func (s *VppAccounts) GetByID(ctx context.Context, id int) (*ResourceVPPAccount, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

	var result ResourceVPPAccount

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

// Create creates a new VPP account.
// URL: POST /JSSResource/vppaccounts/id/0
// Returns the created VPP account with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createvppadminaccountbyid
func (s *VppAccounts) Create(ctx context.Context, req *RequestVPPAccount) (*ResourceVPPAccount, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicVPPAccounts)

	var result ResourceVPPAccount

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

// UpdateByID updates the specified VPP account by ID.
// URL: PUT /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatevppadminaccountbyid
func (s *VppAccounts) UpdateByID(ctx context.Context, id int, req *RequestVPPAccount) (*ResourceVPPAccount, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP account ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

	var result ResourceVPPAccount

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

// DeleteByID removes the specified VPP account by ID.
// URL: DELETE /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletevppadminaccountbyid
func (s *VppAccounts) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

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
