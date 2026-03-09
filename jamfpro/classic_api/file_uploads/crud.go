package file_uploads

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// fileUploadFormFieldName is the form field name expected by the Classic API file uploads endpoint.
// See: https://developer.jamf.com/jamf-pro/reference/fileuploads
const fileUploadFormFieldName = "name"

type (
	// FileUploadsServiceInterface defines the interface for Classic API file upload operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
	FileUploadsServiceInterface interface {
		// CreateAttachment uploads a file to a specific resource in Jamf Pro.
		//
		// resource must be one of ValidFileUploadResources.
		// idType specifies whether identifier is an ID or name.
		// identifier is the resource ID (e.g. "123") or name.
		// filePath is the path to the file on disk.
		// forceIpaUpload, when true and resource is mobiledeviceapplicationsipa,
		// appends ?FORCE_IPA_UPLOAD=true to the request.
		//
		// Note: peripherals resource only supports ResourceIDTypeID.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
		CreateAttachment(ctx context.Context, resource string, idType ResourceIDType, identifier string, filePath string, forceIpaUpload bool) (*resty.Response, error)
	}

	// Service handles communication with the file-uploads-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
	FileUploads struct {
		client interfaces.HTTPClient
	}
)

var _ FileUploadsServiceInterface = (*FileUploads)(nil)

// NewService returns a new file uploads Service backed by the provided HTTP client.
func NewFileUploads(client interfaces.HTTPClient) *FileUploads {
	return &FileUploads{client: client}
}

// CreateAttachment uploads a file to a specific resource in Jamf Pro.
// URL: POST /JSSResource/fileuploads/{resource}/{idType}/{identifier}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
func (s *FileUploads) CreateAttachment(ctx context.Context, resource string, idType ResourceIDType, identifier string, filePath string, forceIpaUpload bool) (*resty.Response, error) {
	// Validate resource
	validResource := false
	for _, r := range ValidFileUploadResources {
		if r == resource {
			validResource = true
			break
		}
	}
	if !validResource {
		return nil, fmt.Errorf("invalid resource type: %s", resource)
	}

	// Validate idType
	if idType != ResourceIDTypeID && idType != ResourceIDTypeName {
		return nil, fmt.Errorf("invalid ID type: %s", idType)
	}

	// For peripherals, only ID is supported
	if resource == "peripherals" && idType == ResourceIDTypeName {
		return nil, fmt.Errorf("peripherals resource only supports ID type")
	}

	if identifier == "" {
		return nil, fmt.Errorf("identifier cannot be empty")
	}

	if filePath == "" {
		return nil, fmt.Errorf("file path cannot be empty")
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("stat file: %w", err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("file path must point to a file, not a directory")
	}

	fileName := info.Name()
	if fileName == "" {
		fileName = filepath.Base(filePath)
	}

	// Construct endpoint
	endpoint := fmt.Sprintf("%s/%s/%s/%s", EndpointFileUploads, resource, idType, identifier)

	// Add query parameter for IPA upload if specified
	if forceIpaUpload && resource == "mobiledeviceapplicationsipa" {
		endpoint = fmt.Sprintf("%s?FORCE_IPA_UPLOAD=true", endpoint)
	}

	headers := map[string]string{
		"Accept": mime.ApplicationXML,
	}

	var result any
	resp, err := s.client.PostMultipart(ctx, endpoint, fileUploadFormFieldName, fileName, f, info.Size(), nil, headers, nil, &result)
	if err != nil {
		return resp, fmt.Errorf("failed to upload file: %w", err)
	}
	return resp, nil
}
