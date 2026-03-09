# Client Build Options Examples

This directory contains comprehensive examples demonstrating various ways to configure and build the Jamf Pro API client.

## Basic Configuration

### 01. Load Configuration from JSON File
**File**: `01_basic_from_file/main.go`

Loads authentication configuration from a JSON file containing instance domain, auth method, and credentials.

```go
authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
jamfClient, err := jamfpro.NewClient(authConfig)
```

### 02. Load Configuration from Environment Variables
**File**: `02_basic_from_env/main.go`

Builds authentication configuration from environment variables:
- `JAMFPRO_INSTANCE_DOMAIN`
- `JAMFPRO_AUTH_METHOD` (oauth2 or basic)
- `JAMFPRO_CLIENT_ID` / `JAMFPRO_CLIENT_SECRET` (for OAuth2)
- `JAMFPRO_USERNAME` / `JAMFPRO_PASSWORD` (for Basic Auth)

```go
authConfig := jamfpro.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(authConfig)
```

### 03. OAuth2 Authentication (Manual)
**File**: `03_oauth2_manual/main.go`

Manually construct OAuth2 authentication configuration.

```go
authConfig := &jamfpro.AuthConfig{
    InstanceDomain: "https://your-instance.jamfcloud.com",
    AuthMethod:     jamfpro.AuthMethodOAuth2,
    ClientID:       "your-client-id",
    ClientSecret:   "your-client-secret",
}
```

### 04. Basic Authentication (Manual)
**File**: `04_basic_auth_manual/main.go`

Manually construct Basic authentication configuration.

```go
authConfig := &jamfpro.AuthConfig{
    InstanceDomain: "https://your-instance.jamfcloud.com",
    AuthMethod:     jamfpro.AuthMethodBasic,
    Username:       "your-username",
    Password:       "your-password",
}
```

## Advanced Configuration

### 05. Custom Client Options
**File**: `05_with_custom_options/main.go`

Configure timeout, retries, concurrency limits, request delays, and debug mode.

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    jamfpro.WithTimeout(60*time.Second),
    jamfpro.WithRetryCount(5),
    jamfpro.WithMaxConcurrentRequests(3),
    jamfpro.WithMandatoryRequestDelay(500*time.Millisecond),
    jamfpro.WithDebug(),
)
```

### 06. OpenTelemetry Tracing
**File**: `06_with_opentelemetry/main.go`

OpenTelemetry instrumentation is **always enabled** in the Jamf Pro client. Simply configure global OTel providers before creating the client, and the HTTP transport will automatically capture traces and metrics.

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Set up global tracer provider
exporter, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())
tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
otel.SetTracerProvider(tp)

// Client automatically uses the global provider
jamfClient, err := jamfpro.NewClient(authConfig)
```

If no global providers are configured, the instrumentation is a zero-overhead no-op.

### 07. Custom Logger
**File**: `07_with_custom_logger/main.go`

Use a custom zap logger for structured logging.

```go
logger, err := zap.NewDevelopment()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    jamfpro.WithLogger(logger),
)
```

### 08. RSQL Filtering (Basic)
**File**: `08_rsql_filtering/main.go`

Use the RSQL builder to filter API results.

```go
filter := jamfClient.
    GetTransport().
    RSQLBuilder().
    EqualTo("general.name", "MacBook*").
    And().
    GreaterThan("hardware.totalRamMegabytes", "8192").
    Build()

rsqlQuery := map[string]string{
    "filter": filter,
}
```

### 09. Bulk Operations with Throttling
**File**: `09_bulk_operations_with_throttling/main.go`

Configure concurrency limits and request delays for bulk operations to avoid rate limits.

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    jamfpro.WithMaxConcurrentRequests(3),
    jamfpro.WithMandatoryRequestDelay(1*time.Second),
    jamfpro.WithTotalRetryDuration(5*time.Minute),
)
```

### 10. HTTP Proxy
**File**: `10_with_proxy/main.go`

Configure HTTP proxy for all requests.

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    jamfpro.WithProxy("http://proxy.example.com:8080"),
)
```

### 11. Insecure TLS (Testing Only)
**File**: `11_insecure_skip_verify/main.go`

Disable TLS certificate verification for testing environments.

```go
jamfClient, err := jamfpro.NewClient(
    authConfig,
    jamfpro.WithInsecureSkipVerify(),
)
```

### 12. Advanced RSQL Filtering
**File**: `12_rsql_advanced_filtering/main.go`

Complex RSQL filters with grouping, OR logic, and multiple conditions.

```go
filter := jamfClient.
    GetTransport().
    RSQLBuilder().
    OpenGroup().
    EqualTo("hardware.make", "Apple").
    And().
    GreaterThan("hardware.totalRamMegabytes", "16384").
    CloseGroup().
    Or().
    OpenGroup().
    EqualTo("operatingSystem.name", "macOS").
    And().
    Contains("general.name", "MacBook").
    CloseGroup().
    Build()
```

### 13. Token Management
**File**: `13_token_management/main.go`

Manually manage bearer token lifecycle with keep-alive and invalidation.

```go
jamfClient.GetTransport().KeepAliveToken()

jamfClient.GetTransport().InvalidateToken()
```

### 14. Comprehensive Example
**File**: `14_comprehensive_example/main.go`

Combines multiple features: custom logger, timeouts, retries, concurrency limits, RSQL filtering, and token management.

## Available Client Options

All options are available via the `jamfpro` package:

### Basic Configuration
- `jamfpro.WithBaseURL(baseURL string)` - Override base URL
- `jamfpro.WithTimeout(timeout time.Duration)` - Set request timeout
- `jamfpro.WithUserAgent(userAgent string)` - Set custom user agent

### Retry Configuration
- `jamfpro.WithRetryCount(count int)` - Configure retry attempts
- `jamfpro.WithRetryWaitTime(waitTime time.Duration)` - Set retry wait time
- `jamfpro.WithRetryMaxWaitTime(maxWaitTime time.Duration)` - Set max retry wait
- `jamfpro.WithTotalRetryDuration(d time.Duration)` - Max retry duration

### Throttling & Concurrency
- `jamfpro.WithMaxConcurrentRequests(n int)` - Limit concurrent requests
- `jamfpro.WithMandatoryRequestDelay(d time.Duration)` - Fixed delay between requests

### Observability
- `jamfpro.WithLogger(logger *zap.Logger)` - Use custom logger
- `jamfpro.WithDebug()` - Enable debug logging

**Note**: OpenTelemetry instrumentation is always enabled. Configure global OTel providers (via `otel.SetTracerProvider()`, etc.) before creating the client, and the HTTP transport will automatically capture traces and metrics. If no global providers are configured, the instrumentation is a zero-overhead no-op.

### Headers & Transport
- `jamfpro.WithGlobalHeader(key, value string)` - Add global header
- `jamfpro.WithGlobalHeaders(headers map[string]string)` - Add multiple headers
- `jamfpro.WithProxy(proxyURL string)` - Configure HTTP proxy
- `jamfpro.WithTLSClientConfig(tlsConfig *tls.Config)` - Custom TLS config
- `jamfpro.WithTransport(transport http.RoundTripper)` - Custom HTTP transport
- `jamfpro.WithInsecureSkipVerify()` - Disable TLS verification (testing only)

## RSQL Reference

RSQL (RESTful Service Query Language) is used to filter API results.

### Operators
- `==` - Equal to
- `!=` - Not equal to
- `<` - Less than
- `<=` - Less than or equal to
- `>` - Greater than
- `>=` - Greater than or equal to
- `=in=` - In list
- `=out=` - Not in list

### Logical Operators
- `;` - AND
- `,` - OR
- `()` - Grouping

### Wildcards
- `*` - Matches any characters

### Examples
```
general.name=="MacBook*"
hardware.totalRamMegabytes>8192
general.name=="MacBook*";hardware.totalRamMegabytes>8192
(general.name=="MacBook*",general.name=="iMac*");operatingSystem.name=="macOS"
```

## References

- [Jamf Pro API Documentation](https://developer.jamf.com/jamf-pro/docs)
- [RSQL Filtering Guide](https://developer.jamf.com/jamf-pro/docs/filtering-with-rsql)
- [Jamf Pro API Scalability Best Practices](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-scalability-best-practices)
