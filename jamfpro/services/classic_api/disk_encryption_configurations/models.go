package disk_encryption_configurations

import "encoding/xml"

// ResourceDiskEncryptionConfiguration represents a Jamf Pro Classic API disk encryption configuration resource.
type ResourceDiskEncryptionConfiguration struct {
	XMLName                  xml.Name                          `xml:"disk_encryption_configuration"`
	ID                       int                               `xml:"id"`
	Name                     string                            `xml:"name"`
	KeyType                  string                            `xml:"key_type,omitempty"`
	FileVaultEnabledUsers    string                            `xml:"file_vault_enabled_users,omitempty"`
	InstitutionalRecoveryKey *InstitutionalRecoveryKey         `xml:"institutional_recovery_key,omitempty"`
}

// InstitutionalRecoveryKey holds the institutional recovery key details for a disk encryption configuration.
type InstitutionalRecoveryKey struct {
	Key             string `xml:"key,omitempty"`
	CertificateType string `xml:"certificate_type,omitempty"`
	Password        string `xml:"password,omitempty"`
	Data            string `xml:"data,omitempty"`
}

// ListItemDiskEncryptionConfiguration is the slim representation returned in list responses.
type ListItemDiskEncryptionConfiguration struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ListResponse is the response for ListDiskEncryptionConfigurations (GET /JSSResource/diskencryptionconfigurations).
type ListResponse struct {
	XMLName xml.Name                               `xml:"disk_encryption_configurations"`
	Size    int                                    `xml:"size"`
	Results []ListItemDiskEncryptionConfiguration  `xml:"disk_encryption_configuration"`
}

// CreateUpdateResponse is the response body for create and update operations,
// which return only the assigned resource ID.
type CreateUpdateResponse struct {
	XMLName xml.Name `xml:"disk_encryption_configuration"`
	ID      int      `xml:"id"`
}

// RequestDiskEncryptionConfiguration is the body for creating or updating a disk encryption configuration.
// The ID field is not included; the target is specified via the URL path.
type RequestDiskEncryptionConfiguration struct {
	XMLName                  xml.Name                  `xml:"disk_encryption_configuration"`
	Name                     string                    `xml:"name"`
	KeyType                  string                    `xml:"key_type,omitempty"`
	FileVaultEnabledUsers    string                    `xml:"file_vault_enabled_users,omitempty"`
	InstitutionalRecoveryKey *InstitutionalRecoveryKey `xml:"institutional_recovery_key,omitempty"`
}
