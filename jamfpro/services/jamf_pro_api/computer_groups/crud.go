package computer_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// ComputerGroupsServiceInterface defines the interface for computer group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
	ComputerGroupsServiceInterface interface {
		// Smart groups
		ListSmartGroupsV2(ctx context.Context, queryParams map[string]string) (*ListSmartResponse, *interfaces.Response, error)
		GetSmartGroupByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error)
		CreateSmartGroupV2(ctx context.Context, req *RequestSmartGroup) (*CreateSmartResponse, *interfaces.Response, error)
		UpdateSmartGroupV2(ctx context.Context, id string, req *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error)
		DeleteSmartGroupV2(ctx context.Context, id string) (*interfaces.Response, error)

		// Static groups
		ListStaticGroupsV2(ctx context.Context, queryParams map[string]string) (*ListStaticResponse, *interfaces.Response, error)
		GetStaticGroupByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error)
		CreateStaticGroupV2(ctx context.Context, req *RequestStaticGroup) (*CreateStaticResponse, *interfaces.Response, error)
		UpdateStaticGroupByIDV2(ctx context.Context, id string, req *RequestStaticGroup) (*ResourceStaticGroup, *interfaces.Response, error)
		DeleteStaticGroupByIDV2(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the computer groups-related methods of the Jamf Pro API.
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

// ListSmartGroupsV2 returns all smart computer groups.
// URL: GET /api/v2/computer-groups/smart-groups
func (s *Service) ListSmartGroupsV2(ctx context.Context, queryParams map[string]string) (*ListSmartResponse, *interfaces.Response, error) {
	var result ListSmartResponse

	resp, err := s.client.Get(ctx, EndpointSmartGroupsV2, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetSmartGroupByIDV2 returns the specified smart group by ID.
// URL: GET /api/v2/computer-groups/smart-groups/{id}
func (s *Service) GetSmartGroupByIDV2(ctx context.Context, id string) (*ResourceSmartGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	var result ResourceSmartGroup

	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateSmartGroupV2 creates a new smart computer group.
// URL: POST /api/v2/computer-groups/smart-groups
func (s *Service) CreateSmartGroupV2(ctx context.Context, req *RequestSmartGroup) (*CreateSmartResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateSmartResponse

	resp, err := s.client.Post(ctx, EndpointSmartGroupsV2, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSmartGroupV2 updates the specified smart group by ID.
// URL: PUT /api/v2/computer-groups/smart-groups/{id}
func (s *Service) UpdateSmartGroupV2(ctx context.Context, id string, req *RequestSmartGroup) (*ResourceSmartGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	var result ResourceSmartGroup

	resp, err := s.client.Put(ctx, endpoint, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSmartGroupV2 removes the specified smart group by ID.
// URL: DELETE /api/v2/computer-groups/smart-groups/{id}
func (s *Service) DeleteSmartGroupV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV2, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// Static Groups CRUD
// -----------------------------------------------------------------------------

// ListStaticGroupsV2 returns all static computer groups.
// URL: GET /api/v2/computer-groups/static-groups
func (s *Service) ListStaticGroupsV2(ctx context.Context, queryParams map[string]string) (*ListStaticResponse, *interfaces.Response, error) {
	var result ListStaticResponse

	resp, err := s.client.Get(ctx, EndpointStaticGroupsV2, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetStaticGroupByIDV2 returns the specified static group by ID.
// URL: GET /api/v2/computer-groups/static-groups/{id}
func (s *Service) GetStaticGroupByIDV2(ctx context.Context, id string) (*ResourceStaticGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticGroup

	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateStaticGroupV2 creates a new static computer group with membership.
// URL: POST /api/v2/computer-groups/static-groups
func (s *Service) CreateStaticGroupV2(ctx context.Context, req *RequestStaticGroup) (*CreateStaticResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateStaticResponse

	resp, err := s.client.Post(ctx, EndpointStaticGroupsV2, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateStaticGroupByIDV2 updates the membership of the specified static group (PATCH).
// URL: PATCH /api/v2/computer-groups/static-groups/{id}
func (s *Service) UpdateStaticGroupByIDV2(ctx context.Context, id string, req *RequestStaticGroup) (*ResourceStaticGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticGroup

	resp, err := s.client.Patch(ctx, endpoint, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteStaticGroupByIDV2 removes the specified static group by ID.
// URL: DELETE /api/v2/computer-groups/static-groups/{id}
func (s *Service) DeleteStaticGroupByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
