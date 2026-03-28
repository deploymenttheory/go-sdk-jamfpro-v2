package sso_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SSO settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso
	SsoSettings struct {
		client client.Client
	}
)

func NewSsoSettings(client client.Client) *SsoSettings {
	return &SsoSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Settings Operations
// -----------------------------------------------------------------------------

// GetV3 retrieves current Jamf SSO settings.
// URL: GET /api/v3/sso
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso
func (s *SsoSettings) GetV3(ctx context.Context) (*ResourceSsoSettings, *resty.Response, error) {
	var result ResourceSsoSettings

	endpoint := constants.EndpointJamfProSsoV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV3 updates SSO settings.
// URL: PUT /api/v3/sso
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-sso
func (s *SsoSettings) UpdateV3(ctx context.Context, request *ResourceSsoSettings) (*ResourceSsoSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if _, ok := validConfigurationTypes[request.ConfigurationType]; !ok {
		return nil, nil, fmt.Errorf("invalid configurationType %q: must be one of SAML, OIDC, OIDC_WITH_SAML", request.ConfigurationType)
	}

	if request.OidcSettings != nil {
		if _, ok := validUserMappings[request.OidcSettings.UserMapping]; !ok {
			return nil, nil, fmt.Errorf("invalid userMapping %q: must be one of USERNAME, EMAIL", request.OidcSettings.UserMapping)
		}
	}

	if request.SamlSettings != nil {
		if request.SamlSettings.MetadataSource != "" {
			if _, ok := validMetadataSources[request.SamlSettings.MetadataSource]; !ok {
				return nil, nil, fmt.Errorf("invalid metadataSource %q: must be one of URL, FILE, UNKNOWN", request.SamlSettings.MetadataSource)
			}
		}
		if request.SamlSettings.IdpProviderType != "" {
			if _, ok := validIdpProviderTypes[request.SamlSettings.IdpProviderType]; !ok {
				return nil, nil, fmt.Errorf("invalid idpProviderType %q: must be one of ADFS, OKTA, GOOGLE, SHIBBOLETH, ONELOGIN, PING, CENTRIFY, AZURE, OTHER", request.SamlSettings.IdpProviderType)
			}
		}
		if request.SamlSettings.UserMapping != "" {
			if _, ok := validUserMappings[request.SamlSettings.UserMapping]; !ok {
				return nil, nil, fmt.Errorf("invalid userMapping %q: must be one of USERNAME, EMAIL", request.SamlSettings.UserMapping)
			}
		}
	}

	var result ResourceSsoSettings

	endpoint := constants.EndpointJamfProSsoV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetEnrollmentCustomizationDependenciesV3 retrieves SSO enrollment customization dependencies.
// URL: GET /api/v3/sso/dependencies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-dependencies
func (s *SsoSettings) GetEnrollmentCustomizationDependenciesV3(ctx context.Context) (*ResponseSsoEnrollmentCustomizationDependencies, *resty.Response, error) {
	var result ResponseSsoEnrollmentCustomizationDependencies

	endpoint := constants.EndpointJamfProDependenciesV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DisableV3 disables SSO.
// URL: POST /api/v3/sso/disable
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-disable
func (s *SsoSettings) DisableV3(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProDisableV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV3 returns the history for SSO settings.
// URL: GET /api/v3/sso/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-history
func (s *SsoSettings) GetHistoryV3(ctx context.Context, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error) {
	var result HistoryListResponse

	endpoint := constants.EndpointJamfProHistoryV3

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryEntry
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteV3 adds a note to the history for SSO settings.
// URL: POST /api/v3/sso/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-sso-history
func (s *SsoSettings) AddHistoryNoteV3(ctx context.Context, request *AddHistoryNoteRequest) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProHistoryV3

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DownloadMetadataV3 downloads the SAML metadata file.
// URL: GET /api/v3/sso/metadata/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-sso-metadata-download
func (s *SsoSettings) DownloadMetadataV3(ctx context.Context) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProMetadataDownloadV3

	resp, data, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
