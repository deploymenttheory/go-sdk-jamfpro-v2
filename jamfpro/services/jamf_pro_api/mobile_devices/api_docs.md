Get Mobile Device objects
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices

Gets Mobile Device objects.

Query Params
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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
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
name
string
serialNumber
string
wifiMacAddress
string
udid
string
phoneNumber
string
model
string
modelIdentifier
string
username
string
type
string
enum
ios tvos watchos visionos unknown

managementId
string
softwareUpdateDeviceId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/mobile-devices?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "iPad",
      "serialNumber": "DMQVGC0DHLA0",
      "wifiMacAddress": "C4:84:66:92:78:00",
      "udid": "0dad565fb40b010a9e490440188063a378721069",
      "phoneNumber": "651-555-5555 Ext111",
      "model": "iPad 5th Generation (Wi-Fi)",
      "modelIdentifier": "iPad6,11",
      "username": "admin",
      "type": "ios",
      "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
      "softwareUpdateDeviceId": "J132AP"
    }
  ]
}
-----

Return paginated Mobile Device Inventory records
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/detail

Return paginated Mobile Device Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Response

200
Successful response

Response body
object
totalCount
integer
results
array of objects

iOS

tvOS

watchOS
object
mobileDeviceId
string
length ≥ 1
deviceType
string
required
Based on the value of this type either ios, appleTv, watch or visionOS objects will be populated.

hardware
object

hardware object
capacityMb
integer
availableSpaceMb
integer
usedSpacePercentage
integer
batteryLevel
integer
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

serialNumber
string
wifiMacAddress
string
bluetoothMacAddress
string
modemFirmwareVersion
string
model
string
modelIdentifier
string
modelNumber
string
bluetoothLowEnergyCapable
boolean
deviceId
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
userAndLocation
object

userAndLocation object
username
string
realName
string
emailAddress
string
position
string
phoneNumber
string
departmentId
string
buildingId
string
room
string
building
string
department
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
managementStatus
string
validationStatus
boolean
bundleSize
string
dynamicSize
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDate
date-time
profiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
removable
boolean
lastInstalled
date-time
groups
array of objects
object
groupId
string
groupName
string
length ≥ 1
groupDescription
string
length ≥ 0
smart
boolean
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
general
object

general object
udid
string
displayName
string
assetTag
string
siteId
string
lastInventoryUpdateDate
date-time
osVersion
string
osRapidSecurityResponse
string
osBuild
string
osSupplementalBuildVersion
string
softwareUpdateDeviceId
string
ipAddress
string
managed
boolean
supervised
boolean
deviceOwnershipType
string
enum
The enrollment method used for the device. Note: The PersonalDeviceProfile enrollment method was removed as of 11.25.

Institutional UserEnrollment AccountDrivenUserEnrollment AccountDrivenDeviceEnrollment

enrollmentMethodPrestage
object

enrollmentMethodPrestage object
mobileDevicePrestageId
string
profileName
string
enrollmentSessionTokenValid
boolean
lastEnrolledDate
date-time
mdmProfileExpirationDate
date-time
timeZone
string
IANA time zone database name

declarativeDeviceManagementEnabled
boolean
managementId
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
lastLoggedInUsernameSelfService
string | null
lastLoggedInUsernameSelfServiceTimestamp
date-time | null
diagnosticAndUsageReportingEnabled
boolean
appAnalyticsEnabled
boolean
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
lastCloudBackupDate
date-time
itunesStoreAccountActive
boolean
security
object
This section only available for Ios type.


security object
dataProtected
boolean
blockLevelEncryptionCapable
boolean
fileLevelEncryptionCapable
boolean
passcodePresent
boolean
passcodeCompliant
boolean
passcodeCompliantWithProfile
boolean
hardwareEncryption
integer
activationLockEnabled
boolean
jailBreakDetected
boolean
attestationStatus
string
enum
PENDING SUCCESS CERTIFICATE_INVALID DEVICE_PROPERTIES_MISMATCH MDA_UNSUPPORTED_DUE_TO_HARDWARE MDA_UNSUPPORTED_DUE_TO_SOFTWARE

lastAttestationAttemptDate
date-time
lastSuccessfulAttestationDate
date-time
passcodeLockGracePeriodEnforcedSeconds
integer
personalDeviceProfileCurrent
boolean
Defaults to false
deprecated
Deprecated as of 11.25. This field always returns false

lostModeEnabled
boolean
lostModePersistent
boolean
lostModeMessage
string
lostModePhoneNumber
string
lostModeFootnote
string
lostModeLocation
object

lostModeLocation object
lastLocationUpdate
date-time
lostModeLocationHorizontalAccuracyMeters
number
lostModeLocationVerticalAccuracyMeters
number
lostModeLocationAltitudeMeters
number
lostModeLocationSpeedMetersPerSecond
number
lostModeLocationCourseDegrees
number
lostModeLocationTimestamp
string
bootstrapTokenEscrowed
string
enum
Indicates the bootstrap token escrow status for the device

ESCROWED NOT_ESCROWED NOT_SUPPORTED

provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time

Return paginated Mobile Device Inventory records
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/detail

Return paginated Mobile Device Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Response

200
Successful response

Response body
object
totalCount
integer
results
array of objects

iOS

tvOS
object
mobileDeviceId
string
length ≥ 1
deviceType
string
required
Based on the value of this type either ios, appleTv, watch or visionOS objects will be populated.

hardware
object

hardware object
capacityMb
integer
availableSpaceMb
integer
usedSpacePercentage
integer
batteryLevel
integer
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

serialNumber
string
wifiMacAddress
string
bluetoothMacAddress
string
modemFirmwareVersion
string
model
string
modelIdentifier
string
modelNumber
string
bluetoothLowEnergyCapable
boolean
deviceId
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
userAndLocation
object

userAndLocation object
username
string
realName
string
emailAddress
string
position
string
phoneNumber
string
departmentId
string
buildingId
string
room
string
building
string
department
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
managementStatus
string
validationStatus
boolean
bundleSize
string
dynamicSize
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDate
date-time
profiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
removable
boolean
lastInstalled
date-time
groups
array of objects
object
groupId
string
groupName
string
length ≥ 1
groupDescription
string
length ≥ 0
smart
boolean
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
general
object

general object
purchasing
object

purchasing object
userProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
removable
boolean
lastInstalled
date-time
username
string

watchOS

Return paginated Mobile Device Inventory records
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/detail

Return paginated Mobile Device Inventory records

Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, General section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Response

200
Successful response

Response body
object
totalCount
integer
results
array of objects

iOS
object
mobileDeviceId
string
length ≥ 1
deviceType
string
required
Based on the value of this type either ios, appleTv, watch or visionOS objects will be populated.

hardware
object

hardware object
capacityMb
integer
availableSpaceMb
integer
usedSpacePercentage
integer
batteryLevel
integer
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

serialNumber
string
wifiMacAddress
string
bluetoothMacAddress
string
modemFirmwareVersion
string
model
string
modelIdentifier
string
modelNumber
string
bluetoothLowEnergyCapable
boolean
deviceId
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
userAndLocation
object

userAndLocation object
username
string
realName
string
emailAddress
string
position
string
phoneNumber
string
departmentId
string
buildingId
string
room
string
building
string
department
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
managementStatus
string
validationStatus
boolean
bundleSize
string
dynamicSize
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDate
date-time
profiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
removable
boolean
lastInstalled
date-time
groups
array of objects
object
groupId
string
groupName
string
length ≥ 1
groupDescription
string
length ≥ 0
smart
boolean
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
general
object

general object
security
object
This section only available for Ios type.


security object
ebooks
array of objects
object
author
string
title
string
version
string
kind
string
managementState
string
network
object
This section only avaiable for Ios type.


network object
serviceSubscriptions
array of objects
object
carrierSettingsVersion
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
subscriberCarrierNetwork
string
eid
string
iccid
string
imei
string
dataPreferred
boolean
roaming
boolean
voicePreferred
boolean
label
string
labelId
string
The unique identifier for this subscription.

meid
string
phoneNumber
string
slot
string
The description of the slot that contains the SIM representing this subscription.

provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
sharedUsers
array of objects
object
managedAppleId
string
loggedIn
boolean
dataToSync
boolean
purchasing
object

purchasing object
purchased
boolean
leased
boolean
poNumber
string
vendor
string
appleCareId
string
purchasePrice
string
purchasingAccount
string
poDate
date-time
warrantyExpiresDate
date-time
leaseExpiresDate
date-time
lifeExpectancy
integer
purchasingContact
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
inventoryDisplay
string
userProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
removable
boolean
lastInstalled
date-time
username
string

tvOS

watchOS

{
  "totalCount": 2,
  "results": [
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "sharedIpad": false,
        "diagnosticAndUsageReportingEnabled": false,
        "appAnalyticsEnabled": false,
        "residentUsers": 0,
        "quotaSize": 1024,
        "temporarySessionOnly": false,
        "temporarySessionTimeout": 30,
        "userSessionTimeout": 30,
        "syncedToComputer": 30,
        "maximumSharediPadUsersStored": 16,
        "lastBackupDate": "2022-10-17T11:48:56.307Z",
        "deviceLocatorServiceEnabled": false,
        "doNotDisturbEnabled": false,
        "cloudBackupEnabled": false,
        "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
        "locationServicesForSelfServiceMobileEnabled": false,
        "itunesStoreAccountActive": false,
        "exchangeDeviceId": "eas-1",
        "tethered": false
      },
      "security": {
        "dataProtected": false,
        "blockLevelEncryptionCapable": true,
        "fileLevelEncryptionCapable": true,
        "passcodePresent": false,
        "passcodeCompliant": true,
        "passcodeCompliantWithProfile": true,
        "hardwareEncryption": 3,
        "activationLockEnabled": false,
        "jailBreakDetected": false,
        "attestationStatus": "SUCCESS",
        "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
        "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
        "passcodeLockGracePeriodEnforcedSeconds": 3,
        "lostModeEnabled": false,
        "lostModePersistent": false,
        "lostModeMessage": "Lost phone",
        "lostModePhoneNumber": "555-555-5555",
        "lostModeFootnote": "Note",
        "lostModeLocation": {
          "lastLocationUpdate": "2019-02-04T21:09:31.661Z",
          "lostModeLocationHorizontalAccuracyMeters": 7,
          "lostModeLocationVerticalAccuracyMeters": 5,
          "lostModeLocationAltitudeMeters": 7.9,
          "lostModeLocationSpeedMetersPerSecond": 10,
          "lostModeLocationCourseDegrees": 15,
          "lostModeLocationTimestamp": "2023-04-21 12:30:00 UTC"
        },
        "bootstrapTokenEscrowed": "NOT_SUPPORTED"
      },
      "ebooks": [
        {
          "author": "Homer J Simpson",
          "title": "The Odyssey",
          "version": "0.1",
          "kind": "PDF",
          "managementState": "Managed"
        }
      ],
      "network": {
        "cellularTechnology": "Unknown",
        "voiceRoamingEnabled": false,
        "imei": "59 105109 176278 3",
        "iccid": "8991101200003204514",
        "meid": "15302309236898",
        "eid": "12547444452496388545569920380795",
        "carrierSettingsVersion": "33.1",
        "currentCarrierNetwork": "Verizon Wireless",
        "currentMobileCountryCode": "311",
        "currentMobileNetworkCode": "480",
        "homeCarrierNetwork": "Verizon",
        "homeMobileCountryCode": "US",
        "homeMobileNetworkCode": "480",
        "dataRoamingEnabled": true,
        "roaming": false,
        "personalHotspotEnabled": false,
        "phoneNumber": "555-555-5555 ext 5",
        "preferredVoiceNumber": "555-555-5555"
      },
      "serviceSubscriptions": [
        {
          "carrierSettingsVersion": "47.1",
          "currentCarrierNetwork": "T-Mobile Wi-Fi",
          "currentMobileCountryCode": "310",
          "currentMobileNetworkCode": "260",
          "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
          "eid": "89049032007008882600085727376656",
          "iccid": "8901 2605 7071 8002 130",
          "imei": "35 882334 083223 0",
          "dataPreferred": true,
          "roaming": true,
          "voicePreferred": true,
          "label": "Primary",
          "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
          "meid": "35882334083223",
          "phoneNumber": "+15128145868",
          "slot": "CTSubscriptionSlotOne"
        }
      ],
      "provisioningProfiles": [
        {
          "displayName": "jamfnation",
          "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
          "expirationDate": "2018-10-24T21:57:37Z"
        }
      ],
      "sharedUsers": [
        {
          "managedAppleId": "astark@jamf.edu",
          "loggedIn": true,
          "dataToSync": true
        }
      ],
      "purchasing": {
        "purchased": true,
        "leased": false,
        "poNumber": "8675309",
        "vendor": "Apple",
        "appleCareId": "9546567.0",
        "purchasePrice": "$399",
        "purchasingAccount": "IT Budget",
        "poDate": "2019-02-04T21:09:31.661Z",
        "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
        "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
        "lifeExpectancy": 7,
        "purchasingContact": "Nick in IT",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userProfiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z",
          "username": "admin"
        }
      ]
    },
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "airPlayPassword": "1234",
        "locales": "null",
        "languages": "english"
      },
      "purchasing": {
        "purchased": true,
        "leased": false,
        "poNumber": "8675309",
        "vendor": "Apple",
        "appleCareId": "9546567.0",
        "purchasePrice": "$399",
        "purchasingAccount": "IT Budget",
        "poDate": "2019-02-04T21:09:31.661Z",
        "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
        "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
        "lifeExpectancy": 7,
        "purchasingContact": "Nick in IT",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userProfiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z",
          "username": "admin"
        }
      ]
    },
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "diagnosticAndUsageReportingEnabled": false,
        "appAnalyticsEnabled": false,
        "deviceLocatorServiceEnabled": false,
        "doNotDisturbEnabled": false,
        "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
        "itunesStoreAccountActive": false
      },
      "security": {
        "dataProtected": false,
        "blockLevelEncryptionCapable": true,
        "fileLevelEncryptionCapable": true,
        "passcodePresent": false,
        "passcodeCompliant": true,
        "passcodeCompliantWithProfile": true,
        "hardwareEncryption": 3,
        "activationLockEnabled": false,
        "jailBreakDetected": false,
        "attestationStatus": "SUCCESS",
        "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
        "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
        "passcodeLockGracePeriodEnforcedSeconds": 3,
        "lostModeEnabled": false,
        "lostModePersistent": false,
        "lostModeMessage": "Lost phone",
        "lostModePhoneNumber": "555-555-5555",
        "lostModeFootnote": "Note",
        "lostModeLocation": {
          "lastLocationUpdate": "2019-02-04T21:09:31.661Z",
          "lostModeLocationHorizontalAccuracyMeters": 7,
          "lostModeLocationVerticalAccuracyMeters": 5,
          "lostModeLocationAltitudeMeters": 7.9,
          "lostModeLocationSpeedMetersPerSecond": 10,
          "lostModeLocationCourseDegrees": 15,
          "lostModeLocationTimestamp": "2023-04-21 12:30:00 UTC"
        },
        "bootstrapTokenEscrowed": "NOT_SUPPORTED"
      },
      "provisioningProfiles": [
        {
          "displayName": "jamfnation",
          "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
          "expirationDate": "2018-10-24T21:57:37Z"
        }
      ]
    }
  ]
}
-----
Get Mobile Device
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}

Get MobileDevice

Path Params
id
string
required
instance id of mobile device record

Responses

200
Successful response

Response body
object
id
string
name
string
serialNumber
string
wifiMacAddress
string
udid
string
phoneNumber
string
model
string
modelIdentifier
string
username
string
type
string
enum
ios tvos watchos visionos unknown

managementId
string
softwareUpdateDeviceId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-devices/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "iPad",
  "serialNumber": "DMQVGC0DHLA0",
  "wifiMacAddress": "C4:84:66:92:78:00",
  "udid": "0dad565fb40b010a9e490440188063a378721069",
  "phoneNumber": "651-555-5555 Ext111",
  "model": "iPad 5th Generation (Wi-Fi)",
  "modelIdentifier": "iPad6,11",
  "username": "admin",
  "type": "ios",
  "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
  "softwareUpdateDeviceId": "J132AP"
}
-----

Update fields on a mobile device that are allowed to be modified by users
patch
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}

Updates fields on a mobile device that are allowed to be modified by users.

Path Params
id
string
required
instance id of mobile device record

1
Body Params
name
string
Mobile Device Name. When updated, Jamf Pro sends an MDM settings command to the device (device must be supervised).

Jan's Mobile Device
enforceName
boolean
Enforce the mobile device name. Device must be supervised. If set to true, Jamf Pro will revert the Mobile Device Name to the â€˜nameâ€™ value each time the device checks in.


true
assetTag
string
8675309
siteId
string
1
timeZone
string
IANA time zone database name

Europe/Warsaw
location
object

location object
username
string
admin
realName
string
IT Bob
emailAddress
string
ITBob@jamf.com
position
string
IT Team Lead
phoneNumber
string
555-555-5555
departmentId
string
1
buildingId
string
1
room
string
4th Floor - Quad 3
updatedExtensionAttributes
array of objects

object

name
string
Example EA
type
string
enum

STRING
Allowed:

STRING

INTEGER

DATE
value
array of strings

string

EA Value

string

EA Value

ADD string
extensionAttributeCollectionAllowed
boolean

true

ADD object
ios
object

ios object
purchasing
object

purchasing object
purchased
boolean

true
leased
boolean

false
poNumber
string
8675309
vendor
string
Apple
appleCareId
string
9546567.0
purchasePrice
string
$399
purchasingAccount
string
IT Budget
poDate
date-time
2019-02-04T21:09:31.661Z
warrantyExpiresDate
date-time
2019-02-04T21:09:31.661Z
leaseExpiresDate
date-time
2019-02-04T21:09:31.661Z
lifeExpectancy
integer
7
purchasingContact
string
Nick in IT
tvos
object

tvos object
airplayPassword
password
•••••
purchasing
object

purchasing object
purchased
boolean

true
leased
boolean

true
poNumber
string
8675309
vendor
string
Apple
appleCareId
string
9546567.0
purchasePrice
string
$399
purchasingAccount
string
IT Budget
poDate
date-time
2019-02-04T21:09:31.661Z
warrantyExpiresDate
date-time
2019-02-04T21:09:31.661Z
leaseExpiresDate
date-time
2019-02-04T21:09:31.661Z
lifeExpectancy
integer
7
purchasingContact
string
Nick in IT
Responses

200
Successful response

Response body
object
id
string
name
string
Mobile device name.

enforceName
boolean
Enforce the mobile device name. Device must be supervised. If set to true, Jamf Pro will revert the Mobile Device Name to the â€˜nameâ€™ value each time the device checks in.

assetTag
string
lastInventoryUpdateTimestamp
date-time
osVersion
string
osBuild
string
osSupplementalBuildVersion
string
Collected for iOS 16 and iPadOS 16.1 or later

osRapidSecurityResponse
string
Collected for iOS 16 and iPadOS 16.1 or later

softwareUpdateDeviceId
string
serialNumber
string
udid
string
ipAddress
string
wifiMacAddress
string
bluetoothMacAddress
string
managed
boolean
timeZone
string
initialEntryTimestamp
date-time
lastEnrollmentTimestamp
date-time
mdmProfileExpirationTimestamp
date-time
deviceOwnershipLevel
string
enrollmentMethod
string
enrollmentSessionTokenValid
boolean
declarativeDeviceManagementEnabled
boolean
site
object
id
string
name
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
location
object
username
string
realName
string
emailAddress
string
position
string
phoneNumber
string
departmentId
string
buildingId
string
room
string
type
string
enum
Based on the value of this either iOS, tvOS, watch or visionOS objects will be populated.

ios tvos watchos visionos unknown

ios
object
will be populated if the type is ios or visionos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
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

lastBackupTimestamp
date-time
capacityMb
integer
availableMb
integer
percentageUsed
integer
shared
boolean
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
cloudBackupEnabled
boolean
lastCloudBackupTimestamp
date-time
locationServicesEnabled
boolean
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
computer
object

computer object
purchasing
object

purchasing object
security
object

security object
network
object

network object
serviceSubscriptions
array of objects
object
carrierSettingsVersion
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
subscriberCarrierNetwork
string
eid
string
iccid
string
imei
string
dataPreferred
boolean
roaming
boolean
voicePreferred
boolean
label
string
labelId
string
The unique identifier for this subscription.

meid
string
phoneNumber
string
slot
string
The description of the slot that contains the SIM representing this subscription.

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

ebooks
array of objects
object
author
string
title
string
version
string
mdmCapableUsers
array of objects
object
userShortName
string
managementId
string
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string
tvos
object
will be populated if the type is appleTv.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
airplayPassword
password
deviceId
string
locales
string
purchasing
object

purchasing object
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
watchos
object
will be populated if the type is watchos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
capacityMb
integer
availableMb
integer
percentageUsed
integer
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
lastCloudBackupTimestamp
date-time
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
security
object

security object
applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string
visionos
object
will be populated if the type is ios or visionos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
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

lastBackupTimestamp
date-time
capacityMb
integer
availableMb
integer
percentageUsed
integer
shared
boolean
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
cloudBackupEnabled
boolean
lastCloudBackupTimestamp
date-time
locationServicesEnabled
boolean
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
computer
object

computer object
purchasing
object

purchasing object
security
object

security object
network
object

network object
serviceSubscriptions
array of objects
object
carrierSettingsVersion
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
subscriberCarrierNetwork
string
eid
string
iccid
string
imei
string
dataPreferred
boolean
roaming
boolean
voicePreferred
boolean
label
string
labelId
string
The unique identifier for this subscription.

meid
string
phoneNumber
string
slot
string
The description of the slot that contains the SIM representing this subscription.

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

ebooks
array of objects
object
author
string
title
string
version
string
mdmCapableUsers
array of objects
object
userShortName
string
managementId
string
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-devices/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data @- <<EOF
{
  "location": {
    "username": "admin",
    "realName": "IT Bob",
    "emailAddress": "ITBob@jamf.com",
    "position": "IT Team Lead",
    "phoneNumber": "555-555-5555",
    "departmentId": "1",
    "buildingId": "1",
    "room": "4th Floor - Quad 3"
  },
  "ios": {
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    }
  },
  "tvos": {
    "purchasing": {
      "purchased": true,
      "leased": true,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "airplayPassword": "12345"
  },
  "name": "Jan's Mobile Device",
  "enforceName": true,
  "assetTag": "8675309",
  "siteId": "1",
  "timeZone": "Europe/Warsaw",
  "updatedExtensionAttributes": [
    {
      "name": "Example EA",
      "type": "STRING",
      "value": [
        "EA Value",
        "EA Value"
      ],
      "extensionAttributeCollectionAllowed": true
    }
  ]
}
EOF

{
  "id": "1",
  "name": "Jon's iPad",
  "enforceName": false,
  "assetTag": "12345",
  "lastInventoryUpdateTimestamp": "2018-10-15T16:39:56Z",
  "osVersion": "11.4",
  "osBuild": "15F79",
  "osSupplementalBuildVersion": "20B101",
  "osRapidSecurityResponse": "(a)",
  "softwareUpdateDeviceId": "J132AP",
  "serialNumber": "DMQVGC0DHLF0",
  "udid": "0dad565fb40b010a9e490440188063a378721069",
  "ipAddress": "10.0.0.1",
  "wifiMacAddress": "ee:00:7c:f0:e5:ff",
  "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
  "managed": true,
  "timeZone": "Europe/Warsaw",
  "initialEntryTimestamp": "2018-10-15T16:39:56.307Z",
  "lastEnrollmentTimestamp": "2018-10-15T16:39:56.307Z",
  "mdmProfileExpirationTimestamp": "2018-10-15T16:39:56.307Z",
  "deviceOwnershipLevel": "institutional",
  "enrollmentMethod": "User-initiated - no invitation",
  "enrollmentSessionTokenValid": false,
  "declarativeDeviceManagementEnabled": true,
  "site": {
    "id": "1",
    "name": "Eau Claire"
  },
  "extensionAttributes": [
    {
      "id": "1",
      "name": "Example EA",
      "type": "STRING",
      "value": [
        "EA Value"
      ],
      "extensionAttributeCollectionAllowed": false
    }
  ],
  "location": {
    "username": "admin",
    "realName": "IT Bob",
    "emailAddress": "ITBob@jamf.com",
    "position": "IT Team Lead",
    "phoneNumber": "555-555-5555",
    "departmentId": "1",
    "buildingId": "1",
    "room": "4th Floor - Quad 3"
  },
  "type": "ios",
  "ios": {
    "model": "iPad 5th Generation (Wi-Fi)",
    "modelIdentifier": "ipad6,11",
    "modelNumber": "MP2F2LL",
    "supervised": true,
    "batteryLevel": 100,
    "batteryHealth": "UNKNOWN",
    "lastBackupTimestamp": "2018-10-15T16:39:56Z",
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "shared": false,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "cloudBackupEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "locationServicesEnabled": false,
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "computer": {
      "id": "1",
      "name": "A name"
    },
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "network": {
      "cellularTechnology": "Unknown",
      "voiceRoamingEnabled": false,
      "imei": "59 105109 176278 3",
      "iccid": "8991101200003204514",
      "meid": "15302309236898",
      "eid": "12547444452496388545569920380795",
      "carrierSettingsVersion": "33.1",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "roaming": false,
      "personalHotspotEnabled": false,
      "phoneNumber": "555-555-5555 ext 5",
      "preferredVoiceNumber": "555-555-5555"
    },
    "serviceSubscriptions": [
      {
        "carrierSettingsVersion": "47.1",
        "currentCarrierNetwork": "T-Mobile Wi-Fi",
        "currentMobileCountryCode": "310",
        "currentMobileNetworkCode": "260",
        "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
        "eid": "89049032007008882600085727376656",
        "iccid": "8901 2605 7071 8002 130",
        "imei": "35 882334 083223 0",
        "dataPreferred": true,
        "roaming": true,
        "voicePreferred": true,
        "label": "Primary",
        "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
        "meid": "35882334083223",
        "phoneNumber": "+15128145868",
        "slot": "CTSubscriptionSlotOne"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "ebooks": [
      {
        "author": "Homer J Simpson",
        "title": "The Odyssey",
        "version": "0.1"
      }
    ],
    "mdmCapableUsers": [
      {
        "userShortName": "testUser",
        "managementId": "8835fe1c-eb9d-42e5-8277-f708ae011a9d"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  },
  "tvos": {
    "model": "Apple TV 3rd Generation Rev 2",
    "modelIdentifier": "AppleTV3,2",
    "modelNumber": "MD199LL",
    "supervised": true,
    "airplayPassword": "1234",
    "deviceId": "1",
    "locales": "null",
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ]
  },
  "watchos": {
    "model": "Apple Watch Ultra",
    "modelIdentifier": "Watch6,18",
    "modelNumber": "A2622",
    "supervised": true,
    "batteryLevel": 100,
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  },
  "visionos": {
    "model": "iPad 5th Generation (Wi-Fi)",
    "modelIdentifier": "ipad6,11",
    "modelNumber": "MP2F2LL",
    "supervised": true,
    "batteryLevel": 100,
    "batteryHealth": "UNKNOWN",
    "lastBackupTimestamp": "2018-10-15T16:39:56Z",
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "shared": false,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "cloudBackupEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "locationServicesEnabled": false,
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "computer": {
      "id": "1",
      "name": "A name"
    },
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "network": {
      "cellularTechnology": "Unknown",
      "voiceRoamingEnabled": false,
      "imei": "59 105109 176278 3",
      "iccid": "8991101200003204514",
      "meid": "15302309236898",
      "eid": "12547444452496388545569920380795",
      "carrierSettingsVersion": "33.1",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "roaming": false,
      "personalHotspotEnabled": false,
      "phoneNumber": "555-555-5555 ext 5",
      "preferredVoiceNumber": "555-555-5555"
    },
    "serviceSubscriptions": [
      {
        "carrierSettingsVersion": "47.1",
        "currentCarrierNetwork": "T-Mobile Wi-Fi",
        "currentMobileCountryCode": "310",
        "currentMobileNetworkCode": "260",
        "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
        "eid": "89049032007008882600085727376656",
        "iccid": "8901 2605 7071 8002 130",
        "imei": "35 882334 083223 0",
        "dataPreferred": true,
        "roaming": true,
        "voicePreferred": true,
        "label": "Primary",
        "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
        "meid": "35882334083223",
        "phoneNumber": "+15128145868",
        "slot": "CTSubscriptionSlotOne"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "ebooks": [
      {
        "author": "Homer J Simpson",
        "title": "The Odyssey",
        "version": "0.1"
      }
    ],
    "mdmCapableUsers": [
      {
        "userShortName": "testUser",
        "managementId": "8835fe1c-eb9d-42e5-8277-f708ae011a9d"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  }
}
-----
Get Mobile Device
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}/detail

Get MobileDevice

Path Params
id
string
required
instance id of mobile device record

1
Responses

200
Successful response

Response body
object
id
string
name
string
Mobile device name.

enforceName
boolean
Enforce the mobile device name. Device must be supervised. If set to true, Jamf Pro will revert the Mobile Device Name to the â€˜nameâ€™ value each time the device checks in.

assetTag
string
lastInventoryUpdateTimestamp
date-time
osVersion
string
osBuild
string
osSupplementalBuildVersion
string
Collected for iOS 16 and iPadOS 16.1 or later

osRapidSecurityResponse
string
Collected for iOS 16 and iPadOS 16.1 or later

softwareUpdateDeviceId
string
serialNumber
string
udid
string
ipAddress
string
wifiMacAddress
string
bluetoothMacAddress
string
managed
boolean
timeZone
string
initialEntryTimestamp
date-time
lastEnrollmentTimestamp
date-time
mdmProfileExpirationTimestamp
date-time
deviceOwnershipLevel
string
enrollmentMethod
string
enrollmentSessionTokenValid
boolean
declarativeDeviceManagementEnabled
boolean
site
object
id
string
name
string
extensionAttributes
array of objects
object
id
string
name
string
type
string
enum
STRING INTEGER DATE

value
array of strings
extensionAttributeCollectionAllowed
boolean
location
object
username
string
realName
string
emailAddress
string
position
string
phoneNumber
string
departmentId
string
buildingId
string
room
string
type
string
enum
Based on the value of this either iOS, tvOS, watch or visionOS objects will be populated.

ios tvos watchos visionos unknown

ios
object
will be populated if the type is ios or visionos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
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

lastBackupTimestamp
date-time
capacityMb
integer
availableMb
integer
percentageUsed
integer
shared
boolean
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
cloudBackupEnabled
boolean
lastCloudBackupTimestamp
date-time
locationServicesEnabled
boolean
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
computer
object

computer object
purchasing
object

purchasing object
security
object

security object
network
object

network object
serviceSubscriptions
array of objects
object
carrierSettingsVersion
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
subscriberCarrierNetwork
string
eid
string
iccid
string
imei
string
dataPreferred
boolean
roaming
boolean
voicePreferred
boolean
label
string
labelId
string
The unique identifier for this subscription.

meid
string
phoneNumber
string
slot
string
The description of the slot that contains the SIM representing this subscription.

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

ebooks
array of objects
object
author
string
title
string
version
string
mdmCapableUsers
array of objects
object
userShortName
string
managementId
string
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string
tvos
object
will be populated if the type is appleTv.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
airplayPassword
password
deviceId
string
locales
string
purchasing
object

purchasing object
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
watchos
object
will be populated if the type is watchos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
capacityMb
integer
availableMb
integer
percentageUsed
integer
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
lastCloudBackupTimestamp
date-time
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
security
object

security object
applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string
visionos
object
will be populated if the type is ios or visionos.

model
string
modelIdentifier
string
modelNumber
string
supervised
boolean
batteryLevel
integer
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

lastBackupTimestamp
date-time
capacityMb
integer
availableMb
integer
percentageUsed
integer
shared
boolean
deviceLocatorServiceEnabled
boolean
doNotDisturbEnabled
boolean
cloudBackupEnabled
boolean
lastCloudBackupTimestamp
date-time
locationServicesEnabled
boolean
iTunesStoreAccountActive
boolean
bleCapable
boolean
unlockToken
string
computer
object

computer object
purchasing
object

purchasing object
security
object

security object
network
object

network object
serviceSubscriptions
array of objects
object
carrierSettingsVersion
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
subscriberCarrierNetwork
string
eid
string
iccid
string
imei
string
dataPreferred
boolean
roaming
boolean
voicePreferred
boolean
label
string
labelId
string
The unique identifier for this subscription.

meid
string
phoneNumber
string
slot
string
The description of the slot that contains the SIM representing this subscription.

applications
array of objects
object
identifier
string
name
string
version
string
shortVersion
string
certificates
array of objects
object
commonName
string
identity
boolean
expirationDateEpoch
date-time
subjectName
string
serialNumber
string
sha1Fingerprint
string
issuedDateEpoch
string
certificateStatus
string
enum
EXPIRING EXPIRED REVOKED PENDING_REVOKE ISSUED

lifecycleStatus
string
enum
ACTIVE INACTIVE

ebooks
array of objects
object
author
string
title
string
version
string
mdmCapableUsers
array of objects
object
userShortName
string
managementId
string
configurationProfiles
array of objects
object
displayName
string
version
string
uuid
string
identifier
string
provisioningProfiles
array of objects
object
displayName
string
uuid
string
expirationDate
date-time
attachments
array of objects
object
name
string
id
string
managementId
string
groups
array of objects
object
groupId
string
groupName
string
length ≥ 1
groupDescription
string
length ≥ 0
smart
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-devices/1/detail \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Jon's iPad",
  "enforceName": false,
  "assetTag": "12345",
  "lastInventoryUpdateTimestamp": "2018-10-15T16:39:56Z",
  "osVersion": "11.4",
  "osBuild": "15F79",
  "osSupplementalBuildVersion": "20B101",
  "osRapidSecurityResponse": "(a)",
  "softwareUpdateDeviceId": "J132AP",
  "serialNumber": "DMQVGC0DHLF0",
  "udid": "0dad565fb40b010a9e490440188063a378721069",
  "ipAddress": "10.0.0.1",
  "wifiMacAddress": "ee:00:7c:f0:e5:ff",
  "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
  "managed": true,
  "timeZone": "Europe/Warsaw",
  "initialEntryTimestamp": "2018-10-15T16:39:56.307Z",
  "lastEnrollmentTimestamp": "2018-10-15T16:39:56.307Z",
  "mdmProfileExpirationTimestamp": "2018-10-15T16:39:56.307Z",
  "deviceOwnershipLevel": "institutional",
  "enrollmentMethod": "User-initiated - no invitation",
  "enrollmentSessionTokenValid": false,
  "declarativeDeviceManagementEnabled": true,
  "site": {
    "id": "1",
    "name": "Eau Claire"
  },
  "extensionAttributes": [
    {
      "id": "1",
      "name": "Example EA",
      "type": "STRING",
      "value": [
        "EA Value"
      ],
      "extensionAttributeCollectionAllowed": false
    }
  ],
  "location": {
    "username": "admin",
    "realName": "IT Bob",
    "emailAddress": "ITBob@jamf.com",
    "position": "IT Team Lead",
    "phoneNumber": "555-555-5555",
    "departmentId": "1",
    "buildingId": "1",
    "room": "4th Floor - Quad 3"
  },
  "type": "ios",
  "ios": {
    "model": "iPad 5th Generation (Wi-Fi)",
    "modelIdentifier": "ipad6,11",
    "modelNumber": "MP2F2LL",
    "supervised": true,
    "batteryLevel": 100,
    "batteryHealth": "UNKNOWN",
    "lastBackupTimestamp": "2018-10-15T16:39:56Z",
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "shared": false,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "cloudBackupEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "locationServicesEnabled": false,
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "computer": {
      "id": "1",
      "name": "A name"
    },
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "network": {
      "cellularTechnology": "Unknown",
      "voiceRoamingEnabled": false,
      "imei": "59 105109 176278 3",
      "iccid": "8991101200003204514",
      "meid": "15302309236898",
      "eid": "12547444452496388545569920380795",
      "carrierSettingsVersion": "33.1",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "roaming": false,
      "personalHotspotEnabled": false,
      "phoneNumber": "555-555-5555 ext 5",
      "preferredVoiceNumber": "555-555-5555"
    },
    "serviceSubscriptions": [
      {
        "carrierSettingsVersion": "47.1",
        "currentCarrierNetwork": "T-Mobile Wi-Fi",
        "currentMobileCountryCode": "310",
        "currentMobileNetworkCode": "260",
        "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
        "eid": "89049032007008882600085727376656",
        "iccid": "8901 2605 7071 8002 130",
        "imei": "35 882334 083223 0",
        "dataPreferred": true,
        "roaming": true,
        "voicePreferred": true,
        "label": "Primary",
        "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
        "meid": "35882334083223",
        "phoneNumber": "+15128145868",
        "slot": "CTSubscriptionSlotOne"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "ebooks": [
      {
        "author": "Homer J Simpson",
        "title": "The Odyssey",
        "version": "0.1"
      }
    ],
    "mdmCapableUsers": [
      {
        "userShortName": "testUser",
        "managementId": "8835fe1c-eb9d-42e5-8277-f708ae011a9d"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  },
  "tvos": {
    "model": "Apple TV 3rd Generation Rev 2",
    "modelIdentifier": "AppleTV3,2",
    "modelNumber": "MD199LL",
    "supervised": true,
    "airplayPassword": "1234",
    "deviceId": "1",
    "locales": "null",
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ]
  },
  "watchos": {
    "model": "Apple Watch Ultra",
    "modelIdentifier": "Watch6,18",
    "modelNumber": "A2622",
    "supervised": true,
    "batteryLevel": 100,
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  },
  "visionos": {
    "model": "iPad 5th Generation (Wi-Fi)",
    "modelIdentifier": "ipad6,11",
    "modelNumber": "MP2F2LL",
    "supervised": true,
    "batteryLevel": 100,
    "batteryHealth": "UNKNOWN",
    "lastBackupTimestamp": "2018-10-15T16:39:56Z",
    "capacityMb": 27503,
    "availableMb": 26646,
    "percentageUsed": 3,
    "shared": false,
    "deviceLocatorServiceEnabled": false,
    "doNotDisturbEnabled": false,
    "cloudBackupEnabled": false,
    "lastCloudBackupTimestamp": "2018-10-15T16:39:56.307Z",
    "locationServicesEnabled": false,
    "iTunesStoreAccountActive": false,
    "bleCapable": false,
    "unlockToken": "VU5MT0NLVE9LRU4=",
    "computer": {
      "id": "1",
      "name": "A name"
    },
    "purchasing": {
      "purchased": true,
      "leased": false,
      "poNumber": "8675309",
      "vendor": "Apple",
      "appleCareId": "9546567.0",
      "purchasePrice": "$399",
      "purchasingAccount": "IT Budget",
      "poDate": "2019-02-04T21:09:31.661Z",
      "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
      "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
      "lifeExpectancy": 7,
      "purchasingContact": "Nick in IT"
    },
    "security": {
      "dataProtected": false,
      "blockLevelEncryptionCapable": true,
      "fileLevelEncryptionCapable": true,
      "passcodePresent": false,
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "hardwareEncryption": 3,
      "activationLockEnabled": false,
      "jailBreakDetected": false,
      "attestationStatus": "SUCCESS",
      "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
      "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
      "bootstrapToken": "dGVzdCB0b2tlbg==",
      "bootstrapTokenEscrowed": "NOT_SUPPORTED"
    },
    "network": {
      "cellularTechnology": "Unknown",
      "voiceRoamingEnabled": false,
      "imei": "59 105109 176278 3",
      "iccid": "8991101200003204514",
      "meid": "15302309236898",
      "eid": "12547444452496388545569920380795",
      "carrierSettingsVersion": "33.1",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "roaming": false,
      "personalHotspotEnabled": false,
      "phoneNumber": "555-555-5555 ext 5",
      "preferredVoiceNumber": "555-555-5555"
    },
    "serviceSubscriptions": [
      {
        "carrierSettingsVersion": "47.1",
        "currentCarrierNetwork": "T-Mobile Wi-Fi",
        "currentMobileCountryCode": "310",
        "currentMobileNetworkCode": "260",
        "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
        "eid": "89049032007008882600085727376656",
        "iccid": "8901 2605 7071 8002 130",
        "imei": "35 882334 083223 0",
        "dataPreferred": true,
        "roaming": true,
        "voicePreferred": true,
        "label": "Primary",
        "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
        "meid": "35882334083223",
        "phoneNumber": "+15128145868",
        "slot": "CTSubscriptionSlotOne"
      }
    ],
    "applications": [
      {
        "identifier": "com.apple.airport.mobileairportutility",
        "name": "AirPort Utility",
        "version": "135.24",
        "shortVersion": "7.0"
      }
    ],
    "certificates": [
      {
        "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
        "identity": false,
        "expirationDateEpoch": "2030-10-31T18:04:13Z",
        "subjectName": "CN=Fleet Docker Jamf Pro JSS Built-in Certificate Authority",
        "serialNumber": "5c28fdae",
        "sha1Fingerprint": "050cfe8ec9d170be7bf8f1a3cac2c52f3c6ddb20",
        "issuedDateEpoch": "2022-05-23T14:54:10Z",
        "certificateStatus": "ISSUED",
        "lifecycleStatus": "ACTIVE"
      }
    ],
    "ebooks": [
      {
        "author": "Homer J Simpson",
        "title": "The Odyssey",
        "version": "0.1"
      }
    ],
    "mdmCapableUsers": [
      {
        "userShortName": "testUser",
        "managementId": "8835fe1c-eb9d-42e5-8277-f708ae011a9d"
      }
    ],
    "configurationProfiles": [
      {
        "displayName": "Test WiFi",
        "version": "1",
        "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
        "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA"
      }
    ],
    "provisioningProfiles": [
      {
        "displayName": "jamfnation",
        "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
        "expirationDate": "2018-10-24T21:57:37Z"
      }
    ],
    "attachments": [
      {
        "name": "Bob's Attachment",
        "id": "1"
      }
    ]
  },
  "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
  "groups": [
    {
      "groupId": "1",
      "groupName": "Test Group",
      "groupDescription": "Test Group Description",
      "smart": false
    }
  ]
}
-----
Erase a Mobile Device
post
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}/erase

Erase a Mobile Device

Path Params
id
string
required
Id of the Mobile Device to erase

1
Body Params
Options for eraseDevice command

preserveDataPlan
boolean
Defaults to false
If 'true', preserve the data plan on an iPhone or iPad with eSIM functionality, if one exists.


false
disallowProximitySetup
boolean
Defaults to false
If 'true', disable Proximity Setup on the next reboot and skip the pane in Setup Assistant.


false
clearActivationLock
boolean
Defaults to false
Clear the activation lock on the device.


false
returnToService
boolean
Defaults to false
If 'true', the device will be returned to service after the erase is complete.


false
Responses

200
Erase command was queued for this device.

Response body
object
deviceId
string
required
Id of the mobile device for which eraseDevice command was queued

commandUuid
string
required
Uuid of the queued eraseDevice command

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-devices/1/erase \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "preserveDataPlan": false,
  "disallowProximitySetup": false,
  "clearActivationLock": false,
  "returnToService": false
}
'

{
  "deviceId": "1234",
  "commandUuid": "1234-5678-90ab-cdef-1234567890ab"
}
-----
Return paginated Mobile Device Inventory records of all paired devices for the device
get
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}/paired-devices

Return paginated Mobile Device Inventory records of all paired devices for the device

Path Params
id
string
required
instance id of mobile device record

1
Query Params
section
array of strings
Defaults to GENERAL
section of mobile device details, if not specified, Paired Devices section data is returned. Multiple section parameters are supported, e.g. section=GENERAL&section=HARDWARE


string


USER_AND_LOCATION

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is displayName:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, languages, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, serialNumber, supervised, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, languages, lastInventoryUpdateDate, locales, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, serialNumber, supervised, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp, groupId, groupName

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Response

200
Successful response

Response body
object
totalCount
integer
results
array of objects

iOS

tvOS

watchOS

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/mobile-devices/1/paired-devices?section=USER_AND_LOCATION&page=0&page-size=100&sort=displayName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 2,
  "results": [
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "sharedIpad": false,
        "diagnosticAndUsageReportingEnabled": false,
        "appAnalyticsEnabled": false,
        "residentUsers": 0,
        "quotaSize": 1024,
        "temporarySessionOnly": false,
        "temporarySessionTimeout": 30,
        "userSessionTimeout": 30,
        "syncedToComputer": 30,
        "maximumSharediPadUsersStored": 16,
        "lastBackupDate": "2022-10-17T11:48:56.307Z",
        "deviceLocatorServiceEnabled": false,
        "doNotDisturbEnabled": false,
        "cloudBackupEnabled": false,
        "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
        "locationServicesForSelfServiceMobileEnabled": false,
        "itunesStoreAccountActive": false,
        "exchangeDeviceId": "eas-1",
        "tethered": false
      },
      "security": {
        "dataProtected": false,
        "blockLevelEncryptionCapable": true,
        "fileLevelEncryptionCapable": true,
        "passcodePresent": false,
        "passcodeCompliant": true,
        "passcodeCompliantWithProfile": true,
        "hardwareEncryption": 3,
        "activationLockEnabled": false,
        "jailBreakDetected": false,
        "attestationStatus": "SUCCESS",
        "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
        "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
        "passcodeLockGracePeriodEnforcedSeconds": 3,
        "lostModeEnabled": false,
        "lostModePersistent": false,
        "lostModeMessage": "Lost phone",
        "lostModePhoneNumber": "555-555-5555",
        "lostModeFootnote": "Note",
        "lostModeLocation": {
          "lastLocationUpdate": "2019-02-04T21:09:31.661Z",
          "lostModeLocationHorizontalAccuracyMeters": 7,
          "lostModeLocationVerticalAccuracyMeters": 5,
          "lostModeLocationAltitudeMeters": 7.9,
          "lostModeLocationSpeedMetersPerSecond": 10,
          "lostModeLocationCourseDegrees": 15,
          "lostModeLocationTimestamp": "2023-04-21 12:30:00 UTC"
        },
        "bootstrapTokenEscrowed": "NOT_SUPPORTED"
      },
      "ebooks": [
        {
          "author": "Homer J Simpson",
          "title": "The Odyssey",
          "version": "0.1",
          "kind": "PDF",
          "managementState": "Managed"
        }
      ],
      "network": {
        "cellularTechnology": "Unknown",
        "voiceRoamingEnabled": false,
        "imei": "59 105109 176278 3",
        "iccid": "8991101200003204514",
        "meid": "15302309236898",
        "eid": "12547444452496388545569920380795",
        "carrierSettingsVersion": "33.1",
        "currentCarrierNetwork": "Verizon Wireless",
        "currentMobileCountryCode": "311",
        "currentMobileNetworkCode": "480",
        "homeCarrierNetwork": "Verizon",
        "homeMobileCountryCode": "US",
        "homeMobileNetworkCode": "480",
        "dataRoamingEnabled": true,
        "roaming": false,
        "personalHotspotEnabled": false,
        "phoneNumber": "555-555-5555 ext 5",
        "preferredVoiceNumber": "555-555-5555"
      },
      "serviceSubscriptions": [
        {
          "carrierSettingsVersion": "47.1",
          "currentCarrierNetwork": "T-Mobile Wi-Fi",
          "currentMobileCountryCode": "310",
          "currentMobileNetworkCode": "260",
          "subscriberCarrierNetwork": "T-Mobile Wi-Fi",
          "eid": "89049032007008882600085727376656",
          "iccid": "8901 2605 7071 8002 130",
          "imei": "35 882334 083223 0",
          "dataPreferred": true,
          "roaming": true,
          "voicePreferred": true,
          "label": "Primary",
          "labelId": "D1F4AEC5-2FCD-4A6D-A09E-A940F60F856B",
          "meid": "35882334083223",
          "phoneNumber": "+15128145868",
          "slot": "CTSubscriptionSlotOne"
        }
      ],
      "provisioningProfiles": [
        {
          "displayName": "jamfnation",
          "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
          "expirationDate": "2018-10-24T21:57:37Z"
        }
      ],
      "sharedUsers": [
        {
          "managedAppleId": "astark@jamf.edu",
          "loggedIn": true,
          "dataToSync": true
        }
      ],
      "purchasing": {
        "purchased": true,
        "leased": false,
        "poNumber": "8675309",
        "vendor": "Apple",
        "appleCareId": "9546567.0",
        "purchasePrice": "$399",
        "purchasingAccount": "IT Budget",
        "poDate": "2019-02-04T21:09:31.661Z",
        "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
        "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
        "lifeExpectancy": 7,
        "purchasingContact": "Nick in IT",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userProfiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z",
          "username": "admin"
        }
      ]
    },
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "airPlayPassword": "1234",
        "locales": "null",
        "languages": "english"
      },
      "purchasing": {
        "purchased": true,
        "leased": false,
        "poNumber": "8675309",
        "vendor": "Apple",
        "appleCareId": "9546567.0",
        "purchasePrice": "$399",
        "purchasingAccount": "IT Budget",
        "poDate": "2019-02-04T21:09:31.661Z",
        "warrantyExpiresDate": "2019-02-04T21:09:31.661Z",
        "leaseExpiresDate": "2019-02-04T21:09:31.661Z",
        "lifeExpectancy": 7,
        "purchasingContact": "Nick in IT",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userProfiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z",
          "username": "admin"
        }
      ]
    },
    {
      "mobileDeviceId": "1",
      "deviceType": "iOS",
      "hardware": {
        "capacityMb": 100,
        "availableSpaceMb": 30,
        "usedSpacePercentage": 70,
        "batteryLevel": 60,
        "batteryHealth": "UNKNOWN",
        "serialNumber": "5c28fdae",
        "wifiMacAddress": "ee:00:7c:f0:e5:ff",
        "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
        "modemFirmwareVersion": "iPad7,11",
        "model": "iPad 7th Generation (Wi-Fi)",
        "modelIdentifier": "iPad7,11",
        "modelNumber": "MW742LL",
        "bluetoothLowEnergyCapable": false,
        "deviceId": "c6a49c6d-8c09-4d71-a37d-2f6a9dfbb69b",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "userAndLocation": {
        "username": "admin",
        "realName": "IT Bob",
        "emailAddress": "ITBob@jamf.com",
        "position": "IT Team Lead",
        "phoneNumber": "555-555-5555",
        "departmentId": "1",
        "buildingId": "1",
        "room": "room",
        "building": "Building 1",
        "department": "Department 1",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ]
      },
      "applications": [
        {
          "identifier": "com.apple.airport.mobileairportutility",
          "name": "AirPort Utility",
          "version": "135.24",
          "shortVersion": "7.0",
          "managementStatus": "Managed",
          "validationStatus": true,
          "bundleSize": "1024",
          "dynamicSize": "1423"
        }
      ],
      "certificates": [
        {
          "commonName": "3B259E4B-FAD5-4860-B1DD-336ADA786EBA",
          "identity": false,
          "expirationDate": "2019-02-04T21:09:31.661Z"
        }
      ],
      "profiles": [
        {
          "displayName": "Test WiFi",
          "version": "1",
          "uuid": "D29DD9FB-0D5B-422F-A3A2-ABBC5848E949",
          "identifier": "ac2-server4.D0EFAC2D-326C-4BB6-87E6-2BCB88490AAA",
          "removable": true,
          "lastInstalled": "2019-02-04T21:09:31.661Z"
        }
      ],
      "groups": [
        {
          "groupId": "1",
          "groupName": "Test Group",
          "groupDescription": "Test Group Description",
          "smart": false
        }
      ],
      "extensionAttributes": [
        {
          "id": "1",
          "name": "Example EA",
          "type": "STRING",
          "value": [
            "EA Value"
          ],
          "extensionAttributeCollectionAllowed": false,
          "inventoryDisplay": "General"
        }
      ],
      "general": {
        "udid": "0dad565fb40b010a9e490440188063a378721069",
        "displayName": "Banezicron",
        "assetTag": "8675309",
        "siteId": "-1",
        "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
        "osVersion": "11.4",
        "osRapidSecurityResponse": "(a)",
        "osBuild": "15F79",
        "osSupplementalBuildVersion": "22A103310o",
        "softwareUpdateDeviceId": "J132AP",
        "ipAddress": "10.0.0.1",
        "managed": true,
        "supervised": true,
        "deviceOwnershipType": "Institutional",
        "enrollmentMethodPrestage": {
          "mobileDevicePrestageId": "5",
          "profileName": "All Mobiles"
        },
        "enrollmentSessionTokenValid": false,
        "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
        "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
        "timeZone": "Europe/Warsaw",
        "declarativeDeviceManagementEnabled": true,
        "managementId": "9932fad3-29e9-4b71-bc7c-77dcefce819d",
        "extensionAttributes": [
          {
            "id": "1",
            "name": "Example EA",
            "type": "STRING",
            "value": [
              "EA Value"
            ],
            "extensionAttributeCollectionAllowed": false,
            "inventoryDisplay": "General"
          }
        ],
        "lastLoggedInUsernameSelfService": "admin",
        "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
        "diagnosticAndUsageReportingEnabled": false,
        "appAnalyticsEnabled": false,
        "deviceLocatorServiceEnabled": false,
        "doNotDisturbEnabled": false,
        "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
        "itunesStoreAccountActive": false
      },
      "security": {
        "dataProtected": false,
        "blockLevelEncryptionCapable": true,
        "fileLevelEncryptionCapable": true,
        "passcodePresent": false,
        "passcodeCompliant": true,
        "passcodeCompliantWithProfile": true,
        "hardwareEncryption": 3,
        "activationLockEnabled": false,
        "jailBreakDetected": false,
        "attestationStatus": "SUCCESS",
        "lastAttestationAttemptDate": "2019-02-04T21:09:31.661Z",
        "lastSuccessfulAttestationDate": "2019-02-04T21:09:31.661Z",
        "passcodeLockGracePeriodEnforcedSeconds": 3,
        "lostModeEnabled": false,
        "lostModePersistent": false,
        "lostModeMessage": "Lost phone",
        "lostModePhoneNumber": "555-555-5555",
        "lostModeFootnote": "Note",
        "lostModeLocation": {
          "lastLocationUpdate": "2019-02-04T21:09:31.661Z",
          "lostModeLocationHorizontalAccuracyMeters": 7,
          "lostModeLocationVerticalAccuracyMeters": 5,
          "lostModeLocationAltitudeMeters": 7.9,
          "lostModeLocationSpeedMetersPerSecond": 10,
          "lostModeLocationCourseDegrees": 15,
          "lostModeLocationTimestamp": "2023-04-21 12:30:00 UTC"
        },
        "bootstrapTokenEscrowed": "NOT_SUPPORTED"
      },
      "provisioningProfiles": [
        {
          "displayName": "jamfnation",
          "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
          "expirationDate": "2018-10-24T21:57:37Z"
        }
      ]
    }
  ]
}
-----
Unmanage a Mobile Device
post
https://yourServer.jamfcloud.com/api/v2/mobile-devices/{id}/unmanage

Unmanage a Mobile Device

Path Params
id
string
required
Id of the mobile device to remove the MDM profile from

1
Responses

200
Command to remove the mdm profile was queued

Response body
object
deviceId
string
required
Id of the mobile device whose MDM profile was removed

commandUuid
string
required
Uuid of the command queued that removes the MDM profile

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-devices/1/unmanage \
     --header 'accept: application/json'

{
  "deviceId": "1234",
  "commandUuid": "1234-5678-90ab-cdef-1234567890ab"
}
-----
