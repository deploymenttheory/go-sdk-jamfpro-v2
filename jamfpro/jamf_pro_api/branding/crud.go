package branding

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"resty.dev/v3"
)

type (
	// Service handles communication with the branding image download methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
	Branding struct {
		client transport.HTTPClient
	}
)

func NewBranding(client transport.HTTPClient) *Branding {
	return &Branding{client: client}
}

// DownloadBrandingImageV1 downloads a self service branding image as raw binary data.
// URL: GET /api/v1/branding-images/download/{id}
// Uses Accept: image/* to request image response. Returns raw binary (not JSON).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
func (s *Branding) DownloadBrandingImageV1(ctx context.Context, id string) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProBrandingImagesDownloadV1, id)
	headers := map[string]string{
		"Accept": constants.ImageAny,
	}

	resp, body, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}
