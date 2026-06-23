package m2m

// ResourceM2mTenantId is the response for GetTenantIdV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-m2m-tenant-id
type ResourceM2mTenantId struct {
	TenantId string `json:"tenantId"`
}
