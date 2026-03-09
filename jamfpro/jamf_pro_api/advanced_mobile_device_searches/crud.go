package advanced_mobile_device_searches

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// AdvancedMobileDeviceSearchesServiceInterface defines the interface for advanced mobile device search operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
	AdvancedMobileDeviceSearchesServiceInterface interface {
		// ListV1 returns all advanced mobile device searches (Get Advanced Mobile Device Searches).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified advanced mobile device search by ID (Get Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-id
		GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedMobileDeviceSearch, *resty.Response, error)

		// CreateV1 creates a new advanced mobile device search (Create Advanced Mobile Device Search).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches
		CreateV1(ctx context.Context, request *ResourceAdvancedMobileDeviceSearch) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified advanced mobile device search by ID (Update Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-mobile-device-searches-id
		UpdateByIDV1(ctx context.Context, id string, request *ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, *resty.Response, error)

		// DeleteByIDV1 removes the specified advanced mobile device search by ID (Delete Advanced Mobile Device Search by ID).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-mobile-device-searches-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// DeleteMultipleV1 deletes multiple advanced mobile device searches by their IDs (Delete multiple Advanced Mobile Device Searches by their IDs).
		//
		// Sends a POST to /api/v1/advanced-mobile-device-searches/delete-multiple with a body containing search IDs.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches-delete-multiple
		DeleteMultipleV1(ctx context.Context, req *DeleteAdvancedMobileDeviceSearchesByIDRequest) (*resty.Response, error)

		// GetChoicesV1 returns criteria choices for advanced mobile device searches (Get Advanced Mobile Device Search Choices).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-choices
		GetChoicesV1(ctx context.Context, criteria, site, contains string) (*ChoicesResponse, *resty.Response, error)
	}

	// Service handles communication with the advanced mobile device searches-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
	AdvancedMobileDeviceSearches struct {
		client transport.HTTPClient
	}
)

var _ AdvancedMobileDeviceSearchesServiceInterface = (*AdvancedMobileDeviceSearches)(nil)

func NewAdvancedMobileDeviceSearches(client transport.HTTPClient) *AdvancedMobileDeviceSearches {
	return &AdvancedMobileDeviceSearches{client: client}
}

// ListV1 returns all advanced mobile device searches.
// URL: GET /api/v1/advanced-mobile-device-searches
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
func (s *AdvancedMobileDeviceSearches) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProAdvancedMobileDeviceSearchesV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetByIDV1 returns the specified advanced mobile device search by ID.
// URL: GET /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-id
func (s *AdvancedMobileDeviceSearches) GetByIDV1(ctx context.Context, id string) (*ResourceAdvancedMobileDeviceSearch, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedMobileDeviceSearchesV1, id)
	var result ResourceAdvancedMobileDeviceSearch

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new advanced mobile device search.
// URL: POST /api/v1/advanced-mobile-device-searches
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches
func (s *AdvancedMobileDeviceSearches) CreateV1(ctx context.Context, request *ResourceAdvancedMobileDeviceSearch) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("search is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProAdvancedMobileDeviceSearchesV1

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

// UpdateByIDV1 updates the specified advanced mobile device search by ID.
// URL: PUT /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-advanced-mobile-device-searches-id
func (s *AdvancedMobileDeviceSearches) UpdateByIDV1(ctx context.Context, id string, request *ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("search is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedMobileDeviceSearchesV1, id)

	var result ResourceAdvancedMobileDeviceSearch

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

// DeleteByIDV1 removes the specified advanced mobile device search by ID.
// URL: DELETE /api/v1/advanced-mobile-device-searches/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-advanced-mobile-device-searches-id
func (s *AdvancedMobileDeviceSearches) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAdvancedMobileDeviceSearchesV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteMultipleV1 deletes multiple advanced mobile device searches by their IDs (Delete multiple Advanced Mobile Device Searches by their IDs).
// URL: POST /api/v1/advanced-mobile-device-searches/delete-multiple
// Body: JSON with ids (array of search IDs)
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-advanced-mobile-device-searches-delete-multiple
func (s *AdvancedMobileDeviceSearches) DeleteMultipleV1(ctx context.Context, req *DeleteAdvancedMobileDeviceSearchesByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProAdvancedMobileDeviceSearchesV1 + "/delete-multiple"

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

// GetChoicesV1 returns criteria choices for advanced mobile device searches.
// URL: GET /api/v1/advanced-mobile-device-searches/choices?criteria=...&site=...&contains=...
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches-choices
func (s *AdvancedMobileDeviceSearches) GetChoicesV1(ctx context.Context, criteria, site, contains string) (*ChoicesResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/choices?criteria=%s&site=%s&contains=%s",
		constants.EndpointJamfProAdvancedMobileDeviceSearchesV1,
		url.QueryEscape(criteria),
		url.QueryEscape(site),
		url.QueryEscape(contains))
	var result ChoicesResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
