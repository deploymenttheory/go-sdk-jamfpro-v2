package macos_configuration_profiles

// ConfigProfileLevel enum values for ResourceConfigProfile.Level.
//
// The API accepts only these two exact uppercase values. Any other string --
// including the title-case "System" / "User" that real .mobileconfig files
// carry in their PayloadScope key -- is rejected with HTTP 400 INVALID_CONTENT
// "Couldn't map content", an error which names no field.
//
// Level is write-only: it is accepted on create and update but is never
// returned by GET.
const (
	ConfigProfileLevelSystem = "SYSTEM"
	ConfigProfileLevelUser   = "USER"
)

// validConfigProfileLevels is the set of accepted Level values.
var validConfigProfileLevels = map[string]struct{}{
	ConfigProfileLevelSystem: {},
	ConfigProfileLevelUser:   {},
}

// Payload types known to be accepted by this endpoint. These are not the only
// supported types -- the allowlist is undocumented and wider than this set --
// but they are the ones confirmed to create successfully with no type-specific
// required fields. See the package documentation in crud.go for the full
// picture, including types that are recognised but carry their own field
// requirements, and those the server explicitly refuses.
//
// payloadType is not validated client-side; an unsupported value is reported
// by the server as HTTP 400 INVALID_FIELD "Provide a proper payloadType field".
const (
	// PayloadTypeManagedClientPreferences is the Custom Settings payload. It
	// is the only type that requires preferenceDomain.
	PayloadTypeManagedClientPreferences = "com.apple.ManagedClient.preferences"
	PayloadTypeNotificationSettings     = "com.apple.notificationsettings"
	PayloadTypePasswordPolicy           = "com.apple.mobiledevice.passwordpolicy"
)
