package groups

type ListResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []ResourceGroup `json:"results"`
}

type ResourceGroup struct {
	GroupPlatformId  string              `json:"groupPlatformId,omitempty"`
	GroupJamfProId   string              `json:"groupJamfProId,omitempty"`
	GroupName        string              `json:"groupName,omitempty"`
	GroupDescription string              `json:"groupDescription,omitempty"`
	GroupType        string              `json:"groupType,omitempty"`
	Smart            bool                `json:"smart"`
	MembershipCount  int                 `json:"membershipCount"`
	Criteria         []SubsetCriterion   `json:"criteria,omitempty"`
	Assignments      []SubsetAssignment  `json:"assignments,omitempty"`
}

type SubsetCriterion struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen"`
	ClosingParen bool   `json:"closingParen"`
}

type SubsetAssignment struct {
	DeviceID string `json:"deviceId"`
	Selected bool   `json:"selected"`
}

type RequestUpdateGroup struct {
	GroupName        string             `json:"groupName,omitempty"`
	GroupDescription string             `json:"groupDescription,omitempty"`
	Criteria         []SubsetCriterion  `json:"criteria,omitempty"`
	Assignments      []SubsetAssignment `json:"assignments,omitempty"`
}
