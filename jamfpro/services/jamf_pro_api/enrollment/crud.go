package enrollment

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// EnrollmentServiceInterface defines the interface for enrollment operations.
	//
	// Manages enrollment settings, ADUE access groups, and language messaging.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
	EnrollmentServiceInterface interface {
		// GetHistoryV2 retrieves enrollment history with optional sorting.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
		GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// ListAccessGroupsV3 lists all ADUE access groups with pagination support.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
		ListAccessGroupsV3(ctx context.Context, rsqlQuery map[string]string) (*ListResponseAccessGroups, *interfaces.Response, error)

		// GetAccessGroupByIDV3 retrieves an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups-id
		GetAccessGroupByIDV3(ctx context.Context, id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *interfaces.Response, error)

		// CreateAccessGroupV3 creates a new ADUE access group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-access-groups
		CreateAccessGroupV3(ctx context.Context, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*CreateResponse, *interfaces.Response, error)

		// UpdateAccessGroupByIDV3 updates an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-access-groups-id
		UpdateAccessGroupByIDV3(ctx context.Context, id string, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *interfaces.Response, error)

		// DeleteAccessGroupByIDV3 deletes an ADUE access group by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-access-groups-id
		DeleteAccessGroupByIDV3(ctx context.Context, id string) (*interfaces.Response, error)

		// ListLanguageMessagesV3 returns all configured enrollment language messages.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages
		ListLanguageMessagesV3(ctx context.Context) ([]ResourceEnrollmentLanguage, *interfaces.Response, error)

		// GetLanguageMessageV3 retrieves enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages-languageid
		GetLanguageMessageV3(ctx context.Context, languageCode string) (*ResourceEnrollmentLanguage, *interfaces.Response, error)

		// UpdateLanguageMessageV3 updates enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-languages-languageid
		UpdateLanguageMessageV3(ctx context.Context, languageCode string, request *ResourceEnrollmentLanguage) (*ResourceEnrollmentLanguage, *interfaces.Response, error)

		// DeleteLanguageMessageV3 deletes enrollment messaging for a specific language code.
		//
		// Validates the language code against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v3-enrollment-languages-languageid
		DeleteLanguageMessageV3(ctx context.Context, languageCode string) (*interfaces.Response, error)

		// DeleteMultipleLanguageMessagesV3 deletes multiple enrollment language messages by their codes.
		//
		// Validates all language codes against available codes before making the request.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-languages-delete-multiple
		DeleteMultipleLanguageMessagesV3(ctx context.Context, request *RequestDeleteMultipleLanguages) (*interfaces.Response, error)

		// ListLanguageCodesV3 retrieves the list of available languages and their ISO 639-1 codes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-language-codes
		ListLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *interfaces.Response, error)

		// GetV4 retrieves the current enrollment configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v4-enrollment
		GetV4(ctx context.Context) (*ResourceEnrollment, *interfaces.Response, error)

		// UpdateV4 updates the enrollment configuration.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
		UpdateV4(ctx context.Context, request *ResourceEnrollment) (*ResourceEnrollment, *interfaces.Response, error)
	}

	// Service handles communication with the enrollment-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ EnrollmentServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetHistoryV2 retrieves enrollment history with optional sorting.
// URL: GET /api/v2/enrollment/history
// https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
func (s *Service) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/history", EndpointEnrollmentV2)

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

// ListAccessGroupsV3 lists all ADUE access groups with pagination support.
// URL: GET /api/v3/enrollment/access-groups
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups
func (s *Service) ListAccessGroupsV3(ctx context.Context, rsqlQuery map[string]string) (*ListResponseAccessGroups, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/access-groups", EndpointEnrollmentV3)

	var result ListResponseAccessGroups

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

// GetAccessGroupByIDV3 retrieves an ADUE access group by ID.
// URL: GET /api/v3/enrollment/access-groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-access-groups-id
func (s *Service) GetAccessGroupByIDV3(ctx context.Context, id string) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("access group ID is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", EndpointEnrollmentV3, id)

	var result ResourceAccountDrivenUserEnrollmentAccessGroup

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

// CreateAccessGroupV3 creates a new ADUE access group.
// URL: POST /api/v3/enrollment/access-groups
// https://developer.jamf.com/jamf-pro/reference/post_v3-enrollment-access-groups
func (s *Service) CreateAccessGroupV3(ctx context.Context, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*CreateResponse, *interfaces.Response, error) {
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
func (s *Service) UpdateAccessGroupByIDV3(ctx context.Context, id string, request *ResourceAccountDrivenUserEnrollmentAccessGroup) (*ResourceAccountDrivenUserEnrollmentAccessGroup, *interfaces.Response, error) {
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
func (s *Service) DeleteAccessGroupByIDV3(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("access group ID is required")
	}

	endpoint := fmt.Sprintf("%s/access-groups/%s", EndpointEnrollmentV3, id)

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

// ListLanguageMessagesV3 returns all configured enrollment language messages.
// URL: GET /api/v3/enrollment/languages
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages
func (s *Service) ListLanguageMessagesV3(ctx context.Context) ([]ResourceEnrollmentLanguage, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/languages", EndpointEnrollmentV3)

	type response struct {
		TotalCount int                          `json:"totalCount"`
		Results    []ResourceEnrollmentLanguage `json:"results"`
	}

	var result response

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result.Results, resp, nil
}

// GetLanguageMessageV3 retrieves enrollment messaging for a specific language code.
// URL: GET /api/v3/enrollment/languages/{languageCode}
// https://developer.jamf.com/jamf-pro/reference/get_v3-enrollment-languages-languageid
func (s *Service) GetLanguageMessageV3(ctx context.Context, languageCode string) (*ResourceEnrollmentLanguage, *interfaces.Response, error) {
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

	resp, err = s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateLanguageMessageV3 updates enrollment messaging for a specific language code.
// URL: PUT /api/v3/enrollment/languages/{languageCode}
// https://developer.jamf.com/jamf-pro/reference/put_v3-enrollment-languages-languageid
func (s *Service) UpdateLanguageMessageV3(ctx context.Context, languageCode string, request *ResourceEnrollmentLanguage) (*ResourceEnrollmentLanguage, *interfaces.Response, error) {
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
func (s *Service) DeleteLanguageMessageV3(ctx context.Context, languageCode string) (*interfaces.Response, error) {
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
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) DeleteMultipleLanguageMessagesV3(ctx context.Context, request *RequestDeleteMultipleLanguages) (*interfaces.Response, error) {
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
func (s *Service) ListLanguageCodesV3(ctx context.Context) ([]ResourceLanguageCode, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/language-codes", EndpointEnrollmentV3)

	var result []ResourceLanguageCode

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) GetV4(ctx context.Context) (*ResourceEnrollment, *interfaces.Response, error) {
	endpoint := EndpointEnrollmentV4

	var result ResourceEnrollment

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

// UpdateV4 updates the enrollment configuration.
// URL: PUT /api/v4/enrollment
// https://developer.jamf.com/jamf-pro/reference/put_v4-enrollment
func (s *Service) UpdateV4(ctx context.Context, request *ResourceEnrollment) (*ResourceEnrollment, *interfaces.Response, error) {
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
