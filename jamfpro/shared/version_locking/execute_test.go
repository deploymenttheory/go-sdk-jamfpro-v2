package version_locking

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// respWith builds a *resty.Response carrying the given status code.
func respWith(t *testing.T, code int) *resty.Response {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(code)
	}))
	t.Cleanup(srv.Close)
	c := resty.New()
	t.Cleanup(func() { _ = c.Close() })
	r, err := c.R().Get(srv.URL)
	require.NoError(t, err)
	return r
}

func TestUnit_Update_SyncsLocksFromServerBeforeSubmitting(t *testing.T) {
	server := &resource{
		VersionLock: 5,
		Location:    subset{VersionLock: 2},
		Purchasing:  subset{VersionLock: 3},
		Account:     &optionalSubset{VersionLock: 4},
	}
	// Caller supplies nonsense locks; they must be overwritten wholesale.
	request := &resource{
		DisplayName: "new name",
		VersionLock: 999,
		Location:    subset{VersionLock: 999},
		Purchasing:  subset{VersionLock: 999},
		Account:     &optionalSubset{VersionLock: 999},
	}

	var got resource
	_, _, err := Update(context.Background(), request,
		func(context.Context) (*resource, *resty.Response, error) { return server, nil, nil },
		func(_ context.Context, r *resource) (*resource, *resty.Response, error) {
			got = *r
			return r, respWith(t, http.StatusOK), nil
		})

	require.NoError(t, err)
	assert.Equal(t, 5, got.VersionLock)
	assert.Equal(t, 2, got.Location.VersionLock)
	assert.Equal(t, 3, got.Purchasing.VersionLock)
	assert.Equal(t, 4, got.Account.VersionLock)
	assert.Equal(t, "new name", got.DisplayName, "payload preserved")
}

// The reproduced Jamf behaviour: HTTP 500, empty error body, write applied.
// Must be reported as success, and must NOT be resubmitted.
func TestUnit_Update_TreatsAppliedWriteReportedAsErrorAsSuccess(t *testing.T) {
	lock := 5
	submits := 0

	result, _, err := Update(context.Background(), &resource{DisplayName: "x"},
		func(context.Context) (*resource, *resty.Response, error) {
			return &resource{VersionLock: lock}, nil, nil
		},
		func(context.Context, *resource) (*resource, *resty.Response, error) {
			submits++
			lock++ // server applied it despite the error
			return nil, respWith(t, http.StatusInternalServerError), errors.New("500")
		})

	require.NoError(t, err, "advanced lock proves the write landed")
	require.NotNil(t, result)
	assert.Equal(t, 6, result.VersionLock, "refreshed state returned")
	assert.Equal(t, 1, submits, "must not resubmit a write that landed")
}

// A genuine conflict: write rejected, lock unchanged. Retry with a fresh lock.
func TestUnit_Update_RetriesConflictWithFreshLock(t *testing.T) {
	lock := 5
	submits := 0
	var locksSent []int

	_, _, err := Update(context.Background(), &resource{DisplayName: "x"},
		func(context.Context) (*resource, *resty.Response, error) {
			return &resource{VersionLock: lock}, nil, nil
		},
		func(_ context.Context, r *resource) (*resource, *resty.Response, error) {
			submits++
			locksSent = append(locksSent, r.VersionLock)
			if submits == 1 {
				lock = 9 // another writer moved it on; our attempt is rejected
				return nil, respWith(t, http.StatusConflict), errors.New("OPTIMISTIC_LOCK_FAILED")
			}
			return r, respWith(t, http.StatusOK), nil
		})

	require.NoError(t, err)
	assert.Equal(t, 2, submits)
	assert.Equal(t, []int{5, 9}, locksSent, "retry must carry the re-read lock, not the stale one")
}

// A conflict must never be reconciled by inspecting state: a competing writer
// advances the lock too, and treating that as our success silently discards the
// caller's change.
func TestUnit_Update_ConflictIsNotMistakenForSuccessWhenAnotherWriterCommits(t *testing.T) {
	lock := 5
	submits := 0

	_, _, err := Update(context.Background(), &resource{DisplayName: "mine"},
		func(context.Context) (*resource, *resty.Response, error) {
			return &resource{VersionLock: lock}, nil, nil
		},
		func(context.Context, *resource) (*resource, *resty.Response, error) {
			submits++
			lock += 10 // a different client committed; our write was rejected
			return nil, respWith(t, http.StatusConflict), errors.New("OPTIMISTIC_LOCK_FAILED")
		},
		WithMaxAttempts(2))

	require.Error(t, err, "rejected write must not be reported as success")
	assert.Equal(t, 2, submits, "conflict retries rather than short-circuiting")
}

func TestUnit_Update_DoesNotRetryUnrecoverableErrors(t *testing.T) {
	submits := 0
	_, _, err := Update(context.Background(), &resource{},
		func(context.Context) (*resource, *resty.Response, error) {
			return &resource{VersionLock: 1}, nil, nil
		},
		func(context.Context, *resource) (*resource, *resty.Response, error) {
			submits++
			return nil, respWith(t, http.StatusBadRequest), errors.New("validation failed")
		})

	require.Error(t, err)
	assert.Equal(t, 1, submits, "a 400 fails identically however the locks are set")
}

func TestUnit_Update_GivesUpAfterMaxAttempts(t *testing.T) {
	submits := 0
	_, _, err := Update(context.Background(), &resource{},
		func(context.Context) (*resource, *resty.Response, error) {
			return &resource{VersionLock: 1}, nil, nil
		},
		func(context.Context, *resource) (*resource, *resty.Response, error) {
			submits++
			return nil, respWith(t, http.StatusConflict), errors.New("conflict")
		},
		WithMaxAttempts(2))

	require.Error(t, err)
	assert.Equal(t, 2, submits)
}

func TestUnit_Update_PropagatesFetchFailure(t *testing.T) {
	submits := 0
	_, _, err := Update(context.Background(), &resource{},
		func(context.Context) (*resource, *resty.Response, error) {
			return nil, nil, errors.New("GET failed")
		},
		func(context.Context, *resource) (*resource, *resty.Response, error) {
			submits++
			return nil, nil, nil
		})

	require.ErrorContains(t, err, "GET failed")
	assert.Zero(t, submits, "never write without knowing the current lock")
}

func TestUnit_Create_ZeroesAllLocks(t *testing.T) {
	request := &resource{
		VersionLock: 7,
		Location:    subset{VersionLock: 7},
		Purchasing:  subset{VersionLock: 7},
		Account:     &optionalSubset{VersionLock: 7},
	}

	var got resource
	_, _, err := Create(context.Background(), request,
		func(_ context.Context, r *resource) (*resource, *resty.Response, error) {
			got = *r
			return r, nil, nil
		})

	require.NoError(t, err)
	assert.Zero(t, got.VersionLock)
	assert.Zero(t, got.Location.VersionLock)
	assert.Zero(t, got.Purchasing.VersionLock)
	assert.Zero(t, got.Account.VersionLock)
}
