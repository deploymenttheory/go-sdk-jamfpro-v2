package user

// ResourceUser represents the current authenticated user information.
type ResourceUser struct {
	ID                int                    `json:"id"`
	Username          string                 `json:"username"`
	RealName          string                 `json:"realName"`
	Email             string                 `json:"email"`
	Preferences       map[string]interface{} `json:"preferences"`
	IsMultiSiteAdmin  bool                   `json:"isMultiSiteAdmin"`
	AccessLevel       string                 `json:"accessLevel"`
	PrivilegeSet      string                 `json:"privilegeSet"`
	PrivilegesBySite  map[string]interface{} `json:"privilegesBySite"`
	GroupIDs          []int                  `json:"groupIds"`
	CurrentSiteID     int                    `json:"currentSiteId"`
}

// RequestChangePassword is the request body for changing the current user's password.
type RequestChangePassword struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// RequestUpdateSession is the request body for updating the current user's session (change site).
type RequestUpdateSession struct {
	CurrentSiteID int `json:"currentSiteId"`
}
