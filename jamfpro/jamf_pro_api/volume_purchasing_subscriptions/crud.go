package volume_purchasing_subscriptions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// VolumePurchasingSubscriptionsServiceInterface defines the interface for volume purchasing subscription operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
	VolumePurchasingSubscriptionsServiceInterface interface {
		// ListV1 returns all volume purchasing subscription objects (Get Volume Purchasing Subscription objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified volume purchasing subscription by ID (Get specified Volume Purchasing Subscription object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions-id
		GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingSubscription, *resty.Response, error)

		// CreateV1 creates a new volume purchasing subscription (Create Volume Purchasing Subscription record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-subscriptions
		CreateV1(ctx context.Context, request *RequestVolumePurchasingSubscription) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified volume purchasing subscription by ID (Update specified Volume Purchasing Subscription object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-volume-purchasing-subscriptions-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, *resty.Response, error)

		// DeleteByIDV1 removes the specified volume purchasing subscription by ID (Remove specified Volume Purchasing Subscription record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-subscriptions-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the volume purchasing subscriptions-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
	VolumePurchasingSubscriptions struct {
		client transport.HTTPClient
	}
)

var _ VolumePurchasingSubscriptionsServiceInterface = (*VolumePurchasingSubscriptions)(nil)

func NewVolumePurchasingSubscriptions(client transport.HTTPClient) *VolumePurchasingSubscriptions {
	return &VolumePurchasingSubscriptions{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Volume Purchasing Subscriptions CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all volume purchasing subscription objects.
// URL: GET /api/v1/volume-purchasing-subscriptions
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
func (s *VolumePurchasingSubscriptions) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProVolumePurchasingSubscriptionsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified volume purchasing subscription by ID.
// URL: GET /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions-id
func (s *VolumePurchasingSubscriptions) GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingSubscription, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing subscription ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingSubscriptionsV1, id)

	var result ResourceVolumePurchasingSubscription

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new volume purchasing subscription.
// URL: POST /api/v1/volume-purchasing-subscriptions
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-subscriptions
func (s *VolumePurchasingSubscriptions) CreateV1(ctx context.Context, request *RequestVolumePurchasingSubscription) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProVolumePurchasingSubscriptionsV1

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

// UpdateByIDV1 updates the specified volume purchasing subscription by ID.
// URL: PUT /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-volume-purchasing-subscriptions-id
func (s *VolumePurchasingSubscriptions) UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingSubscriptionsV1, id)

	var result ResourceVolumePurchasingSubscription

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

// DeleteByIDV1 removes the specified volume purchasing subscription by ID.
// URL: DELETE /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-subscriptions-id
func (s *VolumePurchasingSubscriptions) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing subscription ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingSubscriptionsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
