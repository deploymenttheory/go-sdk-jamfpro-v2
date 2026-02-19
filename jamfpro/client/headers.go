package client

import "resty.dev/v3"

// applyHeaders applies headers to a request with proper precedence:
// 1. Global headers are applied first
// 2. Per-request headers override global headers with the same key
func (t *Transport) applyHeaders(req *resty.Request, requestHeaders map[string]string) {
	for k, v := range t.globalHeaders {
		if v != "" {
			req.SetHeader(k, v)
		}
	}
	for k, v := range requestHeaders {
		if v != "" {
			req.SetHeader(k, v)
		}
	}
}
