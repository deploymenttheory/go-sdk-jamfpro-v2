package disk_encryption_configurations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// DiskEncryptionConfigurationsServiceInterface defines the interface for Classic API disk encryption configuration operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/diskencryptionconfigurations
	DiskEncryptionConfigurationsServiceInterface interface {
		// ListDiskEncryptionConfigurations returns all disk encryption configurations.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findalldiskencryptionconfigurations
		ListDiskEncryptionConfigurations(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetDiskEncryptionConfigurationByID returns the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyid
		GetDiskEncryptionConfigurationByID(ctx context.Context, id int) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error)

		// GetDiskEncryptionConfigurationByName returns the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyname
		GetDiskEncryptionConfigurationByName(ctx context.Context, name string) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error)

		// CreateDiskEncryptionConfiguration creates a new disk encryption configuration.
		//
		// Returns the created resource ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/creatediskencryptionconfiguration
		CreateDiskEncryptionConfiguration(ctx context.Context, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateDiskEncryptionConfigurationByID updates the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyid
		UpdateDiskEncryptionConfigurationByID(ctx context.Context, id int, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// UpdateDiskEncryptionConfigurationByName updates the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyname
		UpdateDiskEncryptionConfigurationByName(ctx context.Context, name string, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error)

		// DeleteDiskEncryptionConfigurationByID removes the specified disk encryption configuration by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyid
		DeleteDiskEncryptionConfigurationByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteDiskEncryptionConfigurationByName removes the specified disk encryption configuration by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyname
		DeleteDiskEncryptionConfigurationByName(ctx context.Context, name string) (*interfaces.Response, error)
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

// ListDiskEncryptionConfigurations returns all disk encryption configurations.
// URL: GET /JSSResource/diskencryptionconfigurations
// https://developer.jamf.com/jamf-pro/reference/findalldiskencryptionconfigurations
func (s *Service) ListDiskEncryptionConfigurations(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicDiskEncryptionConfigurations, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDiskEncryptionConfigurationByID returns the specified disk encryption configuration by ID.
// URL: GET /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyid
func (s *Service) GetDiskEncryptionConfigurationByID(ctx context.Context, id int) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	var result ResourceDiskEncryptionConfiguration

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDiskEncryptionConfigurationByName returns the specified disk encryption configuration by name.
// URL: GET /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/finddiskencryptionconfigurationsbyname
func (s *Service) GetDiskEncryptionConfigurationByName(ctx context.Context, name string) (*ResourceDiskEncryptionConfiguration, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	var result ResourceDiskEncryptionConfiguration

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateDiskEncryptionConfiguration creates a new disk encryption configuration.
// URL: POST /JSSResource/diskencryptionconfigurations/id/0
// Returns the created resource ID.
// https://developer.jamf.com/jamf-pro/reference/creatediskencryptionconfiguration
func (s *Service) CreateDiskEncryptionConfiguration(ctx context.Context, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicDiskEncryptionConfigurations)

	var result CreateUpdateResponse

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDiskEncryptionConfigurationByID updates the specified disk encryption configuration by ID.
// URL: PUT /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyid
func (s *Service) UpdateDiskEncryptionConfigurationByID(ctx context.Context, id int, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDiskEncryptionConfigurationByName updates the specified disk encryption configuration by name.
// URL: PUT /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatediskencryptionconfigurationbyname
func (s *Service) UpdateDiskEncryptionConfigurationByName(ctx context.Context, name string, req *RequestDiskEncryptionConfiguration) (*CreateUpdateResponse, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("disk encryption configuration name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	var result CreateUpdateResponse

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteDiskEncryptionConfigurationByID removes the specified disk encryption configuration by ID.
// URL: DELETE /JSSResource/diskencryptionconfigurations/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyid
func (s *Service) DeleteDiskEncryptionConfigurationByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("disk encryption configuration ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicDiskEncryptionConfigurations, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteDiskEncryptionConfigurationByName removes the specified disk encryption configuration by name.
// URL: DELETE /JSSResource/diskencryptionconfigurations/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletediskencryptionconfigurationbyname
func (s *Service) DeleteDiskEncryptionConfigurationByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("disk encryption configuration name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicDiskEncryptionConfigurations, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
