package dss_declarations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DSSDeclarationsServiceInterface defines the interface for DSS declarations operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-id
	DSSDeclarationsServiceInterface interface {
		// GetByUUIDV1 returns the specified DSS declaration by UUID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-id
		GetByUUIDV1(ctx context.Context, uuid string) (*ResponseDSSDeclaration, *interfaces.Response, error)
	}

	// Service handles communication with the DSS declarations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-id
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DSSDeclarationsServiceInterface = (*Service)(nil)

// NewService returns a new DSS declarations Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - DSS Declarations Operations
// -----------------------------------------------------------------------------

// GetByUUIDV1 returns the specified DSS declaration by UUID.
// URL: GET /api/v1/dss-declarations/{uuid}
// https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-id
func (s *Service) GetByUUIDV1(ctx context.Context, uuid string) (*ResponseDSSDeclaration, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	var result ResponseDSSDeclaration

	endpoint := fmt.Sprintf("%s/%s", EndpointDSSDeclarationsV1, uuid)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get DSS declaration by UUID %s: %w", uuid, err)
	}

	return &result, resp, nil
}
