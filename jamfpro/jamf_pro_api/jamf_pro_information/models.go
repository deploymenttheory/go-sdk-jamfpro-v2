package jamf_pro_information

// ResourceJamfProInformation represents Jamf Pro feature/capability flags.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-pro-information
type ResourceJamfProInformation struct {
	VppTokenEnabled         *bool `json:"vppTokenEnabled,omitempty"`
	DepAccountEnabled       *bool `json:"depAccountEnabled,omitempty"`
	ByodEnabled             *bool `json:"byodEnabled,omitempty"`
	UserMigrationEnabled    *bool `json:"userMigrationEnabled,omitempty"`
	CloudDeploymentsEnabled *bool `json:"cloudDeploymentsEnabled,omitempty"`
	PatchEnabled            *bool `json:"patchEnabled,omitempty"`
	SsoSamlEnabled          *bool `json:"ssoSamlEnabled,omitempty"`
	SmtpEnabled             *bool `json:"smtpEnabled,omitempty"`
}
