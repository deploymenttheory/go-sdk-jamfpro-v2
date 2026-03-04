package restricted_software

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// RestrictedSoftwareServiceInterface defines the interface for Classic API restricted software operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/restrictedsoftware
	RestrictedSoftwareServiceInterface interface {
		// ListRestrictedSoftware returns all restricted software.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftware
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetRestrictedSoftwareByID returns the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyid
		GetByID(ctx context.Context, id int) (*ResourceRestrictedSoftware, *resty.Response, error)

		// GetRestrictedSoftwareByName returns the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyname
		GetByName(ctx context.Context, name string) (*ResourceRestrictedSoftware, *resty.Response, error)

		// CreateRestrictedSoftware creates new restricted software.
		//
		// Returns the created restricted software's assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createrestrictedsoftwarebyid
		Create(ctx context.Context, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateRestrictedSoftwareByID updates the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyid
		UpdateByID(ctx context.Context, id int, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateRestrictedSoftwareByName updates the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyname
		UpdateByName(ctx context.Context, name string, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteRestrictedSoftwareByID removes the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteRestrictedSoftwareByName removes the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the restricted software-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/restrictedsoftware
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ RestrictedSoftwareServiceInterface = (*Service)(nil)

// NewService returns a new restricted software Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Restricted Software CRUD Operations
// -----------------------------------------------------------------------------

// List returns all restricted software.
// URL: GET /JSSResource/restrictedsoftware
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftware
func (s *Service) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicRestrictedSoftware

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

// GetByID returns the specified restricted software by ID.
// URL: GET /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceRestrictedSoftware, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("restricted software ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRestrictedSoftware, id)

	var result ResourceRestrictedSoftware

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

// GetByName returns the specified restricted software by name.
// URL: GET /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceRestrictedSoftware, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("restricted software name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRestrictedSoftware, name)

	var result ResourceRestrictedSoftware

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

// CreateRestrictedSoftware creates new restricted software.
// URL: POST /JSSResource/restrictedsoftware/id/0
// Returns the created restricted software's assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createrestrictedsoftwarebyid
func (s *Service) Create(ctx context.Context, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicRestrictedSoftware)

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

// UpdateByID updates the specified restricted software by ID.
// URL: PUT /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("restricted software ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRestrictedSoftware, id)

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

// UpdateByName updates the specified restricted software by name.
// URL: PUT /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("restricted software name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRestrictedSoftware, name)

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

// DeleteByID removes the specified restricted software by ID.
// URL: DELETE /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("restricted software ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicRestrictedSoftware, id)

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

// DeleteByName removes the specified restricted software by name.
// URL: DELETE /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("restricted software name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicRestrictedSoftware, name)

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
