package classes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ClassesServiceInterface defines the interface for Classic API class operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
	ClassesServiceInterface interface {
		// List returns all classes.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified class by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclassesbyid
		GetByID(ctx context.Context, id int) (*ResourceClass, *resty.Response, error)

		// GetByName returns the specified class by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclassesbyname
		GetByName(ctx context.Context, name string) (*ResourceClass, *resty.Response, error)

		// Create creates a new class.
		//
		// Returns the created class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createclassbyid
		Create(ctx context.Context, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified class by ID.
		//
		// Returns the updated class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateclassbyid
		UpdateByID(ctx context.Context, id int, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified class by name.
		//
		// Returns the updated class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateclassbyname
		UpdateByName(ctx context.Context, name string, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified class by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteclassbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified class by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteclassbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the classes-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
	Classes struct {
		client transport.HTTPClient
	}
)

var _ ClassesServiceInterface = (*Classes)(nil)

// NewService returns a new classes Service backed by the provided HTTP client.
func NewClasses(client transport.HTTPClient) *Classes {
	return &Classes{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Classes CRUD Operations
// -----------------------------------------------------------------------------

// List returns all classes.
// URL: GET /JSSResource/classes
// https://developer.jamf.com/jamf-pro/reference/findclasses
func (s *Classes) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicClasses

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified class by ID.
// URL: GET /JSSResource/classes/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findclassesbyid
func (s *Classes) GetByID(ctx context.Context, id int) (*ResourceClass, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("class ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicClasses, id)

	var result ResourceClass

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified class by name.
// URL: GET /JSSResource/classes/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findclassesbyname
func (s *Classes) GetByName(ctx context.Context, name string) (*ResourceClass, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("class name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicClasses, name)

	var result ResourceClass

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new class.
// URL: POST /JSSResource/classes/id/0
// Returns the created class ID only.
// https://developer.jamf.com/jamf-pro/reference/createclassbyid
func (s *Classes) Create(ctx context.Context, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicClasses)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified class by ID.
// URL: PUT /JSSResource/classes/id/{id}
// Returns the updated class ID only.
// https://developer.jamf.com/jamf-pro/reference/updateclassbyid
func (s *Classes) UpdateByID(ctx context.Context, id int, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("class ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicClasses, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified class by name.
// URL: PUT /JSSResource/classes/name/{name}
// Returns the updated class ID only.
// https://developer.jamf.com/jamf-pro/reference/updateclassbyname
func (s *Classes) UpdateByName(ctx context.Context, name string, req *RequestClass) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("class name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicClasses, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified class by ID.
// URL: DELETE /JSSResource/classes/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteclassbyid
func (s *Classes) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("class ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicClasses, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified class by name.
// URL: DELETE /JSSResource/classes/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteclassbyname
func (s *Classes) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("class name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicClasses, name)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
