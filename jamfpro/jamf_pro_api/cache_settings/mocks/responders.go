package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CacheSettingsMock struct {
	*mocks.GenericMock
}

func NewCacheSettingsMock() *CacheSettingsMock {
	return &CacheSettingsMock{
		GenericMock: mocks.NewJSONMock("CacheSettingsMock"),
	}
}

func (m *CacheSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/cache-settings", 200, "validate_get.json")
}

func (m *CacheSettingsMock) RegisterPutMock() {
	m.Register("PUT", "/api/v1/cache-settings", 200, "validate_get.json")
}

func (m *CacheSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/cache-settings", 500, "error_internal.json", "")
}

func (m *CacheSettingsMock) RegisterPutErrorMock() {
	m.RegisterError("PUT", "/api/v1/cache-settings", 500, "error_internal.json", "")
}
