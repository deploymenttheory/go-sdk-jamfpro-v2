package enrollment

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.EnrollmentMock) {
	t.Helper()
	mock := mocks.NewEnrollmentMock()
	return NewService(mock), mock
}

// Tests for V2 API - History
func TestUnitGetHistoryV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetHistoryV2Mock()

	result, resp, err := svc.GetHistoryV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
}

// Tests for V3 API - Access Groups
func TestUnitListAccessGroupsV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAccessGroupsV3Mock()

	result, resp, err := svc.ListAccessGroupsV3(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Test Access Group", result.Results[0].Name)
}

func TestUnitGetAccessGroupByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAccessGroupByIDV3Mock()

	result, resp, err := svc.GetAccessGroupByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test Access Group", result.Name)
	assert.True(t, result.EnterpriseEnrollmentEnabled)
}

func TestUnitGetAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAccessGroupByIDV3(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

func TestUnitCreateAccessGroupV3_Success(t *testing.T) {
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

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "2", result.ID)
}

func TestUnitCreateAccessGroupV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateAccessGroupV3(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitUpdateAccessGroupByIDV3_Success(t *testing.T) {
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

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Updated Access Group", result.Name)
	assert.True(t, result.PersonalEnrollmentEnabled)
	assert.True(t, result.RequireEula)
}

func TestUnitUpdateAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	request := &ResourceAccountDrivenUserEnrollmentAccessGroup{Name: "Test"}
	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "", request)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

func TestUnitUpdateAccessGroupByIDV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateAccessGroupByIDV3(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteAccessGroupByIDV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAccessGroupByIDV3Mock()

	resp, err := svc.DeleteAccessGroupByIDV3(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteAccessGroupByIDV3_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAccessGroupByIDV3(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "access group ID is required")
}

// Tests for V3 API - Language Messages
func TestUnitListLanguageMessagesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageMessagesV3Mock()

	result, resp, err := svc.ListLanguageMessagesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 1)
	assert.Equal(t, "en", result[0].LanguageCode)
	assert.Equal(t, "English", result[0].Name)
}

func TestUnitListLanguageCodesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()

	result, resp, err := svc.ListLanguageCodesV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result, 3)
	assert.Equal(t, "en", result[0].Value)
	assert.Equal(t, "English", result[0].Name)
}

func TestUnitGetLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterGetLanguageMessageV3Mock()

	result, resp, err := svc.GetLanguageMessageV3(context.Background(), "en")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "en", result.LanguageCode)
	assert.Equal(t, "English", result.Name)
	assert.Equal(t, "Enrollment", result.Title)
}

func TestUnitUpdateLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterUpdateLanguageMessageV3Mock()

	request := &ResourceEnrollmentLanguage{
		LanguageCode:      "en",
		Name:              "English",
		Title:             "Updated Enrollment",
		LoginDescription:  "Updated description",
		Username:          "Username",
		Password:          "Password",
		LoginButton:       "Login",
	}

	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", request)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "Updated Enrollment", result.Title)
}

func TestUnitUpdateLanguageMessageV3_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateLanguageMessageV3(context.Background(), "en", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteLanguageMessageV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterDeleteLanguageMessageV3Mock()

	resp, err := svc.DeleteLanguageMessageV3(context.Background(), "en")
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteMultipleLanguageMessagesV3_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListLanguageCodesV3Mock()
	mock.RegisterDeleteMultipleLanguageMessagesV3Mock()

	request := &RequestDeleteMultipleLanguages{
		IDs: []string{"en", "es"},
	}

	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteMultipleLanguageMessagesV3_EmptyRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteMultipleLanguageMessagesV3(context.Background(), &RequestDeleteMultipleLanguages{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "at least one language code is required")
}

// Tests for V4 API - Enrollment Settings
func TestUnitGetV4_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetV4Mock()

	result, resp, err := svc.GetV4(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.InstallSingleProfile)
	assert.True(t, result.MacOsEnterpriseEnrollmentEnabled)
	assert.Equal(t, "admin", result.ManagementUsername)
}

func TestUnitUpdateV4_Success(t *testing.T) {
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

	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.InstallSingleProfile)
	assert.True(t, result.RestrictReenrollment)
}

func TestUnitUpdateV4_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV4(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// Tests for Validation
func TestUnitValidateLanguageCode_Valid(t *testing.T) {
	codes := []ResourceLanguageCode{
		{Value: "en", Name: "English"},
		{Value: "es", Name: "Spanish"},
	}

	err := ValidateLanguageCode("en", codes)
	require.NoError(t, err)
}

func TestUnitValidateLanguageCode_Invalid(t *testing.T) {
	codes := []ResourceLanguageCode{
		{Value: "en", Name: "English"},
	}

	err := ValidateLanguageCode("fr", codes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid language code")
}

func TestUnitValidateLanguageCode_Empty(t *testing.T) {
	codes := []ResourceLanguageCode{{Value: "en", Name: "English"}}

	err := ValidateLanguageCode("", codes)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required")
}
