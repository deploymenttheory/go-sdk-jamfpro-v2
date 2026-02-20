# Custom Headers

## What are Custom Headers?

Custom headers allow you to add additional HTTP headers to all requests (global headers) or specific requests. This is useful for adding metadata, tracking identifiers, or custom application context.

## Why Use Custom Headers?

Custom headers help you:

- **Track requests** - Add request IDs for debugging and tracing
- **Add metadata** - Include application version, user context, etc.
- **Compliance** - Include required headers for auditing
- **Integration** - Pass data to intermediate proxies or gateways
- **Correlation** - Link requests across distributed systems

## When to Configure Them

Add custom headers when:

- Need to correlate requests across systems
- Adding application metadata for monitoring
- Working with API gateways that require specific headers
- Meeting compliance requirements for request tracking
- Implementing distributed tracing without OpenTelemetry

## Basic Example

Here's how to add custom headers:

```go
package main

import (
    "context"
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
    authConfig := client.AuthConfigFromEnv()

    // Add a global header to all requests
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithGlobalHeader("X-Application-Name", "MyJamfIntegration"),
        client.WithGlobalHeader("X-Application-Version", "1.0.0"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // All requests now include these headers
    result, _, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found %d buildings", len(result.Results))
}
```

## Configuration Options

### Option 1: Single Global Header

Add one header that applies to all requests:

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Request-ID", requestID),
)
```

**When to use:** Adding a single tracking or metadata header

---

### Option 2: Multiple Global Headers

Add multiple headers at once:

```go
headers := map[string]string{
    "X-Application-Name":    "MyJamfApp",
    "X-Application-Version": "1.0.0",
    "X-Environment":         "production",
    "X-Region":              "us-east-1",
}

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeaders(headers),
)
```

**When to use:** Adding multiple metadata or tracking headers

---

### Option 3: Chain Multiple Headers

Add headers one at a time with multiple options:

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-App-Name", "MyJamfApp"),
    client.WithGlobalHeader("X-App-Version", "1.0.0"),
    client.WithGlobalHeader("X-User-ID", userID),
)
```

**When to use:** Building headers conditionally or from different sources

---

### Option 4: Dynamic Headers

Generate headers dynamically:

```go
import "github.com/google/uuid"

// Generate unique request ID for each client
requestID := uuid.New().String()

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Request-ID", requestID),
    client.WithGlobalHeader("X-Timestamp", time.Now().Format(time.RFC3339)),
)
```

**When to use:** Headers that change per client instance

---

## Common Use Cases

### Use Case 1: Request Tracking

```go
import "github.com/google/uuid"

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Request-ID", uuid.New().String()),
    client.WithGlobalHeader("X-Correlation-ID", correlationID),
)
```

### Use Case 2: Application Metadata

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeaders(map[string]string{
        "X-Application-Name":    "JamfDeviceManager",
        "X-Application-Version": version,
        "X-Build-Number":        buildNumber,
        "X-Environment":         env,
    }),
)
```

### Use Case 3: User Context

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeaders(map[string]string{
        "X-User-ID":      userID,
        "X-Organization": orgID,
        "X-Department":   department,
    }),
)
```

### Use Case 4: API Gateway Integration

```go
// Headers required by API gateway
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeaders(map[string]string{
        "X-Gateway-Key":  gatewayKey,
        "X-API-Version":  "v3",
        "X-Client-Type":  "sdk-go",
    }),
)
```

## Header Naming Conventions

### Standard Patterns

```go
// Use X- prefix for custom headers (traditional)
"X-Application-Name"
"X-Request-ID"
"X-User-Context"

// Or modern convention without X-
"Application-Name"
"Request-ID"
"User-Context"

// Use kebab-case for readability
"X-Application-Name"  // Good
"X-APPLICATION-NAME"  // Less readable
"X_Application_Name"  // Don't use underscores
```

### Reserved Headers

Some headers are automatically set by the SDK:

- `User-Agent` - Set via `WithUserAgent()`
- `Authorization` - Set automatically from AuthConfig
- `Content-Type` - Set automatically based on request body
- `Accept` - Set automatically for JSON responses

## Troubleshooting

### Headers Not Appearing in Requests

**Problem:** Custom headers don't appear in requests

**Solutions:**
```go
// Verify headers are set
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Test", "value"),
    client.WithDebug(), // Enable debug mode to see headers
)

// Check logs for header values
```

### Special Characters in Header Values

**Problem:** Header values with special characters cause errors

**Solution:** URL encode special characters:
```go
import "net/url"

encodedValue := url.QueryEscape("value with spaces")
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Custom", encodedValue),
)
```

## Security Considerations

✅ **Do:**
- Validate header values before setting
- Use headers for non-sensitive metadata
- Document custom headers in API integration guides
- Use standard header naming conventions
- Keep header values concise

❌ **Don't:**
- Put sensitive data in headers (passwords, tokens, PII)
- Use headers for large data payloads
- Include credentials unless encrypted in transit
- Log header values that contain secrets
- Use non-standard or ambiguous header names

## Testing with Custom Headers

```go
func TestCustomHeaders(t *testing.T) {
    authConfig := &client.AuthConfig{
        InstanceDomain: "https://test.jamfcloud.com",
        AuthMethod:     client.AuthMethodOAuth2,
        ClientID:       "test-id",
        ClientSecret:   "test-secret",
    }

    // Test single header
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithGlobalHeader("X-Test", "value"),
    )
    assert.NoError(t, err)

    // Test multiple headers
    headers := map[string]string{
        "X-Test-1": "value1",
        "X-Test-2": "value2",
    }
    jamfClient, err = jamfpro.NewClient(
        authConfig,
        client.WithGlobalHeaders(headers),
    )
    assert.NoError(t, err)
}
```

### Inspecting Headers in Tests

```go
// Use debug mode to see actual headers sent
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Test", "value"),
    client.WithDebug(),
)

// Or inspect via HTTP mock
```

## Examples by Use Case

### With OpenTelemetry

```go
import "go.opentelemetry.io/otel/trace"

// Add trace context to headers
spanCtx := trace.SpanContextFromContext(ctx)
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithGlobalHeader("X-Trace-ID", spanCtx.TraceID().String()),
    client.WithGlobalHeader("X-Span-ID", spanCtx.SpanID().String()),
)
```

### With Request ID Propagation

```go
// Propagate request ID from incoming HTTP request
func handleRequest(w http.ResponseWriter, r *http.Request) {
    requestID := r.Header.Get("X-Request-ID")
    if requestID == "" {
        requestID = uuid.New().String()
    }

    authConfig := client.AuthConfigFromEnv()
    jamfClient, _ := jamfpro.NewClient(
        authConfig,
        client.WithGlobalHeader("X-Request-ID", requestID),
    )

    // Use client...
}
```

## Related Documentation

- [Authentication](authentication.md) - Configure OAuth2 or Basic auth
- [Debugging](debugging.md) - View headers in debug output
- [Logging](logging.md) - Log header values (be careful with sensitive data)
- [OpenTelemetry](opentelemetry.md) - Integrate trace context in headers
