package classes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ClassesServiceInterface defines the interface for Classic API class operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
	ClassesServiceInterface interface {
		// ListClasses returns all classes.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
		ListClasses(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetClassByID returns the specified class by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclassesbyid
		GetClassByID(ctx context.Context, id int) (*ResourceClass, *interfaces.Response, error)

		// GetClassByName returns the specified class by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclassesbyname
		GetClassByName(ctx context.Context, name string) (*ResourceClass, *interfaces.Response, error)

		// CreateClass creates a new class.
		//
		// Returns the created class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createclassbyid
		CreateClass(ctx context.Context, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateClassByID updates the specified class by ID.
		//
		// Returns the updated class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateclassbyid
		UpdateClassByID(ctx context.Context, id int, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateClassByName updates the specified class by name.
		//
		// Returns the updated class ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateclassbyname
		UpdateClassByName(ctx context.Context, name string, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteClassByID removes the specified class by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteclassbyid
		DeleteClassByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteClassByName removes the specified class by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteclassbyname
		DeleteClassByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the classes-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ClassesServiceInterface = (*Service)(nil)

// NewService returns a new classes Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Classes CRUD Operations
// -----------------------------------------------------------------------------

// ListClasses returns all classes.
// URL: GET /JSSResource/classes
// https://developer.jamf.com/jamf-pro/reference/findclasses
func (s *Service) ListClasses(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicClasses

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetClassByID returns the specified class by ID.
// URL: GET /JSSResource/classes/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findclassesbyid
func (s *Service) GetClassByID(ctx context.Context, id int) (*ResourceClass, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("class ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicClasses, id)

	var result ResourceClass

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetClassByName returns the specified class by name.
// URL: GET /JSSResource/classes/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findclassesbyname
func (s *Service) GetClassByName(ctx context.Context, name string) (*ResourceClass, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("class name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicClasses, name)

	var result ResourceClass

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateClass creates a new class.
// URL: POST /JSSResource/classes/id/0
// Returns the created class ID only.
// https://developer.jamf.com/jamf-pro/reference/createclassbyid
func (s *Service) CreateClass(ctx context.Context, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicClasses)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateClassByID updates the specified class by ID.
// URL: PUT /JSSResource/classes/id/{id}
// Returns the updated class ID only.
// https://developer.jamf.com/jamf-pro/reference/updateclassbyid
func (s *Service) UpdateClassByID(ctx context.Context, id int, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("class ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicClasses, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateClassByName updates the specified class by name.
// URL: PUT /JSSResource/classes/name/{name}
// Returns the updated class ID only.
// https://developer.jamf.com/jamf-pro/reference/updateclassbyname
func (s *Service) UpdateClassByName(ctx context.Context, name string, req *RequestClass) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("class name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicClasses, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteClassByID removes the specified class by ID.
// URL: DELETE /JSSResource/classes/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteclassbyid
func (s *Service) DeleteClassByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("class ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicClasses, id)

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

// DeleteClassByName removes the specified class by name.
// URL: DELETE /JSSResource/classes/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteclassbyname
func (s *Service) DeleteClassByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("class name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicClasses, name)

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
