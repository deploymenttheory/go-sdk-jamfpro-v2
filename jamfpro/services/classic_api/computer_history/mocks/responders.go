package mocks

import (
	"context"
	_ "embed"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_get_computer_history.xml
var validateGetComputerHistoryXML []byte

// registeredResponse holds a pre-canned response for a single endpoint.
type registeredResponse struct {
	statusCode int
	rawBody    []byte
	errMsg     string
}

// ComputerHistoryMock is a test double implementing interfaces.HTTPClient for Classic API computer history.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Classic API mocks use XML for serialization to match the Classic API wire format.
type ComputerHistoryMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewComputerHistoryMock returns an empty mock ready for response registration.
func NewComputerHistoryMock() *ComputerHistoryMock {
	return &ComputerHistoryMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *ComputerHistoryMock) RegisterMocks() {
	m.RegisterGetByIDMock()
	m.RegisterGetByIDAndSubsetMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByNameAndSubsetMock()
	m.RegisterGetByUDIDMock()
	m.RegisterGetByUDIDAndSubsetMock()
	m.RegisterGetBySerialNumberMock()
	m.RegisterGetBySerialNumberAndSubsetMock()
	m.RegisterGetByMACAddressMock()
	m.RegisterGetByMACAddressAndSubsetMock()
}

// ---- Success responders ----

// RegisterGetByIDMock registers GET /JSSResource/computerhistory/id/1 → 200.
func (m *ComputerHistoryMock) RegisterGetByIDMock() {
	m.register("GET", "/JSSResource/computerhistory/id/1", 200, "validate_get_computer_history.xml")
}

// RegisterGetByIDAndSubsetMock registers GET /JSSResource/computerhistory/id/1/subset/General → 200.
func (m *ComputerHistoryMock) RegisterGetByIDAndSubsetMock() {
	m.register("GET", "/JSSResource/computerhistory/id/1/subset/General", 200, "validate_get_computer_history.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/computerhistory/name/Test-MacBook-Pro → 200.
func (m *ComputerHistoryMock) RegisterGetByNameMock() {
	m.register("GET", "/JSSResource/computerhistory/name/Test-MacBook-Pro", 200, "validate_get_computer_history.xml")
}

// RegisterGetByNameAndSubsetMock registers GET /JSSResource/computerhistory/name/Test-MacBook-Pro/subset/General → 200.
func (m *ComputerHistoryMock) RegisterGetByNameAndSubsetMock() {
	m.register("GET", "/JSSResource/computerhistory/name/Test-MacBook-Pro/subset/General", 200, "validate_get_computer_history.xml")
}

// RegisterGetByUDIDMock registers GET /JSSResource/computerhistory/udid/00000000-0000-0000-0000-000000000001 → 200.
func (m *ComputerHistoryMock) RegisterGetByUDIDMock() {
	m.register("GET", "/JSSResource/computerhistory/udid/00000000-0000-0000-0000-000000000001", 200, "validate_get_computer_history.xml")
}

// RegisterGetByUDIDAndSubsetMock registers GET /JSSResource/computerhistory/udid/.../subset/General → 200.
func (m *ComputerHistoryMock) RegisterGetByUDIDAndSubsetMock() {
	m.register("GET", "/JSSResource/computerhistory/udid/00000000-0000-0000-0000-000000000001/subset/General", 200, "validate_get_computer_history.xml")
}

// RegisterGetBySerialNumberMock registers GET /JSSResource/computerhistory/serialnumber/C02XYZ123456 → 200.
func (m *ComputerHistoryMock) RegisterGetBySerialNumberMock() {
	m.register("GET", "/JSSResource/computerhistory/serialnumber/C02XYZ123456", 200, "validate_get_computer_history.xml")
}

// RegisterGetBySerialNumberAndSubsetMock registers GET /JSSResource/computerhistory/serialnumber/.../subset/General → 200.
func (m *ComputerHistoryMock) RegisterGetBySerialNumberAndSubsetMock() {
	m.register("GET", "/JSSResource/computerhistory/serialnumber/C02XYZ123456/subset/General", 200, "validate_get_computer_history.xml")
}

// RegisterGetByMACAddressMock registers GET /JSSResource/computerhistory/macaddress/00%3A11%3A22%3A33%3A44%3A55 → 200.
func (m *ComputerHistoryMock) RegisterGetByMACAddressMock() {
	m.register("GET", "/JSSResource/computerhistory/macaddress/00%3A11%3A22%3A33%3A44%3A55", 200, "validate_get_computer_history.xml")
}

// RegisterGetByMACAddressAndSubsetMock registers GET /JSSResource/computerhistory/macaddress/.../subset/General → 200.
func (m *ComputerHistoryMock) RegisterGetByMACAddressAndSubsetMock() {
	m.register("GET", "/JSSResource/computerhistory/macaddress/00%3A11%3A22%3A33%3A44%3A55/subset/General", 200, "validate_get_computer_history.xml")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/computerhistory/id/999 → 404.
func (m *ComputerHistoryMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/computerhistory/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *ComputerHistoryMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}

func (m *ComputerHistoryMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerHistoryMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerHistoryMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerHistoryMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *ComputerHistoryMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *ComputerHistoryMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *ComputerHistoryMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerHistoryMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *ComputerHistoryMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}

func (m *ComputerHistoryMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
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

func (m *ComputerHistoryMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *ComputerHistoryMock) InvalidateToken() error                     { return nil }
func (m *ComputerHistoryMock) KeepAliveToken() error                      { return nil }
func (m *ComputerHistoryMock) GetLogger() *zap.Logger                     { return m.logger }

// ---- Internal helpers ----

func (m *ComputerHistoryMock) register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture == "validate_get_computer_history.xml" {
		body = validateGetComputerHistoryXML
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

func (m *ComputerHistoryMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{
			StatusCode: http.StatusNotFound,
			Status:     "404 Not Found",
			Headers:    http.Header{"Content-Type": {"application/xml"}},
			Body:       []byte(`<error>no mock registered</error>`),
		}, fmt.Errorf("ComputerHistoryMock: no response registered for %s %s", method, path)
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
			return resp, fmt.Errorf("ComputerHistoryMock: unmarshal into result: %w", err)
		}
	}
	return resp, nil
}
