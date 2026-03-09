package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// PatchExternalSourcesMock is a test double implementing transport.HTTPClient for Classic API patch external sources.
type PatchExternalSourcesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewPatchExternalSourcesMock() *PatchExternalSourcesMock {
	return &PatchExternalSourcesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *PatchExternalSourcesMock) RegisterMocks() {
	m.RegisterListPatchExternalSourcesMock()
	m.RegisterGetPatchExternalSourceByIDMock()
	m.RegisterGetPatchExternalSourceByNameMock()
	m.RegisterCreatePatchExternalSourceMock()
	m.RegisterUpdatePatchExternalSourceByIDMock()
	m.RegisterUpdatePatchExternalSourceByNameMock()
	m.RegisterDeletePatchExternalSourceByIDMock()
}

func (m *PatchExternalSourcesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *PatchExternalSourcesMock) RegisterListPatchExternalSourcesMock() {
	m.register("GET", "/JSSResource/patchexternalsources", 200, "validate_list_patch_external_sources.xml")
}
func (m *PatchExternalSourcesMock) RegisterGetPatchExternalSourceByIDMock() {
	m.register("GET", "/JSSResource/patchexternalsources/id/1", 200, "validate_get_patch_external_source.xml")
}
func (m *PatchExternalSourcesMock) RegisterGetPatchExternalSourceByNameMock() {
	m.register("GET", "/JSSResource/patchexternalsources/name/Primary Patch Source", 200, "validate_get_patch_external_source.xml")
}
func (m *PatchExternalSourcesMock) RegisterCreatePatchExternalSourceMock() {
	m.register("POST", "/JSSResource/patchexternalsources/id/0", 201, "validate_create_patch_external_source.xml")
}
func (m *PatchExternalSourcesMock) RegisterUpdatePatchExternalSourceByIDMock() {
	m.register("PUT", "/JSSResource/patchexternalsources/id/1", 200, "validate_update_patch_external_source.xml")
}
func (m *PatchExternalSourcesMock) RegisterUpdatePatchExternalSourceByNameMock() {
	m.register("PUT", "/JSSResource/patchexternalsources/name/Primary Patch Source", 200, "validate_update_patch_external_source.xml")
}
func (m *PatchExternalSourcesMock) RegisterDeletePatchExternalSourceByIDMock() {
	m.register("DELETE", "/JSSResource/patchexternalsources/id/1", 200, "")
}
func (m *PatchExternalSourcesMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/patchexternalsources/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}
func (m *PatchExternalSourcesMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/patchexternalsources/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A patch external source with that name already exists")
}

func (m *PatchExternalSourcesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}
func (m *PatchExternalSourcesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchExternalSourcesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchExternalSourcesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchExternalSourcesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *PatchExternalSourcesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *PatchExternalSourcesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *PatchExternalSourcesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchExternalSourcesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *PatchExternalSourcesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *PatchExternalSourcesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *PatchExternalSourcesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *PatchExternalSourcesMock) InvalidateToken() error                    { return nil }
func (m *PatchExternalSourcesMock) KeepAliveToken() error                     { return nil }
func (m *PatchExternalSourcesMock) GetLogger() *zap.Logger                    { return m.logger }

func (m *PatchExternalSourcesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("PatchExternalSourcesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

func (m *PatchExternalSourcesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("PatchExternalSourcesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *PatchExternalSourcesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("PatchExternalSourcesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("PatchExternalSourcesMock: unmarshal into result: %w", err)
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
