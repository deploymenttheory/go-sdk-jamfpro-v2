package cache_settings

// ResourceCacheSettings represents the cache settings resource.
type ResourceCacheSettings struct {
	ID                         string                               `json:"id,omitempty"`
	Name                       string                               `json:"name,omitempty"`
	CacheType                  string                               `json:"cacheType"`
	TimeToLiveSeconds          int                                  `json:"timeToLiveSeconds"`
	TimeToIdleSeconds          int                                  `json:"timeToIdleSeconds"`
	DirectoryTimeToLiveSeconds int                                  `json:"directoryTimeToLiveSeconds,omitempty"`
	EhcacheMaxBytesLocalHeap   string                               `json:"ehcacheMaxBytesLocalHeap,omitempty"`
	CacheUniqueID              string                               `json:"cacheUniqueId"`
	Elasticache                bool                                 `json:"elasticache,omitempty"`
	MemcachedEndpoints         []CacheSettingsSubsetMemcachedEndpoint `json:"memcachedEndpoints"`
}

// CacheSettingsSubsetMemcachedEndpoint represents a memcached endpoint.
type CacheSettingsSubsetMemcachedEndpoint struct {
	ID                      string `json:"id,omitempty"`
	Name                    string `json:"name,omitempty"`
	HostName                string `json:"hostName,omitempty"`
	Port                    int    `json:"port,omitempty"`
	Enabled                 bool   `json:"enabled,omitempty"`
	JSSCacheConfigurationID int    `json:"jssCacheConfigurationId,omitempty"`
}
