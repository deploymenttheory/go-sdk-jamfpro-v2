package policies

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// PoliciesServiceInterface defines the interface for Classic API policy operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/policies
	PoliciesServiceInterface interface {
		// List returns all policies.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpolicies
		List(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified policy by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbyid
		GetByID(ctx context.Context, id int) (*ResourcePolicy, *resty.Response, error)

		// GetByName returns the specified policy by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbyname
		GetByName(ctx context.Context, name string) (*ResourcePolicy, *resty.Response, error)

		// Create creates a new policy.
		//
		// Returns the created policy ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createpolicybyid
		Create(ctx context.Context, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByID updates the specified policy by ID.
		//
		// Returns the updated policy ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatepolicybyid
		UpdateByID(ctx context.Context, id int, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error)

		// UpdateByName updates the specified policy by name.
		//
		// Returns the updated policy ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatepolicybyname
		UpdateByName(ctx context.Context, name string, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error)

		// DeleteByID removes the specified policy by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletepolicybyid
		DeleteByID(ctx context.Context, id int) (*resty.Response, error)

		// DeleteByName removes the specified policy by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletepolicybyname
		DeleteByName(ctx context.Context, name string) (*resty.Response, error)

		// GetByCreatedBy returns all policies filtered by creator type.
		// Valid values are "jss" (GUI/API) or "casper" (Casper Remote).
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbytype
		GetByCreatedBy(ctx context.Context, createdBy string) (*ListResponse, *resty.Response, error)

		// GetByCategory returns all policies in the specified category.
		// Category may be specified by ID, name, or "None" for policies with no category.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbycategory
		GetByCategory(ctx context.Context, category string) (*ListResponse, *resty.Response, error)

		// GetByIDWithSubset returns a subset of data for the specified policy by ID.
		// Valid subsets: General, Scope, SelfService, PackageConfiguration, Scripts,
		// Printers, DockItems, AccountMaintenance, Reboot, Maintenance, FilesProcesses,
		// UserInteraction, DiskEncryption.
		// Multiple subsets can be combined with ampersand (e.g., "General&Scope").
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbyidsubset
		GetByIDWithSubset(ctx context.Context, id int, subset string) (*ResourcePolicy, *resty.Response, error)

		// GetByNameWithSubset returns a subset of data for the specified policy by name.
		// Valid subsets: General, Scope, SelfService, PackageConfiguration, Scripts,
		// Printers, DockItems, AccountMaintenance, Reboot, Maintenance, FilesProcesses,
		// UserInteraction, DiskEncryption.
		// Multiple subsets can be combined with ampersand (e.g., "General&Scope").
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findpoliciesbynamesubset
		GetByNameWithSubset(ctx context.Context, name string, subset string) (*ResourcePolicy, *resty.Response, error)
	}

	// Service handles communication with the policy-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/policies
	Policies struct {
		client interfaces.HTTPClient
	}
)

var _ PoliciesServiceInterface = (*Policies)(nil)

// NewService returns a new policies Service backed by the provided HTTP client.
func NewPolicies(client interfaces.HTTPClient) *Policies {
	return &Policies{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Policies CRUD Operations
// -----------------------------------------------------------------------------

// List returns all policies.
// URL: GET /JSSResource/policies
// https://developer.jamf.com/jamf-pro/reference/findpolicies
func (s *Policies) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicPolicies

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified policy by ID.
// URL: GET /JSSResource/policies/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbyid
func (s *Policies) GetByID(ctx context.Context, id int) (*ResourcePolicy, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("policy ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPolicies, id)

	var result ResourcePolicy

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified policy by name.
// URL: GET /JSSResource/policies/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbyname
func (s *Policies) GetByName(ctx context.Context, name string) (*ResourcePolicy, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("policy name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPolicies, name)

	var result ResourcePolicy

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new policy.
// URL: POST /JSSResource/policies/id/0
// Returns the created policy ID.
// https://developer.jamf.com/jamf-pro/reference/createpolicybyid
func (s *Policies) Create(ctx context.Context, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error) {
	if policy == nil {
		return nil, nil, fmt.Errorf("policy is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicPolicies)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, policy, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified policy by ID.
// URL: PUT /JSSResource/policies/id/{id}
// Returns the updated policy ID.
// https://developer.jamf.com/jamf-pro/reference/updatepolicybyid
func (s *Policies) UpdateByID(ctx context.Context, id int, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("policy ID must be a positive integer")
	}
	if policy == nil {
		return nil, nil, fmt.Errorf("policy is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPolicies, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, policy, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified policy by name.
// URL: PUT /JSSResource/policies/name/{name}
// Returns the updated policy ID.
// https://developer.jamf.com/jamf-pro/reference/updatepolicybyname
func (s *Policies) UpdateByName(ctx context.Context, name string, policy *ResourcePolicy) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("policy name is required")
	}
	if policy == nil {
		return nil, nil, fmt.Errorf("policy is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPolicies, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, policy, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified policy by ID.
// URL: DELETE /JSSResource/policies/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletepolicybyid
func (s *Policies) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("policy ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicPolicies, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified policy by name.
// URL: DELETE /JSSResource/policies/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletepolicybyname
func (s *Policies) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("policy name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicPolicies, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetByCreatedBy returns all policies filtered by creator type.
// Valid values are "jss" (GUI/API) or "casper" (Casper Remote).
// URL: GET /JSSResource/policies/createdBy/{createdBy}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbytype
func (s *Policies) GetByCreatedBy(ctx context.Context, createdBy string) (*ListResponse, *resty.Response, error) {
	if createdBy == "" {
		return nil, nil, fmt.Errorf("createdBy is required")
	}
	if createdBy != "jss" && createdBy != "casper" {
		return nil, nil, fmt.Errorf("createdBy must be 'jss' or 'casper'")
	}

	endpoint := fmt.Sprintf("%s/createdBy/%s", EndpointClassicPolicies, createdBy)

	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByCategory returns all policies in the specified category.
// Category may be specified by ID, name, or "None" for policies with no category.
// URL: GET /JSSResource/policies/category/{category}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbycategory
func (s *Policies) GetByCategory(ctx context.Context, category string) (*ListResponse, *resty.Response, error) {
	if category == "" {
		return nil, nil, fmt.Errorf("category is required")
	}

	endpoint := fmt.Sprintf("%s/category/%s", EndpointClassicPolicies, category)

	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDWithSubset returns a subset of data for the specified policy by ID.
// Valid subsets: General, Scope, SelfService, PackageConfiguration, Scripts,
// Printers, DockItems, AccountMaintenance, Reboot, Maintenance, FilesProcesses,
// UserInteraction, DiskEncryption.
// Multiple subsets can be combined with ampersand (e.g., "General&Scope").
// URL: GET /JSSResource/policies/id/{id}/subset/{subset}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbyidsubset
func (s *Policies) GetByIDWithSubset(ctx context.Context, id int, subset string) (*ResourcePolicy, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("policy ID must be a positive integer")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", EndpointClassicPolicies, id, subset)

	var result ResourcePolicy

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameWithSubset returns a subset of data for the specified policy by name.
// Valid subsets: General, Scope, SelfService, PackageConfiguration, Scripts,
// Printers, DockItems, AccountMaintenance, Reboot, Maintenance, FilesProcesses,
// UserInteraction, DiskEncryption.
// Multiple subsets can be combined with ampersand (e.g., "General&Scope").
// URL: GET /JSSResource/policies/name/{name}/subset/{subset}
// https://developer.jamf.com/jamf-pro/reference/findpoliciesbynamesubset
func (s *Policies) GetByNameWithSubset(ctx context.Context, name string, subset string) (*ResourcePolicy, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("policy name is required")
	}
	if subset == "" {
		return nil, nil, fmt.Errorf("subset is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", EndpointClassicPolicies, name, subset)

	var result ResourcePolicy

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
