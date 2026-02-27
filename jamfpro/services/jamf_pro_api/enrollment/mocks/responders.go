package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// EnrollmentMock is a test double implementing interfaces.HTTPClient.
type EnrollmentMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewEnrollmentMock returns an empty mock ready for response registration.
func NewEnrollmentMock() *EnrollmentMock {
	return &EnrollmentMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *EnrollmentMock) RegisterMocks() {
	m.RegisterGetADUESessionTokenSettingsV1Mock()
	m.RegisterUpdateADUESessionTokenSettingsV1Mock()
	m.RegisterGetHistoryV2Mock()
	m.RegisterAddHistoryNotesV2Mock()
	m.RegisterExportHistoryV2Mock()
	m.RegisterListAccessGroupsV3Mock()
	m.RegisterGetAccessGroupByIDV3Mock()
	m.RegisterCreateAccessGroupV3Mock()
	m.RegisterUpdateAccessGroupByIDV3Mock()
	m.RegisterDeleteAccessGroupByIDV3Mock()
	m.RegisterListLanguageMessagesV3Mock()
	m.RegisterListLanguageCodesV3Mock()
	m.RegisterListFilteredLanguageCodesV3Mock()
	m.RegisterGetLanguageMessageV3Mock()
	m.RegisterUpdateLanguageMessageV3Mock()
	m.RegisterDeleteLanguageMessageV3Mock()
	m.RegisterDeleteMultipleLanguageMessagesV3Mock()
	m.RegisterGetV4Mock()
	m.RegisterUpdateV4Mock()
}

func (m *EnrollmentMock) register(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *EnrollmentMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("EnrollmentMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("EnrollmentMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *EnrollmentMock) RegisterGetADUESessionTokenSettingsV1Mock() {
	body := []byte(`{
		"enabled": false,
		"expirationIntervalDays": 1,
		"expirationIntervalSeconds": 86400
	}`)
	m.register("GET", "/api/v1/adue-session-token-settings", 200, body)
}

func (m *EnrollmentMock) RegisterUpdateADUESessionTokenSettingsV1Mock() {
	body := []byte(`{
		"enabled": true,
		"expirationIntervalDays": 1,
		"expirationIntervalSeconds": 86400
	}`)
	m.register("PUT", "/api/v1/adue-session-token-settings", 200, body)
}

func (m *EnrollmentMock) RegisterGetHistoryV2Mock() {
	body := []byte(`{
		"totalCount": 2,
		"results": [
			{
				"id": 1,
				"username": "admin",
				"date": "2024-01-15T10:30:00Z",
				"note": "Device enrolled",
				"details": "iPad enrolled via DEP"
			},
			{
				"id": 2,
				"username": "admin",
				"date": "2024-01-16T14:20:00Z",
				"note": "Device re-enrolled",
				"details": "MacBook re-enrolled after reset"
			}
		]
	}`)
	m.register("GET", "/api/v2/enrollment/history", 200, body)
}

func (m *EnrollmentMock) RegisterAddHistoryNotesV2Mock() {
	body := []byte(`{
		"id": "1",
		"href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
	}`)
	m.register("POST", "/api/v2/enrollment/history", 201, body)
}

func (m *EnrollmentMock) RegisterExportHistoryV2Mock() {
	body := []byte("Username,DATE,NOTES,Details\nadmin,2022-02-04T11:56:26.343Z,Edited,Re-enrollment Restricted true")
	m.register("POST", "/api/v2/enrollment/history/export", 200, body)
}

func (m *EnrollmentMock) RegisterListAccessGroupsV3Mock() {
	body := []byte(`{
		"totalCount": 1,
		"results": [
			{
				"id": "1",
				"groupId": "100",
				"ldapServerId": "1",
				"name": "Test Access Group",
				"siteId": "1",
				"enterpriseEnrollmentEnabled": true,
				"personalEnrollmentEnabled": false,
				"accountDrivenUserEnrollmentEnabled": true,
				"requireEula": false
			}
		]
	}`)
	m.register("GET", "/api/v3/enrollment/access-groups", 200, body)
}

func (m *EnrollmentMock) RegisterGetAccessGroupByIDV3Mock() {
	body := []byte(`{
		"id": "1",
		"groupId": "100",
		"ldapServerId": "1",
		"name": "Test Access Group",
		"siteId": "1",
		"enterpriseEnrollmentEnabled": true,
		"personalEnrollmentEnabled": false,
		"accountDrivenUserEnrollmentEnabled": true,
		"requireEula": false
	}`)
	m.register("GET", "/api/v3/enrollment/access-groups/1", 200, body)
}

func (m *EnrollmentMock) RegisterCreateAccessGroupV3Mock() {
	body := []byte(`{
		"id": "2",
		"href": "/api/v3/enrollment/access-groups/2"
	}`)
	m.register("POST", "/api/v3/enrollment/access-groups", 201, body)
}

func (m *EnrollmentMock) RegisterUpdateAccessGroupByIDV3Mock() {
	body := []byte(`{
		"id": "1",
		"groupId": "100",
		"ldapServerId": "1",
		"name": "Updated Access Group",
		"siteId": "1",
		"enterpriseEnrollmentEnabled": true,
		"personalEnrollmentEnabled": true,
		"accountDrivenUserEnrollmentEnabled": true,
		"requireEula": true
	}`)
	m.register("PUT", "/api/v3/enrollment/access-groups/1", 200, body)
}

func (m *EnrollmentMock) RegisterDeleteAccessGroupByIDV3Mock() {
	m.register("DELETE", "/api/v3/enrollment/access-groups/1", 204, []byte{})
}

func (m *EnrollmentMock) RegisterListLanguageMessagesV3Mock() {
	body := []byte(`{
		"totalCount": 1,
		"results": [
			{
				"languageCode": "en",
				"name": "English",
				"title": "Enrollment",
				"loginDescription": "Please log in to enroll",
				"username": "Username",
				"password": "Password",
				"loginButton": "Login"
			}
		]
	}`)
	m.register("GET", "/api/v3/enrollment/languages", 200, body)
}

func (m *EnrollmentMock) RegisterListLanguageCodesV3Mock() {
	body := []byte(`[
		{"value": "en", "name": "English"},
		{"value": "es", "name": "Spanish"},
		{"value": "fr", "name": "French"}
	]`)
	m.register("GET", "/api/v3/enrollment/language-codes", 200, body)
}

func (m *EnrollmentMock) RegisterListFilteredLanguageCodesV3Mock() {
	body := []byte(`[
		{"value": "de", "name": "German"},
		{"value": "ja", "name": "Japanese"}
	]`)
	m.register("GET", "/api/v3/enrollment/filtered-language-codes", 200, body)
}

func (m *EnrollmentMock) RegisterGetLanguageMessageV3Mock() {
	body := []byte(`{
		"languageCode": "en",
		"name": "English",
		"title": "Enrollment",
		"loginDescription": "Please log in to enroll",
		"username": "Username",
		"password": "Password",
		"loginButton": "Login",
		"deviceClassDescription": "Select device type",
		"deviceClassPersonal": "Personal",
		"deviceClassPersonalDescription": "Personal device",
		"deviceClassEnterprise": "Enterprise",
		"deviceClassEnterpriseDescription": "Company device",
		"deviceClassButton": "Continue"
	}`)
	m.register("GET", "/api/v3/enrollment/languages/en", 200, body)
}

func (m *EnrollmentMock) RegisterUpdateLanguageMessageV3Mock() {
	body := []byte(`{
		"languageCode": "en",
		"name": "English",
		"title": "Updated Enrollment",
		"loginDescription": "Updated description",
		"username": "Username",
		"password": "Password",
		"loginButton": "Login"
	}`)
	m.register("PUT", "/api/v3/enrollment/languages/en", 200, body)
}

func (m *EnrollmentMock) RegisterDeleteLanguageMessageV3Mock() {
	m.register("DELETE", "/api/v3/enrollment/languages/en", 204, []byte{})
}

func (m *EnrollmentMock) RegisterDeleteMultipleLanguageMessagesV3Mock() {
	m.register("POST", "/api/v3/enrollment/languages/delete-multiple", 204, []byte{})
}

func (m *EnrollmentMock) RegisterGetV4Mock() {
	body := []byte(`{
		"installSingleProfile": true,
		"signingMdmProfileEnabled": false,
		"mdmSigningCertificate": null,
		"restrictReenrollment": false,
		"flushLocationInformation": true,
		"flushLocationHistoryInformation": false,
		"flushPolicyHistory": false,
		"flushExtensionAttributes": false,
		"flushSoftwareUpdatePlans": false,
		"macOsEnterpriseEnrollmentEnabled": true,
		"managementUsername": "admin",
		"createManagementAccount": true,
		"hideManagementAccount": false,
		"allowSshOnlyManagementAccount": false,
		"ensureSshRunning": false,
		"launchSelfService": true,
		"signQuickAdd": false,
		"iosEnterpriseEnrollmentEnabled": true,
		"iosPersonalEnrollmentEnabled": false,
		"accountDrivenUserEnrollmentEnabled": false,
		"accountDrivenDeviceIosEnrollmentEnabled": false,
		"accountDrivenDeviceMacosEnrollmentEnabled": false,
		"accountDrivenUserVisionosEnrollmentEnabled": false,
		"accountDrivenDeviceVisionosEnrollmentEnabled": false,
		"maidUsernameMergeEnabled": false
	}`)
	m.register("GET", "/api/v4/enrollment", 200, body)
}

func (m *EnrollmentMock) RegisterUpdateV4Mock() {
	body := []byte(`{
		"installSingleProfile": false,
		"signingMdmProfileEnabled": false,
		"restrictReenrollment": true,
		"flushLocationInformation": true,
		"macOsEnterpriseEnrollmentEnabled": true,
		"managementUsername": "admin",
		"createManagementAccount": true
	}`)
	m.register("PUT", "/api/v4/enrollment", 200, body)
}

func (m *EnrollmentMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *EnrollmentMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *EnrollmentMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *EnrollmentMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EnrollmentMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EnrollmentMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *EnrollmentMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		if err := mergePage(resp.Body); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *EnrollmentMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *EnrollmentMock) InvalidateToken() error                    { return nil }
func (m *EnrollmentMock) KeepAliveToken() error                     { return nil }
func (m *EnrollmentMock) GetLogger() *zap.Logger                    { return m.logger }
