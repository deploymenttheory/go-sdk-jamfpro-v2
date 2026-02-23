package macos_configuration_profile_custom_settings

const (
	// EndpointConfigProfilesMacOS is the base endpoint for macOS configuration profile operations.
	// API reference: Undocumented
	EndpointConfigProfilesMacOS = "/api/config-profiles/macos"

	// EndpointCustomSettingsSchemaList is the endpoint for retrieving custom settings schema list.
	// GET /api/config-profiles/macos/custom-settings/v1/schema-list
	EndpointCustomSettingsSchemaList = EndpointConfigProfilesMacOS + "/custom-settings/v1/schema-list"
)
