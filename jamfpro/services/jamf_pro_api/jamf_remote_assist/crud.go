package jamf_remote_assist

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfRemoteAssistServiceInterface defines the interface for Jamf Remote Assist operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
	JamfRemoteAssistServiceInterface interface {
		// ListSessionsV1 retrieves session history items (v1, no pagination).
		//
		// Returns up to 100 latest session history items.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-remote-assist-session
		ListSessionsV1(ctx context.Context) ([]SessionHistory, *interfaces.Response, error)

		// GetSessionByIDV1 retrieves a single session history item by ID (v1).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-remote-assist-session-id
		GetSessionByIDV1(ctx context.Context, id string) (*SessionHistory, *interfaces.Response, error)

		// ListSessionsV2 retrieves session history items with pagination and RSQL filtering (v2).
		//
		// Supports optional RSQL filtering, pagination and sorting via rsqlQuery
		// (keys: filter, sort, page, page-size).
		// Fields allowed in filter: sessionId, deviceId, sessionAdminId.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
		ListSessionsV2(ctx context.Context, rsqlQuery map[string]string) (*ListSessionsResponse, *interfaces.Response, error)

		// GetSessionByIDV2 retrieves a single session history item by ID with details (v2).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session-id
		GetSessionByIDV2(ctx context.Context, id string) (*SessionHistory, *interfaces.Response, error)

		// ExportSessionsV2 exports Jamf Remote Assist sessions history.
		//
		// Returns CSV or JSON format based on Accept header.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-jamf-remote-assist-session-export
		ExportSessionsV2(ctx context.Context, request *ExportSessionsRequest, acceptType string) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the Jamf Remote Assist-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfRemoteAssistServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Remote Assist Operations (V1)
// -----------------------------------------------------------------------------

// ListSessionsV1 retrieves session history items (v1, no pagination).
// URL: GET /api/v1/jamf-remote-assist/session
// Returns up to 100 latest session history items.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-remote-assist-session
func (s *Service) ListSessionsV1(ctx context.Context) ([]SessionHistory, *interfaces.Response, error) {
	endpoint := EndpointSessionV1

	var result []SessionHistory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list jamf remote assist sessions (v1): %w", err)
	}

	return result, resp, nil
}

// GetSessionByIDV1 retrieves a single session history item by ID (v1).
// URL: GET /api/v1/jamf-remote-assist/session/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-remote-assist-session-id
func (s *Service) GetSessionByIDV1(ctx context.Context, id string) (*SessionHistory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("session ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSessionV1, id)

	var result SessionHistory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get jamf remote assist session by ID (v1): %w", err)
	}

	return &result, resp, nil
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Remote Assist Operations (V2)
// -----------------------------------------------------------------------------

// ListSessionsV2 retrieves session history items with pagination and RSQL filtering (v2).
// URL: GET /api/v2/jamf-remote-assist/session
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Fields allowed in filter: sessionId, deviceId, sessionAdminId.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
func (s *Service) ListSessionsV2(ctx context.Context, rsqlQuery map[string]string) (*ListSessionsResponse, *interfaces.Response, error) {
	endpoint := EndpointSessionV2

	var result ListSessionsResponse

	mergePage := func(pageData []byte) error {
		var pageResponse ListSessionsResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list jamf remote assist sessions (v2): %w", err)
	}
	return &result, resp, nil
}

// GetSessionByIDV2 retrieves a single session history item by ID with details (v2).
// URL: GET /api/v2/jamf-remote-assist/session/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session-id
func (s *Service) GetSessionByIDV2(ctx context.Context, id string) (*SessionHistory, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("session ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSessionV2, id)

	var result SessionHistory

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get jamf remote assist session by ID (v2): %w", err)
	}

	return &result, resp, nil
}

// ExportSessionsV2 exports Jamf Remote Assist sessions history.
// URL: POST /api/v2/jamf-remote-assist/session/export
// Returns CSV or JSON format based on Accept header.
// acceptType should be "text/csv" or "application/json".
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-jamf-remote-assist-session-export
func (s *Service) ExportSessionsV2(ctx context.Context, request *ExportSessionsRequest, acceptType string) ([]byte, *interfaces.Response, error) {
	if request == nil {
		request = &ExportSessionsRequest{}
	}

	if acceptType == "" {
		acceptType = "text/csv"
	}

	endpoint := fmt.Sprintf("%s/export", EndpointSessionV2)

	var result []byte

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export jamf remote assist sessions: %w", err)
	}

	return result, resp, nil
}
