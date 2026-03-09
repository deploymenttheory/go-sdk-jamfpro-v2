package enrollment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// EnrollmentServiceInterface defines the interface for enrollment operations.
	//
	// Manages enrollment settings, ADUE access groups, and language messaging.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
	EnrollmentServiceInterface interface {
		// GetADUESessionTokenSettingsV1 retrieves ADUE (Account Driven User Enrollment) session token settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
		GetADUESessionTokenSettingsV1(ctx context.Context) (*ResourceADUESessionTokenSettings, *resty.Response, error)

		// UpdateADUESessionTokenSettingsV1 updates ADUE session token settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-adue-session-token-settings
		UpdateADUESessionTokenSettingsV1(ctx context.Context, request *ResourceADUESessionTokenSettings) (*ResourceADUESessionTokenSettings, *resty.Response, error)

		// GetHistoryV2 retrieves enrollment history with optional sorting.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
		GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNotesV2 adds notes to enrollment history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-history
		AddHistoryNotesV2(ctx context.Context, request *RequestAddHistoryNotes) (*CreateResponse, *resty.Response, error)

		// ExportHistoryV2 exports enrollment history in the specified format (JSON or CSV).
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size, export-fields, export-labels).
		// The acceptHeader determines the export format (application/json or text/csv).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-history-export
		ExportHistoryV2(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, request *RequestExportHistory) ([]byte, *resty.Response, error)

		// ListAccessGroupsV3 lists all ADUE access groups with pagination support.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
		ListAccessGroupsV3(ctx context.Context, rsqlQuery map[string]string) (*ListResponseAccessGroups, *resty.Response, error)

		// GetAccessGroupByIDV3 retrieves an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups-id
		GetAccessGroupByIDV3(ctx context.Context, id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *resty.Response, error)

		// CreateAccessGroupV3 creates a new ADUE access group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-access-groups
		CreateAccessGroupV3(ctx context.Context, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*CreateResponse, *resty.Response, error)

		// UpdateAccessGroupByIDV3 updates an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-access-groups-id
		UpdateAccessGroupByIDV3(ctx context.Context, id string, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *resty.Response, error)

		// DeleteAccessGroupByIDV3 deletes an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-access-groups-id
		DeleteAccessGroupByIDV3(ctx context.Context, id string) (*resty.Response, error)

		// ListLanguageMessagesV3 returns all configured enrollment language messages.
		//
		// Automatically fetches all pages. The API supports pagination and sorting but not RSQL filtering.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages
		ListLanguageMessagesV3(ctx context.Context) (*ListResponseLanguageMessages, *resty.Response, error)

		// GetLanguageMessageV3 retrieves enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages-languageid
		GetLanguageMessageV3(ctx context.Context, languageCode string) (*ResourceEnrollmentLanguage, *resty.Response, error)

		// UpdateLanguageMessageV3 updates enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-languages-languageid
		UpdateLanguageMessageV3(ctx context.Context, languageCode string, request *ResourceEnrollmentLanguage) (*ResourceEnrollmentLanguage, *resty.Response, error)

		// DeleteLanguageMessageV3 deletes enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-languages-languageid
		DeleteLanguageMessageV3(ctx context.Context, languageCode string) (*resty.Response, error)

		// DeleteMultipleLanguageMessagesV3 deletes multiple enrollment language messages by their codes.
		//
		// Validates all language codes against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-languages-delete-multiple
		DeleteMultipleLanguageMessagesV3(ctx context.Context, request *RequestDeleteMultipleLanguages) (*resty.Response, error)

		// ListLanguageCodesV3 retrieves the list of available languages and their ISO 639-1 codes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-language-codes
		ListLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error)

		// ListFilteredLanguageCodesV3 returns language codes not yet added to enrollment.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-filtered-language-codes
		ListFilteredLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error)

		// GetV4 retrieves the current enrollment configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
		GetV4(ctx context.Context) (*ResourceEnrollment, *resty.Response, error)

		// UpdateV4 updates the enrollment configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
		UpdateV4(ctx context.Context, request *ResourceEnrollment) (*ResourceEnrollment, *resty.Response, error)
	}

	// Service handles communication with the enrollment-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
	Enrollment struct {
		client interfaces.HTTPClient
	}
)

var _ EnrollmentServiceInterface = (*Enrollment)(nil)

func NewEnrollment(client interfaces.HTTPClient) *Enrollment {
	return &Enrollment{client: client}
}

// GetADUESessionTokenSettingsV1 retrieves ADUE session token settings.
// URL: GET /api/v1/adue-session-token-settings
// https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
func (s *Enrollment) GetADUESessionTokenSettingsV1(ctx context.Context) (*ResourceADUESessionTokenSettings, *resty.Response, error) {
	endpoint := EndpointADUESessionTokenSettingsV1

	var result ResourceADUESessionTokenSettings

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := EndpointADUESessionTokenSettingsV1

	var result ResourceADUESessionTokenSettings

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

// GetHistoryV2 retrieves enrollment history with optional sorting.
// URL: GET /api/v2/enrollment/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
func (s *Enrollment) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", EndpointEnrollmentV2)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryObject
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
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

	endpoint := fmt.Sprintf("%s/history", EndpointEnrollmentV2)

	var result CreateResponse

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

// ExportHistoryV2 exports enrollment history in the specified format (JSON or CSV).
// URL: POST /api/v2/enrollment/history/export
// acceptHeader should be "application/json" or "text/csv".
// rsqlQuery supports: filter (RSQL), sort, page, page-size, export-fields, export-labels (all optional).
// https://developer.jamf.com/jamf-pro/reference/post_v2-enrollment-history-export
func (s *Enrollment) ExportHistoryV2(ctx context.Context, acceptHeader string, rsqlQuery map[string]string, request *RequestExportHistory) ([]byte, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history/export", EndpointEnrollmentV2)

	if acceptHeader == "" {
		acceptHeader = mime.ApplicationJSON
	}

	headers := map[string]string{
		"Accept":       acceptHeader,
		"Content-Type": mime.ApplicationJSON,
	}

	var body any
	if request != nil {
		body = request
	}

	resp, err := s.client.PostWithQuery(ctx, endpoint, rsqlQuery, body, headers, nil)
	if err != nil {
		return nil, resp, err
	}

	return resp.Bytes(), resp, nil
}

// ListAccessGroupsV3 lists all ADUE access groups with pagination support.
// URL: GET /api/v3/enrollment/access-groups
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
func (s *Enrollment) ListAccessGroupsV3(ctx context.Context, rsqlQuery map[string]string) (*ListResponseAccessGroups, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/access-groups", EndpointEnrollmentV3)

	var result ListResponseAccessGroups

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceAccountDrivenUserEnrollmentAccessGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
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

	endpoint := fmt.Sprintf("%s/access-groups/%s", EndpointEnrollmentV3, id)

	var result ResourceAccountDrivenUserEnrollmentAccessGroup

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/access-groups", EndpointEnrollmentV3)

	var result CreateResponse

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

	endpoint := fmt.Sprintf("%s/access-groups/%s", EndpointEnrollmentV3, id)

	var result ResourceAccountDrivenUserEnrollmentAccessGroup

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

// DeleteAccessGroupByIDV3 deletes an ADUE access group by ID.
// URL: DELETE /api/v3/enrollment/access-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-access-groups-id
func (s *Enrollment) DeleteAccessGroupByIDV3(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("access group ID is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", EndpointEnrollmentV3, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListLanguageMessagesV3 returns all configured enrollment language messages.
// URL: GET /api/v3/enrollment/languages
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages
func (s *Enrollment) ListLanguageMessagesV3(ctx context.Context) (*ListResponseLanguageMessages, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/languages", EndpointEnrollmentV3)

	var result ListResponseLanguageMessages

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceEnrollmentLanguage
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, nil, headers, mergePage)
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

	endpoint := fmt.Sprintf("%s/languages/%s", EndpointEnrollmentV3, languageCode)

	var result ResourceEnrollmentLanguage

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err = s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := fmt.Sprintf("%s/languages/%s", EndpointEnrollmentV3, languageCode)

	var result ResourceEnrollmentLanguage

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err = s.client.Put(ctx, endpoint, request, headers, &result)
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

	endpoint := fmt.Sprintf("%s/languages/%s", EndpointEnrollmentV3, languageCode)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err = s.client.Delete(ctx, endpoint, nil, headers, nil)
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

	endpoint := fmt.Sprintf("%s/languages/delete-multiple", EndpointEnrollmentV3)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err = s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// ListLanguageCodesV3 retrieves the list of available languages and their ISO 639-1 codes.
// URL: GET /api/v3/enrollment/language-codes
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-language-codes
func (s *Enrollment) ListLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/language-codes", EndpointEnrollmentV3)

	var result []ResourceLanguageCode

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// ListFilteredLanguageCodesV3 returns language codes not yet added to enrollment.
// URL: GET /api/v3/enrollment/filtered-language-codes
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-filtered-language-codes
func (s *Enrollment) ListFilteredLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/filtered-language-codes", EndpointEnrollmentV3)

	var result []ResourceLanguageCode

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetV4 retrieves the current enrollment configuration.
// URL: GET /api/v4/enrollment
// https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
func (s *Enrollment) GetV4(ctx context.Context) (*ResourceEnrollment, *resty.Response, error) {
	endpoint := EndpointEnrollmentV4

	var result ResourceEnrollment

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	endpoint := EndpointEnrollmentV4

	var result ResourceEnrollment

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
