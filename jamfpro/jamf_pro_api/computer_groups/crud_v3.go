package computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"
	"resty.dev/v3"
)

// -----------------------------------------------------------------------------
// Smart Groups CRUD (V3) — Jamf Pro 11.28, replaces the V2 surface.
// -----------------------------------------------------------------------------

// ListSmartV3 returns all smart computer groups.
// URL: GET /api/v3/computer-groups/smart-groups
func (s *ComputerGroups) ListSmartV3(ctx context.Context, rsqlQuery map[string]string) (*ListSmartV3Response, *resty.Response, error) {
	var result ListSmartV3Response

	endpoint := constants.EndpointJamfProSmartComputerGroupsV3

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceSmartGroupV3
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil {
		if bodyBytes := resp.Bytes(); len(bodyBytes) > 0 {
			var pageResp struct {
				TotalCount int `json:"totalCount"`
			}
			if err := json.Unmarshal(bodyBytes, &pageResp); err == nil {
				result.TotalCount = pageResp.TotalCount
			}
		}
	}
	if result.TotalCount == 0 {
		result.TotalCount = len(result.Results)
	}

	return &result, resp, nil
}

// GetSmartByIDV3 returns the specified smart group by ID (including criteria).
// URL: GET /api/v3/computer-groups/smart-groups/{id}
func (s *ComputerGroups) GetSmartByIDV3(ctx context.Context, id string) (*ResourceSmartGroupV3, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartComputerGroupsV3, id)

	var result ResourceSmartGroupV3
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateSmartV3 creates a new smart computer group.
// URL: POST /api/v3/computer-groups/smart-groups
func (s *ComputerGroups) CreateSmartV3(ctx context.Context, request *RequestSmartGroupV3) (*CreateSmartResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateSmartGroupV3(request); err != nil {
		return nil, nil, err
	}

	var result CreateSmartResponse
	endpoint := constants.EndpointJamfProSmartComputerGroupsV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSmartByIDV3 updates the specified smart group by ID.
// URL: PUT /api/v3/computer-groups/smart-groups/{id}
func (s *ComputerGroups) UpdateSmartByIDV3(ctx context.Context, id string, request *RequestSmartGroupV3) (*ResourceSmartGroupV3, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateSmartGroupV3(request); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartComputerGroupsV3, id)

	var result ResourceSmartGroupV3
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSmartByIDV3 removes the specified smart group by ID.
// URL: DELETE /api/v3/computer-groups/smart-groups/{id}
func (s *ComputerGroups) DeleteSmartByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartComputerGroupsV3, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetSmartGroupMembershipByIDV3 returns the computer IDs that are members of the smart group.
// URL: GET /api/v3/computer-groups/smart-group-membership/{id}
func (s *ComputerGroups) GetSmartGroupMembershipByIDV3(ctx context.Context, id string) (*SmartGroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartComputerGroupMembershipV3, id)

	var result SmartGroupMembershipResponse
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// Static Groups CRUD (V3) — Jamf Pro 11.28, replaces the V2 surface.
// -----------------------------------------------------------------------------

// ListStaticV3 returns all static computer groups.
// URL: GET /api/v3/computer-groups/static-groups
func (s *ComputerGroups) ListStaticV3(ctx context.Context, rsqlQuery map[string]string) (*ListStaticV3Response, *resty.Response, error) {
	var result ListStaticV3Response

	endpoint := constants.EndpointJamfProStaticComputerGroupsV3

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceStaticGroupV3
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil {
		if bodyBytes := resp.Bytes(); len(bodyBytes) > 0 {
			var pageResp struct {
				TotalCount int `json:"totalCount"`
			}
			if err := json.Unmarshal(bodyBytes, &pageResp); err == nil {
				result.TotalCount = pageResp.TotalCount
			}
		}
	}
	if result.TotalCount == 0 {
		result.TotalCount = len(result.Results)
	}

	return &result, resp, nil
}

// GetStaticByIDV3 returns the specified static group by ID.
// URL: GET /api/v3/computer-groups/static-groups/{id}
func (s *ComputerGroups) GetStaticByIDV3(ctx context.Context, id string) (*ResourceStaticGroupV3, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroupsV3, id)

	var result ResourceStaticGroupV3
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateStaticV3 creates a new static computer group. Assignment IDs are
// deduplicated before sending (the API treats assignments as a set).
// URL: POST /api/v3/computer-groups/static-groups
func (s *ComputerGroups) CreateStaticV3(ctx context.Context, request *RequestStaticGroupV3) (*CreateStaticResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateStaticGroupV3(request); err != nil {
		return nil, nil, err
	}
	request.Assignments = smartgroupvalidation.DedupeStrings(request.Assignments)
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var result CreateStaticResponse
	endpoint := constants.EndpointJamfProStaticComputerGroupsV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateStaticByIDV3 updates the specified static group by ID. Assignment IDs
// are deduplicated before sending.
// URL: PUT /api/v3/computer-groups/static-groups/{id}
func (s *ComputerGroups) UpdateStaticByIDV3(ctx context.Context, id string, request *RequestStaticGroupV3) (*ResourceStaticGroupV3, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateStaticGroupV3(request); err != nil {
		return nil, nil, err
	}
	request.Assignments = smartgroupvalidation.DedupeStrings(request.Assignments)
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroupsV3, id)

	var result ResourceStaticGroupV3
	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteStaticByIDV3 removes the specified static group by ID.
// URL: DELETE /api/v3/computer-groups/static-groups/{id}
func (s *ComputerGroups) DeleteStaticByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroupsV3, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
