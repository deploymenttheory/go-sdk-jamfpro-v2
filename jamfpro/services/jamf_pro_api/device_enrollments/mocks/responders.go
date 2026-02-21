package mocks

import (
	"context"
	_ "embed"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
)

//go:embed validate_list.json
var mockListResponse []byte

//go:embed validate_get.json
var mockGetResponse []byte

//go:embed validate_history.json
var mockHistoryResponse []byte

//go:embed validate_sync_states.json
var mockSyncStatesResponse []byte

//go:embed validate_create.json
var mockCreateResponse []byte

//go:embed validate_disown.json
var mockDisownResponse []byte

type registeredResponse struct {
	method   string
	path     string
	response []byte
	status   int
}

type DeviceEnrollmentsMock struct {
	responses []registeredResponse
}

func NewDeviceEnrollmentsMock() *DeviceEnrollmentsMock {
	return &DeviceEnrollmentsMock{
		responses: make([]registeredResponse, 0),
	}
}

func (m *DeviceEnrollmentsMock) dispatch(method, path string) ([]byte, int, bool) {
	for _, r := range m.responses {
		if r.method == method && strings.HasPrefix(path, r.path) {
			return r.response, r.status, true
		}
	}
	return nil, 0, false
}

func (m *DeviceEnrollmentsMock) Get(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if strPtr, ok := out.(*string); ok {
			*strPtr = string(body)
		} else if bytesPtr, ok := out.(*[]byte); ok {
			*bytesPtr = body
		} else {
			if err := json.Unmarshal(body, out); err != nil {
				return nil, err
			}
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DeviceEnrollmentsMock) Post(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	respBody, status, found := m.dispatch("POST", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if err := json.Unmarshal(respBody, out); err != nil {
			return nil, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DeviceEnrollmentsMock) PostWithQuery(ctx context.Context, endpoint string, queryParams map[string]string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return m.Post(ctx, endpoint, body, headers, out)
}

func (m *DeviceEnrollmentsMock) PostForm(ctx context.Context, endpoint string, formData map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DeviceEnrollmentsMock) PostMultipart(ctx context.Context, endpoint string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, out any) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DeviceEnrollmentsMock) Put(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	respBody, status, found := m.dispatch("PUT", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	if out != nil {
		if err := json.Unmarshal(respBody, out); err != nil {
			return nil, err
		}
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DeviceEnrollmentsMock) Patch(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DeviceEnrollmentsMock) Delete(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	_, status, found := m.dispatch("DELETE", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DeviceEnrollmentsMock) DeleteWithBody(ctx context.Context, endpoint string, body interface{}, headers map[string]string, out interface{}) (*interfaces.Response, error) {
	_, status, found := m.dispatch("DELETE", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil
	}

	return &interfaces.Response{StatusCode: status}, nil
}

func (m *DeviceEnrollmentsMock) GetBytes(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	body, status, found := m.dispatch("GET", endpoint)
	if !found {
		return &interfaces.Response{StatusCode: http.StatusNotFound}, nil, nil
	}

	return &interfaces.Response{StatusCode: status}, body, nil
}

func (m *DeviceEnrollmentsMock) GetPaginated(ctx context.Context, endpoint string, queryParams map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	return &interfaces.Response{StatusCode: http.StatusMethodNotAllowed}, nil
}

func (m *DeviceEnrollmentsMock) RSQLBuilder() interfaces.RSQLFilterBuilder {
	return nil
}

func (m *DeviceEnrollmentsMock) InvalidateToken() error {
	return nil
}

func (m *DeviceEnrollmentsMock) KeepAliveToken() error {
	return nil
}

func (m *DeviceEnrollmentsMock) GetLogger() *zap.Logger {
	return nil
}

func RegisterListMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments",
		response: mockListResponse,
		status:   http.StatusOK,
	})
}

func RegisterGetByIDMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/",
		response: mockGetResponse,
		status:   http.StatusOK,
	})
}

func RegisterGetHistoryMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/",
		response: mockHistoryResponse,
		status:   http.StatusOK,
	})
}

func RegisterGetSyncStatesMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/",
		response: mockSyncStatesResponse,
		status:   http.StatusOK,
	})
}

func RegisterGetPublicKeyMock(mock *DeviceEnrollmentsMock) {
	publicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...\n-----END PUBLIC KEY-----"
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/public-key",
		response: []byte(publicKey),
		status:   http.StatusOK,
	})
}

func RegisterGetLatestSyncStateMock(mock *DeviceEnrollmentsMock) {
	latestSync := `{
  "syncState": "CONNECTION_ERROR",
  "instanceId": "1",
  "timestamp": "2019-04-17T14:08:06.706+0000"
}`
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/",
		response: []byte(latestSync),
		status:   http.StatusOK,
	})
}

func RegisterGetAllSyncStatesMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "GET",
		path:     "/api/v1/device-enrollments/syncs",
		response: mockSyncStatesResponse,
		status:   http.StatusOK,
	})
}

func RegisterCreateWithTokenMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v1/device-enrollments/upload-token",
		response: mockCreateResponse,
		status:   http.StatusCreated,
	})
}

func RegisterUpdateByIDMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "PUT",
		path:     "/api/v1/device-enrollments/",
		response: mockGetResponse,
		status:   http.StatusOK,
	})
}

func RegisterUpdateTokenByIDMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "PUT",
		path:     "/api/v1/device-enrollments/",
		response: mockGetResponse,
		status:   http.StatusOK,
	})
}

func RegisterDeleteByIDMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "DELETE",
		path:     "/api/v1/device-enrollments/",
		response: nil,
		status:   http.StatusNoContent,
	})
}

func RegisterDisownDevicesMock(mock *DeviceEnrollmentsMock) {
	mock.responses = append(mock.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v1/device-enrollments/",
		response: mockDisownResponse,
		status:   http.StatusOK,
	})
}

func RegisterAddHistoryNotesMock(mock *DeviceEnrollmentsMock) {
	addHistoryResp := `{
  "id": "1",
  "href": "/api/v1/device-enrollments/1/history"
}`
	mock.responses = append(mock.responses, registeredResponse{
		method:   "POST",
		path:     "/api/v1/device-enrollments/",
		response: []byte(addHistoryResp),
		status:   http.StatusCreated,
	})
}
