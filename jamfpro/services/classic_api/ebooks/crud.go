package ebooks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for Classic API ebook operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ebooks
	ServiceInterface interface {
		// List returns all ebooks.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooks
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified ebook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbyid
		GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error)

		// GetByName returns the specified ebook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbyname
		GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error)

		// GetByNameAndSubset returns a specific subset of an ebook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbynamedatasubset
		GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error)

		// Create creates a new ebook.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createebookbyid
		Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified ebook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateebookbyid
		UpdateByID(ctx context.Context, id int, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified ebook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateebookbyname
		UpdateByName(ctx context.Context, name string, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified ebook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteebookbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified ebook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteebookbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the ebooks-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/ebooks
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService returns a new ebooks Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Ebooks CRUD Operations
// -----------------------------------------------------------------------------

// List returns all ebooks.
// URL: GET /JSSResource/ebooks
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooks
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointEbooks

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

// GetByID returns the specified ebook by ID.
// URL: GET /JSSResource/ebooks/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbyid
func (s *Service) GetByID(ctx context.Context, id int) (*Resource, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("ebook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointEbooks, id)

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

// GetByName returns the specified ebook by name.
// URL: GET /JSSResource/ebooks/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbyname
func (s *Service) GetByName(ctx context.Context, name string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("ebook name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointEbooks, name)

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

// GetByNameAndSubset returns a specific subset of an ebook by name.
// URL: GET /JSSResource/ebooks/name/{name}/subset/{subset}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findebooksbynamedatasubset
func (s *Service) GetByNameAndSubset(ctx context.Context, name, subset string) (*Resource, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("ebook name cannot be empty")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("ebook subset cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", EndpointEbooks, name, subset)

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

// Create creates a new ebook.
// URL: POST /JSSResource/ebooks/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createebookbyid
func (s *Service) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("ebook name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointEbooks)

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

// UpdateByID updates the specified ebook by ID.
// URL: PUT /JSSResource/ebooks/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateebookbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("ebook ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("ebook name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointEbooks, id)

	var out CreateUpdateResponse

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

// UpdateByName updates the specified ebook by name.
// URL: PUT /JSSResource/ebooks/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateebookbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *Resource) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("ebook name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("ebook name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointEbooks, name)

	var out CreateUpdateResponse

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

// DeleteByID removes the specified ebook by ID.
// URL: DELETE /JSSResource/ebooks/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteebookbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("ebook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointEbooks, id)

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

// DeleteByName removes the specified ebook by name.
// URL: DELETE /JSSResource/ebooks/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteebookbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("ebook name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointEbooks, name)

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
