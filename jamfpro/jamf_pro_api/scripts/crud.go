package scripts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
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
		ListScriptsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetScriptByIDV1 returns the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
		GetScriptByIDV1(ctx context.Context, id string) (*ResourceScript, *resty.Response, error)

		// CreateScriptV1 creates a new script record.
		//
		// Returns the created script's ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts
		CreateScriptV1(ctx context.Context, request *RequestScript) (*CreateResponse, *resty.Response, error)

		// UpdateScriptByIDV1 replaces the specified script by ID.
		//
		// Returns the full updated script resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-scripts-id
		UpdateScriptByIDV1(ctx context.Context, id string, request *RequestScript) (*ResourceScript, *resty.Response, error)

		// DeleteScriptByIDV1 removes the specified script by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-scripts-id
		DeleteScriptByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// DownloadScriptByIDV1 downloads the script contents as plain text.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-download
		DownloadScriptByIDV1(ctx context.Context, id string) ([]byte, *resty.Response, error)

		// GetScriptHistoryV1 returns the history object for the specified script.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id-history
		GetScriptHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *resty.Response, error)

		// AddScriptHistoryNotesV1 adds notes to the specified script's history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-scripts-id-history
		AddScriptHistoryNotesV1(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*resty.Response, error)
	}

	// Service handles communication with the scripts-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
	Scripts struct {
		client transport.HTTPClient
	}
)

var _ ScriptsServiceInterface = (*Scripts)(nil)

// NewService returns a new scripts Service backed by the provided HTTP client.
func NewScripts(client transport.HTTPClient) *Scripts {
	return &Scripts{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Scripts CRUD Operations
// -----------------------------------------------------------------------------

// ListScriptsV1 returns a paged list of script objects.
// URL: GET /api/v1/scripts
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts
func (s *Scripts) ListScriptsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProScriptsV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceScript
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
		return nil, resp, fmt.Errorf("failed to list scripts: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetScriptByIDV1 returns the specified script by ID.
// URL: GET /api/v1/scripts/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-scripts-id
func (s *Scripts) GetScriptByIDV1(ctx context.Context, id string) (*ResourceScript, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProScriptsV1, id)

	var result ResourceScript

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Scripts) CreateScriptV1(ctx context.Context, request *RequestScript) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProScriptsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *Scripts) UpdateScriptByIDV1(ctx context.Context, id string, request *RequestScript) (*ResourceScript, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProScriptsV1, id)

	var result ResourceScript

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *Scripts) DeleteScriptByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProScriptsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Scripts) DownloadScriptByIDV1(ctx context.Context, id string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/download", constants.EndpointJamfProScriptsV1, id)

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
func (s *Scripts) GetScriptHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ScriptHistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("script ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProScriptsV1, id)

	var result ScriptHistoryResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Scripts) AddScriptHistoryNotesV1(ctx context.Context, id string, req *AddScriptHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("script ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProScriptsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
