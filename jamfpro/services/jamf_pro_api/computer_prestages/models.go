package computer_prestages

// ListResponse is the response for ListV3.
type ListResponse struct {
	TotalCount int                       `json:"totalCount"`
	Results    []ResourceComputerPrestage `json:"results"`
}

// ResourceDeviceScope is the device scope for a computer prestage (GET /api/v2/computer-prestages/{id}/scope).
type ResourceDeviceScope struct {
	PrestageId  string                            `json:"prestageId"`
	Assignments []DeviceScopeSubsetAssignmentItem `json:"assignments"`
	VersionLock int                               `json:"versionLock"`
}

// ReplaceDeviceScopeRequest is the request body for ReplaceDeviceScopeByIDV2 (PUT /api/v2/computer-prestages/{id}/scope).
type ReplaceDeviceScopeRequest struct {
	SerialNumbers []string `json:"serialNumbers"`
	VersionLock  int      `json:"versionLock"`
}

// DeviceScopeSubsetAssignmentItem represents an assignment within the prestage scope.
type DeviceScopeSubsetAssignmentItem struct {
	SerialNumber   string `json:"serialNumber"`
	AssignmentDate string `json:"assignmentDate"`
	UserAssigned   string `json:"userAssigned"`
}

// CreateResponse is the response for CreateV3 (id and href).
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceComputerPrestage represents a computer prestage (v3 API).
type ResourceComputerPrestage struct {
	ID                                             string                                       `json:"id,omitempty"`
	VersionLock                                    int                                          `json:"versionLock,omitempty"`
	DisplayName                                    string                                       `json:"displayName,omitempty"`
	Mandatory                                      *bool                                        `json:"mandatory,omitempty"`
	MDMRemovable                                   *bool                                        `json:"mdmRemovable,omitempty"`
	SupportPhoneNumber                             string                                       `json:"supportPhoneNumber,omitempty"`
	SupportEmailAddress                            string                                       `json:"supportEmailAddress,omitempty"`
	Department                                     string                                       `json:"department,omitempty"`
	DefaultPrestage                                *bool                                        `json:"defaultPrestage,omitempty"`
	EnrollmentSiteId                               string                                       `json:"enrollmentSiteId,omitempty"`
	KeepExistingSiteMembership                     *bool                                        `json:"keepExistingSiteMembership,omitempty"`
	KeepExistingLocationInformation                *bool                                        `json:"keepExistingLocationInformation,omitempty"`
	RequireAuthentication                          *bool                                        `json:"requireAuthentication,omitempty"`
	AuthenticationPrompt                           string                                       `json:"authenticationPrompt,omitempty"`
	PreventActivationLock                           *bool                                       `json:"preventActivationLock,omitempty"`
	EnableDeviceBasedActivationLock                 *bool                                        `json:"enableDeviceBasedActivationLock,omitempty"`
	DeviceEnrollmentProgramInstanceId               string                                       `json:"deviceEnrollmentProgramInstanceId,omitempty"`
	SkipSetupItems                                 *SkipSetupItems                              `json:"skipSetupItems,omitempty"`
	LocationInformation                            LocationInformation                          `json:"locationInformation,omitempty"`
	PurchasingInformation                          PurchasingInformation                       `json:"purchasingInformation,omitempty"`
	AnchorCertificates                             []string                                     `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationId                       string                                       `json:"enrollmentCustomizationId,omitempty"`
	Language                                        string                                       `json:"language,omitempty"`
	Region                                          string                                       `json:"region,omitempty"`
	AutoAdvanceSetup                                *bool                                        `json:"autoAdvanceSetup,omitempty"`
	InstallProfilesDuringSetup                      *bool                                        `json:"installProfilesDuringSetup,omitempty"`
	PrestageInstalledProfileIds                     []string                                     `json:"prestageInstalledProfileIds,omitempty"`
	CustomPackageIds                                []string                                     `json:"customPackageIds,omitempty"`
	CustomPackageDistributionPointId                string                                       `json:"customPackageDistributionPointId,omitempty"`
	EnableRecoveryLock                              *bool                                        `json:"enableRecoveryLock,omitempty"`
	RecoveryLockPasswordType                        string                                       `json:"recoveryLockPasswordType,omitempty"`
	RecoveryLockPassword                            string                                       `json:"recoveryLockPassword,omitempty"`
	RotateRecoveryLockPassword                      *bool                                        `json:"rotateRecoveryLockPassword,omitempty"`
	PrestageMinimumOsTargetVersionType              string                                       `json:"prestageMinimumOsTargetVersionType,omitempty"`
	MinimumOsSpecificVersion                        string                                       `json:"minimumOsSpecificVersion,omitempty"`
	ProfileUuid                                     string                                       `json:"profileUuid,omitempty"`
	SiteId                                          string                                       `json:"siteId,omitempty"`
	AccountSettings                                 *AccountSettings                             `json:"accountSettings,omitempty"`
	Enabled                                         *bool                                        `json:"enabled,omitempty"`
	SsoForEnrollmentEnabled                         *bool                                        `json:"ssoForEnrollmentEnabled,omitempty"`
	SsoBypassAllowed                                *bool                                        `json:"ssoBypassAllowed,omitempty"`
	SsoEnabled                                      *bool                                        `json:"ssoEnabled,omitempty"`
	SsoForMacOsSelfServiceEnabled                   *bool                                        `json:"ssoForMacOsSelfServiceEnabled,omitempty"`
	TokenExpirationDisabled                         *bool                                        `json:"tokenExpirationDisabled,omitempty"`
	UserAttributeEnabled                            *bool                                        `json:"userAttributeEnabled,omitempty"`
	UserAttributeName                               string                                       `json:"userAttributeName,omitempty"`
	UserMapping                                     string                                       `json:"userMapping,omitempty"`
	EnrollmentSsoForAccountDrivenEnrollmentEnabled *bool                                        `json:"enrollmentSsoForAccountDrivenEnrollmentEnabled,omitempty"`
	GroupEnrollmentAccessEnabled                    *bool                                        `json:"groupEnrollmentAccessEnabled,omitempty"`
	GroupAttributeName                              string                                       `json:"groupAttributeName,omitempty"`
	GroupRdnKey                                     string                                       `json:"groupRdnKey,omitempty"`
	GroupEnrollmentAccessName                        string                                       `json:"groupEnrollmentAccessName,omitempty"`
	IdpProviderType                                 string                                       `json:"idpProviderType,omitempty"`
	OtherProviderTypeName                           string                                       `json:"otherProviderTypeName,omitempty"`
	MetadataSource                                  string                                       `json:"metadataSource,omitempty"`
	SessionTimeout                                  int                                          `json:"sessionTimeout,omitempty"`
	DeviceType                                      string                                       `json:"deviceType,omitempty"`
	OnboardingItems                                 []OnboardingItem                             `json:"onboardingItems,omitempty"`
	PssoEnabled                                     *bool                                        `json:"pssoEnabled,omitempty"`
	PlatformSsoAppBundleId                          string                                       `json:"platformSsoAppBundleId,omitempty"`
}

// SkipSetupItems represents skip setup items subset.
type SkipSetupItems struct {
	Biometric                 *bool `json:"Biometric,omitempty"`
	TermsOfAddress            *bool `json:"TermsOfAddress,omitempty"`
	FileVault                 *bool `json:"FileVault,omitempty"`
	ICloudDiagnostics         *bool `json:"iCloudDiagnostics,omitempty"`
	Diagnostics               *bool `json:"Diagnostics,omitempty"`
	Accessibility             *bool `json:"Accessibility,omitempty"`
	AppleID                   *bool `json:"AppleID,omitempty"`
	ScreenTime                *bool `json:"ScreenTime,omitempty"`
	Siri                      *bool `json:"Siri,omitempty"`
	DisplayTone               *bool `json:"DisplayTone,omitempty"`
	Restore                   *bool `json:"Restore,omitempty"`
	Appearance                *bool `json:"Appearance,omitempty"`
	Privacy                   *bool `json:"Privacy,omitempty"`
	Payment                   *bool `json:"Payment,omitempty"`
	Registration              *bool `json:"Registration,omitempty"`
	TOS                       *bool `json:"TOS,omitempty"`
	ICloudStorage             *bool `json:"iCloudStorage,omitempty"`
	Location                  *bool `json:"Location,omitempty"`
	Intelligence              *bool `json:"Intelligence,omitempty"`
	EnableLockdownMode        *bool `json:"EnableLockdownMode,omitempty"`
	Welcome                   *bool `json:"Welcome,omitempty"`
	Wallpaper                 *bool `json:"Wallpaper,omitempty"`
	SoftwareUpdate            *bool `json:"SoftwareUpdate,omitempty"`
	AdditionalPrivacySettings *bool `json:"AdditionalPrivacySettings,omitempty"`
	OSShowcase                *bool `json:"OSShowcase,omitempty"`
}

// LocationInformation represents location information subset.
type LocationInformation struct {
	ID           string `json:"id,omitempty"`
	VersionLock  int    `json:"versionLock,omitempty"`
	Username     string `json:"username,omitempty"`
	Realname     string `json:"realname,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Email        string `json:"email,omitempty"`
	Room         string `json:"room,omitempty"`
	Position     string `json:"position,omitempty"`
	DepartmentId string `json:"departmentId,omitempty"`
	BuildingId   string `json:"buildingId,omitempty"`
}

// PurchasingInformation represents purchasing information subset.
type PurchasingInformation struct {
	ID                string `json:"id,omitempty"`
	VersionLock        int    `json:"versionLock,omitempty"`
	Leased            *bool  `json:"leased,omitempty"`
	Purchased         *bool  `json:"purchased,omitempty"`
	AppleCareId       string `json:"appleCareId,omitempty"`
	PONumber          string `json:"poNumber,omitempty"`
	Vendor            string `json:"vendor,omitempty"`
	PurchasePrice     string `json:"purchasePrice,omitempty"`
	LifeExpectancy    int    `json:"lifeExpectancy,omitempty"`
	PurchasingAccount string `json:"purchasingAccount,omitempty"`
	PurchasingContact string `json:"purchasingContact,omitempty"`
	LeaseDate         string `json:"leaseDate,omitempty"`
	PODate            string `json:"poDate,omitempty"`
	WarrantyDate      string `json:"warrantyDate,omitempty"`
}

// AccountSettings represents account settings subset.
type AccountSettings struct {
	ID                                      string `json:"id,omitempty"`
	VersionLock                             int    `json:"versionLock,omitempty"`
	PayloadConfigured                       *bool  `json:"payloadConfigured,omitempty"`
	LocalAdminAccountEnabled                *bool  `json:"localAdminAccountEnabled,omitempty"`
	AdminUsername                           string `json:"adminUsername,omitempty"`
	AdminPassword                           string `json:"adminPassword,omitempty"`
	HiddenAdminAccount                      *bool  `json:"hiddenAdminAccount,omitempty"`
	LocalUserManaged                        *bool  `json:"localUserManaged,omitempty"`
	UserAccountType                         string `json:"userAccountType,omitempty"`
	PrefillPrimaryAccountInfoFeatureEnabled *bool  `json:"prefillPrimaryAccountInfoFeatureEnabled,omitempty"`
	PrefillType                             string `json:"prefillType,omitempty"`
	PrefillAccountFullName                  string `json:"prefillAccountFullName,omitempty"`
	PrefillAccountUserName                  string `json:"prefillAccountUserName,omitempty"`
	PreventPrefillInfoFromModification      *bool  `json:"preventPrefillInfoFromModification,omitempty"`
}

// OnboardingItem represents an onboarding item.
type OnboardingItem struct {
	SelfServiceEntityType string `json:"selfServiceEntityType,omitempty"`
	ID                    string `json:"id,omitempty"`
	EntityId              string `json:"entityId,omitempty"`
	Priority              int    `json:"priority,omitempty"`
}
