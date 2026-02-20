package cloud_azure

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CloudAzureServiceInterface defines the interface for Cloud Azure (Azure Cloud IDP) operations.
	// Uses v1 API for all operations. Supports Azure Active Directory integration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
	CloudAzureServiceInterface interface {
		// GetDefaultServerConfigurationV1 returns the default server configuration for Azure Cloud IDP (Get Default Server Configuration).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-defaults-server-configuration
		GetDefaultServerConfigurationV1(ctx context.Context) (*CloudAzureServer, *interfaces.Response, error)

		// GetByIDV1 returns the Azure Cloud IDP configuration by ID (Get Cloud Identity Provider by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-id
		GetByIDV1(ctx context.Context, id string) (*ResourceCloudAzure, *interfaces.Response, error)

		// GetByNameV1 returns the Azure Cloud IDP configuration by display name (searches all providers).
		GetByNameV1(ctx context.Context, name string) (*ResourceCloudAzure, *interfaces.Response, error)

		// CreateV1 creates a new Azure Cloud IDP configuration (Create Cloud Identity Provider).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
		CreateV1(ctx context.Context, request *ResourceCloudAzure) (*ResponseCloudAzureCreated, *interfaces.Response, error)

		// UpdateByIDV1 updates the Azure Cloud IDP configuration by ID (Update Cloud Identity Provider by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cloud-azure-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceCloudAzure) (*ResourceCloudAzure, *interfaces.Response, error)

		// UpdateByNameV1 updates the Azure Cloud IDP configuration by display name.
		UpdateByNameV1(ctx context.Context, name string, request *ResourceCloudAzure) (*ResourceCloudAzure, *interfaces.Response, error)

		// DeleteByIDV1 deletes the Azure Cloud IDP configuration by ID (Delete Cloud Identity Provider by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-cloud-azure-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteByNameV1 deletes the Azure Cloud IDP configuration by display name.
		DeleteByNameV1(ctx context.Context, name string) (*interfaces.Response, error)

		// GetDefaultMappingsV1 returns the default field mappings for Azure Cloud IDP (Get Default Mappings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-defaults-mappings
		GetDefaultMappingsV1(ctx context.Context) (*CloudAzureServerMappings, *interfaces.Response, error)
	}

	// Service handles communication with the Cloud Azure-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CloudAzureServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetDefaultServerConfigurationV1 returns the default server configuration for Azure Cloud IDP.
// URL: GET /api/v1/cloud-azure/defaults/server-configuration
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-defaults-server-configuration
func (s *Service) GetDefaultServerConfigurationV1(ctx context.Context) (*CloudAzureServer, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/defaults/server-configuration", EndpointCloudAzureV1)

	var result CloudAzureServer

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

// GetByIDV1 returns the Azure Cloud IDP configuration by ID.
// URL: GET /api/v1/cloud-azure/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceCloudAzure, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudAzureV1, id)

	var result ResourceCloudAzure

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

// GetByNameV1 returns the Azure Cloud IDP configuration by display name.
// URL: GET /api/v1/cloud-azure (searches all providers)
func (s *Service) GetByNameV1(ctx context.Context, name string) (*ResourceCloudAzure, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	var providers []ResourceCloudAzure

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointCloudAzureV1, nil, headers, &providers)
	if err != nil {
		return nil, resp, err
	}

	for _, provider := range providers {
		if provider.CloudIdPCommon.DisplayName == name {
			return s.GetByIDV1(ctx, provider.Server.ID)
		}
	}

	return nil, resp, fmt.Errorf("cloud azure provider with name %q not found", name)
}

// CreateV1 creates a new Azure Cloud IDP configuration.
// URL: POST /api/v1/cloud-azure
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-azure
func (s *Service) CreateV1(ctx context.Context, request *ResourceCloudAzure) (*ResponseCloudAzureCreated, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResponseCloudAzureCreated

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointCloudAzureV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the Azure Cloud IDP configuration by ID.
// URL: PUT /api/v1/cloud-azure/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-cloud-azure-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *ResourceCloudAzure) (*ResourceCloudAzure, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudAzureV1, id)

	var result ResourceCloudAzure

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

// UpdateByNameV1 updates the Azure Cloud IDP configuration by display name.
func (s *Service) UpdateByNameV1(ctx context.Context, name string, request *ResourceCloudAzure) (*ResourceCloudAzure, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name is required")
	}

	existing, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return nil, resp, err
	}

	return s.UpdateByIDV1(ctx, existing.Server.ID, request)
}

// DeleteByIDV1 deletes the Azure Cloud IDP configuration by ID.
// URL: DELETE /api/v1/cloud-azure/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-cloud-azure-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudAzureV1, id)

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

// DeleteByNameV1 deletes the Azure Cloud IDP configuration by display name.
func (s *Service) DeleteByNameV1(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	existing, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return resp, err
	}

	return s.DeleteByIDV1(ctx, existing.Server.ID)
}

// GetDefaultMappingsV1 returns the default field mappings for Azure Cloud IDP.
// URL: GET /api/v1/cloud-azure/defaults/mappings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-azure-defaults-mappings
func (s *Service) GetDefaultMappingsV1(ctx context.Context) (*CloudAzureServerMappings, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/defaults/mappings", EndpointCloudAzureV1)

	var result CloudAzureServerMappings

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
