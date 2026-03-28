package volume_purchasing_subscriptions

// SubscriptionTrigger* constants represent the trigger types for volume purchasing subscriptions.
const (
	SubscriptionTriggerNoMoreLicenses      = "NO_MORE_LICENSES"
	SubscriptionTriggerRemovedFromAppStore = "REMOVED_FROM_APP_STORE"
)

var validSubscriptionTriggers = map[string]struct{}{
	SubscriptionTriggerNoMoreLicenses:      {},
	SubscriptionTriggerRemovedFromAppStore: {},
}
