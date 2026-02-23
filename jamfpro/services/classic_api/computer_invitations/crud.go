package computer_invitations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInvitationsServiceInterface defines the interface for Classic API computer invitation operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
	ComputerInvitationsServiceInterface interface {
		// List returns all computer invitations.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified computer invitation by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyid
		GetByID(ctx context.Context, id string) (*ResourceComputerInvitation, *interfaces.Response, error)

		// GetByInvitationID returns the specified computer invitation by invitation ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyinvitation
		GetByInvitationID(ctx context.Context, invitationID string) (*ResourceComputerInvitation, *interfaces.Response, error)

		// Create creates a new computer invitation.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerinvitationbyid
		Create(ctx context.Context, req *ResourceComputerInvitation) (*ResourceComputerInvitation, *interfaces.Response, error)

		// DeleteByID removes the specified computer invitation by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerinvitationbyid
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the computer-invitations-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerInvitationsServiceInterface = (*Service)(nil)

// NewService returns a new computer invitations Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Computer Invitations CRUD Operations
// -----------------------------------------------------------------------------

// List returns all computer invitations.
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/computerinvitations
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
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
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyid
func (s *Service) GetByID(ctx context.Context, id string) (*ResourceComputerInvitation, *interfaces.Response, error) {
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
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findcomputerinvitationbyinvitation
func (s *Service) GetByInvitationID(ctx context.Context, invitationID string) (*ResourceComputerInvitation, *interfaces.Response, error) {
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
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputerinvitationbyid
func (s *Service) Create(ctx context.Context, req *ResourceComputerInvitation) (*ResourceComputerInvitation, *interfaces.Response, error) {
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
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletecomputerinvitationbyid
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
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
