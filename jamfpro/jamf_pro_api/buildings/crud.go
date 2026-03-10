package buildings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the buildings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
	Buildings struct {
		client client.Client
	}
)

func NewBuildings(client client.Client) *Buildings {
	return &Buildings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Buildings CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all building objects (Get Building objects).
// URL: GET /api/v1/buildings
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings
func (s *Buildings) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProBuildingsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceBuilding
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Buildings) GetByIDV1(ctx context.Context, id string) (*ResourceBuilding, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBuildingsV1, id)

	var result ResourceBuilding

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Buildings) CreateV1(ctx context.Context, request *RequestBuilding) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProBuildingsV1

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

// UpdateByIDV1 updates the specified building by ID (Update specified Building object).
// URL: PUT /api/v1/buildings/{id}
// Body: JSON with name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country
// https://developer.jamf.com/jamf-pro/reference/put_v1-buildings-id
func (s *Buildings) UpdateByIDV1(ctx context.Context, id string, request *RequestBuilding) (*ResourceBuilding, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBuildingsV1, id)

	var result ResourceBuilding

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

// DeleteByIDV1 removes the specified building by ID (Remove specified Building record).
// URL: DELETE /api/v1/buildings/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-buildings-id
func (s *Buildings) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBuildingsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *Buildings) DeleteBuildingsByIDV1(ctx context.Context, req *DeleteBuildingsByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProBuildingsV1 + "/delete-multiple"

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

// GetBuildingHistoryV1 returns the history object for the specified building.
// URL: GET /api/v1/buildings/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id-history
func (s *Buildings) GetBuildingHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProBuildingsV1, id)

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
		"Accept": constants.ApplicationJSON,
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
func (s *Buildings) AddBuildingHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("building ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProBuildingsV1, id)

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

// ExportV1 exports the buildings collection in the specified format (JSON or CSV).
// URL: POST /api/v1/buildings/export
// Query Params: page, page-size, sort, filter, export-fields, export-labels (optional)
// Body: optional ExportRequest to override query params when URI exceeds ~2k chars
// Accept: text/csv or application/json
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-export
func (s *Buildings) ExportV1(ctx context.Context, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProBuildingsV1 + "/export"

	if acceptType == "" {
		acceptType = constants.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": constants.ApplicationJSON,
	}

	var body any
	if req != nil {
		body = req
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export buildings: %w", err)
	}

	return resp.Bytes(), resp, nil
}

// ExportHistoryV1 exports the history for the specified building in the specified format (JSON or CSV).
// URL: POST /api/v1/buildings/{id}/history/export
// Path param: id (building ID)
// Query Params: page, page-size, sort, filter, export-fields, export-labels (optional)
// Body: optional ExportRequest to override query params when URI exceeds ~2k chars
// Accept: text/csv or application/json
// https://developer.jamf.com/jamf-pro/reference/post_v1-buildings-id-history-export
func (s *Buildings) ExportHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string, req *ExportRequest, acceptType string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("building ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history/export", constants.EndpointJamfProBuildingsV1, id)

	if acceptType == "" {
		acceptType = constants.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptType,
		"Content-Type": constants.ApplicationJSON,
	}

	var body any
	if req != nil {
		body = req
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to export building history: %w", err)
	}

	return resp.Bytes(), resp, nil
}
