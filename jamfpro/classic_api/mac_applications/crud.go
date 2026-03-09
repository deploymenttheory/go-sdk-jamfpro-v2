package mac_applications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Mac applications-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/macapplications
	MacApplications struct {
		client transport.HTTPClient
	}
)

// NewService returns a new Mac applications Service backed by the provided HTTP client.
func NewMacApplications(client transport.HTTPClient) *MacApplications {
	return &MacApplications{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mac Applications CRUD Operations
// -----------------------------------------------------------------------------

// List returns all Mac applications.
//
// URL: GET /JSSResource/macapplications
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplications
func (s *MacApplications) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicMacApplications

	var out ListResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByID returns the specified Mac application by ID.
//
// URL: GET /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyid
func (s *MacApplications) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacApplications, id)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByName returns the specified Mac application by name.
//
// URL: GET /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyname
func (s *MacApplications) GetByName(ctx context.Context, name string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacApplications, name)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByIDAndSubset returns a specific subset of a Mac application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyiddatasubset
func (s *MacApplications) GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mac application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", constants.EndpointClassicMacApplications, id, subset)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// GetByNameAndSubset returns a specific subset of a Mac application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbynamedatasubset
func (s *MacApplications) GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mac application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", constants.EndpointClassicMacApplications, name, subset)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// Create creates a new Mac application.
//
// URL: POST /JSSResource/macapplications/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmacapplicationbyid
func (s *MacApplications) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicMacApplications)

	var out CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified Mac application by ID.
//
// URL: PUT /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyid
func (s *MacApplications) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacApplications, id)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified Mac application by name.
//
// URL: PUT /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyname
func (s *MacApplications) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacApplications, name)

	var out Resource

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified Mac application by ID.
//
// URL: DELETE /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyid
func (s *MacApplications) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mac application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicMacApplications, id)

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

// DeleteByName removes the specified Mac application by name.
//
// URL: DELETE /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyname
func (s *MacApplications) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mac application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicMacApplications, name)

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
