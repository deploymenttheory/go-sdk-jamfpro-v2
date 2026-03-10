package self_service_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the self-service settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	SelfServiceSettings struct {
		client client.Client
	}
)

func NewSelfServiceSettings(client client.Client) *SelfServiceSettings {
	return &SelfServiceSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves self-service settings.
// URL: GET /api/v1/self-service/settings
func (s *SelfServiceSettings) Get(ctx context.Context) (*ResourceSelfServiceSettings, *resty.Response, error) {
	var result ResourceSelfServiceSettings

	endpoint := constants.EndpointJamfProSelfServiceSettingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Update updates self-service settings.
// URL: PUT /api/v1/self-service/settings
func (s *SelfServiceSettings) Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSelfServiceSettings

	endpoint := constants.EndpointJamfProSelfServiceSettingsV1

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

// GetHistoryV1 returns the paginated history for Self Service settings.
// URL: GET /api/v1/self-service/settings/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings-history
func (s *SelfServiceSettings) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProSelfServiceSettingsHistoryV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get self service settings history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotesV1 adds a note to the Self Service settings history.
// URL: POST /api/v1/self-service/settings/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-settings-history
func (s *SelfServiceSettings) AddHistoryNotesV1(ctx context.Context, req *AddHistoryNotesRequest) (*AddHistoryNotesResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	var result AddHistoryNotesResponse
	endpoint := constants.EndpointJamfProSelfServiceSettingsHistoryV1

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
