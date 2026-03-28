package onboarding

// OnboardingItem.selfServiceEntityType constants.
const (
	SelfServiceEntityTypeOsXPolicy        = "OS_X_POLICY"
	SelfServiceEntityTypeOsXConfigProfile = "OS_X_CONFIG_PROFILE"
	SelfServiceEntityTypeOsXMacApp        = "OS_X_MAC_APP"
	SelfServiceEntityTypeOsXAppInstaller  = "OS_X_APP_INSTALLER"
	SelfServiceEntityTypeOsXEbook         = "OS_X_EBOOK"
	SelfServiceEntityTypeOsXPatchPolicy   = "OS_X_PATCH_POLICY"
	SelfServiceEntityTypeUnknown          = "UNKNOWN"
)

var validSelfServiceEntityTypes = map[string]struct{}{
	SelfServiceEntityTypeOsXPolicy:        {},
	SelfServiceEntityTypeOsXConfigProfile: {},
	SelfServiceEntityTypeOsXMacApp:        {},
	SelfServiceEntityTypeOsXAppInstaller:  {},
	SelfServiceEntityTypeOsXEbook:         {},
	SelfServiceEntityTypeOsXPatchPolicy:   {},
	SelfServiceEntityTypeUnknown:          {},
}
