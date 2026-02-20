package policy_properties

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// PolicyPropertiesServiceInterface defines the interface for policy properties operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
	PolicyPropertiesServiceInterface interface {
		// Get returns the current policy properties (Get Policy Properties).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
		Get(ctx context.Context) (*ResourcePolicyProperties, *interfaces.Response, error)

		// Update updates policy properties (Update Policy Properties).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-policy-properties
		Update(ctx context.Context, request *ResourcePolicyProperties) (*ResourcePolicyProperties, *interfaces.Response, error)
	}

	// Service handles communication with the policy properties-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PolicyPropertiesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Policy Properties Operations
// -----------------------------------------------------------------------------

// Get returns the current policy properties.
// URL: GET /api/v1/policy-properties
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
func (s *Service) Get(ctx context.Context) (*ResourcePolicyProperties, *interfaces.Response, error) {
	var result ResourcePolicyProperties
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointPolicyPropertiesV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// Update updates policy properties.
// URL: PUT /api/v1/policy-properties
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-policy-properties
func (s *Service) Update(ctx context.Context, request *ResourcePolicyProperties) (*ResourcePolicyProperties, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourcePolicyProperties
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, EndpointPolicyPropertiesV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
