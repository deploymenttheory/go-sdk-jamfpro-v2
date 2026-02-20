# Authentication

## Overview

The Jamf Pro SDK uses bearer token authentication. You configure either **OAuth2** (client credentials) or **Basic** (username/password); the SDK exchanges these for a bearer token and refreshes it automatically before expiry.

## Why Use Proper Authentication?

- **Secure credentials** — Avoid hardcoding secrets in source code
- **Automatic token refresh** — Tokens are refreshed within a buffer period so requests do not fail mid-session
- **Support multiple environments** — Use different credentials for dev, staging, and production
- **Audit and safety** — Use `HideSensitiveData: true` so tokens are not written to logs

## When to Use It

Use proper authentication whenever:

- Calling the Jamf Pro API from any application
- Deploying to production
- Storing code in version control
- Running automated tests or CI/CD against Jamf Pro

## Basic Example

Recommended: load credentials from the environment and create the client with `NewClientFromEnv()`:

```go
package main

import (
	"context"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	// Require credentials from environment
	if os.Getenv("INSTANCE_DOMAIN") == "" || os.Getenv("AUTH_METHOD") == "" {
		log.Fatal("INSTANCE_DOMAIN and AUTH_METHOD are required")
	}

	jamfClient, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Use the client — authentication is automatic
	result, _, err := jamfClient.Reenrollment.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Re-enrollment configured: %+v", result)
}
```

**Run:**

```bash
export INSTANCE_DOMAIN="https://your-instance.jamfcloud.com"
export AUTH_METHOD="oauth2"
export CLIENT_ID="your-client-id"
export CLIENT_SECRET="your-client-secret"
go run main.go
```

## Token Lifecycle

The SDK obtains a bearer token at startup and refreshes it automatically before it expires. The refresh buffer defaults to 5 minutes; you can change it via `AuthConfig.TokenRefreshBufferPeriod` or in a config file with `token_refresh_buffer_period_seconds`.

- **OAuth2:** `POST /api/v1/oauth/token` (client credentials)
- **Basic:** `POST /api/v1/auth/token` (username/password exchanged for token)

No application code is required for refresh; it is handled inside the transport.

## Configuration Options

### Option 1: Environment Variables (Recommended for production)

Required:

- `INSTANCE_DOMAIN` — Jamf Pro base URL (e.g. `https://your-instance.jamfcloud.com`)
- `AUTH_METHOD` — `oauth2` or `basic`

For OAuth2:

- `CLIENT_ID`
- `CLIENT_SECRET`

For Basic:

- `BASIC_AUTH_USERNAME`
- `BASIC_AUTH_PASSWORD`

Optional:

- `TOKEN_REFRESH_BUFFER_SECONDS` — seconds before expiry to refresh (default 300)
- `HIDE_SENSITIVE_DATA` — set to `true` to redact bearer tokens in logs

```go
jamfClient, err := jamfpro.NewClientFromEnv()
```

**When to use:** Production and CI; keeps secrets out of code and supports 12-factor style configuration.

**Example (Linux/macOS):**

```bash
export INSTANCE_DOMAIN="https://your-instance.jamfcloud.com"
export AUTH_METHOD="oauth2"
export CLIENT_ID="your-client-id"
export CLIENT_SECRET="your-client-secret"
```

**Example (Kubernetes secret):**

```bash
kubectl create secret generic jamf-credentials \
  --from-literal=instance_domain="https://..." \
  --from-literal=auth_method="oauth2" \
  --from-literal=client_id="..." \
  --from-literal=client_secret="..."
```

---

### Option 2: Config File

Load credentials from a JSON file (do not commit this file):

```go
authConfig, err := client.LoadAuthConfigFromFile("clientconfig.json")
if err != nil {
	log.Fatal(err)
}
jamfClient, err := jamfpro.NewClient(authConfig)
```

**File format (`clientconfig.json`):**

```json
{
  "instance_domain": "https://your-instance.jamfcloud.com",
  "auth_method": "oauth2",
  "client_id": "your-client-id",
  "client_secret": "your-client-secret",
  "token_refresh_buffer_period_seconds": 300,
  "hide_sensitive_data": true
}
```

For Basic auth use `"auth_method": "basic"` with `basic_auth_username` and `basic_auth_password`.

**When to use:** Local development or per-developer configuration.

**Add to `.gitignore`:**

```
clientconfig.json
*.local.json
```

---

### Option 3: AuthConfig struct

Build `AuthConfig` programmatically (e.g. from a secret manager):

```go
authConfig := &client.AuthConfig{
	InstanceDomain:           "https://your-instance.jamfcloud.com",
	AuthMethod:               client.AuthMethodOAuth2,
	ClientID:                 secretClientID,
	ClientSecret:             secretClientSecret,
	TokenRefreshBufferPeriod: 5 * time.Minute,
	HideSensitiveData:        true,
}
jamfClient, err := jamfpro.NewClient(authConfig, client.WithLogger(logger))
```

**When to use:** When credentials come from a vault or another runtime source.

---

### Option 4: From environment into AuthConfig

Use the same env vars as `NewClientFromEnv` but get an `AuthConfig` for validation or custom client setup:

```go
authConfig := client.AuthConfigFromEnv()
if err := authConfig.Validate(); err != nil {
	log.Fatal(err)
}
jamfClient, err := jamfpro.NewClient(authConfig, client.WithLogger(zapLogger))
```

## Security Best Practices

### Do

- Store credentials in environment variables or a secret manager in production
- Use different credentials per environment
- Set `HideSensitiveData: true` in production to avoid logging tokens
- Add config files that contain secrets to `.gitignore`
- Revoke and rotate credentials if they are exposed

### Do not

- Hardcode client IDs, client secrets, or passwords in source code
- Commit credential files to version control
- Share credentials in plaintext (email, chat, etc.)
- Use production credentials in development or tests
- Log credentials or bearer tokens

## Troubleshooting

### Authentication failed (401 / invalid token)

**Causes:**

- Invalid or expired OAuth2 client ID/secret or Basic username/password
- Wrong `INSTANCE_DOMAIN` (e.g. missing `https://`, or wrong tenant)
- `AUTH_METHOD` typo (must be `oauth2` or `basic`)

**Check:**

- Verify env vars: `echo $INSTANCE_DOMAIN $AUTH_METHOD`
- For OAuth2, confirm the client is enabled in Jamf Pro (Settings → Global Management → API Roles & Clients)
- For Basic, ensure the account has API access and is not locked

### Initial token fetch failed

**Symptoms:** `NewClient` or `NewClientFromEnv` returns an error such as "initial token fetch failed".

**Check:**

- Network connectivity to `INSTANCE_DOMAIN`
- Correct URL (no trailing path; use `https://instance.jamfcloud.com` not `https://instance.jamfcloud.com/api`)
- Credentials are set and valid (see above)

## Testing with Authentication

### Unit tests

Use mocks that implement the HTTP client interface so tests do not call the real API. See service packages under `jamfpro/services/jamf_pro_api/.../mocks` for examples.

### Acceptance tests

The SDK’s acceptance tests use environment variables (same as production). Set `INSTANCE_DOMAIN`, `AUTH_METHOD`, and the appropriate credentials, then run:

```bash
go test -v -count=1 ./jamfpro/acceptance/...
```

## Related Documentation

- [Quick Start](quick-start.md) — Installation and first API call
- [Jamf Pro API Authentication](https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes) — Official auth overview
