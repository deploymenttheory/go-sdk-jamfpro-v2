package client

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// OTelConfig holds OpenTelemetry configuration options.
type OTelConfig struct {
	// TracerProvider is the OpenTelemetry tracer provider to use.
	// If nil, the global tracer provider will be used.
	TracerProvider trace.TracerProvider

	// Propagators is the propagator to use for context propagation.
	// If nil, the global propagator will be used.
	Propagators propagation.TextMapPropagator

	// ServiceName is the name of the service for tracing spans.
	// Defaults to "jamfpro-client".
	ServiceName string

	// SpanNameFormatter allows customizing span names.
	// If nil, defaults to "HTTP {method}" format.
	SpanNameFormatter func(operation string, req *http.Request) string
}

// DefaultOTelConfig returns a default OpenTelemetry configuration.
func DefaultOTelConfig() *OTelConfig {
	return &OTelConfig{
		TracerProvider: otel.GetTracerProvider(),
		Propagators:    otel.GetTextMapPropagator(),
		ServiceName:    "jamfpro-client",
	}
}

// EnableTracing wraps the HTTP client transport with OpenTelemetry instrumentation.
// This provides automatic tracing for all HTTP requests made by the client.
//
// The instrumentation captures:
// - HTTP method, URL, status code
// - Request and response headers (configurable)
// - Error details
// - Request/response timing
//
// All spans follow OpenTelemetry semantic conventions for HTTP clients.
func (t *Transport) EnableTracing(config *OTelConfig) error {
	if config == nil {
		config = DefaultOTelConfig()
	}

	httpClient := t.client.Client()
	if httpClient == nil {
		return nil
	}

	transport := httpClient.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	opts := []otelhttp.Option{
		otelhttp.WithTracerProvider(config.TracerProvider),
		otelhttp.WithPropagators(config.Propagators),
	}
	if config.SpanNameFormatter != nil {
		opts = append(opts, otelhttp.WithSpanNameFormatter(config.SpanNameFormatter))
	}

	instrumentedTransport := otelhttp.NewTransport(transport, opts...)
	httpClient.Transport = instrumentedTransport

	t.logger.Info("OpenTelemetry tracing enabled",
		zap.String("service_name", config.ServiceName))

	return nil
}
