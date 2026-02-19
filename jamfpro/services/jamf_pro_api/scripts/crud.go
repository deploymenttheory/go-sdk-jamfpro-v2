package scripts

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
)

type (
	// ScriptsServiceInterface defines the interface for script operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
	ScriptsServiceInterface interface {
		// ListScripts returns a paged list of script objects.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
		ListScripts(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetScriptByID returns the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
		GetScriptByID(ctx context.Context, id string) (*ResourceScript, *interfaces.Response, error)

		// CreateScript creates a new script record.
		//
		// Returns the created script's ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts
		CreateScript(ctx context.Context, req *RequestScript) (*CreateResponse, *interfaces.Response, error)

		// UpdateScriptByID replaces the specified script by ID.
		//
		// Returns the full updated script resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-scripts-id
		UpdateScriptByID(ctx context.Context, id string, req *RequestScript) (*ResourceScript, *interfaces.Response, error)

		// DeleteScriptByID removes the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-scripts-id
		DeleteScriptByID(ctx context.Context, id string) (*interfaces.Response, error)

		// DownloadScriptByID downloads the script contents as plain text.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-download
		DownloadScriptByID(ctx context.Context, id string) ([]byte, *interfaces.Response, error)

		// GetScriptHistory returns the history object for the specified script.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-history
		GetScriptHistory(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *interfaces.Response, error)

		// AddScriptHistoryNotes adds notes to the specified script's history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts-id-history
		AddScriptHistoryNotes(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*interfaces.Response, error)
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

// ListScripts returns a paged list of script objects.
// URL: GET /api/v1/scripts
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
func (s *Service) ListScripts(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointScriptsV1, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetScriptByID returns the specified script by ID.
// URL: GET /api/v1/scripts/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
func (s *Service) GetScriptByID(ctx context.Context, id string) (*ResourceScript, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ResourceScript

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateScript creates a new script record.
// URL: POST /api/v1/scripts
// https://developer.jamf.com/jamf-pro/reference/post_v1-scripts
func (s *Service) CreateScript(ctx context.Context, req *RequestScript) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result CreateResponse

	resp, err := s.client.Post(ctx, EndpointScriptsV1, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateScriptByID replaces the specified script by ID.
// URL: PUT /api/v1/scripts/{id}
// Returns the full updated script resource.
// https://developer.jamf.com/jamf-pro/reference/put_v1-scripts-id
func (s *Service) UpdateScriptByID(ctx context.Context, id string, req *RequestScript) (*ResourceScript, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ResourceScript

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteScriptByID removes the specified script by ID.
// URL: DELETE /api/v1/scripts/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-scripts-id
func (s *Service) DeleteScriptByID(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DownloadScriptByID downloads the script contents as plain text.
// URL: GET /api/v1/scripts/{id}/download
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-download
func (s *Service) DownloadScriptByID(ctx context.Context, id string) ([]byte, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept": "text/plain",
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

// GetScriptHistory returns the history object for the specified script.
// URL: GET /api/v1/scripts/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-history
func (s *Service) GetScriptHistory(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	var result ScriptHistoryResponse

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddScriptHistoryNotes adds notes to the specified script's history.
// URL: POST /api/v1/scripts/{id}/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-scripts-id-history
func (s *Service) AddScriptHistoryNotes(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointScriptsV1, id)

	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
