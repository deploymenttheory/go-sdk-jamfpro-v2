package mobile_device_enrollment_profile

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// MobileDeviceEnrollmentProfileServiceInterface defines the interface for mobile device enrollment profile operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-enrollment-profile-id-download-profile
	MobileDeviceEnrollmentProfileServiceInterface interface {
		// GetDownloadProfileV1 retrieves the MDM Enrollment Profile for the specified device as a binary file
		// (Apple Aspen config format).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-enrollment-profile-id-download-profile
		GetDownloadProfileV1(ctx context.Context, id string) ([]byte, *resty.Response, error)
	}

	// Service handles communication with the mobile device enrollment profile-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-enrollment-profile-id-download-profile
	MobileDeviceEnrollmentProfile struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDeviceEnrollmentProfileServiceInterface = (*MobileDeviceEnrollmentProfile)(nil)

// NewService returns a new mobile device enrollment profile Service backed by the provided HTTP client.
func NewMobileDeviceEnrollmentProfile(client interfaces.HTTPClient) *MobileDeviceEnrollmentProfile {
	return &MobileDeviceEnrollmentProfile{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Mobile Device Enrollment Profile Operations
// -----------------------------------------------------------------------------

// GetDownloadProfileV1 retrieves the MDM Enrollment Profile for the specified device.
// URL: GET /api/v1/mobile-device-enrollment-profile/{id}/download-profile
// Response: 200 OK with binary body (application/x-apple-aspen-config).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-enrollment-profile-id-download-profile
func (s *MobileDeviceEnrollmentProfile) GetDownloadProfileV1(ctx context.Context, id string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download-profile", EndpointMobileDeviceEnrollmentProfileV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationXAppleAspenConfig,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to download profile for ID %s: %w", id, err)
	}

	return data, resp, nil
}
