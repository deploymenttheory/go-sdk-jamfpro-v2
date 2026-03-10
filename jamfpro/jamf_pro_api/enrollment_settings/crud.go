package enrollment_settings

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the enrollment settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
	EnrollmentSettings struct {
		client client.Client
	}
)

func NewEnrollmentSettings(client client.Client) *EnrollmentSettings {
	return &EnrollmentSettings{client: client}
}

// GetV4 retrieves the current enrollment settings.
// URL: GET /api/v4/enrollment
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
func (s *EnrollmentSettings) GetV4(ctx context.Context) (*ResourceEnrollmentSettingsV4, *resty.Response, error) {
	var result ResourceEnrollmentSettingsV4
	endpoint := constants.EndpointJamfProEnrollmentSettingsV4

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
