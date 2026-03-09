package client

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

// applyOpenTelemetry wraps the HTTP transport with OpenTelemetry instrumentation.
// This is always enabled and uses the global OpenTelemetry providers set via:
//   - otel.SetTracerProvider()
//   - otel.SetMeterProvider()
//   - otel.SetTextMapPropagator()
//
// If no global providers are configured, the instrumentation is a no-op.
//
// The instrumentation automatically captures:
// - HTTP method, URL, status code
// - Request and response headers
// - Error details
// - Request/response timing
// - Metrics (request duration, body size, etc.)
//
// All telemetry follows OpenTelemetry semantic conventions for HTTP clients.
// See: https://opentelemetry.io/docs/languages/go/getting-started/
func (t *Transport) applyOpenTelemetry() {
	httpClient := t.client.Client()
	if httpClient == nil {
		return
	}

	transport := httpClient.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	instrumentedTransport := otelhttp.NewTransport(transport)
	httpClient.Transport = instrumentedTransport

	t.logger.Debug("OpenTelemetry HTTP instrumentation enabled (uses global providers)")
}
