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
}

type MobileDevicePrestagesMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewMobileDevicePrestagesMock() *MobileDevicePrestagesMock {
	return &MobileDevicePrestagesMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *MobileDevicePrestagesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *MobileDevicePrestagesMock) RegisterListMock() {
	m.register("GET", "/api/v3/mobile-device-prestages", 200, "validate_list.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterCreateMock() {
	m.register("POST", "/api/v3/mobile-device-prestages", 200, "validate_create.json")
}

func (m *MobileDevicePrestagesMock) RegisterUpdateByIDMock(id string) {
	m.register("PUT", "/api/v3/mobile-device-prestages/"+id, 200, "validate_get.json")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v3/mobile-device-prestages/"+id, 200, "")
}

func (m *MobileDevicePrestagesMock) RegisterGetScopeByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterEmptyListMock() {
	m.register("GET", "/api/v3/mobile-device-prestages", 200, "validate_empty_list.json")
}

func (m *MobileDevicePrestagesMock) RegisterReplaceScopeByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.register("PUT", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

// RegisterReplaceScopePutOnlyMock registers only PUT (no GET) to test scope fetch error.
func (m *MobileDevicePrestagesMock) RegisterReplaceScopePutOnlyMock(id string) {
	m.register("PUT", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterAddScopeByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.register("POST", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterRemoveScopeByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/scope", 200, "validate_scope.json")
	m.register("POST", "/api/v2/mobile-device-prestages/"+id+"/scope/delete-multiple", 200, "validate_scope.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetAllSyncsMock() {
	m.register("GET", "/api/v2/mobile-device-prestages/syncs", 200, "validate_syncs.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetSyncsByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/syncs", 200, "validate_syncs.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetLatestSyncByIDMock(id string) {
	m.register("GET", "/api/v2/mobile-device-prestages/"+id+"/syncs/latest", 200, "validate_sync_latest.json")
}

func (m *MobileDevicePrestagesMock) RegisterGetAttachmentsByIDMock(id string) {
	m.register("GET", "/api/v3/mobile-device-prestages/"+id+"/attachments", 200, "validate_attachments.json")
}

func (m *MobileDevicePrestagesMock) RegisterUploadAttachmentMock(id string) {
	m.register("POST", "/api/v3/mobile-device-prestages/"+id+"/attachments", 200, "validate_attachment_upload.json")
}

func (m *MobileDevicePrestagesMock) RegisterDeleteAttachmentsByIDMock(id string) {
	m.register("POST", "/api/v3/mobile-device-prestages/"+id+"/attachments/delete-multiple", 200, "")
}

func (m *MobileDevicePrestagesMock) RegisterGetHistoryByIDMock(id string) {
	m.register("GET", "/api/v3/mobile-device-prestages/"+id+"/history", 200, "validate_history.json")
}

func (m *MobileDevicePrestagesMock) RegisterAddHistoryNoteByIDMock(id string) {
	m.register("POST", "/api/v3/mobile-device-prestages/"+id+"/history", 200, "validate_add_history_note.json")
}

func (m *MobileDevicePrestagesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("MobileDevicePrestagesMock: no response for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *MobileDevicePrestagesMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *MobileDevicePrestagesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *MobileDevicePrestagesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *MobileDevicePrestagesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *MobileDevicePrestagesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *MobileDevicePrestagesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *MobileDevicePrestagesMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *MobileDevicePrestagesMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
		if err := mergePage(page.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *MobileDevicePrestagesMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *MobileDevicePrestagesMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *MobileDevicePrestagesMock) InvalidateToken() error                    { return nil }
func (m *MobileDevicePrestagesMock) KeepAliveToken() error                     { return nil }
func (m *MobileDevicePrestagesMock) GetLogger() *zap.Logger                    { return m.logger }
