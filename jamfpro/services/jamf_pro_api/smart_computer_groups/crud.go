package smart_computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// SmartComputerGroupsServiceInterface defines the interface for smart computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	SmartComputerGroupsServiceInterface interface {
		// List returns a paginated list of all smart computer groups.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
		List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups-id
		GetByID(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error)

		// GetByName returns a smart computer group by name (client-side filter over list).
		GetByName(ctx context.Context, name string) (*ListItem, *interfaces.Response, error)

		// GetMembership returns the computer IDs that are members of the specified smart group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-group-membership-id
		GetMembership(ctx context.Context, id string) (*MembershipResponse, *interfaces.Response, error)

		// Create creates a new smart computer group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-smart-groups
		Create(ctx context.Context, request *RequestSmartGroup) (*CreateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-smart-groups-id
		UpdateByID(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error)

		// DeleteByID removes the specified smart computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-smart-groups-id
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the smart computer groups methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SmartComputerGroupsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// CRUD Operations
// -----------------------------------------------------------------------------

// List returns a paginated list of all smart computer groups.
// URL: GET /api/v2/computer-groups/smart-groups
func (s *Service) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var group ListItem
				if err := mapstructure.Decode(item, &group); err != nil {
					return fmt.Errorf("failed to decode smart computer group: %w", err)
				}
				result.Results = append(result.Results, group)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, EndpointSmartGroupsV2, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list smart computer groups: %w", err)
	}

	return &result, resp, nil
}

// GetByID returns the specified smart computer group by ID.
// URL: GET /api/v2/computer-groups/smart-groups/{id}
func (s *Service) GetByID(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	var result ResourceSmartGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns a smart computer group by name.
func (s *Service) GetByName(ctx context.Context, name string) (*ListItem, *interfaces.Response, error) {
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
func (s *Service) GetMembership(ctx context.Context, id string) (*MembershipResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupMembershipV2, id)

	var result MembershipResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new smart computer group.
// URL: POST /api/v2/computer-groups/smart-groups
func (s *Service) Create(ctx context.Context, request *RequestSmartGroup) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointSmartGroupsV2, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified smart computer group by ID.
// URL: PUT /api/v2/computer-groups/smart-groups/{id}
func (s *Service) UpdateByID(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error) {
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
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart computer group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
