package apilifecycle

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// deprecationOnce dedupes deprecation warnings so each function label warns at
// most once per process. Keyed by funcLabel -> *sync.Once.
var deprecationOnce sync.Map

// DeprecationWarning logs a single zap.Warn the first time a deprecated SDK
// function is invoked in this process. It is a no-op when logger is nil.
//
//	funcLabel:    fully-qualified SDK function, e.g.
//	              "jamf_pro_api/groups.Groups.ListV1"
//	deprecatedIn: the Jamf Pro version the capability was deprecated in,
//	              e.g. "11.28"
//	replacement:  optional migration guidance (may be ""), e.g.
//	              "use the v2 endpoints"
//
// Deduplication is process-global and keyed by funcLabel, so tests that assert
// on the emitted warning must use a label unique to the test.
func DeprecationWarning(logger *zap.Logger, funcLabel, deprecatedIn, replacement string) {
	if logger == nil {
		return
	}
	onceAny, _ := deprecationOnce.LoadOrStore(funcLabel, &sync.Once{})
	onceAny.(*sync.Once).Do(func() {
		fields := []zap.Field{
			zap.String("function", funcLabel),
			zap.String("deprecated_in_jamf_pro_version", deprecatedIn),
		}
		if replacement != "" {
			fields = append(fields, zap.String("replacement", replacement))
		}
		logger.Warn("SDK function is deprecated", fields...)
	})
}

// ServerVersionProvider is the subset of client.Client used by the removal
// guard. Both client.Client (the Transport) and the test mock satisfy it.
type ServerVersionProvider interface {
	// ServerVersion returns the connected Jamf Pro server's parsed version.
	ServerVersion(ctx context.Context) (Version, error)
	// GetLogger returns the configured zap logger (may be nil).
	GetLogger() *zap.Logger
}

// EnsureSupported is the centralised removals guard. It fetches the connected
// Jamf Pro server version and returns a *RemovedError when the server version
// is >= removedIn (the capability has been removed server-side).
//
// Fail-open policy: when the server version cannot be determined (network
// error, empty or unparseable version), EnsureSupported logs a warning and
// returns nil, allowing the call to proceed. The Jamf Pro server is the
// authoritative gate and will itself reject a removed capability; failing
// closed would let a transient /jamf-pro-version outage break otherwise valid
// calls. Flip the err branch below to return the error if fail-closed is
// preferred.
func EnsureSupported(ctx context.Context, c ServerVersionProvider, funcLabel string, removedIn Version) error {
	sv, err := c.ServerVersion(ctx)
	if err != nil {
		if lg := c.GetLogger(); lg != nil {
			lg.Warn("removal guard could not determine server version; allowing call",
				zap.String("function", funcLabel),
				zap.String("removed_in_jamf_pro_version", removedIn.String()),
				zap.Error(err),
			)
		}
		return nil
	}
	if sv.AtLeast(removedIn) {
		return &RemovedError{Function: funcLabel, RemovedIn: removedIn, ServerVer: sv}
	}
	return nil
}

// RemovedError is returned by EnsureSupported when a capability has been
// removed from the connected Jamf Pro server.
type RemovedError struct {
	Function  string
	RemovedIn Version
	ServerVer Version
}

func (e *RemovedError) Error() string {
	return fmt.Sprintf(
		"capability %s was removed in Jamf Pro %s and is not supported by the connected server (version %s)",
		e.Function, e.RemovedIn.String(), e.ServerVer.String(),
	)
}

// IsRemoved reports whether err is, or wraps, a *RemovedError.
func IsRemoved(err error) bool {
	var re *RemovedError
	return errors.As(err, &re)
}
