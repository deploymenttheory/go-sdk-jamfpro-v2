package enrollment_settings

// ResourceEnrollmentSettingsV4 represents the enrollment settings resource (v4).
type ResourceEnrollmentSettingsV4 struct {
	InstallSingleProfile                         bool                                `json:"installSingleProfile"`
	SigningMdmProfileEnabled                     bool                                `json:"signingMdmProfileEnabled"`
	MdmSigningCertificate                        *EnrollmentSubsetCertificate        `json:"mdmSigningCertificate"`
	MdmSigningCertificateDetails                 EnrollmentSubsetCertificateDetails  `json:"mdmSigningCertificateDetails"`
	RestrictReenrollment                         bool                                `json:"restrictReenrollment"`
	FlushLocationInformation                     bool                                `json:"flushLocationInformation"`
	FlushLocationHistoryInformation              bool                                `json:"flushLocationHistoryInformation"`
	FlushPolicyHistory                           bool                                `json:"flushPolicyHistory"`
	FlushExtensionAttributes                     bool                                `json:"flushExtensionAttributes"`
	FlushSoftwareUpdatePlans                     bool                                `json:"flushSoftwareUpdatePlans"`
	FlushMdmCommandsOnReenroll                   string                              `json:"flushMdmCommandsOnReenroll"`
	MacOsEnterpriseEnrollmentEnabled             bool                                `json:"macOsEnterpriseEnrollmentEnabled"`
	ManagementUsername                           string                              `json:"managementUsername"`
	CreateManagementAccount                      bool                                `json:"createManagementAccount"`
	HideManagementAccount                        bool                                `json:"hideManagementAccount"`
	AllowSshOnlyManagementAccount                bool                                `json:"allowSshOnlyManagementAccount"`
	EnsureSshRunning                             bool                                `json:"ensureSshRunning"`
	LaunchSelfService                            bool                                `json:"launchSelfService"`
	SignQuickAdd                                 bool                                `json:"signQuickAdd"`
	DeveloperCertificateIdentity                 *EnrollmentSubsetCertificate        `json:"developerCertificateIdentity"`
	DeveloperCertificateIdentityDetails          EnrollmentSubsetCertificateDetails  `json:"developerCertificateIdentityDetails"`
	IosEnterpriseEnrollmentEnabled               bool                                `json:"iosEnterpriseEnrollmentEnabled"`
	IosPersonalEnrollmentEnabled                 bool                                `json:"iosPersonalEnrollmentEnabled"`
	PersonalDeviceEnrollmentType                 string                              `json:"personalDeviceEnrollmentType"`
	AccountDrivenUserEnrollmentEnabled           bool                                `json:"accountDrivenUserEnrollmentEnabled"`
	AccountDrivenDeviceIosEnrollmentEnabled      bool                                `json:"accountDrivenDeviceIosEnrollmentEnabled"`
	AccountDrivenDeviceMacosEnrollmentEnabled    bool                                `json:"accountDrivenDeviceMacosEnrollmentEnabled"`
	AccountDrivenUserVisionosEnrollmentEnabled   bool                                `json:"accountDrivenUserVisionosEnrollmentEnabled"`
	AccountDrivenDeviceVisionosEnrollmentEnabled bool                                `json:"accountDrivenDeviceVisionosEnrollmentEnabled"`
	MaidUsernameMergeEnabled                     bool                                `json:"maidUsernameMergeEnabled"`
}

// EnrollmentSubsetCertificate represents a certificate reference.
type EnrollmentSubsetCertificate struct {
	Filename string `json:"filename"`
	Md5Sum   string `json:"md5Sum"`
}

// EnrollmentSubsetCertificateDetails represents certificate details.
type EnrollmentSubsetCertificateDetails struct {
	Subject      string `json:"subject"`
	SerialNumber string `json:"serialNumber"`
}
