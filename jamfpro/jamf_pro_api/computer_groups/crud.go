package computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ComputerGroupsServiceInterface defines the interface for computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	ComputerGroupsServiceInterface interface {
		// ListSmartV2 returns all smart computer groups (Get Smart Computer Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
		ListSmartV2(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *resty.Response, error)

		// GetSmartByIDV2 returns the specified smart computer group by ID (Get specified Smart Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups-id
		GetSmartByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *resty.Response, error)

		// CreateSmartV2 creates a new smart computer group (Create Smart Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-smart-groups
		CreateSmartV2(ctx context.Context, request *RequestSmartGroup) (*CreateSmartResponse, *resty.Response, error)

		// UpdateSmartV2 updates the specified smart computer group by ID (Update specified Smart Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-smart-groups-id
		UpdateSmartV2(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *resty.Response, error)

		// DeleteSmartV2 removes the specified smart computer group by ID (Remove specified Smart Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-smart-groups-id
		DeleteSmartV2(ctx context.Context, id string) (*resty.Response, error)

		// ListStaticV2 returns all static computer groups (Get Static Computer Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
		ListStaticV2(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *resty.Response, error)

		// GetStaticByIDV2 returns the specified static computer group by ID (Get specified Static Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups-id
		GetStaticByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *resty.Response, error)

		// CreateStaticV2 creates a new static computer group (Create Static Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-static-groups
		CreateStaticV2(ctx context.Context, request *RequestStaticGroup) (*CreateStaticResponse, *resty.Response, error)

		// UpdateStaticByIDV2 updates the specified static computer group by ID (Update specified Static Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-static-groups-id
		UpdateStaticByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*ResourceStaticGroup, *resty.Response, error)

		// DeleteStaticByIDV2 removes the specified static computer group by ID (Remove specified Static Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-static-groups-id
		DeleteStaticByIDV2(ctx context.Context, id string) (*resty.Response, error)

		// ListAllV1 returns a simple list of all computer groups (id, name, description, smartGroup).
		// No pagination, just a simple list.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-groups
		ListAllV1(ctx context.Context) ([]ResourceGroupV1, *resty.Response, error)

		// GetSmartGroupMembershipByIDV2 returns the membership (computer IDs) for a smart group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-group-membership-id
		GetSmartGroupMembershipByIDV2(ctx context.Context, id string) (*SmartGroupMembershipResponse, *resty.Response, error)
	}

	// Service handles communication with the computer groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	ComputerGroups struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerGroupsServiceInterface = (*ComputerGroups)(nil)

func NewComputerGroups(client interfaces.HTTPClient) *ComputerGroups {
	return &ComputerGroups{client: client}
}

// -----------------------------------------------------------------------------
// Smart Groups CRUD
// -----------------------------------------------------------------------------

// ListSmartV2 returns all smart computer groups.
// URL: GET /api/v2/computer-groups/smart-groups
func (s *ComputerGroups) ListSmartV2(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *resty.Response, error) {
	var result ListSmartResponse

	endpoint := EndpointSmartGroupsV2

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceSmartGroup
		result.Results = []ResourceSmartGroup{}

		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
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

// GetSmartByIDV2 returns the specified smart group by ID.
// URL: GET /api/v2/computer-groups/smart-groups/{id}
func (s *ComputerGroups) GetSmartByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
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

// CreateSmartV2 creates a new smart computer group.
// URL: POST /api/v2/computer-groups/smart-groups
func (s *ComputerGroups) CreateSmartV2(ctx context.Context, request *RequestSmartGroup) (*CreateSmartResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateSmartResponse

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

// UpdateSmartV2 updates the specified smart group by ID.
// URL: PUT /api/v2/computer-groups/smart-groups/{id}
func (s *ComputerGroups) UpdateSmartV2(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *resty.Response, error) {
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

// DeleteSmartV2 removes the specified smart group by ID.
// URL: DELETE /api/v2/computer-groups/smart-groups/{id}
func (s *ComputerGroups) DeleteSmartV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart group ID is required")
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

// -----------------------------------------------------------------------------
// Static Groups CRUD
// -----------------------------------------------------------------------------

// ListStaticV2 returns all static computer groups.
// URL: GET /api/v2/computer-groups/static-groups
func (s *ComputerGroups) ListStaticV2(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *resty.Response, error) {
	var result ListStaticResponse

	result.Results = []ResourceStaticGroup{}

	endpoint := EndpointStaticGroupsV2

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceStaticGroup
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
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

// GetStaticByIDV2 returns the specified static group by ID.
// URL: GET /api/v2/computer-groups/static-groups/{id}
func (s *ComputerGroups) GetStaticByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticGroup

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateStaticV2 creates a new static computer group with membership.
// URL: POST /api/v2/computer-groups/static-groups
func (s *ComputerGroups) CreateStaticV2(ctx context.Context, request *RequestStaticGroup) (*CreateStaticResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateStaticResponse

	endpoint := EndpointStaticGroupsV2

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

// UpdateStaticByIDV2 updates the specified static group by ID.
// URL: PUT /api/v2/computer-groups/static-groups/{id}
func (s *ComputerGroups) UpdateStaticByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*ResourceStaticGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticGroup

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

// DeleteStaticByIDV2 removes the specified static group by ID.
// URL: DELETE /api/v2/computer-groups/static-groups/{id}
func (s *ComputerGroups) DeleteStaticByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// V1 and Additional V2 Endpoints
// -----------------------------------------------------------------------------

// ListAllV1 returns a simple list of all computer groups (id, name, description, smartGroup).
// URL: GET /api/v1/computer-groups
func (s *ComputerGroups) ListAllV1(ctx context.Context) ([]ResourceGroupV1, *resty.Response, error) {
	var result []ResourceGroupV1

	endpoint := EndpointComputerGroupsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetSmartGroupMembershipByIDV2 returns the membership (computer IDs) for a smart group.
// URL: GET /api/v2/computer-groups/smart-group-membership/{id}
func (s *ComputerGroups) GetSmartGroupMembershipByIDV2(ctx context.Context, id string) (*SmartGroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupMembershipV2, id)

	var result SmartGroupMembershipResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
