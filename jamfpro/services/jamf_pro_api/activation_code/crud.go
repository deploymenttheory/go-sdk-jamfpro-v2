package activation_code

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ActivationCodeServiceInterface defines the interface for activation code operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-activation-code-history
	ActivationCodeServiceInterface interface {
		// UpdateV1 updates the activation code in Jamf Pro.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-activation-code
		UpdateV1(ctx context.Context, req *ActivationCodeRequest) (*interfaces.Response, error)

		// UpdateOrganizationNameV1 updates the organization name in Jamf Pro.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-activation-code-organization-name
		UpdateOrganizationNameV1(ctx context.Context, req *OrganizationNameRequest) (*interfaces.Response, error)

		// GetHistoryV1 retrieves activation code history with optional RSQL filtering.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-activation-code-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNoteV1 adds a note to activation code history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-activation-code-history
		AddHistoryNoteV1(ctx context.Context, req *HistoryNoteRequest) (*HistoryEntry, *interfaces.Response, error)

		// ExportHistoryV1 exports activation code history in specified format (JSON or CSV).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-activation-code-history-export
		ExportHistoryV1(ctx context.Context, queryParams map[string]string, req *HistoryExportRequest) (*HistoryExportResponse, *interfaces.Response, error)
	}

	// Service provides methods for interacting with activation code endpoints.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-activation-code-history
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ActivationCodeServiceInterface = (*Service)(nil)

// NewService creates a new activation_code service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Activation Code Operations
// -----------------------------------------------------------------------------

// UpdateV1 updates the activation code in Jamf Pro.
// URL: PUT /api/v1/activation-code
// https://developer.jamf.com/jamf-pro/reference/put_v1-activation-code
func (s *Service) UpdateV1(ctx context.Context, req *ActivationCodeRequest) (*interfaces.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointActivationCodeV1

	headers := map[string]string{
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// UpdateOrganizationNameV1 updates the organization name in Jamf Pro.
// URL: PATCH /api/v1/activation-code/organization-name
// https://developer.jamf.com/jamf-pro/reference/patch_v1-activation-code-organization-name
func (s *Service) UpdateOrganizationNameV1(ctx context.Context, req *OrganizationNameRequest) (*interfaces.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointActivationCodeOrganizationNameV1

	headers := map[string]string{
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV1 retrieves activation code history with automatic pagination and optional RSQL filtering.
// URL: GET /api/v1/activation-code/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/get_v1-activation-code-history
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := EndpointActivationCodeHistoryV1

	var result HistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageResponse HistoryResponse
		if err := json.Unmarshal(pageData, &pageResponse); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResponse.Results...)
		result.TotalCount = pageResponse.TotalCount
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a note to activation code history.
// URL: POST /api/v1/activation-code/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-activation-code-history
func (s *Service) AddHistoryNoteV1(ctx context.Context, req *HistoryNoteRequest) (*HistoryEntry, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointActivationCodeHistoryV1

	var result HistoryEntry

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

// ExportHistoryV1 exports activation code history in specified format (JSON or CSV).
// URL: POST /api/v1/activation-code/history/export
// acceptType should be "application/json" or "text/csv"
// The request body is optional and can override query parameters if URI exceeds 2,000 characters.
// https://developer.jamf.com/jamf-pro/reference/post_v1-activation-code-history-export
func (s *Service) ExportHistoryV1(ctx context.Context, queryParams map[string]string, req *HistoryExportRequest) (*HistoryExportResponse, *interfaces.Response, error) {
	endpoint := EndpointActivationCodeHistoryExportV1

	var result HistoryExportResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	var body any
	if req != nil {
		body = req
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, queryParams, body, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
