package gsx_connection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// GSXConnectionServiceInterface defines the interface for GSX connection operations.
	//
	// Manages GSX (Global Service Exchange) connection settings for Apple repair services.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
	GSXConnectionServiceInterface interface {
		// GetV1 retrieves the GSX connection settings.
		//
		// Returns current configuration including keystore details and service account info.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
		GetV1(ctx context.Context) (*ResourceGSXConnection, *interfaces.Response, error)

		// UpdateV1 updates the GSX connection settings via PATCH.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-gsx-connection
		UpdateV1(ctx context.Context, request *ResourceGSXConnection) (*ResourceGSXConnection, *interfaces.Response, error)

		// GetHistoryV1 retrieves GSX connection history with optional sorting.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)
	}

	// Service handles communication with the GSX connection-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ GSXConnectionServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the GSX connection settings.
// URL: GET /api/v1/gsx-connection
// https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection
func (s *Service) GetV1(ctx context.Context) (*ResourceGSXConnection, *interfaces.Response, error) {
	var result ResourceGSXConnection

	endpoint := EndpointGSXConnectionV1
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates the GSX connection settings via PATCH.
// URL: PATCH /api/v1/gsx-connection
// https://developer.jamf.com/jamf-pro/reference/patch_v1-gsx-connection
func (s *Service) UpdateV1(ctx context.Context, request *ResourceGSXConnection) (*ResourceGSXConnection, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceGSXConnection

	endpoint := EndpointGSXConnectionV1
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV1 retrieves GSX connection history with optional sorting.
// URL: GET /api/v1/gsx-connection/history
// https://developer.jamf.com/jamf-pro/reference/get_v1-gsx-connection-history
func (s *Service) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/history", EndpointGSXConnectionV1)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageResults []HistoryObject
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
		return nil, resp, fmt.Errorf("failed to get GSX connection history: %w", err)
	}

	result.TotalCount = len(result.Results)

	return &result, resp, nil
}
