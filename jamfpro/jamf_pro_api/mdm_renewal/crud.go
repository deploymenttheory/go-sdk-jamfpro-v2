package mdm_renewal

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the MDM renewal-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mdm-renewal-device-common-details-clientmanagementid
	MdmRenewal struct {
		client client.Client
	}
)

func NewMdmRenewal(client client.Client) *MdmRenewal {
	return &MdmRenewal{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - MDM Renewal Operations
// -----------------------------------------------------------------------------

// UpdateDeviceCommonDetailsV1 partially updates device common details.
// URL: PATCH /api/v1/mdm-renewal/device-common-details
// Body: RequestDeviceCommonDetailsUpdate (clientManagementId required in body)
// Response: 204 No Content
func (s *MdmRenewal) UpdateDeviceCommonDetailsV1(ctx context.Context, request *RequestDeviceCommonDetailsUpdate) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if request.ClientManagementID == "" {
		return nil, fmt.Errorf("clientManagementId is required")
	}

	endpoint := constants.EndpointJamfProDeviceCommonDetailsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Patch(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetDeviceCommonDetailsV1 returns device common details for the given client management ID.
// URL: GET /api/v1/mdm-renewal/device-common-details/{clientManagementId}
func (s *MdmRenewal) GetDeviceCommonDetailsV1(ctx context.Context, clientManagementID string) (*DeviceCommonDetails, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDeviceCommonDetailsV1, clientManagementID)

	var result DeviceCommonDetails

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetRenewalStrategiesV1 returns MDM renewal errors and strategies for the given client management ID.
// URL: GET /api/v1/mdm-renewal/renewal-strategies/{clientManagementId}
func (s *MdmRenewal) GetRenewalStrategiesV1(ctx context.Context, clientManagementID string) ([]RenewalErrorWithStrategies, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProRenewalStrategiesV1, clientManagementID)

	var result []RenewalErrorWithStrategies

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// DeleteRenewalStrategiesV1 deletes all renewal strategies and errors for the given client management ID.
// URL: DELETE /api/v1/mdm-renewal/renewal-strategies/{clientManagementId}
// Response: 204 No Content
func (s *MdmRenewal) DeleteRenewalStrategiesV1(ctx context.Context, clientManagementID string) (*resty.Response, error) {
	if clientManagementID == "" {
		return nil, fmt.Errorf("clientManagementId is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProRenewalStrategiesV1, clientManagementID)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
