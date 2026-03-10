package printers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the printer-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/printers
	Printers struct {
		client client.Client
	}
)

// NewService returns a new printers Service backed by the provided HTTP client.
func NewPrinters(client client.Client) *Printers {
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

	endpoint := constants.EndpointClassicPrinters

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

// GetByID returns the specified printer by ID.
// URL: GET /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyid
func (s *Printers) GetByID(ctx context.Context, id int) (*ResourcePrinter, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("printer ID must be a positive integer")
	}

	var result ResourcePrinter

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPrinters, id)

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

// GetByName returns the specified printer by name.
// URL: GET /JSSResource/printers/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findprintersbyname
func (s *Printers) GetByName(ctx context.Context, name string) (*ResourcePrinter, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("printer name is required")
	}

	var result ResourcePrinter

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicPrinters, name)

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

// Create creates a new printer.
// URL: POST /JSSResource/printers/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/createprinterbyid
func (s *Printers) Create(ctx context.Context, req *RequestPrinter) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicPrinters)

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

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPrinters, id)

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

	var result CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicPrinters, name)

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

// DeleteByID removes the specified printer by ID.
// URL: DELETE /JSSResource/printers/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deleteprinterbyid
func (s *Printers) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("printer ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicPrinters, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

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

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicPrinters, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
