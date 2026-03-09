package jamf_management_framework

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// JamfManagementFrameworkServiceInterface defines the interface for Jamf Management Framework operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-management-framework-redeploy-id
	JamfManagementFrameworkServiceInterface interface {
		// RedeployV1 redeploys the Jamf Management Framework for an enrolled device.
		//
		// POST /api/v1/jamf-management-framework/redeploy/{id}
		//
		// Returns 201 Created with deviceId and commandUuid when the command is successfully queued.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-management-framework-redeploy-id
		RedeployV1(ctx context.Context, computerID string) (*RedeployResponse, *resty.Response, error)
	}

	// Service handles communication with the Jamf Management Framework-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-management-framework-redeploy-id
	JamfManagementFramework struct {
		client interfaces.HTTPClient
	}
)

var _ JamfManagementFrameworkServiceInterface = (*JamfManagementFramework)(nil)

func NewJamfManagementFramework(client interfaces.HTTPClient) *JamfManagementFramework {
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

	endpoint := fmt.Sprintf("%s/redeploy/%s", EndpointJamfManagementFrameworkV1, id)

	var result RedeployResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to redeploy jamf management framework: %w", err)
	}

	return &result, resp, nil
}
