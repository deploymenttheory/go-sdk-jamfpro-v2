package digicert

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the DigiCert Trust Lifecycle Manager-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
	Digicert struct {
		client client.Client
	}
)

func NewDigicert(client client.Client) *Digicert {
	return &Digicert{client: client}
}

// Create creates a new DigiCert Trust Lifecycle Manager configuration.
// URL: POST /api/v1/pki/digicert/trust-lifecycle-manager
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-digicert-trust-lifecycle-manager
func (s *Digicert) Create(ctx context.Context, request *ResourceDigicertTrustLifecycleManager) (*ResponseDigicertTrustLifecycleManagerCreated, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProTrustLifecycleManagerV1

	var result ResponseDigicertTrustLifecycleManagerCreated

	headers := map[string]string{
		"Accept":       constants.AcceptAny,
		"Content-Type": constants.ApplicationJSON,
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
func (s *Digicert) GetByID(ctx context.Context, id string) (*ResponseDigicertTrustLifecycleManager, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProTrustLifecycleManagerV1, id)

	var result ResponseDigicertTrustLifecycleManager

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Digicert) UpdateByID(ctx context.Context, id string, request *ResourceDigicertTrustLifecycleManager) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProTrustLifecycleManagerV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationMergePatchJSON,
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
func (s *Digicert) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProTrustLifecycleManagerV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Digicert) ValidateClientCertificate(ctx context.Context, request *ValidateClientCertificateRequest) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/validate-client-certificate", constants.EndpointJamfProTrustLifecycleManagerV1)

	headers := map[string]string{
		"Content-Type": constants.ApplicationJSON,
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
func (s *Digicert) GetConnectionStatusByID(ctx context.Context, id string) (*ConnectionStatusResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection-status", constants.EndpointJamfProTrustLifecycleManagerV1, id)

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

// GetDependenciesByID returns the list of dependencies for a DigiCert Trust Lifecycle Manager configuration.
// URL: GET /api/v1/pki/digicert/trust-lifecycle-manager/{id}/dependencies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-digicert-trust-lifecycle-manager-id-dependencies
func (s *Digicert) GetDependenciesByID(ctx context.Context, id string) (*DependenciesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependencies", constants.EndpointJamfProTrustLifecycleManagerV1, id)

	var result DependenciesResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
