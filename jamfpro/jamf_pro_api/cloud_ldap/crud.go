package cloud_ldap

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Cloud LDAP-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-cloud-ldaps
	CloudLdap struct {
		client transport.HTTPClient
	}
)

func NewCloudLdap(client transport.HTTPClient) *CloudLdap {
	return &CloudLdap{client: client}
}

// GetDefaultMappingsV2 returns the default field mappings for the specified provider.
// URL: GET /api/v2/cloud-ldaps/defaults/{providerName}/mappings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-defaults-provider-mappings
func (s *CloudLdap) GetDefaultMappingsV2(ctx context.Context, providerName string) (*ResponseDefaultMappings, *resty.Response, error) {
	if providerName == "" {
		return nil, nil, fmt.Errorf("providerName is required")
	}

	endpoint := fmt.Sprintf("%s/defaults/%s/mappings", constants.EndpointJamfProCloudLdapV2, providerName)

	var result ResponseDefaultMappings

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) GetDefaultServerConfigurationV2(ctx context.Context, providerName string) (*ResponseDefaultServerConfiguration, *resty.Response, error) {
	if providerName == "" {
		return nil, nil, fmt.Errorf("providerName is required")
	}

	endpoint := fmt.Sprintf("%s/defaults/%s/server-configuration", constants.EndpointJamfProCloudLdapV2, providerName)

	var result ResponseDefaultServerConfiguration

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) CreateV2(ctx context.Context, request *ResourceCloudLdap) (*ResponseCloudLdapCreated, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProCloudLdapV2

	var result ResponseCloudLdapCreated

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

// GetByIDV2 returns the Cloud LDAP configuration by ID.
// URL: GET /api/v2/cloud-ldaps/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-cloud-ldaps-id
func (s *CloudLdap) GetByIDV2(ctx context.Context, id string) (*ResourceCloudLdap, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCloudLdapV2, id)

	var result ResourceCloudLdap

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) UpdateByIDV2(ctx context.Context, id string, request *ResourceCloudLdap) (*ResourceCloudLdap, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCloudLdapV2, id)

	var result ResourceCloudLdap

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *CloudLdap) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCloudLdapV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) GetBindConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/bind", constants.EndpointJamfProCloudLdapV2, id)

	var result ConnectionPoolStats

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) GetSearchConnectionPoolStatsByIDV2(ctx context.Context, id string) (*ConnectionPoolStats, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/search", constants.EndpointJamfProCloudLdapV2, id)

	var result ConnectionPoolStats

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) TestConnectionByIDV2(ctx context.Context, id string) (*ConnectionStatusResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection/status", constants.EndpointJamfProCloudLdapV2, id)

	var result ConnectionStatusResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) GetMappingsByIDV2(ctx context.Context, id string) (*CloudLdapMappings, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/mappings", constants.EndpointJamfProCloudLdapV2, id)

	var result CloudLdapMappings

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *CloudLdap) UpdateMappingsByIDV2(ctx context.Context, id string, request *CloudLdapMappings) (*CloudLdapMappings, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/mappings", constants.EndpointJamfProCloudLdapV2, id)

	var result CloudLdapMappings

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
