# TLS/SSL Configuration

## What is TLS Configuration?

TLS (Transport Layer Security) configuration controls how the SDK establishes secure HTTPS connections to the Jamf Pro API. You can customize certificate validation, use mutual TLS, and set minimum TLS versions.

## Why Configure TLS?

TLS configuration helps you:

- **Use custom certificates** - Work with private CAs or self-signed certificates
- **Enable mutual TLS** - Use client certificates for enhanced authentication
- **Meet security requirements** - Enforce minimum TLS versions (TLS 1.2, 1.3)
- **Corporate environments** - Integrate with enterprise certificate infrastructures
- **Compliance** - Meet regulatory requirements for encryption

## When to Configure It

Configure TLS when:

- Using private or internal CAs
- Required to use client certificates
- Working behind corporate proxies with SSL inspection
- Meeting compliance requirements (PCI-DSS, HIPAA, etc.)
- Enforcing specific TLS versions
- Testing with self-signed certificates (development only)

## Basic Example

Here's how to configure basic TLS settings:

```go
package main

import (
    "crypto/tls"
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
    authConfig := client.AuthConfigFromEnv()

    // Create client with minimum TLS 1.2
    tlsConfig := &tls.Config{
        MinVersion: tls.VersionTLS12,
    }

    jamfClient, err := jamfpro.NewClient(
        authConfig,
        client.WithTLSClientConfig(tlsConfig),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Use client normally
    // All connections now use TLS 1.2 or higher
}
```

## Configuration Options

### Option 1: Minimum TLS Version

Enforce a minimum TLS version for all connections:

```go
import "crypto/tls"

// Require TLS 1.2 or higher (recommended)
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTLSClientConfig(tlsConfig),
)

// Require TLS 1.3 (most secure)
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS13,
}

jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTLSClientConfig(tlsConfig),
)
```

**When to use:**
- TLS 1.2: Industry standard, widely compatible
- TLS 1.3: Maximum security, modern systems

**Default:** System default (usually TLS 1.2+)

---

### Option 2: Custom TLS Configuration

Full control over TLS settings:

```go
import "crypto/tls"

tlsConfig := &tls.Config{
    MinVersion:         tls.VersionTLS12,
    MaxVersion:         tls.VersionTLS13,
    InsecureSkipVerify: false, // NEVER use true in production
    CipherSuites: []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
    },
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTLSClientConfig(tlsConfig),
)
```

**When to use:**
- Specific cipher suite requirements
- Custom security policies
- Advanced TLS configuration needs

---

### Option 3: Disable Certificate Verification (DEVELOPMENT ONLY)

⚠️ **WARNING**: Only for development/testing with self-signed certificates!

```go
// NEVER use this in production!
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithInsecureSkipVerify(),
)
```

**When to use:** Testing with self-signed certificates in development

**⚠️ Security Risk:** Disables certificate validation, vulnerable to MITM attacks

---

## Common Scenarios

### Scenario 1: Corporate Proxy with SSL Inspection

```go
import "crypto/tls"

// Load corporate CA certificate (if needed)
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTLSClientConfig(tlsConfig),
    client.WithProxy("http://proxy.company.com:8080"),
)
```

### Scenario 2: High Security Environment

```go
import "crypto/tls"

// Enforce TLS 1.3 with strong ciphers
tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS13,
    CipherSuites: []uint16{
        tls.TLS_AES_256_GCM_SHA384,
        tls.TLS_CHACHA20_POLY1305_SHA256,
    },
}

authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithTLSClientConfig(tlsConfig),
)
```

### Scenario 3: Development with Self-Signed Certificates

```go
// Temporarily skip verification (NOT for production)
authConfig := client.AuthConfigFromEnv()
jamfClient, err := jamfpro.NewClient(
    authConfig,
    client.WithInsecureSkipVerify(),
)
```

## Security Best Practices

✅ **Do:**
- Use TLS 1.2 or higher
- Validate certificates in production
- Use strong cipher suites
- Keep TLS libraries up to date

❌ **Don't:**
- Disable certificate verification in production
- Use self-signed certificates in production
- Use outdated TLS versions (1.0, 1.1)
- Ignore certificate expiration warnings

## Related Documentation

- [Proxy Support](proxy.md) - Configure proxies (often used with custom CAs)
- [Authentication](authentication.md) - OAuth2 and Basic auth configuration
- [Debugging](debugging.md) - Debug TLS handshake issues
- [Go TLS Package](https://pkg.go.dev/crypto/tls)
