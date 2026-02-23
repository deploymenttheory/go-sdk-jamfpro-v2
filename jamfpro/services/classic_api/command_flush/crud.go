package command_flush

import (
	"context"
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

// CommandFlushServiceInterface defines the operations available in the command flush service.
// Doc: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
// Doc: https://developer.jamf.com/jamf-pro/reference/commandflush-1
type CommandFlushServiceInterface interface {
	// FlushByIDAndStatus clears MDM commands for a specific device or group by ID and status.
	// idType: computers, computergroups, mobiledevices, or mobiledevicegroups
	// status: Pending, Failed, or Pending+Failed
	FlushByIDAndStatus(ctx context.Context, idType string, id string, status string) (*interfaces.Response, error)

	// FlushWithXML clears MDM commands using an XML request body for batch operations.
	FlushWithXML(ctx context.Context, req *RequestCommandFlush) (*interfaces.Response, error)
}

// Service provides access to command flush operations.
// Doc: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
// Doc: https://developer.jamf.com/jamf-pro/reference/commandflush-1
type Service struct {
	client interfaces.HTTPClient
}

// NewService creates a new command flush service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// FlushByIDAndStatus clears MDM commands for a specific device or group by ID and status.
// idType: computers, computergroups, mobiledevices, or mobiledevicegroups
// status: Pending, Failed, or Pending+Failed
// Doc: https://developer.jamf.com/jamf-pro/reference/createcommandflushwithidandstatus
func (s *Service) FlushByIDAndStatus(ctx context.Context, idType string, id string, status string) (*interfaces.Response, error) {
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

	resp, err := s.client.Delete(ctx, endpoint, headers, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to clear %s MDM commands for %s %s: %w", status, idType, id, err)
	}

	return resp, nil
}

// FlushWithXML clears MDM commands using an XML request body for batch operations.
// Doc: https://developer.jamf.com/jamf-pro/reference/commandflush-1
func (s *Service) FlushWithXML(ctx context.Context, req *RequestCommandFlush) (*interfaces.Response, error) {
	if err := validateCommandFlushRequest(req); err != nil {
		return nil, err
	}

	endpoint := EndpointCommandFlush

	xmlData, err := xml.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal command flush request: %w", err)
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	bodyMap := map[string]string{
		"body": string(xmlData),
	}

	resp, err := s.client.Delete(ctx, endpoint, headers, bodyMap, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to flush commands with XML request: %w", err)
	}

	return resp, nil
}
