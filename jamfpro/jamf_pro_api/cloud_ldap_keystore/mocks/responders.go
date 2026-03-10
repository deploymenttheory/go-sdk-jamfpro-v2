package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudLdapKeystoreMock struct {
	*mocks.GenericMock
}

func NewCloudLdapKeystoreMock() *CloudLdapKeystoreMock {
	return &CloudLdapKeystoreMock{
		GenericMock: mocks.NewJSONMock("CloudLdapKeystoreMock"),
	}
}

func (m *CloudLdapKeystoreMock) RegisterValidateMock() {
	m.Register("POST", "/api/v1/ldap-keystore/verify", 200, "validate_keystore.json")
}
