package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SsoCertificateMock struct {
	*mocks.GenericMock
}

func NewSsoCertificateMock() *SsoCertificateMock {
	return &SsoCertificateMock{
		GenericMock: mocks.NewJSONMock("SsoCertificateMock"),
	}
}

func (m *SsoCertificateMock) RegisterGetMock() {
	m.Register("GET", "/api/v2/sso/cert", 200, "validate_get.json")
}

func (m *SsoCertificateMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/sso/cert", 200, "validate_get.json")
}

func (m *SsoCertificateMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v2/sso/cert", 200, "")
}

func (m *SsoCertificateMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v2/sso/cert", 200, "validate_get.json")
}

func (m *SsoCertificateMock) RegisterDownloadMock() {
	m.Register("GET", "/api/v2/sso/cert/download", 200, "")
}

func (m *SsoCertificateMock) RegisterParseMock() {
	m.Register("POST", "/api/v2/sso/cert/parse", 200, "validate_parse.json")
}
