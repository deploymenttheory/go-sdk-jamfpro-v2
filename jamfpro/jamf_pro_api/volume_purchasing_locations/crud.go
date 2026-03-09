package volume_purchasing_locations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// VolumePurchasingLocationsServiceInterface defines the interface for volume purchasing location operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
	VolumePurchasingLocationsServiceInterface interface {
		// ListV1 returns all volume purchasing location objects (Get Volume Purchasing Location objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified volume purchasing location by ID (Get specified Volume Purchasing Location object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id
		GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingLocation, *resty.Response, error)

		// CreateV1 creates a new volume purchasing location (Create Volume Purchasing Location record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations
		CreateV1(ctx context.Context, request *RequestVolumePurchasingLocation) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified volume purchasing location by ID (Update specified Volume Purchasing Location object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-volume-purchasing-locations-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingLocation) (*ResourceVolumePurchasingLocation, *resty.Response, error)

		// DeleteByIDV1 removes the specified volume purchasing location by ID (Remove specified Volume Purchasing Location record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-locations-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// ReclaimVolumePurchasingLocationByIDV1 reclaims the specified volume purchasing location by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-reclaim
		ReclaimVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// GetContentV1 returns the content for the specified volume purchasing location.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id-content
		GetContentV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ContentListResponse, *resty.Response, error)

		// GetHistoryV1 returns the history for the specified volume purchasing location.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id-history
		GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error)

		// AddHistoryNotesV1 adds history notes to the specified volume purchasing location.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-history
		AddHistoryNotesV1(ctx context.Context, id string, request *AddHistoryNotesRequest) (*resty.Response, error)

		// RevokeVolumePurchasingLocationLicensesByIDV1 revokes licenses for the specified volume purchasing location.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-revoke-licenses
		RevokeVolumePurchasingLocationLicensesByIDV1(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the volume purchasing locations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
	VolumePurchasingLocations struct {
		client interfaces.HTTPClient
	}
)

var _ VolumePurchasingLocationsServiceInterface = (*VolumePurchasingLocations)(nil)

func NewVolumePurchasingLocations(client interfaces.HTTPClient) *VolumePurchasingLocations {
	return &VolumePurchasingLocations{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Volume Purchasing Locations CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all volume purchasing location objects.
// URL: GET /api/v1/volume-purchasing-locations
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
func (s *VolumePurchasingLocations) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointVolumePurchasingLocationsV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceVolumePurchasingLocation
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list volume purchasing locations: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified volume purchasing location by ID.
// URL: GET /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id
func (s *VolumePurchasingLocations) GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingLocation, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new volume purchasing location.
// URL: POST /api/v1/volume-purchasing-locations
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations
func (s *VolumePurchasingLocations) CreateV1(ctx context.Context, request *RequestVolumePurchasingLocation) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointVolumePurchasingLocationsV1

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

// UpdateByIDV1 updates the specified volume purchasing location by ID (PATCH).
// URL: PATCH /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-volume-purchasing-locations-id
func (s *VolumePurchasingLocations) UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingLocation) (*ResourceVolumePurchasingLocation, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified volume purchasing location by ID.
// URL: DELETE /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-locations-id
func (s *VolumePurchasingLocations) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ReclaimVolumePurchasingLocationByIDV1 reclaims the specified volume purchasing location by ID.
// URL: POST /api/v1/volume-purchasing-locations/{id}/reclaim
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-reclaim
func (s *VolumePurchasingLocations) ReclaimVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/reclaim", EndpointVolumePurchasingLocationsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetContentV1 returns the content for the specified volume purchasing location.
// URL: GET /api/v1/volume-purchasing-locations/{id}/content
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id-content
func (s *VolumePurchasingLocations) GetContentV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ContentListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/content", EndpointVolumePurchasingLocationsV1, id)

	var result ContentListResponse
	result.Results = []VolumePurchasingSubsetContent{}

	mergePage := func(pageData []byte) error {
		var pageItems []VolumePurchasingSubsetContent
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get volume purchasing location content: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetHistoryV1 returns the history for the specified volume purchasing location.
// URL: GET /api/v1/volume-purchasing-locations/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id-history
func (s *VolumePurchasingLocations) GetHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointVolumePurchasingLocationsV1, id)

	var result HistoryListResponse
	result.Results = []HistoryEntry{}

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryEntry
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get volume purchasing location history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotesV1 adds history notes to the specified volume purchasing location.
// URL: POST /api/v1/volume-purchasing-locations/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-history
func (s *VolumePurchasingLocations) AddHistoryNotesV1(ctx context.Context, id string, request *AddHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing location ID is required")
	}

	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointVolumePurchasingLocationsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RevokeVolumePurchasingLocationLicensesByIDV1 revokes licenses for the specified volume purchasing location.
// URL: POST /api/v1/volume-purchasing-locations/{id}/revoke-licenses
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-revoke-licenses
func (s *VolumePurchasingLocations) RevokeVolumePurchasingLocationLicensesByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/revoke-licenses", EndpointVolumePurchasingLocationsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
