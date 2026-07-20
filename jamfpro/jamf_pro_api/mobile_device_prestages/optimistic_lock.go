package mobile_device_prestages

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/version_locking"
	"resty.dev/v3"
)

// fetchForUpdate supplies current server state to the version locking layer.
//
// When the caller has already fetched the resource (a name lookup, say), that
// copy is handed back for the first attempt so the common path costs one
// round-trip. Any retry re-reads from the API, since the whole point of
// retrying is to pick up locks that have moved on.
func (s *MobileDevicePrestages) fetchForUpdate(id string, prefetched *ResourceMobileDevicePrestage) version_locking.Fetch[ResourceMobileDevicePrestage] {
	return func(ctx context.Context) (*ResourceMobileDevicePrestage, *resty.Response, error) {
		if prefetched != nil {
			current := prefetched
			prefetched = nil
			return current, nil, nil
		}
		current, resp, err := s.GetByIDV3(ctx, id)
		if err != nil {
			return nil, resp, fmt.Errorf("failed to fetch current prestage for version locking: %w", err)
		}
		return current, resp, nil
	}
}

// putByID returns the raw PUT submission for a prestage, without any version
// lock handling — version_locking.Update owns that.
func (s *MobileDevicePrestages) putByID(id string) version_locking.Submit[ResourceMobileDevicePrestage] {
	return func(ctx context.Context, prestage *ResourceMobileDevicePrestage) (*ResourceMobileDevicePrestage, *resty.Response, error) {
		endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProMobileDevicePrestagesV3, id)

		var result ResourceMobileDevicePrestage

		// The body carries version locks the server consumes on the first write,
		// so a transport-level replay would resubmit spent locks. Retry decisions
		// belong to version_locking.Update, which re-reads state first.
		resp, err := s.client.NewRequest(ctx).
			SetHeader("Accept", constants.ApplicationJSON).
			SetHeader("Content-Type", constants.ApplicationJSON).
			SetBody(prestage).
			SetResult(&result).
			DisableRetry().
			Put(endpoint)
		if err != nil {
			return nil, resp, err
		}

		return &result, resp, nil
	}
}
