package static_computer_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// StaticComputerGroupsServiceInterface defines the interface for static computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
	StaticComputerGroupsServiceInterface interface {
		// ListV2 returns all static computer groups.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
		ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV2 returns the specified static computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups-id
		GetByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error)

		// GetByNameV2 returns the specified static computer group by name.
		GetByNameV2(ctx context.Context, name string) (*ResourceStaticGroup, *interfaces.Response, error)

		// CreateV2 creates a new static computer group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-static-groups
		CreateV2(ctx context.Context, request *RequestStaticGroup) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV2 updates the specified static computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-static-groups-id
		UpdateByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*RequestStaticGroup, *interfaces.Response, error)

		// DeleteByIDV2 removes the specified static computer group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-static-groups-id
		DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the static computer groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ StaticComputerGroupsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV2 returns all static computer groups.
// URL: GET /api/v2/computer-groups/static-groups
func (s *Service) ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointStaticGroupsV2

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
				var group ResourceStaticGroup
				if err := mapstructure.Decode(item, &group); err != nil {
					return fmt.Errorf("failed to decode static computer group: %w", err)
				}
				result.Results = append(result.Results, group)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list static computer groups: %w", err)
	}

	return &result, resp, nil
}

// GetByIDV2 returns the specified static computer group by ID.
// URL: GET /api/v2/computer-groups/static-groups/{id}
func (s *Service) GetByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static computer group ID is required")
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

// GetByNameV2 returns the specified static computer group by name.
func (s *Service) GetByNameV2(ctx context.Context, name string) (*ResourceStaticGroup, *interfaces.Response, error) {
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
func (s *Service) CreateV2(ctx context.Context, request *RequestStaticGroup) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var result CreateResponse

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

// UpdateByIDV2 updates the specified static computer group by ID.
// URL: PUT /api/v2/computer-groups/static-groups/{id}
func (s *Service) UpdateByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*RequestStaticGroup, *interfaces.Response, error) {
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

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

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

// DeleteByIDV2 removes the specified static computer group by ID.
// URL: DELETE /api/v2/computer-groups/static-groups/{id}
func (s *Service) DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static computer group ID is required")
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
