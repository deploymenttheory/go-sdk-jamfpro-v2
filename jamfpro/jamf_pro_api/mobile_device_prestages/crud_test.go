package mobile_device_prestages

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_prestages/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListV3 tests listing all mobile device prestages.
func TestUnit_MobileDevicePrestages_ListV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.ListV3(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Test iPad Prestage", result.Results[0].DisplayName)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Test iPhone Prestage", result.Results[1].DisplayName)
	assert.True(t, *result.Results[0].Mandatory)
	assert.Equal(t, "IT", result.Results[0].Department)
}

// TestGetByIDV3 tests retrieving a mobile device prestage by ID.
func TestUnit_MobileDevicePrestages_GetByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test iPad Prestage", result.DisplayName)
	assert.True(t, *result.Mandatory)
	assert.False(t, *result.MdmRemovable)
	assert.Equal(t, "555-1234", result.SupportPhoneNumber)
	assert.Equal(t, "support@example.com", result.SupportEmailAddress)
	assert.Equal(t, "IT", result.Department)
	assert.True(t, *result.RequireAuthentication)
	assert.Equal(t, "Please authenticate", result.AuthenticationPrompt)
	assert.True(t, *result.PreventActivationLock)
	assert.Equal(t, "1", result.DeviceEnrollmentProgramInstanceID)
	assert.Equal(t, "testuser", result.LocationInformation.Username)
	assert.Equal(t, "Test User", result.LocationInformation.Realname)
	assert.Equal(t, "PREFIX_SUFFIX", result.Names.AssignNamesUsing)
	assert.Equal(t, "iPad-", result.Names.DeviceNamePrefix)
	assert.True(t, *result.Names.ManageNames)
	assert.Equal(t, 1, result.VersionLock)
}

// TestGetByNameV3 tests retrieving a mobile device prestage by display name.
func TestUnit_MobileDevicePrestages_GetByNameV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetByNameV3(context.Background(), "Test iPad Prestage")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test iPad Prestage", result.DisplayName)
	assert.Equal(t, "IT", result.Department)
}

// TestGetByNameV3_NotFound tests retrieving a mobile device prestage by name when not found.
func TestUnit_MobileDevicePrestages_GetByNameV3_NotFound(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetByNameV3(context.Background(), "Nonexistent Prestage")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "not found")
}

// TestCreateV3 tests creating a new mobile device prestage.
func TestUnit_MobileDevicePrestages_CreateV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterCreateMock()

	svc := NewMobileDevicePrestages(mock)

	mandatory := true
	mdmRemovable := false
	supervised := true

	prestage := &ResourceMobileDevicePrestage{
		DisplayName:                       "New Test Prestage",
		Mandatory:                         &mandatory,
		MdmRemovable:                      &mdmRemovable,
		SupportPhoneNumber:                "555-9999",
		SupportEmailAddress:               "support@newtest.com",
		Department:                        "Engineering",
		DeviceEnrollmentProgramInstanceID: "1",
		Supervised:                        &supervised,
		LocationInformation: SubsetLocationInformation{
			Username: "newuser",
			Realname: "New User",
		},
		PurchasingInformation: SubsetPurchasingInformation{},
		Names: SubsetNames{
			AssignNamesUsing: "STATIC",
			SingleDeviceName: "TestDevice",
		},
	}

	result, resp, err := svc.CreateV3(context.Background(), prestage)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v3/mobile-device-prestages/3")
}

// TestUpdateByIDV3 tests updating a mobile device prestage by ID.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetByIDMock("1")    // internal GET for version lock
	mock.RegisterUpdateByIDMock("1") // PUT

	svc := NewMobileDevicePrestages(mock)

	mandatory := false
	mdmRemovable := true
	supervised := true

	prestage := &ResourceMobileDevicePrestage{
		DisplayName:                       "Updated Prestage",
		Mandatory:                         &mandatory,
		MdmRemovable:                      &mdmRemovable,
		SupportPhoneNumber:                "555-8888",
		SupportEmailAddress:               "updated@example.com",
		Department:                        "Updated Dept",
		DeviceEnrollmentProgramInstanceID: "1",
		Supervised:                        &supervised,
		LocationInformation:               SubsetLocationInformation{},
		PurchasingInformation:             SubsetPurchasingInformation{},
		Names: SubsetNames{
			AssignNamesUsing: "STATIC",
		},
	}

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", prestage)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test iPad Prestage", result.DisplayName)
}

// TestUpdateByNameV3 tests updating a mobile device prestage by display name.
// The list response provides the version lock – no separate GET is needed.
func TestUnit_MobileDevicePrestages_UpdateByNameV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()       // GetByNameV3 (also provides version lock)
	mock.RegisterUpdateByIDMock("1") // PUT

	svc := NewMobileDevicePrestages(mock)

	mandatory := false
	supervised := true

	prestage := &ResourceMobileDevicePrestage{
		DisplayName:                       "Updated Prestage",
		Mandatory:                         &mandatory,
		DeviceEnrollmentProgramInstanceID: "1",
		Supervised:                        &supervised,
		LocationInformation:               SubsetLocationInformation{},
		PurchasingInformation:             SubsetPurchasingInformation{},
		Names: SubsetNames{
			AssignNamesUsing: "STATIC",
		},
	}

	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test iPad Prestage", prestage)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
}

// TestDeleteByIDV3 tests deleting a mobile device prestage by ID.
func TestUnit_MobileDevicePrestages_DeleteByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	resp, err := svc.DeleteByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestDeleteByNameV3 tests deleting a mobile device prestage by display name.
func TestUnit_MobileDevicePrestages_DeleteByNameV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()
	mock.RegisterDeleteByIDMock("2")

	svc := NewMobileDevicePrestages(mock)
	resp, err := svc.DeleteByNameV3(context.Background(), "Test iPhone Prestage")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestGetScopeByIDV2 tests retrieving device scope for a mobile device prestage.
func TestUnit_MobileDevicePrestages_GetScopeByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetScopeByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetScopeByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.PrestageId)
	assert.Len(t, result.Assignments, 2)
	assert.Equal(t, "C02ABCDEFGH", result.Assignments[0].SerialNumber)
	assert.Equal(t, "2024-01-15T10:30:00Z", result.Assignments[0].AssignmentDate)
	assert.Equal(t, "admin@example.com", result.Assignments[0].UserAssigned)
	assert.Equal(t, "C02XYZABCDE", result.Assignments[1].SerialNumber)
	assert.Equal(t, 0, result.VersionLock)
}

// TestListV3_Empty tests listing mobile device prestages when the list is empty.
func TestUnit_MobileDevicePrestages_ListV3_Empty(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterEmptyListMock()

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.ListV3(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 0, result.TotalCount)
	assert.Len(t, result.Results, 0)
}

// TestCreateV3_NilPrestage tests creating with a nil prestage object.
func TestUnit_MobileDevicePrestages_CreateV3_NilPrestage(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.CreateV3(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "prestage is required")
}

// TestCreateV3_EmptyDisplayName tests creating with an empty display name.
func TestUnit_MobileDevicePrestages_CreateV3_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	prestage := &ResourceMobileDevicePrestage{
		DisplayName: "",
	}

	result, resp, err := svc.CreateV3(context.Background(), prestage)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

// TestUpdateByIDV3_NilPrestage tests updating with a nil prestage object.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_NilPrestage(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "prestage is required")
}

// TestUpdateByIDV3_EmptyID tests updating with an empty ID.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	prestage := &ResourceMobileDevicePrestage{
		DisplayName: "Test",
	}

	result, resp, err := svc.UpdateByIDV3(context.Background(), "", prestage)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestUpdateByIDV3_EmptyDisplayName tests updating with an empty display name.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_EmptyDisplayName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	prestage := &ResourceMobileDevicePrestage{
		DisplayName: "",
	}

	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", prestage)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "display name is required")
}

// TestUpdateByNameV3_EmptyName tests updating with an empty name.
func TestUnit_MobileDevicePrestages_UpdateByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	prestage := &ResourceMobileDevicePrestage{
		DisplayName: "Test",
	}

	result, resp, err := svc.UpdateByNameV3(context.Background(), "", prestage)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// TestDeleteByIDV3_EmptyID tests deleting with an empty ID.
func TestUnit_MobileDevicePrestages_DeleteByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	resp, err := svc.DeleteByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

// TestDeleteByNameV3_EmptyName tests deleting with an empty name.
func TestUnit_MobileDevicePrestages_DeleteByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	resp, err := svc.DeleteByNameV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "name is required")
}

// TestGetScopeByIDV2_EmptyID tests getting scope with an empty ID.
func TestUnit_MobileDevicePrestages_GetScopeByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetScopeByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByIDV3_EmptyID tests getting by ID with an empty ID.
func TestUnit_MobileDevicePrestages_GetByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// TestGetByNameV3_EmptyName tests getting by name with an empty name.
func TestUnit_MobileDevicePrestages_GetByNameV3_EmptyName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetByNameV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "name is required")
}

// Optimistic locking tests

// TestUnit_MobileDevicePrestages_UpdateByIDV3_VersionLockPropagated verifies that
// UpdateByIDV3 fetches the current resource and injects the versionLock (and
// sub-resource locks) into the caller-supplied prestage before issuing the PUT.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_VersionLockPropagated(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetByIDMock("1")    // returns versionLock=1
	mock.RegisterUpdateByIDMock("1") // PUT

	svc := NewMobileDevicePrestages(mock)
	prestage := &ResourceMobileDevicePrestage{DisplayName: "Updated Prestage"}
	_, _, err := svc.UpdateByIDV3(context.Background(), "1", prestage)
	require.NoError(t, err)
	// EnsureVersionLock must have copied versionLock=1 from the GET response.
	assert.Equal(t, 1, prestage.VersionLock)
	// Sub-resource locks are propagated from the current resource (0 in fixture).
	assert.Equal(t, 0, prestage.LocationInformation.VersionLock)
	assert.Equal(t, 0, prestage.PurchasingInformation.VersionLock)
}

// TestUnit_MobileDevicePrestages_UpdateByNameV3_VersionLockPropagated verifies that
// UpdateByNameV3 reuses the versionLock obtained from the list response during
// the name lookup, avoiding a second round-trip.
func TestUnit_MobileDevicePrestages_UpdateByNameV3_VersionLockPropagated(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterListMock()          // "Test iPad Prestage" carries versionLock=1
	mock.RegisterUpdateByIDMock("1") // PUT

	svc := NewMobileDevicePrestages(mock)
	prestage := &ResourceMobileDevicePrestage{DisplayName: "Updated Prestage"}
	_, _, err := svc.UpdateByNameV3(context.Background(), "Test iPad Prestage", prestage)
	require.NoError(t, err)
	assert.Equal(t, 1, prestage.VersionLock)
}

// TestUnit_MobileDevicePrestages_UpdateByIDV3_FetchVersionLockError verifies that
// UpdateByIDV3 returns an error when the internal GET (for version locking) fails.
func TestUnit_MobileDevicePrestages_UpdateByIDV3_FetchVersionLockError(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No GET mock registered – the internal fetch will fail.

	svc := NewMobileDevicePrestages(mock)
	prestage := &ResourceMobileDevicePrestage{DisplayName: "Updated Prestage"}
	result, resp, err := svc.UpdateByIDV3(context.Background(), "1", prestage)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current prestage for version locking")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// TestUnit_MobileDevicePrestages_UpdateByNameV3_FetchVersionLockError verifies that
// UpdateByNameV3 returns an error when the list call (used for name lookup and
// version lock retrieval) fails.
func TestUnit_MobileDevicePrestages_UpdateByNameV3_FetchVersionLockError(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No list mock registered – GetByNameV3 will fail.

	svc := NewMobileDevicePrestages(mock)
	prestage := &ResourceMobileDevicePrestage{DisplayName: "Updated Prestage"}
	result, resp, err := svc.UpdateByNameV3(context.Background(), "Test iPad Prestage", prestage)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get mobile device prestage by name")
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

// --- Scope operations ---

func TestUnit_MobileDevicePrestages_ReplaceScopeByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterReplaceScopeByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	req := &RequestReplaceScope{SerialNumbers: []string{"C02NEW123"}, VersionLock: 0}
	result, resp, err := svc.ReplaceScopeByIDV2(context.Background(), "1", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.PrestageId)
	assert.Len(t, result.Assignments, 2)
}

func TestUnit_MobileDevicePrestages_ReplaceScopeByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestReplaceScope{SerialNumbers: []string{"C02NEW123"}}

	result, resp, err := svc.ReplaceScopeByIDV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_ReplaceScopeByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.ReplaceScopeByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDevicePrestages_AddScopeByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterAddScopeByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	req := &RequestAddScope{SerialNumbers: []string{"C02NEW456"}, VersionLock: 0}
	result, resp, err := svc.AddScopeByIDV2(context.Background(), "1", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.PrestageId)
}

func TestUnit_MobileDevicePrestages_AddScopeByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestAddScope{SerialNumbers: []string{"C02NEW456"}}

	result, resp, err := svc.AddScopeByIDV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_AddScopeByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.AddScopeByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDevicePrestages_RemoveScopeByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterRemoveScopeByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	req := &RequestRemoveScope{SerialNumbers: []string{"C02ABCDEFGH"}, VersionLock: 0}
	result, resp, err := svc.RemoveScopeByIDV2(context.Background(), "1", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.PrestageId)
}

func TestUnit_MobileDevicePrestages_RemoveScopeByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestRemoveScope{SerialNumbers: []string{"C02ABCDEFGH"}}

	result, resp, err := svc.RemoveScopeByIDV2(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_RemoveScopeByIDV2_NilRequest(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.RemoveScopeByIDV2(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

// --- Sync operations ---

func TestUnit_MobileDevicePrestages_GetAllSyncsV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetAllSyncsMock()

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetAllSyncsV2(context.Background())

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "1", result[0].PrestageId)
	assert.Equal(t, "COMPLETE", result[0].SyncState)
	assert.Equal(t, "2", result[1].PrestageId)
	assert.Equal(t, "PENDING", result[1].SyncState)
}

func TestUnit_MobileDevicePrestages_GetSyncsByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetSyncsByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetSyncsByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 2)
}

func TestUnit_MobileDevicePrestages_GetSyncsByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetSyncsByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_GetLatestSyncByIDV2_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetLatestSyncByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetLatestSyncByIDV2(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.PrestageId)
	assert.Equal(t, "COMPLETE", result.SyncState)
	assert.Equal(t, "2024-01-15T10:30:00Z", result.Timestamp)
}

func TestUnit_MobileDevicePrestages_GetLatestSyncByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetLatestSyncByIDV2(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

// --- Attachments ---

func TestUnit_MobileDevicePrestages_GetAttachmentsByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetAttachmentsByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetAttachmentsByIDV3(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "att-1", result[0].ID)
	assert.Equal(t, "profile.mobileconfig", result[0].Name)
	assert.Equal(t, "att-2", result[1].ID)
}

func TestUnit_MobileDevicePrestages_GetAttachmentsByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetAttachmentsByIDV3(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_UploadAttachmentV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterUploadAttachmentMock("1")

	svc := NewMobileDevicePrestages(mock)
	content := []byte("test mobileconfig content")
	reader := bytes.NewReader(content)
	result, resp, err := svc.UploadAttachmentV3(context.Background(), "1", reader, int64(len(content)), "test.mobileconfig")

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "att-3", result.ID)
	assert.Equal(t, "uploaded.mobileconfig", result.Name)
	assert.Equal(t, "com.apple.mobileconfig", result.FileType)
}

func TestUnit_MobileDevicePrestages_UploadAttachmentV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	reader := bytes.NewReader([]byte("content"))

	result, resp, err := svc.UploadAttachmentV3(context.Background(), "", reader, 7, "test.mobileconfig")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_UploadAttachmentV3_NilReader(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.UploadAttachmentV3(context.Background(), "1", nil, 0, "test.mobileconfig")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file reader is required")
}

func TestUnit_MobileDevicePrestages_UploadAttachmentV3_EmptyFileName(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	reader := bytes.NewReader([]byte("content"))

	result, resp, err := svc.UploadAttachmentV3(context.Background(), "1", reader, 7, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "file name is required")
}

func TestUnit_MobileDevicePrestages_UploadAttachmentFromFileV3_Success(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.mobileconfig")
	require.NoError(t, os.WriteFile(tmpFile, []byte("test content"), 0644))

	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterUploadAttachmentMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.UploadAttachmentFromFileV3(context.Background(), "1", tmpFile)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "att-3", result.ID)
}

func TestUnit_MobileDevicePrestages_UploadAttachmentFromFileV3_FileNotFound(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.UploadAttachmentFromFileV3(context.Background(), "1", "/nonexistent/path/file.mobileconfig")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "open attachment file")
}

func TestUnit_MobileDevicePrestages_DeleteAttachmentsByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterDeleteAttachmentsByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	req := &RequestDeleteAttachments{IDs: []string{"att-1", "att-2"}}
	resp, err := svc.DeleteAttachmentsByIDV3(context.Background(), "1", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestUnit_MobileDevicePrestages_DeleteAttachmentsByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestDeleteAttachments{IDs: []string{"att-1"}}

	resp, err := svc.DeleteAttachmentsByIDV3(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_DeleteAttachmentsByIDV3_NilRequest(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	resp, err := svc.DeleteAttachmentsByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// --- History ---

func TestUnit_MobileDevicePrestages_GetHistoryByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterGetHistoryByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.GetHistoryByIDV3(context.Background(), "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Initial setup", result.Results[0].Note)
}

func TestUnit_MobileDevicePrestages_GetHistoryByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.GetHistoryByIDV3(context.Background(), "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_AddHistoryNoteByIDV3_Success(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterAddHistoryNoteByIDMock("1")

	svc := NewMobileDevicePrestages(mock)
	req := &RequestAddHistoryNote{Note: "Test history note"}
	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", req)

	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, result)
	assert.Equal(t, "1", result.ID)
	assert.Contains(t, result.Href, "/api/v3/mobile-device-prestages/1/history")
}

func TestUnit_MobileDevicePrestages_AddHistoryNoteByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestAddHistoryNote{Note: "Test note"}

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevicePrestages_AddHistoryNoteByIDV3_NilRequest(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDevicePrestages_AddHistoryNoteByIDV3_EmptyNote(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	svc := NewMobileDevicePrestages(mock)
	req := &RequestAddHistoryNote{Note: ""}

	result, resp, err := svc.AddHistoryNoteByIDV3(context.Background(), "1", req)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "note is required")
}

// --- Scope version lock fetch errors ---

func TestUnit_MobileDevicePrestages_ReplaceScopeByIDV2_ScopeFetchError(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	mock.RegisterReplaceScopePutOnlyMock("999") // Only PUT, no GET - version lock fetch will fail

	svc := NewMobileDevicePrestages(mock)
	req := &RequestReplaceScope{SerialNumbers: []string{"C02NEW123"}}
	result, resp, err := svc.ReplaceScopeByIDV2(context.Background(), "999", req)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "failed to fetch current device scope for version locking")
}

func TestUnit_MobileDevicePrestages_DeleteByNameV3_GetByNameFails(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No list mock - GetByNameV3 will fail when resolving name to ID

	svc := NewMobileDevicePrestages(mock)
	resp, err := svc.DeleteByNameV3(context.Background(), "Nonexistent")

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get mobile device prestage by name")
}

func TestUnit_MobileDevicePrestages_CreateV3_APIFails(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No create mock registered - Post will fail

	svc := NewMobileDevicePrestages(mock)
	prestage := &ResourceMobileDevicePrestage{DisplayName: "Valid Name"}
	result, resp, err := svc.CreateV3(context.Background(), prestage)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no response")
}

func TestUnit_MobileDevicePrestages_ListV3_APIFails(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No list mock - GetPaginated will fail

	svc := NewMobileDevicePrestages(mock)
	result, resp, err := svc.ListV3(context.Background())

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "no response")
}

func TestUnit_MobileDevicePrestages_DeleteByIDV3_APIFails(t *testing.T) {
	mock := mocks.NewMobileDevicePrestagesMock()
	// No delete mock for this ID

	svc := NewMobileDevicePrestages(mock)
	resp, err := svc.DeleteByIDV3(context.Background(), "1")

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response")
}
