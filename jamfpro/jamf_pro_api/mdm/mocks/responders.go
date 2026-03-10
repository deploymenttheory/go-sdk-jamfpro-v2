package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

// MDMMock is a test double implementing client.Client.
type MDMMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewMDMMock returns an empty mock ready for response registration.
func NewMDMMock() *MDMMock {
	return &MDMMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *MDMMock) RegisterMocks() {
	m.RegisterListCommandsMock()
	m.RegisterBlankPushMock()
	m.RegisterSendCommandMock()
	m.RegisterDeployPackageMock()
	m.RegisterRenewProfileMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *MDMMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *MDMMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("MDMMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MDMMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("MDMMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *MDMMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("MDMMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("MDMMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("could not get caller info")
	}
	dir := filepath.Dir(f)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

func (m *MDMMock) RegisterListCommandsMock() {
	m.register("GET", "/api/v2/mdm/commands", 200, "validate_list_commands.json")
}

func (m *MDMMock) RegisterBlankPushMock() {
	m.register("POST", "/api/v2/mdm/blank-push", 200, "validate_blank_push.json")
}

func (m *MDMMock) RegisterSendCommandMock() {
	m.register("POST", "/api/v2/mdm/commands", 200, "validate_send_command.json")
}

func (m *MDMMock) RegisterDeployPackageMock() {
	m.register("POST", "/api/v1/deploy-package?verbose=true", 200, "validate_deploy_package.json")
}

func (m *MDMMock) RegisterRenewProfileMock() {
	m.register("POST", "/api/v1/mdm/renew-profile", 200, "validate_renew_profile.json")
}

func (m *MDMMock) RegisterNotFoundErrorMock() {
	m.registerError("POST", "/api/v2/mdm/commands", 404, "error_not_found.json")
}

func (m *MDMMock) RegisterBlankPushErrorMock() {
	m.registerError("POST", "/api/v2/mdm/blank-push", 500, "error_not_found.json")
}

func (m *MDMMock) RegisterDeployPackageErrorMock() {
	m.registerError("POST", "/api/v1/deploy-package?verbose=true", 500, "error_not_found.json")
}

func (m *MDMMock) RegisterRenewProfileErrorMock() {
	m.registerError("POST", "/api/v1/mdm/renew-profile", 500, "error_not_found.json")
}

func (m *MDMMock) RegisterListCommandsErrorMock() {
	m.registerError("GET", "/api/v2/mdm/commands", 500, "error_not_found.json")
}

func (m *MDMMock) RegisterListCommandsInvalidJSONMock() {
	m.register("GET", "/api/v2/mdm/commands", 200, "validate_list_commands_invalid.json")
}

func (m *MDMMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *MDMMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MDMMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MDMMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MDMMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *MDMMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *MDMMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *MDMMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MDMMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *MDMMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *MDMMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *MDMMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *MDMMock) InvalidateToken() error                { return nil }
func (m *MDMMock) KeepAliveToken() error                 { return nil }
func (m *MDMMock) GetLogger() *zap.Logger                { return m.logger }
