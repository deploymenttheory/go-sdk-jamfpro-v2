package smart_computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SmartComputerGroupsServiceInterface defines the interface for smart computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	SmartComputerGroupsServiceInterface interface {
		// List returns a paginated list of all smart computer groups.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
		List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups-id
		GetByID(ctx context.Context, id string) (*ResourceSmartGroup, *resty.Response, error)

		// GetByName returns a smart computer group by name (client-side filter over list).
		GetByName(ctx context.Context, name string) (*ListItem, *resty.Response, error)

		// GetMembership returns the computer IDs that are members of the specified smart group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-group-membership-id
		GetMembership(ctx context.Context, id string) (*MembershipResponse, *resty.Response, error)

		// Create creates a new smart computer group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-smart-groups
		Create(ctx context.Context, request *RequestSmartGroup) (*CreateResponse, *resty.Response, error)

		// UpdateByID updates the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-smart-groups-id
		UpdateByID(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *resty.Response, error)

		// DeleteByID removes the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-smart-groups-id
		DeleteByID(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the smart computer groups methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	SmartComputerGroups struct {
		client interfaces.HTTPClient
	}
)

var _ SmartComputerGroupsServiceInterface = (*SmartComputerGroups)(nil)

func NewSmartComputerGroups(client interfaces.HTTPClient) *SmartComputerGroups {
	return &SmartComputerGroups{client: client}
}

// -----------------------------------------------------------------------------
// CRUD Operations
// -----------------------------------------------------------------------------

// List returns a paginated list of all smart computer groups.
// URL: GET /api/v2/computer-groups/smart-groups
func (s *SmartComputerGroups) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		result.Results = []ListItem{}

		var pageResults []ListItem
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	endpoint := EndpointSmartGroupsV2
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list smart computer groups: %w", err)
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

// GetByID returns the specified smart computer group by ID.
// URL: GET /api/v2/computer-groups/smart-groups/{id}
func (s *SmartComputerGroups) GetByID(ctx context.Context, id string) (*ResourceSmartGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	var result ResourceSmartGroup

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns a smart computer group by name.
func (s *SmartComputerGroups) GetByName(ctx context.Context, name string) (*ListItem, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("smart computer group name is required")
	}

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf("name==\"%s\"", name),
	}

	list, resp, err := s.List(ctx, rsqlQuery)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list smart computer groups: %w", err)
	}

	if len(list.Results) == 0 {
		return nil, resp, fmt.Errorf("smart computer group with name %q not found", name)
	}

	return &list.Results[0], resp, nil
}

// GetMembership returns the computer IDs that are members of the specified smart group.
// URL: GET /api/v2/computer-groups/smart-group-membership/{id}
func (s *SmartComputerGroups) GetMembership(ctx context.Context, id string) (*MembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupMembershipV2, id)

	var result MembershipResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new smart computer group.
// URL: POST /api/v2/computer-groups/smart-groups
func (s *SmartComputerGroups) Create(ctx context.Context, request *RequestSmartGroup) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointSmartGroupsV2
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified smart computer group by ID.
// URL: PUT /api/v2/computer-groups/smart-groups/{id}
func (s *SmartComputerGroups) UpdateByID(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	var result ResourceSmartGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified smart computer group by ID.
// URL: DELETE /api/v2/computer-groups/smart-groups/{id}
func (s *SmartComputerGroups) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
