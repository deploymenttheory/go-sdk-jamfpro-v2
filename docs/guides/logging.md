# Structured Logging

## What is Structured Logging?

The Jamf Pro SDK uses [zap](https://github.com/uber-go/zap) for high-performance structured logging. Structured logs use key-value pairs instead of formatted strings, making logs easier to parse, search, and analyze.

## Why Use Structured Logging?

Structured logging helps you:

- **Debug issues faster** - Search logs by specific fields (e.g., all 404 errors)
- **Monitor in production** - Send logs to aggregation systems (Splunk, ELK, Datadog)
- **Track performance** - Log request durations, response sizes, status codes
- **Audit API usage** - Track which endpoints are called and when
- **Correlate with tracing** - Combine with OpenTelemetry for complete observability

## When to Use It

Configure logging when:

- Running in production environments
- Debugging issues in development
- Monitoring API usage and performance
- Meeting compliance or audit requirements
- Integrating with log aggregation systems

## Basic Example

```go
package main

import (
    "context"
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
    "go.uber.org/zap"
)

func main() {
    // Step 1: Create a zap logger
    logger, err := zap.NewProduction()
    if err != nil {
        log.Fatal(err)
    }
    defer logger.Sync()

    // Step 2: Create client with custom logger
    authConfig := client.AuthConfigFromEnv()
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithLogger(logger),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Step 3: Use the client - logging happens automatically
    result, resp, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        logger.Error("Failed to list buildings",
            zap.Error(err),
            zap.Int("status_code", resp.StatusCode),
        )
        return
    }

    logger.Info("Successfully retrieved buildings",
        zap.Int("building_count", len(result.Results)),
    )
}
```

**Output (JSON):**
```json
{
  "level": "info",
  "ts": 1704891234.567,
  "caller": "client/transport.go:110",
  "msg": "Jamf Pro client created",
  "instance_domain": "https://your-instance.jamfcloud.com"
}
```

## Alternative Configuration Options

### Option 1: Production Logger

Use zap's production configuration for JSON logs:

```go
logger, _ := zap.NewProduction()

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)
```

**When to use:** Production environments, log aggregation systems

**Output:** JSON-formatted logs optimized for parsing

---

### Option 2: Development Logger

Use zap's development configuration for human-readable logs:

```go
logger, _ := zap.NewDevelopment()

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)
```

**When to use:** Local development, debugging

**Output:** Pretty-printed, colorized logs with stack traces

---

### Option 3: Custom Logger Configuration

Build a custom logger with specific settings:

```go
import "go.uber.org/zap/zapcore"

// Custom configuration
config := zap.Config{
    Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
    Development: false,
    Encoding:    "json",
    EncoderConfig: zapcore.EncoderConfig{
        TimeKey:        "timestamp",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "message",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    },
    OutputPaths:      []string{"stdout"},
    ErrorOutputPaths: []string{"stderr"},
}

logger, _ := config.Build()

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)
```

**When to use:** When you need specific log formats or output destinations

---

### Option 4: Log Level Control

Set the minimum log level:

```go
// Only log warnings and errors
config := zap.NewProductionConfig()
config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

logger, _ := config.Build()

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)
```

**Log levels:**
- `DebugLevel` - Very verbose, includes all details
- `InfoLevel` - General informational messages (default)
- `WarnLevel` - Warning messages
- `ErrorLevel` - Error messages only
- `FatalLevel` - Fatal errors (program exits)

**When to use:** Control log verbosity in different environments

---

### Option 5: Multiple Output Destinations

Write logs to multiple locations:

```go
// Write to both file and stdout
config := zap.NewProductionConfig()
config.OutputPaths = []string{
    "stdout",
    "/var/log/myapp/jamfpro.log",
}

logger, _ := config.Build()

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
)
```

**When to use:** When you need logs in both console and files

---

## What Gets Logged

The SDK automatically logs:

### Client Creation

```json
{
  "level": "info",
  "msg": "Jamf Pro client created",
  "instance_domain": "https://your-instance.jamfcloud.com",
  "auth_method": "oauth2"
}
```

### Configuration Changes

```json
{
  "level": "info",
  "msg": "HTTP timeout configured",
  "timeout": "30s"
}
```

### Token Operations

```json
{
  "level": "info",
  "msg": "Bearer token obtained",
  "expires_in": "30m"
}
```

```json
{
  "level": "info",
  "msg": "Bearer token refreshed"
}
```

## Adding Custom Logs

Add your own structured logs throughout your application:

```go
func processBuildings(logger *zap.Logger, jamfClient *jamfpro.Client) error {
    logger.Info("Starting building list retrieval")

    result, resp, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        logger.Error("Building list retrieval failed",
            zap.Error(err),
            zap.Int("status_code", resp.StatusCode),
        )
        return err
    }

    logger.Info("Building processing complete",
        zap.Int("building_count", len(result.Results)),
    )

    return nil
}
```

## Common Logging Patterns

### Pattern 1: Request/Response Logging

```go
logger.Info("API request",
    zap.String("method", "GET"),
    zap.String("endpoint", "/api/v1/buildings"),
)

result, resp, err := jamfClient.Buildings.ListV1(ctx, nil)

logger.Info("API response",
    zap.Int("status_code", resp.StatusCode),
    zap.Int("result_count", len(result.Results)),
)
```

### Pattern 2: Error Context Logging

```go
if err != nil {
    logger.Error("Operation failed",
        zap.Error(err),
        zap.String("operation", "list_buildings"),
        zap.String("instance", instanceDomain),
        zap.String("user_id", userID),
        zap.String("request_id", requestID),
    )
}
```

### Pattern 3: Performance Logging

```go
import "time"

start := time.Now()
result, _, err := jamfClient.Buildings.ListV1(ctx, nil)
duration := time.Since(start)

logger.Info("API call completed",
    zap.Duration("duration", duration),
    zap.Bool("success", err == nil),
    zap.String("endpoint", "/api/v1/buildings"),
    zap.Int("result_count", len(result.Results)),
)
```

## Integration with Log Aggregation

### Splunk

```go
config := zap.NewProductionConfig()
config.OutputPaths = []string{
    "/var/log/app/jamfpro.log", // Splunk monitors this file
}

logger, _ := config.Build()
```

### ELK Stack (Elasticsearch, Logstash, Kibana)

```go
// Use JSON format for Logstash parsing
logger, _ := zap.NewProduction()

// Add application metadata
logger = logger.With(
    zap.String("app", "jamfintegration"),
    zap.String("env", "production"),
    zap.String("version", "1.0.0"),
)

authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(authConfig, client.WithLogger(logger))
```

### Datadog

```go
// Configure for Datadog agent
config := zap.NewProductionConfig()
config.EncoderConfig.TimeKey = "@timestamp"
config.OutputPaths = []string{
    "/var/log/app/jamfpro.json", // Datadog agent watches this
}

logger, _ := config.Build()
```

## Debug Mode

Enable debug mode for detailed HTTP request/response logging:

```go
authConfig := client.AuthConfigFromEnv()
jamfClient, _ := jamfpro.NewClient(
    authConfig,
    client.WithLogger(logger),
    client.WithDebug(), // Enables detailed HTTP logging
)
```

**Output includes:**
- Full HTTP requests and responses
- Request/response headers
- Request/response bodies
- Timing information

**⚠️ Warning:** Debug mode logs sensitive data. Only use in development.

## Log Sampling

Reduce log volume in high-traffic scenarios:

```go
import "go.uber.org/zap/zapcore"

// Sample: Log 1 in every 100 info messages
config := zap.NewProductionConfig()
config.Sampling = &zap.SamplingConfig{
    Initial:    100,
    Thereafter: 100,
}

logger, _ := config.Build()
```

**When to use:** High-volume production environments

## Testing with Logs

### Capture Logs in Tests

```go
import "go.uber.org/zap/zaptest"

func TestMyFunction(t *testing.T) {
    // Create test logger that writes to test output
    logger := zaptest.NewLogger(t)

    authConfig := &client.AuthConfig{
        InstanceDomain: "https://test.jamfcloud.com",
        AuthMethod:     client.AuthMethodOAuth2,
        ClientID:       "test-id",
        ClientSecret:   "test-secret",
    }

    jamfClient, _ := jamfpro.NewClient(authConfig, client.WithLogger(logger))

    // Logs will appear in test output
    // ...
}
```

### Verify Log Output

```go
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "go.uber.org/zap/zaptest/observer"
)

func TestLogging(t *testing.T) {
    // Create observable logger for testing
    observedCore, observedLogs := observer.New(zapcore.InfoLevel)
    logger := zap.New(observedCore)

    authConfig := &client.AuthConfig{
        InstanceDomain: "https://test.jamfcloud.com",
        AuthMethod:     client.AuthMethodOAuth2,
        ClientID:       "test-id",
        ClientSecret:   "test-secret",
    }

    jamfClient, _ := jamfpro.NewClient(authConfig, client.WithLogger(logger))

    // Perform operations...

    // Verify logs
    logs := observedLogs.All()
    assert.Equal(t, 1, len(logs))
    assert.Contains(t, logs[0].Message, "Jamf Pro client created")
}
```

## Performance Considerations

- **Zap is fast** - <1μs per log message in production
- **JSON encoding** - Efficient for log aggregation
- **Sampling** - Reduces I/O in high-volume scenarios
- **Buffering** - Logs are buffered for better performance

## Related Documentation

- [OpenTelemetry Tracing](opentelemetry.md) - Combine logs with traces
- [Debugging](debugging.md) - Use debug mode for detailed inspection
- [Authentication](authentication.md) - Configure API access
- [Zap Documentation](https://pkg.go.dev/go.uber.org/zap)
