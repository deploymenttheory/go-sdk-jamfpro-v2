package ldap_servers

import "encoding/xml"

// ResourceLDAPServer represents a Jamf Pro Classic API LDAP server resource.
type ResourceLDAPServer struct {
	XMLName          xml.Name   `xml:"ldap_server"`
	Connection       Connection `xml:"connection"`
	MappingsForUsers Mappings   `xml:"mappings_for_users"`
}

// ListItem represents an LDAP server item in a list response.
type ListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListLDAPServers (GET /JSSResource/ldapservers).
type ListResponse struct {
	XMLName xml.Name   `xml:"ldap_servers"`
	Size    int        `xml:"size"`
	Results []ListItem `xml:"ldap_server"`
}

// RequestLDAPServer is the body for creating or updating an LDAP server.
// The ID field is not included in Connection; it's specified via the URL path.
type RequestLDAPServer struct {
	XMLName          xml.Name          `xml:"ldap_server"`
	Connection       RequestConnection `xml:"connection"`
	MappingsForUsers Mappings          `xml:"mappings_for_users"`
}

// CreateResponse is the response for create operations which returns only ID and Name.
type CreateResponse struct {
	XMLName xml.Name `xml:"ldap_server"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

// Connection contains the connection settings for an LDAP server.
type Connection struct {
	ID                 int     `xml:"id"`
	Name               string  `xml:"name"`
	Hostname           string  `xml:"hostname"`
	ServerType         string  `xml:"server_type"`
	Port               int     `xml:"port"`
	UseSSL             bool    `xml:"use_ssl"`
	AuthenticationType string  `xml:"authentication_type"`
	Account            Account `xml:"account"`
	OpenCloseTimeout   int     `xml:"open_close_timeout"`
	SearchTimeout      int     `xml:"search_timeout"`
	ReferralResponse   string  `xml:"referral_response"`
	UseWildcards       bool    `xml:"use_wildcards"`
}

// RequestConnection is the connection section for create/update requests (excludes ID).
type RequestConnection struct {
	Name               string  `xml:"name"`
	Hostname           string  `xml:"hostname"`
	ServerType         string  `xml:"server_type"`
	Port               int     `xml:"port"`
	UseSSL             bool    `xml:"use_ssl"`
	AuthenticationType string  `xml:"authentication_type"`
	Account            Account `xml:"account"`
	OpenCloseTimeout   int     `xml:"open_close_timeout"`
	SearchTimeout      int     `xml:"search_timeout"`
	ReferralResponse   string  `xml:"referral_response"`
	UseWildcards       bool    `xml:"use_wildcards"`
}

// Account contains the credentials for LDAP server authentication.
type Account struct {
	DistinguishedUsername string `xml:"distinguished_username"`
	Password              string `xml:"password"`
}

// Mappings contains all LDAP mapping configurations.
type Mappings struct {
	UserMappings                UserMappings                `xml:"user_mappings"`
	UserGroupMappings           UserGroupMappings           `xml:"user_group_mappings"`
	UserGroupMembershipMappings UserGroupMembershipMappings `xml:"user_group_membership_mappings"`
}

// UserMappings contains the mappings for user attributes.
type UserMappings struct {
	MapObjectClassToAnyOrAll string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses            string `xml:"object_classes"`
	SearchBase               string `xml:"search_base"`
	SearchScope              string `xml:"search_scope"`
	MapUserID                string `xml:"map_user_id"`
	MapUsername              string `xml:"map_username"`
	MapRealName              string `xml:"map_realname"`
	MapEmailAddress          string `xml:"map_email_address"`
	AppendToEmailResults     string `xml:"append_to_email_results"`
	MapDepartment            string `xml:"map_department"`
	MapBuilding              string `xml:"map_building"`
	MapRoom                  string `xml:"map_room"`
	MapPhone                 string `xml:"map_phone"`
	MapPosition              string `xml:"map_position"`
	MapUserUUID              string `xml:"map_user_uuid"`
}

// UserGroupMappings contains the mappings for user group attributes.
type UserGroupMappings struct {
	MapObjectClassToAnyOrAll string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses            string `xml:"object_classes"`
	SearchBase               string `xml:"search_base"`
	SearchScope              string `xml:"search_scope"`
	MapGroupID               string `xml:"map_group_id"`
	MapGroupName             string `xml:"map_group_name"`
	MapGroupUUID             string `xml:"map_group_uuid"`
}

// UserGroupMembershipMappings contains the mappings for user group membership.
type UserGroupMembershipMappings struct {
	UserGroupMembershipStoredIn                      string `xml:"user_group_membership_stored_in"`
	MapGroupMembershipToUserField                    string `xml:"map_group_membership_to_user_field"`
	AppendToUsername                                 string `xml:"append_to_username"`
	UseDN                                            bool   `xml:"use_dn"`
	RecursiveLookups                                 bool   `xml:"recursive_lookups"`
	GroupMembershipEnabledWhenUserMembershipSelected bool   `xml:"group_membership_enabled_when_user_membership_selected"`
	MapUserMembershipToGroupField                    string `xml:"map_user_membership_to_group_field"`
	MapUserMembershipUseDN                           bool   `xml:"map_user_membership_use_dn"`
	MapObjectClassToAnyOrAll                         string `xml:"map_object_class_to_any_or_all"`
	ObjectClasses                                    string `xml:"object_classes"`
	SearchBase                                       string `xml:"search_base"`
	SearchScope                                      string `xml:"search_scope"`
	Username                                         string `xml:"username"`
	GroupID                                          string `xml:"group_id"`
	UserGroupMembershipUseLDAPCompare                bool   `xml:"user_group_membership_use_ldap_compare"`
	MembershipScopingOptimization                    bool   `xml:"membership_scoping_optimization"`
}
