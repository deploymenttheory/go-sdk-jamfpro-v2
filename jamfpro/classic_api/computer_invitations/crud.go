package computer_invitations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the computer-invitations-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
	ComputerInvitations struct {
		client client.Client
	}
)

// NewService returns a new computer invitations Service backed by the provided HTTP client.
func NewComputerInvitations(client client.Client) *ComputerInvitations {
	return &ComputerInvitations{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer Invitations CRUD Operations
// -----------------------------------------------------------------------------

// List returns all computer invitations.
// URL: GET /JSSResource/computerinvitations
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
func (s *ComputerInvitations) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var out ListResponse

	endpoint := constants.EndpointClassicComputerInvitations

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

// GetByID returns the specified computer invitation by ID.
// URL: GET /JSSResource/computerinvitations/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyid
func (s *ComputerInvitations) GetByID(ctx context.Context, id string) (*ResourceComputerInvitation, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer invitation ID cannot be empty")
	}

	var out ResourceComputerInvitation

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicComputerInvitations, id)

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

// GetByInvitationID returns the specified computer invitation by invitation ID.
// URL: GET /JSSResource/computerinvitations/invitation/{invitationID}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyinvitation
func (s *ComputerInvitations) GetByInvitationID(ctx context.Context, invitationID string) (*ResourceComputerInvitation, *resty.Response, error) {
	if invitationID == "" {
		return nil, nil, fmt.Errorf("computer invitation invitation ID cannot be empty")
	}

	var out ResourceComputerInvitation

	endpoint := fmt.Sprintf("%s/invitation/%s", constants.EndpointClassicComputerInvitations, invitationID)

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

// Create creates a new computer invitation.
// URL: POST /JSSResource/computerinvitations/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerinvitationbyid
func (s *ComputerInvitations) Create(ctx context.Context, req *ResourceComputerInvitation) (*ResourceComputerInvitation, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var out ResourceComputerInvitation

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicComputerInvitations)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		SetResult(&out).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &out, resp, nil
}

// DeleteByID removes the specified computer invitation by ID.
// URL: DELETE /JSSResource/computerinvitations/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerinvitationbyid
func (s *ComputerInvitations) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer invitation ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", constants.EndpointClassicComputerInvitations, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
