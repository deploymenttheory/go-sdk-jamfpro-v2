package printers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// PrintersServiceInterface defines the interface for Classic API printer operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	PrintersServiceInterface interface {
		// ListPrinters returns all printers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallprinters
		ListPrinters(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetPrinterByID returns the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyid
		GetPrinterByID(ctx context.Context, id int) (*ResourcePrinter, *interfaces.Response, error)

		// GetPrinterByName returns the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findprintersbyname
		GetPrinterByName(ctx context.Context, name string) (*ResourcePrinter, *interfaces.Response, error)

		// CreatePrinter creates a new printer.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createprinterbyid
		CreatePrinter(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdatePrinterByID updates the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyid
		UpdatePrinterByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdatePrinterByName updates the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateprinterbyname
		UpdatePrinterByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeletePrinterByID removes the specified printer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
		DeletePrinterByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeletePrinterByName removes the specified printer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deleteprinterbyname
		DeletePrinterByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// ListPrinters returns all printers.
// URL: GET /JSSResource/printers
// https://developer.jamf.com/jamf-pro/reference/findallprinters
func (s *Service) ListPrinters(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicPrinters, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPrinterByID returns the specified printer by ID.
// URL: GET /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyid
func (s *Service) GetPrinterByID(ctx context.Context, id int) (*ResourcePrinter, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("printer ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

	var result ResourcePrinter

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPrinterByName returns the specified printer by name.
// URL: GET /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyname
func (s *Service) GetPrinterByName(ctx context.Context, name string) (*ResourcePrinter, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("printer name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

	var result ResourcePrinter

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreatePrinter creates a new printer.
// URL: POST /JSSResource/printers/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/createprinterbyid
func (s *Service) CreatePrinter(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicPrinters)

	var result CreateUpdateResponse

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdatePrinterByID updates the specified printer by ID.
// URL: PUT /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updateprinterbyid
func (s *Service) UpdatePrinterByID(ctx context.Context, id int, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("printer ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdatePrinterByName updates the specified printer by name.
// URL: PUT /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updateprinterbyname
func (s *Service) UpdatePrinterByName(ctx context.Context, name string, req *RequestPrinter) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("printer name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeletePrinterByID removes the specified printer by ID.
// URL: DELETE /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
func (s *Service) DeletePrinterByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("printer ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPrinters, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeletePrinterByName removes the specified printer by name.
// URL: DELETE /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deleteprinterbyname
func (s *Service) DeletePrinterByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("printer name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPrinters, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
