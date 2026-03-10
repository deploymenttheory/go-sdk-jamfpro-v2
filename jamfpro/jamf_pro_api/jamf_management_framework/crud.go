package jamf_management_framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Management Framework-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-management-framework-redeploy-id
	JamfManagementFramework struct {
		client client.Client
	}
)

func NewJamfManagementFramework(client client.Client) *JamfManagementFramework {
	return &JamfManagementFramework{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Management Framework Operations
// -----------------------------------------------------------------------------

// RedeployV1 redeploys the Jamf Management Framework for an enrolled device.
// URL: POST /api/v1/jamf-management-framework/redeploy/{id}
// Path param: id (computer ID)
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-management-framework-redeploy-id
func (s *JamfManagementFramework) RedeployV1(ctx context.Context, computerID string) (*RedeployResponse, *resty.Response, error) {
	id := strings.TrimSpace(computerID)
	if id == "" {
		return nil, nil, fmt.Errorf("computer ID is required")
	}

	endpoint := fmt.Sprintf("%s/redeploy/%s", constants.EndpointJamfProJamfManagementFrameworkV1, id)

	var result RedeployResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to redeploy jamf management framework: %w", err)
	}

	return &result, resp, nil
}
