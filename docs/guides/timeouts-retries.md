# Timeouts & Retries

## What are Timeouts and Retries?

Timeouts control how long the SDK waits for an API response before giving up. Retries automatically retry failed requests when they encounter transient errors like network issues or rate limits.

## Why Use Timeouts and Retries?

Proper timeout and retry configuration helps you:

- **Prevent hanging requests** - Avoid waiting indefinitely for responses
- **Handle transient failures** - Automatically recover from temporary network issues
- **Respect rate limits** - Retry with backoff when hitting API quotas
- **Improve reliability** - Make your application more resilient to intermittent failures
- **Control resource usage** - Free up resources from slow or failing requests

## When to Use It

Configure timeouts and retries when:

- Making API calls over unreliable networks
- Running long-lived services that need resilience
- Implementing critical workflows that must handle transient failures
- Dealing with rate-limited APIs
- Running in production environments where reliability is critical

## Basic Example

Here's how to configure timeouts and retries:

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
    authConfig := client.AuthConfigFromEnv()

    // Create client with timeout and retry configuration
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithTimeout(30*time.Second),  // 30 second timeout
        client.WithRetryCount(3),             // Retry up to 3 times
    )
    if err != nil {
        log.Fatal(err)
    }

    // Use the client - timeouts and retries are automatic
    result, _, err := jamfClient.Buildings.ListV1(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found %d buildings", len(result.Results))
}
```

**What happens:**
- If the request takes longer than 30 seconds, it times out
- If the request fails with a retryable error, it automatically retries up to 3 times
- Retries use exponential backoff to avoid overwhelming the server

## Alternative Configuration Options

### Option 1: Custom Timeout

Set a timeout appropriate for your use case:

```go
// Short timeout for quick operations
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTimeout(10*time.Second),
)

// Longer timeout for bulk operations
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTimeout(5*time.Minute),
)
```

**When to use:**
- Short timeouts (5-15s): Simple lookups, list operations
- Medium timeouts (30-60s): Updates, creates
- Long timeouts (2-5min): Bulk operations, large data transfers

**Default:** 120 seconds (2 minutes)

---

### Option 2: Retry Configuration

Configure retry behavior for different scenarios.

**Available retry options:**
- `WithRetryCount(n)` - How many times to retry (default: 3)
- `WithRetryWaitTime(d)` - Initial wait time before first retry (default: 2s)
- `WithRetryMaxWaitTime(d)` - Maximum wait time between retries (default: 10s)

The wait time doubles with each retry (exponential backoff) up to the maximum.

```go
import "time"

// Conservative: Few retries, quick backoff
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithRetryCount(2),                      // Retry twice
    client.WithRetryWaitTime(1*time.Second),       // Wait 1s initially
    client.WithRetryMaxWaitTime(5*time.Second),    // Max wait 5s
)

// Aggressive: More retries, longer backoff
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithRetryCount(5),                      // Retry 5 times
    client.WithRetryWaitTime(3*time.Second),       // Wait 3s initially
    client.WithRetryMaxWaitTime(30*time.Second),   // Max wait 30s
)
```

**When to use:**
- Conservative: Rate-limited APIs, quick failures preferred
- Aggressive: Unreliable networks, high importance operations

**Defaults:**
- Retry count: 3
- Wait time: 2 seconds
- Max wait time: 10 seconds

---

### Option 3: Total Retry Duration

Set a maximum wall-clock time for all retry attempts:

```go
import "time"

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithRetryCount(10),
    client.WithTotalRetryDuration(2*time.Minute), // Max 2 minutes total
)
```

**When to use:** When you need a hard deadline for operations regardless of retry count

---

## Retry Behavior

### What Gets Retried

The SDK automatically retries:
- ✅ Network errors (connection refused, timeout, etc.)
- ✅ 5xx server errors (500, 502, 503, 504)
- ✅ 429 rate limit errors
- ✅ Request timeout errors

### What Doesn't Get Retried

The SDK does NOT retry:
- ❌ 4xx client errors (400, 401, 403, 404) - these won't succeed on retry
- ❌ Successful responses (2xx)
- ❌ Context cancellation
- ❌ Invalid request configuration

### Exponential Backoff

Retries use exponential backoff with jitter to prevent overwhelming servers:

**With default settings (2s initial, 10s max):**
```
Retry 1: Wait ~2s   (base wait time)
Retry 2: Wait ~4s   (2x backoff)
Retry 3: Wait ~8s   (4x backoff, approaching max)
```

## Related Documentation

- [Authentication](authentication.md) - Configure API access
- [Logging](logging.md) - Log timeout and retry events
- [Debugging](debugging.md) - Debug timeout issues
