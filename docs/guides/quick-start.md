# Quick Start Guide

Get up and running with the Jamf Pro Go SDK in minutes.

## Prerequisites

- Go 1.21 or higher
- A Jamf Pro instance and credentials:
  - **OAuth2 (recommended):** Client ID and Client Secret from Jamf Pro → Settings → Global Management → API Roles & Clients
  - **Basic auth:** Jamf Pro admin username and password (exchanged for a bearer token by the SDK)

## Installation

```bash
go get github.com/deploymenttheory/go-sdk-jamfpro-v2
```

## Your First API Call

Here's a complete example that fetches re-enrollment settings:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	// Step 1: Create the client from environment variables
	jamfClient, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Step 2: Make an API call using a service
	result, resp, err := jamfClient.Reenrollment.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Step 3: Use the results
	fmt.Printf("Re-enrollment flush policy history: %v\n", result.FlushPolicyHistory)
	fmt.Printf("Status Code: %d\n", resp.StatusCode)
}
```

**Run it:**

```bash
export INSTANCE_DOMAIN="https://your-instance.jamfcloud.com"
export AUTH_METHOD="oauth2"
export CLIENT_ID="your-client-id"
export CLIENT_SECRET="your-client-secret"
go run main.go
```

## Common Operations

### List and get by ID (e.g. API Integrations)

```go
// List (paginated)
list, resp, err := jamfClient.ApiIntegrations.ListV1(ctx, nil)
if err != nil {
	log.Fatal(err)
}
for _, item := range list.Results {
	fmt.Println(item.DisplayName)
}

// Get by ID
item, resp, err := jamfClient.ApiIntegrations.GetByIDV1(ctx, "1")
if err != nil {
	log.Fatal(err)
}
fmt.Println(item.DisplayName)
```

### Create and delete (e.g. Categories)

```go
import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"

created, resp, err := jamfClient.Categories.CreateV1(ctx, &categories.RequestCategory{
	Name: "My Category",
	Priority: 1,
})
if err != nil {
	log.Fatal(err)
}
id := created.ID

// Later: delete
_, err = jamfClient.Categories.DeleteByIDV1(ctx, id)
```

### List computer prestages and get device scope

```go
list, _, err := jamfClient.ComputerPrestages.ListV3(ctx, nil)
if err != nil {
	log.Fatal(err)
}

scope, _, err := jamfClient.ComputerPrestages.GetDeviceScopeByIDV2(ctx, list.Results[0].ID)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Prestage %s scope version: %d\n", scope.PrestageId, scope.VersionLock)
```

## Authentication

The SDK supports two authentication methods; both result in a bearer token that is automatically refreshed.

### From environment variables

Required:

- `INSTANCE_DOMAIN` — e.g. `https://your-instance.jamfcloud.com`
- `AUTH_METHOD` — `oauth2` or `basic`

For OAuth2:

- `CLIENT_ID`
- `CLIENT_SECRET`

For Basic:

- `BASIC_AUTH_USERNAME`
- `BASIC_AUTH_PASSWORD`

Optional:

- `TOKEN_REFRESH_BUFFER_SECONDS` — default 300
- `HIDE_SENSITIVE_DATA` — set to `true` to redact tokens in logs

```go
jamfClient, err := jamfpro.NewClientFromEnv()
```

### From a config file

```go
authConfig, err := client.LoadAuthConfigFromFile("/path/to/clientconfig.json")
if err != nil {
	log.Fatal(err)
}
jamfClient, err := jamfpro.NewClient(authConfig)
```

See [Authentication](authentication.md) for file format, secret managers, and best practices.

## Error Handling

Always check errors and inspect the response when needed:

```go
result, resp, err := jamfClient.ApiIntegrations.GetByIDV1(ctx, id)

if err != nil {
	// Check for specific error types if your client exposes them
	log.Printf("API error: %v", err)
	return
}

if resp != nil && resp.StatusCode == 404 {
	log.Println("Resource not found")
	return
}

// Success
fmt.Println(result.DisplayName)
```

## Response Metadata

Service methods return `(result, resp, err)`. The `resp` value (e.g. `*interfaces.Response`) carries HTTP metadata:

```go
result, resp, err := jamfClient.Categories.ListV1(ctx, nil)
// resp.StatusCode, resp.Headers, resp.Body (if needed)
```

## Next Steps

### Configuration guides

- **[Authentication](authentication.md)** — OAuth2 vs Basic, environment variables, config files, and secure credential handling

### Client options

- **Logger:** `jamfpro.NewClient(authConfig, client.WithLogger(zapLogger))`
- **OpenTelemetry:** After creating the client, call `jamfClient.EnableTracing(otelConfig)` for HTTP tracing

### Examples and API coverage

- Browse [examples/jamf_pro_api/](../../examples/jamf_pro_api/) and [examples/classic_api/](../../examples/classic_api/) for service-specific samples
- [Jamf Pro API Reference](https://developer.jamf.com/jamf-pro/reference)

## Getting Help

- **[Full documentation](../../README.md)** — SDK overview and service list
- **[Jamf Pro API docs](https://developer.jamf.com/jamf-pro/reference)** — Official API reference
- **[GitHub Issues](https://github.com/deploymenttheory/go-sdk-jamfpro-v2/issues)** — Bugs and feature requests
- **[GoDoc](https://pkg.go.dev/github.com/deploymenttheory/go-sdk-jamfpro-v2)** — Package documentation
