package accounts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the accounts-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
	Accounts struct {
		client client.Client
	}
)

// NewService returns a new accounts Service backed by the provided HTTP client.
func NewAccounts(client client.Client) *Accounts {
	return &Accounts{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - User Accounts CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all user accounts with automatic pagination, sorting, and filtering support.
// URL: GET /api/v1/accounts
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
func (s *Accounts) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProAccountsV1

	var result ListResponse

	mergePage := func(pageData []byte) error {
		var items []ResourceAccount
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the user account for the given id.
// URL: GET /api/v1/accounts/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-accounts-id
func (s *Accounts) GetByIDV1(ctx context.Context, id string) (*ResourceAccount, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("account ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAccountsV1, id)

	var result ResourceAccount

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 adds a new user account.
// URL: POST /api/v1/accounts
// https://developer.jamf.com/jamf-pro/reference/post_v1-accounts
func (s *Accounts) CreateV1(ctx context.Context, req *RequestAccount) (*CreateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if req.AccessLevel != "" {
		if _, ok := validAccessLevels[req.AccessLevel]; !ok {
			return nil, nil, fmt.Errorf("invalid accessLevel %q: must be one of FullAccess, SiteAccess, GroupBasedAccess", req.AccessLevel)
		}
	}
	if req.PrivilegeLevel != "" {
		if _, ok := validPrivilegeLevels[req.PrivilegeLevel]; !ok {
			return nil, nil, fmt.Errorf("invalid privilegeLevel %q: must be one of ADMINISTRATOR, AUDITOR, ENROLLMENT, CUSTOM", req.PrivilegeLevel)
		}
	}
	if req.AccountStatus != "" {
		if _, ok := validAccountStatuses[req.AccountStatus]; !ok {
			return nil, nil, fmt.Errorf("invalid accountStatus %q: must be one of Enabled, Disabled", req.AccountStatus)
		}
	}
	if req.AccountType != "" {
		if _, ok := validAccountTypes[req.AccountType]; !ok {
			return nil, nil, fmt.Errorf("invalid accountType %q: must be one of DEFAULT, FEDERATED", req.AccountType)
		}
	}

	endpoint := constants.EndpointJamfProAccountsV1

	var result CreateResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the user account for the given id.
// URL: PUT /api/v1/accounts/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-accounts-id
func (s *Accounts) UpdateByIDV1(ctx context.Context, id string, req *RequestAccount) (*ResourceAccount, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("account ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if req.AccessLevel != "" {
		if _, ok := validAccessLevels[req.AccessLevel]; !ok {
			return nil, nil, fmt.Errorf("invalid accessLevel %q: must be one of FullAccess, SiteAccess, GroupBasedAccess", req.AccessLevel)
		}
	}
	if req.PrivilegeLevel != "" {
		if _, ok := validPrivilegeLevels[req.PrivilegeLevel]; !ok {
			return nil, nil, fmt.Errorf("invalid privilegeLevel %q: must be one of ADMINISTRATOR, AUDITOR, ENROLLMENT, CUSTOM", req.PrivilegeLevel)
		}
	}
	if req.AccountStatus != "" {
		if _, ok := validAccountStatuses[req.AccountStatus]; !ok {
			return nil, nil, fmt.Errorf("invalid accountStatus %q: must be one of Enabled, Disabled", req.AccountStatus)
		}
	}
	if req.AccountType != "" {
		if _, ok := validAccountTypes[req.AccountType]; !ok {
			return nil, nil, fmt.Errorf("invalid accountType %q: must be one of DEFAULT, FEDERATED", req.AccountType)
		}
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAccountsV1, id)

	var result ResourceAccount

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 deletes the user account for the given id.
// URL: DELETE /api/v1/accounts/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-accounts-id
func (s *Accounts) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("account ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAccountsV1, id)

	resp, err := s.client.NewRequest(ctx).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
