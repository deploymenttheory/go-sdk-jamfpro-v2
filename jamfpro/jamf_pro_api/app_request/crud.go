package app_request

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the app request-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
	AppRequest struct {
		client client.Client
	}
)

func NewAppRequest(client client.Client) *AppRequest {
	return &AppRequest{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Form Input Fields CRUD Operations
// -----------------------------------------------------------------------------

// ListFormInputFieldsV1 returns all form input field objects (Search for Form Input Fields).
// URL: GET /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields
func (s *AppRequest) ListFormInputFieldsV1(ctx context.Context, rsqlQuery map[string]string) (*FormInputFieldListResponse, *resty.Response, error) {
	var result FormInputFieldListResponse

	endpoint := constants.EndpointJamfProFormInputFieldsV1

	mergePage := func(pageData []byte) error {
		var items []ResourceFormInputField
		if err := json.Unmarshal(pageData, &items); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, items...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

	if err != nil {
		return nil, resp, fmt.Errorf("failed to list form input fields: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// ReplaceFormInputFieldsV1 replaces all form input fields (Replace all Form Input Fields).
// URL: PUT /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields
func (s *AppRequest) ReplaceFormInputFieldsV1(ctx context.Context, request []RequestFormInputField) ([]ResourceFormInputField, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProFormInputFieldsV1

	var result []ResourceFormInputField

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// CreateFormInputFieldV1 creates a new form input field record (Create Form Input Field record).
// URL: POST /api/v1/app-request/form-input-fields
// https://developer.jamf.com/jamf-pro/reference/post_v1-app-request-form-input-fields
func (s *AppRequest) CreateFormInputFieldV1(ctx context.Context, request *RequestFormInputField) (*ResourceFormInputField, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProFormInputFieldsV1

	var result ResourceFormInputField

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

// GetFormInputFieldByIDV1 returns the specified form input field by ID (Get specified Form Input Field object).
// URL: GET /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-app-request-form-input-fields-id
func (s *AppRequest) GetFormInputFieldByIDV1(ctx context.Context, id int) (*ResourceFormInputField, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", constants.EndpointJamfProFormInputFieldsV1, id)

	var result ResourceFormInputField

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateFormInputFieldByIDV1 updates the specified form input field by ID (Update specified Form Input Field object).
// URL: PUT /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-form-input-fields-id
func (s *AppRequest) UpdateFormInputFieldByIDV1(ctx context.Context, id int, request *RequestFormInputField) (*ResourceFormInputField, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%d", constants.EndpointJamfProFormInputFieldsV1, id)

	var result ResourceFormInputField

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

// DeleteFormInputFieldByIDV1 removes the specified form input field by ID (Remove specified Form Input Field record).
// URL: DELETE /api/v1/app-request/form-input-fields/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-app-request-form-input-fields-id
func (s *AppRequest) DeleteFormInputFieldByIDV1(ctx context.Context, id int) (*resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%d", constants.EndpointJamfProFormInputFieldsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)

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
func (s *AppRequest) GetSettingsV1(ctx context.Context) (*ResourceAppRequestSettings, *resty.Response, error) {
	var result ResourceAppRequestSettings

	endpoint := constants.EndpointJamfProAppRequestSettingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSettingsV1 updates the app request settings (Update Application Request Settings).
// URL: PUT /api/v1/app-request/settings
// https://developer.jamf.com/jamf-pro/reference/put_v1-app-request-settings
func (s *AppRequest) UpdateSettingsV1(ctx context.Context, request *ResourceAppRequestSettings) (*ResourceAppRequestSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAppRequestSettings

	endpoint := constants.EndpointJamfProAppRequestSettingsV1

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
