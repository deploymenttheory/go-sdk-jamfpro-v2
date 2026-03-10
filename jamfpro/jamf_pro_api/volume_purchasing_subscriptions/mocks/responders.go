package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type VolumePurchasingSubscriptionsMock struct {
	*mocks.GenericMock
}

func NewVolumePurchasingSubscriptionsMock() *VolumePurchasingSubscriptionsMock {
	return &VolumePurchasingSubscriptionsMock{
		GenericMock: mocks.NewJSONMock("VolumePurchasingSubscriptionsMock"),
	}
}

func (m *VolumePurchasingSubscriptionsMock) RegisterListMock() {
	m.Register("GET", "/api/v1/volume-purchasing-subscriptions", 200, "validate_list.json")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/volume-purchasing-subscriptions/1", 200, "validate_get.json")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/volume-purchasing-subscriptions", 201, "validate_create.json")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/volume-purchasing-subscriptions/1", 200, "validate_update.json")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/volume-purchasing-subscriptions/1", 204, "")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-subscriptions/999", 404, "error_not_found.json", "")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterUpdateNotFoundErrorMock() {
	m.RegisterError("PUT", "/api/v1/volume-purchasing-subscriptions/999", 404, "error_not_found.json", "")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterDeleteNotFoundErrorMock() {
	m.RegisterError("DELETE", "/api/v1/volume-purchasing-subscriptions/999", 404, "error_not_found.json", "")
}

func (m *VolumePurchasingSubscriptionsMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/volume-purchasing-subscriptions", 500, "error_api.json", "")
}
