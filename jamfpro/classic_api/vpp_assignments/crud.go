package vpp_assignments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the VPP assignments-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/vppassignments
	VppAssignments struct {
		client client.Client
	}
)

// NewService returns a new VPP assignments Service backed by the provided HTTP client.
func NewVppAssignments(client client.Client) *VppAssignments {
	return &VppAssignments{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - VPP Assignments CRUD Operations
// -----------------------------------------------------------------------------

// List returns all VPP assignments.
//
// URL: GET /JSSResource/vppassignments
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findvppassignments
func (s *VppAssignments) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicVPPAssignments

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

// GetByID returns the specified VPP assignment by ID.
//
// URL: GET /JSSResource/vppassignments/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findvppassignmentbyid
func (s *VppAssignments) GetByID(ctx context.Context, id int) (*Resource, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP assignment ID must be a positive integer")
	}

	var out Resource

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAssignments, id)

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

// Create creates a new VPP assignment.
//
// URL: POST /JSSResource/vppassignments/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createvppassignmentbyid
func (s *VppAssignments) Create(ctx context.Context, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("VPP assignment name is required")
	}

	requestBody := &RequestVPPAssignment{Resource: req}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicVPPAssignments)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(requestBody).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// UpdateByID updates the specified VPP assignment by ID.
//
// URL: PUT /JSSResource/vppassignments/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatevppassignmentbyid
func (s *VppAssignments) UpdateByID(ctx context.Context, id int, req *Resource) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("VPP assignment ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.General.Name == "" {
		return nil, nil, fmt.Errorf("VPP assignment name is required")
	}

	requestBody := &RequestVPPAssignment{Resource: req}

	var out CreateUpdateResponse

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAssignments, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(requestBody).
		SetResult(&out).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}
	return &out, resp, nil
}

// DeleteByID removes the specified VPP assignment by ID.
//
// URL: DELETE /JSSResource/vppassignments/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletevppassignmentbyid
func (s *VppAssignments) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("VPP assignment ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicVPPAssignments, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}
	return resp, nil
}
