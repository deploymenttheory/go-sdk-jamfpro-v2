package mocks

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// SoftwareUpdateServersMock is a test double implementing client.Client for Classic API software update servers.
type SoftwareUpdateServersMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

func NewSoftwareUpdateServersMock() *SoftwareUpdateServersMock {
	return &SoftwareUpdateServersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

func (m *SoftwareUpdateServersMock) RegisterMocks() {
	m.RegisterListSoftwareUpdateServersMock()
	m.RegisterGetSoftwareUpdateServerByIDMock()
	m.RegisterGetSoftwareUpdateServerByNameMock()
	m.RegisterCreateSoftwareUpdateServerMock()
	m.RegisterUpdateSoftwareUpdateServerByIDMock()
	m.RegisterUpdateSoftwareUpdateServerByNameMock()
	m.RegisterDeleteSoftwareUpdateServerByIDMock()
	m.RegisterDeleteSoftwareUpdateServerByNameMock()
}

func (m *SoftwareUpdateServersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SoftwareUpdateServersMock) RegisterListSoftwareUpdateServersMock() {
	m.register("GET", "/JSSResource/softwareupdateservers", 200, "validate_list_software_update_servers.xml")
}
func (m *SoftwareUpdateServersMock) RegisterGetSoftwareUpdateServerByIDMock() {
	m.register("GET", "/JSSResource/softwareupdateservers/id/1", 200, "validate_get_software_update_server.xml")
}
func (m *SoftwareUpdateServersMock) RegisterGetSoftwareUpdateServerByNameMock() {
	m.register("GET", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "validate_get_software_update_server.xml")
}
func (m *SoftwareUpdateServersMock) RegisterCreateSoftwareUpdateServerMock() {
	m.register("POST", "/JSSResource/softwareupdateservers/id/0", 201, "validate_create_software_update_server.xml")
}
func (m *SoftwareUpdateServersMock) RegisterUpdateSoftwareUpdateServerByIDMock() {
	m.register("PUT", "/JSSResource/softwareupdateservers/id/1", 200, "validate_update_software_update_server.xml")
}
func (m *SoftwareUpdateServersMock) RegisterUpdateSoftwareUpdateServerByNameMock() {
	m.register("PUT", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "validate_update_software_update_server.xml")
}
func (m *SoftwareUpdateServersMock) RegisterDeleteSoftwareUpdateServerByIDMock() {
	m.register("DELETE", "/JSSResource/softwareupdateservers/id/1", 200, "")
}
func (m *SoftwareUpdateServersMock) RegisterDeleteSoftwareUpdateServerByNameMock() {
	m.register("DELETE", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "")
}
func (m *SoftwareUpdateServersMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/softwareupdateservers/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}
func (m *SoftwareUpdateServersMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/softwareupdateservers/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A software update server with that name already exists")
}

func (m *SoftwareUpdateServersMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}
func (m *SoftwareUpdateServersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SoftwareUpdateServersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SoftwareUpdateServersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SoftwareUpdateServersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *SoftwareUpdateServersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *SoftwareUpdateServersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *SoftwareUpdateServersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SoftwareUpdateServersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *SoftwareUpdateServersMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}
func (m *SoftwareUpdateServersMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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
func (m *SoftwareUpdateServersMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilder(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	})
}
func (m *SoftwareUpdateServersMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *SoftwareUpdateServersMock) InvalidateToken() error                { return nil }
func (m *SoftwareUpdateServersMock) KeepAliveToken() error                 { return nil }
func (m *SoftwareUpdateServersMock) GetLogger() *zap.Logger                { return m.logger }

func (m *SoftwareUpdateServersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("SoftwareUpdateServersMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *SoftwareUpdateServersMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("SoftwareUpdateServersMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

func (m *SoftwareUpdateServersMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {constants.ApplicationXML}}
		return mockhelpers.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("SoftwareUpdateServersMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {constants.ApplicationXML}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("SoftwareUpdateServersMock: unmarshal into result: %w", err)
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
