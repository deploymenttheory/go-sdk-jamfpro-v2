package advanced_mobile_device_searches

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// AdvancedMobileDeviceSearchesServiceInterface defines the interface for advanced mobile device search operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
	AdvancedMobileDeviceSearchesServiceInterface interface {
		// ListV1 returns all advanced mobile device searches (Get Advanced Mobile Device Searches).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
		ListV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified advanced mobile device search by ID (Get Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-id
		GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedMobileDeviceSearch, *interfaces.Response, error)

		// CreateV1 creates a new advanced mobile device search (Create Advanced Mobile Device Search).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches
		CreateV1(ctx context.Context, search *ResourceAdvancedMobileDeviceSearch) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified advanced mobile device search by ID (Update Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-mobile-device-searches-id
		UpdateByIDV1(ctx context.Context, id string, search *ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified advanced mobile device search by ID (Delete Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-mobile-device-searches-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// GetChoicesV1 returns criteria choices for advanced mobile device searches (Get Advanced Mobile Device Search Choices).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-choices
		GetChoicesV1(ctx context.Context, criteria, site, contains string) (*ChoicesResponse, *interfaces.Response, error)
	}

	// Service handles communication with the advanced mobile device searches-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AdvancedMobileDeviceSearchesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all advanced mobile device searches.
// URL: GET /api/v1/advanced-mobile-device-searches
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
func (s *Service) ListV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	resp, err := s.client.Get(ctx, EndpointAdvancedMobileDeviceSearchesV1, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the specified advanced mobile device search by ID.
// URL: GET /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedMobileDeviceSearch, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointAdvancedMobileDeviceSearchesV1, id)
	var result ResourceAdvancedMobileDeviceSearch
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV1 creates a new advanced mobile device search.
// URL: POST /api/v1/advanced-mobile-device-searches
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches
func (s *Service) CreateV1(ctx context.Context, search *ResourceAdvancedMobileDeviceSearch) (*CreateResponse, *interfaces.Response, error) {
	if search == nil {
		return nil, nil, fmt.Errorf("search is required")
	}
	var result CreateResponse
	resp, err := s.client.Post(ctx, EndpointAdvancedMobileDeviceSearchesV1, search, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateByIDV1 updates the specified advanced mobile device search by ID.
// URL: PUT /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-mobile-device-searches-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, search *ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if search == nil {
		return nil, nil, fmt.Errorf("search is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointAdvancedMobileDeviceSearchesV1, id)
	var result ResourceAdvancedMobileDeviceSearch
	resp, err := s.client.Put(ctx, endpoint, search, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteByIDV1 removes the specified advanced mobile device search by ID.
// URL: DELETE /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-mobile-device-searches-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointAdvancedMobileDeviceSearchesV1, id)
	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetChoicesV1 returns criteria choices for advanced mobile device searches.
// URL: GET /api/v1/advanced-mobile-device-searches/choices?criteria=...&site=...&contains=...
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-choices
func (s *Service) GetChoicesV1(ctx context.Context, criteria, site, contains string) (*ChoicesResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/choices?criteria=%s&site=%s&contains=%s",
		EndpointAdvancedMobileDeviceSearchesV1,
		url.QueryEscape(criteria),
		url.QueryEscape(site),
		url.QueryEscape(contains))
	var result ChoicesResponse
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
