# Proxy Support

## What is Proxy Support?

Proxy support allows the SDK to route all HTTP traffic through an intermediate proxy server. This is essential for corporate environments, privacy requirements, or network architectures that mandate proxy usage.

## Why Use a Proxy?

Proxy configuration helps you:

- **Corporate requirements** - Route traffic through corporate proxies
- **Access control** - Comply with network security policies
- **Privacy** - Mask client IP addresses
- **Logging & monitoring** - Centralize traffic inspection
- **Regional access** - Route through specific geographic locations

## When to Configure It

Configure a proxy when:

- Working in corporate environments with mandatory proxies
- Behind firewalls that require proxy for external access
- Need to route traffic through specific geographic locations
- Required for compliance or security policies
- Testing proxy behavior in development

## Basic Example

Here's how to configure a proxy:

```go
package main

import (
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
    authConfig := client.AuthConfigFromEnv()

    // Configure client with HTTP proxy
    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithProxy("http://proxy.company.com:8080"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // All requests now route through the proxy
}
```

## Configuration Options

### Option 1: HTTP Proxy

Configure a standard HTTP proxy:

```go
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithProxy("http://proxy.example.com:8080"),
)
```

**When to use:** Most common proxy type, standard corporate proxies

---

### Option 2: HTTPS Proxy

Use an HTTPS proxy for encrypted proxy connections:

```go
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithProxy("https://secure-proxy.example.com:8443"),
)
```

**When to use:** When proxy connection itself needs to be encrypted

---

## Related Documentation

- [TLS Configuration](tls-configuration.md) - Configure certificates for proxy SSL inspection
- [Authentication](authentication.md) - OAuth2 and Basic auth configuration
- [Timeouts & Retries](timeouts-retries.md) - Adjust timeouts for proxy connections
- [Debugging](debugging.md) - Debug proxy connection issues
