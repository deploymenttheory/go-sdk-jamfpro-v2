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

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	// errMsg causes the mock to return an error (simulating API or transport failures).
	errMsg string
}

// PrintersMock is a test double implementing interfaces.HTTPClient for Classic API printers.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type PrintersMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewPrintersMock returns an empty mock ready for response registration.
func NewPrintersMock() *PrintersMock {
	return &PrintersMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *PrintersMock) RegisterMocks() {
	m.RegisterListPrintersMock()
	m.RegisterGetPrinterByIDMock()
	m.RegisterGetPrinterByNameMock()
	m.RegisterCreatePrinterMock()
	m.RegisterUpdatePrinterByIDMock()
	m.RegisterUpdatePrinterByNameMock()
	m.RegisterDeletePrinterByIDMock()
	m.RegisterDeletePrinterByNameMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *PrintersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListPrintersMock registers GET /JSSResource/printers → 200.
func (m *PrintersMock) RegisterListPrintersMock() {
	m.register("GET", "/JSSResource/printers", 200, "validate_list_printers.xml")
}

// RegisterGetPrinterByIDMock registers GET /JSSResource/printers/id/1 → 200.
func (m *PrintersMock) RegisterGetPrinterByIDMock() {
	m.register("GET", "/JSSResource/printers/id/1", 200, "validate_get_printer.xml")
}

// RegisterGetPrinterByNameMock registers GET /JSSResource/printers/name/Office Printer → 200.
func (m *PrintersMock) RegisterGetPrinterByNameMock() {
	m.register("GET", "/JSSResource/printers/name/Office Printer", 200, "validate_get_printer.xml")
}

// RegisterCreatePrinterMock registers POST /JSSResource/printers/id/0 → 201.
func (m *PrintersMock) RegisterCreatePrinterMock() {
	m.register("POST", "/JSSResource/printers/id/0", 201, "validate_create_printer.xml")
}

// RegisterUpdatePrinterByIDMock registers PUT /JSSResource/printers/id/1 → 200.
func (m *PrintersMock) RegisterUpdatePrinterByIDMock() {
	m.register("PUT", "/JSSResource/printers/id/1", 200, "validate_update_printer.xml")
}

// RegisterUpdatePrinterByNameMock registers PUT /JSSResource/printers/name/Office Printer → 200.
func (m *PrintersMock) RegisterUpdatePrinterByNameMock() {
	m.register("PUT", "/JSSResource/printers/name/Office Printer", 200, "validate_update_printer.xml")
}

// RegisterDeletePrinterByIDMock registers DELETE /JSSResource/printers/id/1 → 200.
func (m *PrintersMock) RegisterDeletePrinterByIDMock() {
	m.register("DELETE", "/JSSResource/printers/id/1", 200, "")
}

// RegisterDeletePrinterByNameMock registers DELETE /JSSResource/printers/name/Office Printer → 200.
func (m *PrintersMock) RegisterDeletePrinterByNameMock() {
	m.register("DELETE", "/JSSResource/printers/name/Office Printer", 200, "")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/printers/id/999 → 404.
func (m *PrintersMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/printers/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/printers/id/0 → 409.
func (m *PrintersMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A printer with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/printers/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A printer with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *PrintersMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *PrintersMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PrintersMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PrintersMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PrintersMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PrintersMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *PrintersMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *PrintersMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PrintersMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PrintersMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *PrintersMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *PrintersMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *PrintersMock) InvalidateToken() error                    { return nil }
func (m *PrintersMock) KeepAliveToken() error                     { return nil }
func (m *PrintersMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

func (m *PrintersMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("PrintersMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *PrintersMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("PrintersMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("PrintersMock: unmarshal into result: %w", err)
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
