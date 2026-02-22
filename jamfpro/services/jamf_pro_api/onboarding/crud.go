package onboarding

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// OnboardingServiceInterface defines the interface for onboarding settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
	OnboardingServiceInterface interface {
		// GetV1 retrieves the current onboarding settings (Get Onboarding Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
		GetV1(ctx context.Context) (*ResponseOnboardingSettings, *interfaces.Response, error)

		// UpdateV1 updates the onboarding settings (Update Onboarding Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-onboarding
		UpdateV1(ctx context.Context, request *ResourceUpdateOnboardingSettings) (*ResponseOnboardingSettings, *interfaces.Response, error)

		// GetEligibleAppsV1 returns the list of eligible apps for onboarding (Get Eligible Apps).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-apps
		GetEligibleAppsV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error)

		// GetEligibleConfigurationProfilesV1 returns the list of eligible configuration profiles (Get Eligible Configuration Profiles).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-configuration-profiles
		GetEligibleConfigurationProfilesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error)

		// GetEligiblePoliciesV1 returns the list of eligible policies (Get Eligible Policies).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-policies
		GetEligiblePoliciesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error)

		// GetHistoryV1 retrieves the onboarding history.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNotesV1 adds notes to the onboarding history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-onboarding-history
		AddHistoryNotesV1(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error)

		// ExportHistoryV1 exports the onboarding history in the specified format (JSON or CSV).
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size, export-fields, export-labels).
		// The Accept header determines the export format (application/json or text/csv).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-onboarding-history-export
		ExportHistoryV1(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, req *RequestExportHistory) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the onboarding-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ OnboardingServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Onboarding Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves the current onboarding settings.
// URL: GET /api/v1/onboarding
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
func (s *Service) GetV1(ctx context.Context) (*ResponseOnboardingSettings, *interfaces.Response, error) {
	var result ResponseOnboardingSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointOnboardingV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the onboarding settings.
// URL: PUT /api/v1/onboarding
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-onboarding
func (s *Service) UpdateV1(ctx context.Context, request *ResourceUpdateOnboardingSettings) (*ResponseOnboardingSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResponseOnboardingSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, EndpointOnboardingV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetEligibleAppsV1 returns the list of eligible apps for onboarding.
// URL: GET /api/v1/onboarding/eligible-apps
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-apps
func (s *Service) GetEligibleAppsV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error) {
	var result ResponseEligibilityList
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointEligibleApps, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetEligibleConfigurationProfilesV1 returns the list of eligible configuration profiles.
// URL: GET /api/v1/onboarding/eligible-configuration-profiles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-configuration-profiles
func (s *Service) GetEligibleConfigurationProfilesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error) {
	var result ResponseEligibilityList
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointEligibleConfigurationProfiles, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetEligiblePoliciesV1 returns the list of eligible policies.
// URL: GET /api/v1/onboarding/eligible-policies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-policies
func (s *Service) GetEligiblePoliciesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *interfaces.Response, error) {
	var result ResponseEligibilityList
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointEligiblePolicies, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetHistoryV1 retrieves the onboarding history.
// URL: GET /api/v1/onboarding/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-history
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/history", EndpointOnboardingV1)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV1 adds notes to the onboarding history.
// URL: POST /api/v1/onboarding/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-onboarding-history
func (s *Service) AddHistoryNotesV1(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", EndpointOnboardingV1)

	var result ResponseAddHistoryNotes

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ExportHistoryV1 exports the onboarding history in the specified format (JSON or CSV).
// URL: POST /api/v1/onboarding/history/export
// acceptHeader should be "application/json" or "text/csv".
// rsqlQuery supports: filter (RSQL), sort, page, page-size, export-fields, export-labels (all optional).
// https://developer.jamf.com/jamf-pro/reference/post_v1-onboarding-history-export
func (s *Service) ExportHistoryV1(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, req *RequestExportHistory) ([]byte, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/history/export", EndpointOnboardingV1)

	if acceptHeader == "" {
		acceptHeader = mime.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptHeader,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, rsqlQuery, headers)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
