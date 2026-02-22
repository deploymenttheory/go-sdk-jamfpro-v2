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

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type ManagedSoftwareUpdatesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewManagedSoftwareUpdatesMock() *ManagedSoftwareUpdatesMock {
	return &ManagedSoftwareUpdatesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ManagedSoftwareUpdatesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetAvailableUpdatesMock() {
	m.register("GET", "/api/v1/managed-software-updates/available-updates", 200, "validate_available_updates.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansMock() {
	m.register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlanByUUIDMock(uuid string) {
	m.register("GET", "/api/v1/managed-software-updates/plans/"+uuid, 200, "validate_plan_detail.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetDeclarationsByPlanUUIDMock(uuid string) {
	m.register("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/declarations", 200, "validate_declarations_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterCreatePlanByDeviceIDMock() {
	m.register("POST", "/api/v1/managed-software-updates/plans", 201, "validate_plan_create.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterCreatePlanByGroupIDMock() {
	m.register("POST", "/api/v1/managed-software-updates/plans/group", 201, "validate_plan_create.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansByGroupIDMock(groupID string) {
	m.register("GET", "/api/v1/managed-software-updates/plans/group/"+groupID+"?group-type=COMPUTER", 200, "validate_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetFeatureToggleMock() {
	m.register("GET", "/api/v1/managed-software-updates/plans/feature-toggle", 200, "validate_feature_toggle_get.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterUpdateFeatureToggleMock() {
	m.register("PUT", "/api/v1/managed-software-updates/plans/feature-toggle", 200, "validate_feature_toggle_response.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetFeatureToggleStatusMock() {
	m.register("GET", "/api/v1/managed-software-updates/plans/feature-toggle/status", 200, "validate_feature_toggle_status.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterForceStopFeatureToggleProcessMock() {
	m.register("POST", "/api/v1/managed-software-updates/plans/feature-toggle/abandon", 200, "validate_error_response.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterEmptyPlansMock() {
	m.register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_empty_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("ManagedSoftwareUpdatesMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *ManagedSoftwareUpdatesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ManagedSoftwareUpdatesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ManagedSoftwareUpdatesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ManagedSoftwareUpdatesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		// Pass the full response body to mergePage for processing
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *ManagedSoftwareUpdatesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ManagedSoftwareUpdatesMock) InvalidateToken() error                     { return nil }
func (m *ManagedSoftwareUpdatesMock) KeepAliveToken() error                      { return nil }
func (m *ManagedSoftwareUpdatesMock) GetLogger() *zap.Logger                     { return m.logger }
