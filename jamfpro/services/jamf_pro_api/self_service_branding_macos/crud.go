package self_service_branding_macos

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// SelfServiceBrandingMacOSServiceInterface defines the interface for self-service branding macOS operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
	SelfServiceBrandingMacOSServiceInterface interface {
		// List returns all self-service branding configurations for macOS.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
		List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified self-service branding configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos-id
		GetByID(ctx context.Context, id string) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error)

		// GetByName returns the specified self-service branding configuration by name.
		//
		// Performs a client-side search over the list of branding configurations.
		GetByName(ctx context.Context, name string) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error)

		// Create creates a new self-service branding configuration for macOS.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-branding-macos
		Create(ctx context.Context, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error)

		// UpdateByID updates the specified self-service branding configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-self-service-branding-macos-id
		UpdateByID(ctx context.Context, id string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error)

		// UpdateByName updates a self-service branding configuration by name.
		//
		// Performs GetByName then UpdateByID.
		UpdateByName(ctx context.Context, name string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error)

		// DeleteByID removes the specified self-service branding configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-self-service-branding-macos-id
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByName removes a self-service branding configuration by name.
		//
		// Performs GetByName then DeleteByID.
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the self-service branding macOS methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SelfServiceBrandingMacOSServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Branding macOS CRUD Operations
// -----------------------------------------------------------------------------

// List returns all self-service branding configurations for macOS.
// URL: GET /api/v1/self-service/branding/macos
func (s *Service) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceSelfServiceBrandingMacOS
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := EndpointSelfServiceBrandingMacOSV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByID returns the specified self-service branding configuration by ID.
// URL: GET /api/v1/self-service/branding/macos/{id}
func (s *Service) GetByID(ctx context.Context, id string) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("self-service branding configuration ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSelfServiceBrandingMacOSV1, id)

	var result ResourceSelfServiceBrandingMacOS

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified self-service branding configuration by name.
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error) {
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
func (s *Service) Create(ctx context.Context, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSelfServiceBrandingMacOS

	endpoint := EndpointSelfServiceBrandingMacOSV1

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

// UpdateByID updates the specified self-service branding configuration by ID.
// URL: PUT /api/v1/self-service/branding/macos/{id}
func (s *Service) UpdateByID(ctx context.Context, id string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSelfServiceBrandingMacOSV1, id)

	var result ResourceSelfServiceBrandingMacOS

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates a self-service branding configuration by name.
func (s *Service) UpdateByName(ctx context.Context, name string, request *ResourceSelfServiceBrandingMacOS) (*ResourceSelfServiceBrandingMacOS, *interfaces.Response, error) {
	target, resp, err := s.GetByName(ctx, name)
	if err != nil {
		return nil, resp, err
	}

	return s.UpdateByID(ctx, target.ID, request)
}

// DeleteByID removes the specified self-service branding configuration by ID.
// URL: DELETE /api/v1/self-service/branding/macos/{id}
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("self-service branding configuration ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSelfServiceBrandingMacOSV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes a self-service branding configuration by name.
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	target, resp, err := s.GetByName(ctx, name)
	if err != nil {
		return resp, err
	}

	return s.DeleteByID(ctx, target.ID)
}
