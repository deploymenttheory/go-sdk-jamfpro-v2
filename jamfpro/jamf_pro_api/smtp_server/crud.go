package smtp_server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SMTP server-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	SmtpServer struct {
		client transport.HTTPClient
	}
)

func NewSmtpServer(client transport.HTTPClient) *SmtpServer {
	return &SmtpServer{client: client}
}

// GetV2 returns the current SMTP server configuration.
// URL: GET /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
func (s *SmtpServer) GetV2(ctx context.Context) (*ResourceSMTPServer, *resty.Response, error) {
	var result ResourceSMTPServer

	endpoint := constants.EndpointJamfProSMTPServerV2

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV2 updates the SMTP server configuration.
// URL: PUT /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-smtp-server
func (s *SmtpServer) UpdateV2(ctx context.Context, request *ResourceSMTPServer) (*ResourceSMTPServer, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSMTPServer

	endpoint := constants.EndpointJamfProSMTPServerV2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationMergePatchJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV1 returns the paginated SMTP server history.
// URL: GET /api/v1/smtp-server/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-smtp-server-history
func (s *SmtpServer) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProSMTPServerHistoryV1
	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get SMTP server history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a note to the SMTP server history.
// URL: POST /api/v1/smtp-server/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-history
func (s *SmtpServer) AddHistoryNoteV1(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	var result AddHistoryNoteResponse
	endpoint := constants.EndpointJamfProSMTPServerHistoryV1
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

// TestV1 tests the SMTP server configuration by sending a test email.
// URL: POST /api/v1/smtp-server/test
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-test
func (s *SmtpServer) TestV1(ctx context.Context, req *TestRequest) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.RecipientEmail == "" {
		return nil, fmt.Errorf("recipientEmail is required")
	}

	endpoint := constants.EndpointJamfProSMTPServerTestV1
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
