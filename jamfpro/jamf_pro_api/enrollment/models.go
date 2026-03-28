package enrollment

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ResourceEnrollment represents the enrollment settings (API v4).
type ResourceEnrollment struct {
	InstallSingleProfile                         bool                           `json:"installSingleProfile"`
	SigningMdmProfileEnabled                     bool                           `json:"signingMdmProfileEnabled"`
	MdmSigningCertificate                        *ResourceEnrollmentCertificate `json:"mdmSigningCertificate"`
	RestrictReenrollment                         bool                           `json:"restrictReenrollment"`
	FlushLocationInformation                     bool                           `json:"flushLocationInformation"`
	FlushLocationHistoryInformation              bool                           `json:"flushLocationHistoryInformation"`
	FlushPolicyHistory                           bool                           `json:"flushPolicyHistory"`
	FlushExtensionAttributes                     bool                           `json:"flushExtensionAttributes"`
	FlushSoftwareUpdatePlans                     bool                           `json:"flushSoftwareUpdatePlans"`
	FlushMdmCommandsOnReenroll                   string                         `json:"flushMdmCommandsOnReenroll,omitempty"`
	MacOsEnterpriseEnrollmentEnabled             bool                           `json:"macOsEnterpriseEnrollmentEnabled"`
	ManagementUsername                           string                         `json:"managementUsername"`
	CreateManagementAccount                      bool                           `json:"createManagementAccount"`
	HideManagementAccount                        bool                           `json:"hideManagementAccount"`
	AllowSshOnlyManagementAccount                bool                           `json:"allowSshOnlyManagementAccount"`
	EnsureSshRunning                             bool                           `json:"ensureSshRunning"`
	LaunchSelfService                            bool                           `json:"launchSelfService"`
	SignQuickAdd                                 bool                           `json:"signQuickAdd"`
	DeveloperCertificateIdentity                 *ResourceEnrollmentCertificate `json:"developerCertificateIdentity"`
	DeveloperCertificateIdentityDetails          ResourceCertificateDetails     `json:"developerCertificateIdentityDetails"`
	MdmSigningCertificateDetails                 ResourceCertificateDetails     `json:"mdmSigningCertificateDetails"`
	IosEnterpriseEnrollmentEnabled               bool                           `json:"iosEnterpriseEnrollmentEnabled"`
	IosPersonalEnrollmentEnabled                 bool                           `json:"iosPersonalEnrollmentEnabled"`
	PersonalDeviceEnrollmentType                 string                         `json:"personalDeviceEnrollmentType,omitempty"`
	AccountDrivenUserEnrollmentEnabled           bool                           `json:"accountDrivenUserEnrollmentEnabled"`
	AccountDrivenDeviceIosEnrollmentEnabled      bool                           `json:"accountDrivenDeviceIosEnrollmentEnabled"`
	AccountDrivenDeviceMacosEnrollmentEnabled    bool                           `json:"accountDrivenDeviceMacosEnrollmentEnabled"`
	AccountDrivenUserVisionosEnrollmentEnabled   bool                           `json:"accountDrivenUserVisionosEnrollmentEnabled"`
	AccountDrivenDeviceVisionosEnrollmentEnabled bool                           `json:"accountDrivenDeviceVisionosEnrollmentEnabled"`
	MaidUsernameMergeEnabled                     bool                           `json:"maidUsernameMergeEnabled"`
}

// ResourceEnrollmentCertificate represents enrollment certificate details.
type ResourceEnrollmentCertificate struct {
	Filename         string `json:"filename"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
	IdentityKeystore string `json:"identityKeystore,omitempty"`
	Md5Sum           string `json:"md5Sum,omitempty"`
}

// ResourceCertificateDetails represents certificate subject and serial number.
type ResourceCertificateDetails struct {
	Subject      string `json:"subject"`
	SerialNumber string `json:"serialNumber"`
}

// ResourceAccountDrivenUserEnrollmentAccessGroup represents an ADUE access group (API v3).
type ResourceAccountDrivenUserEnrollmentAccessGroup struct {
	ID                                 string `json:"id"`
	GroupID                            string `json:"groupId"`
	LdapServerID                       string `json:"ldapServerId"`
	Name                               string `json:"name"`
	SiteID                             string `json:"siteId"`
	EnterpriseEnrollmentEnabled        bool   `json:"enterpriseEnrollmentEnabled"`
	PersonalEnrollmentEnabled          bool   `json:"personalEnrollmentEnabled"`
	AccountDrivenUserEnrollmentEnabled bool   `json:"accountDrivenUserEnrollmentEnabled"`
	RequireEula                        bool   `json:"requireEula"`
}

// ResourceEnrollmentLanguage represents enrollment language messaging (API v3).
type ResourceEnrollmentLanguage struct {
	LanguageCode                     string `json:"languageCode"`
	Name                             string `json:"name"`
	Title                            string `json:"title"`
	LoginDescription                 string `json:"loginDescription"`
	Username                         string `json:"username"`
	Password                         string `json:"password"`
	LoginButton                      string `json:"loginButton"`
	DeviceClassDescription           string `json:"deviceClassDescription"`
	DeviceClassPersonal              string `json:"deviceClassPersonal"`
	DeviceClassPersonalDescription   string `json:"deviceClassPersonalDescription"`
	DeviceClassEnterprise            string `json:"deviceClassEnterprise"`
	DeviceClassEnterpriseDescription string `json:"deviceClassEnterpriseDescription"`
	DeviceClassButton                string `json:"deviceClassButton"`
	PersonalEula                     string `json:"personalEula"`
	EnterpriseEula                   string `json:"enterpriseEula"`
	EulaButton                       string `json:"eulaButton"`
	SiteDescription                  string `json:"siteDescription"`
	CertificateText                  string `json:"certificateText"`
	CertificateButton                string `json:"certificateButton"`
	CertificateProfileName           string `json:"certificateProfileName"`
	CertificateProfileDescription    string `json:"certificateProfileDescription"`
	// Deprecated: PersonalText is deprecated as of Jamf Pro 11.25 and always returns an empty string.
	PersonalText string `json:"personalText"`
	// Deprecated: PersonalButton is deprecated as of Jamf Pro 11.25 and always returns an empty string.
	PersonalButton string `json:"personalButton"`
	// Deprecated: PersonalProfileName is deprecated as of Jamf Pro 11.25 and always returns an empty string.
	PersonalProfileName string `json:"personalProfileName"`
	// Deprecated: PersonalProfileDescription is deprecated as of Jamf Pro 11.25 and always returns an empty string.
	PersonalProfileDescription string `json:"personalProfileDescription"`
	UserEnrollmentText               string `json:"userEnrollmentText"`
	UserEnrollmentButton             string `json:"userEnrollmentButton"`
	UserEnrollmentProfileName        string `json:"userEnrollmentProfileName"`
	UserEnrollmentProfileDescription string `json:"userEnrollmentProfileDescription"`
	EnterpriseText                   string `json:"enterpriseText"`
	EnterpriseButton                 string `json:"enterpriseButton"`
	EnterpriseProfileName            string `json:"enterpriseProfileName"`
	EnterpriseProfileDescription     string `json:"enterpriseProfileDescription"`
	EnterprisePending                string `json:"enterprisePending"`
	QuickAddText                     string `json:"quickAddText"`
	QuickAddButton                   string `json:"quickAddButton"`
	QuickAddName                     string `json:"quickAddName"`
	QuickAddPending                  string `json:"quickAddPending"`
	CompleteMessage                  string `json:"completeMessage"`
	FailedMessage                    string `json:"failedMessage"`
	TryAgainButton                   string `json:"tryAgainButton"`
	CheckNowButton                   string `json:"checkNowButton"`
	CheckEnrollmentMessage           string `json:"checkEnrollmentMessage"`
	LogoutButton                     string `json:"logoutButton"`
}

// ResourceLanguageCode represents an available language code (API v3).
type ResourceLanguageCode struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// HistoryObject is an alias to the shared history item struct.
type HistoryObject = models.SharedHistoryItem

// HistoryResponse is an alias to the shared history response struct.
type HistoryResponse = models.SharedHistoryResponse

// ListResponseAccessGroups is the response for ListAccessGroupsV3.
type ListResponseAccessGroups struct {
	TotalCount int                                              `json:"totalCount"`
	Results    []ResourceAccountDrivenUserEnrollmentAccessGroup `json:"results"`
}

// ListResponseLanguageMessages is the response for ListLanguageMessagesV3.
type ListResponseLanguageMessages struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceEnrollmentLanguage `json:"results"`
}

// CreateResponse is the response for CreateAccessGroupV3.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// RequestDeleteMultipleLanguages represents the request body for deleting multiple languages (API v3).
type RequestDeleteMultipleLanguages struct {
	IDs []string `json:"ids"`
}

// ResourceADUESessionTokenSettings represents ADUE (Account Driven User Enrollment) session token settings (API v1).
// Used by GetADUESessionTokenSettingsV1 and UpdateADUESessionTokenSettingsV1.
type ResourceADUESessionTokenSettings struct {
	Enabled                   bool `json:"enabled"`
	ExpirationIntervalDays    int  `json:"expirationIntervalDays,omitempty"`
	ExpirationIntervalSeconds int  `json:"expirationIntervalSeconds,omitempty"`
}

// RequestAddHistoryNotes is an alias to the shared history note request struct.
type RequestAddHistoryNotes = models.SharedHistoryNoteRequest

// RequestExportHistory represents the optional request body for exporting enrollment history (API v2).
// Overrides query parameters when URI would exceed ~2k characters.
type RequestExportHistory struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []string `json:"fields,omitempty"`
}
