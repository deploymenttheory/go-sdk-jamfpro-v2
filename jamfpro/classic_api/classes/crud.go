package classes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the classes-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findclasses
	Classes struct {
		client client.Client
	}
)

// NewService returns a new classes Service backed by the provided HTTP client.
func NewClasses(client client.Client) *Classes {
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

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

	var result ResourceClass

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicClasses, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

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

	var result ResourceClass

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicClasses, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

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

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicClasses)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

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

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicClasses, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

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

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicClasses, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&result).
		Put(endpoint)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
