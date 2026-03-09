package printers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// PrintersServiceInterface defines the interface for Classic API printer operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	PrintersServiceInterface interface {
		// List returns all printers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprinters
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyid
		GetByID(ctx context.Context, id int) (*ResourcePrinter, *resty.Response, error)

		// GetByName returns the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyname
		GetByName(ctx context.Context, name string) (*ResourcePrinter, *resty.Response, error)

		// Create creates a new printer.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createprinterbyid
		Create(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyid
		UpdateByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyname
		UpdateByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the printer-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	Printers struct {
		client transport.HTTPClient
	}
)

var _ PrintersServiceInterface = (*Printers)(nil)

// NewService returns a new printers Service backed by the provided HTTP client.
func NewPrinters(client transport.HTTPClient) *Printers {
	return &Printers{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Printers CRUD Operations
// -----------------------------------------------------------------------------

// List returns all printers.
// URL: GET /JSSResource/printers
// https://developer.jamf.com/jamf-pro/reference/findprinters
func (s *Printers) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
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
func (s *Printers) GetByID(ctx context.Context, id int) (*ResourcePrinter, *resty.Response, error) {
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
func (s *Printers) GetByName(ctx context.Context, name string) (*ResourcePrinter, *resty.Response, error) {
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
func (s *Printers) Create(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error) {
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
func (s *Printers) UpdateByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error) {
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
func (s *Printers) UpdateByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error) {
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
func (s *Printers) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
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
func (s *Printers) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
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
