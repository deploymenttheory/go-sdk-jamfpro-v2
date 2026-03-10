package adcs_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the AD CS Settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
	AdcsSettings struct {
		client client.Client
	}
)

func NewAdcsSettings(client client.Client) *AdcsSettings {
	return &AdcsSettings{client: client}
}

// CreateV1 creates a new AD CS configuration.
// URL: POST /api/v1/pki/adcs-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
func (s *AdcsSettings) CreateV1(ctx context.Context, request *ResourceAdcsSettings) (*ResponseAdcsSettingsCreated, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProAdcsSettingsV1

	var result ResponseAdcsSettingsCreated

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.AcceptAny).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the AD CS configuration by ID.
// URL: GET /api/v1/pki/adcs-settings/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id
func (s *AdcsSettings) GetByIDV1(ctx context.Context, id string) (*ResponseAdcsSettings, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdcsSettingsV1, id)

	var result ResponseAdcsSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the AD CS configuration by ID using merge-patch semantics.
// URL: PATCH /api/v1/pki/adcs-settings/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-pki-adcs-settings-id
func (s *AdcsSettings) UpdateByIDV1(ctx context.Context, id string, request *ResourceAdcsSettings) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdcsSettingsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationMergePatchJSON).
		SetBody(request).
		Patch(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByIDV1 deletes the AD CS configuration by ID.
// URL: DELETE /api/v1/pki/adcs-settings/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-adcs-settings-id
func (s *AdcsSettings) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdcsSettingsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ValidateServerCertificateV1 validates the AD CS Settings server certificate.
// URL: POST /api/v1/pki/adcs-settings/validate-certificate
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-validate-certificate
func (s *AdcsSettings) ValidateServerCertificateV1(ctx context.Context, request *ValidateCertificateRequest) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/validate-certificate", constants.EndpointJamfProAdcsSettingsV1)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ValidateClientCertificateV1 validates the AD CS Settings client certificate.
// URL: POST /api/v1/pki/adcs-settings/validate-client-certificate
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-validate-client-certificate
func (s *AdcsSettings) ValidateClientCertificateV1(ctx context.Context, request *ValidateCertificateRequest) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/validate-client-certificate", constants.EndpointJamfProAdcsSettingsV1)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetDependenciesByIDV1 returns the list of dependencies for an AD CS Settings configuration.
// URL: GET /api/v1/pki/adcs-settings/{id}/dependencies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id-dependencies
func (s *AdcsSettings) GetDependenciesByIDV1(ctx context.Context, id string) (*DependenciesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", constants.EndpointJamfProAdcsSettingsV1, id)

	var result DependenciesResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryByIDV1 returns the history for an AD CS Settings configuration with automatic pagination.
// URL: GET /api/v1/pki/adcs-settings/{id}/history
// query supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id-history
func (s *AdcsSettings) GetHistoryByIDV1(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProAdcsSettingsV1, id)

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
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a note to the history for an AD CS Settings configuration.
// URL: POST /api/v1/pki/adcs-settings/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-id-history
func (s *AdcsSettings) AddHistoryNoteByIDV1(ctx context.Context, id string, request *HistoryNoteRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProAdcsSettingsV1, id)

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
