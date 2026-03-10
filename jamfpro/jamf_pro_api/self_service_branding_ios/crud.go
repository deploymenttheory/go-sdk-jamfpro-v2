package self_service_branding_ios

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the self-service branding mobile-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-ios
	SelfServiceBrandingIos struct {
		client client.Client
	}
)

func NewSelfServiceBrandingIos(client client.Client) *SelfServiceBrandingIos {
	return &SelfServiceBrandingIos{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self-Service Branding Mobile (iOS) CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all self-service branding mobile configurations.
// URL: GET /api/v1/self-service/branding/ios

// https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-ios
func (s *SelfServiceBrandingIos) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceSelfServiceBrandingMobile
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProSelfServiceBrandingMobileV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified self-service branding mobile configuration by ID.
// URL: GET /api/v1/self-service/branding/ios/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-ios-id
func (s *SelfServiceBrandingIos) GetByIDV1(ctx context.Context, id string) (*ResourceSelfServiceBrandingMobile, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("self-service branding mobile ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMobileV1, id)

	var result ResourceSelfServiceBrandingMobile

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV1 returns the specified self-service branding mobile configuration by name.
func (s *SelfServiceBrandingIos) GetByNameV1(ctx context.Context, name string) (*ResourceSelfServiceBrandingMobile, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("self-service branding mobile name is required")
	}

	list, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for i := range list.Results {
		if list.Results[i].BrandingName == name {
			return &list.Results[i], resp, nil
		}
	}

	return nil, resp, fmt.Errorf("self-service branding mobile with name %q was not found", name)
}

// CreateV1 creates a new self-service branding mobile configuration.
// URL: POST /api/v1/self-service/branding/ios
// https://developer.jamf.com/jamf-pro/reference/post_v1-self-service-branding-ios
func (s *SelfServiceBrandingIos) CreateV1(ctx context.Context, request *ResourceSelfServiceBrandingMobile) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProSelfServiceBrandingMobileV1

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

// UpdateByIDV1 updates the specified self-service branding mobile configuration by ID.
// URL: PUT /api/v1/self-service/branding/ios/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-self-service-branding-ios-id
func (s *SelfServiceBrandingIos) UpdateByIDV1(ctx context.Context, id string, request *ResourceSelfServiceBrandingMobile) (*ResourceSelfServiceBrandingMobile, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMobileV1, id)

	var result ResourceSelfServiceBrandingMobile

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

// UpdateByNameV1 updates a self-service branding mobile configuration by name.
func (s *SelfServiceBrandingIos) UpdateByNameV1(ctx context.Context, name string, request *ResourceSelfServiceBrandingMobile) (*ResourceSelfServiceBrandingMobile, *resty.Response, error) {
	target, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return nil, resp, err
	}

	return s.UpdateByIDV1(ctx, target.ID, request)
}

// DeleteByIDV1 removes the specified self-service branding mobile configuration by ID.
// URL: DELETE /api/v1/self-service/branding/ios/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-self-service-branding-ios-id
func (s *SelfServiceBrandingIos) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("self-service branding mobile ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSelfServiceBrandingMobileV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByNameV1 removes a self-service branding mobile configuration by name.
func (s *SelfServiceBrandingIos) DeleteByNameV1(ctx context.Context, name string) (*resty.Response, error) {
	target, resp, err := s.GetByNameV1(ctx, name)
	if err != nil {
		return resp, err
	}

	return s.DeleteByIDV1(ctx, target.ID)
}
