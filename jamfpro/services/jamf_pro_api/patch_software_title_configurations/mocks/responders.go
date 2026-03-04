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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
}

type PatchSoftwareTitleConfigurationsMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewPatchSoftwareTitleConfigurationsMock() *PatchSoftwareTitleConfigurationsMock {
	return &PatchSoftwareTitleConfigurationsMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *PatchSoftwareTitleConfigurationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterListMock() {
	m.register("GET", "/api/v2/patch-software-title-configurations", 200, "validate_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterCreateMock() {
	m.register("POST", "/api/v2/patch-software-title-configurations", 200, "validate_create.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v2/patch-software-title-configurations/"+id, 200, "validate_get.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v2/patch-software-title-configurations/"+id, 200, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterEmptyListMock() {
	m.register("GET", "/api/v2/patch-software-title-configurations", 200, "validate_empty_list.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDashboardStatusMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/dashboard", 200, "validate_dashboard_status.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddToDashboardMock(id string) {
	m.register("POST", "/api/v2/patch-software-title-configurations/"+id+"/dashboard", 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterRemoveFromDashboardMock(id string) {
	m.register("DELETE", "/api/v2/patch-software-title-configurations/"+id+"/dashboard", 204, "")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDefinitionsMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/definitions", 200, "validate_definitions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetDependenciesMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/dependencies", 200, "validate_dependencies.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterExportReportMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/export-report", 200, "validate_export_report.csv")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetExtensionAttributesMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/extension-attributes", 200, "validate_extension_attributes.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchReportMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/patch-report", 200, "validate_patch_report.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchSummaryMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/patch-summary", 200, "validate_patch_summary.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetHistoryMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/history", 200, "validate_history.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterAddHistoryNoteMock(id string) {
	m.register("POST", "/api/v2/patch-software-title-configurations/"+id+"/history", 201, "validate_add_history_note.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) RegisterGetPatchVersionsMock(id string) {
	m.register("GET", "/api/v2/patch-software-title-configurations/"+id+"/patch-summary/versions", 200, "validate_patch_versions.json")
}

func (m *PatchSoftwareTitleConfigurationsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("PatchSoftwareTitleConfigurationsMock: no response for %s %s", method, path)
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

func (m *PatchSoftwareTitleConfigurationsMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchSoftwareTitleConfigurationsMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *PatchSoftwareTitleConfigurationsMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	bodyBytes := resp.Bytes()
	if mergePage != nil && len(bodyBytes) > 0 {
		// Extract the results array from the response for pagination
		var wrapper struct {
			Results json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &wrapper); err == nil {
			_ = mergePage(wrapper.Results)
		}
	}
	return resp, nil
}
func (m *PatchSoftwareTitleConfigurationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *PatchSoftwareTitleConfigurationsMock) InvalidateToken() error                    { return nil }
func (m *PatchSoftwareTitleConfigurationsMock) KeepAliveToken() error                     { return nil }
func (m *PatchSoftwareTitleConfigurationsMock) GetLogger() *zap.Logger                    { return m.logger }
