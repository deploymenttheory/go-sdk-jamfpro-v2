package patch_policies

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the patch policies-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
	PatchPolicies struct {
		client transport.HTTPClient
	}
)

// NewService returns a new patch policies Service backed by the provided HTTP client.
func NewPatchPolicies(client transport.HTTPClient) *PatchPolicies {
	return &PatchPolicies{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Policies CRUD Operations (v2 - Read Only)
// -----------------------------------------------------------------------------

// ListV2 returns a list of all patch policies with full details.
// URL: GET /api/v2/patch-policies/policy-details
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
func (s *PatchPolicies) ListV2(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProPatchPoliciesPolicyDetails

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageItems []ResourcePatchPolicy
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// ListSummaryV2 returns a list of patch policy summaries.
// URL: GET /api/v2/patch-policies
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies
func (s *PatchPolicies) ListSummaryV2(ctx context.Context, rsqlQuery map[string]string) (*ListSummaryResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProPatchPoliciesV2

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result ListSummaryResponse
	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list patch policy summaries: %w", err)
	}

	return &result, resp, nil
}

// GetByIDV2 returns the patch policy by ID.
// URL: GET /api/v2/patch-policies/policy-details (filtered by ID)
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-policies-policy-details
func (s *PatchPolicies) GetByIDV2(ctx context.Context, id string) (*ResourcePatchPolicy, *resty.Response, error) {
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
func (s *PatchPolicies) GetByNameV2(ctx context.Context, name string) (*ResourcePatchPolicy, *resty.Response, error) {
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
func (s *PatchPolicies) GetDashboardStatusV2(ctx context.Context, id string) (*DashboardStatusResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchPoliciesV2, id)

	var result DashboardStatusResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *PatchPolicies) AddToDashboardV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchPoliciesV2, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *PatchPolicies) RemoveFromDashboardV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchPoliciesV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
