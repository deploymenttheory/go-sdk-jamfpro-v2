package command_flush

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// CommandFlushServiceInterface defines the interface for Classic API command flush operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/commandflush-1
	CommandFlushServiceInterface interface {
		// FlushByIDAndStatus clears MDM commands for a specific device or group by ID and status.
		//
		// Valid idType values: computers, computergroups, mobiledevices, or mobiledevicegroups
		// Valid status values: Pending, Failed, or Pending+Failed
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
		FlushByIDAndStatus(ctx context.Context, idType string, id string, status string) (*resty.Response, error)

		// FlushWithXML clears MDM commands using an XML request body for batch operations.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/commandflush-1
		FlushWithXML(ctx context.Context, req *RequestCommandFlush) (*resty.Response, error)
	}

	// Service handles communication with the command-flush-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/commandflush-1
	CommandFlush struct {
		client transport.HTTPClient
	}
)

var _ CommandFlushServiceInterface = (*CommandFlush)(nil)

// NewService returns a new command flush Service backed by the provided HTTP client.
func NewCommandFlush(client transport.HTTPClient) *CommandFlush {
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

	// URL encode status if it contains +
	encodedStatus := strings.ReplaceAll(status, "+", "%2B")

	endpoint := fmt.Sprintf("%s/%s/id/%s/status/%s", EndpointCommandFlush, idType, id, encodedStatus)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
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

	endpoint := EndpointCommandFlush

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.DeleteWithBody(ctx, endpoint, req, headers, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to flush commands with XML request: %w", err)
	}

	return resp, nil
}
