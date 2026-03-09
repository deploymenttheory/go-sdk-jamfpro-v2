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

// PoliciesMock is a test double implementing interfaces.HTTPClient for Classic API policies.
// Responses are keyed by "METHOD:path" and loaded from XML fixture files in
// the mocks/ directory so that expected shapes are decoupled from test code.
//
// Unlike Jamf Pro API mocks which use JSON, Classic API mocks use XML for
// serialization to match the Classic API wire format.
type PoliciesMock struct {
	responses     map[string]registeredResponse
	logger        *zap.Logger
	LastRSQLQuery map[string]string // captures the rsqlQuery from the most recent Get call
}

// NewPoliciesMock returns an empty mock ready for response registration.
func NewPoliciesMock() *PoliciesMock {
	return &PoliciesMock{
		responses: make(map[string]registeredResponse),
		logger:    zap.NewNop(),
	}
}

// RegisterMocks registers all standard success responses in one call.
func (m *PoliciesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterGetByCreatedByMock()
	m.RegisterGetByCategoryMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
}

// RegisterErrorMocks registers all error responses in one call.
func (m *PoliciesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

// ---- Success responders ----

// RegisterListMock registers GET /JSSResource/policies → 200.
func (m *PoliciesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/policies", 200, "validate_list_policies.xml")
}

// RegisterGetByIDMock registers GET /JSSResource/policies/id/1 → 200.
func (m *PoliciesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/policies/id/1", 200, "validate_get_policy.xml")
}

// RegisterGetByNameMock registers GET /JSSResource/policies/name/Test%20Policy → 200.
func (m *PoliciesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/policies/name/Test Policy", 200, "validate_get_policy.xml")
}

// RegisterCreateMock registers POST /JSSResource/policies/id/0 → 201.
func (m *PoliciesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/policies/id/0", 201, "validate_create_policy.xml")
}

// RegisterUpdateByIDMock registers PUT /JSSResource/policies/id/1 → 200.
func (m *PoliciesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/policies/id/1", 200, "validate_update_policy.xml")
}

// RegisterUpdateByNameMock registers PUT /JSSResource/policies/name/Test Policy → 200.
func (m *PoliciesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/policies/name/Test Policy", 200, "validate_update_policy.xml")
}

// RegisterDeleteByIDMock registers DELETE /JSSResource/policies/id/1 → 200.
func (m *PoliciesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/policies/id/1", 200, "")
}

// RegisterDeleteByNameMock registers DELETE /JSSResource/policies/name/Test Policy → 200.
func (m *PoliciesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/policies/name/Test Policy", 200, "")
}

// RegisterGetByCreatedByMock registers GET /JSSResource/policies/createdBy/jss → 200.
func (m *PoliciesMock) RegisterGetByCreatedByMock() {
	m.Register("GET", "/JSSResource/policies/createdBy/jss", 200, "validate_list_policies.xml")
}

// RegisterGetByCategoryMock registers GET /JSSResource/policies/category/TestCategory → 200.
func (m *PoliciesMock) RegisterGetByCategoryMock() {
	m.Register("GET", "/JSSResource/policies/category/TestCategory", 200, "validate_list_policies.xml")
}

// RegisterGetByIDWithSubsetMock registers GET /JSSResource/policies/id/1/subset/General → 200.
func (m *PoliciesMock) RegisterGetByIDWithSubsetMock() {
	m.Register("GET", "/JSSResource/policies/id/1/subset/General", 200, "validate_get_policy_subset.xml")
}

// RegisterGetByNameWithSubsetMock registers GET /JSSResource/policies/name/Test Policy/subset/General → 200.
func (m *PoliciesMock) RegisterGetByNameWithSubsetMock() {
	m.Register("GET", "/JSSResource/policies/name/Test Policy/subset/General", 200, "validate_get_policy_subset.xml")
}

// ---- Error responders ----

// RegisterNotFoundErrorMock registers GET /JSSResource/policies/id/999 → 404.
func (m *PoliciesMock) RegisterNotFoundErrorMock() {
	body := []byte("<br>An error has occurred.<br>Resource not found<br><br>")
	m.responses["GET:/JSSResource/policies/id/999"] = registeredResponse{
		statusCode: 404,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (404): Resource not found",
	}
}

// RegisterConflictErrorMock registers POST /JSSResource/policies/id/0 → 409
// when the caller wishes to simulate a duplicate-name conflict.
func (m *PoliciesMock) RegisterConflictErrorMock() {
	body := []byte("<br>An error has occurred.<br>A policy with that name already exists.<br><br>")
	m.responses["POST:/JSSResource/policies/id/0"] = registeredResponse{
		statusCode: 409,
		rawBody:    body,
		errMsg:     "Jamf Pro Classic API error (409): A policy with that name already exists",
	}
}

// ---- interfaces.HTTPClient implementation ----

func (m *PoliciesMock) Get(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	m.LastRSQLQuery = rsqlQuery
	return m.dispatch("GET", path, result)
}

func (m *PoliciesMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PoliciesMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PoliciesMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PoliciesMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*resty.Response, error) {
	return m.dispatch("POST", path, result)
}

func (m *PoliciesMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PUT", path, result)
}

func (m *PoliciesMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("PATCH", path, result)
}

func (m *PoliciesMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PoliciesMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*resty.Response, error) {
	return m.dispatch("DELETE", path, result)
}

func (m *PoliciesMock) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string) (*resty.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Bytes(), nil
}

func (m *PoliciesMock) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, _ map[string]string, mergePage func([]byte) error) (*resty.Response, error) {
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

func (m *PoliciesMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *PoliciesMock) InvalidateToken() error                    { return nil }
func (m *PoliciesMock) KeepAliveToken() error                     { return nil }
func (m *PoliciesMock) GetLogger() *zap.Logger                    { return m.logger }

// ---- Internal helpers ----

// registerError stores an error response with externalized XML body.
func (m *PoliciesMock) registerError(method, path string, statusCode int, fixture, errMsg string) {
	body, err := loadMockResponse(fixture)
	if err != nil {
		panic(fmt.Sprintf("PoliciesMock: failed to load error fixture %q: %v", fixture, err))
	}
	m.responses[method+":"+path] = registeredResponse{
		statusCode: statusCode,
		rawBody:    body,
		errMsg:     errMsg,
	}
}

// Register stores a success response keyed by "METHOD:path".
// If fixture is empty, the body is empty (used for 200/204 No Content responses).
// This method is exported so tests can register custom mock responses.
func (m *PoliciesMock) Register(method, path string, statusCode int, fixture string) {
	var body []byte
	if fixture != "" {
		data, err := loadMockResponse(fixture)
		if err != nil {
			panic(fmt.Sprintf("PoliciesMock: failed to load fixture %q: %v", fixture, err))
		}
		body = data
	}
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// dispatch looks up the registered response and either unmarshals the XML body
// into result or returns an error depending on the registration type.
func (m *PoliciesMock) dispatch(method, path string, result any) (*resty.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		headers := http.Header{"Content-Type": {mime.ApplicationXML}}
		return shared.NewMockResponse(http.StatusNotFound, headers, []byte(`<error>no mock registered</error>`)), fmt.Errorf("PoliciesMock: no response registered for %s %s", method, path)
	}

	headers := http.Header{"Content-Type": {mime.ApplicationXML}}
	resp := shared.NewMockResponse(r.statusCode, headers, r.rawBody)

	if r.errMsg != "" {
		return resp, fmt.Errorf("%s", r.errMsg)
	}

	if result != nil && len(r.rawBody) > 0 {
		if err := xml.Unmarshal(r.rawBody, result); err != nil {
			return resp, fmt.Errorf("PoliciesMock: unmarshal into result: %w", err)
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
