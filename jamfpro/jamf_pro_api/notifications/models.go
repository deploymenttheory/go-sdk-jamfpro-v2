package notifications

// NotificationType represents the type of a Jamf Pro notification.
// Valid values for DELETE /api/v1/notifications/{type}/{id}.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-notifications-type-id
const (
	NotificationTypeAPNSCertRevoked                            = "APNS_CERT_REVOKED"
	NotificationTypeAPNSConnectionFailure                      = "APNS_CONNECTION_FAILURE"
	NotificationTypeAppleSchoolManagerTCNotSigned              = "APPLE_SCHOOL_MANAGER_T_C_NOT_SIGNED"
	NotificationTypeBuiltInCAExpired                           = "BUILT_IN_CA_EXPIRED"
	NotificationTypeBuiltInCAExpiring                          = "BUILT_IN_CA_EXPIRING"
	NotificationTypeBuiltInCARenewalFailed                     = "BUILT_IN_CA_RENEWAL_FAILED"
	NotificationTypeBuiltInCARenewalSuccess                    = "BUILT_IN_CA_RENEWAL_SUCCESS"
	NotificationTypeCloudLDAPCertExpired                       = "CLOUD_LDAP_CERT_EXPIRED"
	NotificationTypeCloudLDAPCertWillExpire                    = "CLOUD_LDAP_CERT_WILL_EXPIRE"
	NotificationTypeDEPInstanceExpired                         = "DEP_INSTANCE_EXPIRED"
	NotificationTypeDEPInstanceWillExpire                      = "DEP_INSTANCE_WILL_EXPIRE"
	NotificationTypeDeviceEnrollmentProgramTCNotSigned         = "DEVICE_ENROLLMENT_PROGRAM_T_C_NOT_SIGNED"
	NotificationTypeExceededLicenseCount                       = "EXCEEDED_LICENSE_COUNT"
	NotificationTypeFrequentInventoryCollectionPolicy          = "FREQUENT_INVENTORY_COLLECTION_POLICY"
	NotificationTypeGSXCertExpired                             = "GSX_CERT_EXPIRED"
	NotificationTypeGSXCertWillExpire                          = "GSX_CERT_WILL_EXPIRE"
	NotificationTypeHCLBindError                               = "HCL_BIND_ERROR"
	NotificationTypeHCLError                                   = "HCL_ERROR"
	NotificationTypeInsecureLDAP                               = "INSECURE_LDAP"
	NotificationTypeInvalidReferencesExtAttr                   = "INVALID_REFERENCES_EXT_ATTR"
	NotificationTypeInvalidReferencesPolicies                  = "INVALID_REFERENCES_POLICIES"
	NotificationTypeInvalidReferencesScripts                   = "INVALID_REFERENCES_SCRIPTS"
	NotificationTypeJamfConnectUpdate                          = "JAMF_CONNECT_UPDATE"
	NotificationTypeJamfProtectUpdate                          = "JAMF_PROTECT_UPDATE"
	NotificationTypeJIMError                                   = "JIM_ERROR"
	NotificationTypeLDAPConnectionCheckThroughJIMFailed        = "LDAP_CONNECTION_CHECK_THROUGH_JIM_FAILED"
	NotificationTypeLDAPConnectionCheckThroughJIMSuccessful    = "LDAP_CONNECTION_CHECK_THROUGH_JIM_SUCCESSFUL"
	NotificationTypeMDMExternalSigningCertificateExpired       = "MDM_EXTERNAL_SIGNING_CERTIFICATE_EXPIRED"
	NotificationTypeMDMExternalSigningCertificateExpiring      = "MDM_EXTERNAL_SIGNING_CERTIFICATE_EXPIRING"
	NotificationTypeMDMExternalSigningCertificateExpiringToday = "MDM_EXTERNAL_SIGNING_CERTIFICATE_EXPIRING_TODAY"
	NotificationTypeMIIHeartbeatFailedNotification             = "MII_HEARTBEAT_FAILED_NOTIFICATION"
	NotificationTypeMIIInventoryUploadFailedNotification       = "MII_INVENTORY_UPLOAD_FAILED_NOTIFICATION"
	NotificationTypeMIIUnauthorizedResponseNotification        = "MII_UNATHORIZED_RESPONSE_NOTIFICATION"
	NotificationTypePatchExtentionAttribute                    = "PATCH_EXTENTION_ATTRIBUTE"
	NotificationTypePatchUpdate                                = "PATCH_UPDATE"
	NotificationTypePolicyManagementAccountPayloadSecMultiple  = "POLICY_MANAGEMENT_ACCOUNT_PAYLOAD_SECURITY_MULTIPLE"
	NotificationTypePolicyManagementAccountPayloadSecSingle    = "POLICY_MANAGEMENT_ACCOUNT_PAYLOAD_SECURITY_SINGLE"
	NotificationTypePushCertExpired                            = "PUSH_CERT_EXPIRED"
	NotificationTypePushCertWillExpire                         = "PUSH_CERT_WILL_EXPIRE"
	NotificationTypePushProxyCertExpired                       = "PUSH_PROXY_CERT_EXPIRED"
	NotificationTypeSSOCertExpired                             = "SSO_CERT_EXPIRED"
	NotificationTypeSSOIdPCertExpired                          = "SSO_IDP_CERT_EXPIRED"
	NotificationTypeSSOCertWillExpire                          = "SSO_CERT_WILL_EXPIRE"
	NotificationTypeSSOIdPCertWillExpire                       = "SSO_IDP_CERT_WILL_EXPIRE"
	NotificationTypeTomcatSSLCertExpired                       = "TOMCAT_SSL_CERT_EXPIRED"
	NotificationTypeTomcatSSLCertWillExpire                    = "TOMCAT_SSL_CERT_WILL_EXPIRE"
	NotificationTypeUserInitiatedEnrollmentMgmtAccountSecIssue = "USER_INITIATED_ENROLLMENT_MANAGEMENT_ACCOUNT_SECURITY_ISSUE"
	NotificationTypeUserMAIDDuplicateError                     = "USER_MAID_DUPLICATE_ERROR"
	NotificationTypeUserMAIDMismatchError                      = "USER_MAID_MISMATCH_ERROR"
	NotificationTypeUserMAIDRosterDuplicateError               = "USER_MAID_ROSTER_DUPLICATE_ERROR"
	NotificationTypeVPPAccountExpired                          = "VPP_ACCOUNT_EXPIRED"
	NotificationTypeVPPAccountWillExpire                       = "VPP_ACCOUNT_WILL_EXPIRE"
	NotificationTypeVPPTokenRevoked                            = "VPP_TOKEN_REVOKED"
	NotificationTypeDeviceComplianceConnectionError            = "DEVICE_COMPLIANCE_CONNECTION_ERROR"
	NotificationTypeConditionalAccessConnectionError           = "CONDITIONAL_ACCESS_CONNECTION_ERROR"
	NotificationTypeAzureADMigrationReportGenerated            = "AZURE_AD_MIGRATION_REPORT_GENERATED"
	NotificationTypeBeyondCorpConnectionError                  = "BEYOND_CORP_CONNECTION_ERROR"
	NotificationTypeAppInstallersNewAppVersionAvailable        = "APP_INSTALLERS_NEW_APP_VERSION_AVAILABLE"
	NotificationTypeAppInstallersNewAppVersionDeploymentStarted = "APP_INSTALLERS_NEW_APP_VERSION_DEPLOYMENT_STARTED"
	NotificationTypeAppInstallersAppVersionRemoved             = "APP_INSTALLERS_APP_VERSION_REMOVED"
	NotificationTypeAppInstallersAppTitleRemoved               = "APP_INSTALLERS_APP_TITLE_REMOVED"
	NotificationTypeAppInstallersDeploymentInstallationFailed  = "APP_INSTALLERS_DEPLOYMENT_INSTALLATION_FAILED"
	NotificationTypeSAMLResponseAssertionSigningRequired       = "SAML_RESPONSE_ASSERTION_SIGNING_REQUIRED"
	// NotificationTypeDirectoryCacheAwaitingSync was added in Jamf Pro 11.25.
	NotificationTypeDirectoryCacheAwaitingSync                 = "DIRECTORY_CACHE_AWAITING_SYNC"
)

// ResourceNotification represents a single notification.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
type ResourceNotification struct {
	Type   string         `json:"type"`
	ID     string         `json:"id"`
	Params map[string]any `json:"params"`
}
