package mobile_device_prestages

// ListResponse represents the paginated response for mobile device prestages.
type ListResponse struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceMobileDevicePrestage `json:"results"`
}

// ResourceMobileDevicePrestage represents a mobile device prestage configuration.
type ResourceMobileDevicePrestage struct {
	DisplayName                            string                                          `json:"displayName"`
	Mandatory                              *bool                                           `json:"mandatory"`
	MdmRemovable                           *bool                                           `json:"mdmRemovable"`
	SupportPhoneNumber                     string                                          `json:"supportPhoneNumber"`
	SupportEmailAddress                    string                                          `json:"supportEmailAddress"`
	Department                             string                                          `json:"department"`
	DefaultPrestage                        *bool                                           `json:"defaultPrestage"`
	EnrollmentSiteID                       string                                          `json:"enrollmentSiteId"`
	KeepExistingSiteMembership             *bool                                           `json:"keepExistingSiteMembership"`
	KeepExistingLocationInformation        *bool                                           `json:"keepExistingLocationInformation"`
	RequireAuthentication                  *bool                                           `json:"requireAuthentication"`
	AuthenticationPrompt                   string                                          `json:"authenticationPrompt"`
	PreventActivationLock                  *bool                                           `json:"preventActivationLock"`
	EnableDeviceBasedActivationLock        *bool                                           `json:"enableDeviceBasedActivationLock"`
	DeviceEnrollmentProgramInstanceID      string                                          `json:"deviceEnrollmentProgramInstanceId"`
	SkipSetupItems                         SubsetSkipSetupItems                            `json:"skipSetupItems,omitempty"`
	LocationInformation                    SubsetLocationInformation                       `json:"locationInformation"`
	PurchasingInformation                  SubsetPurchasingInformation                     `json:"purchasingInformation"`
	AnchorCertificates                     []string                                        `json:"anchorCertificates,omitempty"`
	EnrollmentCustomizationID              string                                          `json:"enrollmentCustomizationId,omitempty"`
	Language                               string                                          `json:"language,omitempty"`
	Region                                 string                                          `json:"region,omitempty"`
	AutoAdvanceSetup                       *bool                                           `json:"autoAdvanceSetup"`
	AllowPairing                           *bool                                           `json:"allowPairing"`
	MultiUser                              *bool                                           `json:"multiUser"`
	Supervised                             *bool                                           `json:"supervised"`
	MaximumSharedAccounts                  int                                             `json:"maximumSharedAccounts"`
	ConfigureDeviceBeforeSetupAssistant    *bool                                           `json:"configureDeviceBeforeSetupAssistant"`
	Names                                  SubsetNames                                     `json:"names"`
	SendTimezone                           *bool                                           `json:"sendTimezone"`
	Timezone                               string                                          `json:"timezone"`
	StorageQuotaSizeMegabytes              int                                             `json:"storageQuotaSizeMegabytes"`
	UseStorageQuotaSize                    *bool                                           `json:"useStorageQuotaSize"`
	TemporarySessionOnly                   *bool                                           `json:"temporarySessionOnly"`
	EnforceTemporarySessionTimeout         *bool                                           `json:"enforceTemporarySessionTimeout"`
	TemporarySessionTimeout                *int                                            `json:"temporarySessionTimeout,omitempty"`
	EnforceUserSessionTimeout              *bool                                           `json:"enforceUserSessionTimeout"`
	UserSessionTimeout                     *int                                            `json:"userSessionTimeout,omitempty"`
	ID                                     string                                          `json:"id"`
	ProfileUuid                            string                                          `json:"profileUuid,omitempty"`
	SiteId                                 string                                          `json:"siteId,omitempty"`
	VersionLock                            int                                             `json:"versionLock"`
	PrestageMinimumOsTargetVersionTypeIos  string                                          `json:"prestageMinimumOsTargetVersionTypeIos,omitempty"`
	MinimumOsSpecificVersionIos            string                                          `json:"minimumOsSpecificVersionIos,omitempty"`
	PrestageMinimumOsTargetVersionTypeIpad string                                          `json:"prestageMinimumOsTargetVersionTypeIpad,omitempty"`
	MinimumOsSpecificVersionIpad           string                                          `json:"minimumOsSpecificVersionIpad,omitempty"`
	RTSEnabled                             *bool                                           `json:"rtsEnabled,omitempty"`
	RTSConfigProfileId                     string                                          `json:"rtsConfigProfileId,omitempty"`
	PreserveManagedApps                    *bool                                           `json:"preserveManagedApps,omitempty"`
}

// SubsetSkipSetupItems represents the setup items to skip during enrollment.
type SubsetSkipSetupItems struct {
	Location              *bool `json:"Location"`
	Privacy               *bool `json:"Privacy"`
	Biometric             *bool `json:"Biometric"`
	SoftwareUpdate        *bool `json:"SoftwareUpdate"`
	Diagnostics           *bool `json:"Diagnostics"`
	IMessageAndFaceTime   *bool `json:"iMessageAndFaceTime"`
	Intelligence          *bool `json:"Intelligence"`
	TVRoom                *bool `json:"TVRoom"`
	Passcode              *bool `json:"Passcode"`
	SIMSetup              *bool `json:"SIMSetup"`
	ScreenTime            *bool `json:"ScreenTime"`
	RestoreCompleted      *bool `json:"RestoreCompleted"`
	TVProviderSignIn      *bool `json:"TVProviderSignIn"`
	Siri                  *bool `json:"Siri"`
	Restore               *bool `json:"Restore"`
	ScreenSaver           *bool `json:"ScreenSaver"`
	HomeButtonSensitivity *bool `json:"HomeButtonSensitivity"`
	CloudStorage          *bool `json:"CloudStorage"`
	ActionButton          *bool `json:"ActionButton"`
	TransferData          *bool `json:"TransferData"`
	EnableLockdownMode    *bool `json:"EnableLockdownMode"`
	Zoom                  *bool `json:"Zoom"`
	PreferredLanguage     *bool `json:"PreferredLanguage"`
	VoiceSelection        *bool `json:"VoiceSelection"`
	TVHomeScreenSync      *bool `json:"TVHomeScreenSync"`
	Safety                *bool `json:"Safety"`
	TermsOfAddress        *bool `json:"TermsOfAddress"`
	ExpressLanguage       *bool `json:"ExpressLanguage"`
	CameraButton          *bool `json:"CameraButton"`
	AppleID               *bool `json:"AppleID"`
	DisplayTone           *bool `json:"DisplayTone"`
	WatchMigration        *bool `json:"WatchMigration"`
	UpdateCompleted       *bool `json:"UpdateCompleted"`
	Appearance            *bool `json:"Appearance"`
	Android               *bool `json:"Android"`
	Payment               *bool `json:"Payment"`
	OnBoarding            *bool `json:"OnBoarding"`
	TOS                   *bool `json:"TOS"`
	Welcome               *bool `json:"Welcome"`
	SafetyAndHandling     *bool `json:"SafetyAndHandling"`
	TapToSetup            *bool `json:"TapToSetup"`
	SpokenLanguage        *bool `json:"SpokenLanguage,omitempty"`
	Keyboard              *bool `json:"Keyboard,omitempty"`
	Multitasking          *bool `json:"Multitasking,omitempty"`
	OSShowcase            *bool `json:"OSShowcase,omitempty"`
}

// SubsetLocationInformation represents location information for the device.
type SubsetLocationInformation struct {
	Username     string `json:"username"`
	Realname     string `json:"realname"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Room         string `json:"room"`
	Position     string `json:"position"`
	DepartmentId string `json:"departmentId"`
	BuildingId   string `json:"buildingId"`
	ID           string `json:"id"`
	VersionLock  int    `json:"versionLock"`
}

// SubsetPurchasingInformation represents purchasing information for the device.
type SubsetPurchasingInformation struct {
	ID                string `json:"id"`
	Leased            *bool  `json:"leased"`
	Purchased         *bool  `json:"purchased"`
	AppleCareId       string `json:"appleCareId"`
	PoNumber          string `json:"poNumber"`
	Vendor            string `json:"vendor"`
	PurchasePrice     string `json:"purchasePrice"`
	LifeExpectancy    int    `json:"lifeExpectancy"`
	PurchasingAccount string `json:"purchasingAccount"`
	PurchasingContact string `json:"purchasingContact"`
	LeaseDate         string `json:"leaseDate"`
	PoDate            string `json:"poDate"`
	WarrantyDate      string `json:"warrantyDate"`
	VersionLock       int    `json:"versionLock"`
}

// SubsetNames represents device naming configuration.
type SubsetNames struct {
	AssignNamesUsing       string            `json:"assignNamesUsing"`
	PrestageDeviceNames    []SubsetNamesName `json:"prestageDeviceNames"`
	DeviceNamePrefix       string            `json:"deviceNamePrefix"`
	DeviceNameSuffix       string            `json:"deviceNameSuffix"`
	SingleDeviceName       string            `json:"singleDeviceName"`
	ManageNames            *bool             `json:"manageNames"`
	DeviceNamingConfigured *bool             `json:"deviceNamingConfigured"`
}

// SubsetNamesName represents a device name entry.
type SubsetNamesName struct {
	ID         string `json:"id"`
	DeviceName string `json:"deviceName"`
	Used       *bool  `json:"used"`
}

// CreateResponse represents the response when creating a mobile device prestage.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResourceDeviceScope represents the device scope for a mobile device prestage.
type ResourceDeviceScope struct {
	PrestageId  string                  `json:"prestageId"`
	Assignments []SubsetAssignmentItem  `json:"assignments"`
	VersionLock int                     `json:"versionLock"`
}

// SubsetAssignmentItem represents an assignment within the prestage scope.
type SubsetAssignmentItem struct {
	SerialNumber   string `json:"serialNumber"`
	AssignmentDate string `json:"assignmentDate"`
	UserAssigned   string `json:"userAssigned"`
}
