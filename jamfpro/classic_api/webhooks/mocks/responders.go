package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type WebhooksMock struct {
	*mocks.GenericMock
}

func NewWebhooksMock() *WebhooksMock {
	return &WebhooksMock{
		GenericMock: mocks.NewXMLMock("WebhooksMock"),
	}
}

func (m *WebhooksMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *WebhooksMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *WebhooksMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/webhooks", 200, "validate_list_webhooks.xml")
}

func (m *WebhooksMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/webhooks/id/1", 200, "validate_get_webhook.xml")
}

func (m *WebhooksMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/webhooks/name/Computer Enrolled", 200, "validate_get_webhook.xml")
}

func (m *WebhooksMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/webhooks/id/0", 201, "validate_create_webhook.xml")
}

func (m *WebhooksMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/webhooks/id/1", 200, "validate_update_webhook.xml")
}

func (m *WebhooksMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/webhooks/name/Computer Enrolled", 200, "validate_update_webhook.xml")
}

func (m *WebhooksMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/webhooks/id/1", 200, "")
}

func (m *WebhooksMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/webhooks/name/Computer Enrolled", 200, "")
}

func (m *WebhooksMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/webhooks/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *WebhooksMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/webhooks/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A webhook with that name already exists")
}

