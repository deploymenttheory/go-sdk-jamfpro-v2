package smart_user_groups_test

import (
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/smart_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
	"github.com/stretchr/testify/assert"
)

func TestValidateRequest_NilRequest(t *testing.T) {
	err := smart_user_groups.ValidateRequest(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestValidateRequest_EmptyName(t *testing.T) {
	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []models.SharedSubsetCriteria{
				{Name: "Email Address", SearchType: "like", Value: "@example.com"},
			},
		},
	}
	err := smart_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user group name is required")
}

func TestValidateRequest_IsSmartFalse(t *testing.T) {
	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Test Group",
		IsSmart: false,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []models.SharedSubsetCriteria{
				{Name: "Email Address", SearchType: "like", Value: "@example.com"},
			},
		},
	}
	err := smart_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smart user group must have IsSmart=true")
}

func TestValidateRequest_NoCriteria(t *testing.T) {
	req := &smart_user_groups.RequestSmartUserGroup{
		Name:     "Test Group",
		IsSmart:  true,
		Criteria: nil,
	}
	err := smart_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smart user group must have criteria defined")
}

func TestValidateRequest_EmptyCriteria(t *testing.T) {
	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Test Group",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size:      0,
			Criterion: []models.SharedSubsetCriteria{},
		},
	}
	err := smart_user_groups.ValidateRequest(req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "smart user group must have criteria defined")
}

func TestValidateRequest_Valid(t *testing.T) {
	req := &smart_user_groups.RequestSmartUserGroup{
		Name:    "Test Group",
		IsSmart: true,
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []models.SharedSubsetCriteria{
				{Name: "Email Address", SearchType: "like", Value: "@example.com"},
			},
		},
	}
	err := smart_user_groups.ValidateRequest(req)
	assert.NoError(t, err)
}
