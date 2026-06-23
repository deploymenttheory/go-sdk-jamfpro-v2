package m2m

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// M2M handles communication with the machine-to-machine (M2M) methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-m2m-tenant-id
	M2M struct {
		client client.Client
	}
)

func NewM2M(client client.Client) *M2M {
	return &M2M{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - M2M Operations
// -----------------------------------------------------------------------------

// GetTenantIdV1 returns the tenant ID associated with the Jamf Pro instance.
// URL: GET /api/v1/m2m/tenant-id
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-m2m-tenant-id
func (s *M2M) GetTenantIdV1(ctx context.Context) (*ResourceM2mTenantId, *resty.Response, error) {
	var result ResourceM2mTenantId

	endpoint := constants.EndpointJamfProM2MTenantIdV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
