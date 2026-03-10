package patch_management

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Patch Management-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-management-accept-disclaimer
	PatchManagement struct {
		client client.Client
	}
)

// NewService creates a new Patch Management service.
func NewPatchManagement(client client.Client) *PatchManagement {
	return &PatchManagement{client: client}
}

// AcceptDisclaimerV2 accepts the Patch Management disclaimer.
// URL: POST /api/v2/patch-management-accept-disclaimer
// Must be called before patch management features can be used.
// Performs a POST with no request body and returns no response data.
//
// https://developer.jamf.com/jamf-pro/reference/post_v2-patch-management-accept-disclaimer
func (s *PatchManagement) AcceptDisclaimerV2(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProPatchManagementAcceptDisclaimerV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to accept Patch Management disclaimer: %w", err)
	}

	return resp, nil
}
