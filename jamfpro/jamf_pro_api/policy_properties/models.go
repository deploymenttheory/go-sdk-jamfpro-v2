package policy_properties

// ResourcePolicyProperties is the policy properties resource (get/update).
type ResourcePolicyProperties struct {
	PoliciesRequireNetworkStateChange bool `json:"policiesRequireNetworkStateChange"`
	AllowNetworkStateChangeTriggers   bool `json:"allowNetworkStateChangeTriggers"`
}
