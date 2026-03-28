package volume_purchasing_subscriptions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the volume purchasing subscriptions-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
	VolumePurchasingSubscriptions struct {
		client client.Client
	}
)

func NewVolumePurchasingSubscriptions(client client.Client) *VolumePurchasingSubscriptions {
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetResult(&result).
		Get(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	for _, trigger := range request.Triggers {
		if _, ok := validSubscriptionTriggers[trigger]; !ok {
			return nil, nil, fmt.Errorf("invalid trigger %q: must be one of NO_MORE_LICENSES, REMOVED_FROM_APP_STORE", trigger)
		}
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProVolumePurchasingSubscriptionsV1

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

	for _, trigger := range request.Triggers {
		if _, ok := validSubscriptionTriggers[trigger]; !ok {
			return nil, nil, fmt.Errorf("invalid trigger %q: must be one of NO_MORE_LICENSES, REMOVED_FROM_APP_STORE", trigger)
		}
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProVolumePurchasingSubscriptionsV1, id)

	var result ResourceVolumePurchasingSubscription

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
