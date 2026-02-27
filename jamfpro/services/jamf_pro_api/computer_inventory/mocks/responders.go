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
	errMsg     string
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

func (m *ComputerInventoryMock) registerError(method, path string, errMsg string) {
	m.responses[method+":"+path] = registeredResponse{statusCode: 500, rawBody: nil, errMsg: errMsg}
}

func (m *ComputerInventoryMock) RegisterListMock() {
	m.register("GET", "/api/v3/computers-inventory", 200, "validate_list.json")
}

func (m *ComputerInventoryMock) RegisterCreateMock() {
	m.register("POST", "/api/v3/computers-inventory", 201, "validate_create.json")
}

func (m *ComputerInventoryMock) RegisterGetByIDMock(id string) {
	m.register("GET", "/api/v3/computers-inventory/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterGetDetailByIDMock(id string) {
	m.register("GET", "/api/v3/computers-inventory-detail/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterUpdateByIDMock(id string) {
	m.register("PATCH", "/api/v3/computers-inventory-detail/"+id, 200, "validate_get.json")
}

func (m *ComputerInventoryMock) RegisterDeleteByIDMock(id string) {
	m.register("DELETE", "/api/v3/computers-inventory/"+id, 204, "")
}

func (m *ComputerInventoryMock) RegisterListFileVaultMock() {
	m.register("GET", "/api/v3/computers-inventory/filevault", 200, "validate_filevault_list.json")
}

func (m *ComputerInventoryMock) RegisterGetFileVaultByIDMock(id string) {
	m.register("GET", "/api/v3/computers-inventory/"+id+"/filevault", 200, "validate_filevault.json")
}

func (m *ComputerInventoryMock) RegisterGetRecoveryLockPasswordByIDMock(id string) {
	m.register("GET", "/api/v3/computers-inventory/"+id+"/view-recovery-lock-password", 200, "validate_recovery_lock.json")
}

func (m *ComputerInventoryMock) RegisterUploadAttachmentMock(computerID string) {
	m.register("POST", "/api/v3/computers-inventory/"+computerID+"/attachments", 201, "validate_attachment.json")
}

func (m *ComputerInventoryMock) RegisterGetAttachmentMock(computerID, attachmentID string) {
	path := "/api/v3/computers-inventory/" + computerID + "/attachments/" + attachmentID
	m.register("GET", path, 200, "validate_attachment_get.json")
}

func (m *ComputerInventoryMock) RegisterDeleteAttachmentMock(computerID, attachmentID string) {
	m.register("DELETE", "/api/v3/computers-inventory/"+computerID+"/attachments/"+attachmentID, 204, "")
}

func (m *ComputerInventoryMock) RegisterGetDeviceLockPinMock(id string) {
	m.register("GET", "/api/v3/computers-inventory/"+id+"/view-device-lock-pin", 200, "validate_device_lock_pin.json")
}

func (m *ComputerInventoryMock) RegisterRemoveMDMProfileMock(id string) {
	m.register("POST", "/api/v1/computer-inventory/"+id+"/remove-mdm-profile", 200, "validate_remove_mdm.json")
}

func (m *ComputerInventoryMock) RegisterEraseMock(id string) {
	m.register("POST", "/api/v1/computer-inventory/"+id+"/erase", 204, "")
}

func (m *ComputerInventoryMock) RegisterListErrorMock() {
	m.registerError("GET", "/api/v3/computers-inventory", "simulated ListV3 API error")
}

func (m *ComputerInventoryMock) RegisterListFileVaultErrorMock() {
	m.registerError("GET", "/api/v3/computers-inventory/filevault", "simulated ListFileVault API error")
}

func (m *ComputerInventoryMock) RegisterListInvalidJSONMock() {
	m.register("GET", "/api/v3/computers-inventory", 200, "validate_list_invalid.json")
}

func (m *ComputerInventoryMock) RegisterListFileVaultInvalidJSONMock() {
	m.register("GET", "/api/v3/computers-inventory/filevault", 200, "validate_filevault_list_invalid.json")
}

func (m *ComputerInventoryMock) RegisterCreateErrorMock() {
	m.registerError("POST", "/api/v3/computers-inventory", "simulated CreateV3 API error")
}

func (m *ComputerInventoryMock) RegisterGetByIDErrorMock(id string) {
	m.registerError("GET", "/api/v3/computers-inventory/"+id, "simulated GetByID API error")
}

func (m *ComputerInventoryMock) RegisterDeleteErrorMock(id string) {
	m.registerError("DELETE", "/api/v3/computers-inventory/"+id, "simulated Delete API error")
}

func (m *ComputerInventoryMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("ComputerInventoryMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
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
	if mergePage != nil && resp != nil && len(resp.Body) > 0 {
		if err := mergePage(resp.Body); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}

func (m *ComputerInventoryMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerInventoryMock) InvalidateToken() error                    { return nil }
func (m *ComputerInventoryMock) KeepAliveToken() error                     { return nil }
func (m *ComputerInventoryMock) GetLogger() *zap.Logger                    { return m.logger }
