package buildings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// BuildingsServiceInterface defines the interface for building operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
	BuildingsServiceInterface interface {
		// ListV1 returns all building objects (Get Building objects).
		//
		// Returns a paged list of building objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified building by ID (Get specified Building object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id
		GetByIDV1(ctx context.Context, id string) (*ResourceBuilding, *interfaces.Response, error)

		// CreateV1 creates a new building record (Create Building record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-buildings
		CreateV1(ctx context.Context, request *RequestBuilding) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified building by ID (Update specified Building object).
		//
		// Returns the full updated building resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-buildings-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestBuilding) (*ResourceBuilding, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified building by ID (Remove specified Building record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-buildings-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteBuildingsByIDV1 deletes multiple buildings by their IDs (Delete multiple Buildings by their IDs).
		//
		// Sends a POST to /api/v1/buildings/delete-multiple with a body containing building IDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-delete-multiple
		DeleteBuildingsByIDV1(ctx context.Context, req *DeleteBuildingsByIDRequest) (*interfaces.Response, error)

		// GetBuildingHistoryV1 returns the history object for the specified building.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id-history
		GetBuildingHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddBuildingHistoryNotesV1 adds notes to the specified building history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-id-history
		AddBuildingHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error)

		// ExportV1 exports the buildings collection in the specified format (JSON or CSV).
		//
		// POST /api/v1/buildings/export. Optional query params: page, page-size, sort, filter,
		// export-fields, export-labels. Optional body overrides when URI exceeds ~2k chars.
		// acceptType should be mime.TextCSV or mime.ApplicationJSON.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-export
		ExportV1(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error)

		// ExportHistoryV1 exports the history for the specified building in the specified format (JSON or CSV).
		//
		// POST /api/v1/buildings/{id}/history/export. Optional query params: page, page-size, sort,
		// filter, export-fields, export-labels. Optional body overrides when URI exceeds ~2k chars.
		// acceptType should be mime.TextCSV or mime.ApplicationJSON.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-id-history-export
		ExportHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the buildings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ BuildingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Buildings CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all building objects (Get Building objects).
// URL: GET /api/v1/buildings
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointBuildingsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceBuilding
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list buildings: %w", err)
	}
	
	// Set totalCount to the actual number of results retrieved
	result.TotalCount = len(result.Results)
	
	return &result, resp, nil
}

// GetByIDV1 returns the specified building by ID (Get specified Building object).
// URL: GET /api/v1/buildings/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceBuilding, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointBuildingsV1, id)

	var result ResourceBuilding

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new building record (Create Building record).
// URL: POST /api/v1/buildings
// Body: JSON with name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings
func (s *Service) CreateV1(ctx context.Context, request *RequestBuilding) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointBuildingsV1

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

// UpdateByIDV1 updates the specified building by ID (Update specified Building object).
// URL: PUT /api/v1/buildings/{id}
// Body: JSON with name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country
// https://developer.jamf.com/jamf-pro/reference/put_v1-buildings-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestBuilding) (*ResourceBuilding, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointBuildingsV1, id)

	var result ResourceBuilding

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

// DeleteByIDV1 removes the specified building by ID (Remove specified Building record).
// URL: DELETE /api/v1/buildings/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-buildings-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointBuildingsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteBuildingsByIDV1 deletes multiple buildings by their IDs (Delete multiple Buildings by their IDs).
// URL: POST /api/v1/buildings/delete-multiple
// Body: JSON with ids (array of building IDs)
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-delete-multiple
func (s *Service) DeleteBuildingsByIDV1(ctx context.Context, req *DeleteBuildingsByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointBuildingsV1 + "/delete-multiple"

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

// GetBuildingHistoryV1 returns the history object for the specified building.
// URL: GET /api/v1/buildings/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id-history
func (s *Service) GetBuildingHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointBuildingsV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var items []HistoryObject
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get building history: %w", err)
	}
	
	// Set totalCount to the actual number of results retrieved
	result.TotalCount = len(result.Results)
	
	return &result, resp, nil
}

// AddBuildingHistoryNotesV1 adds notes to the specified building history.
// URL: POST /api/v1/buildings/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-id-history
func (s *Service) AddBuildingHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("building ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointBuildingsV1, id)

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

// ExportV1 exports the buildings collection in the specified format (JSON or CSV).
// URL: POST /api/v1/buildings/export
// Query Params: page, page-size, sort, filter, export-fields, export-labels (optional)
// Body: optional ExportRequest to override query params when URI exceeds ~2k chars
// Accept: text/csv or application/json
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-export
func (s *Service) ExportV1(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error) {
	endpoint := EndpointBuildingsV1 + "/export"

	if acceptType == "" {
		acceptType = mime.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": mime.ApplicationJSON,
	}

	var body any
	if req != nil {
		body = req
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export buildings: %w", err)
	}

	return resp.Body, resp, nil
}

// ExportHistoryV1 exports the history for the specified building in the specified format (JSON or CSV).
// URL: POST /api/v1/buildings/{id}/history/export
// Path param: id (building ID)
// Query Params: page, page-size, sort, filter, export-fields, export-labels (optional)
// Body: optional ExportRequest to override query params when URI exceeds ~2k chars
// Accept: text/csv or application/json
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-id-history-export
func (s *Service) ExportHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history/export", EndpointBuildingsV1, id)

	if acceptType == "" {
		acceptType = mime.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": mime.ApplicationJSON,
	}

	var body any
	if req != nil {
		body = req
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export building history: %w", err)
	}

	return resp.Body, resp, nil
}
