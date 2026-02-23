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
		// GetByID returns the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
		GetByID(ctx context.Context, id int) (*ResourceAccountGroup, *interfaces.Response, error)

		// GetByName returns the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
		GetByName(ctx context.Context, name string) (*ResourceAccountGroup, *interfaces.Response, error)

		// Create creates a new account group.
		//
		// Returns only the created account group's ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createaccountbyid
		Create(ctx context.Context, req *RequestAccountGroup) (*CreateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
		UpdateByID(ctx context.Context, id int, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error)

		// UpdateByName updates the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
		UpdateByName(ctx context.Context, name string, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error)

		// DeleteByID removes the specified account group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified account group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// GetByID returns the specified account group by ID.
// URL: GET /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceAccountGroup, *interfaces.Response, error) {
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

// GetByName returns the specified account group by name.
// URL: GET /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/findaccountsbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceAccountGroup, *interfaces.Response, error) {
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

// Create creates a new account group.
// URL: POST /JSSResource/accounts/groupid/0
// Returns only the created account group's ID.
// https://developer.jamf.com/jamf-pro/reference/createaccountbyid
func (s *Service) Create(ctx context.Context, req *RequestAccountGroup) (*CreateResponse, *interfaces.Response, error) {
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

// UpdateByID updates the specified account group by ID.
// URL: PUT /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error) {
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

// UpdateByName updates the specified account group by name.
// URL: PUT /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/updateaccountbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestAccountGroup) (*ResourceAccountGroup, *interfaces.Response, error) {
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

// DeleteByID removes the specified account group by ID.
// URL: DELETE /JSSResource/accounts/groupid/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
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

// DeleteByName removes the specified account group by name.
// URL: DELETE /JSSResource/accounts/groupname/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteaccountbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
