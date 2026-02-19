package icons

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// IconsServiceInterface defines the interface for icon operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
	IconsServiceInterface interface {
		// GetByIDV1 returns icon metadata by ID (Get Icon).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-id
		GetByIDV1(ctx context.Context, id int) (*ResourceIcon, *interfaces.Response, error)

		// UploadV1 uploads an icon image (Create Icon).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
		UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceIcon, *interfaces.Response, error)
	}

	// Service handles communication with the icons endpoint.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ IconsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetByIDV1 returns icon metadata by ID.
// URL: GET /api/v1/icon/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-id
func (s *Service) GetByIDV1(ctx context.Context, id int) (*ResourceIcon, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", EndpointIconsV1, id)
	var result ResourceIcon
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UploadV1 uploads an icon image (multipart/form-data, field "file", image/png).
// URL: POST /api/v1/icon
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
func (s *Service) UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceIcon, *interfaces.Response, error) {
	if fileName == "" {
		fileName = "icon.png"
	}
	headers := map[string]string{"file": "image/png"}
	var result ResourceIcon
	resp, err := s.client.PostMultipart(ctx, EndpointIconsV1, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UploadV1FromFile opens the file at filePath and uploads it via UploadV1.
func (s *Service) UploadV1FromFile(ctx context.Context, filePath string) (*ResourceIcon, *interfaces.Response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("open icon file: %w", err)
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return nil, nil, fmt.Errorf("stat icon file: %w", err)
	}
	name := info.Name()
	if name == "" {
		name = "icon.png"
	}
	return s.UploadV1(ctx, f, info.Size(), name)
}
