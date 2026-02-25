package mac_applications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API Mac application operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/macapplications
	ServiceInterface interface {
		// List returns all Mac applications.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplications
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified Mac application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyid
		GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error)

		// GetByName returns the specified Mac application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyname
		GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error)

// GetByIDAndSubset returns a specific subset of a Mac application by ID.
// Subset values: General, Scope, SelfService, VPPCodes, VPP.
//
// URL: GET /JSSResource/macapplications/id/{id}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyiddatasubset
		GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error)

// GetByNameAndSubset returns a specific subset of a Mac application by name.
// Subset values: General, Scope, SelfService, VPPCodes, VPP.
//
// URL: GET /JSSResource/macapplications/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbynamedatasubset
		GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error)

		// Create creates a new Mac application.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmacapplicationbyid
		Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified Mac application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyid
		UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error)

		// UpdateByName updates the specified Mac application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyname
		UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error)

		// DeleteByID removes the specified Mac application by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified Mac application by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the Mac applications-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/macapplications
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new Mac applications Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Mac Applications CRUD Operations
// -----------------------------------------------------------------------------

// List returns all Mac applications.
//
// URL: GET /JSSResource/macapplications
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplications
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointMacApplications

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

// GetByID returns the specified Mac application by ID.
//
// URL: GET /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyid
func (s *Service) GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMacApplications, id)

	var out Resource

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

// GetByName returns the specified Mac application by name.
//
// URL: GET /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyname
func (s *Service) GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMacApplications, name)

	var out Resource

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

// GetByIDAndSubset returns a specific subset of a Mac application by ID.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbyiddatasubset
func (s *Service) GetByIDAndSubset(ctx context.Context, id int, subset string) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mac application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", EndpointMacApplications, id, subset)

	var out Resource

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

// GetByNameAndSubset returns a specific subset of a Mac application by name.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findmacapplicationsbynamedatasubset
func (s *Service) GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("mac application subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", EndpointMacApplications, name, subset)

	var out Resource

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

// Create creates a new Mac application.
//
// URL: POST /JSSResource/macapplications/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmacapplicationbyid
func (s *Service) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointMacApplications)

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

// UpdateByID updates the specified Mac application by ID.
//
// URL: PUT /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *Resource) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("mac application ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMacApplications, id)

	var out Resource

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

// UpdateByName updates the specified Mac application by name.
//
// URL: PUT /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatemacapplicationbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *Resource) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("mac application name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("mac application name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMacApplications, name)

	var out Resource

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

// DeleteByID removes the specified Mac application by ID.
//
// URL: DELETE /JSSResource/macapplications/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("mac application ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointMacApplications, id)

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

// DeleteByName removes the specified Mac application by name.
//
// URL: DELETE /JSSResource/macapplications/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletemacapplicationbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("mac application name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointMacApplications, name)

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
