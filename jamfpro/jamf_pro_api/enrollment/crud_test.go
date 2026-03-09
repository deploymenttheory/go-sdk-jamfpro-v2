package enrollment

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Enrollment, *mocks.EnrollmentMock) {
	t.Helper()
	mock := mocks.NewEnrollmentMock()
	return NewEnrollment(mock), mock
}

// Tests for V1 API - ADUE Session Token Settings
func TestUnit_Enrollment_GetADUESessionTokenSettingsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetADUESessionTokenSettingsV1Mock()

	result, resp, err := svc.GetADUESessionTokenSettingsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.False(t, result.Enabled)
	assert.Equal(t, 1, result.ExpirationIntervalDays)
	assert.Equal(t, 86400, result.ExpirationIntervalSeconds)
}

func TestUnit_Enrollment_UpdateADUESessionTokenSettingsV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateADUESessionTokenSettingsV1Mock()

	request := &ResourceADUESessionTokenSettings{
		Enabled:                   true,
		ExpirationIntervalDays:    1,
		ExpirationIntervalSeconds: 86400,
	}

	result, resp, err := svc.UpdateADUESessionTokenSettingsV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.Enabled)
}

func TestUnit_Enrollment_UpdateADUESessionTokenSettingsV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateADUESessionTokenSettingsV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Tests for V2 API - History
func TestUnit_Enrollment_GetHistoryV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryV2Mock()

	result, resp, err := svc.GetHistoryV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

func TestUnit_Enrollment_AddHistoryNotesV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNotesV2Mock()

	request := &RequestAddHistoryNotes{Note: "Test note"}

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.NotEmpty(t, result.Href)
}

func TestUnit_Enrollment_AddHistoryNotesV2_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Enrollment_AddHistoryNotesV2_EmptyNote(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.AddHistoryNotesV2(context.Background(), &RequestAddHistoryNotes{Note: ""})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "note is required")
}

func TestUnit_Enrollment_ExportHistoryV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryV2Mock()

	data, resp, err := svc.ExportHistoryV2(context.Background(), mime.TextCSV, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(data), "Username")
	assert.Contains(t, string(data), "admin")
}

func TestUnit_Enrollment_ExportHistoryV2_WithRequest(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportHistoryV2Mock()

	page, pageSize := 0, 100
	request := &RequestExportHistory{
		Page:     &page,
		PageSize: &pageSize,
		Sort:     []string{"id:desc"},
	}

	data, resp, err := svc.ExportHistoryV2(context.Background(), mime.ApplicationJSON, nil, request)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Enrollment_ListFilteredLanguageCodesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListFilteredLanguageCodesV3Mock()

	result, resp, err := svc.ListFilteredLanguageCodesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, "de", result[0].Value)
	assert.Equal(t, "German", result[0].Name)
	assert.Equal(t, "ja", result[1].Value)
	assert.Equal(t, "Japanese", result[1].Name)
}

// Tests for V3 API - Access Groups
func TestUnit_Enrollment_ListAccessGroupsV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccessGroupsV3Mock()

	result, resp, err := svc.ListAccessGroupsV3(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Test Access Group", result.Results[0].Name)
}

func TestUnit_Enrollment_GetAccessGroupByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccessGroupByIDV3Mock()

	result, resp, err := svc.GetAccessGroupByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Access Group", result.Name)
	assert.True(t, result.EnterpriseEnrollmentEnabled)
}

func TestUnit_Enrollment_GetAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccessGroupByIDV3(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

func TestUnit_Enrollment_CreateAccessGroupV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateAccessGroupV3Mock()

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{
		GroupID:                            "200",
		LdapServerID:                       "1",
		Name:                               "New Access Group",
		SiteID:                             "1",
		EnterpriseEnrollmentEnabled:        true,
		PersonalEnrollmentEnabled:          false,
		AccountDrivenUserEnrollmentEnabled: true,
		RequireEula:                        false,
	}

	result, resp, err := svc.CreateAccessGroupV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "2", result.ID)
}

func TestUnit_Enrollment_CreateAccessGroupV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAccessGroupV3(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Enrollment_UpdateAccessGroupByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateAccessGroupByIDV3Mock()

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{
		Name:                      "Updated Access Group",
		PersonalEnrollmentEnabled: true,
		RequireEula:               true,
	}

	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "1", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Updated Access Group", result.Name)
	assert.True(t, result.PersonalEnrollmentEnabled)
	assert.True(t, result.RequireEula)
}

func TestUnit_Enrollment_UpdateAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{Name: "Test"}
	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "", request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

func TestUnit_Enrollment_UpdateAccessGroupByIDV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Enrollment_DeleteAccessGroupByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccessGroupByIDV3Mock()

	resp, err := svc.DeleteAccessGroupByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Enrollment_DeleteAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccessGroupByIDV3(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

// Tests for V3 API - Language Messages
func TestUnit_Enrollment_ListLanguageMessagesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageMessagesV3Mock()

	result, resp, err := svc.ListLanguageMessagesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "en", result.Results[0].LanguageCode)
	assert.Equal(t, "English", result.Results[0].Name)
}

func TestUnit_Enrollment_ListLanguageCodesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	result, resp, err := svc.ListLanguageCodesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 3)
	assert.Equal(t, "en", result[0].Value)
	assert.Equal(t, "English", result[0].Name)
}

func TestUnit_Enrollment_GetLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterGetLanguageMessageV3Mock()

	result, resp, err := svc.GetLanguageMessageV3(context.Background(), "en")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "en", result.LanguageCode)
	assert.Equal(t, "English", result.Name)
	assert.Equal(t, "Enrollment", result.Title)
}

func TestUnit_Enrollment_UpdateLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterUpdateLanguageMessageV3Mock()

	request := &ResourceEnrollmentLanguage{
		LanguageCode:     "en",
		Name:             "English",
		Title:            "Updated Enrollment",
		LoginDescription: "Updated description",
		Username:         "Username",
		Password:         "Password",
		LoginButton:      "Login",
	}

	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Updated Enrollment", result.Title)
}

func TestUnit_Enrollment_UpdateLanguageMessageV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Enrollment_DeleteLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterDeleteLanguageMessageV3Mock()

	resp, err := svc.DeleteLanguageMessageV3(context.Background(), "en")
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Enrollment_DeleteMultipleLanguageMessagesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterDeleteMultipleLanguageMessagesV3Mock()

	request := &RequestDeleteMultipleLanguages{
		IDs: []string{"en", "es"},
	}

	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Enrollment_DeleteMultipleLanguageMessagesV3_EmptyRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), &RequestDeleteMultipleLanguages{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "at least one language code is required")
}

// Tests for V4 API - Enrollment Settings
func TestUnit_Enrollment_GetV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV4Mock()

	result, resp, err := svc.GetV4(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.InstallSingleProfile)
	assert.True(t, result.MacOsEnterpriseEnrollmentEnabled)
	assert.Equal(t, "admin", result.ManagementUsername)
}

func TestUnit_Enrollment_UpdateV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateV4Mock()

	request := &ResourceEnrollment{
		InstallSingleProfile:             false,
		RestrictReenrollment:             true,
		FlushLocationInformation:         true,
		MacOsEnterpriseEnrollmentEnabled: true,
		ManagementUsername:               "admin",
		CreateManagementAccount:          true,
	}

	result, resp, err := svc.UpdateV4(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.False(t, result.InstallSingleProfile)
	assert.True(t, result.RestrictReenrollment)
}

func TestUnit_Enrollment_UpdateV4_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV4(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Tests for Validation
func TestUnit_Enrollment_ValidateLanguageCode_Valid(t *testing.T) {
	codes := []ResourceLanguageCode{
		{Value: "en", Name: "English"},
		{Value: "es", Name: "Spanish"},
	}

	err := ValidateLanguageCode("en", codes)
	require.NoError(t, err)
}

func TestUnit_Enrollment_ValidateLanguageCode_Invalid(t *testing.T) {
	codes := []ResourceLanguageCode{
		{Value: "en", Name: "English"},
	}

	err := ValidateLanguageCode("fr", codes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnit_Enrollment_ValidateLanguageCode_Empty(t *testing.T) {
	codes := []ResourceLanguageCode{{Value: "en", Name: "English"}}

	err := ValidateLanguageCode("", codes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required")
}

func TestUnit_Enrollment_GetADUESessionTokenSettingsV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetADUESessionTokenSettingsV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_UpdateADUESessionTokenSettingsV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceADUESessionTokenSettings{Enabled: true}
	result, resp, err := svc.UpdateADUESessionTokenSettingsV1(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_AddHistoryNotesV2_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &RequestAddHistoryNotes{Note: "Test"}
	result, resp, err := svc.AddHistoryNotesV2(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_ExportHistoryV2_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.ExportHistoryV2(context.Background(), mime.ApplicationJSON, nil, nil)
	require.Error(t, err)
	assert.Nil(t, data)
	_ = resp
}

func TestUnit_Enrollment_ListFilteredLanguageCodesV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListFilteredLanguageCodesV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_GetHistoryV2_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetHistoryV2(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_ListAccessGroupsV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListAccessGroupsV3(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_GetAccessGroupByIDV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccessGroupByIDV3(context.Background(), "2")
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_CreateAccessGroupV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{Name: "Test"}
	result, resp, err := svc.CreateAccessGroupV3(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_UpdateAccessGroupByIDV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{Name: "Test"}
	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "2", request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_DeleteAccessGroupByIDV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccessGroupByIDV3(context.Background(), "2")
	require.Error(t, err)
	_ = resp
}

func TestUnit_Enrollment_ListLanguageMessagesV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListLanguageMessagesV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_GetLanguageMessageV3_ListCodesError(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetLanguageMessageV3(context.Background(), "en")
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_GetLanguageMessageV3_InvalidCode(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	result, resp, err := svc.GetLanguageMessageV3(context.Background(), "xx")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnit_Enrollment_GetLanguageMessageV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	result, resp, err := svc.GetLanguageMessageV3(context.Background(), "en")
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_UpdateLanguageMessageV3_ListCodesError(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceEnrollmentLanguage{LanguageCode: "en", Name: "English"}
	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_UpdateLanguageMessageV3_InvalidCode(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	request := &ResourceEnrollmentLanguage{LanguageCode: "xx", Name: "Unknown"}
	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "xx", request)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnit_Enrollment_UpdateLanguageMessageV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	request := &ResourceEnrollmentLanguage{LanguageCode: "en", Name: "English"}
	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_DeleteLanguageMessageV3_ListCodesError(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLanguageMessageV3(context.Background(), "en")
	require.Error(t, err)
	_ = resp
}

func TestUnit_Enrollment_DeleteLanguageMessageV3_InvalidCode(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	resp, err := svc.DeleteLanguageMessageV3(context.Background(), "xx")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnit_Enrollment_DeleteLanguageMessageV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	resp, err := svc.DeleteLanguageMessageV3(context.Background(), "en")
	require.Error(t, err)
	_ = resp
}

func TestUnit_Enrollment_DeleteMultipleLanguageMessagesV3_ListCodesError(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &RequestDeleteMultipleLanguages{IDs: []string{"en"}}
	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), request)
	require.Error(t, err)
	_ = resp
}

func TestUnit_Enrollment_DeleteMultipleLanguageMessagesV3_InvalidCode(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	request := &RequestDeleteMultipleLanguages{IDs: []string{"xx"}}
	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnit_Enrollment_DeleteMultipleLanguageMessagesV3_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	request := &RequestDeleteMultipleLanguages{IDs: []string{"en", "es"}}
	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), request)
	require.Error(t, err)
	_ = resp
}

func TestUnit_Enrollment_ListLanguageCodesV3_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ListLanguageCodesV3(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_GetV4_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV4(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_Enrollment_UpdateV4_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceEnrollment{ManagementUsername: "admin"}
	result, resp, err := svc.UpdateV4(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}
