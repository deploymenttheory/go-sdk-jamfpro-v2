package static_user_groups_test

import (
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/static_user_groups"
	"github.com/stretchr/testify/assert"
)

func TestValidateRequest_NilRequest(t *testing.T) {
	err := static_user_groups.ValidateRequest(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestValidateRequest_EmptyName(t *testing.T) {
	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "",
		IsSmart: false,
	}
	err := static_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestValidateRequest_IsSmartTrue(t *testing.T) {
	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "Test Group",
		IsSmart: true,
	}
	err := static_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "static user group must have IsSmart=false")
}

func TestValidateRequest_Valid(t *testing.T) {
	req := &static_user_groups.RequestStaticUserGroup{
		Name:    "Test Group",
		IsSmart: false,
	}
	err := static_user_groups.ValidateRequest(req)
	assert.NoError(t, err)
}
