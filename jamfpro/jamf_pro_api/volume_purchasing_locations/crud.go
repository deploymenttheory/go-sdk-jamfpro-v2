package volume_purchasing_locations

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the volume purchasing locations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
	VolumePurchasingLocations struct {
		client client.Client
	}
)

func NewVolumePurchasingLocations(client client.Client) *VolumePurchasingLocations {
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

	endpoint := constants.EndpointJamfProVolumePurchasingLocationsV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceVolumePurchasingLocation
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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	endpoint := constants.EndpointJamfProVolumePurchasingLocationsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Patch(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s/reclaim", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s/content", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
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

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
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

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s/revoke-licenses", constants.EndpointJamfProVolumePurchasingLocationsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
