Return paginated Computer Inventory records
get
https://yourServer.jamfcloud.com/api/v3/computers-inventory

Return paginated Computer Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of computer details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


string


GENERAL

ADD string
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
sort
array of strings
Defaults to general.name:asc
Sorting criteria in the format: property:asc/desc. Default sort is general.name:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: general.name, udid, id, general.assetTag, general.jamfBinaryVersion, general.lastContactTime, general.lastEnrolledDate, general.lastCloudBackupDate, general.reportDate, general.mdmCertificateExpiration, general.platform, general.lastLoggedInUsernameSelfService, general.lastLoggedInUsernameSelfServiceTimestamp, general.mdmCertificateExpiration, general.platform, general.lastLoggedInUsernameBinary, general.lastLoggedInUsernameBinaryTimestamp hardware.make, hardware.model, operatingSystem.build, operatingSystem.supplementalBuildVersion, operatingSystem.rapidSecurityResponse, operatingSystem.name, operatingSystem.version, userAndLocation.realname, purchasing.lifeExpectancy, purchasing.warrantyDate

Example: sort=udid:desc,general.name:asc.


string

general.name:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter computer inventory collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: general.name, udid, id, general.assetTag, general.barcode1, general.barcode2, general.enrolledViaAutomatedDeviceEnrollment, general.lastIpAddress, general.itunesStoreAccountActive, general.jamfBinaryVersion, general.lastContactTime, general.lastEnrolledDate, general.lastCloudBackupDate, general.reportDate, general.lastReportedIp, general.lastReportedIpV4, general.lastReportedIpV6, general.managementId, general.remoteManagement.managed, general.mdmCapable.capable, general.mdmCertificateExpiration, general.platform, general.supervised, general.userApprovedMdm, general.declarativeDeviceManagementEnabled, general.lastLoggedInUsernameSelfService, general.lastLoggedInUsernameSelfServiceTimestamp, general.mdmCapable.capable, general.mdmCertificateExpiration, general.platform, general.supervised, general.userApprovedMdm, general.declarativeDeviceManagementEnabled, general.lastLoggedInUsernameBinary, general.lastLoggedInUsernameBinaryTimestamp, hardware.bleCapable, hardware.macAddress, hardware.make, hardware.model, hardware.modelIdentifier, hardware.serialNumber, hardware.supportsIosAppInstalls,hardware.appleSilicon, operatingSystem.activeDirectoryStatus, operatingSystem.fileVault2Status, operatingSystem.build, operatingSystem.supplementalBuildVersion, operatingSystem.rapidSecurityResponse, operatingSystem.name, operatingSystem.version, security.activationLockEnabled, security.recoveryLockEnabled,security.firewallEnabled,userAndLocation.buildingId, userAndLocation.departmentId, userAndLocation.email, userAndLocation.realname, userAndLocation.phone, userAndLocation.position,userAndLocation.room, userAndLocation.username, diskEncryption.fileVault2Enabled, purchasing.appleCareId, purchasing.lifeExpectancy, purchasing.purchased, purchasing.leased, purchasing.vendor, purchasing.warrantyDate,

This param can be combined with paging and sorting. Example: filter=general.name=="Orchard"

Response

200
Successful response

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
id
string
length ≥ 1
udid
string
general
object

general object
name
string
length ≥ 1
lastIpAddress
string
lastReportedIp
string
deprecated
Last reported IPv4 address (Deprecated. Use lastReportedIpV4 instead.)

lastReportedIpV4
string
Last reported IPv4 address

lastReportedIpV6
string
jamfBinaryVersion
string
platform
string
barcode1
string
barcode2
string
assetTag
string
remoteManagement
object

remoteManagement object
managed
boolean
managementUsername
string
deprecated
This field always returns null, please use /local-admin-password/ endpoint instead.

supervised
boolean
mdmCapable
object

mdmCapable object
capable
boolean
capableUsers
array of strings
deprecated
Deprecated. Use userManagementInfo instead.

userManagementInfo
array of objects
object
capableUser
string
required
managementId
string
required
reportDate
date-time
lastContactTime
date-time
lastCloudBackupDate
date-time
lastEnrolledDate
date-time
mdmProfileExpiration
date-time
initialEntryDate
date
distributionPoint
string
enrollmentMethod
object

enrollmentMethod object
id
string
objectName
string
objectType
string
site
object

site object
id
string
name
string
itunesStoreAccountActive
boolean
enrolledViaAutomatedDeviceEnrollment
boolean
userApprovedMdm
boolean
declarativeDeviceManagementEnabled
boolean
extensionAttributes
array of objects
object
definitionId
string
An identifier of extension attribute definition.

name
string
A human-readable name by which attribute can be referred to.

description
string | null
An additional explanation of exact attribute meaning, possible values, etc.

enabled
boolean
multiValue
boolean
values
array of strings | null
A value of extension attribute, in some rare cases there may be multiple values present, hence the array.

dataType
string | null
enum
A data type of extension attribute.

STRING INTEGER DATE_TIME

options
array of strings | null
A closed list of possible values (applies to popup input type).

inputType
string | null
enum
The input method. text is most common and means simply free text, popup i a closed list of values from which one or many can be selected and script value is calculated and can never be set directly.

TEXT POPUP SCRIPT LDAP

managementId
string
lastLoggedInUsernameSelfService
string | null
lastLoggedInUsernameSelfServiceTimestamp
date-time | null
lastLoggedInUsernameBinary
string | null
lastLoggedInUsernameBinaryTimestamp
date-time | null
diskEncryption
object

diskEncryption object
bootPartitionEncryptionDetails
object

bootPartitionEncryptionDetails object
partitionName
string
partitionFileVault2State
string
enum
UNKNOWN UNENCRYPTED INELIGIBLE DECRYPTED DECRYPTING ENCRYPTED ENCRYPTING RESTART_NEEDED OPTIMIZING DECRYPTING_PAUSED ENCRYPTING_PAUSED

partitionFileVault2Percent
integer
individualRecoveryKeyValidityStatus
string
enum
VALID INVALID UNKNOWN NOT_APPLICABLE

institutionalRecoveryKeyPresent
boolean
diskEncryptionConfigurationName
string
fileVault2Enabled
boolean
fileVault2EnabledUserNames
array of strings
fileVault2EligibilityMessage
string
purchasing
object

purchasing object
leased
boolean
purchased
boolean
poNumber
string
poDate
date
vendor
string
warrantyDate
date
appleCareId
string
leaseDate
date
purchasePrice
string
lifeExpectancy
integer
purchasingAccount
string
purchasingContact
string
extensionAttributes
array of objects
object
definitionId
string
An identifier of extension attribute definition.

name
string
A human-readable name by which attribute can be referred to.

description
string | null
An additional explanation of exact attribute meaning, possible values, etc.

enabled
boolean
multiValue
boolean
values
array of strings | null
A value of extension attribute, in some rare cases there may be multiple values present, hence the array.

dataType
string | null
enum
A data type of extension attribute.

STRING INTEGER DATE_TIME

options
array of strings | null
A closed list of possible values (applies to popup input type).

inputType
string | null
enum
The input method. text is most common and means simply free text, popup i a closed list of values from which one or many can be selected and script value is calculated and can never be set directly.

TEXT POPUP SCRIPT LDAP

applications
array of objects
object
name
string
path
string
version
string
cfBundleShortVersionString
string
cfBundleVersion
string
macAppStore
boolean
sizeMegabytes
integer
bundleId
string
updateAvailable
boolean
externalVersionId
string
The app's external version ID. It can be used in the iTunes Search API to decide if the app needs to be updated

storage
object

storage object
userAndLocation
object

userAndLocation object
configurationProfiles
array of objects
object
id
string
username
string
lastInstalled
date-time
removable
boolean
displayName
string
profileIdentifier
string
printers
array of objects
object
name
string
type
string
uri
string
location
string
services
array of objects
object
name
string
hardware
object

hardware object
make
string
model
string
modelIdentifier
string
serialNumber
string
processorSpeedMhz
int64
Processor Speed in MHz.

processorCount
integer
coreCount
integer
processorType
string
processorArchitecture
string
busSpeedMhz
int64
cacheSizeKilobytes
int64
Cache Size in KB.

networkAdapterType
string
macAddress
string
altNetworkAdapterType
string
altMacAddress
string
totalRamMegabytes
int64
Total RAM Size in MB.

openRamSlots
integer
Available RAM slots.

batteryCapacityPercent
integer
0 to 100
Remaining percentage of battery power.

batteryHealth
string
enum
Defaults to UNKNOWN
NON_GENUINE: The battery isnâ€™t a genuine Apple battery.
NORMAL: The battery is operating normally.
SERVICE_RECOMMENDED: The system recommends battery service.
UNKNOWN: The system couldnâ€™t determine battery health information.
UNSUPPORTED: The device doesnâ€™t support battery health reporting.
NON_GENUINE NORMAL SERVICE_RECOMMENDED UNKNOWN UNSUPPORTED

smcVersion
string
nicSpeed
string
opticalDrive
string
bootRom
string
bleCapable
boolean
supportsIosAppInstalls
boolean
appleSilicon
boolean
provisioningUdid
string
extensionAttributes
array of objects
object
definitionId
string
An identifier of extension attribute definition.

name
string
A human-readable name by which attribute can be referred to.

description
string | null
An additional explanation of exact attribute meaning, possible values, etc.

enabled
boolean
multiValue
boolean
values
array of strings | null
A value of extension attribute, in some rare cases there may be multiple values present, hence the array.

dataType
string | null
enum
A data type of extension attribute.

STRING INTEGER DATE_TIME

options
array of strings | null
A closed list of possible values (applies to popup input type).

inputType
string | null
enum
The input method. text is most common and means simply free text, popup i a closed list of values from which one or many can be selected and script value is calculated and can never be set directly.

TEXT POPUP SCRIPT LDAP

localUserAccounts
array of objects
object
uid
string
userGuid
string
username
string
fullName
string
admin
boolean
homeDirectory
string
homeDirectorySizeMb
int64
Home directory size in MB.

fileVault2Enabled
boolean
userAccountType
string
enum
LOCAL MOBILE UNKNOWN

passwordMinLength
integer
passwordMaxAge
integer
passwordMinComplexCharacters
integer
passwordHistoryDepth
integer
passwordRequireAlphanumeric
boolean
computerAzureActiveDirectoryId
string
userAzureActiveDirectoryId
string
azureActiveDirectoryId
string
enum
ACTIVATED DEACTIVATED UNRESPONSIVE UNKNOWN

certificates
array of objects
object
commonName
string
identity
boolean
expirationDate
date-time
username
string
lifecycleStatus
string
enum
ACTIVE INACTIVE

certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDate
string
attachments
array of objects
object
id
string
name
string
fileType
string
sizeBytes
int64
File size in bytes

packageReceipts
object
All package receipts are listed by their package name


packageReceipts object
installedByJamfPro
array of strings
installedByInstallerSwu
array of strings
cached
array of strings
security
object

security object
sipStatus
string
enum
NOT_COLLECTED NOT_AVAILABLE DISABLED ENABLED

gatekeeperStatus
string
enum
NOT_COLLECTED DISABLED APP_STORE_AND_IDENTIFIED_DEVELOPERS APP_STORE

xprotectVersion
string
autoLoginDisabled
boolean
remoteDesktopEnabled
boolean
Collected for macOS 10.14.4 or later

activationLockEnabled
boolean
Collected for macOS 10.15.0 or later

recoveryLockEnabled
boolean
firewallEnabled
boolean
secureBootLevel
string
enum
Collected for macOS 10.15.0 or later

NO_SECURITY MEDIUM_SECURITY FULL_SECURITY NOT_SUPPORTED UNKNOWN

externalBootLevel
string
enum
Collected for macOS 10.15.0 or later

ALLOW_BOOTING_FROM_EXTERNAL_MEDIA DISALLOW_BOOTING_FROM_EXTERNAL_MEDIA NOT_SUPPORTED UNKNOWN

bootstrapTokenAllowed
boolean
Collected for macOS 11 or later

bootstrapTokenEscrowedStatus
string
enum
Collected for macOS 11 or later

ESCROWED NOT_ESCROWED NOT_SUPPORTED

lastAttestationAttempt
string
lastSuccessfulAttestation
string
attestationStatus
string
enum
PENDING SUCCESS CERTIFICATE_INVALID DEVICE_PROPERTIES_MISMATCH MDA_UNSUPPORTED_DUE_TO_HARDWARE MDA_UNSUPPORTED_DUE_TO_SOFTWARE

operatingSystem
object

operatingSystem object
name
string
version
string
build
string
supplementalBuildVersion
string
Collected for macOS 13.0 or later

rapidSecurityResponse
string
Collected for macOS 13.0 or later

activeDirectoryStatus
string
fileVault2Status
string
enum
NOT_APPLICABLE NOT_ENCRYPTED BOOT_ENCRYPTED SOME_ENCRYPTED ALL_ENCRYPTED

softwareUpdateDeviceId
string
extensionAttributes
array of objects
object
definitionId
string
An identifier of extension attribute definition.

name
string
A human-readable name by which attribute can be referred to.

description
string | null
An additional explanation of exact attribute meaning, possible values, etc.

enabled
boolean
multiValue
boolean
values
array of strings | null
A value of extension attribute, in some rare cases there may be multiple values present, hence the array.

dataType
string | null
enum
A data type of extension attribute.

STRING INTEGER DATE_TIME

options
array of strings | null
A closed list of possible values (applies to popup input type).

inputType
string | null
enum
The input method. text is most common and means simply free text, popup i a closed list of values from which one or many can be selected and script value is calculated and can never be set directly.

TEXT POPUP SCRIPT LDAP

licensedSoftware
array of objects
object
id
string
name
string
ibeacons
array of objects
object
name
string
softwareUpdates
array of objects
object
name
string
version
string
packageName
string
extensionAttributes
array of objects
object
definitionId
string
An identifier of extension attribute definition.

name
string
A human-readable name by which attribute can be referred to.

description
string | null
An additional explanation of exact attribute meaning, possible values, etc.

enabled
boolean
multiValue
boolean
values
array of strings | null
A value of extension attribute, in some rare cases there may be multiple values present, hence the array.

dataType
string | null
enum
A data type of extension attribute.

STRING INTEGER DATE_TIME

options
array of strings | null
A closed list of possible values (applies to popup input type).

inputType
string | null
enum
The input method. text is most common and means simply free text, popup i a closed list of values from which one or many can be selected and script value is calculated and can never be set directly.

TEXT POPUP SCRIPT LDAP

contentCaching
object

contentCaching object
computerContentCachingInformationId
string
parents
array of objects
object
contentCachingParentId
string
address
string
alerts
object

alerts object
contentCachingParentAlertId
string
addresses
array of strings
className
string
postDate
date-time
details
object

details object
contentCachingParentDetailsId
string
acPower
boolean
cacheSizeBytes
int64
capabilities
object

capabilities object
contentCachingParentCapabilitiesId
string
imports
boolean
namespaces
boolean
personalContent
boolean
queryParameters
boolean
sharedContent
boolean
prioritization
boolean
portable
boolean
localNetwork
array of objects
object
contentCachingParentLocalNetworkId
string
speed
int64
wired
boolean
guid
string
healthy
boolean
port
int64
version
string
alerts
array of objects
object
cacheBytesLimit
int64
className
string
pathPreventingAccess
string
postDate
date-time
reservedVolumeBytes
int64
resource
string
activated
boolean
active
boolean
actualCacheBytesUsed
int64
cacheDetails
array of objects
object
computerContentCachingCacheDetailsId
string
categoryName
string
diskSpaceBytesUsed
int64
cacheBytesFree
int64
cacheBytesLimit
int64
cacheStatus
string
cacheBytesUsed
int64
dataMigrationCompleted
boolean
dataMigrationProgressPercentage
integer
dataMigrationError
object

dataMigrationError object
code
int64
domain
string
userInfo
array of objects
object
key
string
value
string
maxCachePressureLast1HourPercentage
integer
personalCacheBytesFree
int64
personalCacheBytesLimit
int64
personalCacheBytesUsed
int64
port
int64
publicAddress
string
registrationError
string
registrationResponseCode
int64
registrationStarted
date-time
registrationStatus
string
enum
CONTENT_CACHING_FAILED CONTENT_CACHING_PENDING CONTENT_CACHING_SUCCEEDED

restrictedMedia
boolean
serverGuid
string
startupStatus
string
tetheratorStatus
string
enum
CONTENT_CACHING_UNKNOWN CONTENT_CACHING_DISABLED CONTENT_CACHING_ENABLED

totalBytesAreSince
date-time
totalBytesDropped
int64
totalBytesImported
int64
totalBytesReturnedToChildren
int64
totalBytesReturnedToClients
int64
totalBytesReturnedToPeers
int64
totalBytesStoredFromOrigin
int64
totalBytesStoredFromParents
int64
totalBytesStoredFromPeers
int64
groupMemberships
array of objects
object
groupId
string
groupName
string
groupDescription
string
smartGroup
boolean
Indicates that group is smart group

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/computers-inventory?section=GENERAL&page=0&page-size=100&sort=general.name%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "udid": "123",
      "general": {
        "name": "Boalime",
        "lastIpAddress": "247.185.82.186",
        "lastReportedIpV4": "247.185.82.186",
        "lastReportedIpV6": "2001:0db8:85a3:0000:0000:8a2e:0370:7335",
        "jamfBinaryVersion": "9.27",
        "platform": "Mac",
        "barcode1": "5 12345 678900",
        "barcode2": "5 12345 678900",
        "assetTag": "304822",
        "remoteManagement": {
          "managed": true
        },
        "supervised": true,
        "mdmCapable": {
          "capable": true,
          "userManagementInfo": [
            [
              {
                "capableUser": "admin",
                "managementId": "123e4567-e89b-42d3-a456-426614174000"
              },
              {
                "capableUser": "rootadmin",
                "managementId": "123e4567-e89b-42d3-a456-426614174001"
              }
            ]
          ]
        },
        "reportDate": "2018-10-31T18:04:13Z",
        "lastContactTime": "2018-10-31T18:04:13Z",
        "lastCloudBackupDate": "2018-10-31T18:04:13Z",
        "lastEnrolledDate": "2018-10-31T18:04:13Z",
        "mdmProfileExpiration": "2018-10-31T18:04:13Z",
        "initialEntryDate": "2018-10-31",
        "distributionPoint": "distribution point name",
        "enrollmentMethod": {
          "id": "1",
          "objectName": "user@domain.com",
          "objectType": "User-initiated - no invitation"
        },
        "site": {
          "id": "1",
          "name": "Eau Claire"
        },
        "itunesStoreAccountActive": true,
        "enrolledViaAutomatedDeviceEnrollment": true,
        "userApprovedMdm": true,
        "declarativeDeviceManagementEnabled": true,
        "extensionAttributes": [
          {
            "definitionId": "23",
            "name": "Some Attribute",
            "description": "Some Attribute defines how much Foo impacts Bar.",
            "enabled": true,
            "multiValue": true,
            "values": [
              "foo",
              "bar"
            ],
            "dataType": "STRING",
            "options": [
              "foo",
              "bar"
            ],
            "inputType": "TEXT"
          }
        ],
        "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "lastLoggedInUsernameBinary": "admin",
        "lastLoggedInUsernameBinaryTimestamp": "2018-10-31T18:04:13Z"
      },
      "diskEncryption": {
        "bootPartitionEncryptionDetails": {
          "partitionName": "main",
          "partitionFileVault2State": "ENCRYPTING",
          "partitionFileVault2Percent": 100
        },
        "individualRecoveryKeyValidityStatus": "VALID",
        "institutionalRecoveryKeyPresent": true,
        "diskEncryptionConfigurationName": "Test configuration",
        "fileVault2Enabled": true,
        "fileVault2EnabledUserNames": [
          "admin"
        ],
        "fileVault2EligibilityMessage": "Not a boot partition"
      },
      "purchasing": {
        "leased": true,
        "purchased": true,
        "poNumber": "53-1",
        "poDate": "2019-01-01",
        "vendor": "Example Vendor",
        "warrantyDate": "2019-01-01",
        "appleCareId": "abcd",
        "leaseDate": "2019-01-01",
        "purchasePrice": "$500",
        "lifeExpectancy": 5,
        "purchasingAccount": "admin",
        "purchasingContact": "true",
        "extensionAttributes": [
          {
            "definitionId": "23",
            "name": "Some Attribute",
            "description": "Some Attribute defines how much Foo impacts Bar.",
            "enabled": true,
            "multiValue": true,
            "values": [
              "foo",
              "bar"
            ],
            "dataType": "STRING",
            "options": [
              "foo",
              "bar"
            ],
            "inputType": "TEXT"
          }
        ]
      },
      "applications": [
        {
          "name": "Microsoft Word",
          "path": "/usr/local/app",
          "version": "1.0.0",
          "cfBundleShortVersionString": "1.0.0",
          "cfBundleVersion": "1.0.0",
          "macAppStore": true,
          "sizeMegabytes": 25,
          "bundleId": "1",
          "updateAvailable": false,
          "externalVersionId": "1"
        }
      ],
      "storage": {
        "bootDriveAvailableSpaceMegabytes": 3072,
        "disks": [
          {
            "id": "170",
            "device": "disk0",
            "model": "APPLE HDD TOSHIBA MK5065GSXF",
            "revision": "5",
            "serialNumber": "a8598f013366",
            "sizeMegabytes": 262144,
            "smartStatus": "OK",
            "type": "false",
            "partitions": [
              {
                "name": "Foo",
                "sizeMegabytes": 262144,
                "availableMegabytes": 131072,
                "partitionType": "BOOT",
                "percentUsed": 25,
                "fileVault2State": "ENCRYPTING",
                "fileVault2ProgressPercent": 45,
                "lvmManaged": true
              }
            ]
          }
        ]
      },
      "userAndLocation": {
        "username": "Madison Anderson",
        "realname": "13-inch MacBook",
        "email": "email@com.pl",
        "position": "IT Team Lead",
        "phone": "123-456-789",
        "departmentId": "1",
        "buildingId": "1",
        "room": "5",
        "extensionAttributes": [
          {
            "definitionId": "23",
            "name": "Some Attribute",
            "description": "Some Attribute defines how much Foo impacts Bar.",
            "enabled": true,
            "multiValue": true,
            "values": [
              "foo",
              "bar"
            ],
            "dataType": "STRING",
            "options": [
              "foo",
              "bar"
            ],
            "inputType": "TEXT"
          }
        ]
      },
      "configurationProfiles": [
        {
          "id": "1",
          "username": "username",
          "lastInstalled": "2018-10-31T18:04:13Z",
          "removable": true,
          "displayName": "Displayed profile",
          "profileIdentifier": "0ae590fe-9b30-11ea-bb37-0242ac130002"
        }
      ],
      "printers": [
        {
          "name": "My Printer",
          "type": "XYZ 1122",
          "uri": "ipp://10.0.0.5",
          "location": "7th floor"
        }
      ],
      "services": [
        {
          "name": "SomeService"
        }
      ],
      "hardware": {
        "make": "Apple",
        "model": "13-inch MacBook Pro (Mid 2012)",
        "modelIdentifier": "MacBookPro9,2",
        "serialNumber": "C02ZC2QYLVDL",
        "processorSpeedMhz": 2100,
        "processorCount": 2,
        "coreCount": 2,
        "processorType": "Intel Core i5",
        "processorArchitecture": "i386",
        "busSpeedMhz": 2133,
        "cacheSizeKilobytes": 3072,
        "networkAdapterType": "Foo",
        "macAddress": "6A:2C:4B:B7:65:B5",
        "altNetworkAdapterType": "Bar",
        "altMacAddress": "82:45:58:44:dc:01",
        "totalRamMegabytes": 4096,
        "openRamSlots": 0,
        "batteryCapacityPercent": 85,
        "batteryHealth": "UNKNOWN",
        "smcVersion": "2.2f38",
        "nicSpeed": "N/A",
        "opticalDrive": "MATSHITA DVD-R UJ-8A8",
        "bootRom": "MBP91.00D3.B08",
        "bleCapable": false,
        "supportsIosAppInstalls": false,
        "appleSilicon": false,
        "provisioningUdid": "00000AAA888-IH866799UUJD991",
        "extensionAttributes": [
          {
            "definitionId": "23",
            "name": "Some Attribute",
            "description": "Some Attribute defines how much Foo impacts Bar.",
            "enabled": true,
            "multiValue": true,
            "values": [
              "foo",
              "bar"
            ],
            "dataType": "STRING",
            "options": [
              "foo",
              "bar"
            ],
            "inputType": "TEXT"
          }
        ]
      },
      "localUserAccounts": [
        {
          "uid": "501",
          "userGuid": "844F1177-0CF5-40C6-901F-38EDD9969C1C",
          "username": "jamf",
          "fullName": "John Jamf",
          "admin": true,
          "homeDirectory": "/Users/jamf",
          "homeDirectorySizeMb": 131072,
          "fileVault2Enabled": true,
          "userAccountType": "LOCAL",
          "passwordMinLength": 4,
          "passwordMaxAge": 5,
          "passwordMinComplexCharacters": 5,
          "passwordHistoryDepth": 5,
          "passwordRequireAlphanumeric": true,
          "computerAzureActiveDirectoryId": "1",
          "userAzureActiveDirectoryId": "1",
          "azureActiveDirectoryId": "ACTIVATED"
        }
      ],
      "certificates": [
        {
          "commonName": "jamf.com",
          "identity": true,
          "expirationDate": "2030-10-31T18:04:13Z",
          "username": "test",
          "lifecycleStatus": "ACTIVE",
          "certificateStatus": "ISSUED",
          "subjectName": "CN=jamf.com",
          "serialNumber": "40f3d9fb",
          "sha1Fingerprint": "ed361458724d06082b2314acdb82e1f586f085f5",
          "issuedDate": "2022-05-23T14:54:10Z"
        }
      ],
      "attachments": [
        {
          "id": "1",
          "name": "Attachment.pdf",
          "fileType": "application/pdf",
          "sizeBytes": 1024
        }
      ],
      "packageReceipts": {
        "installedByJamfPro": [
          "com.jamf.protect.JamfProtect"
        ],
        "installedByInstallerSwu": [
          "com.apple.pkg.Core"
        ],
        "cached": [
          "com.jamf.protect.JamfProtect"
        ]
      },
      "security": {
        "sipStatus": "ENABLED",
        "gatekeeperStatus": "APP_STORE_AND_IDENTIFIED_DEVELOPERS",
        "xprotectVersion": "1.2.3",
        "autoLoginDisabled": false,
        "remoteDesktopEnabled": true,
        "activationLockEnabled": true,
        "recoveryLockEnabled": true,
        "firewallEnabled": true,
        "secureBootLevel": "FULL_SECURITY",
        "externalBootLevel": "ALLOW_BOOTING_FROM_EXTERNAL_MEDIA",
        "bootstrapTokenAllowed": true,
        "bootstrapTokenEscrowedStatus": "ESCROWED",
        "lastAttestationAttempt": "1970-01-01T00:00:00Z",
        "lastSuccessfulAttestation": "1970-01-01T00:00:00Z",
        "attestationStatus": "PENDING"
      },
      "operatingSystem": {
        "name": "Mac OS X",
        "version": "10.9.5",
        "build": "13A603",
        "supplementalBuildVersion": "13A953",
        "rapidSecurityResponse": "(a)",
        "activeDirectoryStatus": "Not Bound",
        "fileVault2Status": "ALL_ENCRYPTED",
        "softwareUpdateDeviceId": "J132AP",
        "extensionAttributes": [
          {
            "definitionId": "23",
            "name": "Some Attribute",
            "description": "Some Attribute defines how much Foo impacts Bar.",
            "enabled": true,
            "multiValue": true,
            "values": [
              "foo",
              "bar"
            ],
            "dataType": "STRING",
            "options": [
              "foo",
              "bar"
            ],
            "inputType": "TEXT"
          }
        ]
      },
      "licensedSoftware": [
        {
          "id": "1",
          "name": "Microsoft Word"
        }
      ],
      "ibeacons": [
        {
          "name": "room A"
        }
      ],
      "softwareUpdates": [
        {
          "name": "BEdit",
          "version": "1.15.2",
          "packageName": "com.apple.pkg.AdditionalEssentials"
        }
      ],
      "extensionAttributes": [
        {
          "definitionId": "23",
          "name": "Some Attribute",
          "description": "Some Attribute defines how much Foo impacts Bar.",
          "enabled": true,
          "multiValue": true,
          "values": [
            "foo",
            "bar"
          ],
          "dataType": "STRING",
          "options": [
            "foo",
            "bar"
          ],
          "inputType": "TEXT"
        }
      ],
      "contentCaching": {
        "computerContentCachingInformationId": "1",
        "parents": [
          {
            "contentCachingParentId": "1",
            "address": "SomeAddress",
            "alerts": {
              "contentCachingParentAlertId": "1",
              "addresses": [],
              "className": "SomeClass",
              "postDate": "2018-10-31T18:04:13Z"
            },
            "details": {
              "contentCachingParentDetailsId": "1",
              "acPower": true,
              "cacheSizeBytes": 0,
              "capabilities": {
                "contentCachingParentCapabilitiesId": "1",
                "imports": true,
                "namespaces": true,
                "personalContent": true,
                "queryParameters": true,
                "sharedContent": true,
                "prioritization": true
              },
              "portable": true,
              "localNetwork": [
                {
                  "contentCachingParentLocalNetworkId": "1",
                  "speed": 5000,
                  "wired": true
                }
              ]
            },
            "guid": "CD1E1291-4AF9-4468-B5D5-0F780C13DB2F",
            "healthy": true,
            "port": 0,
            "version": "1"
          }
        ],
        "alerts": [
          {
            "cacheBytesLimit": 0,
            "className": "SomeClass",
            "pathPreventingAccess": "/some/path",
            "postDate": "2018-10-31T18:04:13Z",
            "reservedVolumeBytes": 0,
            "resource": "SomeResource"
          }
        ],
        "activated": false,
        "active": false,
        "actualCacheBytesUsed": 0,
        "cacheDetails": [
          {
            "computerContentCachingCacheDetailsId": "1",
            "categoryName": "SomeCategory",
            "diskSpaceBytesUsed": 0
          }
        ],
        "cacheBytesFree": 23353884672,
        "cacheBytesLimit": 0,
        "cacheStatus": "OK",
        "cacheBytesUsed": 0,
        "dataMigrationCompleted": false,
        "dataMigrationProgressPercentage": 0,
        "dataMigrationError": {
          "code": 0,
          "domain": "SomeDomain",
          "userInfo": [
            {
              "key": "foo",
              "value": "bar"
            }
          ]
        },
        "maxCachePressureLast1HourPercentage": 0,
        "personalCacheBytesFree": 23353884672,
        "personalCacheBytesLimit": 0,
        "personalCacheBytesUsed": 0,
        "port": 0,
        "publicAddress": "SomeAddress",
        "registrationError": "NOT_ACTIVATED",
        "registrationResponseCode": 403,
        "registrationStarted": "2018-10-31T18:04:13Z",
        "registrationStatus": "CONTENT_CACHING_FAILED",
        "restrictedMedia": false,
        "serverGuid": "CD1E1291-4AF9-4468-B5D5-0F780C13DB2F",
        "startupStatus": "FAILED",
        "tetheratorStatus": "CONTENT_CACHING_DISABLED",
        "totalBytesAreSince": "2018-10-31T18:04:13Z",
        "totalBytesDropped": 0,
        "totalBytesImported": 0,
        "totalBytesReturnedToChildren": 0,
        "totalBytesReturnedToClients": 0,
        "totalBytesReturnedToPeers": 0,
        "totalBytesStoredFromOrigin": 0,
        "totalBytesStoredFromParents": 0,
        "totalBytesStoredFromPeers": 0
      },
      "groupMemberships": [
        {
          "groupId": "1",
          "groupName": "groupOne",
          "groupDescription": "groupOne description",
          "smartGroup": true
        }
      ]
    }
  ]
}
-----
