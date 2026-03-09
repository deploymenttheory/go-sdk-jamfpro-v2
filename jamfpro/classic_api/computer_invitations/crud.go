package computer_invitations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ComputerInvitationsServiceInterface defines the interface for Classic API computer invitation operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
	ComputerInvitationsServiceInterface interface {
		// List returns all computer invitations.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified computer invitation by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyid
		GetByID(ctx context.Context, id string) (*ResourceComputerInvitation, *resty.Response, error)

		// GetByInvitationID returns the specified computer invitation by invitation ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyinvitation
		GetByInvitationID(ctx context.Context, invitationID string) (*ResourceComputerInvitation, *resty.Response, error)

		// Create creates a new computer invitation.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerinvitationbyid
		Create(ctx context.Context, req *ResourceComputerInvitation) (*ResourceComputerInvitation, *resty.Response, error)

		// DeleteByID removes the specified computer invitation by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerinvitationbyid
		DeleteByID(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the computer-invitations-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
	ComputerInvitations struct {
		client transport.HTTPClient
	}
)

var _ ComputerInvitationsServiceInterface = (*ComputerInvitations)(nil)

// NewService returns a new computer invitations Service backed by the provided HTTP client.
func NewComputerInvitations(client transport.HTTPClient) *ComputerInvitations {
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
	endpoint := EndpointComputerInvitations

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

// GetByID returns the specified computer invitation by ID.
// URL: GET /JSSResource/computerinvitations/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyid
func (s *ComputerInvitations) GetByID(ctx context.Context, id string) (*ResourceComputerInvitation, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("computer invitation ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", EndpointComputerInvitations, id)

	var out ResourceComputerInvitation

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

// GetByInvitationID returns the specified computer invitation by invitation ID.
// URL: GET /JSSResource/computerinvitations/invitation/{invitationID}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyinvitation
func (s *ComputerInvitations) GetByInvitationID(ctx context.Context, invitationID string) (*ResourceComputerInvitation, *resty.Response, error) {
	if invitationID == "" {
		return nil, nil, fmt.Errorf("computer invitation invitation ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/invitation/%s", EndpointComputerInvitations, invitationID)

	var out ResourceComputerInvitation

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

// Create creates a new computer invitation.
// URL: POST /JSSResource/computerinvitations/id/0
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerinvitationbyid
func (s *ComputerInvitations) Create(ctx context.Context, req *ResourceComputerInvitation) (*ResourceComputerInvitation, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointComputerInvitations)

	var out ResourceComputerInvitation

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

// DeleteByID removes the specified computer invitation by ID.
// URL: DELETE /JSSResource/computerinvitations/id/{id}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerinvitationbyid
func (s *ComputerInvitations) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("computer invitation ID cannot be empty")
	}

	endpoint := fmt.Sprintf("%s/id/%s", EndpointComputerInvitations, id)

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
