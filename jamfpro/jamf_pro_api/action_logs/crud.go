package action_logs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"net/url"
	"resty.dev/v3"
)

// ActionLogs provides action history and per-device outcomes.
type ActionLogs struct{ client client.Client }

func NewActionLogs(client client.Client) *ActionLogs { return &ActionLogs{client: client} }

// ListV1 retrieves action logs with pagination.
func (s *ActionLogs) ListV1(ctx context.Context, query map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse
	endpoint := constants.EndpointJamfProActionLogsV1
	mergePage := func(data []byte) error {
		var page []ResourceActionLog
		if err := json.Unmarshal(data, &page); err != nil {
			return err
		}
		result.Results = append(result.Results, page...)
		return nil
	}
	resp, err := s.client.NewRequest(ctx).SetHeader("Accept", constants.ApplicationJSON).SetQueryParams(query).GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetDevicesV1 retrieves action logs device outcomes with pagination.
func (s *ActionLogs) GetDevicesV1(ctx context.Context, id string, query map[string]string) (*DeviceListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	var result DeviceListResponse
	endpoint := fmt.Sprintf("%s/%s/devices", constants.EndpointJamfProActionLogsV1, url.PathEscape(id))
	mergePage := func(data []byte) error {
		var page []ResourceActionLogDevice
		if err := json.Unmarshal(data, &page); err != nil {
			return err
		}
		result.Results = append(result.Results, page...)
		return nil
	}
	resp, err := s.client.NewRequest(ctx).SetHeader("Accept", constants.ApplicationJSON).SetQueryParams(query).GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 retrieves an action log and its status summary.
func (s *ActionLogs) GetByIDV1(ctx context.Context, id string) (*ResourceActionLogDetail, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	var result ResourceActionLogDetail
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProActionLogsV1, url.PathEscape(id))
	resp, err := s.client.NewRequest(ctx).SetHeader("Accept", constants.ApplicationJSON).SetResult(&result).Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
