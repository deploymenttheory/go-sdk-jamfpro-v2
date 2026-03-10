package static_computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the static computer groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
	StaticComputerGroups struct {
		client client.Client
	}
)

func NewStaticComputerGroups(client client.Client) *StaticComputerGroups {
	return &StaticComputerGroups{client: client}
}

// ListV2 returns all static computer groups.
// URL: GET /api/v2/computer-groups/static-groups
func (s *StaticComputerGroups) ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse
	result.Results = []ResourceStaticGroup{}

	endpoint := constants.EndpointJamfProStaticComputerGroups2V2

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceStaticGroup
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list static computer groups: %w", err)
	}

	// Extract totalCount from response body if available
	bodyBytes := resp.Bytes()
	if resp != nil && len(bodyBytes) > 0 {
		var pageResp struct {
			TotalCount int `json:"totalCount"`
		}
		if err := json.Unmarshal(bodyBytes, &pageResp); err == nil {
			result.TotalCount = pageResp.TotalCount
		}
	}

	// Fallback: if totalCount wasn't extracted, use length of results
	if result.TotalCount == 0 {
		result.TotalCount = len(result.Results)
	}

	return &result, resp, nil
}

// GetByIDV2 returns the specified static computer group by ID.
// URL: GET /api/v2/computer-groups/static-groups/{id}
func (s *StaticComputerGroups) GetByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroups2V2, id)

	var result ResourceStaticGroup

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV2 returns the specified static computer group by name.
func (s *StaticComputerGroups) GetByNameV2(ctx context.Context, name string) (*ResourceStaticGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("static computer group name is required")
	}

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf("name==\"%s\"", name),
	}

	list, resp, err := s.ListV2(ctx, rsqlQuery)
	if err != nil {
		return nil, resp, err
	}

	if len(list.Results) == 0 {
		return nil, resp, fmt.Errorf("static computer group with name %q was not found", name)
	}

	return &list.Results[0], resp, nil
}

// CreateV2 creates a new static computer group.
// URL: POST /api/v2/computer-groups/static-groups
func (s *StaticComputerGroups) CreateV2(ctx context.Context, request *RequestStaticGroup) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProStaticComputerGroups2V2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 updates the specified static computer group by ID.
// URL: PUT /api/v2/computer-groups/static-groups/{id}
func (s *StaticComputerGroups) UpdateByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*RequestStaticGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var result RequestStaticGroup

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroups2V2, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV2 removes the specified static computer group by ID.
// URL: DELETE /api/v2/computer-groups/static-groups/{id}
func (s *StaticComputerGroups) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticComputerGroups2V2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
