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

type EnrollmentCustomizationPreviewMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

func NewEnrollmentCustomizationPreviewMock() *EnrollmentCustomizationPreviewMock {
	return &EnrollmentCustomizationPreviewMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func mustMocksDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func (m *EnrollmentCustomizationPreviewMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
		if err != nil {
			panic(fmt.Sprintf("EnrollmentCustomizationPreviewMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *EnrollmentCustomizationPreviewMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := os.ReadFile(filepath.Join(mustMocksDir(), fixture))
	if err != nil {
		panic(fmt.Sprintf("EnrollmentCustomizationPreviewMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *EnrollmentCustomizationPreviewMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("EnrollmentCustomizationPreviewMock: no response registered for %s %s", method, path)
	}
	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)
	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}
	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("EnrollmentCustomizationPreviewMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

func (m *EnrollmentCustomizationPreviewMock) RegisterParseMarkdownMock() {
	m.register("POST", "/api/v1/enrollment-customization/parse-markdown", 200, "validate_parse_markdown.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetAllPanelsMock(id string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/all", 200, "validate_panels_list.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetPanelByIDMock(id, panelID string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 200, "validate_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeletePanelMock(id, panelID string) {
	m.register("DELETE", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateLdapPanelMock(id string) {
	m.register("POST", "/api/v1/enrollment-customization/"+id+"/ldap", 201, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetLdapPanelMock(id, panelID string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 200, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateLdapPanelMock(id, panelID string) {
	m.register("PUT", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 200, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteLdapPanelMock(id, panelID string) {
	m.register("DELETE", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateSsoPanelMock(id string) {
	m.register("POST", "/api/v1/enrollment-customization/"+id+"/sso", 201, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetSsoPanelMock(id, panelID string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 200, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateSsoPanelMock(id, panelID string) {
	m.register("PUT", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 200, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteSsoPanelMock(id, panelID string) {
	m.register("DELETE", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateTextPanelMock(id string) {
	m.register("POST", "/api/v1/enrollment-customization/"+id+"/text", 201, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetTextPanelMock(id, panelID string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 200, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateTextPanelMock(id, panelID string) {
	m.register("PUT", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 200, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteTextPanelMock(id, panelID string) {
	m.register("DELETE", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetTextPanelMarkdownMock(id, panelID string) {
	m.register("GET", "/api/v1/enrollment-customization/"+id+"/text/"+panelID+"/markdown", 200, "validate_text_panel_markdown.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterNotFoundErrorMock(id, panelID string) {
	m.registerError("GET", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 404, "error_not_found.json")
}

func (m *EnrollmentCustomizationPreviewMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if resp != nil {
		bodyBytes := resp.Bytes()
		if mergePage != nil && len(bodyBytes) > 0 {
			if err := mergePage(bodyBytes); err != nil {
				return resp, fmt.Errorf("mergePage failed: %w", err)
			}
		}
	}
	return resp, nil
}

func (m *EnrollmentCustomizationPreviewMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *EnrollmentCustomizationPreviewMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *EnrollmentCustomizationPreviewMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *EnrollmentCustomizationPreviewMock) InvalidateToken() error                    { return nil }
func (m *EnrollmentCustomizationPreviewMock) KeepAliveToken() error                     { return nil }
func (m *EnrollmentCustomizationPreviewMock) GetLogger() *zap.Logger                    { return m.logger }
