package client_checkin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// ClientCheckinServiceInterface defines the interface for client check-in settings (singleton).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
	ClientCheckinServiceInterface interface {
		// GetV3 returns the current client check-in settings (Get Client Check-In settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
		GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error)

		// UpdateV3 updates the client check-in settings (Update Client Check-In object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
		UpdateV3(ctx context.Context, request *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error)

		// GetHistoryV3 returns the client check-in history object (Get Client Check-In history object).
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in-history
		GetHistoryV3(ctx context.Context, rsqlQuery map[string]string) (*ResourceClientCheckinHistory, *interfaces.Response, error)

		// AddHistoryNoteV3 adds a note to the client check-in history (Add a Note to Client Check-In History).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-check-in-history
		AddHistoryNoteV3(ctx context.Context, request *RequestClientCheckinHistoryNote) (*interfaces.Response, error)
	}

	// Service handles communication with the client check-in-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ClientCheckinServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Client Check-In Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current client check-in settings.
// URL: GET /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
func (s *Service) GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	var result ResourceClientCheckinSettings

	endpoint := EndpointClientCheckinV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV3 returns the client check-in history object.
// URL: GET /api/v3/check-in/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in-history
func (s *Service) GetHistoryV3(ctx context.Context, rsqlQuery map[string]string) (*ResourceClientCheckinHistory, *interfaces.Response, error) {
	var result ResourceClientCheckinHistory

	endpoint := EndpointClientCheckinHistoryV3

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
				var entry ResourceClientCheckinHistoryEntry
				if err := mapstructure.Decode(item, &entry); err != nil {
					return fmt.Errorf("failed to decode client check-in history entry: %w", err)
				}
				result.Results = append(result.Results, entry)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get client check-in history: %w", err)
	}

	return &result, resp, nil
}

// AddHistoryNoteV3 adds a note to the client check-in history.
// URL: POST /api/v3/check-in/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-check-in-history
func (s *Service) AddHistoryNoteV3(ctx context.Context, request *RequestClientCheckinHistoryNote) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointClientCheckinHistoryV3

	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UpdateV3 updates the client check-in settings.
// URL: PUT /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
func (s *Service) UpdateV3(ctx context.Context, request *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceClientCheckinSettings

	endpoint := EndpointClientCheckinV3

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
