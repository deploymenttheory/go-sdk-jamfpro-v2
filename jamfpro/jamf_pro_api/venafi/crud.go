package venafi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Venafi PKI-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-pki-venafi
	Venafi struct {
		client client.Client
	}
)

func NewVenafi(client client.Client) *Venafi {
	return &Venafi{client: client}
}

// Create creates a new Venafi PKI configuration.
// URL: POST /api/v1/pki/venafi
func (s *Venafi) Create(ctx context.Context, request *ResourceVenafi) (*ResponseVenafiCreated, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProVenafiV1

	var result ResponseVenafiCreated

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

// GetByID returns the Venafi PKI configuration by ID.
// URL: GET /api/v1/pki/venafi/{id}
func (s *Venafi) GetByID(ctx context.Context, id string) (*ResponseVenafi, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVenafiV1, id)

	var result ResponseVenafi

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the Venafi PKI configuration by ID.
// URL: PATCH /api/v1/pki/venafi/{id}
func (s *Venafi) UpdateByID(ctx context.Context, id string, request *ResourceVenafi) (*ResponseVenafi, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVenafiV1, id)

	var result ResponseVenafi

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID deletes the Venafi PKI configuration by ID.
// URL: DELETE /api/v1/pki/venafi/{id}
func (s *Venafi) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetConnectionStatus tests communication between Jamf Pro and the PKI Proxy Server.
// URL: GET /api/v1/pki/venafi/{id}/connection-status
func (s *Venafi) GetConnectionStatus(ctx context.Context, id string) (*ResponseConnectionStatus, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/connection-status", constants.EndpointJamfProVenafiV1, id)

	var result ResponseConnectionStatus

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDependentProfiles returns configuration profiles using the Venafi CA.
// URL: GET /api/v1/pki/venafi/{id}/dependent-profiles
func (s *Venafi) GetDependentProfiles(ctx context.Context, id string) (*ResponseDependentProfiles, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/dependent-profiles", constants.EndpointJamfProVenafiV1, id)

	var result ResponseDependentProfiles

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistory returns the history for a Venafi configuration.
// URL: GET /api/v1/pki/venafi/{id}/history
func (s *Venafi) GetHistory(ctx context.Context, id string, query map[string]string) (*ResponseHistory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProVenafiV1, id)

	var result ResponseHistory

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, query, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNote adds a note to the history for a Venafi configuration.
// URL: POST /api/v1/pki/venafi/{id}/history
func (s *Venafi) AddHistoryNote(ctx context.Context, id string, request *HistoryNoteRequest) (*ResponseVenafiCreated, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProVenafiV1, id)

	var result ResponseVenafiCreated

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

// GetJamfPublicKey downloads the certificate for securing Jamf Pro to PKI Proxy communication.
// URL: GET /api/v1/pki/venafi/{id}/jamf-public-key
func (s *Venafi) GetJamfPublicKey(ctx context.Context, id string) (*resty.Response, []byte, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/jamf-public-key", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationPEMCertificateChain,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return resp, nil, err
	}

	return resp, data, nil
}

// GetProxyTrustStore downloads the PKI Proxy Server public key.
// URL: GET /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Venafi) GetProxyTrustStore(ctx context.Context, id string) (*resty.Response, []byte, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationPEMCertificateChain,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return resp, nil, err
	}

	return resp, data, nil
}

// RegenerateJamfPublicKeyByIDV1 regenerates the Jamf public key for Venafi configuration.
// URL: POST /api/v1/pki/venafi/{id}/jamf-public-key/regenerate
func (s *Venafi) RegenerateJamfPublicKeyByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/jamf-public-key/regenerate", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UploadProxyTrustStoreByIDV1 uploads the PKI Proxy Server public key (PEM certificate).
// URL: POST /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Venafi) UploadProxyTrustStoreByIDV1(ctx context.Context, id string, pemCertificate []byte) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	if len(pemCertificate) == 0 {
		return nil, fmt.Errorf("pem certificate is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationPEMCertificateChain,
	}

	resp, err := s.client.Post(ctx, endpoint, pemCertificate, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteProxyTrustStoreByIDV1 removes the PKI Proxy Server public key.
// URL: DELETE /api/v1/pki/venafi/{id}/proxy-trust-store
func (s *Venafi) DeleteProxyTrustStoreByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/proxy-trust-store", constants.EndpointJamfProVenafiV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
