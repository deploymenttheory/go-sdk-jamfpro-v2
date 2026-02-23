package usergroups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// UserGroupsServiceInterface defines the interface for Classic API user group operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
	UserGroupsServiceInterface interface {
		// List returns all user groups.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroups
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified user group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyid
		GetByID(ctx context.Context, id int) (*ResourceUserGroup, *interfaces.Response, error)

		// GetByName returns the specified user group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyname
		GetByName(ctx context.Context, name string) (*ResourceUserGroup, *interfaces.Response, error)

		// Create creates a new user group.
		//
		// Returns the created user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createusergroupbyid
		Create(ctx context.Context, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified user group by ID.
		//
		// Returns the updated user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyid
		UpdateByID(ctx context.Context, id int, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified user group by name.
		//
		// Returns the updated user group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyname
		UpdateByName(ctx context.Context, name string, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified user group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified user group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the user-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/usergroups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ UserGroupsServiceInterface = (*Service)(nil)

// NewService returns a new user groups Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - User Groups CRUD Operations
// -----------------------------------------------------------------------------

// List returns all user groups.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroups
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointUserGroups

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

// GetByID returns the specified user group by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceUserGroup, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserGroups, id)

	var out ResourceUserGroup

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

// GetByName returns the specified user group by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusergroupsbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceUserGroup, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserGroups, name)

	var out ResourceUserGroup

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

// Create creates a new user group.
//
// Returns the created user group ID only (Classic API behavior).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createusergroupbyid
func (s *Service) Create(ctx context.Context, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointUserGroups)

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

// UpdateByID updates the specified user group by ID.
//
// Returns the updated user group ID only (Classic API behavior).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserGroups, id)

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

// UpdateByName updates the specified user group by name.
//
// Returns the updated user group ID only (Classic API behavior).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateusergroupbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestUserGroup) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user group name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user group name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserGroups, name)

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

// DeleteByID removes the specified user group by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserGroups, id)

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

// DeleteByName removes the specified user group by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteusergroupbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("user group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserGroups, name)

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
