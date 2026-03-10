package jamf_pro_information

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro information-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-pro-information
	JamfProInformation struct {
		client client.Client
	}
)

func NewJamfProInformation(client client.Client) *JamfProInformation {
	return &JamfProInformation{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Pro Information Operations
// -----------------------------------------------------------------------------

// GetV2 returns Jamf Pro information (feature flags).
// URL: GET /api/v2/jamf-pro-information
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-pro-information
func (s *JamfProInformation) GetV2(ctx context.Context) (*ResourceJamfProInformation, *resty.Response, error) {
	var result ResourceJamfProInformation

	endpoint := constants.EndpointJamfProJamfProInformationV2

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
