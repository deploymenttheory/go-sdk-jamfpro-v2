package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// PackagesMock is a test double implementing interfaces.HTTPClient.
type PackagesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewPackagesMock returns an empty mock ready for response registration.
func NewPackagesMock() *PackagesMock {
	return &PackagesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *PackagesMock) RegisterMocks() {
	m.RegisterListPackagesMock()
	m.RegisterGetPackageMock()
	m.RegisterCreatePackageMock()
	m.RegisterUpdatePackageMock()
	m.RegisterDeletePackageMock()
	m.RegisterDeletePackagesByIDMock()
	m.RegisterGetPackageHistoryMock()
	m.RegisterAddPackageHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *PackagesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PackagesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("PackagesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *PackagesMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("PackagesMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *PackagesMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/json"}},
			Body:       []byte(`{"code":"NOT-FOUND","message":"no mock registered"}`),
		}, fmt.Errorf("PackagesMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/json"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("PackagesMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get working directory: %w", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "mocks", filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}

func (m *PackagesMock) RegisterListPackagesMock() {
	m.register("GET", "/api/v1/packages", 200, "validate_list_packages.json")
}

func (m *PackagesMock) RegisterListPackagesRSQLMock() {
	m.register("GET", "/api/v1/packages", 200, "validate_list_packages_rsql.json")
}

func (m *PackagesMock) RegisterGetPackageMock() {
	m.register("GET", "/api/v1/packages/1", 200, "validate_get_package.json")
}

func (m *PackagesMock) RegisterCreatePackageMock() {
	m.register("POST", "/api/v1/packages", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterUpdatePackageMock() {
	m.register("PUT", "/api/v1/packages/1", 200, "validate_update_package.json")
}

func (m *PackagesMock) RegisterUploadPackageMock() {
	m.register("POST", "/api/v1/packages/1/upload", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterAssignManifestMock() {
	m.register("POST", "/api/v1/packages/1/manifest", 201, "validate_create_package.json")
}

func (m *PackagesMock) RegisterDeleteManifestMock() {
	m.register("DELETE", "/api/v1/packages/1/manifest", 204, "")
}

func (m *PackagesMock) RegisterDeletePackageMock() {
	m.register("DELETE", "/api/v1/packages/1", 204, "")
}

func (m *PackagesMock) RegisterDeletePackagesByIDMock() {
	m.register("POST", "/api/v1/packages/delete-multiple", 204, "")
}

func (m *PackagesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/packages/999", 404, "error_not_found.json")
}

func (m *PackagesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/packages", 409, "error_conflict.json")
}

func (m *PackagesMock) RegisterGetPackageHistoryMock() {
	m.register("GET", "/api/v1/packages/1/history", 200, "validate_get_history.json")
}

func (m *PackagesMock) RegisterAddPackageHistoryNotesMock() {
	m.register("POST", "/api/v1/packages/1/history", 201, "")
}

func (m *PackagesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *PackagesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PackagesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PackagesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PackagesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PackagesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *PackagesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *PackagesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PackagesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PackagesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *PackagesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		if err := mergePage(resp.Body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *PackagesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *PackagesMock) InvalidateToken() error                    { return nil }
func (m *PackagesMock) KeepAliveToken() error                      { return nil }
func (m *PackagesMock) GetLogger() *zap.Logger                     { return m.logger }
