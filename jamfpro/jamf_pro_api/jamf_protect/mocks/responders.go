package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// JamfProtectMock is a test double implementing transport.HTTPClient.
type JamfProtectMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewJamfProtectMock returns an empty mock ready for response registration.
func NewJamfProtectMock() *JamfProtectMock {
	return &JamfProtectMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *JamfProtectMock) RegisterMocks() {
	m.RegisterGetSettingsMock()
	m.RegisterUpdateSettingsMock()
	m.RegisterRegisterMock()
	m.RegisterSyncPlansMock()
	m.RegisterListDeploymentTasksMock()
	m.RegisterRetryDeploymentTasksMock()
	m.RegisterListHistoryMock()
	m.RegisterCreateHistoryNoteMock()
	m.RegisterListPlansMock()
	m.RegisterDeleteIntegrationMock()
}

func (m *JamfProtectMock) register(method, path string, statusCode int, rawJSON string) {
	key := method + " " + path
	m.responses[key] = registeredResponse{
		statusCode: statusCode,
		rawBody:    []byte(rawJSON),
	}
}

// RegisterGetSettingsMock registers a successful response for GetSettingsV1.
func (m *JamfProtectMock) RegisterGetSettingsMock() {
	m.register("GET", "/api/v1/jamf-protect", 200, `{
		"id": "1",
		"protectUrl": "https://protect.example.com",
		"syncStatus": "SYNCED",
		"apiClientId": "client-123",
		"autoInstall": true,
		"lastSyncTime": "2025-01-15T10:30:00Z",
		"apiClientName": "Jamf Pro",
		"registrationId": "reg-456"
	}`)
}

// RegisterUpdateSettingsMock registers a successful response for UpdateSettingsV1.
func (m *JamfProtectMock) RegisterUpdateSettingsMock() {
	m.register("PUT", "/api/v1/jamf-protect", 200, `{
		"id": "1",
		"protectUrl": "https://protect.example.com",
		"syncStatus": "SYNCED",
		"apiClientId": "client-123",
		"autoInstall": true,
		"lastSyncTime": "2025-01-15T10:30:00Z",
		"apiClientName": "Jamf Pro",
		"registrationId": "reg-456"
	}`)
}

// RegisterRegisterMock registers a successful response for RegisterV1.
func (m *JamfProtectMock) RegisterRegisterMock() {
	m.register("POST", "/api/v1/jamf-protect/register", 201, `{
		"id": "1",
		"protectUrl": "https://protect.example.com",
		"syncStatus": "PENDING",
		"apiClientId": "client-123",
		"autoInstall": false,
		"lastSyncTime": "",
		"apiClientName": "Jamf Pro",
		"registrationId": "reg-456"
	}`)
}

// RegisterSyncPlansMock registers a successful response for SyncPlansV1.
func (m *JamfProtectMock) RegisterSyncPlansMock() {
	m.register("POST", "/api/v1/jamf-protect/plans/sync", 204, ``)
}

// RegisterListDeploymentTasksMock registers a successful response for ListDeploymentTasksV1.
func (m *JamfProtectMock) RegisterListDeploymentTasksMock() {
	m.register("GET", "/api/v1/jamf-protect/deployments/deploy-123/tasks", 200, `{
		"totalCount": 2,
		"results": [
			{
				"id": "task-1",
				"status": "COMPLETED",
				"updated": "2025-01-15T10:00:00Z",
				"version": "5.0.0",
				"computerId": "100",
				"computerName": "MacBook-Pro-1"
			},
			{
				"id": "task-2",
				"status": "PENDING",
				"updated": "2025-01-15T10:05:00Z",
				"version": "5.0.0",
				"computerId": "101",
				"computerName": "MacBook-Pro-2"
			}
		]
	}`)
}

// RegisterRetryDeploymentTasksMock registers a successful response for RetryDeploymentTasksV1.
func (m *JamfProtectMock) RegisterRetryDeploymentTasksMock() {
	m.register("POST", "/api/v1/jamf-protect/deployments/deploy-123/tasks/retry", 204, ``)
}

// RegisterListHistoryMock registers a successful response for ListHistoryV1.
func (m *JamfProtectMock) RegisterListHistoryMock() {
	m.register("GET", "/api/v1/jamf-protect/history", 200, `{
		"totalCount": 2,
		"results": [
			{
				"id": 1,
				"username": "admin",
				"date": "2025-01-15T10:00:00Z",
				"note": "Registered Jamf Protect",
				"details": "Initial registration"
			},
			{
				"id": 2,
				"username": "admin",
				"date": "2025-01-15T10:30:00Z",
				"note": "Updated settings",
				"details": "Enabled auto-install"
			}
		]
	}`)
}

// RegisterCreateHistoryNoteMock registers a successful response for CreateHistoryNoteV1.
func (m *JamfProtectMock) RegisterCreateHistoryNoteMock() {
	m.register("POST", "/api/v1/jamf-protect/history", 201, `{
		"id": "3",
		"username": "admin",
		"date": "2025-01-15T11:00:00Z",
		"note": "Test note",
		"details": "Test details"
	}`)
}

// RegisterListPlansMock registers a successful response for ListPlansV1.
func (m *JamfProtectMock) RegisterListPlansMock() {
	m.register("GET", "/api/v1/jamf-protect/plans", 200, `{
		"totalCount": 2,
		"results": [
			{
				"uuid": "plan-uuid-1",
				"id": "plan-1",
				"name": "Standard Protection",
				"description": "Standard security plan",
				"profileName": "StandardProfile",
				"profileId": 10,
				"scopeDescription": "All computers"
			},
			{
				"uuid": "plan-uuid-2",
				"id": "plan-2",
				"name": "Advanced Protection",
				"description": "Advanced security plan",
				"profileName": "AdvancedProfile",
				"profileId": 20,
				"scopeDescription": "Executive computers"
			}
		]
	}`)
}

// RegisterDeleteIntegrationMock registers a successful response for DeleteIntegrationV1.
func (m *JamfProtectMock) RegisterDeleteIntegrationMock() {
	m.register("DELETE", "/api/v1/jamf-protect", 204, ``)
}

// RegisterListDeploymentTasksBadJSONMock registers an invalid JSON response for the tasks endpoint.
func (m *JamfProtectMock) RegisterListDeploymentTasksBadJSONMock() {
	m.register("GET", "/api/v1/jamf-protect/deployments/deploy-123/tasks", 200, "not-valid-json")
}

// RegisterListDeploymentTasksBadResultsMock registers a response with incompatible field types.
func (m *JamfProtectMock) RegisterListDeploymentTasksBadResultsMock() {
	m.register("GET", "/api/v1/jamf-protect/deployments/deploy-123/tasks", 200, `{"totalCount": 1, "results": [{"id": ["not", "a", "string"]}]}`)
}

// RegisterListHistoryBadJSONMock registers an invalid JSON response for the history endpoint.
func (m *JamfProtectMock) RegisterListHistoryBadJSONMock() {
	m.register("GET", "/api/v1/jamf-protect/history", 200, "not-valid-json")
}

// RegisterListHistoryBadResultsMock registers a response with incompatible field types.
func (m *JamfProtectMock) RegisterListHistoryBadResultsMock() {
	m.register("GET", "/api/v1/jamf-protect/history", 200, `{"totalCount": 1, "results": [{"id": ["not", "a", "number"]}]}`)
}

// RegisterListPlansBadJSONMock registers an invalid JSON response for the plans endpoint.
func (m *JamfProtectMock) RegisterListPlansBadJSONMock() {
	m.register("GET", "/api/v1/jamf-protect/plans", 200, "not-valid-json")
}

// RegisterListPlansBadResultsMock registers a response with incompatible field types.
func (m *JamfProtectMock) RegisterListPlansBadResultsMock() {
	m.register("GET", "/api/v1/jamf-protect/plans", 200, `{"totalCount": 1, "results": [{"profileId": "not-a-number"}]}`)
}

// Get implements transport.HTTPClient.
func (m *JamfProtectMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// Post implements transport.HTTPClient.
func (m *JamfProtectMock) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "POST " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for POST %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// PostWithQuery implements transport.HTTPClient.
func (m *JamfProtectMock) PostWithQuery(ctx context.Context, path string, rsqlQuery map[string]string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, body, headers, result)
}

// PostForm implements transport.HTTPClient.
func (m *JamfProtectMock) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	return m.Post(ctx, path, formData, headers, result)
}

// PostMultipart implements transport.HTTPClient.
func (m *JamfProtectMock) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.Post(ctx, path, nil, headers, result)
}

// Put implements transport.HTTPClient.
func (m *JamfProtectMock) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PUT " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PUT %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// Patch implements transport.HTTPClient.
func (m *JamfProtectMock) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	key := "PATCH " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for PATCH %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// Delete implements transport.HTTPClient.
func (m *JamfProtectMock) Delete(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*resty.Response, error) {
	key := "DELETE " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for DELETE %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	if result != nil && len(resp.rawBody) > 0 {
		if err := json.Unmarshal(resp.rawBody, result); err != nil {
			return nil, fmt.Errorf("unmarshal mock response: %w", err)
		}
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), nil
}

// DeleteWithBody implements transport.HTTPClient.
func (m *JamfProtectMock) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*resty.Response, error) {
	return m.Delete(ctx, path, nil, headers, result)
}

// GetBytes implements transport.HTTPClient.
func (m *JamfProtectMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*resty.Response, []byte, error) {
	m.LastRSQLQuery = rsqlQuery
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, nil, fmt.Errorf("%s", resp.errMsg)
	}
	return shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody), resp.rawBody, nil
}

// GetPaginated implements transport.HTTPClient.
func (m *JamfProtectMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*resty.Response, error) {
	key := "GET " + path
	resp, ok := m.responses[key]
	if !ok {
		return nil, fmt.Errorf("no mock registered for GET %s", path)
	}
	if resp.errMsg != "" {
		return nil, fmt.Errorf("%s", resp.errMsg)
	}
	ifaceResp := shared.NewMockResponse(resp.statusCode, http.Header{}, resp.rawBody)
	if mergePage != nil && len(resp.rawBody) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(resp.rawBody, &page); err != nil {
			return ifaceResp, fmt.Errorf("mergePage: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return ifaceResp, fmt.Errorf("mergePage: %w", err)
		}
	}
	return ifaceResp, nil
}

// RSQLBuilder implements transport.HTTPClient.
func (m *JamfProtectMock) RSQLBuilder() transport.RSQLFilterBuilder {
	return nil
}

// InvalidateToken implements transport.HTTPClient.
func (m *JamfProtectMock) InvalidateToken() error {
	return nil
}

// KeepAliveToken implements transport.HTTPClient.
func (m *JamfProtectMock) KeepAliveToken() error {
	return nil
}

// GetLogger implements transport.HTTPClient.
func (m *JamfProtectMock) GetLogger() *zap.Logger {
	return m.logger
}
