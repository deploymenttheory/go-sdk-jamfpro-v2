package self_service_plus_settings

// ResourceSelfServicePlusSettings is the self-service plus settings resource (get/update).
type ResourceSelfServicePlusSettings struct {
	Enabled bool `json:"enabled"`
}

// ResourceFeatureToggleEnabled is the response for the Self Service Plus feature toggle enabled endpoint.
type ResourceFeatureToggleEnabled struct {
	Enabled bool `json:"enabled"`
}
