# Jamf Pro SDK Acceptance Test Strategies

This document defines standardized test patterns for Jamf Pro API acceptance tests. Following these patterns ensures consistency, maintainability, and comprehensive coverage across all service tests.

## Table of Contents

1. [Test Pattern Categories](#test-pattern-categories)
2. [Pattern Definitions](#pattern-definitions)
3. [Pattern Selection Guide](#pattern-selection-guide)
4. [Common Guidelines](#common-guidelines)
5. [Examples](#examples)

---

## Test Pattern Categories

### 1. Full CRUD Lifecycle Pattern

**Use for:** Services that support complete Create, Read, Update, Delete operations on resources (e.g., Buildings, Categories, Departments, Scripts, Packages)

**Operations tested:**
1. Create resource
2. List resources (verify creation)
3. GetByID (retrieve created resource)
4. Update resource
5. GetByID (verify update)
6. History operations (if supported):
   - AddHistoryNotes
   - GetHistory
7. Delete resource

**Test structure:**
```go
func TestAcceptance_ServiceName_Lifecycle(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName
    ctx := context.Background()

    // 1. Create
    acc.LogTestStage(t, "Create", "Creating test resource")
    createReq := &servicename.RequestType{
        Name: acc.UniqueName("acc-test-resource"),
        // ... other required fields
    }
    created, createResp, err := svc.CreateV1(ctx, createReq)
    require.NoError(t, err, "Create should not return an error")
    require.NotNil(t, created)
    assert.Equal(t, 201, createResp.StatusCode)
    assert.NotEmpty(t, created.ID)

    resourceID := created.ID
    acc.LogTestSuccess(t, "Resource created with ID=%s", resourceID)

    // Register cleanup
    acc.Cleanup(t, func() {
        cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        _, delErr := svc.DeleteByIDV1(cleanupCtx, resourceID)
        acc.LogCleanupDeleteError(t, "resource", resourceID, delErr)
    })

    // 2. List — verify creation
    acc.LogTestStage(t, "List", "Listing resources to verify creation")
    ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
    defer cancel2()

    list, listResp, err := svc.ListV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
    require.NoError(t, err)
    require.NotNil(t, list)
    assert.Equal(t, 200, listResp.StatusCode)

    found := false
    for _, item := range list.Results {
        if item.ID == resourceID {
            found = true
            assert.Equal(t, createReq.Name, item.Name)
            break
        }
    }
    assert.True(t, found, "newly created resource should appear in list")
    acc.LogTestSuccess(t, "Resource ID=%s found in list (%d total)", resourceID, list.TotalCount)

    // 3. GetByID
    acc.LogTestStage(t, "GetByID", "Fetching resource by ID=%s", resourceID)
    fetched, fetchResp, err := svc.GetByIDV1(ctx, resourceID)
    require.NoError(t, err)
    require.NotNil(t, fetched)
    assert.Equal(t, 200, fetchResp.StatusCode)
    assert.Equal(t, resourceID, fetched.ID)
    assert.Equal(t, createReq.Name, fetched.Name)
    acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

    // 4. Update
    acc.LogTestStage(t, "Update", "Updating resource ID=%s", resourceID)
    updateReq := &servicename.RequestType{
        Name: acc.UniqueName("acc-test-resource-updated"),
        // ... other fields
    }
    updated, updateResp, err := svc.UpdateByIDV1(ctx, resourceID, updateReq)
    require.NoError(t, err)
    require.NotNil(t, updated)
    assert.Equal(t, 200, updateResp.StatusCode)
    acc.LogTestSuccess(t, "Resource updated: ID=%s", resourceID)

    // 5. GetByID — verify update
    fetched2, _, err := svc.GetByIDV1(ctx, resourceID)
    require.NoError(t, err)
    assert.Equal(t, updateReq.Name, fetched2.Name)
    acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

    // 6. History operations (if supported)
    acc.LogTestStage(t, "History", "Adding history note and fetching history")
    noteReq := &servicename.AddHistoryNotesRequest{
        Note: fmt.Sprintf("Acceptance test note at %s", time.Now().Format(time.RFC3339)),
    }
    noteResp, err := svc.AddHistoryNotesV1(ctx, resourceID, noteReq)
    require.NoError(t, err)
    assert.Contains(t, []int{200, 201}, noteResp.StatusCode)
    acc.LogTestSuccess(t, "History note added")

    history, histResp, err := svc.GetHistoryV1(ctx, resourceID, nil)
    require.NoError(t, err)
    assert.Equal(t, 200, histResp.StatusCode)
    assert.GreaterOrEqual(t, history.TotalCount, 1)
    acc.LogTestSuccess(t, "History entries: %d", history.TotalCount)

    // 7. Delete
    acc.LogTestStage(t, "Delete", "Deleting resource ID=%s", resourceID)
    deleteResp, err := svc.DeleteByIDV1(ctx, resourceID)
    require.NoError(t, err)
    assert.Equal(t, 204, deleteResp.StatusCode)
    acc.LogTestSuccess(t, "Resource ID=%s deleted", resourceID)
}
```

---

### 2. Settings/Configuration Pattern

**Use for:** Services that manage singleton settings/configurations that cannot be created or deleted (e.g., ClientCheckin, CacheSettings, DeviceCommunicationSettings, EnrollmentSettings)

**Operations tested:**

1. Get current settings (store original state)
2. Update settings to test values
3. Get settings again (verify update applied)
4. Update settings back to original values
5. Get settings final time (verify restoration)

**Test structure:**

```go
func TestAcceptance_SettingsName_GetAndUpdate(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.SettingsName
    ctx := context.Background()

    // 1. Get and store current settings
    acc.LogTestStage(t, "Get", "Fetching current settings")
    original, resp, err := svc.GetV1(ctx)
    require.NoError(t, err)
    require.NotNil(t, original)
    assert.Equal(t, 200, resp.StatusCode)
    acc.LogTestSuccess(t, "Retrieved current settings")

    // 2. Update settings to test values
    acc.LogTestStage(t, "Update", "Updating settings to test values")
    updateReq := &settingsname.RequestSettings{
        // Modify one or more settings
        SomeSetting: !original.SomeSetting, // Toggle a boolean for example
        OtherSetting: original.OtherSetting + 10, // Modify a numeric value
    }

    updateResp, err := svc.UpdateV1(ctx, updateReq)
    require.NoError(t, err)
    require.NotNil(t, updateResp)
    assert.Equal(t, 200, updateResp.StatusCode)
    acc.LogTestSuccess(t, "Settings updated to test values")

    // 3. Get and verify test update applied
    acc.LogTestStage(t, "Get", "Verifying test update applied")
    updated, getResp, err := svc.GetV1(ctx)
    require.NoError(t, err)
    require.NotNil(t, updated)
    assert.Equal(t, 200, getResp.StatusCode)
    assert.Equal(t, updateReq.SomeSetting, updated.SomeSetting)
    assert.Equal(t, updateReq.OtherSetting, updated.OtherSetting)
    acc.LogTestSuccess(t, "Test update verified")

    // 4. Restore original settings
    acc.LogTestStage(t, "Restore", "Restoring original settings")
    restoreResp, err := svc.UpdateV1(ctx, original)
    require.NoError(t, err)
    require.NotNil(t, restoreResp)
    assert.Equal(t, 200, restoreResp.StatusCode)
    acc.LogTestSuccess(t, "Original settings restored")

    // 5. Get and verify restoration
    acc.LogTestStage(t, "Get", "Verifying restoration")
    restored, finalResp, err := svc.GetV1(ctx)
    require.NoError(t, err)
    require.NotNil(t, restored)
    assert.Equal(t, 200, finalResp.StatusCode)
    assert.Equal(t, original.SomeSetting, restored.SomeSetting)
    assert.Equal(t, original.OtherSetting, restored.OtherSetting)
    acc.LogTestSuccess(t, "Restoration verified - settings back to original state")

    // Safety net: Cleanup to restore settings in case test fails mid-execution
    acc.Cleanup(t, func() {
        cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        _, restoreErr := svc.UpdateV1(cleanupCtx, original)
        if restoreErr != nil {
            acc.LogTestWarning(t, "Cleanup: failed to restore original settings: %v", restoreErr)
        }
    })
}
```

---

### 3. Read-Only Information Pattern

**Use for:** Services that only provide read access to system information (e.g., JamfProVersion, JamfProInformation, StartupStatus, TimeZones, Locales)

**Operations tested:**
1. Get information
2. Verify response structure and required fields

**Test structure:**
```go
func TestAcceptance_ServiceName_GetV1(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName
    ctx := context.Background()

    result, resp, err := svc.GetV1(ctx)
    require.NoError(t, err)
    require.NotNil(t, result)
    require.NotNil(t, resp)
    assert.Equal(t, 200, resp.StatusCode)

    // Verify expected fields are present and valid
    assert.NotEmpty(t, result.SomeRequiredField)
    assert.GreaterOrEqual(t, result.SomeNumericField, 0)
}
```

---

### 4. Read-Only with Existing Data Pattern

**Use for:** Services where you can only read pre-existing data but cannot create test data (e.g., CloudIdp, MobileDeviceGroups when testing read operations)

**Operations tested:**
1. List resources
2. If resources exist: GetByID, GetByName, or other read operations
3. Skip test if no data available

**Test structure:**
```go
func TestAcceptance_ServiceName_ListAndGet(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName
    ctx := context.Background()

    // 1. List resources
    list, _, err := svc.ListV1(ctx, nil)
    if err != nil {
        t.Skipf("Failed to list resources (may not be supported on this tenant): %v", err)
        return
    }

    assert.NotNil(t, list)
    assert.GreaterOrEqual(t, list.TotalCount, 0)

    // 2. Skip if no data available
    if len(list.Results) == 0 {
        t.Skip("No resources available for testing")
        return
    }

    // 3. Test GetByID with first resource
    firstID := list.Results[0].ID
    result, _, err := svc.GetByIDV1(ctx, firstID)
    require.NoError(t, err)
    assert.Equal(t, firstID, result.ID)
    assert.NotEmpty(t, result.Name)
}
```

---

### 5. RSQL Filter Pattern

**Use for:** Services that support RSQL filtering in List operations

**Operations tested:**
1. Create a resource with unique identifiable data
2. List with RSQL filter to find that specific resource
3. Verify filter works correctly

**Test structure:**
```go
func TestAcceptance_ServiceName_ListWithRSQLFilter(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName
    ctx := context.Background()

    // 1. Create resource with unique name
    name := acc.UniqueName("acc-rsql-test")
    createReq := &servicename.RequestType{
        Name: name,
        // ... other fields
    }

    created, _, err := svc.CreateV1(ctx, createReq)
    require.NoError(t, err)
    require.NotNil(t, created)

    resourceID := created.ID
    acc.LogTestSuccess(t, "Created resource ID=%s name=%q", resourceID, name)

    acc.Cleanup(t, func() {
        cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        _, delErr := svc.DeleteByIDV1(cleanupCtx, resourceID)
        acc.LogCleanupDeleteError(t, "resource", resourceID, delErr)
    })

    // 2. List with RSQL filter
    rsqlQuery := map[string]string{
        "filter": fmt.Sprintf(`name=="%s"`, name),
    }

    list, listResp, err := svc.ListV1(ctx, rsqlQuery)
    require.NoError(t, err)
    require.NotNil(t, list)
    assert.Equal(t, 200, listResp.StatusCode)

    // 3. Verify filtered results
    found := false
    for _, item := range list.Results {
        if item.ID == resourceID {
            found = true
            assert.Equal(t, name, item.Name)
            break
        }
    }
    assert.True(t, found, "resource should appear in RSQL-filtered results")
    acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target resource found=%v", list.TotalCount, found)
}
```

---

### 6. Bulk Operations Pattern

**Use for:** Services that support bulk delete or other bulk operations

**Operations tested:**
1. Create multiple resources
2. Perform bulk operation (e.g., bulk delete)
3. Verify operation succeeded

**Test structure:**
```go
func TestAcceptance_ServiceName_BulkDelete(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName
    ctx := context.Background()

    // 1. Create multiple resources
    ids := make([]string, 0, 2)
    for i := 0; i < 2; i++ {
        req := &servicename.RequestType{
            Name: acc.UniqueName(fmt.Sprintf("acc-bulk-delete-%d", i)),
            // ... other fields
        }

        ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
        created, resp, err := svc.CreateV1(ctx1, req)
        cancel1()
        require.NoError(t, err, "Create %d should succeed", i)
        require.NotNil(t, created)
        assert.Equal(t, 201, resp.StatusCode)
        ids = append(ids, created.ID)
        acc.LogTestSuccess(t, "Bulk test: created resource ID=%s", created.ID)
    }

    // Safety net cleanup
    acc.Cleanup(t, func() {
        cleanCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
        defer cancel()
        for _, id := range ids {
            _, delErr := svc.DeleteByIDV1(cleanCtx, id)
            acc.LogCleanupDeleteError(t, "resource", id, delErr)
        }
    })

    // 2. Bulk delete
    acc.LogTestStage(t, "BulkDelete", "Deleting %d resources: %v", len(ids), ids)

    ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
    defer cancel2()

    bulkResp, err := svc.DeleteResourcesByIDV1(ctx2, &servicename.DeleteByIDRequest{IDs: ids})
    require.NoError(t, err, "Bulk delete should not return an error")
    require.NotNil(t, bulkResp)
    assert.Equal(t, 204, bulkResp.StatusCode)
    acc.LogTestSuccess(t, "Bulk delete of %d resources succeeded", len(ids))

    // 3. Verify deletion
    for _, id := range ids {
        ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
        _, _, getErr := svc.GetByIDV1(ctx3, id)
        cancel3()
        assert.Error(t, getErr, "deleted resource ID=%s should return error on Get", id)
    }
}
```

---

### 7. Validation Errors Pattern

**Use for:** Testing client-side validation without making network calls

**Operations tested:**
1. Test various validation scenarios (empty IDs, nil requests, etc.)
2. Verify appropriate errors are returned

**Test structure:**
```go
func TestAcceptance_ServiceName_ValidationErrors(t *testing.T) {
    acc.RequireClient(t)
    svc := acc.Client.ServiceName

    t.Run("GetByID_EmptyID", func(t *testing.T) {
        _, _, err := svc.GetByIDV1(context.Background(), "")
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "ID is required")
    })

    t.Run("Create_NilRequest", func(t *testing.T) {
        _, _, err := svc.CreateV1(context.Background(), nil)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "request is required")
    })

    t.Run("UpdateByID_EmptyID", func(t *testing.T) {
        _, _, err := svc.UpdateByIDV1(context.Background(), "", &servicename.RequestType{Name: "x"})
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "id is required")
    })

    t.Run("DeleteByID_EmptyID", func(t *testing.T) {
        _, err := svc.DeleteByIDV1(context.Background(), "")
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "ID is required")
    })
}
```

---

## Pattern Selection Guide

| Service Type | Recommended Pattern(s) |
|--------------|------------------------|
| Standard CRUD resources (Buildings, Categories, Departments, Scripts, Packages) | Pattern 1 (Full CRUD Lifecycle) + Pattern 5 (RSQL Filter - **MANDATORY if List supports RSQL**) + Pattern 6 (Bulk Operations if supported) + Pattern 7 (Validation) |
| Settings/Configuration (ClientCheckin, CacheSettings, EnrollmentSettings) | Pattern 2 (Settings/Configuration) |
| Read-only system info (JamfProVersion, StartupStatus, Locales) | Pattern 3 (Read-Only Information) |
| Cloud services with existing data (CloudIdp, CloudAzure) | Pattern 4 (Read-Only with Existing Data) + Pattern 5 (RSQL Filter - **MANDATORY if List supports RSQL**) |
| Special operations (CSA token exchange, certificate validation) | Custom tests based on specific operations |

### MANDATORY: RSQL Testing for List Operations

**If a service's List operation supports RSQL filtering, you MUST include both:**

1. **Generic List Test** - List without filters to verify basic pagination and results
2. **RSQL Filter Test** - List with RSQL filter to verify query functionality (Pattern 5)

**Why this is mandatory:**
- RSQL filtering is a critical feature for production API consumers
- Filtering logic can break independently of basic List operations
- Verifies the API correctly processes and applies filter criteria
- Tests that filter results match expectations

---

## Common Guidelines

### 1. Test Naming

- **Lifecycle tests:** `TestAcceptance_ServiceName_Lifecycle`
- **Settings tests:** `TestAcceptance_ServiceName_GetAndUpdate`
- **Read-only tests:** `TestAcceptance_ServiceName_GetV1`
- **Generic list tests:** `TestAcceptance_ServiceName_List` (embedded in Lifecycle test)
- **RSQL filter tests:** `TestAcceptance_ServiceName_ListWithRSQLFilter` (**MANDATORY if RSQL supported**)
- **Bulk operation tests:** `TestAcceptance_ServiceName_BulkDelete`
- **Validation tests:** `TestAcceptance_ServiceName_ValidationErrors`

### 2. Resource Naming

- Use `acc.UniqueName("acc-test-{service}")` to generate unique names
- This prevents conflicts in shared test environments
- Helps identify test-created resources

### 3. Cleanup

- **Always** register cleanup with `acc.Cleanup(t, func() {...})`
- Use a timeout context for cleanup operations
- Log cleanup errors with `acc.LogCleanupDeleteError()`
- Cleanup runs even if test fails

### 4. Logging

- Use `acc.LogTestStage()` for major test phases
- Use `acc.LogTestSuccess()` for successful operations
- Logs are visible when `JAMF_VERBOSE=true`
- Logs integrate with GitHub Actions annotations

### 5. Context Handling

- Use `context.Background()` for main test context
- Create timeout contexts for individual operations: `context.WithTimeout(ctx, acc.Config.RequestTimeout)`
- Always defer `cancel()` after creating timeout contexts

### 6. Assertions

- Use `require.*` for critical assertions (test should stop if these fail)
- Use `assert.*` for non-critical assertions (test continues after failure)
- Always check error, response, and result values

### 7. Tenant Compatibility

- Use `t.Skipf()` for operations that may not be supported on all tenants
- Check for empty lists before testing GetByID with existing data
- Handle cloud services gracefully when not configured

### 8. Response Status Codes

- **201:** Create operations
- **200:** Get, Update, List operations
- **204:** Delete operations
- Always assert on status codes

### 9. List Operation Testing (MANDATORY)

**For ANY service with a List operation:**

1. **Generic List Test** - Always include a basic list operation
   - Test without filters or with minimal pagination parameters
   - Verify response structure (TotalCount, Results array)
   - Verify at least basic fields are populated

2. **RSQL Filter Test** - If the API supports RSQL filtering, you MUST include:
   - Create a resource with unique, filterable data
   - Use RSQL query to filter for that specific resource
   - Verify filtered results contain only expected matches
   - This is **NOT optional** - RSQL support must be tested

**How to determine if RSQL is supported:**
- Check API documentation for List endpoint
- Look for `filter` query parameter support
- If uncertain, attempt a simple filter and document behavior
- Most Jamf Pro API v1/v2 List endpoints support RSQL

**Example:**
```go
// CORRECT: Tests both generic list AND RSQL filter
func TestAcceptance_Buildings_Lifecycle(t *testing.T) {
    // ... create building ...

    // Generic list test (part of lifecycle)
    list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "200"})
    // ... verify results ...
}

func TestAcceptance_Buildings_ListWithRSQLFilter(t *testing.T) {
    // ... create building with unique name ...

    // RSQL filter test (separate test)
    list, _, err := svc.ListV1(ctx, map[string]string{"filter": `name=="unique-name"`})
    // ... verify filtered results ...
}

// INCORRECT: Only tests generic list, missing RSQL test
func TestAcceptance_Buildings_Lifecycle(t *testing.T) {
    // ... only includes generic list, no separate RSQL test ...
}
```

---

## Examples

### Example: Full CRUD Service (Buildings)

See `buildings_test.go` for a complete implementation of:
- Pattern 1: Full CRUD Lifecycle
- Pattern 5: RSQL Filter
- Pattern 7: Validation Errors

### Example: Settings Service (ClientCheckin)

See `client_checkin_test.go` for Pattern 2: Settings/Configuration

### Example: Read-Only Service (JamfProVersion)

See `jamf_pro_version_test.go` for Pattern 3: Read-Only Information

### Example: Cloud Service (CloudIdp)

See `cloud_idp_test.go` for Pattern 4: Read-Only with Existing Data

---

## Version History

- **v1.0** (2026-02-21): Initial test strategies document
  - Defined 7 core test patterns
  - Established naming conventions and guidelines
  - Documented common practices

---

## Contributing

When adding new acceptance tests:

1. Identify which pattern(s) apply to your service
2. Follow the appropriate template structure
3. Include all recommended test functions for that pattern
4. Add logging at appropriate stages
5. Implement proper cleanup
6. Test both success and error scenarios

If your service doesn't fit existing patterns, document the new pattern and add it to this guide.
