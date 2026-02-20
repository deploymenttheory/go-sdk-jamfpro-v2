package login_customization

// ResourceLoginCustomizationV1 represents the login customization resource (get/update).
type ResourceLoginCustomizationV1 struct {
	RampInstance            bool   `json:"rampInstance"`
	IncludeCustomDisclaimer bool   `json:"includeCustomDisclaimer"`
	DisclaimerHeading       string `json:"disclaimerHeading"`
	DisclaimerMainText      string `json:"disclaimerMainText"`
	ActionText              string `json:"actionText"`
}
