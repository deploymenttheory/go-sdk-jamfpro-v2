package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"resty.dev/v3"
)

// jamfPaginatedPage is the common Jamf Pro API v1 paginated response shape.
type jamfPaginatedPage struct {
	TotalCount int             `json:"totalCount"`
	Results    json.RawMessage `json:"results"`
}

// executePaginated implements requestExecutor for Transport.
// It transparently fetches all pages of a paginated Jamf Pro API endpoint,
// calling mergePage with each page's results array.
//
// Query parameters already set on req (filter, sort) are used as the base
// for every page request. page and page-size are managed internally.
// A new resty.Request is created per page so the loop is safe for retries.
//
// Jamf Pro paginated endpoints return a JSON envelope with "totalCount" and
// "results". Page numbering is zero-based; page-size defaults to DefaultPageSize.
//
// Pagination is only available on endpoints that explicitly support it.
// Example: GET /api/v3/computers-inventory
// See: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
func (t *Transport) executePaginated(req *resty.Request, path string, mergePage func([]byte) error) (*resty.Response, error) {
	// Build initial page params from query params already set on the request.
	// The caller has set filter/sort via SetQueryParam(s); we manage page/page-size.
	currentParams := make(map[string]string)
	for k, vs := range req.QueryParams {
		if len(vs) > 0 {
			currentParams[k] = vs[0]
		}
	}
	if currentParams["page"] == "" {
		currentParams["page"] = "0"
	}
	if currentParams["page-size"] == "" {
		currentParams["page-size"] = strconv.Itoa(DefaultPageSize)
	}

	// Snapshot per-request headers the service set (e.g. Accept).
	// Client-level headers (User-Agent, global headers, Authorization) are
	// applied automatically by resty to every new request created from t.client.R().
	templateHeaders := make(map[string]string)
	for k, vs := range req.Header {
		if len(vs) > 0 {
			templateHeaders[k] = vs[0]
		}
	}

	ctx := req.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var lastResp *resty.Response
	for {
		var pageResp jamfPaginatedPage
		pageReq := t.client.R().
			SetContext(ctx).
			SetResult(&pageResp).
			SetResponseBodyUnlimitedReads(true)
		for k, v := range currentParams {
			if v != "" {
				pageReq.SetQueryParam(k, v)
			}
		}
		for k, v := range templateHeaders {
			if v != "" {
				pageReq.SetHeader(k, v)
			}
		}

		resp, err := t.executeRequest(pageReq, "GET", path)
		lastResp = resp
		if err != nil {
			return lastResp, err
		}

		if err := mergePage(pageResp.Results); err != nil {
			return lastResp, fmt.Errorf("merge page: %w", err)
		}

		pageNum, _ := strconv.Atoi(currentParams["page"])
		pageSize, _ := strconv.Atoi(currentParams["page-size"])
		if pageSize <= 0 {
			pageSize = DefaultPageSize
		}
		if len(pageResp.Results) == 0 || (pageNum+1)*pageSize >= pageResp.TotalCount {
			break
		}
		currentParams["page"] = strconv.Itoa(pageNum + 1)
	}
	return lastResp, nil
}
