package onboarding

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the onboarding-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
	Onboarding struct {
		client client.Client
	}
)

func NewOnboarding(client client.Client) *Onboarding {
	return &Onboarding{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Onboarding Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves the current onboarding settings.
// URL: GET /api/v1/onboarding
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding
func (s *Onboarding) GetV1(ctx context.Context) (*ResponseOnboardingSettings, *resty.Response, error) {
	var result ResponseOnboardingSettings
	endpoint := constants.EndpointJamfProOnboardingV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the onboarding settings.
// URL: PUT /api/v1/onboarding
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-onboarding
func (s *Onboarding) UpdateV1(ctx context.Context, request *ResourceUpdateOnboardingSettings) (*ResponseOnboardingSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	for _, item := range request.OnboardingItems {
		if _, ok := validSelfServiceEntityTypes[item.SelfServiceEntityType]; !ok {
			return nil, nil, fmt.Errorf("invalid selfServiceEntityType %q: must be one of OS_X_POLICY, OS_X_CONFIG_PROFILE, OS_X_MAC_APP, OS_X_APP_INSTALLER, OS_X_EBOOK, OS_X_PATCH_POLICY, UNKNOWN", item.SelfServiceEntityType)
		}
	}

	var result ResponseOnboardingSettings
	endpoint := constants.EndpointJamfProOnboardingV1

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

// GetEligibleAppsV1 returns the list of eligible apps for onboarding.
// URL: GET /api/v1/onboarding/eligible-apps
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-apps
func (s *Onboarding) GetEligibleAppsV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *resty.Response, error) {
	var result ResponseEligibilityList

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEligibilityListItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProEligibleApps

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetEligibleConfigurationProfilesV1 returns the list of eligible configuration profiles.
// URL: GET /api/v1/onboarding/eligible-configuration-profiles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-configuration-profiles
func (s *Onboarding) GetEligibleConfigurationProfilesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *resty.Response, error) {
	var result ResponseEligibilityList

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEligibilityListItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProEligibleConfigurationProfiles

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetEligiblePoliciesV1 returns the list of eligible policies.
// URL: GET /api/v1/onboarding/eligible-policies
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-eligible-policies
func (s *Onboarding) GetEligiblePoliciesV1(ctx context.Context, query map[string]string) (*ResponseEligibilityList, *resty.Response, error) {
	var result ResponseEligibilityList

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEligibilityListItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProEligiblePolicies

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetHistoryV1 retrieves the onboarding history.
// URL: GET /api/v1/onboarding/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-onboarding-history
func (s *Onboarding) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProOnboardingV1)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceHistoryEntry
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

// AddHistoryNotesV1 adds notes to the onboarding history.
// URL: POST /api/v1/onboarding/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-onboarding-history
func (s *Onboarding) AddHistoryNotesV1(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProOnboardingV1)

	var result ResponseAddHistoryNotes

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
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
func (s *Onboarding) ExportHistoryV1(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, req *RequestExportHistory) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history/export", constants.EndpointJamfProOnboardingV1)

	if acceptHeader == "" {
		acceptHeader = constants.ApplicationJSON
	}

	resp, data, err := s.client.NewRequest(ctx).
		SetHeader("Accept", acceptHeader).
		SetQueryParams(rsqlQuery).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
