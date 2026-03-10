package dss_declarations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the DSS declarations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-declarationid
	DssDeclarations struct {
		client client.Client
	}
)

// NewService returns a new DSS declarations Service backed by the provided HTTP client.
func NewDssDeclarations(client client.Client) *DssDeclarations {
	return &DssDeclarations{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - DSS Declarations Operations
// -----------------------------------------------------------------------------

// GetByUUIDV1 returns the specified DSS declaration by UUID.
// URL: GET /api/v1/dss-declarations/{uuid}
// https://developer.jamf.com/jamf-pro/reference/get_v1-dss-declarations-declarationid
func (s *DssDeclarations) GetByUUIDV1(ctx context.Context, uuid string) (*ResponseDSSDeclaration, *resty.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("uuid is required")
	}

	var result ResponseDSSDeclaration

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDSSDeclarationsV1, uuid)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get DSS declaration by UUID %s: %w", uuid, err)
	}

	return &result, resp, nil
}
