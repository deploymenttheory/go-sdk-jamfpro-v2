package buildings

import (
	"context"
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
