package app_request

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AppRequestServiceInterface defines the interface for app request operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
	AppRequestServiceInterface interface {
		// ListFormInputFieldsV1 returns all form input field objects (Search for Form Input Fields).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
		ListFormInputFieldsV1(ctx context.Context, rsqlQuery map[string]string) (*FormInputFieldListResponse, *interfaces.Response, error)

		// ReplaceFormInputFieldsV1 replaces all form input fields (Replace all Form Input Fields).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields
		ReplaceFormInputFieldsV1(ctx context.Context, request []RequestFormInputField) ([]ResourceFormInputField, *interfaces.Response, error)

		// CreateFormInputFieldV1 creates a new form input field record (Create Form Input Field record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-app-request-form-input-fields
		CreateFormInputFieldV1(ctx context.Context, request *RequestFormInputField) (*ResourceFormInputField, *interfaces.Response, error)

		// GetFormInputFieldByIDV1 returns the specified form input field by ID (Get specified Form Input Field object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields-id
		GetFormInputFieldByIDV1(ctx context.Context, id int) (*ResourceFormInputField, *interfaces.Response, error)

		// UpdateFormInputFieldByIDV1 updates the specified form input field by ID (Update specified Form Input Field object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields-id
		UpdateFormInputFieldByIDV1(ctx context.Context, id int, request *RequestFormInputField) (*ResourceFormInputField, *interfaces.Response, error)

		// DeleteFormInputFieldByIDV1 removes the specified form input field by ID (Remove specified Form Input Field record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-app-request-form-input-fields-id
		DeleteFormInputFieldByIDV1(ctx context.Context, id int) (*interfaces.Response, error)

		// GetSettingsV1 retrieves the app request settings (Get Application Request Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-settings
		GetSettingsV1(ctx context.Context) (*ResourceAppRequestSettings, *interfaces.Response, error)

		// UpdateSettingsV1 updates the app request settings (Update Application Request Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-settings
		UpdateSettingsV1(ctx context.Context, request *ResourceAppRequestSettings) (*ResourceAppRequestSettings, *interfaces.Response, error)
	}

	// Service handles communication with the app request-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AppRequestServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Form Input Fields CRUD Operations
// -----------------------------------------------------------------------------

// ListFormInputFieldsV1 returns all form input field objects (Search for Form Input Fields).
// URL: GET /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
func (s *Service) ListFormInputFieldsV1(ctx context.Context, rsqlQuery map[string]string) (*FormInputFieldListResponse, *interfaces.Response, error) {
	var result FormInputFieldListResponse

	endpoint := EndpointFormInputFieldsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceFormInputField
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list form input fields: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// ReplaceFormInputFieldsV1 replaces all form input fields (Replace all Form Input Fields).
// URL: PUT /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields
func (s *Service) ReplaceFormInputFieldsV1(ctx context.Context, request []RequestFormInputField) ([]ResourceFormInputField, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointFormInputFieldsV1

	var result []ResourceFormInputField

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// CreateFormInputFieldV1 creates a new form input field record (Create Form Input Field record).
// URL: POST /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/post_v1-app-request-form-input-fields
func (s *Service) CreateFormInputFieldV1(ctx context.Context, request *RequestFormInputField) (*ResourceFormInputField, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointFormInputFieldsV1

	var result ResourceFormInputField

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

// GetFormInputFieldByIDV1 returns the specified form input field by ID (Get specified Form Input Field object).
// URL: GET /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields-id
func (s *Service) GetFormInputFieldByIDV1(ctx context.Context, id int) (*ResourceFormInputField, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", EndpointFormInputFieldsV1, id)

	var result ResourceFormInputField

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateFormInputFieldByIDV1 updates the specified form input field by ID (Update specified Form Input Field object).
// URL: PUT /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields-id
func (s *Service) UpdateFormInputFieldByIDV1(ctx context.Context, id int, request *RequestFormInputField) (*ResourceFormInputField, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%d", EndpointFormInputFieldsV1, id)

	var result ResourceFormInputField

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

// DeleteFormInputFieldByIDV1 removes the specified form input field by ID (Remove specified Form Input Field record).
// URL: DELETE /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-app-request-form-input-fields-id
func (s *Service) DeleteFormInputFieldByIDV1(ctx context.Context, id int) (*interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", EndpointFormInputFieldsV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// Jamf Pro API - App Request Settings Operations
// -----------------------------------------------------------------------------

// GetSettingsV1 retrieves the app request settings (Get Application Request Settings).
// URL: GET /api/v1/app-request/settings
// https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-settings
func (s *Service) GetSettingsV1(ctx context.Context) (*ResourceAppRequestSettings, *interfaces.Response, error) {
	var result ResourceAppRequestSettings

	endpoint := EndpointSettingsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSettingsV1 updates the app request settings (Update Application Request Settings).
// URL: PUT /api/v1/app-request/settings
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-settings
func (s *Service) UpdateSettingsV1(ctx context.Context, request *ResourceAppRequestSettings) (*ResourceAppRequestSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAppRequestSettings

	endpoint := EndpointSettingsV1

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
