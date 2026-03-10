package computers

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the computers-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computers
	Computers struct {
		client client.Client
	}
)

// NewService returns a new computers Service backed by the provided HTTP client.
func NewComputers(client client.Client) *Computers {
	return &Computers{client: client}
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
func (s *Computers) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicComputers

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByID returns the specified computer by ID.
// URL: GET /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyid
func (s *Computers) GetByID(ctx context.Context, id string) (*ResponseComputer, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer ID cannot be empty")
	}

	var out ResponseComputer

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicComputers, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// GetByName returns the specified computer by name.
// URL: GET /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputersbyname
func (s *Computers) GetByName(ctx context.Context, name string) (*ResponseComputer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}

	var out ResponseComputer

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputers, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&out).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// Create creates a new computer.
// URL: POST /JSSResource/computers
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerbyid
func (s *Computers) Create(ctx context.Context, computer *ResponseComputer) (*ResponseComputer, *resty.Response, error) {
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}
	if computer.General.Name == "" {
		return nil, nil, fmt.Errorf("computer name is required")
	}

	var out ResponseComputer

	endpoint := constants.EndpointClassicComputers

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(computer).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByID updates the specified computer by ID.
// URL: PUT /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyid
func (s *Computers) UpdateByID(ctx context.Context, id string, computer *ResponseComputer) (*ResponseComputer, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer ID cannot be empty")
	}
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}

	applySiteDefault(computer)

	var out ResponseComputer

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicComputers, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(computer).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// UpdateByName updates the specified computer by name.
// URL: PUT /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputerbyname
func (s *Computers) UpdateByName(ctx context.Context, name string, computer *ResponseComputer) (*ResponseComputer, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer name cannot be empty")
	}
	if computer == nil {
		return nil, nil, fmt.Errorf("computer is required")
	}

	applySiteDefault(computer)

	var out ResponseComputer

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputers, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(computer).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// DeleteByID removes the specified computer by ID.
// URL: DELETE /JSSResource/computers/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyid
func (s *Computers) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicComputers, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified computer by name.
// URL: DELETE /JSSResource/computers/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerbyname
func (s *Computers) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("computer name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputers, name)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
