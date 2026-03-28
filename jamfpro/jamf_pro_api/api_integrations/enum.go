package api_integrations

// ApiIntegration.appType constants.
const (
	AppTypeClientCredentials = "CLIENT_CREDENTIALS"
	AppTypeNativeAppOauth    = "NATIVE_APP_OAUTH"
	AppTypeNone              = "NONE"
)

var validAppTypes = map[string]struct{}{
	AppTypeClientCredentials: {},
	AppTypeNativeAppOauth:    {},
	AppTypeNone:              {},
}
