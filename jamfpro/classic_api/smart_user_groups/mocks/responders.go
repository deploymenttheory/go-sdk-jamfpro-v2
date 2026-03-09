package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

type SmartUserGroupsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewSmartUserGroupsMock() *SmartUserGroupsMock {
	return &SmartUserGroupsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *SmartUserGroupsMock) RegisterMocks() {
	m.RegisterListSmartUserGroupsMock()
	m.RegisterGetSmartUserGroupByIDMock()
	m.RegisterGetSmartUserGroupByNameMock()
	m.RegisterCreateSmartUserGroupMock()
	m.RegisterUpdateSmartUserGroupByIDMock()
	m.RegisterUpdateSmartUserGroupByNameMock()
	m.RegisterDeleteSmartUserGroupByIDMock()
	m.RegisterDeleteSmartUserGroupByNameMock()
}

func (m *SmartUserGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SmartUserGroupsMock) RegisterListSmartUserGroupsMock() {
	m.register("GET", "/JSSResource/usergroups", 200, "validate_list_user_groups.xml")
}

func (m *SmartUserGroupsMock) RegisterGetSmartUserGroupByIDMock() {
	m.register("GET", "/JSSResource/usergroups/id/1", 200, "validate_get_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterGetSmartUserGroupByNameMock() {
	m.register("GET", "/JSSResource/usergroups/name/All Users", 200, "validate_get_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterCreateSmartUserGroupMock() {
	m.register("POST", "/JSSResource/usergroups/id/0", 201, "validate_create_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterUpdateSmartUserGroupByIDMock() {
	m.register("PUT", "/JSSResource/usergroups/id/1", 200, "validate_update_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterUpdateSmartUserGroupByNameMock() {
	m.register("PUT", "/JSSResource/usergroups/name/All Users", 200, "validate_update_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterDeleteSmartUserGroupByIDMock() {
	m.register("DELETE", "/JSSResource/usergroups/id/1", 200, "")
}

func (m *SmartUserGroupsMock) RegisterDeleteSmartUserGroupByNameMock() {
	m.register("DELETE", "/JSSResource/usergroups/name/All Users", 200, "")
}

func (m *SmartUserGroupsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/usergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *SmartUserGroupsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/usergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user group with that name already exists")
}

func (m *SmartUserGroupsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *SmartUserGroupsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartUserGroupsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartUserGroupsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartUserGroupsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *SmartUserGroupsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *SmartUserGroupsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *SmartUserGroupsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SmartUserGroupsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *SmartUserGroupsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *SmartUserGroupsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		body := resp.Bytes()
		if err := mergePage(body); err != nil {
			return resp, err
		}
	}
	return resp, nil
}

func (m *SmartUserGroupsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *SmartUserGroupsMock) InvalidateToken() error                    { return nil }
func (m *SmartUserGroupsMock) KeepAliveToken() error                     { return nil }
func (m *SmartUserGroupsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *SmartUserGroupsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("SmartUserGroupsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SmartUserGroupsMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("SmartUserGroupsMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

func (m *SmartUserGroupsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("SmartUserGroupsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("SmartUserGroupsMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func loadMockResponse(filename string) ([]byte, error) {
	_, callerPath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("runtime.Caller failed")
	}
	dir := filepath.Dir(callerPath)
	data, err := os.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return nil, fmt.Errorf("read fixture %s: %w", filename, err)
	}
	return data, nil
}
