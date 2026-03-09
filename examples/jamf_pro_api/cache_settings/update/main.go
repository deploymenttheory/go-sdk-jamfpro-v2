package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cache_settings"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	settings := cache_settings.ResourceCacheSettings{
		CacheType:                  "ehcache",
		TimeToLiveSeconds:          3600,
		TimeToIdleSeconds:          1800,
		DirectoryTimeToLiveSeconds: 600,
		EhcacheMaxBytesLocalHeap:   "256M",
		CacheUniqueID:              "jamf-pro-cache",
		Elasticache:                false,
		MemcachedEndpoints:         []cache_settings.CacheSettingsSubsetMemcachedEndpoint{},
	}

	updated, _, err := jamfClient.JamfProAPI.CacheSettings.UpdateV1(context.Background(), &settings)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated cache settings:\n" + string(out))
}
