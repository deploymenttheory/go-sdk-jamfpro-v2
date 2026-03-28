package self_service_settings

// SelfServiceInteractionSettings.defaultLandingPage constants.
const (
	DefaultLandingPageHome          = "HOME"
	DefaultLandingPageBrowse        = "BROWSE"
	DefaultLandingPageHistory       = "HISTORY"
	DefaultLandingPageNotifications = "NOTIFICATIONS"
)

// SelfServiceLoginSettings.authType constants.
const (
	AuthTypeBasic = "Basic"
	AuthTypeSaml  = "Saml"
)

// SelfServiceLoginSettings.userLoginLevel constants.
const (
	UserLoginLevelNotRequired = "NotRequired"
	UserLoginLevelAnonymous   = "Anonymous"
	UserLoginLevelRequired    = "Required"
)

var validUserLoginLevels = map[string]struct{}{
	UserLoginLevelNotRequired: {},
	UserLoginLevelAnonymous:   {},
	UserLoginLevelRequired:    {},
}

var validAuthTypes = map[string]struct{}{
	AuthTypeBasic: {},
	AuthTypeSaml:  {},
}

var validDefaultLandingPages = map[string]struct{}{
	DefaultLandingPageHome:          {},
	DefaultLandingPageBrowse:        {},
	DefaultLandingPageHistory:       {},
	DefaultLandingPageNotifications: {},
}
