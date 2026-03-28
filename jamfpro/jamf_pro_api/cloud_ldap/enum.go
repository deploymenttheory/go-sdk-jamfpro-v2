package cloud_ldap

// CloudLdapServerConnectionType enum values for CloudLdapServer.ConnectionType.
const (
	CloudLdapServerConnectionTypeLDAPS    = "LDAPS"
	CloudLdapServerConnectionTypeStartTLS = "START_TLS"
)

// ObjectClassLimitation enum values.
const (
	ObjectClassLimitationAnyObjectClasses = "ANY_OBJECT_CLASSES"
	ObjectClassLimitationAllObjectClasses = "ALL_OBJECT_CLASSES"
)

// SearchScope enum values.
const (
	SearchScopeAllSubtrees    = "ALL_SUBTREES"
	SearchScopeFirstLevelOnly = "FIRST_LEVEL_ONLY"
)

var validConnectionTypes = map[string]struct{}{
	CloudLdapServerConnectionTypeLDAPS:    {},
	CloudLdapServerConnectionTypeStartTLS: {},
}

var validObjectClassLimitations = map[string]struct{}{
	ObjectClassLimitationAnyObjectClasses: {},
	ObjectClassLimitationAllObjectClasses: {},
}

var validSearchScopes = map[string]struct{}{
	SearchScopeAllSubtrees:    {},
	SearchScopeFirstLevelOnly: {},
}
