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

// RegisterErrorMock registers an error response for testing error paths.
func (m *ManagedSoftwareUpdatesMock) RegisterErrorMock(method, path, errMsg string) {
	m.responses[method+":"+path] = registeredResponse{errMsg: errMsg}
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

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlanEventsByUUIDMock(uuid string) {
	m.register("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/events", 200, "validate_plan_events.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesMock() {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses.json")
}

// RegisterGetUpdateStatusesArrayFormatMock uses raw array format (Real API style).
func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesArrayFormatMock() {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses_array.json")
}

// RegisterGetPlansInvalidMock returns invalid JSON to exercise mergePage error path.
func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansInvalidMock() {
	m.register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_plans_list_invalid.json")
}

// RegisterGetUpdateStatusesInvalidMock returns invalid JSON to exercise mergePage error path.
func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesInvalidMock() {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses_invalid.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByComputerGroupMock(id string) {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses/computer-groups/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByComputerMock(id string) {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses/computers/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByMobileDeviceGroupMock(id string) {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses/mobile-device-groups/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByMobileDeviceMock(id string) {
	m.register("GET", "/api/v1/managed-software-updates/update-statuses/mobile-devices/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ManagedSoftwareUpdatesMock: no response for %s %s", method, path)
	}
	if r.errMsg != "" {
		return nil, fmt.Errorf("%s", r.errMsg)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *ManagedSoftwareUpdatesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ManagedSoftwareUpdatesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ManagedSoftwareUpdatesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ManagedSoftwareUpdatesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *ManagedSoftwareUpdatesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return nil, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		var page struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &page); err != nil {
			return nil, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return nil, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *ManagedSoftwareUpdatesMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *ManagedSoftwareUpdatesMock) InvalidateToken() error                    { return nil }
func (m *ManagedSoftwareUpdatesMock) KeepAliveToken() error                     { return nil }
func (m *ManagedSoftwareUpdatesMock) GetLogger() *zap.Logger                    { return m.logger }
