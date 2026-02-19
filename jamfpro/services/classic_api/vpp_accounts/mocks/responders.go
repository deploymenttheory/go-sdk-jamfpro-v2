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

// VPPAccountsMock is a test double implementing interfaces.HTTPClient for Classic API VPP accounts.
type VPPAccountsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewVPPAccountsMock() *VPPAccountsMock {
	return &VPPAccountsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *VPPAccountsMock) RegisterMocks() {
	m.RegisterListVPPAccountsMock()
	m.RegisterGetVPPAccountByIDMock()
	m.RegisterCreateVPPAccountMock()
	m.RegisterUpdateVPPAccountByIDMock()
	m.RegisterDeleteVPPAccountByIDMock()
}

func (m *VPPAccountsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *VPPAccountsMock) RegisterListVPPAccountsMock() {
	m.register("GET", "/JSSResource/vppaccounts", 200, "validate_list_vpp_accounts.xml")
}
func (m *VPPAccountsMock) RegisterGetVPPAccountByIDMock() {
	m.register("GET", "/JSSResource/vppaccounts/id/1", 200, "validate_get_vpp_account.xml")
}
func (m *VPPAccountsMock) RegisterCreateVPPAccountMock() {
	m.register("POST", "/JSSResource/vppaccounts/id/0", 201, "validate_create_vpp_account.xml")
}
func (m *VPPAccountsMock) RegisterUpdateVPPAccountByIDMock() {
	m.register("PUT", "/JSSResource/vppaccounts/id/1", 200, "validate_update_vpp_account.xml")
}
func (m *VPPAccountsMock) RegisterDeleteVPPAccountByIDMock() {
	m.register("DELETE", "/JSSResource/vppaccounts/id/1", 200, "")
}
func (m *VPPAccountsMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/vppaccounts/id/999"] = registeredResponse{
		statusCode: 404, rawBody: body,
		errMsg: "Jamf Pro Classic API error (404): Resource not found",
	}
}
func (m *VPPAccountsMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A VPP account with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/vppaccounts/id/0"] = registeredResponse{
		statusCode: 409, rawBody: body,
		errMsg: "Jamf Pro Classic API error (409): A VPP account with that name already exists",
	}
}

func (m *VPPAccountsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}
func (m *VPPAccountsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VPPAccountsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VPPAccountsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VPPAccountsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *VPPAccountsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *VPPAccountsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *VPPAccountsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *VPPAccountsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *VPPAccountsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *VPPAccountsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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
func (m *VPPAccountsMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *VPPAccountsMock) InvalidateToken() error                    { return nil }
func (m *VPPAccountsMock) KeepAliveToken() error                     { return nil }
func (m *VPPAccountsMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *VPPAccountsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("VPPAccountsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *VPPAccountsMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("VPPAccountsMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("VPPAccountsMock: unmarshal into result: %w", err)
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
