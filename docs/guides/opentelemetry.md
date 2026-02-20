# OpenTelemetry Tracing

## What is OpenTelemetry Tracing?

OpenTelemetry tracing provides distributed tracing capabilities for your Jamf Pro API calls. It automatically captures detailed information about each HTTP request, including timing, status codes, errors, and request/response metadata.

## Why Use OpenTelemetry?

OpenTelemetry tracing helps you:

- **Monitor performance** - Track how long API calls take and identify bottlenecks
- **Debug issues** - See the complete flow of requests across your application
- **Track errors** - Automatically capture and report API errors with full context
- **Improve observability** - Integrate with platforms like Jaeger, Zipkin, DataDog, Honeycomb, etc.
- **Understand dependencies** - Visualize how your application interacts with Jamf Pro

## When to Use It

Use OpenTelemetry tracing when:

- Running in production environments where observability is critical
- Debugging complex issues that span multiple services
- Monitoring API performance and identifying slow requests
- Tracking error rates and failure patterns
- Meeting compliance or SLA requirements for observability

## Basic Example

Here's a simple example showing how to enable OpenTelemetry tracing:

```go
package main

import (
    "context"
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
    // Step 1: Initialize OpenTelemetry exporter
    exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        log.Fatal(err)
    }

    // Step 2: Create tracer provider
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
    )
    defer tp.Shutdown(context.Background())

    // Step 3: Set as global tracer provider
    otel.SetTracerProvider(tp)

    // Step 4: Create Jamf Pro client
    authConfig := client.AuthConfigFromEnv()
    jamfClient, err := jamfpro.NewClient(authConfig)
    if err != nil {
        log.Fatal(err)
    }

    // Step 5: Enable tracing on the client
    if err := jamfClient.EnableTracing(nil); err != nil {
        log.Fatal(err)
    }

    // Step 6: Use the client normally - tracing happens automatically!
    result, _, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }

    log.Printf("Found %d buildings", len(result.Results))

    // Traces are automatically exported - check your console output!
}
```

**What you get:**

- All HTTP requests are automatically traced
- Spans include method, URL, status code, timing
- Errors are automatically recorded
- Zero code changes needed in your business logic

## Alternative Configuration Options

### Option 1: Using Default Configuration

The simplest approach uses the global OpenTelemetry tracer provider:

```go
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
if err != nil {
    log.Fatal(err)
}

// nil uses otel.GetTracerProvider() and otel.GetTextMapPropagator()
if err := jamfClient.EnableTracing(nil); err != nil {
    log.Fatal(err)
}
```

**When to use:** For most applications where you've already configured OpenTelemetry globally.

---

### Option 2: Custom Tracer Provider

Provide a specific tracer provider for more control:

```go
// Create your own tracer provider
myTracerProvider := trace.NewTracerProvider(
    trace.WithBatcher(myExporter),
    trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.1))), // Sample 10%
)

// Configure the client to use it
otelConfig := &client.OTelConfig{
    TracerProvider: myTracerProvider,
    ServiceName:    "my-jamf-integration",
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
if err != nil {
    log.Fatal(err)
}

if err := jamfClient.EnableTracing(otelConfig); err != nil {
    log.Fatal(err)
}
```

**When to use:** When you need different tracing configurations for different clients, or want to override the global tracer provider.

---

### Option 3: Custom Span Naming

Customize how spans are named for better organization in your tracing UI:

```go
otelConfig := &client.OTelConfig{
    SpanNameFormatter: func(operation string, req *http.Request) string {
        // Custom format: "Jamf Pro: GET /api/v1/buildings"
        return fmt.Sprintf("Jamf Pro: %s %s", req.Method, req.URL.Path)
    },
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
if err != nil {
    log.Fatal(err)
}

if err := jamfClient.EnableTracing(otelConfig); err != nil {
    log.Fatal(err)
}
```

**When to use:** When you want more descriptive span names in your tracing dashboard (e.g., Jaeger, Zipkin).

---

### Option 4: Custom Propagators

Control how trace context is propagated across service boundaries:

```go
import "go.opentelemetry.io/otel/propagation"

otelConfig := &client.OTelConfig{
    Propagators: propagation.NewCompositeTextMapPropagator(
        propagation.TraceContext{},
        propagation.Baggage{},
    ),
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
if err != nil {
    log.Fatal(err)
}

if err := jamfClient.EnableTracing(otelConfig); err != nil {
    log.Fatal(err)
}
```

**When to use:** When integrating with systems that use specific trace context formats (W3C Trace Context, B3, etc.).

---

## Integration with Popular Backends

### Jaeger

```go
import (
    "go.opentelemetry.io/otel/exporters/jaeger"
    "go.opentelemetry.io/otel/sdk/trace"
)

exporter, _ := jaeger.New(jaeger.WithCollectorEndpoint(
    jaeger.WithEndpoint("http://jaeger:14268/api/traces"),
))

tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
otel.SetTracerProvider(tp)

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(authConfig)
jamfClient.EnableTracing(nil)
```

### OTLP (OpenTelemetry Protocol)

```go
import "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

exporter, _ := otlptracegrpc.New(context.Background(),
    otlptracegrpc.WithEndpoint("otel-collector:4317"),
    otlptracegrpc.WithInsecure(),
)

tp := trace.NewTracerProvider(trace.WithBatcher(exporter))
otel.SetTracerProvider(tp)

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(authConfig)
jamfClient.EnableTracing(nil)
```

## What Gets Traced

Each HTTP request creates a span with the following information:

| Attribute | Description | Example |
|-----------|-------------|---------|
| `http.method` | HTTP method | `GET`, `POST`, `PUT`, `DELETE` |
| `http.url` | Full URL | `https://instance.jamfcloud.com/api/v1/buildings` |
| `http.status_code` | Response status | `200`, `404`, `401` |
| `http.request_content_length` | Request size in bytes | `1024` |
| `http.response_content_length` | Response size in bytes | `4096` |
| Span duration | Request timing | `245ms` |
| Span status | Success or error | `Ok`, `Error` |

All attributes follow [OpenTelemetry semantic conventions](https://opentelemetry.io/docs/specs/semconv/http/) for HTTP clients.

## Disabling Tracing

To disable tracing, simply don't call `EnableTracing()`:

```go
// No tracing - client works normally without instrumentation
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
```

## Performance Considerations

- **Minimal overhead**: OpenTelemetry adds microseconds of latency per request
- **Async export**: Spans are batched and exported in the background
- **Sampling**: Use sampling to reduce overhead in high-traffic scenarios:
  ```go
  trace.WithSampler(trace.TraceIDRatioBased(0.1)) // Sample 10%
  ```
- **No-op when disabled**: Without `EnableTracing()`, there's zero tracing overhead

## Complete Production Example

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
    "go.opentelemetry.io/otel/sdk/trace"
    "go.uber.org/zap"
)

func main() {
    // Initialize logger
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    // Initialize OpenTelemetry
    exporter, _ := otlptracegrpc.New(context.Background(),
        otlptracegrpc.WithEndpoint("otel-collector:4317"),
        otlptracegrpc.WithInsecure(),
    )
    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithSampler(trace.ParentBased(trace.TraceIDRatioBased(0.1))),
    )
    defer tp.Shutdown(context.Background())
    otel.SetTracerProvider(tp)

    // Create Jamf Pro client with logging and tracing
    authConfig := client.AuthConfigFromEnv()
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithLogger(logger),
        client.WithTimeout(30*time.Second),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Enable tracing
    if err := jamfClient.EnableTracing(&client.OTelConfig{
        ServiceName: "jamf-integration",
    }); err != nil {
        log.Fatal(err)
    }

    // Use client normally
    ctx := context.Background()
    result, _, err := jamfClient.Buildings.ListV1(ctx, nil)
    if err != nil {
        logger.Error("Failed to list buildings", zap.Error(err))
        return
    }

    logger.Info("Buildings retrieved",
        zap.Int("count", len(result.Results)),
    )
}
```

## Related Documentation

- [Structured Logging](logging.md) - Combine tracing with logging for complete observability
- [Debugging](debugging.md) - Debug mode for detailed request inspection
- [Authentication](authentication.md) - Configure API access
- [OpenTelemetry Go Documentation](https://opentelemetry.io/docs/languages/go/)
