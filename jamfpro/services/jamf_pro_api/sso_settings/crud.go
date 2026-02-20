package sso_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
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

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointSsoV3, nil, headers, &result)
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

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointSsoV3, request, headers, &result)
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

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointDependenciesV3, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
