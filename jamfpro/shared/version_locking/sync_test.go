package version_locking

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Mirrors the shape of a prestage: locks at the top level, on embedded
// subsets, on an optional pointer subset, and inside a slice.
type subset struct {
	Name        string
	VersionLock int
}

type optionalSubset struct {
	Enabled     bool
	VersionLock int
}

type resource struct {
	DisplayName string
	VersionLock int
	Location    subset
	Purchasing  subset
	Account     *optionalSubset
	Items       []subset
	unexported  int //nolint:unused // present to prove reflection skips it
}

func TestUnit_SyncAll_CopiesEveryLockInTree(t *testing.T) {
	current := &resource{
		DisplayName: "server",
		VersionLock: 7,
		Location:    subset{Name: "loc", VersionLock: 3},
		Purchasing:  subset{Name: "pur", VersionLock: 4},
		Account:     &optionalSubset{Enabled: true, VersionLock: 5},
		Items:       []subset{{Name: "a", VersionLock: 11}, {Name: "b", VersionLock: 12}},
	}
	request := &resource{
		DisplayName: "client",
		VersionLock: 0,
		Location:    subset{Name: "loc-new", VersionLock: 999},
		Purchasing:  subset{Name: "pur-new", VersionLock: 999},
		Account:     &optionalSubset{Enabled: false, VersionLock: 999},
		Items:       []subset{{Name: "a2", VersionLock: 999}, {Name: "b2", VersionLock: 999}},
	}

	SyncAll(current, request)

	assert.Equal(t, 7, request.VersionLock, "top-level lock")
	assert.Equal(t, 3, request.Location.VersionLock, "embedded subset lock")
	assert.Equal(t, 4, request.Purchasing.VersionLock, "second embedded subset lock")
	assert.Equal(t, 5, request.Account.VersionLock, "pointer subset lock")
	assert.Equal(t, 11, request.Items[0].VersionLock, "slice element lock")
	assert.Equal(t, 12, request.Items[1].VersionLock, "slice element lock")

	// Everything that is not a lock must survive untouched.
	assert.Equal(t, "client", request.DisplayName)
	assert.Equal(t, "loc-new", request.Location.Name)
	assert.False(t, request.Account.Enabled)
}

func TestUnit_SyncAll_NilPointerSubsetIsSafe(t *testing.T) {
	current := &resource{VersionLock: 2, Account: nil}
	request := &resource{VersionLock: 0, Account: nil}

	require.NotPanics(t, func() { SyncAll(current, request) })
	assert.Equal(t, 2, request.VersionLock)

	// Server has the subset, client does not: must not panic or invent one.
	current2 := &resource{VersionLock: 2, Account: &optionalSubset{VersionLock: 9}}
	request2 := &resource{VersionLock: 0, Account: nil}
	require.NotPanics(t, func() { SyncAll(current2, request2) })
	assert.Nil(t, request2.Account)
	assert.Equal(t, 2, request2.VersionLock)
}

func TestUnit_SyncAll_MismatchedSliceLengths(t *testing.T) {
	current := &resource{Items: []subset{{VersionLock: 1}}}
	request := &resource{Items: []subset{{VersionLock: 999}, {VersionLock: 888}}}

	require.NotPanics(t, func() { SyncAll(current, request) })
	assert.Equal(t, 1, request.Items[0].VersionLock, "overlapping element synced")
	assert.Equal(t, 888, request.Items[1].VersionLock, "extra client element untouched")
}

func TestUnit_SyncAll_RejectsMismatchedTypes(t *testing.T) {
	request := &resource{VersionLock: 42}
	require.NotPanics(t, func() { SyncAll(&subset{VersionLock: 1}, request) })
	assert.Equal(t, 42, request.VersionLock, "no cross-type copying")
}

func TestUnit_ZeroAll_ZeroesEveryLockForCreate(t *testing.T) {
	request := &resource{
		VersionLock: 9,
		Location:    subset{VersionLock: 9},
		Purchasing:  subset{VersionLock: 9},
		Account:     &optionalSubset{VersionLock: 9},
		Items:       []subset{{VersionLock: 9}},
		DisplayName: "keep me",
	}

	ZeroAll(request)

	assert.Equal(t, NewResourceVersionLock, request.VersionLock)
	assert.Equal(t, NewResourceVersionLock, request.Location.VersionLock)
	assert.Equal(t, NewResourceVersionLock, request.Purchasing.VersionLock)
	assert.Equal(t, NewResourceVersionLock, request.Account.VersionLock)
	assert.Equal(t, NewResourceVersionLock, request.Items[0].VersionLock)
	assert.Equal(t, "keep me", request.DisplayName)
}

func TestUnit_TopLock(t *testing.T) {
	got, ok := TopLock(&resource{VersionLock: 6})
	assert.True(t, ok)
	assert.Equal(t, 6, got)

	_, ok = TopLock(nil)
	assert.False(t, ok)

	var nilRes *resource
	_, ok = TopLock(nilRes)
	assert.False(t, ok)

	_, ok = TopLock(&struct{ Name string }{})
	assert.False(t, ok, "struct without a lock")
}
