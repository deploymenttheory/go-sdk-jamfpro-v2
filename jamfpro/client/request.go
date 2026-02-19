package client

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// toInterfaceResponse converts a resty.Response to interfaces.Response.
func toInterfaceResponse(resp *resty.Response) *interfaces.Response {
	if resp == nil {
		return &interfaces.Response{
			Headers: make(http.Header),
		}
	}
	return &interfaces.Response{
		StatusCode: resp.StatusCode(),
		Status:     resp.Status(),
		Headers:    resp.Header(),
		Body:       []byte(resp.String()),
		Duration:   resp.Duration(),
		ReceivedAt: resp.ReceivedAt(),
		Size:       resp.Size(),
	}
}

// Get executes a GET request.
func (t *Transport) Get(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	for k, v := range rsqlQuery {
		if v != "" {
			req.SetQueryParam(k, v)
		}
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "GET", path)
}

// Post executes a POST request with JSON body.
func (t *Transport) Post(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if body != nil {
		req.SetBody(body)
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "POST", path)
}

// PostWithQuery executes a POST request with both body and query parameters.
func (t *Transport) PostWithQuery(ctx context.Context, path string, queryParams map[string]string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	for k, v := range queryParams {
		if v != "" {
			req.SetQueryParam(k, v)
		}
	}
	if body != nil {
		req.SetBody(body)
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "POST", path)
}

// PostForm executes a POST request with form-urlencoded data.
func (t *Transport) PostForm(ctx context.Context, path string, formData map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if formData != nil {
		req.SetFormData(formData)
	}
	for k, v := range t.globalHeaders {
		if v != "" && k != "Content-Type" {
			req.SetHeader(k, v)
		}
	}
	for k, v := range headers {
		if v != "" && k != "Content-Type" {
			req.SetHeader(k, v)
		}
	}
	return t.executeRequest(req, "POST", path)
}

// PostMultipart executes a POST request with multipart/form-data.
func (t *Transport) PostMultipart(ctx context.Context, path string, fileField string, fileName string, fileReader io.Reader, fileSize int64, formFields map[string]string, headers map[string]string, progressCallback interfaces.MultipartProgressCallback, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if fileReader != nil && fileName != "" && fileField != "" {
		field := &resty.MultipartField{
			Name:        fileField,
			FileName:    fileName,
			ContentType: "application/octet-stream",
			Reader:      fileReader,
			FileSize:    fileSize,
		}
		if progressCallback != nil {
			field.ProgressCallback = func(p resty.MultipartFieldProgress) {
				progressCallback(p.Name, p.FileName, p.Written, p.FileSize)
			}
		}
		req.SetMultipartFields(field)
	}
	if len(formFields) > 0 {
		req.SetMultipartFormData(formFields)
	}
	for k, v := range t.globalHeaders {
		if v != "" && k != "Content-Type" {
			req.SetHeader(k, v)
		}
	}
	for k, v := range headers {
		if v != "" && k != "Content-Type" {
			req.SetHeader(k, v)
		}
	}
	return t.executeRequest(req, "POST", path)
}

// Put executes a PUT request.
func (t *Transport) Put(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if body != nil {
		req.SetBody(body)
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "PUT", path)
}

// Patch executes a PATCH request.
func (t *Transport) Patch(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if body != nil {
		req.SetBody(body)
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "PATCH", path)
}

// Delete executes a DELETE request.
func (t *Transport) Delete(ctx context.Context, path string, queryParams map[string]string, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	for k, v := range queryParams {
		if v != "" {
			req.SetQueryParam(k, v)
		}
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "DELETE", path)
}

// DeleteWithBody executes a DELETE request with a JSON body.
func (t *Transport) DeleteWithBody(ctx context.Context, path string, body any, headers map[string]string, result any) (*interfaces.Response, error) {
	req := t.client.R().SetContext(ctx).SetResult(result)
	if body != nil {
		req.SetBody(body)
	}
	t.applyHeaders(req, headers)
	return t.executeRequest(req, "DELETE", path)
}

// GetBytes performs a GET request and returns raw bytes without unmarshaling.
func (t *Transport) GetBytes(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string) (*interfaces.Response, []byte, error) {
	req := t.client.R().SetContext(ctx)
	for k, v := range rsqlQuery {
		if v != "" {
			req.SetQueryParam(k, v)
		}
	}
	t.applyHeaders(req, headers)
	t.logger.Debug("Executing bytes request", zap.String("method", "GET"), zap.String("path", path))
	resp, err := req.Get(path)
	ifaceResp := toInterfaceResponse(resp)
	if err != nil {
		t.logger.Error("Bytes request failed", zap.String("path", path), zap.Error(err))
		return ifaceResp, nil, fmt.Errorf("bytes request failed: %w", err)
	}
	if resp.IsError() {
		return ifaceResp, nil, ParseErrorResponse(
			[]byte(resp.String()),
			resp.StatusCode(),
			resp.Status(),
			"GET",
			path,
			t.logger,
		)
	}
	return ifaceResp, []byte(resp.String()), nil
}

