package cloud_ldap

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CloudLdapServiceInterface defines the interface for Cloud LDAP operations.
	// Uses v2 API for all operations. Supports Google and Azure LDAP integration.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
	CloudLdapServiceInterface interface {
		// GetDefaultMappingsV2 returns the default field mappings for the specified provider (Get Default Mappings).
		//
		// providerName should be "GOOGLE" or "AZURE".
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-defaults-provider-mappings
		GetDefaultMappingsV2(ctx context.Context, providerName string) (*ResponseDefaultMappings, *interfaces.Response, error)

		// GetDefaultServerConfigurationV2 returns the default server configuration for the specified provider (Get Default Server Configuration).
		//
		// providerName should be "GOOGLE" or "AZURE".
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-defaults-provider-server-configuration
		GetDefaultServerConfigurationV2(ctx context.Context, providerName string) (*ResponseDefaultServerConfiguration, *interfaces.Response, error)

		// CreateV2 creates a new Cloud LDAP configuration (Create Cloud LDAP).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
		CreateV2(ctx context.Context, request *ResourceCloudLdap) (*ResponseCloudLdapCreated, *interfaces.Response, error)

		// GetByIDV2 returns the Cloud LDAP configuration by ID (Get Cloud LDAP by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id
		GetByIDV2(ctx context.Context, id string) (*ResourceCloudLdap, *interfaces.Response, error)

		// UpdateByIDV2 updates the Cloud LDAP configuration by ID (Update Cloud LDAP by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-cloud-ldaps-id
		UpdateByIDV2(ctx context.Context, id string, request *ResourceCloudLdap) (*ResourceCloudLdap, *interfaces.Response, error)

		// DeleteByIDV2 deletes the Cloud LDAP configuration by ID (Delete Cloud LDAP by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-cloud-ldaps-id
		DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error)

		// GetBindConnectionPoolStatsByIDV2 returns bind connection pool statistics (Get Bind Connection Pool Statistics).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-bind
		GetBindConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *interfaces.Response, error)

		// GetSearchConnectionPoolStatsByIDV2 returns search connection pool statistics (Get Search Connection Pool Statistics).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-search
		GetSearchConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *interfaces.Response, error)

		// TestConnectionByIDV2 tests the communication with the specified cloud connection (Test Connection).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-status
		TestConnectionByIDV2(ctx context.Context, id string) (*ConnectionStatusResponse, *interfaces.Response, error)

		// GetMappingsByIDV2 returns the mappings configuration for the Cloud LDAP by ID (Get Mappings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-mappings
		GetMappingsByIDV2(ctx context.Context, id string) (*CloudLdapMappings, *interfaces.Response, error)

		// UpdateMappingsByIDV2 updates the mappings configuration for the Cloud LDAP by ID (Update Mappings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-cloud-ldaps-id-mappings
		UpdateMappingsByIDV2(ctx context.Context, id string, request *CloudLdapMappings) (*CloudLdapMappings, *interfaces.Response, error)
	}

	// Service handles communication with the Cloud LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CloudLdapServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetDefaultMappingsV2 returns the default field mappings for the specified provider.
// URL: GET /api/v2/cloud-ldaps/defaults/{providerName}/mappings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-defaults-provider-mappings
func (s *Service) GetDefaultMappingsV2(ctx context.Context, providerName string) (*ResponseDefaultMappings, *interfaces.Response, error) {
	if providerName == "" {
		return nil, nil, fmt.Errorf("providerName is required")
	}

	endpoint := fmt.Sprintf("%s/defaults/%s/mappings", EndpointCloudLdapV2, providerName)

	var result ResponseDefaultMappings

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

// GetDefaultServerConfigurationV2 returns the default server configuration for the specified provider.
// URL: GET /api/v2/cloud-ldaps/defaults/{providerName}/server-configuration
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-defaults-provider-server-configuration
func (s *Service) GetDefaultServerConfigurationV2(ctx context.Context, providerName string) (*ResponseDefaultServerConfiguration, *interfaces.Response, error) {
	if providerName == "" {
		return nil, nil, fmt.Errorf("providerName is required")
	}

	endpoint := fmt.Sprintf("%s/defaults/%s/server-configuration", EndpointCloudLdapV2, providerName)

	var result ResponseDefaultServerConfiguration

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

// CreateV2 creates a new Cloud LDAP configuration.
// URL: POST /api/v2/cloud-ldaps
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
func (s *Service) CreateV2(ctx context.Context, request *ResourceCloudLdap) (*ResponseCloudLdapCreated, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResponseCloudLdapCreated

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, EndpointCloudLdapV2, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV2 returns the Cloud LDAP configuration by ID.
// URL: GET /api/v2/cloud-ldaps/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id
func (s *Service) GetByIDV2(ctx context.Context, id string) (*ResourceCloudLdap, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudLdapV2, id)

	var result ResourceCloudLdap

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

// UpdateByIDV2 updates the Cloud LDAP configuration by ID.
// URL: PUT /api/v2/cloud-ldaps/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-cloud-ldaps-id
func (s *Service) UpdateByIDV2(ctx context.Context, id string, request *ResourceCloudLdap) (*ResourceCloudLdap, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudLdapV2, id)

	var result ResourceCloudLdap

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

// DeleteByIDV2 deletes the Cloud LDAP configuration by ID.
// URL: DELETE /api/v2/cloud-ldaps/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-cloud-ldaps-id
func (s *Service) DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointCloudLdapV2, id)

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

// GetBindConnectionPoolStatsByIDV2 returns bind connection pool statistics.
// URL: GET /api/v2/cloud-ldaps/{id}/connection/bind
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-bind
func (s *Service) GetBindConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/bind", EndpointCloudLdapV2, id)

	var result ConnectionPoolStats

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

// GetSearchConnectionPoolStatsByIDV2 returns search connection pool statistics.
// URL: GET /api/v2/cloud-ldaps/{id}/connection/search
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-search
func (s *Service) GetSearchConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/search", EndpointCloudLdapV2, id)

	var result ConnectionPoolStats

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

// TestConnectionByIDV2 tests the communication with the specified cloud connection.
// URL: GET /api/v2/cloud-ldaps/{id}/connection/status
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-connection-status
func (s *Service) TestConnectionByIDV2(ctx context.Context, id string) (*ConnectionStatusResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/status", EndpointCloudLdapV2, id)

	var result ConnectionStatusResponse

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

// GetMappingsByIDV2 returns the mappings configuration for the Cloud LDAP by ID.
// URL: GET /api/v2/cloud-ldaps/{id}/mappings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id-mappings
func (s *Service) GetMappingsByIDV2(ctx context.Context, id string) (*CloudLdapMappings, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/mappings", EndpointCloudLdapV2, id)

	var result CloudLdapMappings

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

// UpdateMappingsByIDV2 updates the mappings configuration for the Cloud LDAP by ID.
// URL: PUT /api/v2/cloud-ldaps/{id}/mappings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-cloud-ldaps-id-mappings
func (s *Service) UpdateMappingsByIDV2(ctx context.Context, id string, request *CloudLdapMappings) (*CloudLdapMappings, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/mappings", EndpointCloudLdapV2, id)

	var result CloudLdapMappings

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
