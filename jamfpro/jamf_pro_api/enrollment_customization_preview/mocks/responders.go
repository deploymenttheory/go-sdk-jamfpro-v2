package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type EnrollmentCustomizationPreviewMock struct {
	*mocks.GenericMock
}

func NewEnrollmentCustomizationPreviewMock() *EnrollmentCustomizationPreviewMock {
	return &EnrollmentCustomizationPreviewMock{
		GenericMock: mocks.NewJSONMock("EnrollmentCustomizationPreviewMock"),
	}
}

func (m *EnrollmentCustomizationPreviewMock) RegisterParseMarkdownMock() {
	m.Register("POST", "/api/v1/enrollment-customization/parse-markdown", 200, "validate_parse_markdown.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetAllPanelsMock(id string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/all", 200, "validate_panels_list.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetPanelByIDMock(id, panelID string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 200, "validate_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeletePanelMock(id, panelID string) {
	m.Register("DELETE", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateLdapPanelMock(id string) {
	m.Register("POST", "/api/v1/enrollment-customization/"+id+"/ldap", 201, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetLdapPanelMock(id, panelID string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 200, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateLdapPanelMock(id, panelID string) {
	m.Register("PUT", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 200, "validate_ldap_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteLdapPanelMock(id, panelID string) {
	m.Register("DELETE", "/api/v1/enrollment-customization/"+id+"/ldap/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateSsoPanelMock(id string) {
	m.Register("POST", "/api/v1/enrollment-customization/"+id+"/sso", 201, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetSsoPanelMock(id, panelID string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 200, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateSsoPanelMock(id, panelID string) {
	m.Register("PUT", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 200, "validate_sso_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteSsoPanelMock(id, panelID string) {
	m.Register("DELETE", "/api/v1/enrollment-customization/"+id+"/sso/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterCreateTextPanelMock(id string) {
	m.Register("POST", "/api/v1/enrollment-customization/"+id+"/text", 201, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetTextPanelMock(id, panelID string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 200, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterUpdateTextPanelMock(id, panelID string) {
	m.Register("PUT", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 200, "validate_text_panel.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeleteTextPanelMock(id, panelID string) {
	m.Register("DELETE", "/api/v1/enrollment-customization/"+id+"/text/"+panelID, 204, "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetTextPanelMarkdownMock(id, panelID string) {
	m.Register("GET", "/api/v1/enrollment-customization/"+id+"/text/"+panelID+"/markdown", 200, "validate_text_panel_markdown.json")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterParseMarkdownErrorMock() {
	m.RegisterError("POST", "/api/v1/enrollment-customization/parse-markdown", 400, "error_bad_request.json", "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetAllPanelsErrorMock(id string) {
	m.RegisterError("GET", "/api/v1/enrollment-customization/"+id+"/all", 404, "error_not_found.json", "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterGetPanelByIDErrorMock(id, panelID string) {
	m.RegisterError("GET", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 404, "error_not_found.json", "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterDeletePanelErrorMock(id, panelID string) {
	m.RegisterError("DELETE", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 404, "error_not_found.json", "")
}

func (m *EnrollmentCustomizationPreviewMock) RegisterNotFoundErrorMock(id, panelID string) {
	m.RegisterError("GET", "/api/v1/enrollment-customization/"+id+"/all/"+panelID, 404, "error_not_found.json", "")
}
