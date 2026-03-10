package static_user_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the static-user-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
	StaticUserGroups struct {
		client client.Client
	}
)

// NewService returns a new static user groups Service backed by the provided HTTP client.
func NewStaticUserGroups(client client.Client) *StaticUserGroups {
	return &StaticUserGroups{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Static User Groups CRUD Operations
// -----------------------------------------------------------------------------

// List returns all user groups (both smart and static).
//
// URL: GET /JSSResource/usergroups
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroups
func (s *StaticUserGroups) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicStaticUserGroups

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByID returns the specified static user group by ID.
//
// URL: GET /JSSResource/usergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyid
func (s *StaticUserGroups) GetByID(ctx context.Context, id int) (*ResourceStaticUserGroup, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user group ID must be a positive integer")
	}

	var out ResourceStaticUserGroup

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicStaticUserGroups, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByName returns the specified static user group by name.
//
// URL: GET /JSSResource/usergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyname
func (s *StaticUserGroups) GetByName(ctx context.Context, name string) (*ResourceStaticUserGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user group name is required")
	}

	var out ResourceStaticUserGroup

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicStaticUserGroups, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new static user group.
//
// Returns the created user group ID only (Classic API behavior).
//
// URL: POST /JSSResource/usergroups/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createusergroupbyid
func (s *StaticUserGroups) Create(ctx context.Context, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicStaticUserGroups)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified static user group by ID.
//
// Returns the updated user group ID only (Classic API behavior).
//
// URL: PUT /JSSResource/usergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyid
func (s *StaticUserGroups) UpdateByID(ctx context.Context, id int, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user group ID must be a positive integer")
	}
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicStaticUserGroups, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified static user group by name.
//
// Returns the updated user group ID only (Classic API behavior).
//
// URL: PUT /JSSResource/usergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyname
func (s *StaticUserGroups) UpdateByName(ctx context.Context, name string, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user group name is required")
	}
	if err := ValidateRequest(req); err != nil {
		return nil, nil, err
	}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicStaticUserGroups, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified static user group by ID.
//
// URL: DELETE /JSSResource/usergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyid
func (s *StaticUserGroups) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicStaticUserGroups, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified static user group by name.
//
// URL: DELETE /JSSResource/usergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyname
func (s *StaticUserGroups) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("user group name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicStaticUserGroups, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
