package enrollment_settings

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the enrollment settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
	EnrollmentSettings struct {
		client transport.HTTPClient
	}
)

func NewEnrollmentSettings(client transport.HTTPClient) *EnrollmentSettings {
	return &EnrollmentSettings{client: client}
}

// GetV4 retrieves the current enrollment settings.
// URL: GET /api/v4/enrollment
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
func (s *EnrollmentSettings) GetV4(ctx context.Context) (*ResourceEnrollmentSettingsV4, *resty.Response, error) {
	var result ResourceEnrollmentSettingsV4
	endpoint := constants.EndpointJamfProEnrollmentSettingsV4
	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
