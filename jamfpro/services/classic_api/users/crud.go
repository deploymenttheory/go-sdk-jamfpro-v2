package users

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// UsersServiceInterface defines the interface for Classic API user operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/users
	UsersServiceInterface interface {
		// List returns all users.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusers
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified user by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyid
		GetByID(ctx context.Context, id int) (*ResourceUser, *interfaces.Response, error)

		// GetByName returns the specified user by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyname
		GetByName(ctx context.Context, name string) (*ResourceUser, *interfaces.Response, error)

		// GetByEmail returns users matching the specified email.
		// Note: Returns a list response even for a single match.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyemailaddress
		GetByEmail(ctx context.Context, email string) (*ListResponse, *interfaces.Response, error)

		// Create creates a new user.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createuserbyid
		Create(ctx context.Context, req *RequestUser) (*ResourceUser, *interfaces.Response, error)

		// UpdateByID updates the specified user by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyid
		UpdateByID(ctx context.Context, id int, req *RequestUser) (*ResourceUser, *interfaces.Response, error)

		// UpdateByName updates the specified user by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyname
		UpdateByName(ctx context.Context, name string, req *RequestUser) (*ResourceUser, *interfaces.Response, error)

		// UpdateByEmail updates the specified user by email.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserbyemail
		UpdateByEmail(ctx context.Context, email string, req *RequestUser) (*ResourceUser, *interfaces.Response, error)

		// DeleteByID removes the specified user by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified user by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)

		// DeleteByEmail removes the specified user by email.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyemail
		DeleteByEmail(ctx context.Context, email string) (*interfaces.Response, error)
	}

	// Service handles communication with the users-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/users
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ UsersServiceInterface = (*Service)(nil)

// NewService returns a new users Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Users CRUD Operations
// -----------------------------------------------------------------------------

// List returns all users.
//
// URL: GET /JSSResource/users
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusers
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointUsers

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

// GetByID returns the specified user by ID.
//
// URL: GET /JSSResource/users/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceUser, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUsers, id)

	var out ResourceUser

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

// GetByName returns the specified user by name.
//
// URL: GET /JSSResource/users/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceUser, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUsers, url.PathEscape(name))

	var out ResourceUser

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

// GetByEmail returns users matching the specified email.
// Note: The API returns a list response (<users>) even for a single match.
// Because email addresses may not be unique, this operation may return a list of users.
//
// URL: GET /JSSResource/users/email/{email}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findusersbyemailaddress
func (s *Service) GetByEmail(ctx context.Context, email string) (*ListResponse, *interfaces.Response, error) {
	if email == "" {
		return nil, nil, fmt.Errorf("user email cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/email/%s", EndpointUsers, url.PathEscape(email))

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

// Create creates a new user.
//
// URL: POST /JSSResource/users/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createuserbyid
func (s *Service) Create(ctx context.Context, req *RequestUser) (*ResourceUser, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointUsers)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*RequestUser
	}{
		RequestUser: req,
	}

	var out ResourceUser

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, &requestBody, headers, &out)
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
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestUser) (*ResourceUser, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUsers, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*RequestUser
	}{
		RequestUser: req,
	}

	var out ResourceUser

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, &requestBody, headers, &out)
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
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestUser) (*ResourceUser, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUsers, url.PathEscape(name))

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*RequestUser
	}{
		RequestUser: req,
	}

	var out ResourceUser

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, &requestBody, headers, &out)
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
func (s *Service) UpdateByEmail(ctx context.Context, email string, req *RequestUser) (*ResourceUser, *interfaces.Response, error) {
	if email == "" {
		return nil, nil, fmt.Errorf("user email cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user name is required in request")
	}

	endpoint := fmt.Sprintf("%s/email/%s", EndpointUsers, url.PathEscape(email))

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*RequestUser
	}{
		RequestUser: req,
	}

	var out ResourceUser

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, &requestBody, headers, &out)
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
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUsers, id)

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

// DeleteByName removes the specified user by name.
//
// URL: DELETE /JSSResource/users/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("user name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUsers, url.PathEscape(name))

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

// DeleteByEmail removes the specified user by email.
//
// URL: DELETE /JSSResource/users/email/{email}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserbyemail
func (s *Service) DeleteByEmail(ctx context.Context, email string) (*interfaces.Response, error) {
	if email == "" {
		return nil, fmt.Errorf("user email cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/email/%s", EndpointUsers, url.PathEscape(email))

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
