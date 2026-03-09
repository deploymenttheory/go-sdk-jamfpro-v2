package patch_software_title_configurations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// PatchSoftwareTitleConfigurationsServiceInterface defines the interface for patch software title configuration operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
	PatchSoftwareTitleConfigurationsServiceInterface interface {
		// ListV2 returns all patch software title configurations.
		//
		// This endpoint retrieves all patch software title configurations with their details.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
		ListV2(ctx context.Context) (*ListResponse, *resty.Response, error)

		// GetByIDV2 returns the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
		GetByIDV2(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error)

		// GetByNameV2 returns the patch software title configuration by display name.
		//
		// This is a convenience method that calls ListV2 and filters by DisplayName.
		GetByNameV2(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error)

		// CreateV2 creates a new patch software title configuration.
		// Returns CreateResponse (id, href).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-software-title-configurations
		CreateV2(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV2 updates the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-patch-software-title-configurations-id
		UpdateByIDV2(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error)

		// UpdateByNameV2 updates the patch software title configuration by display name.
		UpdateByNameV2(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error)

		// DeleteByIDV2 deletes the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-software-title-configurations-id
		DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error)

		// DeleteByNameV2 deletes the patch software title configuration by display name.
		DeleteByNameV2(ctx context.Context, name string) (*resty.Response, error)

		// GetDashboardStatusByIDV2 returns whether the software title configuration is on the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-dashboard
		GetDashboardStatusByIDV2(ctx context.Context, id string) (*ResourceDashboardStatus, *resty.Response, error)

		// AddToDashboardByIDV2 adds the software title configuration to the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-software-title-configurations-id-dashboard
		AddToDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error)

		// RemoveFromDashboardByIDV2 removes the software title configuration from the dashboard.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-software-title-configurations-id-dashboard
		RemoveFromDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error)

		// GetDefinitionsByIDV2 returns paginated patch software title definitions.
		//
		// Query params: page, page-size, sort, filter (RSQL).
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-definitions
		GetDefinitionsByIDV2(ctx context.Context, id string, query map[string]string) (*DefinitionsResponse, *resty.Response, error)

		// GetDependenciesByIDV2 returns paginated dependencies (smart groups) for the configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-dependencies
		GetDependenciesByIDV2(ctx context.Context, id string, query map[string]string) (*DependenciesResponse, *resty.Response, error)

		// ExportReportByIDV2 exports patch reporting data as CSV bytes.
		//
		// Query params: filter (RSQL), columns-to-export (comma-separated column names).
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-export-report
		ExportReportByIDV2(ctx context.Context, id string, query map[string]string) ([]byte, *resty.Response, error)

		// GetExtensionAttributesByIDV2 returns extension attributes for the software title.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-extension-attributes
		GetExtensionAttributesByIDV2(ctx context.Context, id string) ([]ResourceExtensionAttribute, *resty.Response, error)

		// GetPatchReportByIDV2 returns paginated patch report for the configuration.
		//
		// Query params: page, page-size, sort, filter (RSQL).
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-patch-report
		GetPatchReportByIDV2(ctx context.Context, id string, query map[string]string) (*PatchReportResponse, *resty.Response, error)

		// GetPatchSummaryByIDV2 returns the patch summary for the configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-patch-summary
		GetPatchSummaryByIDV2(ctx context.Context, id string) (*ResourcePatchSummary, *resty.Response, error)

		// GetHistoryByIDV2 returns paginated history for the configuration.
		//
		// Query params: page, page-size, sort, filter (RSQL).
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-history
		GetHistoryByIDV2(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNoteByIDV2 adds a history note to the configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-software-title-configurations-id-history
		AddHistoryNoteByIDV2(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error)

		// GetPatchVersionsByIDV2 returns patch versions for the configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id-patch-summary-versions
		GetPatchVersionsByIDV2(ctx context.Context, id string) ([]ResourcePatchVersion, *resty.Response, error)
	}

	// Service handles communication with the patch software title configurations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
	PatchSoftwareTitleConfigurations struct {
		client transport.HTTPClient
	}
)

var _ PatchSoftwareTitleConfigurationsServiceInterface = (*PatchSoftwareTitleConfigurations)(nil)

// NewService returns a new patch software title configurations Service backed by the provided HTTP client.
func NewPatchSoftwareTitleConfigurations(client transport.HTTPClient) *PatchSoftwareTitleConfigurations {
	return &PatchSoftwareTitleConfigurations{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Software Title Configurations CRUD Operations (v2)
// -----------------------------------------------------------------------------

// ListV2 returns all patch software title configurations.
// URL: GET /api/v2/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
func (s *PatchSoftwareTitleConfigurations) ListV2(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV2 returns the patch software title configuration by ID.
// URL: GET /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) GetByIDV2(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV2 returns the patch software title configuration by display name.
// This is a convenience method that calls ListV2 and filters by DisplayName.
func (s *PatchSoftwareTitleConfigurations) GetByNameV2(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	list, resp, err := s.ListV2(ctx)
	if err != nil {
		return nil, resp, err
	}

	for i := range *list {
		if (*list)[i].DisplayName == name {
			return &(*list)[i], resp, nil
		}
	}

	return nil, resp, fmt.Errorf("patch software title configuration with name %q not found", name)
}

// CreateV2 creates a new patch software title configuration.
// URL: POST /api/v2/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/post_v2-patch-software-title-configurations
func (s *PatchSoftwareTitleConfigurations) CreateV2(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *resty.Response, error) {
	if config == nil {
		return nil, nil, fmt.Errorf("config is required")
	}

	if config.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	if config.SoftwareTitleID == "" {
		return nil, nil, fmt.Errorf("software title id is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, config, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 updates the patch software title configuration by ID.
// URL: PATCH /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/patch_v2-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) UpdateByIDV2(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if config == nil {
		return nil, nil, fmt.Errorf("config is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, config, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByNameV2 updates the patch software title configuration by display name.
func (s *PatchSoftwareTitleConfigurations) UpdateByNameV2(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV2(ctx, name)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get patch software title configuration by name: %w", err)
	}

	return s.UpdateByIDV2(ctx, target.ID, config)
}

// DeleteByIDV2 deletes the patch software title configuration by ID.
// URL: DELETE /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV2 deletes the patch software title configuration by display name.
func (s *PatchSoftwareTitleConfigurations) DeleteByNameV2(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV2(ctx, name)
	if err != nil {
		return resp, fmt.Errorf("failed to get patch software title configuration by name: %w", err)
	}

	return s.DeleteByIDV2(ctx, target.ID)
}

// -----------------------------------------------------------------------------
// Dashboard Management
// -----------------------------------------------------------------------------

// GetDashboardStatusByIDV2 returns whether the software title configuration is on the dashboard.
// URL: GET /api/v2/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) GetDashboardStatusByIDV2(ctx context.Context, id string) (*ResourceDashboardStatus, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourceDashboardStatus

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddToDashboardByIDV2 adds the software title configuration to the dashboard.
// URL: POST /api/v2/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) AddToDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

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

// RemoveFromDashboardByIDV2 removes the software title configuration from the dashboard.
// URL: DELETE /api/v2/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) RemoveFromDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// Data Retrieval
// -----------------------------------------------------------------------------

// GetDefinitionsByIDV2 returns paginated patch software title definitions.
// URL: GET /api/v2/patch-software-title-configurations/{id}/definitions
func (s *PatchSoftwareTitleConfigurations) GetDefinitionsByIDV2(ctx context.Context, id string, query map[string]string) (*DefinitionsResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/definitions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result DefinitionsResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourceDefinition
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	// Get totalCount from response
	var rawResp struct {
		TotalCount int `json:"totalCount"`
	}
	bodyBytes := resp.Bytes()
	if resp != nil && len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &rawResp)
		result.TotalCount = rawResp.TotalCount
	}
	if result.TotalCount == 0 && len(result.Results) > 0 {
	}

	return &result, resp, nil
}

// GetDependenciesByIDV2 returns paginated dependencies (smart groups) for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/dependencies
func (s *PatchSoftwareTitleConfigurations) GetDependenciesByIDV2(ctx context.Context, id string, query map[string]string) (*DependenciesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result DependenciesResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourceDependency
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	var rawResp struct {
		TotalCount int `json:"totalCount"`
	}
	bodyBytes := resp.Bytes()
	if resp != nil && len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &rawResp)
		result.TotalCount = rawResp.TotalCount
	}
	if result.TotalCount == 0 && len(result.Results) > 0 {
	}

	return &result, resp, nil
}

// ExportReportByIDV2 exports patch reporting data as CSV bytes.
// URL: GET /api/v2/patch-software-title-configurations/{id}/export-report
func (s *PatchSoftwareTitleConfigurations) ExportReportByIDV2(ctx context.Context, id string, query map[string]string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/export-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	headers := map[string]string{
		"Accept": constants.TextCSV,
	}

	resp, body, err := s.client.GetBytes(ctx, endpoint, query, headers)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

// GetExtensionAttributesByIDV2 returns extension attributes for the software title.
// URL: GET /api/v2/patch-software-title-configurations/{id}/extension-attributes
func (s *PatchSoftwareTitleConfigurations) GetExtensionAttributesByIDV2(ctx context.Context, id string) ([]ResourceExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/extension-attributes", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result []ResourceExtensionAttribute

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetPatchReportByIDV2 returns paginated patch report for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/patch-report
func (s *PatchSoftwareTitleConfigurations) GetPatchReportByIDV2(ctx context.Context, id string, query map[string]string) (*PatchReportResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result PatchReportResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourcePatchReportItem
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	var rawResp struct {
		TotalCount int `json:"totalCount"`
	}
	bodyBytes := resp.Bytes()
	if resp != nil && len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &rawResp)
		result.TotalCount = rawResp.TotalCount
	}
	if result.TotalCount == 0 && len(result.Results) > 0 {
	}

	return &result, resp, nil
}

// GetPatchSummaryByIDV2 returns the patch summary for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/patch-summary
func (s *PatchSoftwareTitleConfigurations) GetPatchSummaryByIDV2(ctx context.Context, id string) (*ResourcePatchSummary, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSummary

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// History Management
// -----------------------------------------------------------------------------

// GetHistoryByIDV2 returns paginated history for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/history
func (s *PatchSoftwareTitleConfigurations) GetHistoryByIDV2(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourceHistoryItem
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	var rawResp struct {
		TotalCount int `json:"totalCount"`
	}
	bodyBytes := resp.Bytes()
	if resp != nil && len(bodyBytes) > 0 {
		_ = json.Unmarshal(bodyBytes, &rawResp)
		result.TotalCount = rawResp.TotalCount
	}
	if result.TotalCount == 0 && len(result.Results) > 0 {
	}

	return &result, resp, nil
}

// AddHistoryNoteByIDV2 adds a history note to the configuration.
// URL: POST /api/v2/patch-software-title-configurations/{id}/history
func (s *PatchSoftwareTitleConfigurations) AddHistoryNoteByIDV2(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResponseAddHistoryNote

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// Versions
// -----------------------------------------------------------------------------

// GetPatchVersionsByIDV2 returns patch versions for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/patch-summary/versions
func (s *PatchSoftwareTitleConfigurations) GetPatchVersionsByIDV2(ctx context.Context, id string) ([]ResourcePatchVersion, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary/versions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result []ResourcePatchVersion

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
