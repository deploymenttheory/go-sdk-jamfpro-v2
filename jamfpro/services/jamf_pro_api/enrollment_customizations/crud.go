package enrollment_customizations

import (
	"context"
	"fmt"
	"io"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// EnrollmentCustomizationsServiceInterface defines the interface for enrollment customization operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
	EnrollmentCustomizationsServiceInterface interface {
		// ListV2 returns a paged list of enrollment customization objects.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
		ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV2 returns the specified enrollment customization by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id
		GetByIDV2(ctx context.Context, id string) (*ResourceEnrollmentCustomization, *interfaces.Response, error)

		// GetByNameV2 returns the specified enrollment customization by display name.
		//
		// This is a convenience method that calls ListV2 and filters by name.
		GetByNameV2(ctx context.Context, name string) (*ResourceEnrollmentCustomization, *interfaces.Response, error)

		// CreateV2 creates a new enrollment customization record.
		//
		// Returns the created enrollment customization's ID and href.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations
		CreateV2(ctx context.Context, request *ResourceEnrollmentCustomization) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV2 replaces the specified enrollment customization by ID.
		//
		// Returns the full updated enrollment customization resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-enrollment-customizations-id
		UpdateByIDV2(ctx context.Context, id string, request *ResourceEnrollmentCustomization) (*ResourceEnrollmentCustomization, *interfaces.Response, error)

		// DeleteByIDV2 removes the specified enrollment customization by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-enrollment-customizations-id
		DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error)

		// GetHistoryV2 returns the history object for the specified enrollment customization.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-history
		GetHistoryV2(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNotesV2 adds notes to the specified enrollment customization's history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-id-history
		AddHistoryNotesV2(ctx context.Context, id string, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error)

		// GetPrestagesV2 retrieves the list of prestages using this enrollment customization.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-prestages
		GetPrestagesV2(ctx context.Context, id string) (*PrestagesResponse, *interfaces.Response, error)

		// UploadImageV2 uploads an image for enrollment customizations.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-images
		UploadImageV2(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ImageUploadResponse, *interfaces.Response, error)

		// GetImageByIdV2 retrieves an enrollment customization image by ID.
		//
		// Returns the image file data as bytes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-images-id
		GetImageByIdV2(ctx context.Context, id string) ([]byte, *interfaces.Response, error)
	}

	// Service handles communication with the enrollment customizations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ EnrollmentCustomizationsServiceInterface = (*Service)(nil)

// NewService returns a new enrollment customizations Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Enrollment Customizations CRUD Operations
// -----------------------------------------------------------------------------

// ListV2 returns a paged list of enrollment customization objects.
// URL: GET /api/v2/enrollment-customizations
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
func (s *Service) ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointEnrollmentCustomizationsV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV2 returns the specified enrollment customization by ID.
// URL: GET /api/v2/enrollment-customizations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id
func (s *Service) GetByIDV2(ctx context.Context, id string) (*ResourceEnrollmentCustomization, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointEnrollmentCustomizationsV2, id)

	var result ResourceEnrollmentCustomization

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV2 returns the specified enrollment customization by display name.
// This is a convenience method that calls ListV2 and filters by name.
func (s *Service) GetByNameV2(ctx context.Context, name string) (*ResourceEnrollmentCustomization, *interfaces.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("enrollment customization name is required")
	}

	list, resp, err := s.ListV2(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, customization := range list.Results {
		if customization.DisplayName == name {
			return s.GetByIDV2(ctx, customization.ID)
		}
	}

	return nil, resp, fmt.Errorf("enrollment customization with name %q not found", name)
}

// CreateV2 creates a new enrollment customization record.
// URL: POST /api/v2/enrollment-customizations
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations
func (s *Service) CreateV2(ctx context.Context, request *ResourceEnrollmentCustomization) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	var result CreateResponse

	endpoint := EndpointEnrollmentCustomizationsV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 replaces the specified enrollment customization by ID.
// URL: PUT /api/v2/enrollment-customizations/{id}
// Returns the full updated enrollment customization resource.
// https://developer.jamf.com/jamf-pro/reference/put_v2-enrollment-customizations-id
func (s *Service) UpdateByIDV2(ctx context.Context, id string, request *ResourceEnrollmentCustomization) (*ResourceEnrollmentCustomization, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointEnrollmentCustomizationsV2, id)

	var result ResourceEnrollmentCustomization

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV2 removes the specified enrollment customization by ID.
// URL: DELETE /api/v2/enrollment-customizations/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v2-enrollment-customizations-id
func (s *Service) DeleteByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointEnrollmentCustomizationsV2, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV2 returns the history object for the specified enrollment customization.
// URL: GET /api/v2/enrollment-customizations/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-history
func (s *Service) GetHistoryV2(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointEnrollmentCustomizationsV2, id)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to the specified enrollment customization's history.
// URL: POST /api/v2/enrollment-customizations/{id}/history
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-id-history
func (s *Service) AddHistoryNotesV2(ctx context.Context, id string, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointEnrollmentCustomizationsV2, id)

	var result ResponseAddHistoryNotes

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPrestagesV2 retrieves the list of prestages using this enrollment customization.
// URL: GET /api/v2/enrollment-customizations/{id}/prestages
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-prestages
func (s *Service) GetPrestagesV2(ctx context.Context, id string) (*PrestagesResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/prestages", EndpointEnrollmentCustomizationsV2, id)

	var result PrestagesResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadImageV2 uploads an image for enrollment customizations.
// URL: POST /api/v2/enrollment-customizations/images
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-images
func (s *Service) UploadImageV2(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ImageUploadResponse, *interfaces.Response, error) {
	if fileReader == nil {
		return nil, nil, fmt.Errorf("file reader is required")
	}
	if fileName == "" {
		return nil, nil, fmt.Errorf("file name is required")
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ImageUploadResponse

	resp, err := s.client.PostMultipart(ctx, EndpointEnrollmentCustomizationsImagesV2, "file", fileName, fileReader, fileSize, nil, headers, nil, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetImageByIdV2 retrieves an enrollment customization image by ID.
// URL: GET /api/v2/enrollment-customizations/images/{id}
// Returns the image file data as bytes.
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-images-id
func (s *Service) GetImageByIdV2(ctx context.Context, id string) ([]byte, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("image ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointEnrollmentCustomizationsImagesV2, id)

	headers := map[string]string{
		"Accept": "image/*",
	}

	resp, data, err := s.client.GetBytes(ctx, endpoint, nil, headers)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
