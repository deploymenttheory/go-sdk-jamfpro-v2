package sso_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// SsoSettingsServiceInterface defines the interface for SSO settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-sso
	SsoSettingsServiceInterface interface {
		// GetV3 retrieves current Jamf SSO settings (Get SSO Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso
		GetV3(ctx context.Context) (*ResourceSsoSettings, *interfaces.Response, error)

		// UpdateV3 updates SSO settings (Update SSO Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-sso
		UpdateV3(ctx context.Context, request *ResourceSsoSettings) (*ResourceSsoSettings, *interfaces.Response, error)

		// GetEnrollmentCustomizationDependenciesV3 retrieves SSO enrollment customization dependencies (Get SSO Enrollment Customization Dependencies).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-dependencies
		GetEnrollmentCustomizationDependenciesV3(ctx context.Context) (*ResponseSsoEnrollmentCustomizationDependencies, *interfaces.Response, error)

		// DisableV3 disables SSO (Disable SSO).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-disable
		DisableV3(ctx context.Context) (*interfaces.Response, error)

		// GetHistoryV3 returns the history for SSO settings (Get SSO History).
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-history
		GetHistoryV3(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *interfaces.Response, error)

		// AddHistoryNoteV3 adds a note to the history for SSO settings (Add SSO History Note).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-history
		AddHistoryNoteV3(ctx context.Context, request *AddHistoryNoteRequest) (*CreateResponse, *interfaces.Response, error)

		// DownloadMetadataV3 downloads the SAML metadata file (Download SAML Metadata).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-metadata-download
		DownloadMetadataV3(ctx context.Context) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the SSO settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SsoSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Settings Operations
// -----------------------------------------------------------------------------

// GetV3 retrieves current Jamf SSO settings.
// URL: GET /api/v3/sso
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso
func (s *Service) GetV3(ctx context.Context) (*ResourceSsoSettings, *interfaces.Response, error) {
	var result ResourceSsoSettings

	endpoint := EndpointSsoV3
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV3 updates SSO settings.
// URL: PUT /api/v3/sso
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-sso
func (s *Service) UpdateV3(ctx context.Context, request *ResourceSsoSettings) (*ResourceSsoSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSsoSettings

	endpoint := EndpointSsoV3
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

// GetEnrollmentCustomizationDependenciesV3 retrieves SSO enrollment customization dependencies.
// URL: GET /api/v3/sso/dependencies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-dependencies
func (s *Service) GetEnrollmentCustomizationDependenciesV3(ctx context.Context) (*ResponseSsoEnrollmentCustomizationDependencies, *interfaces.Response, error) {
	var result ResponseSsoEnrollmentCustomizationDependencies

	endpoint := EndpointDependenciesV3
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DisableV3 disables SSO.
// URL: POST /api/v3/sso/disable
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-disable
func (s *Service) DisableV3(ctx context.Context) (*interfaces.Response, error) {
	endpoint := EndpointDisableV3
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV3 returns the history for SSO settings.
// URL: GET /api/v3/sso/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-history
func (s *Service) GetHistoryV3(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *interfaces.Response, error) {
	var result HistoryListResponse

	endpoint := EndpointHistoryV3

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
				var entry HistoryEntry
				if err := mapstructure.Decode(item, &entry); err != nil {
					return fmt.Errorf("failed to decode history entry: %w", err)
				}
				result.Results = append(result.Results, entry)
			}
		}

		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNoteV3 adds a note to the history for SSO settings.
// URL: POST /api/v3/sso/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-history
func (s *Service) AddHistoryNoteV3(ctx context.Context, request *AddHistoryNoteRequest) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointHistoryV3
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

// DownloadMetadataV3 downloads the SAML metadata file.
// URL: GET /api/v3/sso/metadata/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-metadata-download
func (s *Service) DownloadMetadataV3(ctx context.Context) ([]byte, *interfaces.Response, error) {
	var result []byte

	endpoint := EndpointMetadataDownloadV3
	headers := map[string]string{
		"Accept": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
