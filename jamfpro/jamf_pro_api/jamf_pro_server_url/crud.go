package jamf_pro_server_url

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro server URL-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
	JamfProServerUrl struct {
		client client.Client
	}
)

func NewJamfProServerUrl(client client.Client) *JamfProServerUrl {
	return &JamfProServerUrl{client: client}
}

// GetV1 retrieves the Jamf Pro server URL settings.
// URL: GET /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
func (s *JamfProServerUrl) GetV1(ctx context.Context) (*ResourceJamfProServerURL, *resty.Response, error) {
	var result ResourceJamfProServerURL

	endpoint := constants.EndpointJamfProJamfProServerURLV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}

// UpdateV1 updates the Jamf Pro server URL settings.
// URL: PUT /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-pro-server-url
func (s *JamfProServerUrl) UpdateV1(ctx context.Context, request *ResourceJamfProServerURL) (*ResourceJamfProServerURL, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceJamfProServerURL

	endpoint := constants.EndpointJamfProJamfProServerURLV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}

// GetHistoryV1 retrieves the Jamf Pro server URL settings history.
// URL: GET /api/v1/jamf-pro-server-url/history

// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url-history
func (s *JamfProServerUrl) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProJamfProServerURLV1 + "/history"

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
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
		return nil, resp, fmt.Errorf("failed to get Jamf Pro server URL history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// CreateHistoryNoteV1 adds a note to the Jamf Pro server URL settings history.
// URL: POST /api/v1/jamf-pro-server-url/history
// Body: JSON with note
// Returns 201 Created with the created history object
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-pro-server-url-history
func (s *JamfProServerUrl) CreateHistoryNoteV1(ctx context.Context, req *CreateHistoryNoteRequest) (*HistoryObject, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := constants.EndpointJamfProJamfProServerURLV1 + "/history"

	var result HistoryObject

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create Jamf Pro server URL history note: %w", err)
	}

	return &result, resp, nil
}
