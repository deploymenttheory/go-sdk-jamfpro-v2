# Go SDK for Jamf Pro API

[![Go Report Card](https://goreportcard.com/badge/github.com/deploymenttheory/go-sdk-jamfpro-v2)](https://goreportcard.com/report/github.com/deploymenttheory/go-sdk-jamfpro-v2)
[![GoDoc](https://pkg.go.dev/badge/github.com/deploymenttheory/go-sdk-jamfpro-v2)](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-jamfpro-v2)
[![License](https://img.shields.io/github/license/deploymenttheory/go-sdk-jamfpro-v2)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/deploymenttheory/go-sdk-jamfpro-v2)](https://go.dev/)
[![Release](https://img.shields.io/github/v/release/deploymenttheory/go-sdk-jamfpro-v2)](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/releases)
[![Tests](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/workflows/Tests/badge.svg)](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/actions)

A Go client library for the [Jamf Pro API](https://developer.jamf.com/jamf-pro/reference), supporting both the Classic API and the Jamf Pro API (REST). Uses OAuth2 or Basic auth with bearer token exchange, automatic token refresh, and production-ready transport (retries, sticky sessions, logging, optional OpenTelemetry tracing).


## Quick Start

Get started quickly with the SDK using the **[Quick Start Guide](docs/guides/quick-start.md)**, which includes:

- Installation instructions
- Your first API call
- Common operations (list, get, create, update, delete)
- Authentication from environment or config file
- Error handling and response metadata
- Links to configuration guides for production use

## Examples

The [examples directory](examples/) contains working examples for many SDK services:

- **Jamf Pro API:** [examples/jamf_pro_api/](examples/jamf_pro_api/) — API integrations, API roles, buildings, categories, computer groups, computer prestages, departments, dock items, enrollment settings, packages, reenrollment, SSO settings, volume purchasing, and more
- **Classic API:** [examples/classic_api/](examples/classic_api/) — Network segments, printers, restricted software, webhooks, and other Classic endpoints

Each example includes a complete `main.go` you can run with your Jamf Pro credentials.


## HTTP Client Configuration

The SDK includes a powerful HTTP client with production-ready configuration options:

- **[Authentication](docs/guides/authentication.md)** - OAuth2 and Basic auth with secure credential management
- **[Timeouts & Retries](docs/guides/timeouts-retries.md)** - Configurable timeouts and automatic retry logic with exponential backoff
- **[TLS/SSL Configuration](docs/guides/tls-configuration.md)** - Custom certificates, mutual TLS, and security settings
- **[Proxy Support](docs/guides/proxy.md)** - HTTP/HTTPS/SOCKS5 proxy configuration
- **[Custom Headers](docs/guides/custom-headers.md)** - Global and per-request header management
- **[Structured Logging](docs/guides/logging.md)** - Integration with zap for production logging
- **[OpenTelemetry Tracing](docs/guides/opentelemetry.md)** - Distributed tracing and observability
- **[Debug Mode](docs/guides/debugging.md)** - Detailed request/response inspection

## Configuration Options

### Creating a client

```go
import (
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

// From environment (INSTANCE_DOMAIN, AUTH_METHOD, CLIENT_ID, CLIENT_SECRET or BASIC_AUTH_*)
jamfClient, err := jamfpro.NewClientFromEnv()

// From AuthConfig (e.g. from file or secret manager)
authConfig := client.AuthConfigFromEnv() // or LoadAuthConfigFromFile(path)
jamfClient, err := jamfpro.NewClient(authConfig, client.WithLogger(logger))
```

### AuthConfig fields

```go
&client.AuthConfig{
    InstanceDomain:           "https://your-instance.jamfcloud.com",
    AuthMethod:               client.AuthMethodOAuth2, // or client.AuthMethodBasic
    ClientID:                 "your-client-id",
    ClientSecret:             "your-client-secret",
    TokenRefreshBufferPeriod: 5 * time.Minute,  // refresh before expiry
    HideSensitiveData:        true,              // redact tokens in logs
}
```

### Optional client options

The SDK client supports extensive configuration through functional options. Below is the complete list of available configuration options grouped by category.

#### Basic Configuration

```go
client.WithBaseURL("https://...")                    // Custom base URL
client.WithTimeout(30*time.Second)                   // Request timeout
client.WithRetryCount(3)                             // Number of retry attempts
client.WithRetryWaitTime(2*time.Second)              // Initial retry wait time
client.WithRetryMaxWaitTime(10*time.Second)          // Maximum retry wait time
client.WithTotalRetryDuration(2*time.Minute)         // Total retry budget
```

#### TLS/Security

```go
client.WithTLSClientConfig(tlsConfig)                // Custom TLS configuration
client.WithInsecureSkipVerify()                      // Skip cert verification (dev only!)
```

#### Network

```go
client.WithProxy("http://proxy:8080")                // HTTP/HTTPS/SOCKS5 proxy
client.WithTransport(customTransport)                // Custom HTTP transport
```

#### Headers

```go
client.WithUserAgent("MyApp/1.0")                    // Set User-Agent header
client.WithGlobalHeader("X-Custom-Header", "value")  // Add single global header
client.WithGlobalHeaders(map[string]string{...})     // Add multiple global headers
```

#### Observability

```go
client.WithLogger(zapLogger)                         // Structured logging with zap
jamfClient.EnableTracing(otelConfig)                 // OpenTelemetry distributed tracing (call after NewClient)
client.WithDebug()                                   // Enable debug mode (dev only!)
```

#### Concurrency & Rate Limiting

```go
client.WithMaxConcurrentRequests(5)                  // Limit concurrent requests (Jamf Pro recommendation: ≤5)
client.WithMandatoryRequestDelay(100*time.Millisecond) // Add delay between requests
```

#### Example: Production Configuration

```go
import (
    "time"
    "go.uber.org/zap"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

logger, _ := zap.NewProduction()
authConfig := client.AuthConfigFromEnv()

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTimeout(30*time.Second),
    client.WithRetryCount(3),
    client.WithLogger(logger),
    client.WithMaxConcurrentRequests(5),
    client.WithGlobalHeader("X-Application-Name", "MyJamfIntegration"),
)

// Enable OpenTelemetry tracing (optional)
jamfClient.EnableTracing(&client.OTelConfig{
    ServiceName: "my-jamf-integration",
})
```

See the [configuration guides](docs/guides/) for detailed documentation on each option.

## Documentation

- [Jamf Pro API Reference](https://developer.jamf.com/jamf-pro/reference)
- [GoDoc](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-jamfpro-v2)

## Contributing

Contributions are welcome. Please read our [Contributing Guidelines](CONTRIBUTING.md) before submitting pull requests.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Support

- **Issues:** [GitHub Issues](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/issues)
- **Jamf Pro API docs:** [developer.jamf.com](https://developer.jamf.com/jamf-pro/reference)

## Disclaimer

This is a community SDK and is not affiliated with or endorsed by Jamf LLC.
