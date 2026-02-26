package digicert

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DigicertServiceInterface defines the interface for DigiCert Trust Lifecycle Manager operations.
	// Supports PKI integration with DigiCert Trust Lifecycle Manager for certificate management.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
	DigicertServiceInterface interface {
		// Create creates a new DigiCert Trust Lifecycle Manager configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
		Create(ctx context.Context, request *ResourceDigicertTrustLifecycleManager) (*ResponseDigicertTrustLifecycleManagerCreated, *interfaces.Response, error)

		// GetByID returns the DigiCert Trust Lifecycle Manager configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id
		GetByID(ctx context.Context, id string) (*ResponseDigicertTrustLifecycleManager, *interfaces.Response, error)

		// UpdateByID updates the DigiCert Trust Lifecycle Manager configuration by ID using merge-patch semantics.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-pki-digicert-trust-lifecycle-manager-id
		UpdateByID(ctx context.Context, id string, request *ResourceDigicertTrustLifecycleManager) (*interfaces.Response, error)

		// DeleteByID deletes the DigiCert Trust Lifecycle Manager configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-digicert-trust-lifecycle-manager-id
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)

		// ValidateClientCertificate validates the DigiCert Trust Lifecycle Manager client certificate.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager-validate-client-certificate
		ValidateClientCertificate(ctx context.Context, request *ValidateClientCertificateRequest) (*interfaces.Response, error)

		// GetConnectionStatusByID returns the connection status for a DigiCert Trust Lifecycle Manager configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id-connection-status
		GetConnectionStatusByID(ctx context.Context, id string) (*ConnectionStatusResponse, *interfaces.Response, error)

		// GetDependenciesByID returns the list of dependencies for a DigiCert Trust Lifecycle Manager configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id-dependencies
		GetDependenciesByID(ctx context.Context, id string) (*DependenciesResponse, *interfaces.Response, error)
	}

	// Service handles communication with the DigiCert Trust Lifecycle Manager-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DigicertServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// Create creates a new DigiCert Trust Lifecycle Manager configuration.
// URL: POST /api/v1/pki/digicert/trust-lifecycle-manager
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
func (s *Service) Create(ctx context.Context, request *ResourceDigicertTrustLifecycleManager) (*ResponseDigicertTrustLifecycleManagerCreated, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointTrustLifecycleManagerV1

	var result ResponseDigicertTrustLifecycleManagerCreated

	headers := map[string]string{
		"Accept":       "*/*",
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the DigiCert Trust Lifecycle Manager configuration by ID.
// URL: GET /api/v1/pki/digicert/trust-lifecycle-manager/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id
func (s *Service) GetByID(ctx context.Context, id string) (*ResponseDigicertTrustLifecycleManager, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointTrustLifecycleManagerV1, id)

	var result ResponseDigicertTrustLifecycleManager

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the DigiCert Trust Lifecycle Manager configuration by ID using merge-patch semantics.
// URL: PATCH /api/v1/pki/digicert/trust-lifecycle-manager/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-pki-digicert-trust-lifecycle-manager-id
func (s *Service) UpdateByID(ctx context.Context, id string, request *ResourceDigicertTrustLifecycleManager) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointTrustLifecycleManagerV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationMergePatchJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByID deletes the DigiCert Trust Lifecycle Manager configuration by ID.
// URL: DELETE /api/v1/pki/digicert/trust-lifecycle-manager/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-digicert-trust-lifecycle-manager-id
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointTrustLifecycleManagerV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ValidateClientCertificate validates the DigiCert Trust Lifecycle Manager client certificate.
// URL: POST /api/v1/pki/digicert/trust-lifecycle-manager/validate-client-certificate
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager-validate-client-certificate
func (s *Service) ValidateClientCertificate(ctx context.Context, request *ValidateClientCertificateRequest) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/validate-client-certificate", EndpointTrustLifecycleManagerV1)

	headers := map[string]string{
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetConnectionStatusByID returns the connection status for a DigiCert Trust Lifecycle Manager configuration.
// URL: GET /api/v1/pki/digicert/trust-lifecycle-manager/{id}/connection-status
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id-connection-status
func (s *Service) GetConnectionStatusByID(ctx context.Context, id string) (*ConnectionStatusResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection-status", EndpointTrustLifecycleManagerV1, id)

	var result ConnectionStatusResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDependenciesByID returns the list of dependencies for a DigiCert Trust Lifecycle Manager configuration.
// URL: GET /api/v1/pki/digicert/trust-lifecycle-manager/{id}/dependencies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id-dependencies
func (s *Service) GetDependenciesByID(ctx context.Context, id string) (*DependenciesResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", EndpointTrustLifecycleManagerV1, id)

	var result DependenciesResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
