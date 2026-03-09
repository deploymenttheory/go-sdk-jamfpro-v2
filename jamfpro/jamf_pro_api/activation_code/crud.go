package activation_code

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service provides methods for interacting with activation code endpoints.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-activation-code-history
	ActivationCode struct {
		client transport.HTTPClient
	}
)

// NewService creates a new activation_code service.
func NewActivationCode(client transport.HTTPClient) *ActivationCode {
	return &ActivationCode{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Activation Code Operations
// -----------------------------------------------------------------------------

// UpdateV1 updates the activation code in Jamf Pro.
// URL: PUT /api/v1/activation-code
// https://developer.jamf.com/jamf-pro/reference/put_v1-activation-code
func (s *ActivationCode) UpdateV1(ctx context.Context, req *ActivationCodeRequest) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProActivationCodeV1

	headers := map[string]string{
		"Content-Type": constants.ApplicationJSON,
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
func (s *ActivationCode) UpdateOrganizationNameV1(ctx context.Context, req *OrganizationNameRequest) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProActivationCodeOrganizationNameV1

	headers := map[string]string{
		"Content-Type": constants.ApplicationJSON,
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
func (s *ActivationCode) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProActivationCodeHistoryV1

	var result HistoryResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var items []HistoryEntry
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	// Set TotalCount to the number of results we collected
	result.TotalCount = len(result.Results)

	return &result, resp, nil
}

// AddHistoryNoteV1 adds a note to activation code history.
// URL: POST /api/v1/activation-code/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-activation-code-history
func (s *ActivationCode) AddHistoryNoteV1(ctx context.Context, req *HistoryNoteRequest) (*HistoryEntry, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProActivationCodeHistoryV1

	var result HistoryEntry

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *ActivationCode) ExportHistoryV1(ctx context.Context, queryParams map[string]string, req *HistoryExportRequest) (*HistoryExportResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProActivationCodeHistoryExportV1

	var result HistoryExportResponse

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
