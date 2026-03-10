package policy_properties

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the policy properties-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
	PolicyProperties struct {
		client client.Client
	}
)

func NewPolicyProperties(client client.Client) *PolicyProperties {
	return &PolicyProperties{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Policy Properties Operations
// -----------------------------------------------------------------------------

// Get returns the current policy properties.
// URL: GET /api/v1/policy-properties
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
func (s *PolicyProperties) Get(ctx context.Context) (*ResourcePolicyProperties, *resty.Response, error) {
	var result ResourcePolicyProperties
	endpoint := constants.EndpointJamfProPolicyPropertiesV1
	headers := map[string]string{"Accept": constants.ApplicationJSON, "Content-Type": constants.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// Update updates policy properties.
// URL: PUT /api/v1/policy-properties
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-policy-properties
func (s *PolicyProperties) Update(ctx context.Context, request *ResourcePolicyProperties) (*ResourcePolicyProperties, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourcePolicyProperties
	endpoint := constants.EndpointJamfProPolicyPropertiesV1
	headers := map[string]string{"Accept": constants.ApplicationJSON, "Content-Type": constants.ApplicationJSON}
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
