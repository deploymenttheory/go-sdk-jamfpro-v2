package acceptance

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// SkipIfNotConfigured skips the test when Jamf Pro credentials are not set.
func SkipIfNotConfigured(t *testing.T) {
	t.Helper()
	if !IsConfigured() {
		t.Skip("INSTANCE_DOMAIN or AUTH_METHOD not set, skipping acceptance test")
	}
}

// RequireClient ensures the shared client is initialised, skipping if
// credentials are absent or initialisation fails.
func RequireClient(t *testing.T) {
	t.Helper()
	SkipIfNotConfigured(t)

	if Client == nil {
		err := InitClient()
		require.NoError(t, err, "Failed to initialise Jamf Pro client")
	}
}

// NewContext creates a context with the configured request timeout.
func NewContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), Config.RequestTimeout)
}

// Cleanup registers a cleanup function, skipped when JAMF_SKIP_CLEANUP=true.
func Cleanup(t *testing.T, fn func()) {
	t.Helper()
	if !Config.SkipCleanup {
		t.Cleanup(fn)
	} else if Config.Verbose {
		t.Log("Skipping cleanup (JAMF_SKIP_CLEANUP=true)")
	}
}

// LogTestStage logs a named test stage with optional GitHub Actions annotation.
func LogTestStage(t *testing.T, stage, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::notice title=%s::%s\n", stage, formatted)
	}
	if Config.Verbose {
		t.Logf("[%s] %s", stage, formatted)
	}
}

// LogTestSuccess logs a successful step.
func LogTestSuccess(t *testing.T, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::notice title=Success::%s\n", formatted)
	}
	if Config.Verbose {
		t.Logf("OK: %s", formatted)
	}
}

// LogTestWarning logs a non-fatal warning.
func LogTestWarning(t *testing.T, message string, args ...any) {
	t.Helper()
	formatted := message
	if len(args) > 0 {
		formatted = fmt.Sprintf(message, args...)
	}
	if isGitHubActions() {
		fmt.Printf("::warning title=Warning::%s\n", formatted)
	}
	if Config.Verbose {
		t.Logf("WARNING: %s", formatted)
	}
}

// PollUntil retries fn every interval until it returns true or timeout elapses.
// Used to wait for eventually-consistent API state.
func PollUntil(t *testing.T, timeout, interval time.Duration, fn func() bool) bool {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if fn() {
			return true
		}
		time.Sleep(interval)
	}
	return false
}

func isGitHubActions() bool {
	return os.Getenv("GITHUB_ACTIONS") == "true"
}
