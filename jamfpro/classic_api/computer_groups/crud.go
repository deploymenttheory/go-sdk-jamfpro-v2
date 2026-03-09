package computer_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ComputerGroupsServiceInterface defines the interface for Classic API computer group operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroups
	ComputerGroupsServiceInterface interface {
		// List returns all computer groups.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroups
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified computer group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroupsbyid
		GetByID(ctx context.Context, id int) (*ResourceComputerGroup, *resty.Response, error)

		// GetByName returns the specified computer group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroupsbyname
		GetByName(ctx context.Context, name string) (*ResourceComputerGroup, *resty.Response, error)

		// Create creates a new computer group.
		//
		// Returns the created computer group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputergroupbyid
		Create(ctx context.Context, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified computer group by ID.
		//
		// Returns the updated computer group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputergroupbyid
		UpdateByID(ctx context.Context, id int, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified computer group by name.
		//
		// Returns the updated computer group ID only (Classic API behavior).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputergroupbyname
		UpdateByName(ctx context.Context, name string, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified computer group by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputergroupbyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified computer group by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputergroupbyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)
	}

	// Service handles communication with the computer-groups-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroups
	ComputerGroups struct {
		client transport.HTTPClient
	}
)

var _ ComputerGroupsServiceInterface = (*ComputerGroups)(nil)

// NewService returns a new computer groups Service backed by the provided HTTP client.
func NewComputerGroups(client transport.HTTPClient) *ComputerGroups {
	return &ComputerGroups{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer Groups CRUD Operations
// -----------------------------------------------------------------------------

// List returns all computer groups.
// URL: GET /JSSResource/computergroups
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroups
func (s *ComputerGroups) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointClassicComputerGroups

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

// GetByID returns the specified computer group by ID.
// URL: GET /JSSResource/computergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroupsbyid
func (s *ComputerGroups) GetByID(ctx context.Context, id int) (*ResourceComputerGroup, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("computer group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicComputerGroups, id)

	var out ResourceComputerGroup

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

// GetByName returns the specified computer group by name.
// URL: GET /JSSResource/computergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputergroupsbyname
func (s *ComputerGroups) GetByName(ctx context.Context, name string) (*ResourceComputerGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputerGroups, name)

	var out ResourceComputerGroup

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

// Create creates a new computer group.
//
// Returns the created computer group ID only (Classic API behavior).
// URL: POST /JSSResource/computergroups/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputergroupbyid
func (s *ComputerGroups) Create(ctx context.Context, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("computer group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicComputerGroups)

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

// UpdateByID updates the specified computer group by ID.
//
// Returns the updated computer group ID only (Classic API behavior).
// URL: PUT /JSSResource/computergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputergroupbyid
func (s *ComputerGroups) UpdateByID(ctx context.Context, id int, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("computer group ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("computer group name is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicComputerGroups, id)

	var out CreateUpdateResponse

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

// UpdateByName updates the specified computer group by name.
//
// Returns the updated computer group ID only (Classic API behavior).
// URL: PUT /JSSResource/computergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatecomputergroupbyname
func (s *ComputerGroups) UpdateByName(ctx context.Context, name string, req *RequestComputerGroup) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("computer group name cannot be empty")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("computer group name is required in request")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputerGroups, name)

	var out CreateUpdateResponse

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

// DeleteByID removes the specified computer group by ID.
// URL: DELETE /JSSResource/computergroups/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputergroupbyid
func (s *ComputerGroups) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("computer group ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicComputerGroups, id)

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

// DeleteByName removes the specified computer group by name.
// URL: DELETE /JSSResource/computergroups/name/{name}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputergroupbyname
func (s *ComputerGroups) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("computer group name cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicComputerGroups, name)

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
