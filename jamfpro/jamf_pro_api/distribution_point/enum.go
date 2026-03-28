package distribution_point

// FileSharingConnectionType enum values for ResourceDistributionPoint.FileSharingConnectionType.
const (
	FileSharingConnectionTypeAFP  = "AFP"
	FileSharingConnectionTypeSMB  = "SMB"
	FileSharingConnectionTypeNone = "NONE"
)

// HTTPSSecurityType enum values for ResourceDistributionPoint.HTTPSSecurityType.
const (
	HTTPSSecurityTypeUsernamePassword = "USERNAME_PASSWORD"
	HTTPSSecurityTypeNone             = "NONE"
)

// validFileSharingConnectionTypes is the set of accepted FileSharingConnectionType values.
var validFileSharingConnectionTypes = map[string]struct{}{
	FileSharingConnectionTypeAFP:  {},
	FileSharingConnectionTypeSMB:  {},
	FileSharingConnectionTypeNone: {},
}

// validHTTPSSecurityTypes is the set of accepted HTTPSSecurityType values.
var validHTTPSSecurityTypes = map[string]struct{}{
	HTTPSSecurityTypeUsernamePassword: {},
	HTTPSSecurityTypeNone:             {},
}
