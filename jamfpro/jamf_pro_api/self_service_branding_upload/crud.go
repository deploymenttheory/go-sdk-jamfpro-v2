package self_service_branding_upload

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the self-service branding images methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
	SelfServiceBrandingUpload struct {
		client client.Client
	}
)

func NewSelfServiceBrandingUpload(client client.Client) *SelfServiceBrandingUpload {
	return &SelfServiceBrandingUpload{client: client}
}

// Upload uploads a branding image file (multipart/form-data, field "file").
// URL: POST /api/self-service/branding/images
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
func (s *SelfServiceBrandingUpload) Upload(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceBrandingImage, *resty.Response, error) {
	if fileName == "" {
		fileName = "branding.png"
	}

	endpoint := constants.EndpointJamfProBrandingImages

	var result ResourceBrandingImage

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Content-Type", constants.MultipartFormData).
		SetMultipartFile("file", fileName, fileReader, fileSize, nil).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UploadFromFile opens the file at filePath and uploads it via Upload.
func (s *SelfServiceBrandingUpload) UploadFromFile(ctx context.Context, filePath string) (*ResourceBrandingImage, *resty.Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open branding image file: %w", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat branding image file: %w", err)
	}
	name := info.Name()
	if name == "" {
		name = "branding.png"
	}
	return s.Upload(ctx, f, info.Size(), name)
}
