Return the list of all Mobile Device Groups
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups

Returns the list of all mobile device groups.

Response

200
Success

Response body
array of objects
object
id
integer
name
string
description
string
isSmartGroup
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups \
     --header 'accept: application/json'

[
  {
    "id": 1,
    "name": "All Managed iPads",
    "description": "A group containing all managed iPads",
    "isSmartGroup": true
  }
]
-----
Get Static Group Membership by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-group-membership/{id}

Get Static Group Membership by Id

Path Params
id
string
required
instance id of static-group

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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osRapidSecurityResponse, osSupplementalBuildVersion, osVersion, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be sorted by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be filtered by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Responses

200
Successful response

Response body
object
totalCount
integer
results
array of objects
object
mobileDeviceId
string
length ≥ 1
udid
string
airPlayPassword
password
appAnalyticsEnabled
boolean
assetTag
string
availableSpaceMb
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

bluetoothLowEnergyCapable
boolean
bluetoothMacAddress
string
capacityMb
integer
lostModeEnabledDate
date-time
declarativeDeviceManagementEnabled
boolean
deviceId
string
deviceLocatorServiceEnabled
boolean
deviceOwnershipType
string
devicePhoneNumber
string
diagnosticAndUsageReportingEnabled
boolean
displayName
string
doNotDisturbEnabled
boolean
enrollmentSessionTokenValid
boolean
exchangeDeviceId
string
cloudBackupEnabled
boolean
osBuild
string
osSupplementalBuildVersion
string
osRapidSecurityResponse
string
osVersion
string
ipAddress
string
itunesStoreAccountActive
boolean
jamfParentPairings
integer
languages
string
lastBackupDate
date-time
lastEnrolledDate
date-time
lastCloudBackupDate
date-time
lastInventoryUpdateDate
date-time
locales
string
locationServicesForSelfServiceMobileEnabled
boolean
lostModeEnabled
boolean
managed
boolean
managementId
string
mdmProfileExpirationDate
date-time
model
string
modelIdentifier
string
modelNumber
string
modemFirmwareVersion
string
pairedDevices
integer
quotaSize
integer
residentUsers
integer
serialNumber
string
sharedIpad
boolean
supervised
boolean
tethered
boolean
timeZone
string
usedSpacePercentage
integer
wifiMacAddress
string
building
string
department
string
emailAddress
string
fullName
string
position
string
room
string
userPhoneNumber
string
username
string
appleCareId
string
leaseExpirationDate
date-time
lifeExpectancyYears
integer
poDate
date-time
poNumber
string
purchasePrice
string
purchasedOrLeased
boolean
purchasingAccount
string
purchasingContact
string
vendor
string
warrantyExpirationDate
date-time
activationLockEnabled
boolean
blockEncryptionCapable
boolean
dataProtection
boolean
fileEncryptionCapable
boolean
hardwareEncryptionSupported
boolean
jailbreakStatus
string
passcodeCompliant
boolean
passcodeCompliantWithProfile
boolean
passcodeLockGracePeriodEnforcedSeconds
integer
passcodePresent
boolean
personalDeviceProfileCurrent
boolean
Defaults to false
deprecated
Deprecated as of 11.25. This field always returns false.

carrierSettingsVersion
string
cellularTechnology
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
dataRoamingEnabled
boolean
eid
string
homeCarrierNetwork
string
homeMobileCountryCode
string
homeMobileNetworkCode
string
iccid
string
imei
string
imei2
string
meid
string
personalHotspotEnabled
boolean
preferredVoiceNumber
string
roaming
boolean
voiceRoamingEnabled
string
lastLoggedInUsernameSelfService
string | null
lastLoggedInUsernameSelfServiceTimestamp
date-time | null
extensionAttributeValueList
array of objects
object
displayName
string
value
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/static-group-membership/?page=0&page-size=100&sort=displayName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "mobileDeviceId": "1",
      "udid": "0dad565fb40b010a9e490440188063a378721069",
      "airPlayPassword": "1234",
      "appAnalyticsEnabled": false,
      "assetTag": "8675309",
      "availableSpaceMb": 26646,
      "batteryLevel": 100,
      "batteryHealth": "UNKNOWN",
      "bluetoothLowEnergyCapable": false,
      "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
      "capacityMb": 27503,
      "lostModeEnabledDate": "2022-10-17T11:48:56.307Z",
      "declarativeDeviceManagementEnabled": false,
      "deviceId": "1",
      "deviceLocatorServiceEnabled": false,
      "deviceOwnershipType": "institutional",
      "devicePhoneNumber": "555-555-5555",
      "diagnosticAndUsageReportingEnabled": false,
      "displayName": "iPad",
      "doNotDisturbEnabled": false,
      "enrollmentSessionTokenValid": false,
      "exchangeDeviceId": "TH3YE2RI4D234IS2B6U6IGH95D",
      "cloudBackupEnabled": false,
      "osBuild": "15F79",
      "osSupplementalBuildVersion": "22A103310o",
      "osRapidSecurityResponse": "(a)",
      "osVersion": "11.4",
      "ipAddress": "10.0.0.1",
      "itunesStoreAccountActive": false,
      "jamfParentPairings": 1,
      "languages": "Polish",
      "lastBackupDate": "2022-10-17T11:48:56.307Z",
      "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
      "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
      "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
      "locales": "null",
      "locationServicesForSelfServiceMobileEnabled": false,
      "lostModeEnabled": false,
      "managed": true,
      "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
      "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
      "model": "iPad 7th Generation (Wi-Fi)",
      "modelIdentifier": "iPad7,11",
      "modelNumber": "MW742LL",
      "modemFirmwareVersion": "5.70.01",
      "pairedDevices": 1,
      "quotaSize": 1024,
      "residentUsers": 0,
      "serialNumber": "5c28fdae",
      "sharedIpad": false,
      "supervised": true,
      "tethered": false,
      "timeZone": "Europe/Warsaw",
      "usedSpacePercentage": 3,
      "wifiMacAddress": "C4:84:66:92:78:00",
      "building": "Eau Claire",
      "department": "Support",
      "emailAddress": "support@jamf.com",
      "fullName": "John Smith",
      "position": "IT Team Lead",
      "room": "4th Floor - Quad 3",
      "userPhoneNumber": "555-555-5555",
      "username": "admin",
      "appleCareId": "9546567.0",
      "leaseExpirationDate": "2022-10-17T11:48:56.307Z",
      "lifeExpectancyYears": 7,
      "poDate": "2022-10-17T11:48:56.307Z",
      "poNumber": "8675309",
      "purchasePrice": "$399",
      "purchasedOrLeased": true,
      "purchasingAccount": "IT Budget",
      "purchasingContact": "Nick in IT",
      "vendor": "Apple",
      "warrantyExpirationDate": "2022-10-17T11:48:56.307Z",
      "activationLockEnabled": true,
      "blockEncryptionCapable": false,
      "dataProtection": false,
      "fileEncryptionCapable": false,
      "hardwareEncryptionSupported": false,
      "jailbreakStatus": "Compromised",
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "passcodeLockGracePeriodEnforcedSeconds": 9819083,
      "passcodePresent": true,
      "carrierSettingsVersion": "33.1",
      "cellularTechnology": "Both",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "eid": "89049032007008882600085727376656",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "iccid": "8991101200003204514",
      "imei": "35 882334 083223 0",
      "imei2": "35 585308 072899 3",
      "meid": "35882334083223",
      "personalHotspotEnabled": false,
      "preferredVoiceNumber": "555-555-5555",
      "roaming": false,
      "voiceRoamingEnabled": "Enabled",
      "lastLoggedInUsernameSelfService": "admin",
      "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
      "extensionAttributeValueList": [
        {
          "displayName": "EA Name",
          "value": "EA Value"
        }
      ]
    }
  ]
}
-----
Get Static Groups
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups

Get Static Groups

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
Defaults to groupId:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId.


string

groupId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName=="staticGroup1"

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
groupId
string
groupName
string
groupDescription
string
siteId
string
count
integer
≥ 0
membership count

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/static-groups?page=0&page-size=100&sort=groupId%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "groupId": "6",
      "groupName": "Static iPad Group",
      "groupDescription": "A static group containing only iPads",
      "siteId": "11",
      "count": 15
    }
  ]
}
-----
Create membership of a static group
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups

Create membership of a static group

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
groupName
string
required
length ≥ 1
Static iPads
groupDescription
string
Static iPads
siteId
string
11
assignments
array of objects

ADD object
Responses

201
Static mobile device group created successfully

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/static-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupName": "Static iPads",
  "groupDescription": "Static iPads",
  "siteId": "11"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get Static Group by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups/{id}

Get Static Group by Id

Path Params
id
string
required
instance id of static-group

1
Responses

200
Successful response

Response body
object
groupId
string
groupName
string
groupDescription
string
siteId
string
count
integer
≥ 0
membership count

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/static-groups/1 \
     --header 'accept: application/json'

{
  "groupId": "6",
  "groupName": "Static iPad Group",
  "groupDescription": "A static group containing only iPads",
  "siteId": "11",
  "count": 15
}
-----
Remove Static Group by Id
delete
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups/{id}

Remove Static Group by Id

Path Params
id
string
required
instance id of static-group

Response
204
Static Group successfully removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/static-groups/
-----
Update membership of a static group
patch
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/static-groups/{id}

Update membership of a static group

Path Params
id
string
required
instance id of a static group

1
Body Params
groupName
string
required
length ≥ 1
Static iPads
groupDescription
string
Static iPads
siteId
string
11
assignments
array of objects

object

mobileDeviceId
string
5
selected
boolean
If true the device should be added to the group, if false should be removed from the group


true

ADD object
Responses

200
Successful response

Response body
object
groupId
string
length ≥ 1
groupName
string
required
length ≥ 1
groupDescription
string
siteId
string
assignments
array of objects
object
mobileDeviceId
string
selected
boolean
If true the device should be added to the group, if false should be removed from the group

{
  "groupId": "7",
  "groupName": "Static iPads",
  "groupDescription": "Static iPads",
  "siteId": "11",
  "assignments": [
    {
      "mobileDeviceId": "5",
      "selected": true
    }
  ]
}
-----
Erase all devices in the group
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/{id}/erase

Erase all devices in the group

Path Params
id
string
required
instance id of mobile-device-group

Body Params
preserveDataPlan
boolean

true
disallowProximitySetup
boolean

true
clearActivationLock
boolean

true
returnToService
boolean

true
Response
204
Erase all devices in the group

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups//erase \
     --header 'content-type: application/json' \
     --data '
{
  "returnToService": true,
  "clearActivationLock": true,
  "disallowProximitySetup": true,
  "preserveDataPlan": true
}
'
-----
Get Smart Group Membership by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-group-membership/{id}

Get Smart Group Membership by Id

Path Params
id
string
required
instance id of smart-group

1
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
Defaults to displayName:asc
Sorting criteria in the format: property:asc/desc. Default sort is mobileDeviceId:asc. Multiple sort criteria are supported and must be separated with a comma.

Fields allowed in the sort: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, batteryHealth, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, lostModeEnabledDate, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, enrollmentSessionTokenValid, exchangeDeviceId, cloudBackupEnabled, osBuild, osRapidSecurityResponse, osSupplementalBuildVersion, osVersion, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastBackupDate, lastEnrolledDate, lastCloudBackupDate, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, mdmProfileExpirationDate, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, deviceOwnershipType, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, leaseExpirationDate,lifeExpectancyYears, poDate, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, warrantyExpirationDate, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, hardwareEncryptionSupported, jailbreakStatus, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, cellularTechnology, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, voiceRoamingEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be sorted by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

Example: sort=displayName:desc,username:asc


string

displayName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter mobile device collection. Default filter is empty query - returning all results for the requested page.

Fields allowed in the query: airPlayPassword, appAnalyticsEnabled, assetTag, availableSpaceMb, batteryLevel, bluetoothLowEnergyCapable, bluetoothMacAddress, capacityMb, declarativeDeviceManagementEnabled, deviceId, deviceLocatorServiceEnabled, devicePhoneNumber, diagnosticAndUsageReportingEnabled, displayName, doNotDisturbEnabled, exchangeDeviceId, cloudBackupEnabled, osBuild, osSupplementalBuildVersion, osVersion, osRapidSecurityResponse, ipAddress, itunesStoreAccountActive, mobileDeviceId, managementId, languages, lastInventoryUpdateDate, locales, locationServicesForSelfServiceMobileEnabled, lostModeEnabled, managed, model, modelIdentifier, modelNumber, modemFirmwareVersion, preferredVoiceNumber, quotaSize, residentUsers, serialNumber, sharedIpad, supervised, tethered, timeZone, udid, usedSpacePercentage, wifiMacAddress, building, department, emailAddress, fullName, userPhoneNumber, position, room, username, appleCareId, lifeExpectancyYears, poNumber, purchasePrice, purchasedOrLeased, purchasingAccount, purchasingContact, vendor, activationLockEnabled, blockEncryptionCapable, dataProtection, fileEncryptionCapable, passcodeCompliant, passcodeCompliantWithProfile, passcodeLockGracePeriodEnforcedSeconds, passcodePresent, carrierSettingsVersion, currentCarrierNetwork, currentMobileCountryCode, currentMobileNetworkCode, dataRoamingEnabled, eid, network, homeMobileCountryCode, homeMobileNetworkCode, iccid, imei, imei2, meid, personalHotspotEnabled, roaming, lastLoggedInUsernameSelfService, lastLoggedInUsernameSelfServiceTimestamp

Extension attributes can be filtered by using the format EA+ID where ID is the ID of the extension attribute, for example EA+1!=null

This param can be combined with paging and sorting. Example: filter=displayName=="iPad"

Responses

200
Successful response

Response body
object
totalCount
integer
results
array of objects
object
mobileDeviceId
string
length ≥ 1
udid
string
airPlayPassword
password
appAnalyticsEnabled
boolean
assetTag
string
availableSpaceMb
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

bluetoothLowEnergyCapable
boolean
bluetoothMacAddress
string
capacityMb
integer
lostModeEnabledDate
date-time
declarativeDeviceManagementEnabled
boolean
deviceId
string
deviceLocatorServiceEnabled
boolean
deviceOwnershipType
string
devicePhoneNumber
string
diagnosticAndUsageReportingEnabled
boolean
displayName
string
doNotDisturbEnabled
boolean
enrollmentSessionTokenValid
boolean
exchangeDeviceId
string
cloudBackupEnabled
boolean
osBuild
string
osSupplementalBuildVersion
string
osRapidSecurityResponse
string
osVersion
string
ipAddress
string
itunesStoreAccountActive
boolean
jamfParentPairings
integer
languages
string
lastBackupDate
date-time
lastEnrolledDate
date-time
lastCloudBackupDate
date-time
lastInventoryUpdateDate
date-time
locales
string
locationServicesForSelfServiceMobileEnabled
boolean
lostModeEnabled
boolean
managed
boolean
managementId
string
mdmProfileExpirationDate
date-time
model
string
modelIdentifier
string
modelNumber
string
modemFirmwareVersion
string
pairedDevices
integer
quotaSize
integer
residentUsers
integer
serialNumber
string
sharedIpad
boolean
supervised
boolean
tethered
boolean
timeZone
string
usedSpacePercentage
integer
wifiMacAddress
string
building
string
department
string
emailAddress
string
fullName
string
position
string
room
string
userPhoneNumber
string
username
string
appleCareId
string
leaseExpirationDate
date-time
lifeExpectancyYears
integer
poDate
date-time
poNumber
string
purchasePrice
string
purchasedOrLeased
boolean
purchasingAccount
string
purchasingContact
string
vendor
string
warrantyExpirationDate
date-time
activationLockEnabled
boolean
blockEncryptionCapable
boolean
dataProtection
boolean
fileEncryptionCapable
boolean
hardwareEncryptionSupported
boolean
jailbreakStatus
string
passcodeCompliant
boolean
passcodeCompliantWithProfile
boolean
passcodeLockGracePeriodEnforcedSeconds
integer
passcodePresent
boolean
personalDeviceProfileCurrent
boolean
Defaults to false
deprecated
Deprecated as of 11.25. This field always returns false.

carrierSettingsVersion
string
cellularTechnology
string
currentCarrierNetwork
string
currentMobileCountryCode
string
currentMobileNetworkCode
string
dataRoamingEnabled
boolean
eid
string
homeCarrierNetwork
string
homeMobileCountryCode
string
homeMobileNetworkCode
string
iccid
string
imei
string
imei2
string
meid
string
personalHotspotEnabled
boolean
preferredVoiceNumber
string
roaming
boolean
voiceRoamingEnabled
string
lastLoggedInUsernameSelfService
string | null
lastLoggedInUsernameSelfServiceTimestamp
date-time | null
extensionAttributeValueList
array of objects
object
displayName
string
value
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-group-membership/1?page=0&page-size=100&sort=displayName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "mobileDeviceId": "1",
      "udid": "0dad565fb40b010a9e490440188063a378721069",
      "airPlayPassword": "1234",
      "appAnalyticsEnabled": false,
      "assetTag": "8675309",
      "availableSpaceMb": 26646,
      "batteryLevel": 100,
      "batteryHealth": "UNKNOWN",
      "bluetoothLowEnergyCapable": false,
      "bluetoothMacAddress": "ee:00:7c:f0:e5:aa",
      "capacityMb": 27503,
      "lostModeEnabledDate": "2022-10-17T11:48:56.307Z",
      "declarativeDeviceManagementEnabled": false,
      "deviceId": "1",
      "deviceLocatorServiceEnabled": false,
      "deviceOwnershipType": "institutional",
      "devicePhoneNumber": "555-555-5555",
      "diagnosticAndUsageReportingEnabled": false,
      "displayName": "iPad",
      "doNotDisturbEnabled": false,
      "enrollmentSessionTokenValid": false,
      "exchangeDeviceId": "TH3YE2RI4D234IS2B6U6IGH95D",
      "cloudBackupEnabled": false,
      "osBuild": "15F79",
      "osSupplementalBuildVersion": "22A103310o",
      "osRapidSecurityResponse": "(a)",
      "osVersion": "11.4",
      "ipAddress": "10.0.0.1",
      "itunesStoreAccountActive": false,
      "jamfParentPairings": 1,
      "languages": "Polish",
      "lastBackupDate": "2022-10-17T11:48:56.307Z",
      "lastEnrolledDate": "2022-10-17T11:48:56.307Z",
      "lastCloudBackupDate": "2022-10-17T11:48:56.307Z",
      "lastInventoryUpdateDate": "2022-10-17T11:48:56.307Z",
      "locales": "null",
      "locationServicesForSelfServiceMobileEnabled": false,
      "lostModeEnabled": false,
      "managed": true,
      "managementId": "73226fb6-61df-4c10-9552-eb9bc353d507",
      "mdmProfileExpirationDate": "2022-10-17T11:48:56.307Z",
      "model": "iPad 7th Generation (Wi-Fi)",
      "modelIdentifier": "iPad7,11",
      "modelNumber": "MW742LL",
      "modemFirmwareVersion": "5.70.01",
      "pairedDevices": 1,
      "quotaSize": 1024,
      "residentUsers": 0,
      "serialNumber": "5c28fdae",
      "sharedIpad": false,
      "supervised": true,
      "tethered": false,
      "timeZone": "Europe/Warsaw",
      "usedSpacePercentage": 3,
      "wifiMacAddress": "C4:84:66:92:78:00",
      "building": "Eau Claire",
      "department": "Support",
      "emailAddress": "support@jamf.com",
      "fullName": "John Smith",
      "position": "IT Team Lead",
      "room": "4th Floor - Quad 3",
      "userPhoneNumber": "555-555-5555",
      "username": "admin",
      "appleCareId": "9546567.0",
      "leaseExpirationDate": "2022-10-17T11:48:56.307Z",
      "lifeExpectancyYears": 7,
      "poDate": "2022-10-17T11:48:56.307Z",
      "poNumber": "8675309",
      "purchasePrice": "$399",
      "purchasedOrLeased": true,
      "purchasingAccount": "IT Budget",
      "purchasingContact": "Nick in IT",
      "vendor": "Apple",
      "warrantyExpirationDate": "2022-10-17T11:48:56.307Z",
      "activationLockEnabled": true,
      "blockEncryptionCapable": false,
      "dataProtection": false,
      "fileEncryptionCapable": false,
      "hardwareEncryptionSupported": false,
      "jailbreakStatus": "Compromised",
      "passcodeCompliant": true,
      "passcodeCompliantWithProfile": true,
      "passcodeLockGracePeriodEnforcedSeconds": 9819083,
      "passcodePresent": true,
      "carrierSettingsVersion": "33.1",
      "cellularTechnology": "Both",
      "currentCarrierNetwork": "Verizon Wireless",
      "currentMobileCountryCode": "311",
      "currentMobileNetworkCode": "480",
      "dataRoamingEnabled": true,
      "eid": "89049032007008882600085727376656",
      "homeCarrierNetwork": "Verizon",
      "homeMobileCountryCode": "US",
      "homeMobileNetworkCode": "480",
      "iccid": "8991101200003204514",
      "imei": "35 882334 083223 0",
      "imei2": "35 585308 072899 3",
      "meid": "35882334083223",
      "personalHotspotEnabled": false,
      "preferredVoiceNumber": "555-555-5555",
      "roaming": false,
      "voiceRoamingEnabled": "Enabled",
      "lastLoggedInUsernameSelfService": "admin",
      "lastLoggedInUsernameSelfServiceTimestamp": "2018-10-31T18:04:13Z",
      "extensionAttributeValueList": [
        {
          "displayName": "EA Name",
          "value": "EA Value"
        }
      ]
    }
  ]
}
-----
Get Smart Groups
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups

Get Smart Groups

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
Defaults to groupId:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Available criteria to sort on: groupId, groupName, siteId.


string

groupId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter smart group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupId, groupName, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. This param can be combined with paging and sorting. Example: groupName=="smartGroup1"

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
groupId
string
groupName
string
groupDescription
string
siteId
string
count
integer
≥ 0
membership count

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-groups?page=0&page-size=100&sort=groupId%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 2,
  "results": [
    {
      "groupId": "1",
      "groupName": "Smart iPad Group",
      "groupDescription": "A smart group containing iPads meeting certain criteria",
      "siteId": "11",
      "count": 25
    }
  ]
}
-----
Create a smart group
post
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups

Create a smart group

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
groupName
string
required
length ≥ 1
Smart iPads Group
groupDescription
string
Smart iPads that meet certain criteria
siteId
string
11
criteria
array of objects

ADD object
Responses

201
Smart mobile device group created successfully

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupName": "Smart iPads Group",
  "groupDescription": "Smart iPads that meet certain criteria",
  "siteId": "11"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Get Smart Group by Id
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/{id}

Get Smart Group by Id

Path Params
id
string
required
instance id of smart-group

1
Responses

200
Successful response

Response body
object
groupId
string
groupName
string
groupDescription
string
siteId
string
count
integer
≥ 0
membership count

criteria
array of objects
The criteria used to define the smart group

object
name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

priority
integer
required
≥ 0
The priority order of this criterion

andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

value
string
required
The value to search for

openingParen
boolean
Whether to add an opening parenthesis before this criterion

closingParen
boolean
Whether to add a closing parenthesis after this criterion

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/1 \
     --header 'accept: application/json'

{
  "groupId": "1",
  "groupName": "Smart iPad Group",
  "groupDescription": "A smart group containing iPads meeting certain criteria",
  "siteId": "11",
  "count": 25,
  "criteria": [
    {
      "name": "Model",
      "priority": 1,
      "andOr": "and",
      "searchType": "is",
      "value": "iPad",
      "openingParen": false,
      "closingParen": false
    }
  ]
}
-----
Update a smart group
put
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/{id}

Update a smart group

Path Params
id
string
required
instance id of a smart group

Body Params
groupName
string
required
length ≥ 1
Smart iPads Group
groupDescription
string
Smart iPads that meet certain criteria
siteId
string
11
criteria
array of objects

object

name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

Model
priority
integer
required
≥ 0
The priority order of this criterion

1
andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

and
searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

is
value
string
required
The value to search for

iPad
openingParen
boolean
Whether to add an opening parenthesis before this criterion


true
closingParen
boolean
Whether to add a closing parenthesis after this criterion


false

object

name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

Model
priority
integer
required
≥ 0
The priority order of this criterion

1
andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

and
searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

is
value
string
required
The value to search for

iPad
openingParen
boolean
Whether to add an opening parenthesis before this criterion


true
closingParen
boolean
Whether to add a closing parenthesis after this criterion


true

ADD object
Responses

202
Successful response

Response body
object
groupId
string
length ≥ 1
The unique identifier of the smart group

groupName
string
required
length ≥ 1
groupDescription
string
siteId
string
criteria
array of objects
object
name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

priority
integer
required
≥ 0
The priority order of this criterion

andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

value
string
required
The value to search for

openingParen
boolean
Whether to add an opening parenthesis before this criterion

closingParen
boolean
Whether to add a closing parenthesis after this criterion

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "siteId": "11",
  "groupDescription": "Smart iPads that meet certain criteria",
  "groupName": "Smart iPads Group",
  "criteria": [
    {
      "name": "Model",
      "priority": 1,
      "andOr": "and",
      "searchType": "is",
      "value": "iPad",
      "openingParen": true,
      "closingParen": false
    },
    {
      "name": "Model",
      "priority": 1,
      "andOr": "and",
      "searchType": "is",
      "value": "iPad",
      "openingParen": true,
      "closingParen": true
    }
  ]
}
'
{
  "groupId": "7",
  "groupName": "Smart iPads Group",
  "groupDescription": "Smart iPads that meet certain criteria",
  "siteId": "11",
  "criteria": [
    {
      "name": "Model",
      "priority": 1,
      "andOr": "and",
      "searchType": "is",
      "value": "iPad",
      "openingParen": false,
      "closingParen": false
    }
  ]
}
-----
Remove Smart Group by Id
delete
https://yourServer.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/{id}

Remove Smart Group by Id

Path Params
id
string
required
instance id of smart-group

1
Response
204
Smart Group successfully removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-groups/smart-groups/1
-----