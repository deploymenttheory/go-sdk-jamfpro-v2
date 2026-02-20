package login_customization

// ResourceLoginCustomizationV1 represents the login customization resource (get/update).
type ResourceLoginCustomizationV1 struct {
	RampInstance            bool   `json:"rampInstance,omitempty"`
	IncludeCustomDisclaimer bool   `json:"includeCustomDisclaimer"`
	DisclaimerHeading       string `json:"disclaimerHeading,omitempty"`
	DisclaimerMainText      string `json:"disclaimerMainText,omitempty"`
	ActionText              string `json:"actionText,omitempty"`
}
