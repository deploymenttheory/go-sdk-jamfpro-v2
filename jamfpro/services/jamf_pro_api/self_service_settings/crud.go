package self_service_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// SelfServiceSettingsServiceInterface defines the interface for self-service settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	SelfServiceSettingsServiceInterface interface {
		// Get retrieves self-service settings (Get Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
		Get(ctx context.Context) (*ResourceSelfServiceSettings, *interfaces.Response, error)

		// Update updates self-service settings (Update Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-self-service-settings
		Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *interfaces.Response, error)

		// GetHistoryV1 returns the paginated history for Self Service settings.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNotesV1 adds a note to the Self Service settings history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-settings-history
		AddHistoryNotesV1(ctx context.Context, req *AddHistoryNotesRequest) (*AddHistoryNotesResponse, *interfaces.Response, error)
	}

	// Service handles communication with the self-service settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SelfServiceSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves self-service settings.
// URL: GET /api/v1/self-service/settings
func (s *Service) Get(ctx context.Context) (*ResourceSelfServiceSettings, *interfaces.Response, error) {
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
func (s *Service) Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *interfaces.Response, error) {
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
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	var result HistoryResponse

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
				var history HistoryObject
				if err := mapstructure.Decode(item, &history); err != nil {
					return fmt.Errorf("failed to decode history item: %w", err)
				}
				result.Results = append(result.Results, history)
			}
		}

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

	return &result, resp, nil
}

// AddHistoryNotesV1 adds a note to the Self Service settings history.
// URL: POST /api/v1/self-service/settings/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-settings-history
func (s *Service) AddHistoryNotesV1(ctx context.Context, req *AddHistoryNotesRequest) (*AddHistoryNotesResponse, *interfaces.Response, error) {
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
