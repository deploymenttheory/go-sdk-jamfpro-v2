package adcs_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// AdcsSettingsServiceInterface defines the interface for AD CS Settings operations.
	// Uses v1 API for all operations. Supports certificate management for PKI integration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
	AdcsSettingsServiceInterface interface {
		// CreateV1 creates a new AD CS configuration (Create AD CS Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
		CreateV1(ctx context.Context, request *ResourceAdcsSettings) (*ResponseAdcsSettingsCreated, *resty.Response, error)

		// GetByIDV1 returns the AD CS configuration by ID (Get AD CS Settings by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id
		GetByIDV1(ctx context.Context, id string) (*ResponseAdcsSettings, *resty.Response, error)

		// UpdateByIDV1 updates the AD CS configuration by ID using merge-patch semantics (Update AD CS Settings by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-pki-adcs-settings-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceAdcsSettings) (*resty.Response, error)

		// DeleteByIDV1 deletes the AD CS configuration by ID (Delete AD CS Settings by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-adcs-settings-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// ValidateServerCertificateV1 validates the AD CS Settings server certificate (Validate Server Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-validate-certificate
		ValidateServerCertificateV1(ctx context.Context, request *ValidateCertificateRequest) (*resty.Response, error)

		// ValidateClientCertificateV1 validates the AD CS Settings client certificate (Validate Client Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-validate-client-certificate
		ValidateClientCertificateV1(ctx context.Context, request *ValidateCertificateRequest) (*resty.Response, error)

		// GetDependenciesByIDV1 returns the list of dependencies for an AD CS Settings configuration (Get Dependencies).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id-dependencies
		GetDependenciesByIDV1(ctx context.Context, id string) (*DependenciesResponse, *resty.Response, error)

		// GetHistoryByIDV1 returns the history for an AD CS Settings configuration (Get History).
		//
		// Query params (optional, pass via query): page, page-size, sort, filter.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-adcs-settings-id-history
		GetHistoryByIDV1(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNoteByIDV1 adds a note to the history for an AD CS Settings configuration (Add History Note).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings-id-history
		AddHistoryNoteByIDV1(ctx context.Context, id string, request *HistoryNoteRequest) (*resty.Response, error)
	}

	// Service handles communication with the AD CS Settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-adcs-settings
	AdcsSettings struct {
		client transport.HTTPClient
	}
)

var _ AdcsSettingsServiceInterface = (*AdcsSettings)(nil)

func NewAdcsSettings(client transport.HTTPClient) *AdcsSettings {
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

	headers := map[string]string{
		"Accept": constants.AcceptAny,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationMergePatchJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, nil)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
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

	headers := map[string]string{
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
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

	headers := map[string]string{
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var items []HistoryItem
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
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

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
