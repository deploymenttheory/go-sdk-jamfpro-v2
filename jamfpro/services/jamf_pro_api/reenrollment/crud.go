package reenrollment

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ReenrollmentServiceInterface defines the interface for re-enrollment settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
	ReenrollmentServiceInterface interface {
		// Get retrieves re-enrollment settings.
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
		Get(ctx context.Context) (*ResourceReenrollmentSettings, *interfaces.Response, error)

		// Update updates re-enrollment settings.
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-reenrollment
		Update(ctx context.Context, request *ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, *interfaces.Response, error)

		// GetHistory returns paginated re-enrollment history (page, page-size, sort).
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment-history
		GetHistory(ctx context.Context, query map[string]string) (*ReenrollmentHistoryResponse, *interfaces.Response, error)

		// AddHistoryNotes adds a note to re-enrollment history.
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-reenrollment-history
		AddHistoryNotes(ctx context.Context, request *AddReenrollmentHistoryNotesRequest) (*ReenrollmentHistoryObject, *interfaces.Response, error)

		// ExportHistory exports re-enrollment history. query may include page, page-size, sort, filter, export-fields, export-labels. body may override when URI exceeds ~2k chars. Uses Accept: text/csv,application/json.
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-reenrollment-history-export
		ExportHistory(ctx context.Context, query map[string]string, body *ExportReenrollmentHistoryRequest) (*interfaces.Response, []byte, error)
	}

	// Service handles communication with the re-enrollment settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ReenrollmentServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Re-enrollment Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves re-enrollment settings.
// URL: GET /api/v1/reenrollment
func (s *Service) Get(ctx context.Context) (*ResourceReenrollmentSettings, *interfaces.Response, error) {
	var result ResourceReenrollmentSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointReenrollmentV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Update updates re-enrollment settings.
// URL: PUT /api/v1/reenrollment
func (s *Service) Update(ctx context.Context, request *ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceReenrollmentSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointReenrollmentV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistory returns paginated re-enrollment history.
// URL: GET /api/v1/reenrollment/history
func (s *Service) GetHistory(ctx context.Context, query map[string]string) (*ReenrollmentHistoryResponse, *interfaces.Response, error) {
	var result ReenrollmentHistoryResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointReenrollmentHistoryV1, query, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// AddHistoryNotes adds a note to re-enrollment history.
// URL: POST /api/v1/reenrollment/history
func (s *Service) AddHistoryNotes(ctx context.Context, request *AddReenrollmentHistoryNotesRequest) (*ReenrollmentHistoryObject, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ReenrollmentHistoryObject

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Post(ctx, EndpointReenrollmentHistoryV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// ExportHistory exports re-enrollment history. query: page, page-size, sort, filter, export-fields, export-labels. body optional (overrides query when URI would exceed ~2k chars; use page, pageSize, sort, filter, fields). Uses Accept: text/csv,application/json and Content-Type: application/json when body is sent.
// URL: POST /api/v1/reenrollment/history/export
func (s *Service) ExportHistory(ctx context.Context, query map[string]string, body *ExportReenrollmentHistoryRequest) (*interfaces.Response, []byte, error) {
	headers := map[string]string{"Accept": "text/csv,application/json"}
	var sendBody any
	if body != nil {
		sendBody = body
		headers["Content-Type"] = mime.ApplicationJSON
	}
	resp, err := s.client.PostWithQuery(ctx, EndpointReenrollmentHistoryExport, query, sendBody, headers, nil)
	if err != nil {
		return nil, nil, err
	}
	return resp, resp.Body, nil
}
