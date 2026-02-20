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

## SDK Services

### Jamf Pro API (REST)

- **Access & Account:** Access Management Settings, Account Preferences, API Integrations, API Roles, API Role Privileges
- **Enrollment & SSO:** Enrollment Settings, Reenrollment, Service Discovery Enrollment, SSO Certificate, SSO Settings, Adue Session Token Settings, Login Customization
- **Inventory & Groups:** Buildings, Categories, Computer Extension Attributes, Computer Groups (static/smart), Computer Prestages, Departments, Mobile Device Extension Attributes, Mobile Device Groups
- **Configuration & Distribution:** Cache Settings, Certificate Authority, Client Check-in, Cloud Distribution Point, Dock Items, Packages, Policy Properties, Scripts, Volume Purchasing Locations, Volume Purchasing Subscriptions
- **Self Service & Notifications:** Self Service Settings, Self Service Plus Settings, Notifications, Onboarding, Return to Service, Startup Status
- **Infrastructure:** Bookmarks, Icons, Jamf Pro Information, Jamf Pro Version, LDAP, Locales, SMTPServer, Time Zones
- **Other:** App Installers, Device Communication Settings, Policy Properties

### Classic API

- Accounts, Account Groups, Activation Code, Advanced Computer/User Searches, Allowed File Extensions, BYO Profiles, Classes, Directory Bindings, Disk Encryption Configurations, IBeacons, LDAP Servers, Network Segments, Patch External Sources, Printers, Removeable Mac Addresses, Restricted Software, Sites, Software Update Servers, VPP Accounts, Webhooks

## HTTP Client Configuration

The SDK uses a transport layer with bearer token auth, retries, and optional observability:

- **[Authentication](docs/guides/authentication.md)** — OAuth2 and Basic auth, environment variables, config files, and secure credential handling
- **Client options** — `WithLogger` (zap), `EnableTracing` (OpenTelemetry); pass options into `jamfpro.NewClient(authConfig, options...)` or `jamfpro.NewClientFromEnv(options...)`

The transport applies idempotent retries, sticky-session cookies for Jamf Cloud, and adaptive throttling. See [Quick Start](docs/guides/quick-start.md) for a minimal client and [Authentication](docs/guides/authentication.md) for production patterns.

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

```go
client.WithLogger(zapLogger)   // Structured logging (zap)
jamfClient.EnableTracing(otelConfig)  // OpenTelemetry HTTP tracing (call after NewClient)
```

See the [configuration guides](docs/guides/) for detailed documentation.

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
