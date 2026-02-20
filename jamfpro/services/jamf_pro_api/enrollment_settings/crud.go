package enrollment_settings

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// EnrollmentSettingsServiceInterface defines the interface for enrollment settings operations (v4).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
	EnrollmentSettingsServiceInterface interface {
		// GetV4 retrieves the current enrollment settings (Get Enrollment Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
		GetV4(ctx context.Context) (*ResourceEnrollmentSettingsV4, *interfaces.Response, error)
	}

	// Service handles communication with the enrollment settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ EnrollmentSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV4 retrieves the current enrollment settings.
// URL: GET /api/v4/enrollment
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
func (s *Service) GetV4(ctx context.Context) (*ResourceEnrollmentSettingsV4, *interfaces.Response, error) {
	var result ResourceEnrollmentSettingsV4
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointEnrollmentSettingsV4, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
