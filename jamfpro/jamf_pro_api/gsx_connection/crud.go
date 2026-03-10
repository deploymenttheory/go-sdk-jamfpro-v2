package gsx_connection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the GSX connection-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
	GsxConnection struct {
		client client.Client
	}
)

func NewGsxConnection(client client.Client) *GsxConnection {
	return &GsxConnection{client: client}
}

// GetV1 retrieves the GSX connection settings.
// URL: GET /api/v1/gsx-connection
// https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
func (s *GsxConnection) GetV1(ctx context.Context) (*ResourceGSXConnection, *resty.Response, error) {
	var result ResourceGSXConnection

	endpoint := constants.EndpointJamfProGSXConnectionV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// ReplaceV1 replaces the GSX connection settings via PUT.
// URL: PUT /api/v1/gsx-connection
// https://developer.jamf.com/jamf-pro/reference/put_v1-gsx-connection
func (s *GsxConnection) ReplaceV1(ctx context.Context, request *ResourceGSXConnection) (*ResourceGSXConnection, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceGSXConnection

	endpoint := constants.EndpointJamfProGSXConnectionV1

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

// UpdateV1 updates the GSX connection settings via PATCH.
// URL: PATCH /api/v1/gsx-connection
// https://developer.jamf.com/jamf-pro/reference/patch_v1-gsx-connection
func (s *GsxConnection) UpdateV1(ctx context.Context, request *ResourceGSXConnection) (*ResourceGSXConnection, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceGSXConnection

	endpoint := constants.EndpointJamfProGSXConnectionV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationMergePatchJSON).
		SetBody(request).
		SetResult(&result).
		Patch(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV1 retrieves GSX connection history with optional sorting.
// URL: GET /api/v1/gsx-connection/history
// https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection-history
func (s *GsxConnection) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProGSXConnectionV1)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	req := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON)
	if rsqlQuery != nil {
		req = req.SetQueryParams(rsqlQuery)
	}
	resp, err := req.GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get GSX connection history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteV1 adds a history note to the GSX connection.
// URL: POST /api/v1/gsx-connection/history
// https://developer.jamf.com/jamf-pro/reference/post_v1-gsx-connection-history
func (s *GsxConnection) AddHistoryNoteV1(ctx context.Context, request *AddHistoryNoteRequest) (*AddHistoryNoteResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result AddHistoryNoteResponse

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProGSXConnectionV1)

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

// TestV1 tests the functionality of the GSX connection.
// URL: POST /api/v1/gsx-connection/test
// https://developer.jamf.com/jamf-pro/reference/post_v1-gsx-connection-test
func (s *GsxConnection) TestV1(ctx context.Context) (*resty.Response, error) {
	endpoint := fmt.Sprintf("%s/test", constants.EndpointJamfProGSXConnectionV1)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
