package smtp_server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SMTP server-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
	SmtpServer struct {
		client client.Client
	}
)

func NewSmtpServer(client client.Client) *SmtpServer {
	return &SmtpServer{client: client}
}

// GetV2 returns the current SMTP server configuration.
// URL: GET /api/v2/smtp-server
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-smtp-server
func (s *SmtpServer) GetV2(ctx context.Context) (*ResourceSMTPServer, *resty.Response, error) {
	var result ResourceSMTPServer

	endpoint := constants.EndpointJamfProSMTPServerV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	if _, ok := validAuthenticationTypes[request.AuthenticationType]; !ok {
		return nil, nil, fmt.Errorf("invalid authenticationType %q: must be one of NONE, BASIC, GRAPH_API, GOOGLE_MAIL", request.AuthenticationType)
	}

	if request.ConnectionSettings != nil {
		if _, ok := validEncryptionTypes[request.ConnectionSettings.EncryptionType]; !ok {
			return nil, nil, fmt.Errorf("invalid encryptionType %q: must be one of NONE, SSL, TLS_1_2, TLS_1_1, TLS_1, TLS_1_3", request.ConnectionSettings.EncryptionType)
		}
	}

	var result ResourceSMTPServer

	endpoint := constants.EndpointJamfProSMTPServerV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationMergePatchJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
