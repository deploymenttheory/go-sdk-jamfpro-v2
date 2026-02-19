package webhooks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// WebhooksServiceInterface defines the interface for Classic API webhook operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/webhooks
	WebhooksServiceInterface interface {
		// ListWebhooks returns all webhooks.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findallwebhooks
		ListWebhooks(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// GetWebhookByID returns the specified webhook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findwebhooksbyid
		GetWebhookByID(ctx context.Context, id int) (*ResourceWebhook, *interfaces.Response, error)

		// GetWebhookByName returns the specified webhook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findwebhooksbyname
		GetWebhookByName(ctx context.Context, name string) (*ResourceWebhook, *interfaces.Response, error)

		// CreateWebhook creates a new webhook.
		//
		// Returns the created webhook with its assigned ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createwebhookbyid
		CreateWebhook(ctx context.Context, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error)

		// UpdateWebhookByID updates the specified webhook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatewebhookbyid
		UpdateWebhookByID(ctx context.Context, id int, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error)

		// UpdateWebhookByName updates the specified webhook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updatewebhookbyname
		UpdateWebhookByName(ctx context.Context, name string, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error)

		// DeleteWebhookByID removes the specified webhook by ID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletewebhookbyid
		DeleteWebhookByID(ctx context.Context, id int) (*interfaces.Response, error)

		// DeleteWebhookByName removes the specified webhook by name.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/deletewebhookbyname
		DeleteWebhookByName(ctx context.Context, name string) (*interfaces.Response, error)
	}

	// Service handles communication with the webhook-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/webhooks
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ WebhooksServiceInterface = (*Service)(nil)

// NewService returns a new webhooks Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Webhooks CRUD Operations
// -----------------------------------------------------------------------------

// ListWebhooks returns all webhooks.
// URL: GET /JSSResource/webhooks
// https://developer.jamf.com/jamf-pro/reference/findallwebhooks
func (s *Service) ListWebhooks(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointClassicWebhooks, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetWebhookByID returns the specified webhook by ID.
// URL: GET /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findwebhooksbyid
func (s *Service) GetWebhookByID(ctx context.Context, id int) (*ResourceWebhook, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("webhook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicWebhooks, id)

	var result ResourceWebhook

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetWebhookByName returns the specified webhook by name.
// URL: GET /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findwebhooksbyname
func (s *Service) GetWebhookByName(ctx context.Context, name string) (*ResourceWebhook, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("webhook name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicWebhooks, name)

	var result ResourceWebhook

	resp, err := s.client.Get(ctx, endpoint, nil, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateWebhook creates a new webhook.
// URL: POST /JSSResource/webhooks/id/0
// Returns the created webhook with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createwebhookbyid
func (s *Service) CreateWebhook(ctx context.Context, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", EndpointClassicWebhooks)

	var result ResourceWebhook

	resp, err := s.client.Post(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateWebhookByID updates the specified webhook by ID.
// URL: PUT /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatewebhookbyid
func (s *Service) UpdateWebhookByID(ctx context.Context, id int, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("webhook ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicWebhooks, id)

	var result ResourceWebhook

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateWebhookByName updates the specified webhook by name.
// URL: PUT /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatewebhookbyname
func (s *Service) UpdateWebhookByName(ctx context.Context, name string, req *RequestWebhook) (*ResourceWebhook, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("webhook name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicWebhooks, name)

	var result ResourceWebhook

	resp, err := s.client.Put(ctx, endpoint, req, shared.XMLHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteWebhookByID removes the specified webhook by ID.
// URL: DELETE /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletewebhookbyid
func (s *Service) DeleteWebhookByID(ctx context.Context, id int) (*interfaces.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("webhook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", EndpointClassicWebhooks, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteWebhookByName removes the specified webhook by name.
// URL: DELETE /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletewebhookbyname
func (s *Service) DeleteWebhookByName(ctx context.Context, name string) (*interfaces.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("webhook name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", EndpointClassicWebhooks, name)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.XMLHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
