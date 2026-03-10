package cloud_idp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Cloud Identity Provider-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
	CloudIdp struct {
		client client.Client
	}
)

func NewCloudIdp(client client.Client) *CloudIdp {
	return &CloudIdp{client: client}
}

// ListV1 returns all Cloud Identity Provider configurations with automatic pagination.
// URL: GET /api/v1/cloud-idp
// query supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp
func (s *CloudIdp) ListV1(ctx context.Context, query map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var items []ResourceCloudIdProvider
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	endpoint := constants.EndpointJamfProCloudIdpV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the Cloud Identity Provider configuration by ID.
// URL: GET /api/v1/cloud-idp/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id
func (s *CloudIdp) GetByIDV1(ctx context.Context, id string) (*ResourceCloudIdProviderDetails, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCloudIdpV1, id)

	var result ResourceCloudIdProviderDetails

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV1 returns the Cloud Identity Provider configuration by display name.
// URL: GET /api/v1/cloud-idp (searches first page)
// Undocumented
func (s *CloudIdp) GetByNameV1(ctx context.Context, name string) (*ResourceCloudIdProviderDetails, *resty.Response, error) {
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
// Accept header: pass constants.TextCSV for CSV export, constants.ApplicationJSON or omit for JSON (default).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-export
func (s *CloudIdp) ExportV1(ctx context.Context, query map[string]string, request *ExportRequest, accept ...string) (*resty.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/export", constants.EndpointJamfProCloudIdpV1)

	acceptHeader := constants.ApplicationJSON
	if len(accept) > 0 && accept[0] != "" {
		acceptHeader = accept[0]
	}

	reqBuilder := s.client.NewRequest(ctx).
		SetHeader("Accept", acceptHeader).
		SetHeader("Content-Type", constants.ApplicationJSON)

	if request != nil {
		reqBuilder = reqBuilder.SetBody(request)
	} else if query != nil {
		reqBuilder = reqBuilder.SetQueryParams(query)
	}

	resp, err := reqBuilder.Post(endpoint)
	if err != nil {
		return resp, nil, err
	}

	return resp, resp.Bytes(), nil
}

// GetHistoryByIDV1 returns the history for a Cloud Identity Provider configuration.
// URL: GET /api/v1/cloud-idp/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-idp-id-history
func (s *CloudIdp) GetHistoryByIDV1(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProCloudIdpV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var items []HistoryItem
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get cloud IDP history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a note to the history for a Cloud Identity Provider configuration.
// URL: POST /api/v1/cloud-idp/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-history
func (s *CloudIdp) AddHistoryNoteByIDV1(ctx context.Context, id string, request *HistoryNoteRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProCloudIdpV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// TestGroupSearchByIDV1 performs a test group search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-group
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-group
func (s *CloudIdp) TestGroupSearchByIDV1(ctx context.Context, id string, request *TestGroupSearchRequest) (*TestGroupSearchResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-group", constants.EndpointJamfProCloudIdpV1, id)

	var result TestGroupSearchResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// TestUserSearchByIDV1 performs a test user search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-user
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user
func (s *CloudIdp) TestUserSearchByIDV1(ctx context.Context, id string, request *TestUserSearchRequest) (*TestUserSearchResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-user", constants.EndpointJamfProCloudIdpV1, id)

	var result TestUserSearchResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// TestUserMembershipByIDV1 performs a test user membership search to verify configuration and mappings.
// URL: POST /api/v1/cloud-idp/{id}/test-user-membership
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-idp-id-test-user-membership
func (s *CloudIdp) TestUserMembershipByIDV1(ctx context.Context, id string, request *TestUserMembershipRequest) (*TestUserMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/test-user-membership", constants.EndpointJamfProCloudIdpV1, id)

	var result TestUserMembershipResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
