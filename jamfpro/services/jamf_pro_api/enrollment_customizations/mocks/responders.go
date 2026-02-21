package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_list.json
var validateListJSON []byte

//go:embed validate_get.json
var validateGetJSON []byte

//go:embed validate_create.json
var validateCreateJSON []byte

//go:embed validate_history.json
var validateHistoryJSON []byte

//go:embed validate_add_history_notes.json
var validateAddHistoryNotesJSON []byte

//go:embed validate_prestages.json
var validatePrestagesJSON []byte

type registeredResponse struct {
	method   string
	path     string
	response []byte
	status   int
}

type EnrollmentCustomizationsMock struct {
	responses []registeredResponse
}

func NewEnrollmentCustomizationsMock() *EnrollmentCustomizationsMock {
	return &EnrollmentCustomizationsMock{
		responses: make([]registeredResponse, 0),
	}
}

func (m *EnrollmentCustomizationsMock) dispatch(method, path string) ([]byte, int, bool) {
	var bestMatch *registeredResponse
	longestMatch := 0

	for i := range m.responses {
		r := &m.responses[i]
		if r.method == method && strings.HasPrefix(path, r.path) {
			if len(r.path) > longestMatch {
				longestMatch = len(r.path)
				bestMatch = r
			}
		}
	}

	if bestMatch != nil {
		return bestMatch.response, bestMatch.status, true
	}
	return nil, 0, false
}

func (m *EnrollmentCustomizationsMock) Get(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(body)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = body
		} else if err := json.Unmarshal(body, out); err != nil {
			return &interfaces.Response{StatusCode: status}, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *EnrollmentCustomizationsMock) Post(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	respBody, status, found := m.dispatch("POST", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(respBody)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = respBody
		} else if err := json.Unmarshal(respBody, out); err != nil {
			return &interfaces.Response{StatusCode: status}, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *EnrollmentCustomizationsMock) PostWithQuery(ctx context.Context, endpoint string, queryParams map[string]string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *EnrollmentCustomizationsMock) PostForm(ctx context.Context, endpoint string, formData map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *EnrollmentCustomizationsMock) PostMultipart(ctx context.Context, endpoint string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out interface{}) (*interfaces.Response, error) {
	respBody, status, found := m.dispatch("POST_MULTIPART", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(respBody)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = respBody
		} else if err := json.Unmarshal(respBody, out); err != nil {
			return &interfaces.Response{StatusCode: status}, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *EnrollmentCustomizationsMock) Put(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	respBody, status, found := m.dispatch("PUT", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(respBody)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = respBody
		} else if err := json.Unmarshal(respBody, out); err != nil {
			return &interfaces.Response{StatusCode: status}, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *EnrollmentCustomizationsMock) Patch(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *EnrollmentCustomizationsMock) Delete(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	_, status, found := m.dispatch("DELETE", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *EnrollmentCustomizationsMock) DeleteWithBody(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *EnrollmentCustomizationsMock) GetBytes(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil, nil
	}

	return &interfaces.Response{StatusCode: status}, body, nil
}

func (m *EnrollmentCustomizationsMock) GetPaginated(ctx context.Context, endpoint string, rsqlQuery map[string]string, headers map[string]string, mergePage func(page []byte) error) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *EnrollmentCustomizationsMock) GetLogger() *zap.Logger {
	return nil
}

func (m *EnrollmentCustomizationsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *EnrollmentCustomizationsMock) InvalidateToken() error {
	return nil
}

func (m *EnrollmentCustomizationsMock) KeepAliveToken() error {
	return nil
}

func (m *EnrollmentCustomizationsMock) RegisterListMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/enrollment-customizations",
		response: validateListJSON,
		status:   http.StatusOK,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterGetByIDMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/enrollment-customizations/",
		response: validateGetJSON,
		status:   http.StatusOK,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterCreateMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v2/enrollment-customizations",
		response: validateCreateJSON,
		status:   http.StatusCreated,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterUpdateMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "PUT",
		path:     "/api/v2/enrollment-customizations/",
		response: validateGetJSON,
		status:   http.StatusOK,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterDeleteMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "DELETE",
		path:     "/api/v2/enrollment-customizations/",
		response: nil,
		status:   http.StatusNoContent,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterGetHistoryMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/enrollment-customizations/",
		response: validateHistoryJSON,
		status:   http.StatusOK,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterAddHistoryNotesMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v2/enrollment-customizations/",
		response: validateAddHistoryNotesJSON,
		status:   http.StatusCreated,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterGetPrestagesMock() {
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v2/enrollment-customizations/",
		response: validatePrestagesJSON,
		status:   http.StatusOK,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterUploadImageMock() {
	response := []byte(`{"id":"123","url":"https://example.jamfcloud.com/enrollment-customizations/images/123"}`)
	m.responses = append(m.responses, registeredResponse{
		method:   "POST_MULTIPART",
		path:     "/api/v2/enrollment-customizations/images",
		response: response,
		status:   http.StatusCreated,
	})
}

func (m *EnrollmentCustomizationsMock) RegisterGetImageByIdMock(id string) {
	// Mock image data (small PNG header)
	imageData := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	m.responses = append(m.responses, registeredResponse{
		method:   "GET",
		path:     fmt.Sprintf("/api/v2/enrollment-customizations/images/%s", id),
		response: imageData,
		status:   http.StatusOK,
	})
}
