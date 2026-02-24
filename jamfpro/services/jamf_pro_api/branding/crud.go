package branding

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
)

type (
	// BrandingServiceInterface defines the interface for branding image operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
	BrandingServiceInterface interface {
		// DownloadBrandingImageV1 downloads a self service branding image as raw binary data.
		// Uses Accept: image/* header. Returns raw bytes (not JSON).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
		DownloadBrandingImageV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the branding image download methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ BrandingServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// DownloadBrandingImageV1 downloads a self service branding image as raw binary data.
// URL: GET /api/v1/branding-images/download/{id}
// Uses Accept: image/* to request image response. Returns raw binary (not JSON).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-branding-images-download-id
func (s *Service) DownloadBrandingImageV1(ctx context.Context, id string) ([]byte, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", EndpointBrandingImagesDownloadV1, id)
	headers := map[string]string{"Accept": "image/*"}

	resp, body, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}
