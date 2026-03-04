package apns_client_push_status

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// APNSClientPushStatusServiceInterface defines the interface for APNS client push status operations.
type APNSClientPushStatusServiceInterface interface {
	// ListV1 retrieves MDM clients with push notifications disabled.
	ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

	// EnableAllClientsV1 creates a request to enable push notifications for all MDM clients with push disabled.
	EnableAllClientsV1(ctx context.Context) (*resty.Response, error)

	// GetEnableAllClientsStatusV1 retrieves the status of the most recent enable-all-clients request.
	GetEnableAllClientsStatusV1(ctx context.Context) (*EnableAllClientsStatusResponse, *resty.Response, error)

	// EnableClientV1 enables push notifications for a single MDM client by management ID.
	EnableClientV1(ctx context.Context, req *EnableClientRequest) (*resty.Response, error)
}

// Service provides methods for interacting with APNS client push status endpoints.
type Service struct {
	client interfaces.HTTPClient
}

var _ APNSClientPushStatusServiceInterface = (*Service)(nil)

// NewService creates a new apns_client_push_status service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 retrieves MDM clients with push notifications disabled with optional RSQL filtering.
// See: https://developer.jamf.com/jamf-pro/reference/get_v1-apns-client-push-status
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := EndpointAPNSClientPushStatusV1
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageResults []PushStatusEntry
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get APNS client push status: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// EnableAllClientsV1 creates a request to enable push notifications for all MDM clients with push disabled.
// POST /api/v1/apns-client-push-status/enable-all-clients
// This is an asynchronous operation; use GetEnableAllClientsStatusV1 to check progress.
func (s *Service) EnableAllClientsV1(ctx context.Context) (*resty.Response, error) {
	endpoint := EndpointEnableAllClientsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to enable all APNS clients: %w", err)
	}

	return resp, nil
}

// GetEnableAllClientsStatusV1 retrieves the status of the most recent enable-all-clients request.
// GET /api/v1/apns-client-push-status/enable-all-clients/status
// Returns 404 if no recent request exists.
func (s *Service) GetEnableAllClientsStatusV1(ctx context.Context) (*EnableAllClientsStatusResponse, *resty.Response, error) {
	endpoint := EndpointEnableAllClientsStatusV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result EnableAllClientsStatusResponse
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get enable-all-clients status: %w", err)
	}

	return &result, resp, nil
}

// EnableClientV1 enables push notifications for a single MDM client by management ID.
// POST /api/v1/apns-client-push-status/enable-client
// Returns 204 No Content on success.
func (s *Service) EnableClientV1(ctx context.Context, req *EnableClientRequest) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.ManagementID == "" {
		return nil, fmt.Errorf("managementId is required")
	}

	endpoint := EndpointEnableClientV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to enable APNS client: %w", err)
	}

	return resp, nil
}
