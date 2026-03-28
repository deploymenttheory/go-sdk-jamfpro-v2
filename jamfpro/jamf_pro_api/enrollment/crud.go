package enrollment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the enrollment-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
	Enrollment struct {
		client client.Client
	}
)

func NewEnrollment(client client.Client) *Enrollment {
	return &Enrollment{client: client}
}

// GetADUESessionTokenSettingsV1 retrieves ADUE session token settings.
// URL: GET /api/v1/adue-session-token-settings
// https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
func (s *Enrollment) GetADUESessionTokenSettingsV1(ctx context.Context) (*ResourceADUESessionTokenSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProADUESessionTokenSettingsV1

	var result ResourceADUESessionTokenSettings

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateADUESessionTokenSettingsV1 updates ADUE session token settings.
// URL: PUT /api/v1/adue-session-token-settings
// https://developer.jamf.com/jamf-pro/reference/put_v1-adue-session-token-settings
func (s *Enrollment) UpdateADUESessionTokenSettingsV1(ctx context.Context, request *ResourceADUESessionTokenSettings) (*ResourceADUESessionTokenSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProADUESessionTokenSettingsV1

	var result ResourceADUESessionTokenSettings

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

// GetHistoryV2 retrieves enrollment history with optional sorting.
// URL: GET /api/v2/enrollment/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
func (s *Enrollment) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEnrollmentV2)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to enrollment history.
// URL: POST /api/v2/enrollment/history
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-history
func (s *Enrollment) AddHistoryNotesV2(ctx context.Context, request *RequestAddHistoryNotes) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if request.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEnrollmentV2)

	var result CreateResponse

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

// ExportHistoryV2 exports enrollment history in the specified format (JSON or CSV).
// URL: POST /api/v2/enrollment/history/export
// acceptHeader should be "application/json" or "text/csv".
// rsqlQuery supports: filter (RSQL), sort, page, page-size, export-fields, export-labels (all optional).
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-history-export
func (s *Enrollment) ExportHistoryV2(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, request *RequestExportHistory) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history/export", constants.EndpointJamfProEnrollmentV2)

	if acceptHeader == "" {
		acceptHeader = constants.ApplicationJSON
	}

	req := s.client.NewRequest(ctx).
		SetHeader("Accept", acceptHeader).
		SetHeader("Content-Type", constants.ApplicationJSON)

	if rsqlQuery != nil {
		req = req.SetQueryParams(rsqlQuery)
	}

	if request != nil {
		req = req.SetBody(request)
	}

	resp, err := req.Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return resp.Bytes(), resp, nil
}

// ListAccessGroupsV3 lists all ADUE access groups with pagination support.
// URL: GET /api/v3/enrollment/access-groups
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
func (s *Enrollment) ListAccessGroupsV3(ctx context.Context, rsqlQuery map[string]string) (*ListResponseAccessGroups, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/access-groups", constants.EndpointJamfProEnrollmentV3)

	var result ListResponseAccessGroups

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceAccountDrivenUserEnrollmentAccessGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, err
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetAccessGroupByIDV3 retrieves an ADUE access group by ID.
// URL: GET /api/v3/enrollment/access-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups-id
func (s *Enrollment) GetAccessGroupByIDV3(ctx context.Context, id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("access group ID is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", constants.EndpointJamfProEnrollmentV3, id)

	var result ResourceAccountDrivenUserEnrollmentAccessGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateAccessGroupV3 creates a new ADUE access group.
// URL: POST /api/v3/enrollment/access-groups
// https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-access-groups
func (s *Enrollment) CreateAccessGroupV3(ctx context.Context, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups", constants.EndpointJamfProEnrollmentV3)

	var result CreateResponse

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

// UpdateAccessGroupByIDV3 updates an ADUE access group by ID.
// URL: PUT /api/v3/enrollment/access-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-access-groups-id
func (s *Enrollment) UpdateAccessGroupByIDV3(ctx context.Context, id string, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("access group ID is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", constants.EndpointJamfProEnrollmentV3, id)

	var result ResourceAccountDrivenUserEnrollmentAccessGroup

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

// DeleteAccessGroupByIDV3 deletes an ADUE access group by ID.
// URL: DELETE /api/v3/enrollment/access-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-access-groups-id
func (s *Enrollment) DeleteAccessGroupByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("access group ID is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", constants.EndpointJamfProEnrollmentV3, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListLanguageMessagesV3 returns all configured enrollment language messages.
// URL: GET /api/v3/enrollment/languages
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages
func (s *Enrollment) ListLanguageMessagesV3(ctx context.Context) (*ListResponseLanguageMessages, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/languages", constants.EndpointJamfProEnrollmentV3)

	var result ListResponseLanguageMessages

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEnrollmentLanguage
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list enrollment language messages: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetLanguageMessageV3 retrieves enrollment messaging for a specific language code.
// URL: GET /api/v3/enrollment/languages/{languageCode}
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages-languageid
func (s *Enrollment) GetLanguageMessageV3(ctx context.Context, languageCode string) (*ResourceEnrollmentLanguage, *resty.Response, error) {
	// Retrieve available language codes for validation
	languageCodes, resp, err := s.ListLanguageCodesV3(ctx)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to retrieve language codes for validation: %w", err)
	}

	// Validate the language code
	if err := ValidateLanguageCode(languageCode, languageCodes); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/languages/%s", constants.EndpointJamfProEnrollmentV3, languageCode)

	var result ResourceEnrollmentLanguage

	resp, err = s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateLanguageMessageV3 updates enrollment messaging for a specific language code.
// URL: PUT /api/v3/enrollment/languages/{languageCode}
// https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-languages-languageid
func (s *Enrollment) UpdateLanguageMessageV3(ctx context.Context, languageCode string, request *ResourceEnrollmentLanguage) (*ResourceEnrollmentLanguage, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	// Retrieve available language codes for validation
	languageCodes, resp, err := s.ListLanguageCodesV3(ctx)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to retrieve language codes for validation: %w", err)
	}

	// Validate the language code
	if err := ValidateLanguageCode(languageCode, languageCodes); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/languages/%s", constants.EndpointJamfProEnrollmentV3, languageCode)

	var result ResourceEnrollmentLanguage

	resp, err = s.client.NewRequest(ctx).
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

// DeleteLanguageMessageV3 deletes enrollment messaging for a specific language code.
// URL: DELETE /api/v3/enrollment/languages/{languageCode}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-languages-languageid
func (s *Enrollment) DeleteLanguageMessageV3(ctx context.Context, languageCode string) (*resty.Response, error) {
	// Retrieve available language codes for validation
	languageCodes, resp, err := s.ListLanguageCodesV3(ctx)
	if err != nil {
		return resp, fmt.Errorf("failed to retrieve language codes for validation: %w", err)
	}

	// Validate the language code
	if err := ValidateLanguageCode(languageCode, languageCodes); err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/languages/%s", constants.EndpointJamfProEnrollmentV3, languageCode)

	resp, err = s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteMultipleLanguageMessagesV3 deletes multiple enrollment language messages by their codes.
// URL: POST /api/v3/enrollment/languages/delete-multiple
// https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-languages-delete-multiple
func (s *Enrollment) DeleteMultipleLanguageMessagesV3(ctx context.Context, request *RequestDeleteMultipleLanguages) (*resty.Response, error) {
	if request == nil || len(request.IDs) == 0 {
		return nil, fmt.Errorf("request with at least one language code is required")
	}

	// Retrieve available language codes for validation
	languageCodes, resp, err := s.ListLanguageCodesV3(ctx)
	if err != nil {
		return resp, fmt.Errorf("failed to retrieve language codes for validation: %w", err)
	}

	// Validate all language codes
	for _, code := range request.IDs {
		if err := ValidateLanguageCode(code, languageCodes); err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("%s/languages/delete-multiple", constants.EndpointJamfProEnrollmentV3)

	resp, err = s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListLanguageCodesV3 retrieves the list of available languages and their ISO 639-1 codes.
// URL: GET /api/v3/enrollment/language-codes
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-language-codes
func (s *Enrollment) ListLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/language-codes", constants.EndpointJamfProEnrollmentV3)

	var result []ResourceLanguageCode

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// ListFilteredLanguageCodesV3 returns language codes not yet added to enrollment.
// URL: GET /api/v3/enrollment/filtered-language-codes
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-filtered-language-codes
func (s *Enrollment) ListFilteredLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/filtered-language-codes", constants.EndpointJamfProEnrollmentV3)

	var result []ResourceLanguageCode

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetV4 retrieves the current enrollment configuration.
// URL: GET /api/v4/enrollment
// https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
func (s *Enrollment) GetV4(ctx context.Context) (*ResourceEnrollment, *resty.Response, error) {
	endpoint := constants.EndpointJamfProEnrollmentV4

	var result ResourceEnrollment

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV4 updates the enrollment configuration.
// URL: PUT /api/v4/enrollment
// https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
func (s *Enrollment) UpdateV4(ctx context.Context, request *ResourceEnrollment) (*ResourceEnrollment, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.FlushMdmCommandsOnReenroll != "" {
		if _, ok := validFlushMdmCommandsOnReenroll[request.FlushMdmCommandsOnReenroll]; !ok {
			return nil, nil, fmt.Errorf("invalid flushMdmCommandsOnReenroll %q: must be one of DELETE_NOTHING, DELETE_ERRORS, DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED, DELETE_EVERYTHING", request.FlushMdmCommandsOnReenroll)
		}
	}

	endpoint := constants.EndpointJamfProEnrollmentV4

	var result ResourceEnrollment

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
