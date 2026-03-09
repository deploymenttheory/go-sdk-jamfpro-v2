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
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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

// RemoveableMacAddressesMock is a test double implementing transport.HTTPClient for Classic API removeable MAC addresses.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type RemoveableMacAddressesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewRemoveableMacAddressesMock returns an empty mock ready for response registration.
func NewRemoveableMacAddressesMock() *RemoveableMacAddressesMock {
	return &RemoveableMacAddressesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *RemoveableMacAddressesMock) RegisterMocks() {
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
func (m *RemoveableMacAddressesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/removablemacaddresses → 200.
func (m *RemoveableMacAddressesMock) RegisterListMock() {
	m.register("GET", "/JSSResource/removablemacaddresses", 200, "validate_list_removeable_mac_addresses.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/removablemacaddresses/id/1 → 200.
func (m *RemoveableMacAddressesMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/removablemacaddresses/id/1", 200, "validate_get_removeable_mac_address.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF → 200.
func (m *RemoveableMacAddressesMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "validate_get_removeable_mac_address.xml")
}

// RegisterCreateMock registers POST /JSSResource/removablemacaddresses/id/0 → 201.
func (m *RemoveableMacAddressesMock) RegisterCreateMock() {
	m.register("POST", "/JSSResource/removablemacaddresses/id/0", 201, "validate_create_removeable_mac_address.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/removablemacaddresses/id/1 → 200.
func (m *RemoveableMacAddressesMock) RegisterUpdateByIDMock() {
	m.register("PUT", "/JSSResource/removablemacaddresses/id/1", 200, "validate_update_removeable_mac_address.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF → 200.
func (m *RemoveableMacAddressesMock) RegisterUpdateByNameMock() {
	m.register("PUT", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "validate_update_removeable_mac_address.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/removablemacaddresses/id/1 → 200.
func (m *RemoveableMacAddressesMock) RegisterDeleteByIDMock() {
	m.register("DELETE", "/JSSResource/removablemacaddresses/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF → 200.
func (m *RemoveableMacAddressesMock) RegisterDeleteByNameMock() {
	m.register("DELETE", "/JSSResource/removablemacaddresses/name/AA:BB:CC:DD:EE:FF", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/removablemacaddresses/id/999 → 404.
func (m *RemoveableMacAddressesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/removablemacaddresses/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/removablemacaddresses/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *RemoveableMacAddressesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A removeable MAC address with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/removablemacaddresses/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A removeable MAC address with that name already exists",
	}
}

// ---- transport.HTTPClient implementation ----

func (m *RemoveableMacAddressesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *RemoveableMacAddressesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RemoveableMacAddressesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RemoveableMacAddressesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RemoveableMacAddressesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ transport.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *RemoveableMacAddressesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *RemoveableMacAddressesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *RemoveableMacAddressesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *RemoveableMacAddressesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *RemoveableMacAddressesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *RemoveableMacAddressesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *RemoveableMacAddressesMock) RSQLBuilder() transport.RSQLFilterBuilder { return nil }
func (m *RemoveableMacAddressesMock) InvalidateToken() error                    { return nil }
func (m *RemoveableMacAddressesMock) KeepAliveToken() error                     { return nil }
func (m *RemoveableMacAddressesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
func (m *RemoveableMacAddressesMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("RemoveableMacAddressesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *RemoveableMacAddressesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("RemoveableMacAddressesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("RemoveableMacAddressesMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}

// loadMockResponse reads an XML fixture file from the mocks/ directory
// adjacent to the test being run.
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
