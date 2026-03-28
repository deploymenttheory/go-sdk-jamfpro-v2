package accounts

// UserAccount.accessLevel constants.
const (
	AccessLevelFullAccess       = "FullAccess"
	AccessLevelSiteAccess       = "SiteAccess"
	AccessLevelGroupBasedAccess = "GroupBasedAccess"
)

// UserAccount.privilegeLevel constants.
const (
	PrivilegeLevelAdministrator = "ADMINISTRATOR"
	PrivilegeLevelAuditor       = "AUDITOR"
	PrivilegeLevelEnrollment    = "ENROLLMENT"
	PrivilegeLevelCustom        = "CUSTOM"
)

// UserAccount.accountStatus constants.
const (
	AccountStatusEnabled  = "Enabled"
	AccountStatusDisabled = "Disabled"
)

// UserAccount.accountType constants.
const (
	AccountTypeDefault   = "DEFAULT"
	AccountTypeFederated = "FEDERATED"
)

var validAccessLevels = map[string]struct{}{
	AccessLevelFullAccess:       {},
	AccessLevelSiteAccess:       {},
	AccessLevelGroupBasedAccess: {},
}

var validPrivilegeLevels = map[string]struct{}{
	PrivilegeLevelAdministrator: {},
	PrivilegeLevelAuditor:       {},
	PrivilegeLevelEnrollment:    {},
	PrivilegeLevelCustom:        {},
}

var validAccountStatuses = map[string]struct{}{
	AccountStatusEnabled:  {},
	AccountStatusDisabled: {},
}

var validAccountTypes = map[string]struct{}{
	AccountTypeDefault:   {},
	AccountTypeFederated: {},
}
