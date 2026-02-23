package disk_encryption_configurations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DiskEncryptionConfigurationsServiceInterface defines the interface for Classic API disk encryption configuration operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/diskencryptionconfigurations
	DiskEncryptionConfigurationsServiceInterface interface {
		// List returns all disk encryption configurations.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurations
		List(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetByID returns the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyid
		GetByID(ctx context.Context, id int) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error)

		// GetByName returns the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyname
		GetByName(ctx context.Context, name string) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error)

		// Create creates a new disk encryption configuration.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/creatediskencryptionconfigurationbyid
		Create(ctx context.Context, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByID updates the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyid
		UpdateByID(ctx context.Context, id int, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateByName updates the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyname
		UpdateByName(ctx context.Context, name string, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteByID removes the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyid
		DeleteByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteByName removes the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyname
		DeleteByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the disk encryption configuration-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/diskencryptionconfigurations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DiskEncryptionConfigurationsServiceInterface = (*Service)(nil)

// NewService returns a new disk encryption configurations Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Disk Encryption Configurations CRUD Operations
// -----------------------------------------------------------------------------

// List returns all disk encryption configurations.
// URL: GET /JSSResource/diskencryptionconfigurations
// https://developer.jamf.com/jamf-pro/reference/findalldiskencryptionconfigurations
func (s *Service) List(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointClassicDiskEncryptionConfigurations

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) GetByID(ctx context.Context, id int) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	var result ResourceDiskEncryptionConfiguration

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) GetByName(ctx context.Context, name string) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	var result ResourceDiskEncryptionConfiguration

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) Create(ctx context.Context, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicDiskEncryptionConfigurations)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) UpdateByID(ctx context.Context, id int, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) UpdateByName(ctx context.Context, name string, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	var result CreateUpdateResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) DeleteByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
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
func (s *Service) DeleteByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
