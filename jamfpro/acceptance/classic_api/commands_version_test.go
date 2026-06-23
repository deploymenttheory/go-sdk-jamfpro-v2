package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
)

// serverVersionOrSkip fetches and parses the connected Jamf Pro server version,
// skipping the test if it cannot be determined. Shared by the command-removal
// acceptance tests.
func serverVersionOrSkip(t *testing.T, ctx context.Context) apilifecycle.Version {
	t.Helper()
	res, _, err := acc.Client.JamfProAPI.JamfProVersion.GetV1(ctx)
	if err != nil {
		t.Skipf("could not get Jamf Pro version: %v", err)
	}
	if res == nil || res.Version == nil || *res.Version == "" {
		t.Skip("Jamf Pro version is empty")
	}
	v, err := apilifecycle.Parse(*res.Version)
	if err != nil {
		t.Skipf("could not parse Jamf Pro version %q: %v", *res.Version, err)
	}
	return v
}
