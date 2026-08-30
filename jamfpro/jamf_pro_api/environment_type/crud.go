package environment_type

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the environment type-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-environment-type
	EnvironmentType struct {
		client client.Client
	}
)

func NewEnvironmentType(client client.Client) *EnvironmentType {
	return &EnvironmentType{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Environment Type Operations
// -----------------------------------------------------------------------------

// GetV2 returns the cloud services environment type that this Jamf Pro instance
// is configured to communicate with.
// URL: GET /api/v2/environment-type
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-environment-type
func (s *EnvironmentType) GetV2(ctx context.Context) (*ResourceEnvironmentType, *resty.Response, error) {
	var result ResourceEnvironmentType

	endpoint := constants.EndpointJamfProEnvironmentTypeV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
