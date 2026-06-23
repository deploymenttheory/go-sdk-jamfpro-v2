package smartgroupvalidation_test

import (
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ValidateAndOr(t *testing.T) {
	require.NoError(t, smartgroupvalidation.ValidateAndOr("and", "or", "AND", "Or", ""))
	require.Error(t, smartgroupvalidation.ValidateAndOr("and", "nand"))
	require.Error(t, smartgroupvalidation.ValidateAndOr("xor"))
}

func TestUnit_ValidateGroupName(t *testing.T) {
	require.NoError(t, smartgroupvalidation.ValidateGroupName("ok"))
	require.Error(t, smartgroupvalidation.ValidateGroupName(""))
	require.NoError(t, smartgroupvalidation.ValidateGroupName(strings.Repeat("a", 255)))
	require.Error(t, smartgroupvalidation.ValidateGroupName(strings.Repeat("a", 256)))
}

func TestUnit_DedupeStrings(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, smartgroupvalidation.DedupeStrings([]string{"1", "2", "2", "3", "1"}))
	assert.Nil(t, smartgroupvalidation.DedupeStrings(nil))
	assert.Equal(t, []string{}, smartgroupvalidation.DedupeStrings([]string{}))
}
