package static_user_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// StaticUserGroupsServiceInterface defines the interface for Classic API static user group operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
	StaticUserGroupsServiceInterface interface {
		// List returns all user groups (both smart and static).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroups
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified static user group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyid
		GetByID(ctx context.Context, id int) (*ResourceStaticUserGroup, *resty.Response, error)

		// GetByName returns the specified static user group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyname
		GetByName(ctx context.Context, name string) (*ResourceStaticUserGroup, *resty.Response, error)

		// Create creates a new static user group.
		//
		// Returns the created user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createusergroupbyid
		Create(ctx context.Context, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified static user group by ID.
		//
		// Returns the updated user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyid
		UpdateByID(ctx context.Context, id int, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified static user group by name.
		//
		// Returns the updated user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyname
		UpdateByName(ctx context.Context, name string, req *RequestStaticUserGroup) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified static user group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified static user group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the static-user-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
	StaticUserGroups struct {
		client transport.HTTPClient
	}
)

var _ StaticUserGroupsServiceInterface = (*StaticUserGroups)(nil)

// NewService returns a new static user groups Service backed by the provided HTTP client.
func NewStaticUserGroups(client transport.HTTPClient) *StaticUserGroups {
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
	endpoint := EndpointStaticUserGroups

	var out ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointStaticUserGroups, id)

	var out ResourceStaticUserGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointStaticUserGroups, name)

	var out ResourceStaticUserGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
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

	endpoint := fmt.Sprintf("%s/id/0", EndpointStaticUserGroups)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointStaticUserGroups, id)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointStaticUserGroups, name)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
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

	endpoint := fmt.Sprintf("%s/id/%d", EndpointStaticUserGroups, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
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

	endpoint := fmt.Sprintf("%s/name/%s", EndpointStaticUserGroups, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
