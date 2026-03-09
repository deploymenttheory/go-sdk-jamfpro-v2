package conditional_access

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Conditional Access-related methods of the Jamf Pro API.
	ConditionalAccess struct {
		client transport.HTTPClient
	}
)

func NewConditionalAccess(client transport.HTTPClient) *ConditionalAccess {
	return &ConditionalAccess{client: client}
}

// GetDeviceComplianceFeatureToggleV1 returns the enablement state of the device compliance feature for Conditional Access.
// URL: GET /api/v1/conditional-access/device-compliance/feature-toggle
// Note: This endpoint is undocumented.
func (s *ConditionalAccess) GetDeviceComplianceFeatureToggleV1(ctx context.Context) (*ResourceDeviceComplianceStatus, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/device-compliance/feature-toggle", constants.EndpointJamfProConditionalAccessV1)

	var result ResourceDeviceComplianceStatus

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDeviceComplianceInformationComputerV1 returns compliance information for a computer device.
// URL: GET /api/v1/conditional-access/device-compliance-information/computer/{deviceId}
func (s *ConditionalAccess) GetDeviceComplianceInformationComputerV1(ctx context.Context, deviceId string) ([]ResourceDeviceComplianceInfo, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/device-compliance-information/computer/%s", constants.EndpointJamfProConditionalAccessV1, deviceId)

	var result []ResourceDeviceComplianceInfo

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetDeviceComplianceInformationMobileV1 returns compliance information for a mobile device.
// URL: GET /api/v1/conditional-access/device-compliance-information/mobile/{deviceId}
func (s *ConditionalAccess) GetDeviceComplianceInformationMobileV1(ctx context.Context, deviceId string) ([]ResourceDeviceComplianceInfo, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/device-compliance-information/mobile/%s", constants.EndpointJamfProConditionalAccessV1, deviceId)

	var result []ResourceDeviceComplianceInfo

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
