package patch_software_title_configurations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"resty.dev/v3"
)

type (
	// Service handles communication with the patch software title configurations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
	PatchSoftwareTitleConfigurations struct {
		client client.Client
	}
)

// NewService returns a new patch software title configurations Service backed by the provided HTTP client.
func NewPatchSoftwareTitleConfigurations(client client.Client) *PatchSoftwareTitleConfigurations {
	return &PatchSoftwareTitleConfigurations{client: client}
}

// deprecatedV2Replacement is the migration hint logged for the deprecated V2
// patch-software-title-configuration methods.
const deprecatedV2Replacement = "use the v3 patch-software-title-configurations endpoints (\u2026V3 methods)"

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Software Title Configurations CRUD Operations (v2)
// -----------------------------------------------------------------------------

// ListV2 returns all patch software title configurations.
// URL: GET /api/v2/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
//
// Deprecated: deprecated in Jamf Pro 11.30; use ListV3.
func (s *PatchSoftwareTitleConfigurations) ListV2(ctx context.Context) (*ListResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.ListV2", "11.30", deprecatedV2Replacement)

	var result ListResponse

	endpoint := constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV2 returns the patch software title configuration by ID.
// URL: GET /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetByIDV2(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV2 returns the patch software title configuration by display name.
// This is a convenience method that calls ListV2 and filters by DisplayName.
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetByNameV3.
func (s *PatchSoftwareTitleConfigurations) GetByNameV2(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetByNameV2", "11.30", deprecatedV2Replacement)

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
//
// Deprecated: deprecated in Jamf Pro 11.30; use CreateV3.
func (s *PatchSoftwareTitleConfigurations) CreateV2(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.CreateV2", "11.30", deprecatedV2Replacement)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(config).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 updates the patch software title configuration by ID.
// URL: PATCH /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/patch_v2-patch-software-title-configurations-id
//
// Deprecated: deprecated in Jamf Pro 11.30; use UpdateByIDV3.
func (s *PatchSoftwareTitleConfigurations) UpdateByIDV2(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.UpdateByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if config == nil {
		return nil, nil, fmt.Errorf("config is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(config).
		SetResult(&result).
		Patch(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByNameV2 updates the patch software title configuration by display name.
//
// Deprecated: deprecated in Jamf Pro 11.30; use UpdateByNameV3.
func (s *PatchSoftwareTitleConfigurations) UpdateByNameV2(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.UpdateByNameV2", "11.30", deprecatedV2Replacement)

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
//
// Deprecated: deprecated in Jamf Pro 11.30; use DeleteByIDV3.
func (s *PatchSoftwareTitleConfigurations) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.DeleteByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV2 deletes the patch software title configuration by display name.
//
// Deprecated: deprecated in Jamf Pro 11.30; use DeleteByNameV3.
func (s *PatchSoftwareTitleConfigurations) DeleteByNameV2(ctx context.Context, name string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.DeleteByNameV2", "11.30", deprecatedV2Replacement)

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
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetDashboardStatusByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetDashboardStatusByIDV2(ctx context.Context, id string) (*ResourceDashboardStatus, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetDashboardStatusByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourceDashboardStatus

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddToDashboardByIDV2 adds the software title configuration to the dashboard.
// URL: POST /api/v2/patch-software-title-configurations/{id}/dashboard
//
// Deprecated: deprecated in Jamf Pro 11.30; use AddToDashboardByIDV3.
func (s *PatchSoftwareTitleConfigurations) AddToDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.AddToDashboardByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveFromDashboardByIDV2 removes the software title configuration from the dashboard.
// URL: DELETE /api/v2/patch-software-title-configurations/{id}/dashboard
//
// Deprecated: deprecated in Jamf Pro 11.30; use RemoveFromDashboardByIDV3.
func (s *PatchSoftwareTitleConfigurations) RemoveFromDashboardByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.RemoveFromDashboardByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetDefinitionsByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetDefinitionsByIDV2(ctx context.Context, id string, query map[string]string) (*DefinitionsResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetDefinitionsByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/definitions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result DefinitionsResponse

	mergePage := func(pageData []byte) error {
		var page []ResourceDefinition
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
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

// GetDependenciesByIDV2 returns paginated dependencies (smart groups) for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/dependencies
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetDependenciesByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetDependenciesByIDV2(ctx context.Context, id string, query map[string]string) (*DependenciesResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetDependenciesByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result DependenciesResponse

	mergePage := func(pageData []byte) error {
		var page []ResourceDependency
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use ExportReportByIDV3.
func (s *PatchSoftwareTitleConfigurations) ExportReportByIDV2(ctx context.Context, id string, query map[string]string) ([]byte, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.ExportReportByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/export-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.TextCSV).
		SetQueryParams(query).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

// GetExtensionAttributesByIDV2 returns extension attributes for the software title.
// URL: GET /api/v2/patch-software-title-configurations/{id}/extension-attributes
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetExtensionAttributesByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetExtensionAttributesByIDV2(ctx context.Context, id string) ([]ResourceExtensionAttribute, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetExtensionAttributesByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/extension-attributes", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result []ResourceExtensionAttribute

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetPatchReportByIDV2 returns paginated patch report for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/patch-report
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetPatchReportByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetPatchReportByIDV2(ctx context.Context, id string, query map[string]string) (*PatchReportResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetPatchReportByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result PatchReportResponse

	mergePage := func(pageData []byte) error {
		var page []ResourcePatchReportItem
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetPatchSummaryByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetPatchSummaryByIDV2(ctx context.Context, id string) (*ResourcePatchSummary, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetPatchSummaryByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSummary

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetHistoryByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetHistoryByIDV2(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetHistoryByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var page []ResourceHistoryItem
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, page...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
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
//
// Deprecated: deprecated in Jamf Pro 11.30; use AddHistoryNoteByIDV3.
func (s *PatchSoftwareTitleConfigurations) AddHistoryNoteByIDV2(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.AddHistoryNoteByIDV2", "11.30", deprecatedV2Replacement)

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

// -----------------------------------------------------------------------------
// Versions
// -----------------------------------------------------------------------------

// GetPatchVersionsByIDV2 returns patch versions for the configuration.
// URL: GET /api/v2/patch-software-title-configurations/{id}/patch-summary/versions
//
// Deprecated: deprecated in Jamf Pro 11.30; use GetPatchVersionsByIDV3.
func (s *PatchSoftwareTitleConfigurations) GetPatchVersionsByIDV2(ctx context.Context, id string) ([]ResourcePatchVersion, *resty.Response, error) {
	apilifecycle.DeprecationWarning(s.client.GetLogger(), "jamf_pro_api/patch_software_title_configurations.PatchSoftwareTitleConfigurations.GetPatchVersionsByIDV2", "11.30", deprecatedV2Replacement)

	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary/versions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV2, id)

	var result []ResourcePatchVersion

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
