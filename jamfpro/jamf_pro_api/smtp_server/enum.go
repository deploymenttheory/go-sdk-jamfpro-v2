package smtp_server

// SmtpServer.authenticationType constants.
const (
	AuthenticationTypeNone       = "NONE"
	AuthenticationTypeBasic      = "BASIC"
	AuthenticationTypeGraphApi   = "GRAPH_API"
	AuthenticationTypeGoogleMail = "GOOGLE_MAIL"
)

// SmtpConnectionSettings.encryptionType constants.
const (
	EncryptionTypeNone  = "NONE"
	EncryptionTypeSsl   = "SSL"
	EncryptionTypeTls12 = "TLS_1_2"
	EncryptionTypeTls11 = "TLS_1_1"
	EncryptionTypeTls1  = "TLS_1"
	EncryptionTypeTls13 = "TLS_1_3"
)

var validAuthenticationTypes = map[string]struct{}{
	AuthenticationTypeNone:       {},
	AuthenticationTypeBasic:      {},
	AuthenticationTypeGraphApi:   {},
	AuthenticationTypeGoogleMail: {},
}

var validEncryptionTypes = map[string]struct{}{
	EncryptionTypeNone:  {},
	EncryptionTypeSsl:   {},
	EncryptionTypeTls12: {},
	EncryptionTypeTls11: {},
	EncryptionTypeTls1:  {},
	EncryptionTypeTls13: {},
}
