package patch_policies

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// PatchPoliciesServiceInterface defines the interface for patch policy operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
	PatchPoliciesServiceInterface interface {
		// ListV2 returns a list of all patch policies.
		//
		// This endpoint retrieves all patch policies with their details.
		// The v2 API is read-only for patch policies. Create/Update/Delete operations
		// require the Classic API (XML).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
		ListV2(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByIDV2 returns the patch policy by ID.
		//
		// This method retrieves all patch policies and filters by ID.
		// The v2 API does not provide a direct endpoint for fetching by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
		GetByIDV2(ctx context.Context, id string) (*ResourcePatchPolicy, *interfaces.Response, error)

		// GetByNameV2 returns the patch policy by name.
		//
		// This method retrieves all patch policies and filters by name.
		// The v2 API does not provide a direct endpoint for fetching by name.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
		GetByNameV2(ctx context.Context, name string) (*ResourcePatchPolicy, *interfaces.Response, error)

		// GetDashboardStatusV2 checks if a patch policy is on the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-id-dashboard
		GetDashboardStatusV2(ctx context.Context, id string) (*DashboardStatusResponse, *interfaces.Response, error)

		// AddToDashboardV2 adds a patch policy to the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-policies-id-dashboard
		AddToDashboardV2(ctx context.Context, id string) (*interfaces.Response, error)

		// RemoveFromDashboardV2 removes a patch policy from the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-policies-id-dashboard
		RemoveFromDashboardV2(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the patch policies-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PatchPoliciesServiceInterface = (*Service)(nil)

// NewService returns a new patch policies Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Policies CRUD Operations (v2 - Read Only)
// -----------------------------------------------------------------------------

// ListV2 returns a list of all patch policies.
// URL: GET /api/v2/patch-policies/policy-details
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
func (s *Service) ListV2(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointPatchPoliciesV2 + "/policy-details"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourcePatchPolicy
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// GetByIDV2 returns the patch policy by ID.
// URL: GET /api/v2/patch-policies/policy-details (filtered by ID)
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
func (s *Service) GetByIDV2(ctx context.Context, id string) (*ResourcePatchPolicy, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	policies, resp, err := s.ListV2(ctx)
	if err != nil {
		return nil, resp, err
	}

	for _, policy := range policies.Results {
		if policy.ID == id {
			return &policy, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("patch policy not found with ID: %s", id)
}

// GetByNameV2 returns the patch policy by name.
// URL: GET /api/v2/patch-policies/policy-details (filtered by name)
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
func (s *Service) GetByNameV2(ctx context.Context, name string) (*ResourcePatchPolicy, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	policies, resp, err := s.ListV2(ctx)
	if err != nil {
		return nil, resp, err
	}

	for _, policy := range policies.Results {
		if policy.Name == name {
			return &policy, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("patch policy not found with name: %s", name)
}

// GetDashboardStatusV2 checks if a patch policy is on the dashboard.
// URL: GET /api/v2/patch-policies/{id}/dashboard
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-id-dashboard
func (s *Service) GetDashboardStatusV2(ctx context.Context, id string) (*DashboardStatusResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", EndpointPatchPoliciesV2, id)

	var result DashboardStatusResponse

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

// AddToDashboardV2 adds a patch policy to the dashboard.
// URL: POST /api/v2/patch-policies/{id}/dashboard
// https://developer.jamf.com/jamf-pro/reference/post_v2-patch-policies-id-dashboard
func (s *Service) AddToDashboardV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", EndpointPatchPoliciesV2, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveFromDashboardV2 removes a patch policy from the dashboard.
// URL: DELETE /api/v2/patch-policies/{id}/dashboard
// https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-policies-id-dashboard
func (s *Service) RemoveFromDashboardV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", EndpointPatchPoliciesV2, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
