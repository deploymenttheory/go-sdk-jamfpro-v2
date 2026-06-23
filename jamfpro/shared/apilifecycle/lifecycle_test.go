package apilifecycle_test

import (
	"context"
	"errors"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

// fakeProvider is a minimal apilifecycle.ServerVersionProvider for guard tests.
type fakeProvider struct {
	version string
	err     error
	logger  *zap.Logger
}

func (f fakeProvider) ServerVersion(context.Context) (apilifecycle.Version, error) {
	if f.err != nil {
		return apilifecycle.Version{}, f.err
	}
	return apilifecycle.Parse(f.version)
}

func (f fakeProvider) GetLogger() *zap.Logger { return f.logger }

func TestUnit_DeprecationWarning_LogsOncePerFunc(t *testing.T) {
	core, logs := observer.New(zap.WarnLevel)
	logger := zap.New(core)

	// Process-global dedup: use a label unique to this test.
	label := "pkg.Type.Func/" + t.Name()
	apilifecycle.DeprecationWarning(logger, label, "11.28", "use the v2 endpoints")
	apilifecycle.DeprecationWarning(logger, label, "11.28", "use the v2 endpoints")

	entries := logs.FilterMessage("SDK function is deprecated").All()
	require.Len(t, entries, 1)
	fields := entries[0].ContextMap()
	assert.Equal(t, label, fields["function"])
	assert.Equal(t, "11.28", fields["deprecated_in_jamf_pro_version"])
	assert.Equal(t, "use the v2 endpoints", fields["replacement"])
}

func TestUnit_DeprecationWarning_NilLoggerNoPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		apilifecycle.DeprecationWarning(nil, "pkg.Type.Func/"+t.Name(), "11.28", "")
	})
}

func TestUnit_EnsureSupported_BelowRemoval_Allows(t *testing.T) {
	err := apilifecycle.EnsureSupported(context.Background(),
		fakeProvider{version: "11.27.9"}, "pkg.F", apilifecycle.MustParse("11.28.0"))
	require.NoError(t, err)
}

func TestUnit_EnsureSupported_AtRemoval_Errors(t *testing.T) {
	err := apilifecycle.EnsureSupported(context.Background(),
		fakeProvider{version: "11.28.0"}, "pkg.F", apilifecycle.MustParse("11.28.0"))
	require.Error(t, err)
	assert.True(t, apilifecycle.IsRemoved(err))

	var re *apilifecycle.RemovedError
	require.True(t, errors.As(err, &re))
	assert.Equal(t, "pkg.F", re.Function)
	assert.Equal(t, "11.28.0", re.RemovedIn.String())
	assert.Equal(t, "11.28.0", re.ServerVer.String())
}

func TestUnit_EnsureSupported_AboveRemoval_Errors(t *testing.T) {
	err := apilifecycle.EnsureSupported(context.Background(),
		fakeProvider{version: "11.29.0"}, "pkg.F", apilifecycle.MustParse("11.28.0"))
	require.Error(t, err)
	assert.True(t, apilifecycle.IsRemoved(err))
}

func TestUnit_EnsureSupported_VersionLookupFails_FailsOpen(t *testing.T) {
	core, logs := observer.New(zap.WarnLevel)
	err := apilifecycle.EnsureSupported(context.Background(),
		fakeProvider{err: errors.New("boom"), logger: zap.New(core)},
		"pkg.F", apilifecycle.MustParse("11.28.0"))
	require.NoError(t, err) // fail-open
	assert.Equal(t, 1, logs.FilterMessage("removal guard could not determine server version; allowing call").Len())
}

func TestUnit_IsRemoved_NonRemovedError(t *testing.T) {
	assert.False(t, apilifecycle.IsRemoved(errors.New("some other error")))
	assert.False(t, apilifecycle.IsRemoved(nil))
}
