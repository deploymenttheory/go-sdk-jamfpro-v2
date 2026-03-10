package vpp_accounts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the VPP account-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppaccounts
	VppAccounts struct {
		client client.Client
	}
)

// NewService returns a new VPP accounts Service backed by the provided HTTP client.
func NewVppAccounts(client client.Client) *VppAccounts {
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

// GetByID returns the specified VPP account by ID.
// URL: GET /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findvppadminaccountbyid
func (s *VppAccounts) GetByID(ctx context.Context, id int) (*ResourceVPPAccount, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	var result ResourceVPPAccount

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

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

// Create creates a new VPP account.
// URL: POST /JSSResource/vppaccounts/id/0
// Returns the created VPP account with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createvppadminaccountbyid
func (s *VppAccounts) Create(ctx context.Context, req *RequestVPPAccount) (*ResourceVPPAccount, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceVPPAccount

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicVPPAccounts)

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

	var result ResourceVPPAccount

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

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

// DeleteByID removes the specified VPP account by ID.
// URL: DELETE /JSSResource/vppaccounts/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletevppadminaccountbyid
func (s *VppAccounts) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("VPP account ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAccounts, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
