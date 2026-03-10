package enrollment_customizations

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the enrollment customizations-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
	EnrollmentCustomizations struct {
		client client.Client
	}
)

// NewService returns a new enrollment customizations Service backed by the provided HTTP client.
func NewEnrollmentCustomizations(client client.Client) *EnrollmentCustomizations {
	return &EnrollmentCustomizations{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Enrollment Customizations CRUD Operations
// -----------------------------------------------------------------------------

// ListV2 returns a paged list of enrollment customization objects.
// URL: GET /api/v2/enrollment-customizations
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations
func (s *EnrollmentCustomizations) ListV2(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProEnrollmentCustomizationsV2

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEnrollmentCustomization
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	req := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON)

	if rsqlQuery != nil {
		req.SetQueryParams(rsqlQuery)
	}

	resp, err := req.GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV2 returns the specified enrollment customization by ID.
// URL: GET /api/v2/enrollment-customizations/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id
func (s *EnrollmentCustomizations) GetByIDV2(ctx context.Context, id string) (*ResourceEnrollmentCustomization, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	var result ResourceEnrollmentCustomization

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByNameV2 returns the specified enrollment customization by display name.
// This is a convenience method that calls ListV2 and filters by name.
func (s *EnrollmentCustomizations) GetByNameV2(ctx context.Context, name string) (*ResourceEnrollmentCustomization, *resty.Response, error) {
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
func (s *EnrollmentCustomizations) CreateV2(ctx context.Context, request *ResourceEnrollmentCustomization) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProEnrollmentCustomizationsV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV2 replaces the specified enrollment customization by ID.
// URL: PUT /api/v2/enrollment-customizations/{id}
// Returns the full updated enrollment customization resource.
// https://developer.jamf.com/jamf-pro/reference/put_v2-enrollment-customizations-id
func (s *EnrollmentCustomizations) UpdateByIDV2(ctx context.Context, id string, request *ResourceEnrollmentCustomization) (*ResourceEnrollmentCustomization, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.DisplayName == "" {
		return nil, nil, fmt.Errorf("display name is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	var result ResourceEnrollmentCustomization

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV2 removes the specified enrollment customization by ID.
// URL: DELETE /api/v2/enrollment-customizations/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v2-enrollment-customizations-id
func (s *EnrollmentCustomizations) DeleteByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryV2 returns the history object for the specified enrollment customization.
// URL: GET /api/v2/enrollment-customizations/{id}/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-history
func (s *EnrollmentCustomizations) GetHistoryV2(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceHistoryEntry
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	req := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON)

	if rsqlQuery != nil {
		req.SetQueryParams(rsqlQuery)
	}

	resp, err := req.GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to the specified enrollment customization's history.
// URL: POST /api/v2/enrollment-customizations/{id}/history
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-id-history
func (s *EnrollmentCustomizations) AddHistoryNotesV2(ctx context.Context, id string, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	var result ResponseAddHistoryNotes

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetPrestagesV2 retrieves the list of prestages using this enrollment customization.
// URL: GET /api/v2/enrollment-customizations/{id}/prestages
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-id-prestages
func (s *EnrollmentCustomizations) GetPrestagesV2(ctx context.Context, id string) (*PrestagesResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("enrollment customization ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/prestages", constants.EndpointJamfProEnrollmentCustomizationsV2, id)

	var result PrestagesResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UploadImageV2 uploads an image for enrollment customizations.
// URL: POST /api/v2/enrollment-customizations/images
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-customizations-images
func (s *EnrollmentCustomizations) UploadImageV2(ctx context.Context, fileReader io.Reader, fileSize int64, fileName string) (*ImageUploadResponse, *resty.Response, error) {
	if fileReader == nil {
		return nil, nil, fmt.Errorf("file reader is required")
	}
	if fileName == "" {
		return nil, nil, fmt.Errorf("file name is required")
	}

	endpoint := constants.EndpointJamfProEnrollmentCustomizationsImagesV2

	var result ImageUploadResponse

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetMultipartFile("file", fileName, fileReader, fileSize, nil).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetImageByIdV2 retrieves an enrollment customization image by ID.
// URL: GET /api/v2/enrollment-customizations/images/{id}
// Returns the image file data as bytes.
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-customizations-images-id
func (s *EnrollmentCustomizations) GetImageByIdV2(ctx context.Context, id string) ([]byte, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("image ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProEnrollmentCustomizationsImagesV2, id)

	resp, data, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ImageAny).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}
