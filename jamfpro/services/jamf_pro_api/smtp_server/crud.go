package smtp_server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// SMTPServerServiceInterface defines the interface for SMTP server operations (singleton).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	SMTPServerServiceInterface interface {
		// GetV2 returns the current SMTP server configuration (Get SMTP server).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
		GetV2(ctx context.Context) (*ResourceSMTPServer, *interfaces.Response, error)

		// UpdateV2 updates the SMTP server configuration (Update SMTP server).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-smtp-server
		UpdateV2(ctx context.Context, request *ResourceSMTPServer) (*ResourceSMTPServer, *interfaces.Response, error)

		// GetHistoryV1 returns the paginated SMTP server history.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-smtp-server-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNoteV1 adds a note to the SMTP server history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-history
		AddHistoryNoteV1(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *interfaces.Response, error)

		// TestV1 tests the SMTP server configuration by sending a test email.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-test
		TestV1(ctx context.Context, req *TestRequest) (*interfaces.Response, error)
	}

	// Service handles communication with the SMTP server-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SMTPServerServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV2 returns the current SMTP server configuration.
// URL: GET /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
func (s *Service) GetV2(ctx context.Context) (*ResourceSMTPServer, *interfaces.Response, error) {
	var result ResourceSMTPServer

	endpoint := EndpointSMTPServerV2

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) UpdateV2(ctx context.Context, request *ResourceSMTPServer) (*ResourceSMTPServer, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSMTPServer

	endpoint := EndpointSMTPServerV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationMergePatchJSON,
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
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageResults []HistoryObject
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	endpoint := EndpointSMTPServerHistoryV1
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) AddHistoryNoteV1(ctx context.Context, req *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	var result AddHistoryNoteResponse
	endpoint := EndpointSMTPServerHistoryV1
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

// TestV1 tests the SMTP server configuration by sending a test email.
// URL: POST /api/v1/smtp-server/test
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-smtp-server-test
func (s *Service) TestV1(ctx context.Context, req *TestRequest) (*interfaces.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.RecipientEmail == "" {
		return nil, fmt.Errorf("recipientEmail is required")
	}

	endpoint := EndpointSMTPServerTestV1
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
