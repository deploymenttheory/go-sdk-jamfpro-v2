package command_flush

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the command-flush-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/commandflush-1
	CommandFlush struct {
		client client.Client
	}
)

// NewService returns a new command flush Service backed by the provided HTTP client.
func NewCommandFlush(client client.Client) *CommandFlush {
	return &CommandFlush{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Command Flush Operations
// -----------------------------------------------------------------------------

// FlushByIDAndStatus clears MDM commands for a specific device or group by ID and status.
//
// Valid idType values: computers, computergroups, mobiledevices, or mobiledevicegroups
// Valid status values: Pending, Failed, or Pending+Failed
// URL: DELETE /JSSResource/commandflush/{idType}/id/{id}/status/{status}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
func (s *CommandFlush) FlushByIDAndStatus(ctx context.Context, idType string, id string, status string) (*resty.Response, error) {
	if err := validateIDType(idType); err != nil {
		return nil, err
	}

	if err := validateStatus(status); err != nil {
		return nil, err
	}

	encodedStatus := strings.ReplaceAll(status, "+", "%2B")

	endpoint := fmt.Sprintf("%s/%s/id/%s/status/%s", constants.EndpointClassicCommandFlush, idType, id, encodedStatus)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Delete(endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to clear %s MDM commands for %s %s: %w", status, idType, id, err)
	}

	return resp, nil
}

// FlushWithXML clears MDM commands using an XML request body for batch operations.
// URL: DELETE /JSSResource/commandflush
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/commandflush-1
func (s *CommandFlush) FlushWithXML(ctx context.Context, req *RequestCommandFlush) (*resty.Response, error) {
	if err := validateCommandFlushRequest(req); err != nil {
		return nil, err
	}

	endpoint := constants.EndpointClassicCommandFlush

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(req).
		Delete(endpoint)

	if err != nil {
		return nil, fmt.Errorf("failed to flush commands with XML request: %w", err)
	}

	return resp, nil
}
