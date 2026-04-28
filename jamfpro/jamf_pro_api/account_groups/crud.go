package account_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// AccountGroups handles communication with the account groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups
	AccountGroups struct {
		client client.Client
	}
)

// NewAccountGroups returns a new AccountGroups service.
func NewAccountGroups(client client.Client) *AccountGroups {
	return &AccountGroups{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Account Groups Operations
// -----------------------------------------------------------------------------

// ListV1 returns all account groups with pagination, sorting, and filtering support.
// URL: GET /api/v1/account-groups
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups
func (s *AccountGroups) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListAccountGroupsResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProAccountGroupsV1

	var result ListAccountGroupsResponse

	mergePage := func(pageData []byte) error {
		var items []ResourceAccountGroup
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
		return nil, resp, fmt.Errorf("failed to list account groups: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the account group for the given id.
// URL: GET /api/v1/account-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups-id
func (s *AccountGroups) GetByIDV1(ctx context.Context, id string) (*ResourceAccountGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("account group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAccountGroupsV1, id)

	var result ResourceAccountGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get account group by ID: %w", err)
	}

	return &result, resp, nil
}
