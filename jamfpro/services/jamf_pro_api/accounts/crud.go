package accounts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AccountsServiceInterface defines the interface for user account operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
	AccountsServiceInterface interface {
		// ListV1 returns all user accounts with pagination, sorting, and filtering support.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the user account for the given id.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts-id
		GetByIDV1(ctx context.Context, id string) (*ResourceAccount, *interfaces.Response, error)

		// CreateV1 adds a new user account.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-accounts
		CreateV1(ctx context.Context, req *RequestAccount) (*CreateResponse, *interfaces.Response, error)

		// DeleteByIDV1 deletes the user account for the given id.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-accounts-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the accounts-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
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
// Jamf Pro API - User Accounts CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all user accounts with automatic pagination, sorting, and filtering support.
// URL: GET /api/v1/accounts
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/get_v1-accounts
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointAccountsV1

	var result ListResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageResponse ListResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the user account for the given id.
// URL: GET /api/v1/accounts/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-accounts-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceAccount, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("account ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointAccountsV1, id)

	var result ResourceAccount

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 adds a new user account.
// URL: POST /api/v1/accounts
// https://developer.jamf.com/jamf-pro/reference/post_v1-accounts
func (s *Service) CreateV1(ctx context.Context, req *RequestAccount) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointAccountsV1

	var result CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 deletes the user account for the given id.
// URL: DELETE /api/v1/accounts/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-accounts-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("account ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointAccountsV1, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
