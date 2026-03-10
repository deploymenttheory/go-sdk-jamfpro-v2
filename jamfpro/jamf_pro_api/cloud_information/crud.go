package cloud_information

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the cloud information-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-information
	CloudInformation struct {
		client client.Client
	}
)

func NewCloudInformation(client client.Client) *CloudInformation {
	return &CloudInformation{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Cloud Information Operations
// -----------------------------------------------------------------------------

// GetV1 returns information related to cloud setup.
// URL: GET /api/v1/cloud-information
// https://developer.jamf.com/jamf-pro/reference/get_v1-cloud-information
func (s *CloudInformation) GetV1(ctx context.Context) (*ResourceCloudInformation, *resty.Response, error) {
	var result ResourceCloudInformation

	endpoint := constants.EndpointJamfProCloudInformationV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
