package reenrollment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the re-enrollment settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
	Reenrollment struct {
		client client.Client
	}
)

func NewReenrollment(client client.Client) *Reenrollment {
	return &Reenrollment{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Re-enrollment Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves re-enrollment settings.
// URL: GET /api/v1/reenrollment
func (s *Reenrollment) Get(ctx context.Context) (*ResourceReenrollmentSettings, *resty.Response, error) {
	var result ResourceReenrollmentSettings

	endpoint := constants.EndpointJamfProReenrollmentV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Update updates re-enrollment settings.
// URL: PUT /api/v1/reenrollment
func (s *Reenrollment) Update(ctx context.Context, request *ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceReenrollmentSettings

	endpoint := constants.EndpointJamfProReenrollmentV1

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

// GetHistory returns paginated re-enrollment history.
// URL: GET /api/v1/reenrollment/history
func (s *Reenrollment) GetHistory(ctx context.Context, query map[string]string) (*ReenrollmentHistoryResponse, *resty.Response, error) {
	var result ReenrollmentHistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ReenrollmentHistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProReenrollmentHistoryV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotes adds a note to re-enrollment history.
// URL: POST /api/v1/reenrollment/history
func (s *Reenrollment) AddHistoryNotes(ctx context.Context, request *AddReenrollmentHistoryNotesRequest) (*ReenrollmentHistoryObject, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ReenrollmentHistoryObject

	endpoint := constants.EndpointJamfProReenrollmentHistoryV1

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

// ExportHistory exports re-enrollment history. query: page, page-size, sort, filter, export-fields, export-labels. body optional (overrides query when URI would exceed ~2k chars; use page, pageSize, sort, filter, fields). Uses Accept: text/csv,application/json and Content-Type: application/json when body is sent.
// URL: POST /api/v1/reenrollment/history/export
func (s *Reenrollment) ExportHistory(ctx context.Context, query map[string]string, body *ExportReenrollmentHistoryRequest) (*resty.Response, []byte, error) {
	endpoint := constants.EndpointJamfProReenrollmentHistoryExport

	var sendBody any
	if body != nil {
		sendBody = body
	}

	req := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.TextCSVApplicationJSON).
		SetQueryParams(query)

	if body != nil {
		req = req.SetHeader("Content-Type", constants.ApplicationJSON).SetBody(sendBody)
	}

	resp, err := req.Post(endpoint)
	if err != nil {
		return nil, nil, err
	}
	return resp, resp.Bytes(), nil
}
