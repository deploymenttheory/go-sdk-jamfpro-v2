package distribution_point

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// DistributionPointServiceInterface defines the interface for Distribution Point operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points
	DistributionPointServiceInterface interface {
		// ListV1 retrieves all distribution points with pagination and RSQL filtering.
		//
		// Supports optional RSQL filtering, pagination and sorting via rsqlQuery
		// (keys: filter, sort, page, page-size).
		// Fields allowed in filter: name, serverName, principal, fileSharingConnectionType, httpsEnabled.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// CreateV1 creates a new distribution point.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points
		CreateV1(ctx context.Context, request *RequestDistributionPoint) (*CreateResponse, *resty.Response, error)

		// DeleteMultipleV1 deletes multiple distribution points at once.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-delete-multiple
		DeleteMultipleV1(ctx context.Context, ids []string) (*resty.Response, error)

		// GetByIDV1 retrieves a single distribution point by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id
		GetByIDV1(ctx context.Context, id string) (*ResourceDistributionPoint, *resty.Response, error)

		// UpdateByIDV1 updates the specified distribution point by ID (full update with PUT).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-distribution-points-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestDistributionPoint) (*ResourceDistributionPoint, *resty.Response, error)

		// DeleteByIDV1 removes the specified distribution point by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-distribution-points-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// PatchByIDV1 partially updates the specified distribution point by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-distribution-points-id
		PatchByIDV1(ctx context.Context, id string, request *RequestDistributionPoint) (*ResourceDistributionPoint, *resty.Response, error)

		// GetHistoryByIDV1 retrieves the history for a distribution point with pagination.
		//
		// Supports optional RSQL filtering, pagination and sorting via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id-history
		GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error)

		// CreateHistoryNoteV1 adds a history note to a distribution point.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-id-history
		CreateHistoryNoteV1(ctx context.Context, id string, note string) (*HistoryEntry, *resty.Response, error)
	}

	// Service handles communication with the Distribution Point-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DistributionPointServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Distribution Point Operations (V1)
// -----------------------------------------------------------------------------

// ListV1 retrieves all distribution points with pagination and RSQL filtering.
// URL: GET /api/v1/distribution-points
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointV1

	mergePage := func(pageData []byte) error {
		var items []ResourceDistributionPoint
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
		return nil, resp, fmt.Errorf("failed to list distribution points: %w", err)
	}

	// Set totalCount to the actual number of results retrieved
	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// CreateV1 creates a new distribution point.
// URL: POST /api/v1/distribution-points
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points
func (s *Service) CreateV1(ctx context.Context, request *RequestDistributionPoint) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create distribution point: %w", err)
	}

	return &result, resp, nil
}

// DeleteMultipleV1 deletes multiple distribution points at once.
// URL: POST /api/v1/distribution-points/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-delete-multiple
func (s *Service) DeleteMultipleV1(ctx context.Context, ids []string) (*resty.Response, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("at least one ID is required")
	}

	endpoint := EndpointV1DeleteMultiple

	request := DeleteMultipleRequest{IDs: ids}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete multiple distribution points: %w", err)
	}

	return resp, nil
}

// GetByIDV1 retrieves a single distribution point by ID.
// URL: GET /api/v1/distribution-points/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceDistributionPoint, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("distribution point ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointV1, id)

	var result ResourceDistributionPoint

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get distribution point by ID: %w", err)
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified distribution point by ID (full update with PUT).
// URL: PUT /api/v1/distribution-points/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-distribution-points-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestDistributionPoint) (*ResourceDistributionPoint, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("distribution point ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointV1, id)

	var result ResourceDistributionPoint

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update distribution point: %w", err)
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified distribution point by ID.
// URL: DELETE /api/v1/distribution-points/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-distribution-points-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("distribution point ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to delete distribution point: %w", err)
	}

	return resp, nil
}

// PatchByIDV1 partially updates the specified distribution point by ID.
// URL: PATCH /api/v1/distribution-points/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-distribution-points-id
func (s *Service) PatchByIDV1(ctx context.Context, id string, request *RequestDistributionPoint) (*ResourceDistributionPoint, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("distribution point ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointV1, id)

	var result ResourceDistributionPoint

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to patch distribution point: %w", err)
	}

	return &result, resp, nil
}

// GetHistoryByIDV1 retrieves the history for a distribution point with pagination.
// URL: GET /api/v1/distribution-points/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-distribution-points-id-history
func (s *Service) GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("distribution point ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointV1, id)

	var result HistoryListResponse

	mergePage := func(pageData []byte) error {
		var items []HistoryEntry
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
		return nil, resp, fmt.Errorf("failed to get distribution point history: %w", err)
	}

	// Set totalCount to the actual number of results retrieved
	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// CreateHistoryNoteV1 adds a history note to a distribution point.
// URL: POST /api/v1/distribution-points/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-distribution-points-id-history
func (s *Service) CreateHistoryNoteV1(ctx context.Context, id string, note string) (*HistoryEntry, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("distribution point ID is required")
	}

	if note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointV1, id)

	request := CreateHistoryNoteRequest{Note: note}

	var result HistoryEntry

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create distribution point history note: %w", err)
	}

	return &result, resp, nil
}
