package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CloudIdpMock struct {
	*mocks.GenericMock
}

func NewCloudIdpMock() *CloudIdpMock {
	return &CloudIdpMock{
		GenericMock: mocks.NewJSONMock("CloudIdpMock"),
	}
}

func (m *CloudIdpMock) RegisterListMock() {
	m.Register("GET", "/api/v1/cloud-idp", 200, "validate_list.json")
}

func (m *CloudIdpMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/cloud-idp/"+id, 200, "validate_get.json")
}

func (m *CloudIdpMock) RegisterExportMock() {
	m.Register("POST", "/api/v1/cloud-idp/export", 200, "validate_list.json")
}

func (m *CloudIdpMock) RegisterGetHistoryByIDMock(id string) {
	m.Register("GET", "/api/v1/cloud-idp/"+id+"/history", 200, "validate_history.json")
}

func (m *CloudIdpMock) RegisterAddHistoryNoteMock(id string) {
	m.Register("POST", "/api/v1/cloud-idp/"+id+"/history", 201, "")
}

func (m *CloudIdpMock) RegisterTestGroupSearchMock(id string) {
	m.Register("POST", "/api/v1/cloud-idp/"+id+"/test-group", 200, "validate_test_group.json")
}

func (m *CloudIdpMock) RegisterTestUserSearchMock(id string) {
	m.Register("POST", "/api/v1/cloud-idp/"+id+"/test-user", 200, "validate_test_user.json")
}

func (m *CloudIdpMock) RegisterTestUserMembershipMock(id string) {
	m.Register("POST", "/api/v1/cloud-idp/"+id+"/test-user-membership", 200, "validate_test_membership.json")
}

func (m *CloudIdpMock) RegisterErrorMocks() {
	m.RegisterError("GET", "/api/v1/cloud-idp", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/cloud-idp/1", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-idp/export", 500, "error_internal.json", "")
	m.RegisterError("GET", "/api/v1/cloud-idp/1/history", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-idp/1/history", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-idp/1/test-group", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-idp/1/test-user", 500, "error_internal.json", "")
	m.RegisterError("POST", "/api/v1/cloud-idp/1/test-user-membership", 500, "error_internal.json", "")
}
