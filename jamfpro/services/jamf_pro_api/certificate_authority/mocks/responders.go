package mocks

import (
	"context"
	"encoding/json"
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
}

// CertificateAuthorityMock is a test double implementing interfaces.HTTPClient.
type CertificateAuthorityMock struct {
	responses map[string]registeredResponse
	logger    *zap.Logger
}

// NewCertificateAuthorityMock returns an empty mock ready for response registration.
func NewCertificateAuthorityMock() *CertificateAuthorityMock {
	return &CertificateAuthorityMock{responses: make(map[string]registeredResponse), logger: zap.NewNop()}
}

func (m *CertificateAuthorityMock) register(method, path string, statusCode int, fixture string) {
	body, _ := os.ReadFile(filepath.Join(mustGetwd(), "mocks", fixture))
	m.responses[method+":"+path] = registeredResponse{statusCode: statusCode, rawBody: body}
}

// RegisterGetActiveCertificateAuthorityMock registers a successful GET response.
func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityMock() {
	m.register("GET", "/api/v1/pki/certificate-authority/active", 200, "validate_get.json")
}

// RegisterGetActiveCertificateAuthorityDERMock registers GET .../active/der.
func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityDERMock() {
	m.register("GET", "/api/v1/pki/certificate-authority/active/der", 200, "validate_active_der.der")
}

// RegisterGetActiveCertificateAuthorityPEMMock registers GET .../active/pem.
func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityPEMMock() {
	m.register("GET", "/api/v1/pki/certificate-authority/active/pem", 200, "validate_active_pem.pem")
}

// RegisterGetCertificateAuthorityByIDMock registers GET .../{id}, .../{id}/der, .../{id}/pem for the given id.
func (m *CertificateAuthorityMock) RegisterGetCertificateAuthorityByIDMock(id string) {
	base := "/api/v1/pki/certificate-authority/" + id
	m.register("GET", base, 200, "validate_get.json")
	m.register("GET", base+"/der", 200, "validate_active_der.der")
	m.register("GET", base+"/pem", 200, "validate_active_pem.pem")
}

func (m *CertificateAuthorityMock) dispatch(method, path string, result any) (*interfaces.Response, error) {
	r, ok := m.responses[method+":"+path]
	if !ok {
		return &interfaces.Response{StatusCode: 404, Headers: http.Header{}, Body: nil}, fmt.Errorf("CertificateAuthorityMock: no response for %s %s", method, path)
	}
	resp := &interfaces.Response{StatusCode: r.statusCode, Status: fmt.Sprintf("%d", r.statusCode), Headers: http.Header{"Content-Type": {"application/json"}}, Body: r.rawBody}
	if result != nil && len(r.rawBody) > 0 {
		_ = json.Unmarshal(r.rawBody, result)
	}
	return resp, nil
}

func mustGetwd() string {
	dir, _ := os.Getwd()
	return dir
}

func (m *CertificateAuthorityMock) Get(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("GET", path, result)
}
func (m *CertificateAuthorityMock) Post(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CertificateAuthorityMock) PostWithQuery(ctx context.Context, path string, _ map[string]string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CertificateAuthorityMock) PostForm(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CertificateAuthorityMock) PostMultipart(ctx context.Context, path string, _ string, _ string, _ io.Reader, _ int64, _ map[string]string, _ map[string]string, _ interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	return m.dispatch("POST", path, result)
}
func (m *CertificateAuthorityMock) Put(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PUT", path, result)
}
func (m *CertificateAuthorityMock) Patch(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("PATCH", path, result)
}
func (m *CertificateAuthorityMock) Delete(ctx context.Context, path string, _ map[string]string, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CertificateAuthorityMock) DeleteWithBody(ctx context.Context, path string, _ any, _ map[string]string, result any) (*interfaces.Response, error) {
	return m.dispatch("DELETE", path, result)
}
func (m *CertificateAuthorityMock) GetBytes(ctx context.Context, path string, _ map[string]string, _ map[string]string) (*interfaces.Response, []byte, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, nil, err
	}
	return resp, resp.Body, nil
}
func (m *CertificateAuthorityMock) GetPaginated(ctx context.Context, path string, _ map[string]string, _ map[string]string, mergePage func([]byte) error) (*interfaces.Response, error) {
	resp, err := m.dispatch("GET", path, nil)
	if err != nil {
		return resp, err
	}
	if mergePage != nil && len(resp.Body) > 0 {
		_ = mergePage(resp.Body)
	}
	return resp, nil
}
func (m *CertificateAuthorityMock) RSQLBuilder() interfaces.RSQLFilterBuilder { return nil }
func (m *CertificateAuthorityMock) InvalidateToken() error                    { return nil }
func (m *CertificateAuthorityMock) KeepAliveToken() error                       { return nil }
func (m *CertificateAuthorityMock) GetLogger() *zap.Logger                      { return m.logger }
