package jamf_pro_server_url

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfProServerURLServiceInterface defines the interface for Jamf Pro server URL operations.
	//
	// Manages the Jamf Pro server URL and unsecured enrollment URL settings.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
	JamfProServerURLServiceInterface interface {
		// GetV1 retrieves the Jamf Pro server URL settings.
		//
		// Returns the configured Jamf Pro server URL and unsecured enrollment URL.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
		GetV1(ctx context.Context) (*ResourceJamfProServerURL, *interfaces.Response, error)

		// UpdateV1 updates the Jamf Pro server URL settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-pro-server-url
		UpdateV1(ctx context.Context, request *ResourceJamfProServerURL) (*ResourceJamfProServerURL, *interfaces.Response, error)

		// GetHistoryV1 retrieves the Jamf Pro server URL settings history.
		//
		// GET /api/v1/jamf-pro-server-url/history. Query params: page, page-size, sort.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// CreateHistoryNoteV1 adds a note to the Jamf Pro server URL settings history.
		//
		// POST /api/v1/jamf-pro-server-url/history. Body: {"note": "string"}. Returns 201 Created.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-pro-server-url-history
		CreateHistoryNoteV1(ctx context.Context, req *CreateHistoryNoteRequest) (*HistoryObject, *interfaces.Response, error)
	}

	// Service handles communication with the Jamf Pro server URL-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfProServerURLServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the Jamf Pro server URL settings.
// URL: GET /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
func (s *Service) GetV1(ctx context.Context) (*ResourceJamfProServerURL, *interfaces.Response, error) {
	var result ResourceJamfProServerURL

	endpoint := EndpointJamfProServerURLV1
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}

// UpdateV1 updates the Jamf Pro server URL settings.
// URL: PUT /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-pro-server-url
func (s *Service) UpdateV1(ctx context.Context, request *ResourceJamfProServerURL) (*ResourceJamfProServerURL, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceJamfProServerURL

	endpoint := EndpointJamfProServerURLV1
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}

// GetHistoryV1 retrieves the Jamf Pro server URL settings history.
// URL: GET /api/v1/jamf-pro-server-url/history

// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url-history
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := EndpointJamfProServerURLV1 + "/history"

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageResponse HistoryResponse
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
		return nil, resp, fmt.Errorf("failed to get Jamf Pro server URL history: %w", err)
	}
	return &result, resp, nil
}

// CreateHistoryNoteV1 adds a note to the Jamf Pro server URL settings history.
// URL: POST /api/v1/jamf-pro-server-url/history
// Body: JSON with note
// Returns 201 Created with the created history object
// https://developer.jamf.com/jamf-pro/reference/post_v1-jamf-pro-server-url-history
func (s *Service) CreateHistoryNoteV1(ctx context.Context, req *CreateHistoryNoteRequest) (*HistoryObject, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := EndpointJamfProServerURLV1 + "/history"

	var result HistoryObject
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create Jamf Pro server URL history note: %w", err)
	}

	return &result, resp, nil
}
