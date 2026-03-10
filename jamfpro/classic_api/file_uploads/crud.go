package file_uploads

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

// fileUploadFormFieldName is the form field name expected by the Classic API file uploads endpoint.
// See: https://developer.jamf.com/jamf-pro/reference/fileuploads
const fileUploadFormFieldName = "name"

// ValidFileUploadResources contains the list of valid resources for file uploads.
var ValidFileUploadResources = []string{
	"computers",
	"mobiledevices",
	"enrollmentprofiles",
	"printers",
	"peripherals",
	"policies",
	"ebooks",
	"mobiledeviceapplications",
	"icon",
	"mobiledeviceapplicationsipa",
	"diskencryptionconfigurations",
}

type (
	// Service handles communication with the file-uploads-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
	FileUploads struct {
		client client.Client
	}
)

// NewService returns a new file uploads Service backed by the provided HTTP client.
func NewFileUploads(client client.Client) *FileUploads {
	return &FileUploads{client: client}
}

// CreateAttachment uploads a file to a specific resource in Jamf Pro.
// URL: POST /JSSResource/fileuploads/{resource}/{idType}/{identifier}
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/fileuploads
func (s *FileUploads) CreateAttachment(ctx context.Context, resource string, idType ResourceIDType, identifier string, filePath string, forceIpaUpload bool) (*resty.Response, error) {
	// Validate resource
	if !slices.Contains(ValidFileUploadResources, resource) {
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

	endpoint := fmt.Sprintf("%s/%s/%s/%s", constants.EndpointClassicFileUploads, resource, idType, identifier)

	// Add query parameter for IPA upload if specified
	if forceIpaUpload && resource == "mobiledeviceapplicationsipa" {
		endpoint = fmt.Sprintf("%s?FORCE_IPA_UPLOAD=true", endpoint)
	}

	headers := map[string]string{
		"Accept": constants.ApplicationXML,
	}

	var result any
	resp, err := s.client.PostMultipart(ctx, endpoint, fileUploadFormFieldName, fileName, f, info.Size(), nil, headers, nil, &result)
	if err != nil {
		return resp, fmt.Errorf("failed to upload file: %w", err)
	}
	return resp, nil
}
