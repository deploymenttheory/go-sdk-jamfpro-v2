package icon

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
	// Service handles communication with the icons-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
	Icon struct {
		client client.Client
	}
)

func NewIcon(client client.Client) *Icon {
	return &Icon{client: client}
}

// GetByIDV1 returns icon metadata by ID.
// URL: GET /api/v1/icon/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-id
func (s *Icon) GetByIDV1(ctx context.Context, id int) (*ResourceIcon, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", constants.EndpointJamfProIconV1, id)
	var result ResourceIcon

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadV1 uploads an icon image (multipart/form-data, field "file").
// URL: POST /api/v1/icon
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-icon
func (s *Icon) UploadV1(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ResourceIcon, *resty.Response, error) {
	if fileName == "" {
		fileName = "icon.png"
	}
	endpoint := constants.EndpointJamfProIconV1

	var result ResourceIcon

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

// UploadV1FromFile opens the file at filePath and uploads it via UploadV1.
func (s *Icon) UploadV1FromFile(ctx context.Context, filePath string) (*ResourceIcon, *resty.Response, error) {
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

// DownloadV1 downloads the icon image bytes. res: original, 300, or 512 (default original).
// scale: 0 = original image, non-0 = scaled to 300. Use Accept: image/*.
// URL: GET /api/v1/icon/download/{id}?res=...&scale=...
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-download-id
func (s *Icon) DownloadV1(ctx context.Context, id int, res, scale string) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", constants.EndpointJamfProIconsDownloadV1, id)
	if res == "" {
		res = "original"
	}

	if scale == "" {
		scale = "0"
	}

	rsqlQuery := map[string]string{"res": res, "scale": scale}

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ImageAny).
		SetQueryParams(rsqlQuery).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return body, resp, nil
}
