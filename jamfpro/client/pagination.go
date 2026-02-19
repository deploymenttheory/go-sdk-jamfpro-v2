package client

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"net/url"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
)

// PaginationLinks contains pagination navigation links (for cursor-style APIs).
type PaginationLinks struct {
	Self string `json:"self"`
	Next string `json:"next,omitempty"`
}

// jamfPaginatedPage is the common Jamf Pro API v1 paginated response shape.
type jamfPaginatedPage struct {
	TotalCount int             `json:"totalCount"`
	Results    json.RawMessage `json:"results"`
}

// GetPaginated executes a paginated GET request, transparently fetching all
// pages and merging them via the caller-supplied mergePage function.
//
// Jamf Pro paginated endpoints return a JSON envelope with "totalCount" and
// "results". Page numbering is zero-based; page-size defaults to DefaultPageSize.
// rsqlQuery may include:
//   - "filter"    – RSQL expression to narrow results (use RSQLBuilder)
//   - "sort"      – sort field and direction (e.g. "general.name:asc")
//   - "page"      – override starting page (default 0)
//   - "page-size" – override page size (default DefaultPageSize)
//
// Pagination is only available on endpoints that explicitly support it.
// Example: GET /api/v3/computers-inventory
// See: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
func (t *Transport) GetPaginated(ctx context.Context, path string, rsqlQuery map[string]string, headers map[string]string, mergePage func(pageData []byte) error) (*interfaces.Response, error) {
	currentParams := make(map[string]string)
	maps.Copy(currentParams, rsqlQuery)
	if currentParams["page"] == "" {
		currentParams["page"] = "0"
	}
	if currentParams["page-size"] == "" {
		currentParams["page-size"] = strconv.Itoa(DefaultPageSize)
	}

	var lastResp *interfaces.Response
	for {
		var pageResp jamfPaginatedPage
		resp, err := t.Get(ctx, path, currentParams, headers, &pageResp)
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
		// Next page if we got a full page and there might be more
		if len(pageResp.Results) == 0 || (pageNum+1)*pageSize >= pageResp.TotalCount {
			break
		}
		currentParams["page"] = strconv.Itoa(pageNum + 1)
	}
	return lastResp, nil
}

// HasNextPage checks if there is a next page (for link-based pagination).
func HasNextPage(links *PaginationLinks) bool {
	return links != nil && links.Next != ""
}

// ExtractParamsFromURL extracts query parameters from a URL string.
func ExtractParamsFromURL(urlStr string) (map[string]string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	params := make(map[string]string)
	for key, values := range parsedURL.Query() {
		if len(values) > 0 {
			params[key] = values[0]
		}
	}
	return params, nil
}
