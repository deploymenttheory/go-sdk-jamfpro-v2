package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ActivationCodeMock struct {
	*mocks.GenericMock
}

func NewActivationCodeMock() *ActivationCodeMock {
	return &ActivationCodeMock{
		GenericMock: mocks.NewXMLMock("ActivationCodeMock"),
	}
}

func (m *ActivationCodeMock) RegisterMocks() {
	m.RegisterGetActivationCodeMock()
	m.RegisterUpdateActivationCodeMock()
}

func (m *ActivationCodeMock) RegisterGetActivationCodeMock() {
	m.Register("GET", "/JSSResource/activationcode", 200, "validate_get_activation_code.xml")
}

func (m *ActivationCodeMock) RegisterUpdateActivationCodeMock() {
	m.Register("PUT", "/JSSResource/activationcode", 200, "validate_update_activation_code.xml")
}

