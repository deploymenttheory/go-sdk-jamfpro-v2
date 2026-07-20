// Package macos_configuration_profiles implements the
// undocumented /api/config-profiles/macos endpoint.
//
// The endpoint appears in no published Jamf Pro OpenAPI spec. The contract
// below was established by probing a live instance running Jamf Pro 11.30.0 on
// 2026-07-20, and may change without notice.
//
// Routes and allowed methods (from OPTIONS):
//
//	/api/config-profiles/macos                                  POST, OPTIONS
//	/api/config-profiles/macos/{uuid}                            GET, HEAD, PUT, DELETE, OPTIONS
//	/api/config-profiles/macos/custom-settings/v1/schema-list     GET
//
// There is no list operation: GET on the collection returns HTTP 405. The path
// is unversioned -- /api/v1/config-profiles/macos returns 404 -- which is why
// the methods in this package carry no V1 suffix.
//
// Scope: this endpoint manages a payload-authoring object, not a scoped,
// deployable configuration profile. Profiles created here do not appear in the
// Classic API /JSSResource/osxconfigurationprofiles list.
//
// Supported payload types are an undocumented allowlist. Probing 61 types
// against 11.30.0 produced three distinct responses, which are the only way to
// tell membership:
//
//   - Accepted: com.apple.ManagedClient.preferences,
//     com.apple.notificationsettings, com.apple.mobiledevice.passwordpolicy.
//   - Recognised, with their own field rules -- an INVALID_CONTENT error
//     naming a type-specific field: com.apple.servicemanagement,
//     com.apple.extensiblesso, com.apple.webcontent-filter,
//     com.apple.security.acme, com.apple.relay.managed,
//     com.apple.dnsSettings.managed, com.apple.dnsProxy.managed.
//   - Known but refused: com.apple.applicationaccess,
//     com.apple.airplay.security, reported as "Payload type is not supported."
//
// Everything else returns INVALID_FIELD "Provide a proper payloadType field".
// Notably this excludes certificates, TCC/PPPC, VPN, FileVault and system
// extension payloads, so the endpoint does not cover every payload a real
// .mobileconfig may contain. The list above is not assumed complete and is
// therefore documented rather than enforced -- validateConfigProfile does not
// check payloadType.
//
// Payload constraints validated client-side by this package, because the
// server reports them as opaque HTTP 400s:
//
//   - preferenceDomain is required by com.apple.ManagedClient.preferences
//     payloads only; other supported types create fine without it.
//   - forced.plist is required whenever forced is set, and must be a complete
//     plist document rather than a bare <dict> fragment.
//   - level accepts only "SYSTEM" or "USER" (see enum.go), and is write-only.
//
// Known server-side defect: DELETE returns 204 for a resource that exists and
// 404 for one that does not, but does not actually remove the object. A
// profile polled every 30s for 10 minutes after a successful DELETE continued
// to return 200 on GET. Callers must not treat a 204 as proof of removal.
package macos_configuration_profiles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the macOS configuration profile custom settings
	// methods of the Jamf Pro API.
	//
	// API reference: Undocumented
	MacosConfigurationProfiles struct {
		client client.Client
	}
)

// NewService creates a new macOS configuration profile custom settings service.
func NewMacosConfigurationProfiles(client client.Client) *MacosConfigurationProfiles {
	return &MacosConfigurationProfiles{client: client}
}

// validateConfigProfile applies the constraints the API enforces server-side,
// so callers receive an error naming the offending field instead of an opaque
// HTTP 400 INVALID_CONTENT.
//
// It deliberately does not parse forced.plist to confirm it is a well-formed
// complete plist document; that risks rejecting payloads the server would
// accept. See the ForcedSettings.Plist field comment.
func validateConfigProfile(profile *ResourceConfigProfile) error {
	if profile.Level != "" {
		if _, ok := validConfigProfileLevels[profile.Level]; !ok {
			return fmt.Errorf("invalid level %q: must be one of SYSTEM, USER", profile.Level)
		}
	}

	for i, item := range profile.PayloadContent {
		// payloadType is deliberately not checked against an allowlist. The
		// server supports a set of types that is undocumented and wider than
		// it first appears, so a client-side list would reject payloads the
		// API accepts. Unsupported types are reported by the server as
		// INVALID_FIELD "Provide a proper payloadType field".

		// preferenceDomain is required only by the Custom Settings payload.
		// Other supported types (com.apple.notificationsettings,
		// com.apple.mobiledevice.passwordpolicy, ...) create fine without it.
		if item.PayloadType == PayloadTypeManagedClientPreferences && item.PreferenceDomain == "" {
			return fmt.Errorf("payloadContent[%d]: preferenceDomain is required for %s payloads", i, PayloadTypeManagedClientPreferences)
		}

		if item.Forced != nil && item.Forced.Plist == "" {
			return fmt.Errorf("payloadContent[%d]: forced.plist is required when forced is set", i)
		}
	}

	return nil
}

// GetSchemaList retrieves the list of custom settings schemas.
// URL: GET /api/config-profiles/macos/custom-settings/v1/schema-list
func (s *MacosConfigurationProfiles) GetSchemaList(ctx context.Context) (*ResponseCustomSettingsSchemaList, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCustomSettingsSchemaList

	var result ResponseCustomSettingsSchemaList

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get custom settings schema list: %w", err)
	}

	return &result, resp, nil
}

// GetByPayloadUUID retrieves a macOS configuration profile by payload UUID.
// URL: GET /api/config-profiles/macos/{id}
func (s *MacosConfigurationProfiles) GetByPayloadUUID(ctx context.Context, id string) (*ResourceConfigProfile, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("payload UUID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProConfigProfilesMacOS, id)

	var result ResourceConfigProfile

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get macOS configuration profile with ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// Create creates a new macOS configuration profile with custom settings schema.
// URL: POST /api/config-profiles/macos
func (s *MacosConfigurationProfiles) Create(ctx context.Context, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *resty.Response, error) {
	if profile == nil {
		return nil, nil, fmt.Errorf("profile is required")
	}

	if err := validateConfigProfile(profile); err != nil {
		return nil, nil, err
	}

	endpoint := constants.EndpointJamfProConfigProfilesMacOS

	var result ResponseConfigProfileCreate

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(profile).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create macOS configuration profile: %w", err)
	}

	return &result, resp, nil
}

// UpdateByPayloadUUID replaces a macOS configuration profile by payload UUID.
// URL: PUT /api/config-profiles/macos/{id}
//
// The server returns {"uuid": "..."} -- the same envelope as Create, not the
// updated resource -- which is why the return type is *ResponseConfigProfileCreate.
// Call GetByPayloadUUID afterwards to read the profile back.
//
// The payloadUUID field in the request body is ignored; the path parameter wins.
func (s *MacosConfigurationProfiles) UpdateByPayloadUUID(ctx context.Context, id string, profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("payload UUID is required")
	}

	if profile == nil {
		return nil, nil, fmt.Errorf("profile is required")
	}

	if err := validateConfigProfile(profile); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProConfigProfilesMacOS, id)

	var result ResponseConfigProfileCreate

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(profile).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update macOS configuration profile with ID %s: %w", id, err)
	}

	return &result, resp, nil
}

// DeleteByPayloadUUID removes a macOS configuration profile by payload UUID.
// URL: DELETE /api/config-profiles/macos/{id}
//
// Returns 204 No Content for a profile that exists, 404 INVALID_ID for one
// that does not.
//
// A 204 does not guarantee the profile was removed. Against Jamf Pro 11.30.0
// the object remained readable via GetByPayloadUUID indefinitely after a
// successful delete -- confirmed by polling for 10 minutes, so this is a
// server-side defect rather than replication lag. Do not assert that a
// subsequent read fails.
func (s *MacosConfigurationProfiles) DeleteByPayloadUUID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("payload UUID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProConfigProfilesMacOS, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to delete macOS configuration profile with ID %s: %w", id, err)
	}

	return resp, nil
}
