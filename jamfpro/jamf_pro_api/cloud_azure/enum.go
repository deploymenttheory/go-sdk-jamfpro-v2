package cloud_azure

// AzureServerConfigurationType enum values for CloudAzureServer.Type.
const (
	AzureServerConfigurationTypePublic  = "PUBLIC"
	AzureServerConfigurationTypeGccHigh = "GCC_HIGH"
)

var validAzureServerConfigurationTypes = map[string]struct{}{
	AzureServerConfigurationTypePublic:  {},
	AzureServerConfigurationTypeGccHigh: {},
}
