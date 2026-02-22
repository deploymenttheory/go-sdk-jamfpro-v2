package patch_software_title_configurations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
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
		ListV2(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByIDV2 returns the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
		GetByIDV2(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error)

		// GetByNameV2 returns the patch software title configuration by display name.
		//
		// This is a convenience method that calls ListV2 and filters by DisplayName.
		GetByNameV2(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error)

		// CreateV2 creates a new patch software title configuration.
		// Returns CreateResponse (id, href).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-software-title-configurations
		CreateV2(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV2 updates the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-patch-software-title-configurations-id
		UpdateByIDV2(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error)

		// UpdateByNameV2 updates the patch software title configuration by display name.
		UpdateByNameV2(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error)

		// DeleteByIDV2 deletes the patch software title configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-patch-software-title-configurations-id
		DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV2 deletes the patch software title configuration by display name.
		DeleteByNameV2(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the patch software title configurations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ PatchSoftwareTitleConfigurationsServiceInterface = (*Service)(nil)

// NewService returns a new patch software title configurations Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Patch Software Title Configurations CRUD Operations (v2)
// -----------------------------------------------------------------------------

// ListV2 returns all patch software title configurations.
// URL: GET /api/v2/patch-software-title-configurations
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations
func (s *Service) ListV2(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var page []ResourcePatchSoftwareTitleConfiguration
		if err := json.Unmarshal(pageData, &page); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result = append(result, page...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, EndpointPatchSoftwareTitleConfigurationsV2, nil, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV2 returns the patch software title configuration by ID.
// URL: GET /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v2-patch-software-title-configurations-id
func (s *Service) GetByIDV2(ctx context.Context, id string) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

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

// GetByNameV2 returns the patch software title configuration by display name.
// This is a convenience method that calls ListV2 and filters by DisplayName.
func (s *Service) GetByNameV2(ctx context.Context, name string) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error) {
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
func (s *Service) CreateV2(ctx context.Context, config *ResourcePatchSoftwareTitleConfiguration) (*CreateResponse, *interfaces.Response, error) {
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

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointPatchSoftwareTitleConfigurationsV2, config, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 updates the patch software title configuration by ID.
// URL: PATCH /api/v2/patch-software-title-configurations/{id}
// https://developer.jamf.com/jamf-pro/reference/patch_v2-patch-software-title-configurations-id
func (s *Service) UpdateByIDV2(ctx context.Context, id string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if config == nil {
		return nil, nil, fmt.Errorf("config is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPatchSoftwareTitleConfigurationsV2, id)

	var result ResourcePatchSoftwareTitleConfiguration

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, config, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByNameV2 updates the patch software title configuration by display name.
func (s *Service) UpdateByNameV2(ctx context.Context, name string, config *ResourcePatchSoftwareTitleConfiguration) (*ResourcePatchSoftwareTitleConfiguration, *interfaces.Response, error) {
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
func (s *Service) DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointPatchSoftwareTitleConfigurationsV2, id)

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

// DeleteByNameV2 deletes the patch software title configuration by display name.
func (s *Service) DeleteByNameV2(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	target, resp, err := s.GetByNameV2(ctx, name)
	if err != nil {
		return resp, fmt.Errorf("failed to get patch software title configuration by name: %w", err)
	}

	return s.DeleteByIDV2(ctx, target.ID)
}
