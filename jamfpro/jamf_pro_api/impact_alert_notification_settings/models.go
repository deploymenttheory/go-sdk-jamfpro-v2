package impact_alert_notification_settings

// ResourceImpactAlertNotificationSettings represents the impact alert notification settings.
type ResourceImpactAlertNotificationSettings struct {
	ScopeableObjectsAlertEnabled             bool `json:"scopeableObjectsAlertEnabled"`
	ScopeableObjectsConfirmationCodeEnabled  bool `json:"scopeableObjectsConfirmationCodeEnabled"`
	DeployableObjectsAlertEnabled            bool `json:"deployableObjectsAlertEnabled"`
	DeployableObjectsConfirmationCodeEnabled bool `json:"deployableObjectsConfirmationCodeEnabled"`
}
