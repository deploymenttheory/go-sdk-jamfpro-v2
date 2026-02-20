# Debug Mode

## What is Debug Mode?

Debug mode enables detailed logging of all HTTP requests and responses, including headers, bodies, and timing information. This helps troubleshoot issues and understand exactly what the SDK is sending to the Jamf Pro API.

## Why Use Debug Mode?

Debug mode helps you:

- **Troubleshoot issues** - See exactly what's being sent and received
- **Verify requests** - Confirm API calls are formatted correctly
- **Inspect responses** - View raw API responses for debugging
- **Monitor traffic** - Understand request/response patterns
- **Learn the API** - See how the SDK interacts with Jamf Pro

## When to Enable It

Enable debug mode when:

- Debugging integration issues
- Investigating unexpected API responses
- Troubleshooting authentication problems
- Verifying request formats
- Learning how the SDK works
- **Only in development** - Never enable in production!

## Basic Example

Here's how to enable debug mode:

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

    // Enable debug mode
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithDebug(),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Make a request - detailed output will be printed
    result, _, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found %d buildings", len(result.Results))
}
```

**Debug Output:**
```
2024-01-15 10:30:45 | GET | https://your-instance.jamfcloud.com/api/v1/buildings?page=0&page-size=100
REQUEST HEADERS:
  Accept-Encoding: gzip
  User-Agent: go-sdk-jamfpro-v2/2.0.0
  Authorization: Bearer ***redacted***
  Content-Type: application/json

RESPONSE:
  Status Code: 200
  Proto: HTTP/2.0
  Duration: 245ms
RESPONSE HEADERS:
  Content-Type: application/json

RESPONSE BODY:
{
  "totalCount": 2,
  "results": [
    {
      "id": "1",
      "name": "Main Office"
    }
  ]
}
```

## Configuration Options

### Option 1: Basic Debug Mode

Enable standard debug output:

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithDebug(),
)
```

**When to use:** General debugging and troubleshooting

**Output includes:**
- Request method and URL
- Request headers (bearer token redacted if HideSensitiveData is true)
- Response status and headers
- Response body
- Request duration

---

### Option 2: Debug Mode with Custom Logger

Combine debug mode with structured logging:

```go
import "go.uber.org/zap"

logger, _ := zap.NewDevelopment()

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
    client.WithDebug(),
)
```

**When to use:** Structured debug output for parsing or analysis

---

### Option 3: Conditional Debug Mode

Enable debug mode based on environment:

```go
var options []client.ClientOption

if os.Getenv("DEBUG") == "true" {
    options = append(options, client.WithDebug())
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig, options...)
```

**When to use:** Toggle debug mode without code changes

```bash
# Enable debug mode
DEBUG=true go run main.go

# Disable debug mode
go run main.go
```

---

## What Gets Logged

### Request Information
```
GET https://your-instance.jamfcloud.com/api/v1/buildings
REQUEST HEADERS:
  User-Agent: go-sdk-jamfpro-v2/2.0.0
  Authorization: Bearer ***redacted***
  Content-Type: application/json
  X-Custom-Header: value
```

### Response Information
```
RESPONSE:
  Status Code: 200 OK
  Proto: HTTP/2.0
  Duration: 234ms

RESPONSE HEADERS:
  Content-Type: application/json

RESPONSE BODY:
{ ... full JSON response ... }
```

### Error Responses
```
RESPONSE:
  Status Code: 404 Not Found
  Duration: 123ms

RESPONSE BODY:
{
  "httpStatus": 404,
  "errors": [
    {
      "code": "NOT_FOUND",
      "description": "Building not found"
    }
  ]
}
```

## Common Debugging Scenarios

### Scenario 1: Authentication Issues

```go
// Enable debug to see authentication process
authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithDebug(),
)

// Check if bearer token is being obtained and used correctly
_, _, err := jamfClient.Buildings.ListV1(ctx, nil)
// Look for "Authorization" header in debug output
```

### Scenario 2: Request Format Verification

```go
// Verify POST request body format
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithDebug(),
)

// Debug shows actual JSON being sent
import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"

_, _, err := jamfClient.Buildings.CreateV1(ctx, &buildings.RequestBuilding{
    Name: "New Building",
})
```

### Scenario 3: Proxy Issues

```go
// Debug proxy connections
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithProxy("http://proxy:8080"),
    client.WithDebug(),
)

// See if requests are going through proxy
_, _, err := jamfClient.Buildings.ListV1(ctx, nil)
```

### Scenario 4: TLS Certificate Issues

```go
// Debug TLS handshake
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithDebug(),
)

// See TLS-related errors in debug output
_, _, err := jamfClient.Buildings.ListV1(ctx, nil)
```

### Scenario 5: Pagination Issues

```go
// Debug pagination parameters
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithDebug(),
)

// See query parameters in URL
result, _, _ := jamfClient.Buildings.ListV1(ctx, map[string]string{
    "page":      "0",
    "page-size": "50",
})
```

## Security Warnings

⚠️ **NEVER enable debug mode in production!**

Debug mode logs sensitive information:
- **Bearer tokens** (partially redacted if HideSensitiveData is true)
- **Request/response bodies** (may contain sensitive data)
- **Headers** (may contain credentials or tokens)
- **URLs** (may contain parameters)

### Safe Debug Practices

✅ **Do:**
- Use only in development/testing
- Clear debug logs before committing
- Use environment variables to toggle debug
- Redact sensitive data from debug logs (set HideSensitiveData: true)
- Limit debug output to necessary information

❌ **Don't:**
- Enable in production
- Commit debug output to version control
- Share debug logs containing secrets
- Log to public systems with debug enabled
- Leave debug mode on continuously

## Disabling Debug Mode

```go
// Simply omit WithDebug() option
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)

// Or conditionally disable
var options []client.ClientOption
if os.Getenv("ENVIRONMENT") != "production" {
    options = append(options, client.WithDebug())
}
jamfClient, err := jamfpro.NewClient(authConfig, options...)
```

## Alternative Debugging Tools

### HTTP Proxies

Use HTTP debugging proxies for advanced inspection:

```bash
# Charles Proxy, mitmproxy, Burp Suite, etc.
mitmproxy -p 8080
```

```go
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithProxy("http://127.0.0.1:8080"),
    client.WithInsecureSkipVerify(), // For proxy SSL inspection
)
```

### Network Monitoring

Use system tools to monitor HTTP traffic:

```bash
# tcpdump
sudo tcpdump -i any -A 'host your-instance.jamfcloud.com'

# Wireshark
# Use GUI to filter: http.host == "your-instance.jamfcloud.com"
```

### Structured Logging

Use structured logging instead of debug mode for production:

```go
import "go.uber.org/zap"

logger, _ := zap.NewProduction()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)

// Log specific operations
logger.Info("Making API call",
    zap.String("endpoint", "/api/v1/buildings"),
    zap.String("method", "GET"),
)
```

## Testing with Debug Mode

```go
func TestWithDebug(t *testing.T) {
    // Enable debug for specific test
    authConfig := &client.AuthConfig{
        InstanceDomain: "https://test.jamfcloud.com",
        AuthMethod:     client.AuthMethodOAuth2,
        ClientID:       "test-id",
        ClientSecret:   "test-secret",
    }

    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithDebug(),
    )
    require.NoError(t, err)

    // Debug output helps verify test behavior
    // ...
}
```

### Capturing Debug Output

```go
import (
    "bytes"
    "log"
)

func TestDebugOutput(t *testing.T) {
    // Capture debug output
    var buf bytes.Buffer
    log.SetOutput(&buf)
    defer log.SetOutput(os.Stderr)

    authConfig := &client.AuthConfig{
        InstanceDomain: "https://test.jamfcloud.com",
        AuthMethod:     client.AuthMethodOAuth2,
        ClientID:       "test-id",
        ClientSecret:   "test-secret",
    }

    jamfClient, _ := jamfpro.NewClient(
        authConfig,
        client.WithDebug(),
    )

    // Make request...

    // Verify debug output
    output := buf.String()
    assert.Contains(t, output, "REQUEST HEADERS")
    assert.Contains(t, output, "RESPONSE")
}
```

## Related Documentation

- [Logging](logging.md) - Structured logging for production
- [OpenTelemetry](opentelemetry.md) - Distributed tracing for observability
- [Authentication](authentication.md) - Debug authentication issues
- [Proxy Support](proxy.md) - Debug proxy connections
- [TLS Configuration](tls-configuration.md) - Debug TLS issues
