package venafi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// VenafiServiceInterface defines the interface for Venafi PKI operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi
	VenafiServiceInterface interface {
		// Create creates a new Venafi PKI configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi
		Create(ctx context.Context, request *ResourceVenafi) (*ResponseVenafiCreated, *interfaces.Response, error)

		// GetByID returns the Venafi PKI configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id
		GetByID(ctx context.Context, id string) (*ResponseVenafi, *interfaces.Response, error)

		// UpdateByID updates the Venafi PKI configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-pki-venafi-id
		UpdateByID(ctx context.Context, id string, request *ResourceVenafi) (*ResponseVenafi, *interfaces.Response, error)

		// DeleteByID deletes the Venafi PKI configuration by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-venafi-id
		DeleteByID(ctx context.Context, id string) (*interfaces.Response, error)

		// GetConnectionStatus tests communication between Jamf Pro and the PKI Proxy Server.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id-connection-status
		GetConnectionStatus(ctx context.Context, id string) (*ResponseConnectionStatus, *interfaces.Response, error)

		// GetDependentProfiles returns configuration profiles using the Venafi CA.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id-dependent-profiles
		GetDependentProfiles(ctx context.Context, id string) (*ResponseDependentProfiles, *interfaces.Response, error)

		// GetHistory returns the history for a Venafi configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id-history
		GetHistory(ctx context.Context, id string, query map[string]string) (*ResponseHistory, *interfaces.Response, error)

		// AddHistoryNote adds a note to the history for a Venafi configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi-id-history
		AddHistoryNote(ctx context.Context, id string, request *HistoryNoteRequest) (*ResponseVenafiCreated, *interfaces.Response, error)

		// GetJamfPublicKey downloads the certificate for securing Jamf Pro to PKI Proxy communication.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id-jamf-public-key
		GetJamfPublicKey(ctx context.Context, id string) (*interfaces.Response, []byte, error)

		// GetProxyTrustStore downloads the PKI Proxy Server public key.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-venafi-id-proxy-trust-store
		GetProxyTrustStore(ctx context.Context, id string) (*interfaces.Response, []byte, error)

		// RegenerateJamfPublicKeyByIDV1 regenerates the Jamf public key for Venafi configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi-id-jamf-public-key-regenerate
		RegenerateJamfPublicKeyByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// UploadProxyTrustStoreByIDV1 uploads the PKI Proxy Server public key (PEM certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi-id-proxy-trust-store
		UploadProxyTrustStoreByIDV1(ctx context.Context, id string, pemCertificate []byte) (*interfaces.Response, error)

		// DeleteProxyTrustStoreByIDV1 removes the PKI Proxy Server public key.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-pki-venafi-id-proxy-trust-store
		DeleteProxyTrustStoreByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the Venafi PKI-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ VenafiServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// Create creates a new Venafi PKI configuration.
// URL: POST /api/v1/pki/venafi
func (s *Service) Create(ctx context.Context, request *ResourceVenafi) (*ResponseVenafiCreated, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointVenafiV1

	var result ResponseVenafiCreated

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

// GetByID returns the Venafi PKI configuration by ID.
// URL: GET /api/v1/pki/venafi/{id}
func (s *Service) GetByID(ctx context.Context, id string) (*ResponseVenafi, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVenafiV1, id)

	var result ResponseVenafi

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the Venafi PKI configuration by ID.
// URL: PATCH /api/v1/pki/venafi/{id}
func (s *Service) UpdateByID(ctx context.Context, id string, request *ResourceVenafi) (*ResponseVenafi, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVenafiV1, id)

	var result ResponseVenafi

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID deletes the Venafi PKI configuration by ID.
// URL: DELETE /api/v1/pki/venafi/{id}
func (s *Service) DeleteByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetConnectionStatus tests communication between Jamf Pro and the PKI Proxy Server.
// URL: GET /api/v1/pki/venafi/{id}/connection-status
func (s *Service) GetConnectionStatus(ctx context.Context, id string) (*ResponseConnectionStatus, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection-status", EndpointVenafiV1, id)

	var result ResponseConnectionStatus

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDependentProfiles returns configuration profiles using the Venafi CA.
// URL: GET /api/v1/pki/venafi/{id}/dependent-profiles
func (s *Service) GetDependentProfiles(ctx context.Context, id string) (*ResponseDependentProfiles, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependent-profiles", EndpointVenafiV1, id)

	var result ResponseDependentProfiles

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistory returns the history for a Venafi configuration.
// URL: GET /api/v1/pki/venafi/{id}/history
func (s *Service) GetHistory(ctx context.Context, id string, query map[string]string) (*ResponseHistory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointVenafiV1, id)

	var result ResponseHistory

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var historyItem HistoryItem
				if err := mapstructure.Decode(item, &historyItem); err != nil {
					return fmt.Errorf("failed to decode history item: %w", err)
				}
				result.Results = append(result.Results, historyItem)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNote adds a note to the history for a Venafi configuration.
// URL: POST /api/v1/pki/venafi/{id}/history
func (s *Service) AddHistoryNote(ctx context.Context, id string, request *HistoryNoteRequest) (*ResponseVenafiCreated, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointVenafiV1, id)

	var result ResponseVenafiCreated

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

// GetJamfPublicKey downloads the certificate for securing Jamf Pro to PKI Proxy communication.
// URL: GET /api/v1/pki/venafi/{id}/jamf-public-key
func (s *Service) GetJamfPublicKey(ctx context.Context, id string) (*interfaces.Response, []byte, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/jamf-public-key", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationPEMCertificateChain,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return resp, nil, err
	}

	return resp, data, nil
}

// GetProxyTrustStore downloads the PKI Proxy Server public key.
// URL: GET /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Service) GetProxyTrustStore(ctx context.Context, id string) (*interfaces.Response, []byte, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationPEMCertificateChain,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return resp, nil, err
	}

	return resp, data, nil
}

// RegenerateJamfPublicKeyByIDV1 regenerates the Jamf public key for Venafi configuration.
// URL: POST /api/v1/pki/venafi/{id}/jamf-public-key/regenerate
func (s *Service) RegenerateJamfPublicKeyByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/jamf-public-key/regenerate", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UploadProxyTrustStoreByIDV1 uploads the PKI Proxy Server public key (PEM certificate).
// URL: POST /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Service) UploadProxyTrustStoreByIDV1(ctx context.Context, id string, pemCertificate []byte) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if len(pemCertificate) == 0 {
		return nil, fmt.Errorf("pem certificate is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationPEMCertificateChain,
	}

	resp, err := s.client.Post(ctx, endpoint, pemCertificate, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteProxyTrustStoreByIDV1 removes the PKI Proxy Server public key.
// URL: DELETE /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Service) DeleteProxyTrustStoreByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", EndpointVenafiV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
