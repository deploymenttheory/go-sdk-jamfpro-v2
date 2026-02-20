package volume_purchasing_subscriptions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// VolumePurchasingSubscriptionsServiceInterface defines the interface for volume purchasing subscription operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
	VolumePurchasingSubscriptionsServiceInterface interface {
		// ListV1 returns all volume purchasing subscription objects (Get Volume Purchasing Subscription objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified volume purchasing subscription by ID (Get specified Volume Purchasing Subscription object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions-id
		GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingSubscription, *interfaces.Response, error)

		// CreateV1 creates a new volume purchasing subscription (Create Volume Purchasing Subscription record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-subscriptions
		CreateV1(ctx context.Context, request *RequestVolumePurchasingSubscription) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified volume purchasing subscription by ID (Update specified Volume Purchasing Subscription object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-volume-purchasing-subscriptions-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified volume purchasing subscription by ID (Remove specified Volume Purchasing Subscription record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-subscriptions-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the volume purchasing subscriptions-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ VolumePurchasingSubscriptionsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Volume Purchasing Subscriptions CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all volume purchasing subscription objects.
// URL: GET /api/v1/volume-purchasing-subscriptions
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointVolumePurchasingSubscriptionsV1

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

// GetByIDV1 returns the specified volume purchasing subscription by ID.
// URL: GET /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceVolumePurchasingSubscription, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("volume purchasing subscription ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingSubscriptionsV1, id)

	var result ResourceVolumePurchasingSubscription

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

// CreateV1 creates a new volume purchasing subscription.
// URL: POST /api/v1/volume-purchasing-subscriptions
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-volume-purchasing-subscriptions
func (s *Service) CreateV1(ctx context.Context, request *RequestVolumePurchasingSubscription) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointVolumePurchasingSubscriptionsV1

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

// UpdateByIDV1 updates the specified volume purchasing subscription by ID.
// URL: PUT /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-volume-purchasing-subscriptions-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestVolumePurchasingSubscription) (*ResourceVolumePurchasingSubscription, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingSubscriptionsV1, id)

	var result ResourceVolumePurchasingSubscription

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

// DeleteByIDV1 removes the specified volume purchasing subscription by ID.
// URL: DELETE /api/v1/volume-purchasing-subscriptions/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-volume-purchasing-subscriptions-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("volume purchasing subscription ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointVolumePurchasingSubscriptionsV1, id)

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
