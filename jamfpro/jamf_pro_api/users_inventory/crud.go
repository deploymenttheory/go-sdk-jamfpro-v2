package users_inventory

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the user inventory methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-users
	UsersInventory struct {
		client client.Client
	}
)

// NewUsersInventory returns a new UsersInventory service.
func NewUsersInventory(client client.Client) *UsersInventory {
	return &UsersInventory{client: client}
}

// ListV1 retrieves a paginated list of users with optional RSQL filtering.
// URL: GET /api/v1/users
// rsqlQuery supports: filter (RSQL), sort, page, page-size, platform (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-users
func (s *UsersInventory) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListUsersResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUsersInventoryV1

	var result ListUsersResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list users: %w", err)
	}

	return &result, resp, nil
}

// GetByIDV1 retrieves a single user by their ID.
// URL: GET /api/v1/users/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-users-id
func (s *UsersInventory) GetByIDV1(ctx context.Context, id string) (*ResourceUser, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("user ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProUsersInventoryV1, id)

	var result ResourceUser

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new user in the inventory.
// URL: POST /api/v1/users
// https://developer.jamf.com/jamf-pro/reference/post_v1-users
func (s *UsersInventory) CreateV1(ctx context.Context, request *RequestUserInventory) (*CreateUserResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProUsersInventoryV1

	var result CreateUserResponse

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

// UpdateByIDV1 updates an existing user in the inventory by ID.
// URL: PUT /api/v1/users/{id}
// Response: 204 No Content
// https://developer.jamf.com/jamf-pro/reference/put_v1-users-id
func (s *UsersInventory) UpdateByIDV1(ctx context.Context, id string, request *RequestUserInventory) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProUsersInventoryV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Put(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByIDV1 deletes a user from the inventory by ID.
// URL: DELETE /api/v1/users/{id}
// Response: 204 No Content
// https://developer.jamf.com/jamf-pro/reference/delete_v1-users-id
func (s *UsersInventory) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProUsersInventoryV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
