package macos_configuration_profile_custom_settings

// ResponseCustomSettingsSchemaList represents the top-level response for custom settings schema list.
// It is a slice of buckets, each containing custom settings domains.
type ResponseCustomSettingsSchemaList []ResourceCustomSettingsBucket

// ResourceCustomSettingsBucket represents a single bucket of custom settings schemas.
type ResourceCustomSettingsBucket struct {
	BucketName            string                                  `json:"bucketName"`
	DisplayName           string                                  `json:"displayName"`
	CustomSettingsDomains map[string]ResourceCustomSettingsDomain `json:"customSettingsDomains"`
}

// ResourceCustomSettingsDomain represents a domain of custom settings.
type ResourceCustomSettingsDomain struct {
	SettingsDomain string                           `json:"settingsDomain"`
	Versions       map[string]ResourceDomainVersion `json:"versions"`
}

// ResourceDomainVersion represents a specific version of a domain.
type ResourceDomainVersion struct {
	Version  string   `json:"version"`
	Variants []string `json:"variants"`
}

// ResourceConfigProfile represents a macOS configuration profile.
type ResourceConfigProfile struct {
	PayloadUUID    string               `json:"payloadUUID"`
	PayloadContent []PayloadContentItem `json:"payloadContent"`
	Level          string               `json:"level,omitempty"`
}

// PayloadContentItem represents an item in the payload content.
type PayloadContentItem struct {
	PayloadType         string          `json:"payloadType"`
	PayloadVersion      int             `json:"payloadVersion"`
	PayloadIdentifier   string          `json:"payloadIdentifier"`
	PayloadUUID         string          `json:"payloadUUID"`
	PayloadOrganization string          `json:"payloadOrganization,omitempty"`
	PreferenceDomain    string          `json:"preferenceDomain,omitempty"`
	Forced              *ForcedSettings `json:"forced,omitempty"`
	PayloadDisplayName  string          `json:"payloadDisplayName,omitempty"`
}

// ForcedSettings represents forced settings in a payload.
type ForcedSettings struct {
	Plist         string `json:"plist,omitempty"`
	JsonSchema    string `json:"jsonSchema,omitempty"`
	SchemaSource  string `json:"schemaSource,omitempty"`
	SchemaDomain  string `json:"schemaDomain,omitempty"`
	SchemaVersion string `json:"schemaVersion,omitempty"`
	SchemaVariant string `json:"schemaVariant,omitempty"`
}

// ResponseConfigProfileCreate represents the response when creating a configuration profile.
type ResponseConfigProfileCreate struct {
	UUID string `json:"uuid"`
}
