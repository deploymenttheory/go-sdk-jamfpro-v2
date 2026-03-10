package self_service_branding_macos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the self-service branding macOS methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
	SelfServiceBrandingMacos struct {
		client client.Client
	}
)

func NewSelfServiceBrandingMacos(client client.Client) *SelfServiceBrandingMacos {
	return &SelfServiceBrandingMacos{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Branding macOS CRUD Operations
// -----------------------------------------------------------------------------

// List returns all self-service branding configurations for macOS.
// URL: GET /api/v1/self-service/branding/macos
func (s *SelfServiceBrandingMacos) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceSelfServiceBrandingMacOS
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProSelfServiceBrandingMacOSV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByID returns the specified self-service branding configuration by ID.
// URL: GET /api/v1/self-service/branding/macos/{id}
func (s *SelfServiceBrandingMacos) GetByID(ctx context.Context, id string) (*ResourceSelfServiceBrandingMacOS, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("self-service branding configuration ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMacOSV1, id)

	var result ResourceSelfServiceBrandingMacOS

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified self-service branding configuration by name.
func (s *SelfServiceBrandingMacos) GetByName(ctx context.Context, name string) (*ResourceSelfServiceBrandingMacOS, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("self-service branding configuration name is required")
	}

	list, resp, err := s.List(ctx, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list self-service branding configurations: %w", err)
	}

	for _, item := range list.Results {
		if item.BrandingName == name {
			return &item, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("self-service branding configuration with name %q not found", name)
}

// Create creates a new self-service branding configuration for macOS.
// URL: POST /api/v1/self-service/branding/macos
func (s *SelfServiceBrandingMacos) Create(ctx context.Context, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSelfServiceBrandingMacOS

	endpoint := constants.EndpointJamfProSelfServiceBrandingMacOSV1

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

// UpdateByID updates the specified self-service branding configuration by ID.
// URL: PUT /api/v1/self-service/branding/macos/{id}
func (s *SelfServiceBrandingMacos) UpdateByID(ctx context.Context, id string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMacOSV1, id)

	var result ResourceSelfServiceBrandingMacOS

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates a self-service branding configuration by name.
func (s *SelfServiceBrandingMacos) UpdateByName(ctx context.Context, name string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *resty.Response, error) {
	target, resp, err := s.GetByName(ctx, name)
	if err != nil {
		return nil, resp, err
	}

	return s.UpdateByID(ctx, target.ID, request)
}

// DeleteByID removes the specified self-service branding configuration by ID.
// URL: DELETE /api/v1/self-service/branding/macos/{id}
func (s *SelfServiceBrandingMacos) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("self-service branding configuration ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMacOSV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes a self-service branding configuration by name.
func (s *SelfServiceBrandingMacos) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	target, resp, err := s.GetByName(ctx, name)
	if err != nil {
		return resp, err
	}

	return s.DeleteByID(ctx, target.ID)
}
