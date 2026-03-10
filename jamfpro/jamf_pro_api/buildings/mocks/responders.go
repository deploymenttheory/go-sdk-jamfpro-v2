package mocks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"go.uber.org/zap"
	"resty.dev/v3"

	mockhelpers "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// BuildingsMock is a test double implementing client.Client.
type BuildingsMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string
}

// NewBuildingsMock returns an empty mock ready for response registration.
func NewBuildingsMock() *BuildingsMock {
	return &BuildingsMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *BuildingsMock) RegisterMocks() {
	m.RegisterListBuildingsMock()
	m.RegisterGetBuildingMock()
	m.RegisterCreateBuildingMock()
	m.RegisterUpdateBuildingMock()
	m.RegisterDeleteBuildingMock()
	m.RegisterDeleteBuildingsByIDMock()
	m.RegisterGetBuildingHistoryMock()
	m.RegisterAddBuildingHistoryNotesMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *BuildingsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *BuildingsMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("BuildingsMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *BuildingsMock) registerError(method, path string, statusCode int, fixture string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("BuildingsMock: failed to load error fixture %q: %v", fixture, err))
	}
	var parsed struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	_ = json.Unmarshal(body, &parsed)
	errMsg := fmt.Sprintf("Jamf Pro API error (%d) [%s]: %s", statusCode, parsed.Code, parsed.Message)
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body, errMsg: errMsg}
}

func (m *BuildingsMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return nil, fmt.Errorf("BuildingsMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {"application/json"}}
	resp := mockhelpers.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := json.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("BuildingsMock: unmarshal into result: %w", err)
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

func (m *BuildingsMock) RegisterListBuildingsMock() {
	m.register("GET", "/api/v1/buildings", 200, "validate_list_buildings.json")
}

func (m *BuildingsMock) RegisterListBuildingsRSQLMock() {
	m.register("GET", "/api/v1/buildings", 200, "validate_list_buildings_rsql.json")
}

func (m *BuildingsMock) RegisterGetBuildingMock() {
	m.register("GET", "/api/v1/buildings/1", 200, "validate_get_building.json")
}

func (m *BuildingsMock) RegisterCreateBuildingMock() {
	m.register("POST", "/api/v1/buildings", 201, "validate_create_building.json")
}

func (m *BuildingsMock) RegisterUpdateBuildingMock() {
	m.register("PUT", "/api/v1/buildings/1", 200, "validate_update_building.json")
}

func (m *BuildingsMock) RegisterDeleteBuildingMock() {
	m.register("DELETE", "/api/v1/buildings/1", 204, "")
}

func (m *BuildingsMock) RegisterDeleteBuildingsByIDMock() {
	m.register("POST", "/api/v1/buildings/delete-multiple", 204, "")
}

func (m *BuildingsMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/api/v1/buildings/999", 404, "error_not_found.json")
}

func (m *BuildingsMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/api/v1/buildings", 409, "error_conflict.json")
}

func (m *BuildingsMock) RegisterGetBuildingHistoryMock() {
	m.register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history.json")
}

func (m *BuildingsMock) RegisterGetBuildingHistoryNullDetailsMock() {
	m.register("GET", "/api/v1/buildings/1/history", 200, "validate_get_history_null_details.json")
}

func (m *BuildingsMock) RegisterAddBuildingHistoryNotesMock() {
	m.register("POST", "/api/v1/buildings/1/history", 201, "")
}

func (m *BuildingsMock) RegisterExportBuildingsMock() {
	m.register("POST", "/api/v1/buildings/export", 200, "validate_list_buildings.json")
}

func (m *BuildingsMock) RegisterExportBuildingHistoryMock() {
	m.register("POST", "/api/v1/buildings/1/history/export", 200, "validate_get_history.json")
}

func (m *BuildingsMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *BuildingsMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BuildingsMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BuildingsMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BuildingsMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ client.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *BuildingsMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *BuildingsMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *BuildingsMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *BuildingsMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *BuildingsMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *BuildingsMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil {
		bodyBytes := resp.Bytes()
		var pageResp struct {
			TotalCount int             `json:"totalCount"`
			Results    json.RawMessage `json:"results"`
		}
		if err := json.Unmarshal(bodyBytes, &pageResp); err != nil {
			return resp, fmt.Errorf("unmarshal page response: %w", err)
		}

		if err := mergePage(pageResp.Results); err != nil {
			return resp, fmt.Errorf("mergePage failed: %w", err)
		}
	}
	return resp, nil
}
func (m *BuildingsMock) NewRequest(ctx context.Context) *client.RequestBuilder {
	return client.NewMockRequestBuilderWithQueryCapture(ctx, func(method, path string, result any) (*resty.Response, error) {
		return m.dispatch(method, path, result)
	}, &m.LastRSQLQuery)
}

func (m *BuildingsMock) RSQLBuilder() client.RSQLFilterBuilder { return nil }
func (m *BuildingsMock) InvalidateToken() error                { return nil }
func (m *BuildingsMock) KeepAliveToken() error                 { return nil }
func (m *BuildingsMock) GetLogger() *zap.Logger                { return m.logger }
