package users

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the users-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/users
	Users struct {
		client client.Client
	}
)

// NewService returns a new users Service backed by the provided HTTP client.
func NewUsers(client client.Client) *Users {
	return &Users{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Users CRUD Operations
// -----------------------------------------------------------------------------

// List returns all users.
//
// URL: GET /JSSResource/users
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusers
func (s *Users) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicUsers

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

// GetByID returns the specified user by ID.
//
// URL: GET /JSSResource/users/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyid
func (s *Users) GetByID(ctx context.Context, id int) (*ResourceUser, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user ID must be a positive integer")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicUsers, id)

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

// GetByName returns the specified user by name.
//
// URL: GET /JSSResource/users/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyname
func (s *Users) GetByName(ctx context.Context, name string) (*ResourceUser, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user name cannot be empty")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicUsers, url.PathEscape(name))

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

// GetByEmail returns users matching the specified email.
// Note: The API returns a list response (<users>) even for a single match.
// Because email addresses may not be unique, this operation may return a list of users.
//
// URL: GET /JSSResource/users/email/{email}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyemailaddress
func (s *Users) GetByEmail(ctx context.Context, email string) (*ListResponse, *resty.Response, error) {
	if email == "" {
		return nil, nil, fmt.Errorf("user email cannot be empty")
	}

	var out ListResponse

	endpoint := fmt.Sprintf("%s/email/%s", constants.EndpointClassicUsers, url.PathEscape(email))

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

// Create creates a new user.
//
// URL: POST /JSSResource/users/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createuserbyid
func (s *Users) Create(ctx context.Context, req *RequestUser) (*ResourceUser, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicUsers)

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

// UpdateByID updates the specified user by ID.
//
// URL: PUT /JSSResource/users/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyid
func (s *Users) UpdateByID(ctx context.Context, id int, req *RequestUser) (*ResourceUser, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicUsers, id)

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

// UpdateByName updates the specified user by name.
//
// URL: PUT /JSSResource/users/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyname
func (s *Users) UpdateByName(ctx context.Context, name string, req *RequestUser) (*ResourceUser, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required in request")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicUsers, url.PathEscape(name))

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

// UpdateByEmail updates the specified user by email.
//
// URL: PUT /JSSResource/users/email/{email}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyemail
func (s *Users) UpdateByEmail(ctx context.Context, email string, req *RequestUser) (*ResourceUser, *resty.Response, error) {
	if email == "" {
		return nil, nil, fmt.Errorf("user email cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required in request")
	}

	var out ResourceUser

	endpoint := fmt.Sprintf("%s/email/%s", constants.EndpointClassicUsers, url.PathEscape(email))

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

// DeleteByID removes the specified user by ID.
//
// URL: DELETE /JSSResource/users/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyid
func (s *Users) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicUsers, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteByName removes the specified user by name.
//
// URL: DELETE /JSSResource/users/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyname
func (s *Users) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("user name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicUsers, url.PathEscape(name))

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}
	return resp, nil
}

// DeleteByEmail removes the specified user by email.
//
// URL: DELETE /JSSResource/users/email/{email}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyemail
func (s *Users) DeleteByEmail(ctx context.Context, email string) (*resty.Response, error) {
	if email == "" {
		return nil, fmt.Errorf("user email cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/email/%s", constants.EndpointClassicUsers, url.PathEscape(email))

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}
	return resp, nil
}
