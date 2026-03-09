package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// AppRequestMock is a test double implementing transport.HTTPClient.
// Responses are keyed by "METHOD:path" and loaded from JSON fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
type AppRequestMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewAppRequestMock returns an empty mock ready for response registration.
func NewAppRequestMock() *AppRequestMock {
	return &AppRequestMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *AppRequestMock) RegisterMocks() {
	m.RegisterListFormInputFieldsMock()
	m.RegisterReplaceFormInputFieldsMock()
	m.RegisterCreateFormInputFieldMock()
	m.RegisterGetFormInputFieldMock()
	m.RegisterUpdateFormInputFieldMock()
	m.RegisterDeleteFormInputFieldMock()
	m.RegisterGetSettingsMock()
	m.RegisterUpdateSettingsMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *AppRequestMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

// ---- Success responders ----

// RegisterListFormInputFieldsMock registers GET /api/v1/app-request/form-input-fields → 200.
func (m *AppRequestMock) RegisterListFormInputFieldsMock() {
	m.register("GET", "/api/v1/app-request/form-input-fields", 200, "validate_list_form_input_fields.json")
}

// RegisterReplaceFormInputFieldsMock registers PUT /api/v1/app-request/form-input-fields → 200.
func (m *AppRequestMock) RegisterReplaceFormInputFieldsMock() {
	m.register("PUT", "/api/v1/app-request/form-input-fields", 200, "validate_replace_form_input_fields.json")
}

// RegisterCreateFormInputFieldMock registers POST /api/v1/app-request/form-input-fields → 201.
func (m *AppRequestMock) RegisterCreateFormInputFieldMock() {
	m.register("POST", "/api/v1/app-request/form-input-fields", 201, "validate_create_form_input_field.json")
}

// RegisterGetFormInputFieldMock registers GET /api/v1/app-request/form-input-fields/1 → 200.
func (m *AppRequestMock) RegisterGetFormInputFieldMock() {
	m.register("GET", "/api/v1/app-request/form-input-fields/1", 200, "validate_get_form_input_field.json")
}

// RegisterUpdateFormInputFieldMock registers PUT /api/v1/app-request/form-input-fields/1 → 200.
func (m *AppRequestMock) RegisterUpdateFormInputFieldMock() {
	m.register("PUT", "/api/v1/app-request/form-input-fields/1", 200, "validate_get_form_input_field.json")
}

// RegisterDeleteFormInputFieldMock registers DELETE /api/v1/app-request/form-input-fields/1 → 204.
func (m *AppRequestMock) RegisterDeleteFormInputFieldMock() {
	m.register("DELETE", "/api/v1/app-request/form-input-fields/1", 204, "")
}

// RegisterGetSettingsMock registers GET /api/v1/app-request/settings → 200.
func (m *AppRequestMock) RegisterGetSettingsMock() {
	m.register("GET", "/api/v1/app-request/settings", 200, "validate_get_settings.json")
}

// RegisterUpdateSettingsMock registers PUT /api/v1/app-request/settings → 200.
func (m *AppRequestMock) RegisterUpdateSettingsMock() {
	m.register("PUT", "/api/v1/app-request/settings", 200, "validate_update_settings.json")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /api/v1/app-request/form-input-fields/999 → 404.
func (m *AppRequestMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/app-request/form-input-fields/999", 404, "error_not_found.json")
}

// ---- transport.HTTPClient implementation ----

func (m *AppRequestMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *AppRequestMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AppRequestMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AppRequestMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AppRequestMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *AppRequestMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *AppRequestMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *AppRequestMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AppRequestMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *AppRequestMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *AppRequestMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}

	bodyBytes := resp.Bytes()
	var wrapper struct {
		Results json.RawMessage `json:"results"`
	}
	if err := json.Unmarshal(bodyBytes, &wrapper); err != nil {
		return resp, fmt.Errorf("failed to extract results: %w", err)
	}

	if mergePage != nil {
		if err := mergePage(wrapper.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *AppRequestMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *AppRequestMock) InvalidateToken() error                    { return nil }
func (m *AppRequestMock) KeepAliveToken() error                     { return nil }
func (m *AppRequestMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 204 No Content responses).
func (m *AppRequestMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("AppRequestMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// registerError stores an error response. The fixture body is returned alongside the error.
func (m *AppRequestMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("AppRequestMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

// dispatch looks up the registered response and either unmarshals the body
// into result or returns an error depending on the registration type.
func (m *AppRequestMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {"application/json"}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`)), fmt.Errorf("AppRequestMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("AppRequestMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads a JSON fixture file from the mocks/ directory
// adjacent to the test being run.
func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working directory: %w", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "mocks", filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}
