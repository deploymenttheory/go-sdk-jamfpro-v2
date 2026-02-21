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

type ComputerInventoryMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewComputerInventoryMock() *ComputerInventoryMock {
	return &ComputerInventoryMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *ComputerInventoryMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		body, _ = os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerInventoryMock) RegisterListMock() {
	m.register("GET", "/api/v1/computers-inventory", 200, "validate_list.json")
}

func (m *ComputerInventoryMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v1/computers-inventory/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v1/computers-inventory/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v1/computers-inventory/"+id, 204, "")
}

func (m *ComputerInventoryMock) RegisterListFileVaultMock() {
	m.register("GET", "/api/v1/computers-inventory/filevault", 200, "validate_filevault_list.json")
}

func (m *ComputerInventoryMock) RegisterGetFileVaultByIDMock(id string) {
	m.register("GET", "/api/v1/computers-inventory/"+id+"/filevault", 200, "validate_filevault.json")
}

func (m *ComputerInventoryMock) RegisterGetRecoveryLockPasswordByIDMock(id string) {
	m.register("GET", "/api/v1/computers-inventory/"+id+"/view-recovery-lock-password", 200, "validate_recovery_lock.json")
}

func (m *ComputerInventoryMock) RegisterDeleteAttachmentMock(computerID, attachmentID string) {
	m.register("DELETE", "/api/v1/computers-inventory/"+computerID+"/attachments/"+attachmentID, 204, "")
}

func (m *ComputerInventoryMock) RegisterRemoveMDMProfileMock(id string) {
	m.register("POST", "/api/v1/computers-inventory/"+id+"/remove-mdm-profile", 200, "validate_remove_mdm.json")
}

func (m *ComputerInventoryMock) RegisterEraseMock(id string) {
	m.register("POST", "/api/v1/computers-inventory/"+id+"/erase", 204, "")
}

func (m *ComputerInventoryMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("ComputerInventoryMock: no response for %s %s", method, path)
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

func (m *ComputerInventoryMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *ComputerInventoryMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *ComputerInventoryMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *ComputerInventoryMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *ComputerInventoryMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerInventoryMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *ComputerInventoryMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *ComputerInventoryMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *ComputerInventoryMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerInventoryMock) InvalidateToken() error                     { return nil }
func (m *ComputerInventoryMock) KeepAliveToken() error                      { return nil }
func (m *ComputerInventoryMock) GetLogger() *zap.Logger                     { return m.logger }
