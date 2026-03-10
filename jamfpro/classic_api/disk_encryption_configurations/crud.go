package disk_encryption_configurations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the disk encryption configuration-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/diskencryptionconfigurations
	DiskEncryptionConfigurations struct {
		client client.Client
	}
)

// NewService returns a new disk encryption configurations Service backed by the provided HTTP client.
func NewDiskEncryptionConfigurations(client client.Client) *DiskEncryptionConfigurations {
	return &DiskEncryptionConfigurations{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Disk Encryption Configurations CRUD Operations
// -----------------------------------------------------------------------------

// List returns all disk encryption configurations.
// URL: GET /JSSResource/diskencryptionconfigurations
// https://developer.jamf.com/jamf-pro/reference/findalldiskencryptionconfigurations
func (s *DiskEncryptionConfigurations) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicDiskEncryptionConfigurations

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified disk encryption configuration by ID.
// URL: GET /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyid
func (s *DiskEncryptionConfigurations) GetByID(ctx context.Context, id int) (*ResourceDiskEncryptionConfiguration, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDiskEncryptionConfigurations, id)

	var result ResourceDiskEncryptionConfiguration

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified disk encryption configuration by name.
// URL: GET /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyname
func (s *DiskEncryptionConfigurations) GetByName(ctx context.Context, name string) (*ResourceDiskEncryptionConfiguration, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDiskEncryptionConfigurations, name)

	var result ResourceDiskEncryptionConfiguration

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new disk encryption configuration.
// URL: POST /JSSResource/diskencryptionconfigurations/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/creatediskencryptionconfigurationbyid
func (s *DiskEncryptionConfigurations) Create(ctx context.Context, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicDiskEncryptionConfigurations)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified disk encryption configuration by ID.
// URL: PUT /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyid
func (s *DiskEncryptionConfigurations) UpdateByID(ctx context.Context, id int, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDiskEncryptionConfigurations, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified disk encryption configuration by name.
// URL: PUT /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyname
func (s *DiskEncryptionConfigurations) UpdateByName(ctx context.Context, name string, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDiskEncryptionConfigurations, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified disk encryption configuration by ID.
// URL: DELETE /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyid
func (s *DiskEncryptionConfigurations) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicDiskEncryptionConfigurations, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified disk encryption configuration by name.
// URL: DELETE /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyname
func (s *DiskEncryptionConfigurations) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicDiskEncryptionConfigurations, name)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
