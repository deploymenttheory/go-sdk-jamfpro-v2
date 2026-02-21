package conditional_access

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ConditionalAccessServiceInterface defines the interface for Conditional Access operations.
	// Uses v1 API. Manages device compliance feature enablement for Conditional Access.
	//
	// Note: This API is undocumented in the official Jamf Pro API reference.
	ConditionalAccessServiceInterface interface {
		// GetDeviceComplianceFeatureToggleV1 returns the enablement state of the device compliance feature for Conditional Access.
		//
		// Note: This endpoint is undocumented.
		GetDeviceComplianceFeatureToggleV1(ctx context.Context) (*ResourceDeviceComplianceStatus, *interfaces.Response, error)
	}

	// Service handles communication with the Conditional Access-related methods of the Jamf Pro API.
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ConditionalAccessServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetDeviceComplianceFeatureToggleV1 returns the enablement state of the device compliance feature for Conditional Access.
// URL: GET /api/v1/conditional-access/device-compliance/feature-toggle
// Note: This endpoint is undocumented.
func (s *Service) GetDeviceComplianceFeatureToggleV1(ctx context.Context) (*ResourceDeviceComplianceStatus, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/device-compliance/feature-toggle", EndpointConditionalAccessV1)

	var result ResourceDeviceComplianceStatus

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
