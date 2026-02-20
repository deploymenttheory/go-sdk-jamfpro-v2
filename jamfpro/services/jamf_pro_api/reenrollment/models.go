package reenrollment

// ResourceReenrollmentSettings is the re-enrollment settings resource (get/update).
type ResourceReenrollmentSettings struct {
	FlushPolicyHistory              bool   `json:"isFlushPolicyHistoryEnabled"`
	FlushLocationInformation        bool   `json:"isFlushLocationInformationEnabled"`
	FlushLocationInformationHistory bool   `json:"isFlushLocationInformationHistoryEnabled"`
	FlushExtensionAttributes        bool   `json:"isFlushExtensionAttributesEnabled"`
	FlushSoftwareUpdatePlans        bool   `json:"isFlushSoftwareUpdatePlansEnabled"`
	FlushMdmQueue                   string `json:"flushMDMQueue"`
}
