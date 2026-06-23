package mobile_devices

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// MobileDevices handles communication with the Mobile Devices (v2) methods of
	// the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices
	MobileDevices struct {
		client client.Client
	}
)

func NewMobileDevices(client client.Client) *MobileDevices {
	return &MobileDevices{client: client}
}

// ListV2 returns a paginated list of basic mobile device records.
// URL: GET /api/v2/mobile-devices
// query supports: sort, filter, page, page-size (all optional). page and
// page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices
func (s *MobileDevices) ListV2(ctx context.Context, query map[string]string) (*MobileDeviceListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProMobileDevicesV2

	var result MobileDeviceListResponse

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceMobileDevice
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV2 returns the specified basic mobile device record by ID.
// URL: GET /api/v2/mobile-devices/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices-id
func (s *MobileDevices) GetByIDV2(ctx context.Context, id string) (*ResourceMobileDevice, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProMobileDevicesV2, id)

	var result ResourceMobileDevice

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDetailV2 returns a paginated list of full mobile device inventory records.
// URL: GET /api/v2/mobile-devices/detail
// query supports: section (repeatable; a map[string]string holds a single value
// — callers needing multiple sections may pass a comma-joined value), sort,
// filter, page, page-size, and exception-handling (STRICT default / LENIENT,
// Jamf Pro 11.29+). page and page-size are managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices-detail
func (s *MobileDevices) GetDetailV2(ctx context.Context, query map[string]string) (*MobileDeviceDetailListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProMobileDevicesDetailV2

	var result MobileDeviceDetailListResponse

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceMobileDeviceDetail
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil {
		if bodyBytes := resp.Bytes(); len(bodyBytes) > 0 {
			var pageResp struct {
				TotalCount int `json:"totalCount"`
			}
			if err := json.Unmarshal(bodyBytes, &pageResp); err == nil {
				result.TotalCount = pageResp.TotalCount
			}
		}
	}
	if result.TotalCount == 0 {
		result.TotalCount = len(result.Results)
	}

	return &result, resp, nil
}

// GetDetailByIDV2 returns the full mobile device inventory record for the
// specified ID.
// URL: GET /api/v2/mobile-devices/{id}/detail
// query supports: section (all sections returned if not specified).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices-id-detail
func (s *MobileDevices) GetDetailByIDV2(ctx context.Context, id string, query map[string]string) (*ResourceMobileDeviceDetailsV2, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/detail", constants.EndpointJamfProMobileDevicesV2, id)

	var result ResourceMobileDeviceDetailsV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(query).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPairedDevicesByIDV2 returns a paginated list of full mobile device
// inventory records of all paired devices for the specified device.
// URL: GET /api/v2/mobile-devices/{id}/paired-devices
// query supports: section, sort, filter, page, page-size. page and page-size are
// managed internally by GetPaginated.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-devices-id-paired-devices
func (s *MobileDevices) GetPairedDevicesByIDV2(ctx context.Context, id string, query map[string]string) (*MobileDeviceDetailListResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/paired-devices", constants.EndpointJamfProMobileDevicesV2, id)

	var result MobileDeviceDetailListResponse

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceMobileDeviceDetail
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetQueryParams(query).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil {
		if bodyBytes := resp.Bytes(); len(bodyBytes) > 0 {
			var pageResp struct {
				TotalCount int `json:"totalCount"`
			}
			if err := json.Unmarshal(bodyBytes, &pageResp); err == nil {
				result.TotalCount = pageResp.TotalCount
			}
		}
	}
	if result.TotalCount == 0 {
		result.TotalCount = len(result.Results)
	}

	return &result, resp, nil
}
