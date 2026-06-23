package groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// RequestUpdateGroupV2 is the GroupUpdateDtoV2 PATCH body for the v2 unified
// group endpoint. Its shape matches RequestUpdateGroup (the criteria already use
// the V2 openingParen/closingParen form).
type RequestUpdateGroupV2 = RequestUpdateGroup

// -----------------------------------------------------------------------------
// Unified Groups CRUD (V2) — Jamf Pro 11.28, replaces the V1 surface.
// -----------------------------------------------------------------------------

// ListV2 retrieves a paginated list of unified (computer + mobile) groups.
// URL: GET /api/v2/groups
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
func (s *Groups) ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProGroupsV2

	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list groups: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV2 retrieves a group by its platform ID (groupPlatformId).
// URL: GET /api/v2/groups/{id}
func (s *Groups) GetByIDV2(ctx context.Context, id string) (*ResourceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV2, id)

	var result ResourceGroup
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 updates a group by its platform ID using the GroupUpdateDtoV2 body.
// URL: PATCH /api/v2/groups/{id}
func (s *Groups) UpdateByIDV2(ctx context.Context, id string, req *RequestUpdateGroupV2) (*ResourceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if err := validateGroupUpdate(req); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV2, id)

	var result ResourceGroup
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Patch(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV2 removes the specified group by its platform ID.
// URL: DELETE /api/v2/groups/{id}
func (s *Groups) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
