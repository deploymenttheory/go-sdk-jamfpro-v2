package login_customization

// ResourceLoginCustomizationV1 represents the login customization resource (get/update).
// disclaimerHeading max 20 characters; disclaimerMainText and actionText are required on PUT.
type ResourceLoginCustomizationV1 struct {
	RampInstance            bool   `json:"rampInstance,omitempty"`
	IncludeCustomDisclaimer bool   `json:"includeCustomDisclaimer"`
	DisclaimerHeading       string `json:"disclaimerHeading,omitempty"` // max 20 chars
	DisclaimerMainText      string `json:"disclaimerMainText"`          // required on PUT
	ActionText              string `json:"actionText"`                  // required on PUT
}
