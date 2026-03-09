package self_service_branding_upload

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"resty.dev/v3"
)

type (
	// SelfServiceBrandingUploadServiceInterface defines the interface for self-service branding image upload operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
	SelfServiceBrandingUploadServiceInterface interface {
		// Upload uploads a branding image file (multipart/form-data, field "file").
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
		Upload(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceBrandingImage, *resty.Response, error)

		// UploadFromFile opens the file at filePath and uploads it via Upload.
		UploadFromFile(ctx context.Context, filePath string) (*ResourceBrandingImage, *resty.Response, error)
	}

	// Service handles communication with the self-service branding images methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
	SelfServiceBrandingUpload struct {
		client transport.HTTPClient
	}
)

var _ SelfServiceBrandingUploadServiceInterface = (*SelfServiceBrandingUpload)(nil)

func NewSelfServiceBrandingUpload(client transport.HTTPClient) *SelfServiceBrandingUpload {
	return &SelfServiceBrandingUpload{client: client}
}

// Upload uploads a branding image file (multipart/form-data, field "file").
// URL: POST /api/self-service/branding/images
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_self-service-branding-images
func (s *SelfServiceBrandingUpload) Upload(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceBrandingImage, *resty.Response, error) {
	if fileName == "" {
		fileName = "branding.png"
	}
	headers := map[string]string{"Content-Type": "multipart/form-data"}
	var result ResourceBrandingImage
	resp, err := s.client.PostMultipart(ctx, EndpointBrandingImages, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
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
