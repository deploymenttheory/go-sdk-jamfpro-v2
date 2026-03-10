package client

import (
	"context"

	"go.uber.org/zap"
)

// Client is the interface service implementations depend on.
// The Transport struct in this package satisfies this interface.
type Client interface {
	// NewRequest returns a RequestBuilder that the service layer uses to
	// construct a complete request — headers, body, query params, result
	// target — before executing it via Get/Post/Put/Patch/Delete/GetBytes/GetPaginated.
	// Auth, retry, throttling, and concurrency limiting are applied by
	// the transport at execution time.
	NewRequest(ctx context.Context) *RequestBuilder

	// RSQLBuilder returns a new RSQL filter expression builder.
	// Pass the Build() result as SetQueryParam("filter", ...) to filter paginated results.
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
//	    → general.name=="MacBook Pro";hardware.totalRamMegabytes>"8192"
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

	// Build returns the completed RSQL filter string ready to use as SetQueryParam("filter", ...).
	Build() string
	// IsEmpty reports whether no expressions have been added yet.
	IsEmpty() bool
}
