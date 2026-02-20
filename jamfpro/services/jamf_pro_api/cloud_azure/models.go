package cloud_azure

// ResourceCloudAzure represents the Azure Cloud Identity Provider configuration.
type ResourceCloudAzure struct {
	CloudIdPCommon CloudIdPCommon       `json:"cloudIdPCommon"`
	Server         CloudAzureServer     `json:"server"`
}

// CloudIdPCommon contains common Cloud Identity Provider fields.
type CloudIdPCommon struct {
	DisplayName  string `json:"displayName"`
	ProviderName string `json:"providerName"`
}

// CloudAzureServer represents the Azure-specific server configuration.
type CloudAzureServer struct {
	ID                                       string                  `json:"id"`
	TenantId                                 string                  `json:"tenantId"`
	Enabled                                  bool                    `json:"enabled"`
	Migrated                                 bool                    `json:"migrated"`
	Mappings                                 CloudAzureServerMappings `json:"mappings"`
	SearchTimeout                            int                     `json:"searchTimeout"`
	TransitiveMembershipEnabled              bool                    `json:"transitiveMembershipEnabled"`
	TransitiveMembershipUserField            string                  `json:"transitiveMembershipUserField"`
	TransitiveDirectoryMembershipEnabled     bool                    `json:"transitiveDirectoryMembershipEnabled"`
	MembershipCalculationOptimizationEnabled bool                    `json:"membershipCalculationOptimizationEnabled"`
	Code                                     string                  `json:"code"`
}

// CloudAzureServerMappings defines field mappings for Azure Cloud IDP.
type CloudAzureServerMappings struct {
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	RealName   string `json:"realName"`
	Email      string `json:"email"`
	Department string `json:"department"`
	Building   string `json:"building"`
	Room       string `json:"room"`
	Phone      string `json:"phone"`
	Position   string `json:"position"`
	GroupId    string `json:"groupId"`
	GroupName  string `json:"groupName"`
}

// ResponseCloudAzureCreated represents the response received after creating an Azure Cloud Identity Provider.
type ResponseCloudAzureCreated struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
