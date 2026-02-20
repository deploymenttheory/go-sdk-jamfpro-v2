package computer_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerGroupsServiceInterface defines the interface for computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	ComputerGroupsServiceInterface interface {
		// ListSmartV2 returns all smart computer groups (Get Smart Computer Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
		ListSmartV2(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *interfaces.Response, error)

		// GetSmartByIDV2 returns the specified smart computer group by ID (Get specified Smart Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups-id
		GetSmartByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error)

		// CreateSmartV2 creates a new smart computer group (Create Smart Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-smart-groups
		CreateSmartV2(ctx context.Context, request *RequestSmartGroup) (*CreateSmartResponse, *interfaces.Response, error)

		// UpdateSmartV2 updates the specified smart computer group by ID (Update specified Smart Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-smart-groups-id
		UpdateSmartV2(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error)

		// DeleteSmartV2 removes the specified smart computer group by ID (Remove specified Smart Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-smart-groups-id
		DeleteSmartV2(ctx context.Context, id string) (*interfaces.Response, error)

		// ListStaticV2 returns all static computer groups (Get Static Computer Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
		ListStaticV2(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *interfaces.Response, error)

		// GetStaticByIDV2 returns the specified static computer group by ID (Get specified Static Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups-id
		GetStaticByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error)

		// CreateStaticV2 creates a new static computer group (Create Static Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-groups-static-groups
		CreateStaticV2(ctx context.Context, request *RequestStaticGroup) (*CreateStaticResponse, *interfaces.Response, error)

		// UpdateStaticByIDV2 updates the specified static computer group by ID (Update specified Static Computer Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-computer-groups-static-groups-id
		UpdateStaticByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*ResourceStaticGroup, *interfaces.Response, error)

		// DeleteStaticByIDV2 removes the specified static computer group by ID (Remove specified Static Computer Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-groups-static-groups-id
		DeleteStaticByIDV2(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the computer groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerGroupsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Smart Groups CRUD
// -----------------------------------------------------------------------------

// ListSmartV2 returns all smart computer groups.
// URL: GET /api/v2/computer-groups/smart-groups
func (s *Service) ListSmartV2(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *interfaces.Response, error) {
	var result ListSmartResponse

	endpoint := EndpointSmartGroupsV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSmartByIDV2 returns the specified smart group by ID.
// URL: GET /api/v2/computer-groups/smart-groups/{id}
func (s *Service) GetSmartByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
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

// CreateSmartV2 creates a new smart computer group.
// URL: POST /api/v2/computer-groups/smart-groups
func (s *Service) CreateSmartV2(ctx context.Context, request *RequestSmartGroup) (*CreateSmartResponse, *interfaces.Response, error) {
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
func (s *Service) UpdateSmartV2(ctx context.Context, id string, request *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error) {
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
func (s *Service) DeleteSmartV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart group ID is required")
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

// -----------------------------------------------------------------------------
// Static Groups CRUD
// -----------------------------------------------------------------------------

// ListStaticV2 returns all static computer groups.
// URL: GET /api/v2/computer-groups/static-groups
func (s *Service) ListStaticV2(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *interfaces.Response, error) {
	var result ListStaticResponse

	endpoint := EndpointStaticGroupsV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetStaticByIDV2 returns the specified static group by ID.
// URL: GET /api/v2/computer-groups/static-groups/{id}
func (s *Service) GetStaticByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticGroup

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

// CreateStaticV2 creates a new static computer group with membership.
// URL: POST /api/v2/computer-groups/static-groups
func (s *Service) CreateStaticV2(ctx context.Context, request *RequestStaticGroup) (*CreateStaticResponse, *interfaces.Response, error) {
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

// UpdateStaticByIDV2 updates the membership of the specified static group (PATCH).
// URL: PATCH /api/v2/computer-groups/static-groups/{id}
func (s *Service) UpdateStaticByIDV2(ctx context.Context, id string, request *RequestStaticGroup) (*ResourceStaticGroup, *interfaces.Response, error) {
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

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteStaticByIDV2 removes the specified static group by ID.
// URL: DELETE /api/v2/computer-groups/static-groups/{id}
func (s *Service) DeleteStaticByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

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
