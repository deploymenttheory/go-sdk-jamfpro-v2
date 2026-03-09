package api_authorization

// ResourceAuthV1 represents the current authorization details response from GET /api/v1/auth.
type ResourceAuthV1 struct {
	Account             AuthAccountV1     `json:"account"`
	AccountGroups       []AuthAccountGroupV1 `json:"accountGroups"`
	Sites               []AuthSiteV1      `json:"sites"`
	AuthenticationType  string            `json:"authenticationType"`
}

// AuthAccountV1 represents the account object within the auth response.
type AuthAccountV1 struct {
	ID                 string                    `json:"id"`
	Username           string                    `json:"username"`
	RealName           string                    `json:"realName"`
	Email              string                    `json:"email"`
	Preferences        AuthAccountPreferencesV1  `json:"preferences"`
	MultiSiteAdmin     bool                      `json:"multiSiteAdmin"`
	AccessLevel        string                    `json:"accessLevel"`
	PrivilegeSet       string                    `json:"privilegeSet"`
	PrivilegesBySite   map[string][]string      `json:"privilegesBySite"`
	GroupIDs           []int                     `json:"groupIds"`
	CurrentSiteID      string                    `json:"currentSiteId"`
}

// AuthAccountPreferencesV1 represents user preferences within the account.
type AuthAccountPreferencesV1 struct {
	Language             string `json:"language"`
	DateFormat           string `json:"dateFormat"`
	Region               string `json:"region"`
	Timezone             string `json:"timezone"`
	DisableRelativeDates bool   `json:"disableRelativeDates"`
}

// AuthAccountGroupV1 represents an account group within the auth response.
type AuthAccountGroupV1 struct {
	AccessLevel   string `json:"accessLevel"`
	PrivilegeSet  string `json:"privilegeSet"`
	SiteID        int    `json:"siteId"`
	Privileges    string `json:"privileges"`
	MemberUserIDs []int  `json:"memberUserIds"`
}

// AuthSiteV1 represents a site within the auth response.
type AuthSiteV1 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
