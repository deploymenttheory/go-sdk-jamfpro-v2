package apilifecycle_test

import (
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Parse_Valid(t *testing.T) {
	cases := []struct {
		in   string
		want apilifecycle.Version
	}{
		{"11.28.0", apilifecycle.Version{Major: 11, Minor: 28, Patch: 0}},
		{"11.29.1", apilifecycle.Version{Major: 11, Minor: 29, Patch: 1}},
		{"11.28", apilifecycle.Version{Major: 11, Minor: 28, Patch: 0}},
		{"11", apilifecycle.Version{Major: 11, Minor: 0, Patch: 0}},
		{"  11.27.2  ", apilifecycle.Version{Major: 11, Minor: 27, Patch: 2}},
		{"11.28.0-t1776264729651", apilifecycle.Version{Major: 11, Minor: 28, Patch: 0}},
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			got, err := apilifecycle.Parse(c.in)
			require.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func TestUnit_Parse_Empty(t *testing.T) {
	_, err := apilifecycle.Parse("   ")
	require.Error(t, err)
}

func TestUnit_Version_CompareAndAtLeast(t *testing.T) {
	v1127 := apilifecycle.MustParse("11.27.0")
	v1128 := apilifecycle.MustParse("11.28.0")
	v11281 := apilifecycle.MustParse("11.28.1")

	assert.Equal(t, -1, v1127.Compare(v1128))
	assert.Equal(t, 1, v1128.Compare(v1127))
	assert.Equal(t, 0, v1128.Compare(v1128))
	assert.Equal(t, -1, v1128.Compare(v11281))

	assert.False(t, v1127.AtLeast(v1128))
	assert.True(t, v1128.AtLeast(v1128))
	assert.True(t, v11281.AtLeast(v1128))
}

func TestUnit_Version_String(t *testing.T) {
	assert.Equal(t, "11.28.0", apilifecycle.MustParse("11.28").String())
}

func TestUnit_MustParse_PanicsOnEmpty(t *testing.T) {
	assert.Panics(t, func() { apilifecycle.MustParse("") })
}
