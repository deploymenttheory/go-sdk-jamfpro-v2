package jamf_remote_assist

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Remote Assist-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session
	JamfRemoteAssist struct {
		client client.Client
	}
)

func NewJamfRemoteAssist(client client.Client) *JamfRemoteAssist {
	return &JamfRemoteAssist{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Remote Assist Operations (V1)
// -----------------------------------------------------------------------------

// ListSessionsV1 retrieves session history items (v1, no pagination).
// URL: GET /api/v1/jamf-remote-assist/session
// Returns up to 100 latest session history items.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-remote-assist-session
func (s *JamfRemoteAssist) ListSessionsV1(ctx context.Context) ([]SessionHistory, *resty.Response, error) {
	endpoint := constants.EndpointJamfProSessionV1

	var result []SessionHistory

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *JamfRemoteAssist) GetSessionByIDV1(ctx context.Context, id string) (*SessionHistory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("session ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSessionV1, id)

	var result SessionHistory

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *JamfRemoteAssist) ListSessionsV2(ctx context.Context, rsqlQuery map[string]string) (*ListSessionsResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProSessionV2

	var result ListSessionsResponse

	mergePage := func(pageData []byte) error {
		var pageItems []SessionHistory
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list jamf remote assist sessions (v2): %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetSessionByIDV2 retrieves a single session history item by ID with details (v2).
// URL: GET /api/v2/jamf-remote-assist/session/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-remote-assist-session-id
func (s *JamfRemoteAssist) GetSessionByIDV2(ctx context.Context, id string) (*SessionHistory, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("session ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSessionV2, id)

	var result SessionHistory

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *JamfRemoteAssist) ExportSessionsV2(ctx context.Context, request *ExportSessionsRequest, acceptType string) ([]byte, *resty.Response, error) {
	if request == nil {
		request = &ExportSessionsRequest{}
	}

	if acceptType == "" {
		acceptType = "text/csv"
	}

	endpoint := fmt.Sprintf("%s/export", constants.EndpointJamfProSessionV2)

	var result []byte

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export jamf remote assist sessions: %w", err)
	}

	return result, resp, nil
}
