package printers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// PrintersServiceInterface defines the interface for Classic API printer operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	PrintersServiceInterface interface {
		// List returns all printers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprinters
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyid
		GetByID(ctx context.Context, id int) (*ResourcePrinter, *interfaces.Response, error)

		// GetByName returns the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyname
		GetByName(ctx context.Context, name string) (*ResourcePrinter, *interfaces.Response, error)

		// Create creates a new printer.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createprinterbyid
		Create(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyid
		UpdateByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyname
		UpdateByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the printer-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PrintersServiceInterface = (*Service)(nil)

// NewService returns a new printers Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Printers CRUD Operations
// -----------------------------------------------------------------------------

// List returns all printers.
// URL: GET /JSSResource/printers
// https://developer.jamf.com/jamf-pro/reference/findprinters
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicPrinters

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

// GetByID returns the specified printer by ID.
// URL: GET /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyid
func (s *Service) GetByID(ctx context.Context, id int) (*ResourcePrinter, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("printer ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

	var result ResourcePrinter

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

// GetByName returns the specified printer by name.
// URL: GET /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResourcePrinter, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("printer name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

	var result ResourcePrinter

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

// Create creates a new printer.
// URL: POST /JSSResource/printers/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/createprinterbyid
func (s *Service) Create(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicPrinters)

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

// UpdateByID updates the specified printer by ID.
// URL: PUT /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateprinterbyid
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("printer ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

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

// UpdateByName updates the specified printer by name.
// URL: PUT /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateprinterbyname
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("printer name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

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

// DeleteByID removes the specified printer by ID.
// URL: DELETE /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("printer ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

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

// DeleteByName removes the specified printer by name.
// URL: DELETE /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteprinterbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("printer name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

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
