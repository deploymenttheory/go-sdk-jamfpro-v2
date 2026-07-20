package patch_software_title_configurations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Software Title Configurations CRUD Operations (v3)
// -----------------------------------------------------------------------------

// ListV3 returns all patch software title configurations.
// URL: GET /api/v3/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/get_v3-patch-software-title-configurations
func (s *PatchSoftwareTitleConfigurations) ListV3(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3

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

// GetByIDV3 returns the patch software title configuration by ID.
// URL: GET /api/v3/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v3-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) GetByIDV3(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// GetByNameV3 returns the patch software title configuration by display name.
// This is a convenience method that calls ListV3 and filters by DisplayName.
func (s *PatchSoftwareTitleConfigurations) GetByNameV3(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	list, resp, err := s.ListV3(ctx)
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

// CreateV3 creates a new patch software title configuration.
// URL: POST /api/v3/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/post_v3-patch-software-title-configurations
func (s *PatchSoftwareTitleConfigurations) CreateV3(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *resty.Response, error) {
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

	endpoint := constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3

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

// UpdateByIDV3 updates the patch software title configuration by ID.
// URL: PATCH /api/v3/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/patch_v3-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) UpdateByIDV3(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if config == nil {
		return nil, nil, fmt.Errorf("config is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// UpdateByNameV3 updates the patch software title configuration by display name.
func (s *PatchSoftwareTitleConfigurations) UpdateByNameV3(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get patch software title configuration by name: %w", err)
	}

	return s.UpdateByIDV3(ctx, target.ID, config)
}

// DeleteByIDV3 deletes the patch software title configuration by ID.
// URL: DELETE /api/v3/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-patch-software-title-configurations-id
func (s *PatchSoftwareTitleConfigurations) DeleteByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV3 deletes the patch software title configuration by display name.
func (s *PatchSoftwareTitleConfigurations) DeleteByNameV3(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV3(ctx, name)
	if err != nil {
		return resp, fmt.Errorf("failed to get patch software title configuration by name: %w", err)
	}

	return s.DeleteByIDV3(ctx, target.ID)
}

// -----------------------------------------------------------------------------
// Dashboard Management
// -----------------------------------------------------------------------------

// GetDashboardStatusByIDV3 returns whether the software title configuration is on the dashboard.
// URL: GET /api/v3/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) GetDashboardStatusByIDV3(ctx context.Context, id string) (*ResourceDashboardStatus, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// AddToDashboardByIDV3 adds the software title configuration to the dashboard.
// URL: POST /api/v3/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) AddToDashboardByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveFromDashboardByIDV3 removes the software title configuration from the dashboard.
// URL: DELETE /api/v3/patch-software-title-configurations/{id}/dashboard
func (s *PatchSoftwareTitleConfigurations) RemoveFromDashboardByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dashboard", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// GetDefinitionsByIDV3 returns paginated patch software title definitions.
// URL: GET /api/v3/patch-software-title-configurations/{id}/definitions
func (s *PatchSoftwareTitleConfigurations) GetDefinitionsByIDV3(ctx context.Context, id string, query map[string]string) (*DefinitionsResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/definitions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

	return &result, resp, nil
}

// GetDependenciesByIDV3 returns paginated dependencies (smart groups) for the configuration.
// URL: GET /api/v3/patch-software-title-configurations/{id}/dependencies
func (s *PatchSoftwareTitleConfigurations) GetDependenciesByIDV3(ctx context.Context, id string, query map[string]string) (*DependenciesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

	return &result, resp, nil
}

// ExportReportByIDV3 exports patch reporting data as CSV bytes.
// URL: GET /api/v3/patch-software-title-configurations/{id}/export-report
//
// Note: Jamf Pro 11.30 renamed the lastContactTime column to lastCheckIn. Callers
// passing an explicit columns-to-export query param must request lastCheckIn; the
// default column set was updated to match.
func (s *PatchSoftwareTitleConfigurations) ExportReportByIDV3(ctx context.Context, id string, query map[string]string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/export-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.TextCSV).
		SetQueryParams(query).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return body, resp, nil
}

// GetExtensionAttributesByIDV3 returns extension attributes for the software title.
// URL: GET /api/v3/patch-software-title-configurations/{id}/extension-attributes
func (s *PatchSoftwareTitleConfigurations) GetExtensionAttributesByIDV3(ctx context.Context, id string) ([]ResourceExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/extension-attributes", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// GetPatchReportByIDV3 returns paginated patch report for the configuration.
// URL: GET /api/v3/patch-software-title-configurations/{id}/patch-report
func (s *PatchSoftwareTitleConfigurations) GetPatchReportByIDV3(ctx context.Context, id string, query map[string]string) (*PatchReportResponseV3, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-report", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

	var result PatchReportResponseV3

	mergePage := func(pageData []byte) error {
		var page []ResourcePatchReportItemV3
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

	return &result, resp, nil
}

// GetPatchSummaryByIDV3 returns the patch summary for the configuration.
// URL: GET /api/v3/patch-software-title-configurations/{id}/patch-summary
func (s *PatchSoftwareTitleConfigurations) GetPatchSummaryByIDV3(ctx context.Context, id string) (*ResourcePatchSummary, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// GetHistoryByIDV3 returns paginated history for the configuration.
// URL: GET /api/v3/patch-software-title-configurations/{id}/history
func (s *PatchSoftwareTitleConfigurations) GetHistoryByIDV3(ctx context.Context, id string, query map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

	return &result, resp, nil
}

// AddHistoryNoteByIDV3 adds a history note to the configuration.
// URL: POST /api/v3/patch-software-title-configurations/{id}/history
func (s *PatchSoftwareTitleConfigurations) AddHistoryNoteByIDV3(ctx context.Context, id string, request *RequestAddHistoryNote) (*ResponseAddHistoryNote, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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

// GetPatchVersionsByIDV3 returns patch versions for the configuration.
// URL: GET /api/v3/patch-software-title-configurations/{id}/patch-summary/versions
func (s *PatchSoftwareTitleConfigurations) GetPatchVersionsByIDV3(ctx context.Context, id string) ([]ResourcePatchVersion, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/patch-summary/versions", constants.EndpointJamfProPatchSoftwareTitleConfigurationsV3, id)

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
