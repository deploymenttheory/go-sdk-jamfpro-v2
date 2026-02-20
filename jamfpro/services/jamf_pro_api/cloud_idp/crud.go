package cloud_idp

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CloudIdpServiceInterface defines the interface for Cloud Identity Provider operations.
	// Uses v1 API for all operations. Supports listing, exporting, testing, and history operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
	CloudIdpServiceInterface interface {
		// ListV1 returns all Cloud Identity Provider configurations (Get Cloud Identity Providers).
		//
		// Query params (optional, pass via query): page, page-size, sort, filter.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
		ListV1(ctx context.Context, query map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the Cloud Identity Provider configuration by ID (Get Cloud Identity Provider Configuration by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id
		GetByIDV1(ctx context.Context, id string) (*ResourceCloudIdProviderDetails, *interfaces.Response, error)

		// GetByNameV1 returns the Cloud Identity Provider configuration by display name (searches first page of ListV1).
		GetByNameV1(ctx context.Context, name string) (*ResourceCloudIdProviderDetails, *interfaces.Response, error)

		// ExportV1 exports Cloud Identity Providers collection (Export Cloud Identity Providers).
		//
		// Query params (optional, pass via query): export-fields, export-labels, page, page-size, sort, filter.
		// Request body can override query parameters to avoid URI length limits.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-export
		ExportV1(ctx context.Context, query map[string]string, request *ExportRequest) (*interfaces.Response, []byte, error)

		// GetHistoryByIDV1 returns the history for a Cloud Identity Provider configuration (Get History).
		//
		// Query params (optional, pass via query): page, page-size, sort, filter.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id-history
		GetHistoryByIDV1(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNoteByIDV1 adds a note to the history for a Cloud Identity Provider configuration (Add History Note).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-history
		AddHistoryNoteByIDV1(ctx context.Context, id string, request *HistoryNoteRequest) (*interfaces.Response, error)

		// TestGroupSearchByIDV1 performs a test group search to verify configuration and mappings (Test Group Search).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-group
		TestGroupSearchByIDV1(ctx context.Context, id string, request *TestGroupSearchRequest) (*TestGroupSearchResponse, *interfaces.Response, error)

		// TestUserSearchByIDV1 performs a test user search to verify configuration and mappings (Test User Search).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user
		TestUserSearchByIDV1(ctx context.Context, id string, request *TestUserSearchRequest) (*TestUserSearchResponse, *interfaces.Response, error)

		// TestUserMembershipByIDV1 performs a test user membership search to verify configuration and mappings (Test User Membership).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user-membership
		TestUserMembershipByIDV1(ctx context.Context, id string, request *TestUserMembershipRequest) (*TestUserMembershipResponse, *interfaces.Response, error)
	}

	// Service handles communication with the Cloud Identity Provider-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CloudIdpServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all Cloud Identity Provider configurations.
// URL: GET /api/v1/cloud-idp
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
func (s *Service) ListV1(ctx context.Context, query map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointCloudIdpV1, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the Cloud Identity Provider configuration by ID.
// URL: GET /api/v1/cloud-idp/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceCloudIdProviderDetails, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudIdpV1, id)

	var result ResourceCloudIdProviderDetails

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

// GetByNameV1 returns the Cloud Identity Provider configuration by display name.
// URL: GET /api/v1/cloud-idp (searches first page)
func (s *Service) GetByNameV1(ctx context.Context, name string) (*ResourceCloudIdProviderDetails, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	list, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, provider := range list.Results {
		if provider.DisplayName == name {
			return s.GetByIDV1(ctx, provider.ID)
		}
	}

	return nil, resp, fmt.Errorf("cloud identity provider with name %q not found", name)
}

// ExportV1 exports Cloud Identity Providers collection.
// URL: POST /api/v1/cloud-idp/export
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-export
func (s *Service) ExportV1(ctx context.Context, query map[string]string, request *ExportRequest) (*interfaces.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/export", EndpointCloudIdpV1)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var resp *interfaces.Response
	var err error

	if request != nil {
		resp, err = s.client.Post(ctx, endpoint, request, headers, nil)
	} else {
		resp, err = s.client.PostWithQuery(ctx, endpoint, query, nil, headers, nil)
	}

	if err != nil {
		return resp, nil, err
	}

	return resp, resp.Body, nil
}

// GetHistoryByIDV1 returns the history for a Cloud Identity Provider configuration.
// URL: GET /api/v1/cloud-idp/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id-history
func (s *Service) GetHistoryByIDV1(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCloudIdpV1, id)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a note to the history for a Cloud Identity Provider configuration.
// URL: POST /api/v1/cloud-idp/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-history
func (s *Service) AddHistoryNoteByIDV1(ctx context.Context, id string, request *HistoryNoteRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointCloudIdpV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// TestGroupSearchByIDV1 performs a test group search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-group
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-group
func (s *Service) TestGroupSearchByIDV1(ctx context.Context, id string, request *TestGroupSearchRequest) (*TestGroupSearchResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-group", EndpointCloudIdpV1, id)

	var result TestGroupSearchResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// TestUserSearchByIDV1 performs a test user search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-user
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user
func (s *Service) TestUserSearchByIDV1(ctx context.Context, id string, request *TestUserSearchRequest) (*TestUserSearchResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-user", EndpointCloudIdpV1, id)

	var result TestUserSearchResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// TestUserMembershipByIDV1 performs a test user membership search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-user-membership
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user-membership
func (s *Service) TestUserMembershipByIDV1(ctx context.Context, id string, request *TestUserMembershipRequest) (*TestUserMembershipResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-user-membership", EndpointCloudIdpV1, id)

	var result TestUserMembershipResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
