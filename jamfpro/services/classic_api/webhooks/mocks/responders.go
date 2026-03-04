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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// WebhooksMock is a test double implementing interfaces.HTTPClient for Classic API webhooks.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type WebhooksMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewWebhooksMock returns an empty mock ready for response registration.
func NewWebhooksMock() *WebhooksMock {
	return &WebhooksMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *WebhooksMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *WebhooksMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/webhooks → 200.
func (m *WebhooksMock) RegisterListMock() {
	m.register("GET", "/JSSResource/webhooks", 200, "validate_list_webhooks.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/webhooks/id/1 → 200.
func (m *WebhooksMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/webhooks/id/1", 200, "validate_get_webhook.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/webhooks/name/Computer Enrolled → 200.
func (m *WebhooksMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/webhooks/name/Computer Enrolled", 200, "validate_get_webhook.xml")
}

// RegisterCreateMock registers POST /JSSResource/webhooks/id/0 → 201.
func (m *WebhooksMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/webhooks/id/0", 201, "validate_create_webhook.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/webhooks/id/1 → 200.
func (m *WebhooksMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/webhooks/id/1", 200, "validate_update_webhook.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/webhooks/name/Computer Enrolled → 200.
func (m *WebhooksMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/webhooks/name/Computer Enrolled", 200, "validate_update_webhook.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/webhooks/id/1 → 200.
func (m *WebhooksMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/webhooks/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/webhooks/name/Computer Enrolled → 200.
func (m *WebhooksMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/webhooks/name/Computer Enrolled", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/webhooks/id/999 → 404.
func (m *WebhooksMock) RegisterNotFoundErrorMock() {
	m.registerError("GET", "/JSSResource/webhooks/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

// RegisterConflictErrorMock registers POST /JSSResource/webhooks/id/0 → 409.
func (m *WebhooksMock) RegisterConflictErrorMock() {
	m.registerError("POST", "/JSSResource/webhooks/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A webhook with that name already exists")
}

// ---- interfaces.HTTPClient implementation ----

func (m *WebhooksMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *WebhooksMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *WebhooksMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *WebhooksMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *WebhooksMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *WebhooksMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *WebhooksMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *WebhooksMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *WebhooksMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *WebhooksMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *WebhooksMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *WebhooksMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *WebhooksMock) InvalidateToken() error                    { return nil }
func (m *WebhooksMock) KeepAliveToken() error                     { return nil }
func (m *WebhooksMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

func (m *WebhooksMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("WebhooksMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *WebhooksMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("WebhooksMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

func (m *WebhooksMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("WebhooksMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("WebhooksMock: unmarshal into result: %w", err)
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
