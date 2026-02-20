package cloud_ldap

// ResourceCloudLdap represents the Cloud LDAP configuration.
type ResourceCloudLdap struct {
	CloudIdPCommon *CloudIdPCommon    `json:"cloudIdPCommon"`
	Server         *CloudLdapServer   `json:"server"`
	Mappings       *CloudLdapMappings `json:"mappings,omitempty"`
}

// CloudIdPCommon contains common Cloud Identity Provider fields.
type CloudIdPCommon struct {
	ID           string `json:"id,omitempty"`
	ProviderName string `json:"providerName"`
	DisplayName  string `json:"displayName"`
}

// CloudLdapServer represents the LDAP server configuration.
type CloudLdapServer struct {
	Enabled                                  bool               `json:"enabled"`
	Keystore                                 *CloudLdapKeystore `json:"keystore"`
	UseWildcards                             bool               `json:"useWildcards"`
	ConnectionType                           string             `json:"connectionType"`
	ServerUrl                                string             `json:"serverUrl"`
	DomainName                               string             `json:"domainName"`
	Port                                     int                `json:"port"`
	ConnectionTimeout                        int                `json:"connectionTimeout"`
	SearchTimeout                            int                `json:"searchTimeout"`
	MembershipCalculationOptimizationEnabled bool               `json:"membershipCalculationOptimizationEnabled,omitempty"`
}

// CloudLdapKeystore represents the keystore configuration for LDAP.
type CloudLdapKeystore struct {
	Type           string `json:"type,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Subject        string `json:"subject,omitempty"`
	FileName       string `json:"fileName,omitempty"`
	Password       string `json:"password,omitempty"`
	FileBytes      string `json:"fileBytes,omitempty"`
}

// CloudLdapMappings defines field mappings for Cloud LDAP.
type CloudLdapMappings struct {
	UserMappings       CloudLdapUserMappings       `json:"userMappings"`
	GroupMappings      CloudLdapGroupMappings      `json:"groupMappings"`
	MembershipMappings CloudLdapMembershipMappings `json:"membershipMappings"`
}

// CloudLdapUserMappings defines user field mappings.
type CloudLdapUserMappings struct {
	ObjectClassLimitation string `json:"objectClassLimitation"`
	ObjectClasses         string `json:"objectClasses"`
	SearchBase            string `json:"searchBase"`
	SearchScope           string `json:"searchScope"`
	AdditionalSearchBase  string `json:"additionalSearchBase"`
	UserID                string `json:"userID"`
	Username              string `json:"username"`
	RealName              string `json:"realName"`
	EmailAddress          string `json:"emailAddress"`
	Department            string `json:"department"`
	Building              string `json:"building"`
	Room                  string `json:"room"`
	Phone                 string `json:"phone"`
	Position              string `json:"position"`
	UserUuid              string `json:"userUuid"`
}

// CloudLdapGroupMappings defines group field mappings.
type CloudLdapGroupMappings struct {
	ObjectClassLimitation string `json:"objectClassLimitation"`
	ObjectClasses         string `json:"objectClasses"`
	SearchBase            string `json:"searchBase"`
	SearchScope           string `json:"searchScope"`
	GroupID               string `json:"groupID"`
	GroupName             string `json:"groupName"`
	GroupUuid             string `json:"groupUuid"`
}

// CloudLdapMembershipMappings defines membership field mappings.
type CloudLdapMembershipMappings struct {
	GroupMembershipMapping string `json:"groupMembershipMapping"`
}

// ResponseCloudLdapCreated represents the response received after creating a Cloud LDAP configuration.
type ResponseCloudLdapCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseDefaultMappings represents the default mappings for Cloud LDAP.
type ResponseDefaultMappings struct {
	UserMappings       CloudLdapUserMappings       `json:"userMappings"`
	GroupMappings      CloudLdapGroupMappings      `json:"groupMappings"`
	MembershipMappings CloudLdapMembershipMappings `json:"membershipMappings"`
}

// ResponseDefaultServerConfiguration represents the default server configuration for Cloud LDAP.
type ResponseDefaultServerConfiguration struct {
	ID                                       string             `json:"id"`
	Enabled                                  bool               `json:"enabled"`
	ServerUrl                                string             `json:"serverUrl"`
	DomainName                               string             `json:"domainName"`
	Port                                     int                `json:"port"`
	Keystore                                 *CloudLdapKeystore `json:"keystore,omitempty"`
	ConnectionTimeout                        int                `json:"connectionTimeout"`
	SearchTimeout                            int                `json:"searchTimeout"`
	UseWildcards                             bool               `json:"useWildcards"`
	ConnectionType                           string             `json:"connectionType"`
	MembershipCalculationOptimizationEnabled bool               `json:"membershipCalculationOptimizationEnabled"`
}

// ConnectionPoolStats represents connection pool statistics for Cloud LDAP.
type ConnectionPoolStats struct {
	NumConnectionsClosedDefunct           int64 `json:"numConnectionsClosedDefunct"`
	NumConnectionsClosedExpired           int64 `json:"numConnectionsClosedExpired"`
	NumConnectionsClosedUnneeded          int64 `json:"numConnectionsClosedUnneeded"`
	NumFailedCheckouts                    int64 `json:"numFailedCheckouts"`
	NumFailedConnectionAttempts           int64 `json:"numFailedConnectionAttempts"`
	NumReleasedValid                      int64 `json:"numReleasedValid"`
	NumSuccessfulCheckouts                int64 `json:"numSuccessfulCheckouts"`
	NumSuccessfulCheckoutsNewConnection   int64 `json:"numSuccessfulCheckoutsNewConnection"`
	NumSuccessfulConnectionAttempts       int64 `json:"numSuccessfulConnectionAttempts"`
	MaximumAvailableConnections           int64 `json:"maximumAvailableConnections"`
	NumSuccessfulCheckoutsWithoutWaiting  int64 `json:"numSuccessfulCheckoutsWithoutWaiting"`
	NumSuccessfulCheckoutsAfterWaiting    int64 `json:"numSuccessfulCheckoutsAfterWaiting"`
	NumAvailableConnections               int64 `json:"numAvailableConnections"`
}

// ConnectionStatusResponse represents the connection status test result.
type ConnectionStatusResponse struct {
	Status string `json:"status"`
}
