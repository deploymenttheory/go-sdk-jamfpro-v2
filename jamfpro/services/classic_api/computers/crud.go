package computers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputersServiceInterface defines the interface for Classic API computer operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computers
	ComputersServiceInterface interface {
		// List returns all computers.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputers
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified computer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyid
		GetByID(ctx context.Context, id string) (*ResponseComputer, *interfaces.Response, error)

		// GetByName returns the specified computer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyname
		GetByName(ctx context.Context, name string) (*ResponseComputer, *interfaces.Response, error)

		// Create creates a new computer.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerbyid
		Create(ctx context.Context, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error)

		// UpdateByID updates the specified computer by ID.
		// If Site.ID == 0 && Site.Name == "", sets ID = -1 and Name = "none".
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyid
		UpdateByID(ctx context.Context, id string, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error)

		// UpdateByName updates the specified computer by name.
		// If Site.ID == 0 && Site.Name == "", sets ID = -1 and Name = "none".
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyname
		UpdateByName(ctx context.Context, name string, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error)

		// DeleteByID removes the specified computer by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyid
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByName removes the specified computer by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the computers-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computers
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputersServiceInterface = (*Service)(nil)

// NewService returns a new computers Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// applySiteDefault sets Site to ID=-1, Name="none" when both are empty/zero.
func applySiteDefault(computer *ResponseComputer) {
	if computer.General.Site.ID == 0 && computer.General.Site.Name == "" {
		computer.General.Site.ID = -1
		computer.General.Site.Name = "none"
	}
}

// -----------------------------------------------------------------------------
// Classic API - Computers CRUD Operations
// -----------------------------------------------------------------------------

// List returns all computers.
// URL: GET /JSSResource/computers
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputers
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointComputers

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

// GetByID returns the specified computer by ID.
// URL: GET /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyid
func (s *Service) GetByID(ctx context.Context, id string) (*ResponseComputer, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", EndpointComputers, id)

	var out ResponseComputer

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

// GetByName returns the specified computer by name.
// URL: GET /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyname
func (s *Service) GetByName(ctx context.Context, name string) (*ResponseComputer, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointComputers, name)

	var out ResponseComputer

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

// Create creates a new computer.
// URL: POST /JSSResource/computers
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerbyid
func (s *Service) Create(ctx context.Context, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error) {
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}
	if computer.General.Name == "" {
		return nil, nil, fmt.Errorf("computer name is required")
	}

	endpoint := EndpointComputers

	var out ResponseComputer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, computer, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified computer by ID.
// URL: PUT /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyid
func (s *Service) UpdateByID(ctx context.Context, id string, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer ID cannot be empty")
	}
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}

	applySiteDefault(computer)

	endpoint := fmt.Sprintf("%s/id/%s", EndpointComputers, id)

	var out ResponseComputer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, computer, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByName updates the specified computer by name.
// URL: PUT /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyname
func (s *Service) UpdateByName(ctx context.Context, name string, computer *ResponseComputer) (*ResponseComputer, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}

	applySiteDefault(computer)

	endpoint := fmt.Sprintf("%s/name/%s", EndpointComputers, name)

	var out ResponseComputer

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, computer, headers, &out)
	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified computer by ID.
// URL: DELETE /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyid
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", EndpointComputers, id)

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

// DeleteByName removes the specified computer by name.
// URL: DELETE /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyname
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("computer name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointComputers, name)

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
