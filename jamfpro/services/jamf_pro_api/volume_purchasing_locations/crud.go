package volume_purchasing_locations

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// VolumePurchasingLocationsServiceInterface defines the interface for volume purchasing location operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
	VolumePurchasingLocationsServiceInterface interface {
		// ListVolumePurchasingLocationsV1 returns all volume purchasing location objects (Get Volume Purchasing Location objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
		ListVolumePurchasingLocationsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetVolumePurchasingLocationByIDV1 returns the specified volume purchasing location by ID (Get specified Volume Purchasing Location object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id
		GetVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingLocation, *interfaces.Response, error)

		// CreateVolumePurchasingLocationV1 creates a new volume purchasing location (Create Volume Purchasing Location record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations
		CreateVolumePurchasingLocationV1(ctx context.Context, req *RequestVolumePurchasingLocation) (*CreateResponse, *interfaces.Response, error)

		// UpdateVolumePurchasingLocationByIDV1 updates the specified volume purchasing location by ID (Update specified Volume Purchasing Location object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-volume-purchasing-locations-id
		UpdateVolumePurchasingLocationByIDV1(ctx context.Context, id string, req *RequestVolumePurchasingLocation) (*ResourceVolumePurchasingLocation, *interfaces.Response, error)

		// DeleteVolumePurchasingLocationByIDV1 removes the specified volume purchasing location by ID (Remove specified Volume Purchasing Location record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-locations-id
		DeleteVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// ReclaimVolumePurchasingLocationByIDV1 reclaims the specified volume purchasing location by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-reclaim
		ReclaimVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// GetContentV1 returns the content for the specified volume purchasing location.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id-content
		GetContentV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ContentListResponse, *interfaces.Response, error)
	}

	// Service handles communication with the volume purchasing locations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ VolumePurchasingLocationsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Volume Purchasing Locations CRUD Operations
// -----------------------------------------------------------------------------

// ListVolumePurchasingLocationsV1 returns all volume purchasing location objects.
// URL: GET /api/v1/volume-purchasing-locations
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations
func (s *Service) ListVolumePurchasingLocationsV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointVolumePurchasingLocationsV1

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

// GetVolumePurchasingLocationByIDV1 returns the specified volume purchasing location by ID.
// URL: GET /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-locations-id
func (s *Service) GetVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingLocation, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

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

// CreateVolumePurchasingLocationV1 creates a new volume purchasing location.
// URL: POST /api/v1/volume-purchasing-locations
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations
func (s *Service) CreateVolumePurchasingLocationV1(ctx context.Context, req *RequestVolumePurchasingLocation) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointVolumePurchasingLocationsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateVolumePurchasingLocationByIDV1 updates the specified volume purchasing location by ID (PATCH).
// URL: PATCH /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-volume-purchasing-locations-id
func (s *Service) UpdateVolumePurchasingLocationByIDV1(ctx context.Context, id string, req *RequestVolumePurchasingLocation) (*ResourceVolumePurchasingLocation, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

	var result ResourceVolumePurchasingLocation

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteVolumePurchasingLocationByIDV1 removes the specified volume purchasing location by ID.
// URL: DELETE /api/v1/volume-purchasing-locations/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-locations-id
func (s *Service) DeleteVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingLocationsV1, id)

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

// ReclaimVolumePurchasingLocationByIDV1 reclaims the specified volume purchasing location by ID.
// URL: POST /api/v1/volume-purchasing-locations/{id}/reclaim
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-locations-id-reclaim
func (s *Service) ReclaimVolumePurchasingLocationByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
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
func (s *Service) GetContentV1(ctx context.Context, id string, rsqlQuery map[string]string) (*ContentListResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing location ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/content", EndpointVolumePurchasingLocationsV1, id)

	var result ContentListResponse

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
