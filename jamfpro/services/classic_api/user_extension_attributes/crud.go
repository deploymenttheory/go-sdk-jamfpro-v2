package user_extension_attributes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// UserExtensionAttributesServiceInterface defines the interface for Classic API user extension attribute operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/userextensionattributes
	UserExtensionAttributesServiceInterface interface {
		// List returns all user extension attributes.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributes
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified user extension attribute by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributesbyid
		GetByID(ctx context.Context, id int) (*ResourceUserExtensionAttribute, *interfaces.Response, error)

		// GetByName returns the specified user extension attribute by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributesbyname
		GetByName(ctx context.Context, name string) (*ResourceUserExtensionAttribute, *interfaces.Response, error)

		// Create creates a new user extension attribute.
		//
		// Returns the created user extension attribute (full resource).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createuserextensionattributebyid
		Create(ctx context.Context, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error)

		// UpdateByID updates the specified user extension attribute by ID.
		//
		// Returns the updated user extension attribute (full resource).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserextensionattributebyid
		UpdateByID(ctx context.Context, id int, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error)

		// UpdateByName updates the specified user extension attribute by name.
		//
		// Returns the updated user extension attribute (full resource).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserextensionattributebyname
		UpdateByName(ctx context.Context, name string, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error)

		// DeleteByID removes the specified user extension attribute by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserextensionattributebyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified user extension attribute by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserextensionattributebyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the user-extension-attributes-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/userextensionattributes
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ UserExtensionAttributesServiceInterface = (*Service)(nil)

// NewService returns a new user extension attributes Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - User Extension Attributes CRUD Operations
// -----------------------------------------------------------------------------

// List returns all user extension attributes.
//
// URL: GET /JSSResource/userextensionattributes
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributes
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointUserExtensionAttributes

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

// GetByID returns the specified user extension attribute by ID.
//
// URL: GET /JSSResource/userextensionattributes/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributesbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceUserExtensionAttribute, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user extension attribute ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserExtensionAttributes, id)

	var out ResourceUserExtensionAttribute

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

// GetByName returns the specified user extension attribute by name.
//
// URL: GET /JSSResource/userextensionattributes/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finduserextensionattributesbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceUserExtensionAttribute, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user extension attribute name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserExtensionAttributes, name)

	var out ResourceUserExtensionAttribute

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

// Create creates a new user extension attribute.
//
// Returns the created user extension attribute (full resource).
//
// URL: POST /JSSResource/userextensionattributes/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createuserextensionattributebyid
func (s *Service) Create(ctx context.Context, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user extension attribute name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointUserExtensionAttributes)

	var out ResourceUserExtensionAttribute

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

// UpdateByID updates the specified user extension attribute by ID.
//
// Returns the updated user extension attribute (full resource).
//
// URL: PUT /JSSResource/userextensionattributes/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserextensionattributebyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("user extension attribute ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user extension attribute name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserExtensionAttributes, id)

	var out ResourceUserExtensionAttribute

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

// UpdateByName updates the specified user extension attribute by name.
//
// Returns the updated user extension attribute (full resource).
//
// URL: PUT /JSSResource/userextensionattributes/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateuserextensionattributebyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestUserExtensionAttribute) (*ResourceUserExtensionAttribute, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("user extension attribute name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("user extension attribute name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserExtensionAttributes, name)

	var out ResourceUserExtensionAttribute

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

// DeleteByID removes the specified user extension attribute by ID.
//
// URL: DELETE /JSSResource/userextensionattributes/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserextensionattributebyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("user extension attribute ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointUserExtensionAttributes, id)

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

// DeleteByName removes the specified user extension attribute by name.
//
// URL: DELETE /JSSResource/userextensionattributes/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteuserextensionattributebyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("user extension attribute name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointUserExtensionAttributes, name)

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
