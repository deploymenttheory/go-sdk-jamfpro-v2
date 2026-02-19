package cloud_distribution_point

// RequestCloudDistributionPointV1 is the body for CreateV1 and UpdateV1 (PATCH).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-cloud-distribution-point
type RequestCloudDistributionPointV1 struct {
	CdnType                 string `json:"cdnType"`
	Master                  bool   `json:"master"`
	Username                string `json:"username,omitempty"`
	Password                string `json:"password,omitempty"`
	Directory               string `json:"directory,omitempty"`
	UploadUrl               string `json:"uploadUrl,omitempty"`
	DownloadUrl             string `json:"downloadUrl,omitempty"`
	SecondaryAuthRequired   *bool  `json:"secondaryAuthRequired,omitempty"`
	SecondaryAuthStatusCode *int   `json:"secondaryAuthStatusCode,omitempty"`
	SecondaryAuthTimeToLive *int   `json:"secondaryAuthTimeToLive,omitempty"`
	RequireSignedUrls       *bool  `json:"requireSignedUrls,omitempty"`
	KeyPairId               string `json:"keyPairId,omitempty"`
	ExpirationSeconds       *int   `json:"expirationSeconds,omitempty"`
	PrivateKey              string `json:"privateKey,omitempty"`
}

// ResourceCloudDistributionPointV1 is the response for GetV1, CreateV1, UpdateV1.
type ResourceCloudDistributionPointV1 struct {
	HasConnectionSucceeded  bool   `json:"hasConnectionSucceeded"`
	Message                 string `json:"message"`
	InventoryId             string `json:"inventoryId"`
	CdnType                 string `json:"cdnType"`
	Master                  bool   `json:"master"`
	Username                string `json:"username"`
	Directory               string `json:"directory"`
	CdnUrl                  string `json:"cdnUrl"`
	UploadUrl               string `json:"uploadUrl"`
	DownloadUrl             string `json:"downloadUrl"`
	SecondaryAuthRequired   bool   `json:"secondaryAuthRequired"`
	SecondaryAuthStatusCode int    `json:"secondaryAuthStatusCode"`
	SecondaryAuthTimeToLive int    `json:"secondaryAuthTimeToLive"`
	RequireSignedUrls       bool   `json:"requireSignedUrls"`
	KeyPairId               string `json:"keyPairId"`
	ExpirationSeconds       int    `json:"expirationSeconds"`
	PrivateKey              any    `json:"privateKey"`
}

// UploadCapabilityV1 is the response for GetUploadCapabilityV1.
type UploadCapabilityV1 struct {
	PrincipalDistributionTechnology bool `json:"principalDistributionTechnology"`
	DirectUploadCapable             bool `json:"directUploadCapable"`
}

// TestConnectionV1 is the response for GetTestConnectionV1.
type TestConnectionV1 struct {
	HasConnectionSucceeded bool   `json:"hasConnectionSucceeded"`
	Message                string `json:"message"`
}
