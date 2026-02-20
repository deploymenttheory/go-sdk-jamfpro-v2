package scripts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ScriptsServiceInterface defines the interface for script operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
	ScriptsServiceInterface interface {
		// ListScriptsV1 returns a paged list of script objects.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
		ListScriptsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetScriptByIDV1 returns the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
		GetScriptByIDV1(ctx context.Context, id string) (*ResourceScript, *interfaces.Response, error)

		// CreateScriptV1 creates a new script record.
		//
		// Returns the created script's ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts
		CreateScriptV1(ctx context.Context, request *RequestScript) (*CreateResponse, *interfaces.Response, error)

		// UpdateScriptByIDV1 replaces the specified script by ID.
		//
		// Returns the full updated script resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-scripts-id
		UpdateScriptByIDV1(ctx context.Context, id string, request *RequestScript) (*ResourceScript, *interfaces.Response, error)

		// DeleteScriptByIDV1 removes the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-scripts-id
		DeleteScriptByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DownloadScriptByIDV1 downloads the script contents as plain text.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-download
		DownloadScriptByIDV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error)

		// GetScriptHistoryV1 returns the history object for the specified script.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-history
		GetScriptHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *interfaces.Response, error)

		// AddScriptHistoryNotesV1 adds notes to the specified script's history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts-id-history
		AddScriptHistoryNotesV1(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*interfaces.Response, error)
	}

	// Service handles communication with the scripts-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ScriptsServiceInterface = (*Service)(nil)

// NewService returns a new scripts Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Scripts CRUD Operations
// -----------------------------------------------------------------------------

// ListScriptsV1 returns a paged list of script objects.
// URL: GET /api/v1/scripts
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
func (s *Service) ListScriptsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointScriptsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetScriptByIDV1 returns the specified script by ID.
// URL: GET /api/v1/scripts/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
func (s *Service) GetScriptByIDV1(ctx context.Context, id string) (*ResourceScript, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	var result ResourceScript

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

// CreateScriptV1 creates a new script record.
// URL: POST /api/v1/scripts
// https://developer.jamf.com/jamf-pro/reference/post_v1-scripts
func (s *Service) CreateScriptV1(ctx context.Context, request *RequestScript) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointScriptsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateScriptByIDV1 replaces the specified script by ID.
// URL: PUT /api/v1/scripts/{id}
// Returns the full updated script resource.
// https://developer.jamf.com/jamf-pro/reference/put_v1-scripts-id
func (s *Service) UpdateScriptByIDV1(ctx context.Context, id string, request *RequestScript) (*ResourceScript, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	var result ResourceScript

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

// DeleteScriptByIDV1 removes the specified script by ID.
// URL: DELETE /api/v1/scripts/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-scripts-id
func (s *Service) DeleteScriptByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DownloadScriptByIDV1 downloads the script contents as plain text.
// URL: GET /api/v1/scripts/{id}/download
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-download
func (s *Service) DownloadScriptByIDV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download", EndpointScriptsV1, id)

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, map[string]string{"Accept": "text/plain"})
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// GetScriptHistoryV1 returns the history object for the specified script.
// URL: GET /api/v1/scripts/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-history
func (s *Service) GetScriptHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointScriptsV1, id)

	var result ScriptHistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddScriptHistoryNotesV1 adds notes to the specified script's history.
// URL: POST /api/v1/scripts/{id}/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-scripts-id-history
func (s *Service) AddScriptHistoryNotesV1(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
