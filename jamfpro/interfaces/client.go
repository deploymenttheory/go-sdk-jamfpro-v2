package interfaces

import (
	"context"
	"io"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Response represents HTTP response metadata returned alongside data or errors.
type Response struct {
	StatusCode int
	Status     string
	Headers    http.Header
	Body       []byte
	Duration   time.Duration
	ReceivedAt time.Time
	Size       int64
}

// MultipartProgressCallback is called during multipart uploads to report progress.
type MultipartProgressCallback func(fieldName string, fileName string, bytesWritten int64, totalBytes int64)

// RSQLFilterBuilder builds RSQL filter expressions for Jamf Pro API endpoints
// that support the filter query parameter.
//
// RSQL reference: https://developer.jamf.com/jamf-pro/docs/filtering-with-rsql
//
// Usage:
//
//	filter := client.RSQLBuilder().
//	    EqualTo("general.name", "MacBook Pro").
//	    And().
//	    GreaterThan("hardware.totalRamMegabytes", "8192").
//	    Build()
//	// → general.name=="MacBook Pro";hardware.totalRamMegabytes>"8192"
type RSQLFilterBuilder interface {
	// EqualTo produces: field=="value". Wildcards (*) in value are preserved.
	EqualTo(field, value string) RSQLFilterBuilder
	// NotEqualTo produces: field!="value". Wildcards (*) in value are preserved.
	NotEqualTo(field, value string) RSQLFilterBuilder
	// LessThan produces: field<"value".
	LessThan(field, value string) RSQLFilterBuilder
	// LessOrEqual produces: field<="value".
	LessOrEqual(field, value string) RSQLFilterBuilder
	// GreaterThan produces: field>"value".
	GreaterThan(field, value string) RSQLFilterBuilder
	// GreaterOrEqual produces: field>="value".
	GreaterOrEqual(field, value string) RSQLFilterBuilder

	// In produces: field=in=(v1,v2,...).
	In(field string, values ...string) RSQLFilterBuilder
	// NotIn produces: field=out=(v1,v2,...).
	NotIn(field string, values ...string) RSQLFilterBuilder

	// Contains produces: field=="*value*" (substring match, literal * in value is escaped).
	Contains(field, value string) RSQLFilterBuilder
	// StartsWith produces: field=="value*" (prefix match, literal * in value is escaped).
	StartsWith(field, value string) RSQLFilterBuilder
	// EndsWith produces: field=="*value" (suffix match, literal * in value is escaped).
	EndsWith(field, value string) RSQLFilterBuilder

	// And appends a semicolon — logical AND in RSQL.
	And() RSQLFilterBuilder
	// Or appends a comma — logical OR in RSQL.
	Or() RSQLFilterBuilder

	// OpenGroup appends a left parenthesis for grouping sub-expressions.
	OpenGroup() RSQLFilterBuilder
	// CloseGroup appends a right parenthesis.
	CloseGroup() RSQLFilterBuilder

	// Build returns the completed RSQL filter string ready to use as rsqlQuery["filter"].
	Build() string
	// IsEmpty reports whether no expressions have been added yet.
	IsEmpty() bool
}

// HTTPClient is the interface service implementations depend on.
// The Transport struct in the client package satisfies this interface.
type HTTPClient interface {
	// Get executes a GET request and unmarshals the JSON response into the result parameter.
	// rsqlQuery is optional; pass nil or an empty map when no filtering or pagination is needed.
	// Supported keys: "filter" (RSQL expression), "sort", "page", "page-size".
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	Get(
		ctx context.Context, // request context
		path string, // API endpoint path
		rsqlQuery map[string]string, // optional: RSQL filter, sort, page, page-size
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// Post executes a POST request with a JSON body.
	// The body is marshaled to JSON and the response is unmarshaled into the result parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	Post(
		ctx context.Context, // request context
		path string, // API endpoint path
		body any, // request body to marshal as JSON
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// PostWithQuery executes a POST request with both query parameters and a JSON body.
	// The body is marshaled to JSON and the response is unmarshaled into the result parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	PostWithQuery(
		ctx context.Context, // request context
		path string, // API endpoint path
		rsqlQuery map[string]string, // URL query parameters
		body any, // request body to marshal as JSON
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// PostForm executes a POST request with form-urlencoded data.
	// The Content-Type header is automatically set to application/x-www-form-urlencoded.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	PostForm(
		ctx context.Context, // request context
		path string, // API endpoint path
		formData map[string]string, // form fields as key-value pairs
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// PostMultipart executes a POST request with multipart/form-data encoding, typically for file uploads.
	// The Content-Type header is automatically set to multipart/form-data with a boundary.
	// Progress tracking is supported via the optional progressCallback parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	PostMultipart(
		ctx context.Context, // request context
		path string, // API endpoint path
		fileField string, // form field name for the file
		fileName string, // name of the file being uploaded
		fileReader io.Reader, // reader for file content
		fileSize int64, // size of the file in bytes
		formFields map[string]string, // additional form fields
		headers map[string]string, // HTTP headers
		progressCallback MultipartProgressCallback, // optional progress callback
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// Put executes a PUT request with a JSON body.
	// The body is marshaled to JSON and the response is unmarshaled into the result parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	Put(
		ctx context.Context, // request context
		path string, // API endpoint path
		body any, // request body to marshal as JSON
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// Patch executes a PATCH request with a JSON body.
	// The body is marshaled to JSON and the response is unmarshaled into the result parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	Patch(
		ctx context.Context, // request context
		path string, // API endpoint path
		body any, // request body to marshal as JSON
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// Delete executes a DELETE request.
	// Query parameters and headers are applied if provided.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	Delete(
		ctx context.Context, // request context
		path string, // API endpoint path
		rsqlQuery map[string]string, // URL query parameters
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// DeleteWithBody executes a DELETE request with a JSON body (for bulk operations).
	// The body is marshaled to JSON and the response is unmarshaled into the result parameter.
	// Returns response metadata and error. Response is non-nil even on error for accessing headers.
	DeleteWithBody(
		ctx context.Context, // request context
		path string, // API endpoint path
		body any, // request body to marshal as JSON
		headers map[string]string, // HTTP headers
		result any, // pointer to unmarshal response into
	) (*Response, error)

	// GetBytes performs a GET request and returns raw bytes without unmarshaling.
	// Use this for non-JSON responses such as binary files, CSV exports, or raw XML.
	// rsqlQuery is optional; pass nil or an empty map when no filtering is needed.
	GetBytes(
		ctx context.Context, // request context
		path string, // API endpoint path
		rsqlQuery map[string]string, // optional: RSQL filter, sort, page, page-size
		headers map[string]string, // HTTP headers
	) (*Response, []byte, error)

	// GetPaginated transparently fetches all pages of a paginated Jamf Pro API endpoint,
	// calling mergePage with each page's results array. The caller supplies optional
	// "filter" (RSQL) and "sort" params in rsqlQuery; "page" and "page-size" are
	// managed internally. Only available on endpoints that explicitly support pagination.
	// Example: GET /api/v3/computers-inventory
	// See: https://developer.jamf.com/jamf-pro/reference/get_v3-computers-inventory
	GetPaginated(
		ctx context.Context, // request context
		path string, // API endpoint path
		rsqlQuery map[string]string, // optional: filter (RSQL), sort, page, page-size
		headers map[string]string, // HTTP headers
		mergePage func(pageData []byte) error, // callback to merge each page's results
	) (*Response, error)

	// RSQLBuilder returns a new RSQL filter expression builder.
	// Pass the Build() result as rsqlQuery["filter"] to filter paginated results.
	// See: https://developer.jamf.com/jamf-pro/docs/filtering-with-rsql
	RSQLBuilder() RSQLFilterBuilder

	// InvalidateToken explicitly revokes the current bearer token at the Jamf Pro API
	// and clears the local cache. Use before shutdown or credential rotation.
	// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
	InvalidateToken() error

	// KeepAliveToken extends the lifetime of the current bearer token without
	// performing a full re-authentication. Use before long-running operations
	// to prevent mid-operation token expiry.
	// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
	KeepAliveToken() error

	// GetLogger returns the configured zap logger instance.
	GetLogger() *zap.Logger
}
