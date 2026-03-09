package accounts

// ResourceAccount represents a Jamf Pro user account.
type ResourceAccount struct {
	ID                       string `json:"id"`
	Username                 string `json:"username"`
	Realname                 string `json:"realname"`
	Email                    string `json:"email"`
	Phone                    string `json:"phone"`
	LdapServerID             int    `json:"ldapServerId"`
	DistinguishedName        string `json:"distinguishedName"`
	SiteID                   int    `json:"siteId"`
	AccessLevel              string `json:"accessLevel"`              // FullAccess, SiteAccess, GroupBasedAccess
	PrivilegeLevel           string `json:"privilegeLevel"`           // ADMINISTRATOR, AUDITOR, ENROLLMENT, CUSTOM
	LastPasswordChange       string `json:"lastPasswordChange"`       // ISO 8601 timestamp
	ChangePasswordOnNextLogin bool   `json:"changePasswordOnNextLogin"`
	FailedLoginAttempts      int    `json:"failedLoginAttempts"`
	AccountStatus            string `json:"accountStatus"`            // Enabled, Disabled
	AccountType              string `json:"accountType"`              // DEFAULT, FEDERATED
}

// ListResponse is the response for listing user accounts.
type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	Results    []ResourceAccount `json:"results"`
}

// RequestAccount is the body for creating and updating accounts.
type RequestAccount struct {
	PlainPassword             string `json:"plainPassword,omitempty"`
	Password                  string `json:"password,omitempty"`
	Username                  string `json:"username,omitempty"`
	Realname                  string `json:"realname,omitempty"`
	Email                     string `json:"email,omitempty"`
	Phone                     string `json:"phone,omitempty"`
	LdapServerID              int    `json:"ldapServerId,omitempty"`
	DistinguishedName         string `json:"distinguishedName,omitempty"`
	SiteID                    int    `json:"siteId,omitempty"`
	AccessLevel               string `json:"accessLevel,omitempty"`              // FullAccess, SiteAccess, GroupBasedAccess
	PrivilegeLevel            string `json:"privilegeLevel,omitempty"`           // ADMINISTRATOR, AUDITOR, ENROLLMENT, CUSTOM
	ChangePasswordOnNextLogin bool   `json:"changePasswordOnNextLogin,omitempty"`
	AccountStatus             string `json:"accountStatus,omitempty"`            // Enabled, Disabled
	AccountType               string `json:"accountType,omitempty"`              // DEFAULT, FEDERATED
}

// CreateResponse is the response for creating an account.
type CreateResponse struct {
	ID                       string `json:"id"`
	Username                 string `json:"username"`
	Realname                 string `json:"realname"`
	Email                    string `json:"email"`
	Phone                    string `json:"phone"`
	LdapServerID             int    `json:"ldapServerId"`
	DistinguishedName        string `json:"distinguishedName"`
	SiteID                   int    `json:"siteId"`
	AccessLevel              string `json:"accessLevel"`
	PrivilegeLevel           string `json:"privilegeLevel"`
	LastPasswordChange       string `json:"lastPasswordChange"`       // ISO 8601 timestamp
	ChangePasswordOnNextLogin bool   `json:"changePasswordOnNextLogin"`
	FailedLoginAttempts      int    `json:"failedLoginAttempts"`
	AccountStatus            string `json:"accountStatus"`
	AccountType              string `json:"accountType"`
}
