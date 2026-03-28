package cloud_distribution_point

// FileInfoStatus* constants represent the status values for cloud distribution point inventory file info.
const (
	FileInfoStatusReady   = "READY"
	FileInfoStatusPending = "PENDING"
	FileInfoStatusError   = "ERROR"
)

// FileInfoType* constants represent the file types for cloud distribution point inventory file info.
const (
	FileInfoTypeNone            = "NONE"
	FileInfoTypePackage         = "PACKAGE"
	FileInfoTypeEbook           = "EBOOK"
	FileInfoTypeMobileDeviceApp = "MOBILE_DEVICE_APP"
	FileInfoTypeScript          = "SCRIPT"
)

// CdnType enum values for RequestCloudDistributionPointV1.CdnType and ResourceCloudDistributionPointV1.CdnType.
const (
	CdnTypeNone               = "NONE"
	CdnTypeJamfCloud          = "JAMF_CLOUD"
	CdnTypeRackspaceCloudFiles = "RACKSPACE_CLOUD_FILES"
	CdnTypeAmazonS3           = "AMAZON_S3"
	CdnTypeAkamai             = "AKAMAI"
)

// validCdnTypes is the set of accepted cdnType values for request validation.
var validCdnTypes = map[string]struct{}{
	CdnTypeNone:               {},
	CdnTypeJamfCloud:          {},
	CdnTypeRackspaceCloudFiles: {},
	CdnTypeAmazonS3:           {},
	CdnTypeAkamai:             {},
}
