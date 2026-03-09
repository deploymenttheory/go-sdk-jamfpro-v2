package mobile_device_apps

// RequestReinstallAppConfig represents the request body for reinstalling app config.
// The reinstallCode is the $APP_CONFIG_REINSTALL_CODE variable for the specific
// device and app supplied by the managed iOS app's current App Config.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
type RequestReinstallAppConfig struct {
	ReinstallCode string `json:"reinstallCode"`
}
