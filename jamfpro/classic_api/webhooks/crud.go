package webhooks

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the webhook-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/webhooks
	Webhooks struct {
		client client.Client
	}
)

// NewService returns a new webhooks Service backed by the provided HTTP client.
func NewWebhooks(client client.Client) *Webhooks {
	return &Webhooks{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Webhooks CRUD Operations
// -----------------------------------------------------------------------------

// List returns all webhooks.
// URL: GET /JSSResource/webhooks
// https://developer.jamf.com/jamf-pro/reference/findwebhooks
func (s *Webhooks) List(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointClassicWebhooks

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified webhook by ID.
// URL: GET /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/findwebhooksbyid
func (s *Webhooks) GetByID(ctx context.Context, id int) (*ResourceWebhook, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("webhook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicWebhooks, id)

	var result ResourceWebhook

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns the specified webhook by name.
// URL: GET /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/findwebhooksbyname
func (s *Webhooks) GetByName(ctx context.Context, name string) (*ResourceWebhook, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("webhook name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicWebhooks, name)

	var result ResourceWebhook

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new webhook.
// URL: POST /JSSResource/webhooks/id/0
// Returns the created webhook with its assigned ID.
// https://developer.jamf.com/jamf-pro/reference/createwebhookbyid
func (s *Webhooks) Create(ctx context.Context, req *RequestWebhook) (*ResourceWebhook, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/0", constants.EndpointClassicWebhooks)

	var result ResourceWebhook

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified webhook by ID.
// URL: PUT /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/updatewebhookbyid
func (s *Webhooks) UpdateByID(ctx context.Context, id int, req *RequestWebhook) (*ResourceWebhook, *resty.Response, error) {
	if id <= 0 {
		return nil, nil, fmt.Errorf("webhook ID must be a positive integer")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicWebhooks, id)

	var result ResourceWebhook

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByName updates the specified webhook by name.
// URL: PUT /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/updatewebhookbyname
func (s *Webhooks) UpdateByName(ctx context.Context, name string, req *RequestWebhook) (*ResourceWebhook, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("webhook name is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicWebhooks, name)

	var result ResourceWebhook

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified webhook by ID.
// URL: DELETE /JSSResource/webhooks/id/{id}
// https://developer.jamf.com/jamf-pro/reference/deletewebhookbyid
func (s *Webhooks) DeleteByID(ctx context.Context, id int) (*resty.Response, error) {
	if id <= 0 {
		return nil, fmt.Errorf("webhook ID must be a positive integer")
	}

	endpoint := fmt.Sprintf("%s/id/%d", constants.EndpointClassicWebhooks, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteByName removes the specified webhook by name.
// URL: DELETE /JSSResource/webhooks/name/{name}
// https://developer.jamf.com/jamf-pro/reference/deletewebhookbyname
func (s *Webhooks) DeleteByName(ctx context.Context, name string) (*resty.Response, error) {
	if name == "" {
		return nil, fmt.Errorf("webhook name is required")
	}

	endpoint := fmt.Sprintf("%s/name/%s", constants.EndpointClassicWebhooks, name)

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
