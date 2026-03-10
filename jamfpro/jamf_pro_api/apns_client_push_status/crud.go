package apns_client_push_status

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service provides methods for interacting with APNS client push status endpoints.
	ApnsClientPushStatus struct {
		client client.Client
	}
)

// NewService creates a new apns_client_push_status service.
func NewApnsClientPushStatus(client client.Client) *ApnsClientPushStatus {
	return &ApnsClientPushStatus{client: client}
}

// ListV1 retrieves MDM clients with push notifications disabled with optional RSQL filtering.
// See: https://developer.jamf.com/jamf-pro/reference/get_v1-apns-client-push-status
func (s *ApnsClientPushStatus) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProAPNSClientPushStatusV1
	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageResults []PushStatusEntry
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

	if err != nil {
		return nil, resp, fmt.Errorf("failed to get APNS client push status: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// EnableAllClientsV1 creates a request to enable push notifications for all MDM clients with push disabled.
// POST /api/v1/apns-client-push-status/enable-all-clients
// This is an asynchronous operation; use GetEnableAllClientsStatusV1 to check progress.
func (s *ApnsClientPushStatus) EnableAllClientsV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProEnableAllClientsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Post(endpoint)

	if err != nil {
		return resp, fmt.Errorf("failed to enable all APNS clients: %w", err)
	}

	return resp, nil
}

// GetEnableAllClientsStatusV1 retrieves the status of the most recent enable-all-clients request.
// GET /api/v1/apns-client-push-status/enable-all-clients/status
// Returns 404 if no recent request exists.
func (s *ApnsClientPushStatus) GetEnableAllClientsStatusV1(ctx context.Context) (*EnableAllClientsStatusResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProEnableAllClientsStatusV1

	var result EnableAllClientsStatusResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, fmt.Errorf("failed to get enable-all-clients status: %w", err)
	}

	return &result, resp, nil
}

// EnableClientV1 enables push notifications for a single MDM client by management ID.
// POST /api/v1/apns-client-push-status/enable-client
// Returns 204 No Content on success.
func (s *ApnsClientPushStatus) EnableClientV1(ctx context.Context, req *EnableClientRequest) (*resty.Response, error) {
	if req == nil {
		return nil, fmt.Errorf("request is required")
	}
	if req.ManagementID == "" {
		return nil, fmt.Errorf("managementId is required")
	}

	endpoint := constants.EndpointJamfProEnableClientV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		Post(endpoint)

	if err != nil {
		return resp, fmt.Errorf("failed to enable APNS client: %w", err)
	}

	return resp, nil
}
