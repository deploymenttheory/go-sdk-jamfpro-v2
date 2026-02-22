package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_get.json
var validateGetJSON []byte

//go:embed validate_eligible_apps.json
var validateEligibleAppsJSON []byte

//go:embed validate_history.json
var validateHistoryJSON []byte

//go:embed validate_add_history_notes.json
var validateAddHistoryNotesJSON []byte

//go:embed validate_export_history.csv
var validateExportHistoryCSV []byte

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type OnboardingMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewOnboardingMock() *OnboardingMock {
	return &OnboardingMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *OnboardingMock) register(method, path string, statusCode int, body []byte) {
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *OnboardingMock) RegisterGetMock() {
	m.register("GET", "/api/v1/onboarding", 200, validateGetJSON)
}

func (m *OnboardingMock) RegisterUpdateMock() {
	m.register("PUT", "/api/v1/onboarding", 200, validateGetJSON)
}

func (m *OnboardingMock) RegisterGetEligibleAppsMock() {
	m.register("GET", "/api/v1/onboarding/eligible-apps", 200, validateEligibleAppsJSON)
}

func (m *OnboardingMock) RegisterGetEligibleConfigurationProfilesMock() {
	m.register("GET", "/api/v1/onboarding/eligible-configuration-profiles", 200, validateEligibleAppsJSON)
}

func (m *OnboardingMock) RegisterGetEligiblePoliciesMock() {
	m.register("GET", "/api/v1/onboarding/eligible-policies", 200, validateEligibleAppsJSON)
}

func (m *OnboardingMock) RegisterGetHistoryMock() {
	m.register("GET", "/api/v1/onboarding/history", 200, validateHistoryJSON)
}

func (m *OnboardingMock) RegisterAddHistoryNotesMock() {
	m.register("POST", "/api/v1/onboarding/history", 201, validateAddHistoryNotesJSON)
}

func (m *OnboardingMock) RegisterExportHistoryMock() {
	m.register("GET", "/api/v1/onboarding/history/export", 200, validateExportHistoryCSV)
}

func (m *OnboardingMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("OnboardingMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func (m *OnboardingMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *OnboardingMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *OnboardingMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *OnboardingMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *OnboardingMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OnboardingMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *OnboardingMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *OnboardingMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *OnboardingMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *OnboardingMock) InvalidateToken() error                    { return nil }
func (m *OnboardingMock) KeepAliveToken() error                      { return nil }
func (m *OnboardingMock) GetLogger() *zap.Logger                      { return m.logger }
