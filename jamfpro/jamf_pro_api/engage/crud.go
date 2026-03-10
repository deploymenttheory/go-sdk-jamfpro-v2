package engage

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Engage-related methods of the Jamf Pro API.
	//
	// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
	Engage struct {
		client client.Client
	}
)

// NewService returns a new Engage Service backed by the provided HTTP client.
func NewEngage(client client.Client) *Engage {
	return &Engage{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Engage Operations
// -----------------------------------------------------------------------------

// GetV2 retrieves the Engage settings from the Jamf Pro server.
// URL: GET /api/v2/engage
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
func (s *Engage) GetV2(ctx context.Context) (*ResourceEngageSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProEngageV2

	var result ResourceEngageSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV2 updates the Engage settings on the Jamf Pro server.
// URL: PUT /api/v2/engage
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/put_v2-engage
func (s *Engage) UpdateV2(ctx context.Context, settings *ResourceEngageSettings) (*ResourceEngageSettings, *resty.Response, error) {
	if settings == nil {
		return nil, nil, fmt.Errorf("settings cannot be nil")
	}

	endpoint := constants.EndpointJamfProEngageV2

	var result ResourceEngageSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(settings).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV2 returns the history object for Engage settings.
// URL: GET /api/v2/engage/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage-history
func (s *Engage) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEngageV2)

	var result HistoryResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to the Engage settings history.
// URL: POST /api/v2/engage/history
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/post_v2-engage-history
func (s *Engage) AddHistoryNotesV2(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEngageV2)

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
