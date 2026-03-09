package classic_ldap

// ResourceOnPremLdapMappingsV1 represents the LDAP attribute mappings for an OnPrem LDAP configuration.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-classic-ldap-id
type ResourceOnPremLdapMappingsV1 struct {
	UserObjectMapIdTo         string `json:"userObjectMapIdTo"`
	UserObjectMapUsernameTo   string `json:"userObjectMapUsernameTo"`
	UserObjectMapRealNameTo   string `json:"userObjectMapRealNameTo"`
	UserObjectMapEmailTo      string `json:"userObjectMapEmailTo"`
	UserObjectMapDepartmentTo string `json:"userObjectMapDepartmentTo"`
	UserObjectMapBuildingTo   string `json:"userObjectMapBuildingTo"`
	UserObjectMapRoomTo       string `json:"userObjectMapRoomTo"`
	UserObjectMapPhoneTo      string `json:"userObjectMapPhoneTo"`
	UserObjectMapPositionTo   string `json:"userObjectMapPositionTo"`
	UserObjectMapUuidTo       string `json:"userObjectMapUuidTo"`
	UserGroupObjectMapIdTo    string `json:"userGroupObjectMapIdTo"`
	UserGroupObjectMapGroupNameTo string `json:"userGroupObjectMapGroupNameTo"`
	UserGroupObjectMapUuidTo  string `json:"userGroupObjectMapUuidTo"`
}
