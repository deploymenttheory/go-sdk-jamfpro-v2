package sso_settings

// SsoSettings.configurationType constants.
const (
	ConfigurationTypeSaml        = "SAML"
	ConfigurationTypeOidc        = "OIDC"
	ConfigurationTypeOidcWithSaml = "OIDC_WITH_SAML"
)

// SamlSettings.idpProviderType constants.
const (
	IdpProviderTypeAdfs      = "ADFS"
	IdpProviderTypeOkta      = "OKTA"
	IdpProviderTypeGoogle    = "GOOGLE"
	IdpProviderTypeShibboleth = "SHIBBOLETH"
	IdpProviderTypeOnelogin  = "ONELOGIN"
	IdpProviderTypePing      = "PING"
	IdpProviderTypeCentrify  = "CENTRIFY"
	IdpProviderTypeAzure     = "AZURE"
	IdpProviderTypeOther     = "OTHER"
)

// SamlSettings.metadataSource constants.
const (
	MetadataSourceUrl     = "URL"
	MetadataSourceFile    = "FILE"
	MetadataSourceUnknown = "UNKNOWN"
)

// OidcSettings.userMapping constants.
const (
	UserMappingUsername = "USERNAME"
	UserMappingEmail    = "EMAIL"
)

// SsoKeystore.keystoreSetupType constants.
const (
	KeystoreSetupTypeNone      = "NONE"
	KeystoreSetupTypeUploaded  = "UPLOADED"
	KeystoreSetupTypeGenerated = "GENERATED"
)

// SsoKeystore.type constants.
const (
	KeystoreTypePkcs12 = "PKCS12"
	KeystoreTypeJks    = "JKS"
	KeystoreTypeNone   = "NONE"
)

var validConfigurationTypes = map[string]struct{}{
	ConfigurationTypeSaml:         {},
	ConfigurationTypeOidc:         {},
	ConfigurationTypeOidcWithSaml: {},
}

var validUserMappings = map[string]struct{}{
	UserMappingUsername: {},
	UserMappingEmail:    {},
}

var validMetadataSources = map[string]struct{}{
	MetadataSourceUrl:     {},
	MetadataSourceFile:    {},
	MetadataSourceUnknown: {},
}

var validIdpProviderTypes = map[string]struct{}{
	IdpProviderTypeAdfs:       {},
	IdpProviderTypeOkta:       {},
	IdpProviderTypeGoogle:     {},
	IdpProviderTypeShibboleth: {},
	IdpProviderTypeOnelogin:   {},
	IdpProviderTypePing:       {},
	IdpProviderTypeCentrify:   {},
	IdpProviderTypeAzure:      {},
	IdpProviderTypeOther:      {},
}
