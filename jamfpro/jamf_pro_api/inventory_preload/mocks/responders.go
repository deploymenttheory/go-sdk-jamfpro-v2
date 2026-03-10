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
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type InventoryPreloadMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewInventoryPreloadMock() *InventoryPreloadMock {
	return &InventoryPreloadMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *InventoryPreloadMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("InventoryPreloadMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *InventoryPreloadMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("InventoryPreloadMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *InventoryPreloadMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("InventoryPreloadMock: no response registered for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("InventoryPreloadMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *InventoryPreloadMock) RegisterCreateFromCSVMock() {
	m.register("POST", "/api/v2/inventory-preload/csv", 200, "validate_create_from_csv.json")
}

func (m *InventoryPreloadMock) RegisterGetCSVTemplateMock() {
	m.register("GET", "/api/v2/inventory-preload/csv-template", 200, "validate_csv_template.csv")
}

func (m *InventoryPreloadMock) RegisterValidateCSVMock() {
	m.register("POST", "/api/v2/inventory-preload/csv-validate", 200, "validate_csv_validation.json")
}

func (m *InventoryPreloadMock) RegisterGetEAColumnsMock() {
	m.register("GET", "/api/v2/inventory-preload/ea-columns", 200, "validate_ea_columns.json")
}

func (m *InventoryPreloadMock) RegisterExportMock() {
	m.register("POST", "/api/v2/inventory-preload/export", 200, "validate_list_records.json")
}

func (m *InventoryPreloadMock) RegisterListHistoryMock() {
	m.register("GET", "/api/v2/inventory-preload/history", 200, "validate_history_list.json")
}

func (m *InventoryPreloadMock) RegisterAddHistoryNoteMock() {
	m.register("POST", "/api/v2/inventory-preload/history", 201, "validate_add_history_note.json")
}

func (m *InventoryPreloadMock) RegisterListRecordsMock() {
	m.register("GET", "/api/v2/inventory-preload/records", 200, "validate_list_records.json")
}

func (m *InventoryPreloadMock) RegisterListRecordsInvalidJSONMock() {
	m.responses["GET:/api/v2/inventory-preload/records"] = registeredResponse{statusCode: 200, rawBody: []byte(`{invalid json`)}
}

func (m *InventoryPreloadMock) RegisterCreateRecordMock() {
	m.register("POST", "/api/v2/inventory-preload/records", 201, "validate_create_record.json")
}

func (m *InventoryPreloadMock) RegisterDeleteAllRecordsMock() {
	m.register("POST", "/api/v2/inventory-preload/records/delete-all", 200, "")
}

func (m *InventoryPreloadMock) RegisterGetRecordByIDMock(id string) {
	m.register("GET", "/api/v2/inventory-preload/records/"+id, 200, "validate_get_record.json")
}

func (m *InventoryPreloadMock) RegisterUpdateRecordMock(id string) {
	m.register("PUT", "/api/v2/inventory-preload/records/"+id, 200, "validate_get_record.json")
}

func (m *InventoryPreloadMock) RegisterDeleteRecordMock(id string) {
	m.register("DELETE", "/api/v2/inventory-preload/records/"+id, 200, "")
}

func (m *InventoryPreloadMock) RegisterNotFoundErrorMock(id string) {
	m.registerError("GET", "/api/v2/inventory-preload/records/"+id, 404, "error_not_found.json")
	m.registerError("PUT", "/api/v2/inventory-preload/records/"+id, 404, "error_not_found.json")
	m.registerError("DELETE", "/api/v2/inventory-preload/records/"+id, 404, "error_not_found.json")
}

func (m *InventoryPreloadMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *InventoryPreloadMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *InventoryPreloadMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *InventoryPreloadMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *InventoryPreloadMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *InventoryPreloadMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *InventoryPreloadMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *InventoryPreloadMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *InventoryPreloadMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *InventoryPreloadMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *InventoryPreloadMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
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
func (m *InventoryPreloadMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilderWithQueryCapture(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	}, &m.LastRSQLQuery)
}

func (m *InventoryPreloadMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *InventoryPreloadMock) InvalidateToken() error                    { return nil }
func (m *InventoryPreloadMock) KeepAliveToken() error                     { return nil }
func (m *InventoryPreloadMock) GetLogger() *zap.Logger                    { return m.logger }
