package enrollment_customization_preview

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/enrollment_customization_preview/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testID = "1"
const testPanelID = "1"

func setupMockService(t *testing.T) (*EnrollmentCustomizationPreview, *mocks.EnrollmentCustomizationPreviewMock) {
	t.Helper()
	mock := mocks.NewEnrollmentCustomizationPreviewMock()
	return NewEnrollmentCustomizationPreview(mock), mock
}

func TestUnit_EnrollmentCustomizationPreview_NewService(t *testing.T) {
	svc := NewEnrollmentCustomizationPreview(nil)
	require.NotNil(t, svc)
}

func TestUnit_EnrollmentCustomizationPreview_ParseMarkdown_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterParseMarkdownMock()

	req := &RequestParseMarkdown{Markdown: "# Hello"}
	result, resp, err := svc.ParseMarkdown(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, result.Markdown, "<p>")
	assert.Contains(t, result.Markdown, "Hello")
}

func TestUnit_EnrollmentCustomizationPreview_ParseMarkdown_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ParseMarkdown(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_ParseMarkdown_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.ParseMarkdown(context.Background(), &RequestParseMarkdown{Markdown: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_GetAllPanels_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetAllPanelsMock(testID)

	result, resp, err := svc.GetAllPanels(context.Background(), testID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result.Panels, 3)
	assert.Equal(t, "LDAP Panel", result.Panels[0].DisplayName)
	assert.Equal(t, 1, result.Panels[0].ID)
	assert.Equal(t, "LDAP", result.Panels[0].Type)
}

func TestUnit_EnrollmentCustomizationPreview_GetAllPanels_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAllPanels(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetAllPanels_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetAllPanels(context.Background(), testID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_GetPanelByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPanelByIDMock(testID, testPanelID)

	result, resp, err := svc.GetPanelByID(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "LDAP Panel", result.DisplayName)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "LDAP", result.Type)
}

func TestUnit_EnrollmentCustomizationPreview_GetPanelByID_EmptyEnrollmentID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPanelByID(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetPanelByID_EmptyPanelID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPanelByID(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetPanelByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock(testID, "999")

	result, resp, err := svc.GetPanelByID(context.Background(), testID, "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_DeletePanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePanelMock(testID, testPanelID)

	resp, err := svc.DeletePanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_DeletePanel_EmptyEnrollmentID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_DeletePanel_EmptyPanelID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateLdapPanelMock(testID)

	req := &ResourceLdapPanel{
		DisplayName: "LDAP Panel",
		Title:       "LDAP Authentication",
	}
	result, resp, err := svc.CreateLdapPanel(context.Background(), testID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "LDAP Panel", result.DisplayName)
	assert.Equal(t, "LDAP Authentication", result.Title)
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "x", Title: "y"}
	result, resp, err := svc.CreateLdapPanel(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateLdapPanel(context.Background(), testID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "", Title: "y"}
	result, resp, err := svc.CreateLdapPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_EmptyTitle(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "x", Title: ""}
	result, resp, err := svc.CreateLdapPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "title is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetLdapPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetLdapPanelMock(testID, testPanelID)

	result, resp, err := svc.GetLdapPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "LDAP Panel", result.DisplayName)
}

func TestUnit_EnrollmentCustomizationPreview_GetLdapPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetLdapPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.GetLdapPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateLdapPanelMock(testID, testPanelID)

	req := &ResourceLdapPanel{
		DisplayName: "LDAP Panel Updated",
		Title:       "LDAP Auth",
	}
	result, resp, err := svc.UpdateLdapPanel(context.Background(), testID, testPanelID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &ResourceLdapPanel{DisplayName: "x", Title: "y"}

	_, _, err := svc.UpdateLdapPanel(context.Background(), "", testPanelID, req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.UpdateLdapPanel(context.Background(), testID, "", req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateLdapPanel(context.Background(), testID, testPanelID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "", Title: "y"}
	result, resp, err := svc.UpdateLdapPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_EmptyTitle(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "x", Title: ""}
	result, resp, err := svc.UpdateLdapPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "title is required")
}

func TestUnit_EnrollmentCustomizationPreview_DeleteLdapPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteLdapPanelMock(testID, testPanelID)

	resp, err := svc.DeleteLdapPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_DeleteLdapPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLdapPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	resp, err = svc.DeleteLdapPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateSsoPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateSsoPanelMock(testID)

	req := &ResourceSsoPanel{DisplayName: "SSO Panel"}
	result, resp, err := svc.CreateSsoPanel(context.Background(), testID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "SSO Panel", result.DisplayName)
}

func TestUnit_EnrollmentCustomizationPreview_CreateSsoPanel_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSsoPanel{DisplayName: "x"}
	result, resp, err := svc.CreateSsoPanel(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateSsoPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateSsoPanel(context.Background(), testID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateSsoPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSsoPanel{DisplayName: ""}
	result, resp, err := svc.CreateSsoPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetSsoPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetSsoPanelMock(testID, testPanelID)

	result, resp, err := svc.GetSsoPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "SSO Panel", result.DisplayName)
}

func TestUnit_EnrollmentCustomizationPreview_UpdateSsoPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateSsoPanelMock(testID, testPanelID)

	req := &ResourceSsoPanel{DisplayName: "SSO Panel Updated"}
	result, resp, err := svc.UpdateSsoPanel(context.Background(), testID, testPanelID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_UpdateSsoPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateSsoPanel(context.Background(), testID, testPanelID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateSsoPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSsoPanel{DisplayName: ""}
	result, resp, err := svc.UpdateSsoPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateSsoPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &ResourceSsoPanel{DisplayName: "x"}

	_, _, err := svc.UpdateSsoPanel(context.Background(), "", testPanelID, req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.UpdateSsoPanel(context.Background(), testID, "", req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetSsoPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetSsoPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.GetSsoPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetSsoPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetSsoPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_UpdateSsoPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSsoPanel{DisplayName: "x"}
	result, resp, err := svc.UpdateSsoPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeleteSsoPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteSsoPanelMock(testID, testPanelID)

	resp, err := svc.DeleteSsoPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_DeleteSsoPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSsoPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateTextPanelMock(testID)

	req := &ResourceTextPanel{
		DisplayName: "Text Panel",
		Title:       "Welcome",
		Body:        "Welcome to enrollment",
	}
	result, resp, err := svc.CreateTextPanel(context.Background(), testID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "Text Panel", result.DisplayName)
	assert.Equal(t, "Welcome to enrollment", result.Body)
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: "z"}
	result, resp, err := svc.CreateTextPanel(context.Background(), "", req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateTextPanel(context.Background(), testID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "", Title: "y", Body: "z"}
	result, resp, err := svc.CreateTextPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_EmptyTitle(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "", Body: "z"}
	result, resp, err := svc.CreateTextPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "title is required")
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_EmptyBody(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: ""}
	result, resp, err := svc.CreateTextPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "body is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetTextPanelMock(testID, testPanelID)

	result, resp, err := svc.GetTextPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "Text Panel", result.DisplayName)
	assert.Equal(t, "Welcome to enrollment", result.Body)
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateTextPanelMock(testID, testPanelID)

	req := &ResourceTextPanel{
		DisplayName: "Text Panel Updated",
		Title:       "Updated",
		Body:        "Updated body",
	}
	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_EmptyBody(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: ""}
	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "body is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_EmptyDisplayName(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "", Title: "y", Body: "z"}
	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "display name is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_EmptyTitle(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "", Body: "z"}
	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "title is required")
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: "z"}

	_, _, err := svc.UpdateTextPanel(context.Background(), "", testPanelID, req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.UpdateTextPanel(context.Background(), testID, "", req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetTextPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	_, _, err = svc.GetTextPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetTextPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_UpdateTextPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: "z"}
	result, resp, err := svc.UpdateTextPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeletePanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeleteSsoPanel_EmptyPanelID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSsoPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_DeleteSsoPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteSsoPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_GetLdapPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetLdapPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_CreateLdapPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "x", Title: "y"}
	result, resp, err := svc.CreateLdapPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_UpdateLdapPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceLdapPanel{DisplayName: "x", Title: "y"}
	result, resp, err := svc.UpdateLdapPanel(context.Background(), testID, testPanelID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeleteLdapPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteLdapPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_CreateSsoPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceSsoPanel{DisplayName: "x"}
	result, resp, err := svc.CreateSsoPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_CreateTextPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &ResourceTextPanel{DisplayName: "x", Title: "y", Body: "z"}
	result, resp, err := svc.CreateTextPanel(context.Background(), testID, req)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeleteTextPanel_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteTextPanel(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_EnrollmentCustomizationPreview_DeleteTextPanel_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteTextPanelMock(testID, testPanelID)

	resp, err := svc.DeleteTextPanel(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_EnrollmentCustomizationPreview_DeleteTextPanel_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteTextPanel(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	resp, err = svc.DeleteTextPanel(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanelMarkdown_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetTextPanelMarkdownMock(testID, testPanelID)

	result, resp, err := svc.GetTextPanelMarkdown(context.Background(), testID, testPanelID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, result.Markdown, "Welcome")
	assert.Contains(t, result.Markdown, "enrollment guide")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanelMarkdown_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetTextPanelMarkdown(context.Background(), "", testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "enrollment customization ID is required")

	result, resp, err = svc.GetTextPanelMarkdown(context.Background(), testID, "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "panel ID is required")
}

func TestUnit_EnrollmentCustomizationPreview_GetTextPanelMarkdown_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetTextPanelMarkdown(context.Background(), testID, testPanelID)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
