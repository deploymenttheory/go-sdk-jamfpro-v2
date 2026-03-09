package volume_purchasing_subscriptions

// ResourceVolumePurchasingSubscription represents a volume purchasing subscription resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-volume-purchasing-subscriptions-id
type ResourceVolumePurchasingSubscription struct {
	ID                 string                                    `json:"id,omitempty"`
	Name               string                                    `json:"name"`
	Enabled            bool                                     `json:"enabled,omitempty"`
	Triggers           []string                                 `json:"triggers,omitempty"`
	LocationIds        []string                                  `json:"locationIds,omitempty"`
	InternalRecipients []VolumePurchasingSubscriptionInternalRecipient `json:"internalRecipients,omitempty"`
	ExternalRecipients []VolumePurchasingSubscriptionExternalRecipient `json:"externalRecipients,omitempty"`
	SiteId             string                                    `json:"siteId,omitempty"`
}

// VolumePurchasingSubscriptionInternalRecipient represents an internal recipient.
type VolumePurchasingSubscriptionInternalRecipient struct {
	AccountId string `json:"accountId,omitempty"`
	Frequency string `json:"frequency,omitempty"`
}

// VolumePurchasingSubscriptionExternalRecipient represents an external recipient.
type VolumePurchasingSubscriptionExternalRecipient struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// ListResponse is the response for ListVolumePurchasingSubscriptionsV1.
type ListResponse struct {
	TotalCount int                                   `json:"totalCount"`
	Results    []ResourceVolumePurchasingSubscription `json:"results"`
}

// RequestVolumePurchasingSubscription is the body for creating and updating volume purchasing subscriptions.
// Same shape as resource for create/update.
type RequestVolumePurchasingSubscription struct {
	Name               string                                                    `json:"name"`
	Enabled            bool                                                      `json:"enabled,omitempty"`
	Triggers           []string                                                  `json:"triggers,omitempty"`
	LocationIds        []string                                                  `json:"locationIds,omitempty"`
	InternalRecipients []VolumePurchasingSubscriptionInternalRecipient `json:"internalRecipients,omitempty"`
	ExternalRecipients []VolumePurchasingSubscriptionExternalRecipient `json:"externalRecipients,omitempty"`
	SiteId             string                                                    `json:"siteId,omitempty"`
}

// CreateResponse is the response for CreateVolumePurchasingSubscriptionV1.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}
