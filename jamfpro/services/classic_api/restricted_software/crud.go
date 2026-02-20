package restricted_software

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// RestrictedSoftwareServiceInterface defines the interface for Classic API restricted software operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/restrictedsoftware
	RestrictedSoftwareServiceInterface interface {
		// ListRestrictedSoftware returns all restricted software.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftware
		ListRestrictedSoftware(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetRestrictedSoftwareByID returns the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyid
		GetRestrictedSoftwareByID(ctx context.Context, id int) (*ResourceRestrictedSoftware, *interfaces.Response, error)

		// GetRestrictedSoftwareByName returns the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyname
		GetRestrictedSoftwareByName(ctx context.Context, name string) (*ResourceRestrictedSoftware, *interfaces.Response, error)

		// CreateRestrictedSoftware creates new restricted software.
		//
		// Returns the created restricted software's assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createrestrictedsoftwarebyid
		CreateRestrictedSoftware(ctx context.Context, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateRestrictedSoftwareByID updates the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyid
		UpdateRestrictedSoftwareByID(ctx context.Context, id int, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateRestrictedSoftwareByName updates the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyname
		UpdateRestrictedSoftwareByName(ctx context.Context, name string, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteRestrictedSoftwareByID removes the specified restricted software by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyid
		DeleteRestrictedSoftwareByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteRestrictedSoftwareByName removes the specified restricted software by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyname
		DeleteRestrictedSoftwareByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// ListRestrictedSoftware returns all restricted software.
// URL: GET /JSSResource/restrictedsoftware
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftware
func (s *Service) ListRestrictedSoftware(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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

// GetRestrictedSoftwareByID returns the specified restricted software by ID.
// URL: GET /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyid
func (s *Service) GetRestrictedSoftwareByID(ctx context.Context, id int) (*ResourceRestrictedSoftware, *interfaces.Response, error) {
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

// GetRestrictedSoftwareByName returns the specified restricted software by name.
// URL: GET /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findrestrictedsoftwarebyname
func (s *Service) GetRestrictedSoftwareByName(ctx context.Context, name string) (*ResourceRestrictedSoftware, *interfaces.Response, error) {
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
func (s *Service) CreateRestrictedSoftware(ctx context.Context, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateRestrictedSoftwareByID updates the specified restricted software by ID.
// URL: PUT /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyid
func (s *Service) UpdateRestrictedSoftwareByID(ctx context.Context, id int, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// UpdateRestrictedSoftwareByName updates the specified restricted software by name.
// URL: PUT /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updaterestrictedsoftwarebyname
func (s *Service) UpdateRestrictedSoftwareByName(ctx context.Context, name string, req *RequestRestrictedSoftware) (*CreateUpdateResponse, *interfaces.Response, error) {
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

// DeleteRestrictedSoftwareByID removes the specified restricted software by ID.
// URL: DELETE /JSSResource/restrictedsoftware/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyid
func (s *Service) DeleteRestrictedSoftwareByID(ctx context.Context, id int) (*interfaces.Response, error) {
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

// DeleteRestrictedSoftwareByName removes the specified restricted software by name.
// URL: DELETE /JSSResource/restrictedsoftware/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleterestrictedsoftwarebyname
func (s *Service) DeleteRestrictedSoftwareByName(ctx context.Context, name string) (*interfaces.Response, error) {
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
