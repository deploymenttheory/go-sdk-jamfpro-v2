package self_service_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SelfServiceSettingsServiceInterface defines the interface for self-service settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	SelfServiceSettingsServiceInterface interface {
		// Get retrieves self-service settings (Get Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
		Get(ctx context.Context) (*ResourceSelfServiceSettings, *resty.Response, error)

		// Update updates self-service settings (Update Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-self-service-settings
		Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *resty.Response, error)

		// GetHistoryV1 returns the paginated history for Self Service settings.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNotesV1 adds a note to the Self Service settings history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-settings-history
		AddHistoryNotesV1(ctx context.Context, req *AddHistoryNotesRequest) (*AddHistoryNotesResponse, *resty.Response, error)
	}

	// Service handles communication with the self-service settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	SelfServiceSettings struct {
		client transport.HTTPClient
	}
)

var _ SelfServiceSettingsServiceInterface = (*SelfServiceSettings)(nil)

func NewSelfServiceSettings(client transport.HTTPClient) *SelfServiceSettings {
	return &SelfServiceSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves self-service settings.
// URL: GET /api/v1/self-service/settings
func (s *SelfServiceSettings) Get(ctx context.Context) (*ResourceSelfServiceSettings, *resty.Response, error) {
	var result ResourceSelfServiceSettings

	endpoint := EndpointSelfServiceSettingsV1
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := EndpointSelfServiceSettingsV1
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

	endpoint := EndpointSelfServiceSettingsHistoryV1
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
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
	endpoint := EndpointSelfServiceSettingsHistoryV1
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
