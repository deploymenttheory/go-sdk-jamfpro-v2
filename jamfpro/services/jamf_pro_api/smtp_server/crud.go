package smtp_server

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// SMTPServerServiceInterface defines the interface for SMTP server operations (singleton).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	SMTPServerServiceInterface interface {
		// GetSMTPServerV2 returns the current SMTP server configuration (Get SMTP server).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
		GetSMTPServerV2(ctx context.Context) (*ResourceSMTPServer, *interfaces.Response, error)

		// UpdateSMTPServerV2 updates the SMTP server configuration (Update SMTP server).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-smtp-server
		UpdateSMTPServerV2(ctx context.Context, settings *ResourceSMTPServer) (*ResourceSMTPServer, *interfaces.Response, error)
	}

	// Service handles communication with the SMTP server-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SMTPServerServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetSMTPServerV2 returns the current SMTP server configuration.
// URL: GET /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
func (s *Service) GetSMTPServerV2(ctx context.Context) (*ResourceSMTPServer, *interfaces.Response, error) {
	var result ResourceSMTPServer
	resp, err := s.client.Get(ctx, EndpointSMTPServerV2, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateSMTPServerV2 updates the SMTP server configuration.
// URL: PUT /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-smtp-server
func (s *Service) UpdateSMTPServerV2(ctx context.Context, settings *ResourceSMTPServer) (*ResourceSMTPServer, *interfaces.Response, error) {
	if settings == nil {
		return nil, nil, fmt.Errorf("settings is required")
	}
	var result ResourceSMTPServer
	resp, err := s.client.Put(ctx, EndpointSMTPServerV2, settings, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
