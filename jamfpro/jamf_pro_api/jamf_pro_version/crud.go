package jamf_pro_version

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro version-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
	JamfProVersion struct {
		client client.Client
	}
)

func NewJamfProVersion(client client.Client) *JamfProVersion {
	return &JamfProVersion{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Pro Version Operations
// -----------------------------------------------------------------------------

// GetV1 returns the Jamf Pro server version.
// URL: GET /api/v1/jamf-pro-version
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
func (s *JamfProVersion) GetV1(ctx context.Context) (*ResourceJamfProVersion, *resty.Response, error) {
	var result ResourceJamfProVersion

	endpoint := constants.EndpointJamfProJamfProVersionV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
