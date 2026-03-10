package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AppRequestMock struct {
	*mocks.GenericMock
}

func NewAppRequestMock() *AppRequestMock {
	return &AppRequestMock{
		GenericMock: mocks.NewJSONMock("AppRequestMock"),
	}
}

func (m *AppRequestMock) RegisterMocks() {
	m.RegisterListFormInputFieldsMock()
	m.RegisterReplaceFormInputFieldsMock()
	m.RegisterCreateFormInputFieldMock()
	m.RegisterGetFormInputFieldMock()
	m.RegisterUpdateFormInputFieldMock()
	m.RegisterDeleteFormInputFieldMock()
	m.RegisterGetSettingsMock()
	m.RegisterUpdateSettingsMock()
}

func (m *AppRequestMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v1/app-request/form-input-fields", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/app-request/form-input-fields", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/app-request/form-input-fields", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/app-request/form-input-fields/1", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/app-request/form-input-fields/1", 500, "error_internal.json", "")
	m.RegisterError("DELETE", "/api/v1/app-request/form-input-fields/1", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/app-request/settings", 500, "error_internal.json", "")
	m.RegisterError("PUT", "/api/v1/app-request/settings", 500, "error_internal.json", "")
	m.RegisterNotFoundErrorMock()
}

func (m *AppRequestMock) RegisterListFormInputFieldsMock() {
	m.Register("GET", "/api/v1/app-request/form-input-fields", 200, "validate_list_form_input_fields.json")
}

func (m *AppRequestMock) RegisterReplaceFormInputFieldsMock() {
	m.Register("PUT", "/api/v1/app-request/form-input-fields", 200, "validate_replace_form_input_fields.json")
}

func (m *AppRequestMock) RegisterCreateFormInputFieldMock() {
	m.Register("POST", "/api/v1/app-request/form-input-fields", 201, "validate_create_form_input_field.json")
}

func (m *AppRequestMock) RegisterGetFormInputFieldMock() {
	m.Register("GET", "/api/v1/app-request/form-input-fields/1", 200, "validate_get_form_input_field.json")
}

func (m *AppRequestMock) RegisterUpdateFormInputFieldMock() {
	m.Register("PUT", "/api/v1/app-request/form-input-fields/1", 200, "validate_get_form_input_field.json")
}

func (m *AppRequestMock) RegisterDeleteFormInputFieldMock() {
	m.Register("DELETE", "/api/v1/app-request/form-input-fields/1", 204, "")
}

func (m *AppRequestMock) RegisterGetSettingsMock() {
	m.Register("GET", "/api/v1/app-request/settings", 200, "validate_get_settings.json")
}

func (m *AppRequestMock) RegisterUpdateSettingsMock() {
	m.Register("PUT", "/api/v1/app-request/settings", 200, "validate_update_settings.json")
}

func (m *AppRequestMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/app-request/form-input-fields/999", 404, "error_not_found.json", "")
}
