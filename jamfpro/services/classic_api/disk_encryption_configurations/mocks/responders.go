package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// DiskEncryptionConfigurationsMock is a test double implementing interfaces.HTTPClient for Classic API disk encryption configurations.
type DiskEncryptionConfigurationsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewDiskEncryptionConfigurationsMock() *DiskEncryptionConfigurationsMock {
	return &DiskEncryptionConfigurationsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *DiskEncryptionConfigurationsMock) RegisterMocks() {
	m.RegisterListDiskEncryptionConfigurationsMock()
	m.RegisterGetDiskEncryptionConfigurationByIDMock()
	m.RegisterGetDiskEncryptionConfigurationByNameMock()
	m.RegisterCreateDiskEncryptionConfigurationMock()
	m.RegisterUpdateDiskEncryptionConfigurationByIDMock()
	m.RegisterUpdateDiskEncryptionConfigurationByNameMock()
	m.RegisterDeleteDiskEncryptionConfigurationByIDMock()
	m.RegisterDeleteDiskEncryptionConfigurationByNameMock()
}

func (m *DiskEncryptionConfigurationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DiskEncryptionConfigurationsMock) RegisterListDiskEncryptionConfigurationsMock() {
	m.register("GET", "/JSSResource/diskencryptionconfigurations", 200, "validate_list_disk_encryption_configurations.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterGetDiskEncryptionConfigurationByIDMock() {
	m.register("GET", "/JSSResource/diskencryptionconfigurations/id/1", 200, "validate_get_disk_encryption_configuration.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterGetDiskEncryptionConfigurationByNameMock() {
	m.register("GET", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "validate_get_disk_encryption_configuration.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterCreateDiskEncryptionConfigurationMock() {
	m.register("POST", "/JSSResource/diskencryptionconfigurations/id/0", 201, "validate_create_disk_encryption_configuration.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterUpdateDiskEncryptionConfigurationByIDMock() {
	m.register("PUT", "/JSSResource/diskencryptionconfigurations/id/1", 200, "validate_update_disk_encryption_configuration.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterUpdateDiskEncryptionConfigurationByNameMock() {
	m.register("PUT", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "validate_update_disk_encryption_configuration.xml")
}
func (m *DiskEncryptionConfigurationsMock) RegisterDeleteDiskEncryptionConfigurationByIDMock() {
	m.register("DELETE", "/JSSResource/diskencryptionconfigurations/id/1", 200, "")
}
func (m *DiskEncryptionConfigurationsMock) RegisterDeleteDiskEncryptionConfigurationByNameMock() {
	m.register("DELETE", "/JSSResource/diskencryptionconfigurations/name/FileVault Config", 200, "")
}
func (m *DiskEncryptionConfigurationsMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/diskencryptionconfigurations/id/999"] = registeredResponse{
		statusCode: 404, rawBody: body,
		errMsg: "Jamf Pro Classic API error (404): Resource not found",
	}
}
func (m *DiskEncryptionConfigurationsMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A disk encryption configuration with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/diskencryptionconfigurations/id/0"] = registeredResponse{
		statusCode: 409, rawBody: body,
		errMsg: "Jamf Pro Classic API error (409): A disk encryption configuration with that name already exists",
	}
}

func (m *DiskEncryptionConfigurationsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}
func (m *DiskEncryptionConfigurationsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DiskEncryptionConfigurationsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DiskEncryptionConfigurationsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DiskEncryptionConfigurationsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *DiskEncryptionConfigurationsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *DiskEncryptionConfigurationsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *DiskEncryptionConfigurationsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DiskEncryptionConfigurationsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *DiskEncryptionConfigurationsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *DiskEncryptionConfigurationsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *DiskEncryptionConfigurationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *DiskEncryptionConfigurationsMock) InvalidateToken() error                    { return nil }
func (m *DiskEncryptionConfigurationsMock) KeepAliveToken() error                     { return nil }
func (m *DiskEncryptionConfigurationsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *DiskEncryptionConfigurationsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("DiskEncryptionConfigurationsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *DiskEncryptionConfigurationsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("DiskEncryptionConfigurationsMock: no response registered for %s %s", method, path)
	}

	resp := &interfaces.Response{
		StatusCode: r.statusCode,
		Status:     fmt.Sprintf("%d", r.statusCode),
		Headers:    http.Header{"Content-Type": {"application/xml"}},
		Body:       r.rawBody,
	}

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("DiskEncryptionConfigurationsMock: unmarshal into result: %w", err)
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
