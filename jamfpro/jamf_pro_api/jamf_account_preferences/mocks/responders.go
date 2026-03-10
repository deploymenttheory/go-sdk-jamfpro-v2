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

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type JamfAccountPreferencesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewJamfAccountPreferencesMock() *JamfAccountPreferencesMock {
	return &JamfAccountPreferencesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *JamfAccountPreferencesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("JamfAccountPreferencesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *JamfAccountPreferencesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("JamfAccountPreferencesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *JamfAccountPreferencesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("JamfAccountPreferencesMock: no response registered for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("JamfAccountPreferencesMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

const endpointV3 = "/api/v3/account-preferences"

func (m *JamfAccountPreferencesMock) RegisterGetV3Mock() {
	m.register("GET", endpointV3, 200, "validate_get.json")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3Mock() {
	m.register("PATCH", endpointV3, 200, "validate_update.json")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3_204NoContentMock() {
	m.register("PATCH", endpointV3, 204, "")
}

func (m *JamfAccountPreferencesMock) RegisterGetV3ErrorMock() {
	m.registerError("GET", endpointV3, 404, "error_not_found.json")
}

func (m *JamfAccountPreferencesMock) RegisterUpdateV3ErrorMock() {
	m.registerError("PATCH", endpointV3, 500, "error_not_found.json")
}

func (m *JamfAccountPreferencesMock) RegisterInvalidJSONMock() {
	m.responses["GET:"+endpointV3] = registeredResponse{statusCode: 200, rawBody: []byte(`{invalid json`)}
}

func (m *JamfAccountPreferencesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *JamfAccountPreferencesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfAccountPreferencesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfAccountPreferencesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfAccountPreferencesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *JamfAccountPreferencesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *JamfAccountPreferencesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *JamfAccountPreferencesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfAccountPreferencesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *JamfAccountPreferencesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *JamfAccountPreferencesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		if err := mergePage(bodyBytes); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *JamfAccountPreferencesMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *JamfAccountPreferencesMock) InvalidateToken() error                    { return nil }
func (m *JamfAccountPreferencesMock) KeepAliveToken() error                     { return nil }
func (m *JamfAccountPreferencesMock) GetLogger() *zap.Logger                    { return m.logger }
