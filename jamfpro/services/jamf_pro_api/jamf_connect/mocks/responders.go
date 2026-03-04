package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// JamfConnectMock is a test double implementing interfaces.HTTPClient.
type JamfConnectMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewJamfConnectMock returns an empty mock ready for response registration.
func NewJamfConnectMock() *JamfConnectMock {
	return &JamfConnectMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *JamfConnectMock) RegisterMocks() {
	m.RegisterGetSettingsMock()
	m.RegisterListConfigProfilesMock()
	m.RegisterUpdateConfigProfileMock()
	m.RegisterRetryDeploymentTasksMock()
}

func (m *JamfConnectMock) register(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *JamfConnectMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {"application/json"}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`)), fmt.Errorf("JamfConnectMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("JamfConnectMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *JamfConnectMock) RegisterGetSettingsMock() {
	body := []byte(`{
		"id": "1",
		"displayName": "Jamf Connect Settings",
		"description": "Test settings",
		"enabled": true,
		"settings": "test-settings",
		"version": "2.0.0",
		"lastModified": "2024-01-15T10:30:00Z",
		"lastModifiedBy": "admin"
	}`)
	m.register("GET", "/api/v1/jamf-connect", 200, body)
}

func (m *JamfConnectMock) RegisterListConfigProfilesMock() {
	body := []byte(`{
		"totalCount": 2,
		"results": [
			{
				"uuid": "123e4567-e89b-12d3-a456-426614174000",
				"profileId": 1,
				"profileName": "Test Profile 1",
				"scopeDescription": "All Computers",
				"siteId": "1",
				"version": "2.0.0",
				"autoDeploymentType": "INSTALL_AUTOMATICALLY"
			},
			{
				"uuid": "223e4567-e89b-12d3-a456-426614174001",
				"profileId": 2,
				"profileName": "Test Profile 2",
				"scopeDescription": "Marketing Department",
				"siteId": "1",
				"version": "2.1.0",
				"autoDeploymentType": "PROMPT_USERS_TO_INSTALL"
			}
		]
	}`)
	m.register("GET", "/api/v1/jamf-connect/config-profiles", 200, body)
}

func (m *JamfConnectMock) RegisterUpdateConfigProfileMock() {
	body := []byte(`{
		"uuid": "123e4567-e89b-12d3-a456-426614174000",
		"profileId": 1,
		"profileName": "Test Profile 1",
		"scopeDescription": "All Computers",
		"siteId": "1",
		"version": "2.1.0",
		"autoDeploymentType": "INSTALL_AUTOMATICALLY"
	}`)
	m.register("PUT", "/api/v1/jamf-connect/config-profiles/123e4567-e89b-12d3-a456-426614174000", 200, body)
}

func (m *JamfConnectMock) RegisterRetryDeploymentTasksMock() {
	m.register("POST", "/api/v1/jamf-connect/deployments/123e4567-e89b-12d3-a456-426614174000/tasks/retry", 204, []byte{})
}

func (m *JamfConnectMock) RegisterGetDeploymentTasksMock() {
	body := []byte(`{
		"totalCount": 2,
		"results": [
			{"status": "COMPLETED", "updated": "2024-01-15T10:00:00Z", "version": "2.0.0"},
			{"status": "PENDING", "updated": "2024-01-15T11:00:00Z", "version": "2.1.0"}
		]
	}`)
	m.register("GET", "/api/v1/jamf-connect/deployments/dep-123/tasks", 200, body)
}

func (m *JamfConnectMock) RegisterGetHistoryMock() {
	body := []byte(`{
		"totalCount": 1,
		"results": [
			{"id": "1", "username": "admin", "date": "2024-01-15T10:00:00Z", "note": "Config updated", "details": "Profile modified"}
		]
	}`)
	m.register("GET", "/api/v1/jamf-connect/history", 200, body)
}

// RegisterListConfigProfilesBadJSONMock registers a response with invalid JSON body for the config profiles endpoint.
// This is used to test the mergePage unmarshal error path.
func (m *JamfConnectMock) RegisterListConfigProfilesBadJSONMock() {
	m.responses["GET:/api/v1/jamf-connect/config-profiles"] = registeredResponse{
		statusCode: 200,
		rawBody:    []byte("not-valid-json"),
	}
}

// RegisterListConfigProfilesBadResultsMock registers a response where profileId is a string (not an int),
// causing mapstructure.Decode to fail when decoding the profile.
func (m *JamfConnectMock) RegisterListConfigProfilesBadResultsMock() {
	body := []byte(`{"totalCount": 1, "results": [{"profileId": "not-a-number"}]}`)
	m.responses["GET:/api/v1/jamf-connect/config-profiles"] = registeredResponse{
		statusCode: 200,
		rawBody:    body,
	}
}

// RegisterGetDeploymentTasksBadJSONMock registers invalid JSON for the deployment tasks endpoint.
func (m *JamfConnectMock) RegisterGetDeploymentTasksBadJSONMock() {
	m.responses["GET:/api/v1/jamf-connect/deployments/dep-123/tasks"] = registeredResponse{
		statusCode: 200,
		rawBody:    []byte("not-valid-json"),
	}
}

// RegisterGetDeploymentTasksBadResultsMock registers a response where a field has an incompatible type.
func (m *JamfConnectMock) RegisterGetDeploymentTasksBadResultsMock() {
	body := []byte(`{"totalCount": 1, "results": [{"status": ["not", "a", "string"]}]}`)
	m.responses["GET:/api/v1/jamf-connect/deployments/dep-123/tasks"] = registeredResponse{
		statusCode: 200,
		rawBody:    body,
	}
}

// RegisterGetHistoryBadJSONMock registers invalid JSON for the history endpoint.
func (m *JamfConnectMock) RegisterGetHistoryBadJSONMock() {
	m.responses["GET:/api/v1/jamf-connect/history"] = registeredResponse{
		statusCode: 200,
		rawBody:    []byte("not-valid-json"),
	}
}

// RegisterGetHistoryBadResultsMock registers a response where a field has an incompatible type.
func (m *JamfConnectMock) RegisterGetHistoryBadResultsMock() {
	body := []byte(`{"totalCount": 1, "results": [{"id": ["not", "a", "string"]}]}`)
	m.responses["GET:/api/v1/jamf-connect/history"] = registeredResponse{
		statusCode: 200,
		rawBody:    body,
	}
}

func (m *JamfConnectMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *JamfConnectMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfConnectMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfConnectMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfConnectMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfConnectMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *JamfConnectMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *JamfConnectMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfConnectMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfConnectMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *JamfConnectMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &page); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *JamfConnectMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *JamfConnectMock) InvalidateToken() error                    { return nil }
func (m *JamfConnectMock) KeepAliveToken() error                     { return nil }
func (m *JamfConnectMock) GetLogger() *zap.Logger                    { return m.logger }
