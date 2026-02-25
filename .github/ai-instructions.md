# AI Instructions – Go SDK for Jamf Pro API

These instructions guide AI tools to follow the project's conventions and best practices when suggesting code. They cover code formatting, naming conventions, API implementation patterns, test strategies, and documentation requirements. By following these guidelines, AI suggestions should align with the project's style and help contributors produce high-quality, consistent code. Always review existing patterns in the repository—when in doubt, examine similar services or tests for reference.

## Development Setup & Workflow

- **Module Path:** `github.com/deploymenttheory/go-sdk-jamfpro-v2`
- **Go Version:** 1.25.0 or higher
- **HTTP Client:** Uses `resty.dev/v3` for HTTP operations
- **Testing:** Unit tests with mocks and acceptance tests against real Jamf Pro instances
- **Logging:** Structured logging with `go.uber.org/zap`
- **Tracing:** OpenTelemetry support for distributed tracing

### Running Tests

```bash
# Unit tests (with mocks)
go test ./... -v -run TestUnit

# Acceptance tests (requires real Jamf Pro instance)
INSTANCE_DOMAIN=https://your-instance.jamfcloud.com \
AUTH_METHOD=oauth2 \
CLIENT_ID=your-client-id \
CLIENT_SECRET=your-secret \
go test ./... -v -run TestAcceptance
```

### Environment Variables

- `INSTANCE_DOMAIN` – Jamf Pro instance URL (e.g., `https://your-instance.jamfcloud.com`)
- `AUTH_METHOD` – Authentication method (`oauth2` or `basic`)
- `CLIENT_ID` – OAuth2 client ID
- `CLIENT_SECRET` – OAuth2 client secret
- `JAMF_VERBOSE` – Set to `true` for verbose test logging

## Project Structure

### Service Organization

All API services are organized under `jamfpro/services/` with two main categories:

- **Jamf Pro API (REST):** `jamfpro/services/jamf_pro_api/{service_name}/`
  - Modern REST API with versioned endpoints
  - JSON request/response format
  - Supports pagination, RSQL filtering, and structured responses
  
- **Classic API (REST):** `jamfpro/services/classic_api/{service_name}/`
  - Legacy XML-based API
  - Maintained for services not yet migrated to the Jamf Pro API
  - Some operations only available in Classic API

### Service Files

Each service directory MUST contain:

- **`constants.go`** – API endpoint constants
  - Jamf Pro API: `EndpointServiceNameV1`, `EndpointServiceNameV2`
  - Classic API: `EndpointClassicServiceName`
  
- **`crud.go`** – CRUD operation implementations
  - Service interface definition
  - Service struct with HTTP client
  - All CRUD method implementations
  
- **`models.go`** – Data models and request/response structures
  
- **`crud_test.go`** – Unit tests with mocks
  - Named `TestUnit_<ServiceName>_<FunctionName>`
  - Uses external JSON/XML files for mock responses stored in the mocks directory.
  
- **`api_docs.md`** (optional) – Official Jamf API documentation copied verbatim for this service. **When present, treat this as the primary source of truth** for endpoint paths, HTTP methods, request/response schemas, required vs optional fields, enum values, and valid headers. Always read this file before implementing or reviewing any service code.

### Using api_docs.md as the Source of Truth

When `api_docs.md` exists in a service directory, **read it before writing or reviewing any code for that service**. It contains the official Jamf API documentation and is the authoritative reference for all implementation decisions.

**Endpoint validation:**
- Confirm the exact URL paths and HTTP methods match the documented endpoints
- Verify path parameter names (e.g., `{id}`) match the implementation
- Ensure endpoint constants in `constants.go` reflect the documented paths exactly

**Function availability:**
- Only implement functions for operations that exist in `api_docs.md`
- Do not add CRUD functions that have no corresponding endpoint in the documentation
- Non-standard operations (e.g., `client-credentials`) must map to a documented endpoint

**Header validation:**
- Derive `Accept` and `Content-Type` values from the curl examples in `api_docs.md`
- If the curl example includes only `--header 'accept: application/json'` with no `content-type`, do not send a `Content-Type` header (common for GET and body-less DELETE operations)
- If the curl example includes both `accept` and `content-type`, include both headers in the implementation

**Data model validation:**
- Fields marked `required` in `api_docs.md` MUST NOT use `omitempty` in their JSON struct tags
- Fields NOT marked `required` (optional) SHOULD use `omitempty`
- Enum values documented in `api_docs.md` (e.g., `CLIENT_CREDENTIALS`, `NATIVE_APP_OAUTH`, `NONE`) should be represented accurately in models or comments
- Array fields and nested object structures must match the documented schema hierarchy

**Request body validation:**
- Only include fields in request structs that appear under the documented Body Params section
- Required body fields must be validated before the API call (e.g., `if req == nil { return nil, nil, fmt.Errorf("request is required") }`)
- Optional body fields with documented defaults (e.g., `enabled: true`) should be noted in field comments

**Response schema validation:**
- Response structs must include all fields documented in the Response body
- Field names must match the JSON keys shown in the documented example responses
- Cross-reference the example JSON responses to confirm field types (string, integer, boolean, array)

**Mock data validation:**
- Mock JSON/XML files in `mocks/` must include all fields shown in the example responses in `api_docs.md`
- Required fields must be present; optional fields should also be populated for maximal coverage

### Mock Organization

Mock implementations for unit tests:

- **Location:** `jamfpro/services/{api_type}/{service_name}/mocks/`
- **Structure:**
  - `responders.go` – Mock client implementation with HTTP response handlers
  - `validate_get.json` – Mock response for Get operations
  - `validate_list.json` – Mock response for List operations
  - `validate_create.json` – Mock response for Create operations
  - `validate_update.json` – Mock response for Update operations
  - Additional scenario files as needed

### Example Files

Each service should have example implementations:

- **Location:** `examples/{api_type}/{service_name}/{operation}/main.go`
- **Structure:** One example per CRUD operation
  - `get/main.go`
  - `list/main.go`
  - `create/main.go`
  - `update/main.go`
  - `delete/main.go`
- **Pattern:** Examples must match the CRUD function name exactly

### Acceptance Tests

Acceptance tests verify operations against real Jamf Pro instances:

- **Location:** `jamfpro/acceptance/{api_type}/{service_name}_test.go`
- **Naming:** `TestAcceptance_<ServiceName>_<FunctionName>`
- **Strategies:** Follow patterns defined in `jamfpro/acceptance/test_strategies.md`
- **Requirements:** Must handle tenant-specific feature availability gracefully

## Naming Conventions

### Service Names

- Use lowercase with underscores for directory names (e.g., `cloud_information`, `computer_prestages`)
- Service struct named `Service` within the package
- Service interface named `{ServiceName}ServiceInterface`

### Function Names

**Jamf Pro API functions MUST include HTTP method verb and API version:**
- `GetByIDV1`, `GetByIDV2` – Retrieve single resource
- `ListV1`, `ListV2` – List resources (with pagination support)
- `CreateV1`, `CreateV2` – Create new resource
- `UpdateByIDV1`, `UpdateByIDV2` – Update existing resource
- `DeleteByIDV1`, `DeleteByIDV2` – Delete resource

**Classic API functions use verb without version:**
- `GetByID`, `GetByName` – Retrieve operations
- `List` – List operations
- `Create` – Create operations
- `Update`, `UpdateByID`, `UpdateByName` – Update operations
- `Delete`, `DeleteByID`, `DeleteByName` – Delete operations

**When multiple functions don't fit the pattern, use concise naming aligned with the API endpoint.**

Examples of non-standard naming when standard CRUD verbs don't fit:
- **Bulk operations:** `DeleteBuildingsByIDV1`, `DeletePackagesByIDV1` (for `/api/v1/{resource}/delete-multiple` endpoints)
- **File operations:** `UploadV1` (package upload to `/api/v1/packages/{id}/upload`), `DownloadV1` (script download from `/api/v1/scripts/{id}/download`)
- **Manifest operations:** `AssignManifestV1`, `DeleteManifestV1` (for `/api/v1/packages/{id}/manifest` endpoints)
- **History operations:** `GetHistoryV1`, `AddHistoryNotesV1` (for `/{resource}/{id}/history` endpoints)

**INCORRECT naming examples to avoid:**
- ❌ `ListScriptsV1` – Should be `ListV1` (resource context is already in the package name)
- ❌ `GetScriptByIDV1` – Should be `GetByIDV1`
- ❌ `CreateScriptV1` – Should be `CreateV1`
- ❌ `UpdatePackageByIDV1` – Should be `UpdateByIDV1`

**Use standard CRUD verbs (`ListV1`, `GetByIDV1`, `CreateV1`, `UpdateByIDV1`, `DeleteByIDV1`) unless:**
- The operation doesn't map to standard CRUD (upload, download, assign, export, etc.)
- The API endpoint path indicates a specialized action beyond basic CRUD

### Endpoint Constants

**Jamf Pro API:**
```go
const (
    EndpointCloudInformationV1 = "/api/v1/cloud-information"
    EndpointBuildingsV1        = "/api/v1/buildings"
    EndpointBuildingsV2        = "/api/v2/buildings"
)
```

**Classic API:**
```go
const (
    EndpointClassicBuildings = "/JSSResource/buildings"
    EndpointClassicPolicies  = "/JSSResource/policies"
)
```

### Test Naming

**Unit Tests:**
- Pattern: `TestUnit_<ServiceName>_<FunctionName>`
- Examples:
  - `TestUnit_CloudInformation_GetV1`
  - `TestUnit_Buildings_ListV1`
  - `TestUnit_Buildings_ListV1_WithRSQLFilter`
  - `TestUnit_Buildings_CreateV1_ValidationError`

**Acceptance Tests:**
- Pattern: `TestAcceptance_<ServiceName>_<FunctionName>`
- Examples:
  - `TestUnit_Buildings_ListV1`
  - `TestAcceptance_Buildings_ListV1_WithRSQLFilter`
  - `TestAcceptance_ClientCheckin_GetAndUpdate`

For accpetance tests that support full crud use - `TestAcceptance_<ServiceName>_Lifecycle`

### Model Names

- Request models: `Request{ResourceName}` (e.g., `RequestBuilding`, `RequestCategory`)
- Response models: `Resource{ResourceName}` (e.g., `ResourceBuilding`, `ResourceCategory`)
- List responses: `Resource{ResourceName}List` (e.g., `ResourceBuildingList`)

## Code Implementation Guidelines

### Service Structure

```go
package service_name

import (
    "context"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
    // ServiceNameServiceInterface defines the interface for service operations.
    //
    // Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
    ServiceNameServiceInterface interface {
        // GetV1 retrieves a resource by ID.
        //
        // Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
        GetV1(ctx context.Context, id string) (*ResourceType, *interfaces.Response, error)
    }

    // Service handles communication with the service-related methods of the Jamf Pro API.
    //
    // Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
    Service struct {
        client interfaces.HTTPClient
    }
)

var _ ServiceNameServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
    return &Service{client: client}
}
```

### CRUD Operations

**Endpoint Variable Pattern:**

All CRUD operations must follow this pattern:
1. Declare result variable (`var result ResultType`)
2. Assign endpoint constant to local variable (`endpoint := EndpointConstantName` or `endpoint := fmt.Sprintf(...)`)
3. Define headers map
4. Use the `endpoint` variable (not the constant directly) in HTTP client calls

This pattern ensures consistency and makes endpoints easy to trace during debugging.


**Get Operation:**
```go
// GetV1 retrieves a resource by ID.
// URL: GET /api/v1/resource/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-resource-id
func (s *Service) GetV1(ctx context.Context, id string) (*ResourceType, *interfaces.Response, error) {
    var result ResourceType

    endpoint := fmt.Sprintf("%s/%s", EndpointResourceV1, id)

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

**List Operation with Pagination:**
```go
// ListV1 retrieves all resources with pagination support.
// URL: GET /api/v1/resources
// https://developer.jamf.com/jamf-pro/reference/get_v1-resources
func (s *Service) ListV1(ctx context.Context, params map[string]string) (*ResourceList, *interfaces.Response, error) {
    var result ResourceList

    endpoint := EndpointResourcesV1

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    resp, err := s.client.GetPagination(ctx, endpoint, params, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

**Create Operation:**
```go
// CreateV1 creates a new resource.
// URL: POST /api/v1/resources
// https://developer.jamf.com/jamf-pro/reference/post_v1-resources
func (s *Service) CreateV1(ctx context.Context, req *RequestResource) (*ResourceType, *interfaces.Response, error) {
    var result ResourceType

    endpoint := EndpointResourcesV1

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

**Update Operation:**
```go
// UpdateByIDV1 updates an existing resource by ID.
// URL: PUT /api/v1/resources/{id}
// https://developer.jamf.com/jamf-pro/reference/put_v1-resources-id
func (s *Service) UpdateByIDV1(ctx context.Context, id string, req *RequestResource) (*ResourceType, *interfaces.Response, error) {
    var result ResourceType

    endpoint := fmt.Sprintf("%s/%s", EndpointResourcesV1, id)

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

**Delete Operation:**
```go
// DeleteByIDV1 deletes a resource by ID.
// URL: DELETE /api/v1/resources/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-resources-id
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
    endpoint := fmt.Sprintf("%s/%s", EndpointResourcesV1, id)

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    resp, err := s.client.Delete(ctx, endpoint, headers)
    if err != nil {
        return resp, err
    }

    return resp, nil
}
```

### Headers

**CRITICAL: Validate all headers against `api_docs.md` when it exists.** Derive `Accept` and `Content-Type` values from the curl examples in the documentation. Do not add `Content-Type` to requests that don't send a body (e.g., GET, body-less DELETE) unless the documentation curl example explicitly includes it.

Common header patterns:

- **JSON operations:** `Accept: application/json`, `Content-Type: application/json`
- **XML operations (Classic API):** `Accept: application/xml`, `Content-Type: application/xml`
- **File uploads:** `Content-Type: multipart/form-data`

### RSQL Query Support

**MANDATORY: Implement RSQL query parameters for List functions that support filtering.**

```go
// ListV1 supports RSQL filtering via the 'filter' query parameter.
// URL: GET /api/v1/resources?filter=name=="value"
// https://developer.jamf.com/jamf-pro/reference/get_v1-resources
func (s *Service) ListV1(ctx context.Context, params map[string]string) (*ResourceList, *interfaces.Response, error) {
    var result ResourceList

    endpoint := EndpointResourcesV1

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    // params can include: filter, page, page-size, sort
    resp, err := s.client.GetPagination(ctx, endpoint, params, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

## Testing Requirements

### Unit Tests

**MANDATORY Requirements:**

1. **All API functions MUST have unit test coverage**
2. **Test naming:** `TestUnit_<ServiceName>_<FunctionName>`
3. **Mock responses:** Use externalized JSON (Jamf Pro API) or XML (Classic API) files in `mocks/` directory
4. **Maximal field definition:** Mock responses must validate the full data model with all fields populated
5. **Error scenarios:** Test both success and error cases

**Unit Test Structure:**

```go
package service_name

import (
    "context"
    "testing"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_name/mocks"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ServiceMock) {
    t.Helper()
    mock := mocks.NewServiceMock()
    return NewService(mock), mock
}

func TestUnit_ServiceName_GetV1_Success(t *testing.T) {
    svc, mock := setupMockService(t)
    mock.RegisterGetResourceMock()

    result, resp, err := svc.GetV1(context.Background(), "test-id")
    require.NoError(t, err)
    require.NotNil(t, result)
    require.NotNil(t, resp)
    assert.Equal(t, 200, resp.StatusCode)
    assert.Equal(t, "test-id", result.ID)
    assert.NotEmpty(t, result.Name)
}

func TestUnit_ServiceName_GetV1_Error(t *testing.T) {
    svc, mock := setupMockService(t)
    mock.RegisterErrorMock()

    result, resp, err := svc.GetV1(context.Background(), "invalid-id")
    assert.Error(t, err)
    assert.Nil(t, result)
    assert.NotNil(t, resp)
    assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_ServiceName_ListV1_Success(t *testing.T) {
    svc, mock := setupMockService(t)
    mock.RegisterListResourcesMock()

    params := map[string]string{"page": "0", "page-size": "100"}
    result, resp, err := svc.ListV1(context.Background(), params)
    require.NoError(t, err)
    require.NotNil(t, result)
    require.NotNil(t, resp)
    assert.Equal(t, 200, resp.StatusCode)
    assert.GreaterOrEqual(t, result.TotalCount, 0)
}
```

### Acceptance Tests

**MANDATORY Requirements:**

1. **All acceptance tests MUST follow one of the predefined strategies** (see `jamfpro/acceptance/test_strategies.md`)
2. **Test naming:** `TestAcceptance_<ServiceName>_<TestType>`
3. **Handle tenant-specific features:** Use `t.Skipf()` when features are not enabled
4. **RSQL testing:** If List supports RSQL, MUST include `TestAcceptance_<ServiceName>_List<API_Version>_WithRSQLFilter`
5. **Cleanup:** Always register cleanup handlers to delete test resources

**See `jamfpro/acceptance/test_strategies.md` for comprehensive acceptance test patterns including:**
- Full CRUD Lifecycle Pattern
- Settings/Configuration Pattern
- Read-Only Information Pattern
- Read-Only with Existing Data Pattern
- RSQL Filter Pattern (MANDATORY when supported)
- Bulk Operations Pattern
- Validation Errors Pattern

### Test Coverage

- **Unit tests:** All functions must have corresponding unit tests
- **Acceptance tests:** Must pass or have handlers when service is not enabled in the tenant
- **Mock data:** JSON/XML files must be maximal in field definition

## Documentation Requirements

### Function Comments

**MANDATORY: All exported functions MUST have:**

1. **Function description** – What the function does
2. **URL format** – The exact API endpoint called
3. **API documentation link** – Valid Jamf Pro API documentation URL

```go
// GetV1 retrieves a building by ID.
// URL: GET /api/v1/buildings/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-buildings-id
func (s *Service) GetV1(ctx context.Context, id string) (*ResourceBuilding, *interfaces.Response, error) {
    // implementation
}
```

### API Documentation URLs

**Ensure all documentation URLs are valid and accurate:**

- Base URL: `https://developer.jamf.com/jamf-pro/reference/`
- Format: `https://developer.jamf.com/jamf-pro/reference/{operation}_{version}-{resource}`
- Examples:
  - `https://developer.jamf.com/jamf-pro/reference/get_v1-buildings`
  - `https://developer.jamf.com/jamf-pro/reference/post_v1-buildings`
  - `https://developer.jamf.com/jamf-pro/reference/put_v1-buildings-id`
  - `https://developer.jamf.com/jamf-pro/reference/delete_v1-buildings-id`

### Interface Comments

```go
// ServiceNameServiceInterface defines the interface for service operations.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
ServiceNameServiceInterface interface {
    // GetV1 retrieves a resource by ID.
    //
    // Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/...
    GetV1(ctx context.Context, id string) (*ResourceType, *interfaces.Response, error)
}
```

## Examples Requirements

**MANDATORY: All API functions MUST have a corresponding example.**

### Example Structure

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"

    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
    // Initialize client from environment or config file
    authConfig, err := client.AuthConfigFromEnv()
    if err != nil {
        log.Fatalf("Failed to load auth config: %v", err)
    }
    
    jamfClient, err := jamfpro.NewClient(authConfig)
    if err != nil {
        log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
    }

    // Call the API function
    result, resp, err := jamfClient.ServiceName.FunctionName(context.Background(), "param")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // Marshal and display the full response body
    out, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
        log.Fatalf("Failed to marshal result: %v", err)
    }
    
    fmt.Printf("Response (Status: %d):\n%s\n", resp.StatusCode, string(out))
}
```

**Key requirements:**
- Example file name must match the CRUD function name
- Must unmarshall response into JSON or XML to show the full response body
- Must handle errors gracefully
- Should use `json.MarshalIndent` for readable output

## Service Registration

**MANDATORY: All services MUST be registered in `jamfpro/new.go`.**

```go
// In new.go
import (
    "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/service_name"
)

type Client struct {
    // ... other services
    ServiceName *service_name.Service
}

func NewClient(authConfig *client.AuthConfig, opts ...client.Option) (*Client, error) {
    // ... initialization code
    
    return &Client{
        // ... other services
        ServiceName: service_name.NewService(httpClient),
    }, nil
}
```

## Pagination and Filtering

### Pagination Support

**MANDATORY: All List functions that support pagination MUST use the `GetPagination` transport method.**

```go
func (s *Service) ListV1(ctx context.Context, params map[string]string) (*ResourceList, *interfaces.Response, error) {
    var result ResourceList

    endpoint := EndpointResourcesV1

    headers := map[string]string{
        "Accept":       mime.ApplicationJSON,
        "Content-Type": mime.ApplicationJSON,
    }

    // GetPagination handles page, page-size, and other query parameters
    resp, err := s.client.GetPagination(ctx, endpoint, params, headers, &result)
    if err != nil {
        return nil, resp, err
    }

    return &result, resp, nil
}
```

### RSQL Query Implementation

**MANDATORY: Jamf Pro API functions that support RSQL queries MUST have proper implementation and testing.**

For more information on RSQL syntax and usage, see `rsql.md` and `paginated.md` documentation files.

## Quality Checklist

Before submitting code, verify ALL requirements from `reqs.md` are met:

1. ✅ All API functions have an example in `examples/{api_type}/{service_name}/{operation}/main.go`
2. ✅ All functions have correct comment styling with valid API documentation URLs
3. ✅ All unit tests are named `TestUnit_<ServiceName>_<FunctionName>`
4. ✅ All acceptance tests are named `TestAcceptance_<ServiceName>_<FunctionName>`
5. ✅ All List functions supporting pagination use the `GetPagination` transport method
6. ✅ All Jamf Pro API functions supporting RSQL have proper implementation
7. ✅ All unit tests for List, Get, and errors use externalized JSON (Jamf Pro API) or XML (Classic API) mock responses with maximal field definition
8. ✅ All services are registered in `jamfpro/new.go`
9. ✅ All Classic API endpoint constants are named `EndpointClassic<Service_Name>`
10. ✅ All headers accurately reflect the API documentation for each function
11. ✅ All examples unmarshall responses into JSON or XML showing the full response body
12. ✅ All unit tests pass
13. ✅ All acceptance tests pass or have handlers when the service is not enabled in the Jamf Pro tenant (with comments)
14. ✅ All acceptance tests follow one of the predefined acceptance test strategies
15. ✅ All Jamf Pro API functions describe the HTTP method verb with API version (e.g., `GetByIDV1`, `CreateV2`). Classic API uses verb only (e.g., `GetByID`, `Create`)
16. ✅ Each example matches the name of the CRUD function

## Common Patterns

### Error Handling

```go
if err != nil {
    return nil, resp, fmt.Errorf("failed to get resource: %w", err)
}
```

### Context Usage

Always accept `context.Context` as the first parameter for all API operations to support cancellation and timeouts.

### Response Handling

Always return a triple: `(*Result, *interfaces.Response, error)` for operations that return data, or `(*interfaces.Response, error)` for operations that don't (like Delete).

### Validation

Perform client-side validation before making network calls:

```go
if id == "" {
    return nil, nil, fmt.Errorf("ID is required")
}

if req == nil {
    return nil, nil, fmt.Errorf("request is required")
}
```

## Best Practices

- **Keep functions focused:** Each function should do one thing well
- **Use descriptive variable names:** Prefer clarity over brevity
- **Handle errors gracefully:** Always check and return errors with context
- **Document assumptions:** If a function has specific requirements, document them
- **Follow existing patterns:** Review similar services for consistency
- **Test thoroughly:** Unit tests with mocks, acceptance tests with real API
- **Update documentation:** Never edit `api_docs.md` as it is generated from the official Jamf documentation.

## Code Style

- Use `gofmt` for formatting
- Follow standard Go naming conventions
- Keep line length reasonable (aim for 120 characters or less)
- Use blank lines to separate logical blocks
- Group imports: standard library, external packages, internal packages

## Contributing

When implementing new services or modifying existing ones:

1. Read this document thoroughly
2. Review similar services for patterns
3. Check `jamfpro/acceptance/test_strategies.md` for test patterns
4. Verify all quality checklist items before submitting
5. Run both unit and acceptance tests
6. Update documentation and examples
7. Ensure service is registered in `jamfpro/new.go`
