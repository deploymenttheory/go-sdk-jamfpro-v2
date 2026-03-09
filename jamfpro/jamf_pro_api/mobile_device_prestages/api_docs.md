Get all Prestage sync States for all prestages
get
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/syncs

Get all prestage sync states for all prestages

Response

200
Successful response

Response body
array of objects
object
syncState
string
prestageId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages/syncs \
     --header 'accept: application/json'

[
  {
    "syncState": "CONNECTION_ERROR",
    "prestageId": "1",
    "timestamp": "2019-04-17T14:08:06.706+0000"
  }
]
-----
Get Device Scope for a specific Mobile Device Prestage
get
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/scope

Get device scope for a specific mobile device prestage

Path Params
id
string
required
Mobile Device Prestage identifier

Responses

200
Successful response

Response body
object
prestageId
string
assignments
array of objects
object
serialNumber
string
assignmentDate
date-time
userAssigned
string
versionLock
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages//scope \
     --header 'accept: application/json'

{
  "prestageId": "1",
  "assignments": [
    {
      "serialNumber": "XYZ",
      "assignmentDate": "2019-02-04T21:09:31.661Z",
      "userAssigned": "admin"
    }
  ],
  "versionLock": 1
}
-----
Replace Device Scope for a specific Mobile Device Prestage
put
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/scope

Replace device scope for a specific mobile device prestage

Path Params
id
string
required
Mobile Device Prestage identifier

Body Params
Serial Numbers to scope

serialNumbers
array of strings
required

string

DMQVGC0DHLF0

string

DMQVGC0DHLF0

string


ADD string
versionLock
integer
required
1
Responses

200
Successful response

Response body
object
prestageId
string
assignments
array of objects
object
serialNumber
string
assignmentDate
date-time
userAssigned
string
versionLock
integer

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages//scope \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "serialNumbers": [
    "DMQVGC0DHLF0",
    "DMQVGC0DHLF0"
  ],
  "versionLock": 1
}
'

{
  "prestageId": "1",
  "assignments": [
    {
      "serialNumber": "XYZ",
      "assignmentDate": "2019-02-04T21:09:31.661Z",
      "userAssigned": "admin"
    }
  ],
  "versionLock": 1
}
-----
Add Device Scope for a specific Mobile Device Prestage
post
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/scope

Add device scope for a specific mobile device prestage

Path Params
id
string
required
Mobile Device Prestage identifier

Body Params
Serial Numbers to scope

serialNumbers
array of strings
required

string

DMQVGC0DHLF0

ADD string
versionLock
integer
required
1
Responses

200
Successful response

Response body
object
prestageId
string
assignments
array of objects
object
serialNumber
string
assignmentDate
date-time
userAssigned
string
versionLock
integer

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages//scope \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "serialNumbers": [
    "DMQVGC0DHLF0"
  ],
  "versionLock": 1
}
'

{
  "prestageId": "1",
  "assignments": [
    {
      "serialNumber": "XYZ",
      "assignmentDate": "2019-02-04T21:09:31.661Z",
      "userAssigned": "admin"
    }
  ],
  "versionLock": 1
}
-----
Remove Device Scope for a specific Mobile Device Prestage
post
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/scope/delete-multiple

Remove device scope for a specific mobile device prestage

Path Params
id
string
required
Mobile Device Prestage identifier

1
Body Params
Serial Numbers to remove from scope

serialNumbers
array of strings
required

string

DMQVGC0DHLF0

ADD string
versionLock
integer
required
1
Responses

200
Successful response

Response body
object
prestageId
string
assignments
array of objects
object
serialNumber
string
assignmentDate
date-time
userAssigned
string
versionLock
integer

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages/1/scope/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "serialNumbers": [
    "DMQVGC0DHLF0"
  ],
  "versionLock": 1
}
'

{
  "prestageId": "1",
  "assignments": [
    {
      "serialNumber": "XYZ",
      "assignmentDate": "2019-02-04T21:09:31.661Z",
      "userAssigned": "admin"
    }
  ],
  "versionLock": 1
}
-----
Get all prestage sync states for a single prestage
get
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/syncs

Get all prestage sync states for a single prestage

Path Params
id
string
required
Mobile Device Prestage identifier

1
Response

200
Successful response

Response body
array of objects
object
syncState
string
prestageId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages/1/syncs \
     --header 'accept: application/json'

[
  {
    "syncState": "CONNECTION_ERROR",
    "prestageId": "1",
    "timestamp": "2019-04-17T14:08:06.706+0000"
  }
]
-----
Get the latest Sync State for a single Prestage
get
https://yourServer.jamfcloud.com/api/v2/mobile-device-prestages/{id}/syncs/latest

Get the latest sync state for a single prestage

Path Params
id
string
required
Mobile Device Prestage identifier

1
Response

200
Successful response

Response body
object
syncState
string
prestageId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/mobile-device-prestages/1/syncs/latest \
     --header 'accept: application/json'

{
  "syncState": "CONNECTION_ERROR",
  "prestageId": "1",
  "timestamp": "2019-04-17T14:08:06.706+0000"
}
-----

Create a Mobile Device Prestage
post
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages

Create a mobile device prestage

Body Params
Mobile Device Prestage to create. ids defined in this body will be ignored

displayName
string
required
Example Mobile Prestage Name
mandatory
boolean
required

true
mdmRemovable
boolean
required

true
supportPhoneNumber
string
required
5555555555
supportEmailAddress
string
required
example@example.com
department
string
required
Oxbow
defaultPrestage
boolean
required

true
enrollmentSiteId
string
required
-1
keepExistingSiteMembership
boolean
required

true
keepExistingLocationInformation
boolean
required

true
requireAuthentication
boolean
required

true
authenticationPrompt
string
required
LDAP authentication prompt
preventActivationLock
boolean
required

true
enableDeviceBasedActivationLock
boolean
required

true
deviceEnrollmentProgramInstanceId
string
required
5
skipSetupItems
object

skipSetupItems object
boolean
newKey

false
boolean
newKey-1

false

Add Field
locationInformation
object
required

locationInformation object
username
string
required
name
realname
string
required
realName
phone
string
required
123-456-7890
email
string
required
test@jamf.com
room
string
required
room
position
string
required
postion
departmentId
string
required
1
buildingId
string
required
1
id
string
required
-1
versionLock
integer
required
1
purchasingInformation
object
required

purchasingInformation object
id
string
required
-1
leased
boolean
required

true
purchased
boolean
required

true
appleCareId
string
required
abcd
poNumber
string
required
53-1
vendor
string
required
Example Vendor
purchasePrice
string
required
$500
lifeExpectancy
integer
required
5
purchasingAccount
string
required
admin
purchasingContact
string
required
true
leaseDate
string
required
2019-01-01
poDate
string
required
2019-01-01
warrantyDate
string
required
2019-01-01
versionLock
integer
required
1
anchorCertificates
array of strings
The Base64 encoded PEM Certificate


string

xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==

ADD string
enrollmentCustomizationId
string
2
language
string
en
region
string
US
autoAdvanceSetup
boolean
required

true
allowPairing
boolean
required

true
multiUser
boolean
required

true
supervised
boolean
required

true
maximumSharedAccounts
integer
required
10
configureDeviceBeforeSetupAssistant
boolean
required

true
names
object

names object
sendTimezone
boolean
required

true
timezone
string
required
America/Chicago
storageQuotaSizeMegabytes
integer
required
4096
useStorageQuotaSize
boolean
required

true
temporarySessionOnly
boolean

true
enforceTemporarySessionTimeout
boolean

false
temporarySessionTimeout
integer
30
enforceUserSessionTimeout
boolean

false
userSessionTimeout
integer
30
prestageMinimumOsTargetVersionTypeIos
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersionIos
string
length ≥ 0
17.1
prestageMinimumOsTargetVersionTypeIpad
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersionIpad
string
length ≥ 0
17.1
rtsEnabled
boolean

false
rtsConfigProfileId
string
1
preserveManagedApps
boolean
Controls whether managed apps are preserved during Return to Service operations.


true
installAppsDuringEnrollment
boolean
Controls whether apps are installed during the enrollment process.


true
doNotUseProfileFromBackup
boolean
If true, the device does not use the profile when it restores a backup. Default is false. Available in iOS 26 and later, and visionOS 26 and later; otherwise ignored by devices.


true
Response

201
Mobile Device Prestage was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "mandatory": true,
  "mdmRemovable": true,
  "defaultPrestage": true,
  "keepExistingSiteMembership": true,
  "keepExistingLocationInformation": true,
  "requireAuthentication": true,
  "preventActivationLock": true,
  "enableDeviceBasedActivationLock": true,
  "skipSetupItems": {
    "newKey": false,
    "newKey-1": false
  },
  "locationInformation": {
    "username": "name",
    "realname": "realName",
    "phone": "123-456-7890",
    "email": "test@jamf.com",
    "room": "room",
    "position": "postion",
    "departmentId": "1",
    "buildingId": "1",
    "id": "-1",
    "versionLock": 1
  },
  "purchasingInformation": {
    "leased": true,
    "purchased": true,
    "id": "-1",
    "appleCareId": "abcd",
    "poNumber": "53-1",
    "vendor": "Example Vendor",
    "purchasePrice": "$500",
    "lifeExpectancy": 5,
    "purchasingAccount": "admin",
    "purchasingContact": "true",
    "leaseDate": "2019-01-01",
    "poDate": "2019-01-01",
    "warrantyDate": "2019-01-01",
    "versionLock": 1
  },
  "autoAdvanceSetup": true,
  "allowPairing": true,
  "multiUser": true,
  "supervised": true,
  "configureDeviceBeforeSetupAssistant": true,
  "sendTimezone": true,
  "useStorageQuotaSize": true,
  "displayName": "Example Mobile Prestage Name",
  "supportPhoneNumber": "5555555555",
  "supportEmailAddress": "example@example.com",
  "department": "Oxbow",
  "enrollmentSiteId": "-1",
  "authenticationPrompt": "LDAP authentication prompt",
  "deviceEnrollmentProgramInstanceId": "5",
  "anchorCertificates": [
    "xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
  ],
  "enrollmentCustomizationId": "2",
  "language": "en",
  "region": "US",
  "maximumSharedAccounts": 10,
  "timezone": "America/Chicago",
  "storageQuotaSizeMegabytes": 4096,
  "temporarySessionOnly": true,
  "enforceTemporarySessionTimeout": false,
  "temporarySessionTimeout": 30,
  "enforceUserSessionTimeout": false,
  "userSessionTimeout": 30,
  "prestageMinimumOsTargetVersionTypeIos": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIos": "17.1",
  "prestageMinimumOsTargetVersionTypeIpad": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIpad": "17.1",
  "rtsEnabled": false,
  "rtsConfigProfileId": "1",
  "preserveManagedApps": true,
  "installAppsDuringEnrollment": true,
  "doNotUseProfileFromBackup": true
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a Mobile Device Prestage with the supplied id
get
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}

Retrieves a Mobile Device Prestage with the supplied id

Path Params
id
string
required
Mobile Device Prestage identifier

1
Responses

200
Success

Response body
object
displayName
string
required
mandatory
boolean
required
mdmRemovable
boolean
required
supportPhoneNumber
string
required
supportEmailAddress
string
required
department
string
required
defaultPrestage
boolean
required
enrollmentSiteId
string
required
keepExistingSiteMembership
boolean
required
keepExistingLocationInformation
boolean
required
requireAuthentication
boolean
required
authenticationPrompt
string
required
preventActivationLock
boolean
required
enableDeviceBasedActivationLock
boolean
required
deviceEnrollmentProgramInstanceId
string
required
skipSetupItems
object
Has additional fields
locationInformation
object
required
username
string
required
realname
string
required
phone
string
required
email
string
required
room
string
required
position
string
required
departmentId
string
required
buildingId
string
required
id
string
required
versionLock
integer
required
purchasingInformation
object
required
id
string
required
leased
boolean
required
purchased
boolean
required
appleCareId
string
required
poNumber
string
required
vendor
string
required
purchasePrice
string
required
lifeExpectancy
integer
required
purchasingAccount
string
required
purchasingContact
string
required
leaseDate
string
required
poDate
string
required
warrantyDate
string
required
versionLock
integer
required
anchorCertificates
array of strings
The Base64 encoded PEM Certificate

enrollmentCustomizationId
string
language
string
region
string
autoAdvanceSetup
boolean
required
allowPairing
boolean
required
multiUser
boolean
required
supervised
boolean
required
maximumSharedAccounts
integer
required
configureDeviceBeforeSetupAssistant
boolean
required
names
object
assignNamesUsing
string
prestageDeviceNames
array of objects
object
id
string
deviceName
string
used
boolean
deviceNamePrefix
string
deviceNameSuffix
string
singleDeviceName
string
manageNames
boolean
deviceNamingConfigured
boolean
sendTimezone
boolean
required
timezone
string
required
storageQuotaSizeMegabytes
integer
required
useStorageQuotaSize
boolean
required
temporarySessionOnly
boolean
enforceTemporarySessionTimeout
boolean
temporarySessionTimeout
integer
enforceUserSessionTimeout
boolean
userSessionTimeout
integer
prestageMinimumOsTargetVersionTypeIos
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersionIos
string
length ≥ 0
prestageMinimumOsTargetVersionTypeIpad
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersionIpad
string
length ≥ 0
rtsEnabled
boolean
rtsConfigProfileId
string
preserveManagedApps
boolean
Controls whether managed apps are preserved during Return to Service operations.

installAppsDuringEnrollment
boolean
Controls whether apps are installed during the enrollment process.

doNotUseProfileFromBackup
boolean
If true, the device does not use the profile when it restores a backup. Default is false. Available in iOS 26 and later, and visionOS 26 and later; otherwise ignored by devices.

id
string
profileUuid
string
siteId
string
versionLock
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages/1 \
     --header 'accept: application/json'

{
  "displayName": "Example Mobile Prestage Name",
  "mandatory": false,
  "mdmRemovable": true,
  "supportPhoneNumber": "5555555555",
  "supportEmailAddress": "example@example.com",
  "department": "Oxbow",
  "defaultPrestage": false,
  "enrollmentSiteId": "-1",
  "keepExistingSiteMembership": true,
  "keepExistingLocationInformation": true,
  "requireAuthentication": true,
  "authenticationPrompt": "LDAP authentication prompt",
  "preventActivationLock": true,
  "enableDeviceBasedActivationLock": true,
  "deviceEnrollmentProgramInstanceId": "5",
  "skipSetupItems": {
    "Location": true,
    "Privacy": false
  },
  "locationInformation": {
    "username": "name",
    "realname": "realName",
    "phone": "123-456-7890",
    "email": "test@jamf.com",
    "room": "room",
    "position": "postion",
    "departmentId": "1",
    "buildingId": "1",
    "id": "-1",
    "versionLock": 1
  },
  "purchasingInformation": {
    "id": "-1",
    "leased": true,
    "purchased": true,
    "appleCareId": "abcd",
    "poNumber": "53-1",
    "vendor": "Example Vendor",
    "purchasePrice": "$500",
    "lifeExpectancy": 5,
    "purchasingAccount": "admin",
    "purchasingContact": "true",
    "leaseDate": "2019-01-01",
    "poDate": "2019-01-01",
    "warrantyDate": "2019-01-01",
    "versionLock": 1
  },
  "anchorCertificates": [
    "xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
  ],
  "enrollmentCustomizationId": "2",
  "language": "en",
  "region": "US",
  "autoAdvanceSetup": true,
  "allowPairing": true,
  "multiUser": true,
  "supervised": true,
  "maximumSharedAccounts": 10,
  "configureDeviceBeforeSetupAssistant": true,
  "names": {
    "assignNamesUsing": "List of Names",
    "prestageDeviceNames": [
      {
        "id": "1",
        "deviceName": "iPad",
        "used": false
      }
    ],
    "deviceNamePrefix": "prefix",
    "deviceNameSuffix": "suffix",
    "singleDeviceName": "name",
    "manageNames": true,
    "deviceNamingConfigured": true
  },
  "sendTimezone": true,
  "timezone": "America/Chicago",
  "storageQuotaSizeMegabytes": 4096,
  "useStorageQuotaSize": true,
  "temporarySessionOnly": false,
  "enforceTemporarySessionTimeout": false,
  "temporarySessionTimeout": 30,
  "enforceUserSessionTimeout": false,
  "userSessionTimeout": 30,
  "prestageMinimumOsTargetVersionTypeIos": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIos": "17.1",
  "prestageMinimumOsTargetVersionTypeIpad": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIpad": "17.1",
  "rtsEnabled": false,
  "rtsConfigProfileId": "1",
  "preserveManagedApps": false,
  "installAppsDuringEnrollment": true,
  "doNotUseProfileFromBackup": true,
  "id": "1",
  "profileUuid": "29d-a8d8f-b8sdjndf-dsa9",
  "siteId": "5",
  "versionLock": 0
}
-----
Update a Mobile Device Prestage
put
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}

Updates a Mobile Device Prestage

Path Params
id
string
required
Mobile Device Prestage identifier

Body Params
Mobile Device Prestage to update

displayName
string
required
Example Mobile Prestage Name
mandatory
boolean
required

true
mdmRemovable
boolean
required

true
supportPhoneNumber
string
required
5555555555
supportEmailAddress
string
required
example@example.com
department
string
required
Oxbow
defaultPrestage
boolean
required

true
enrollmentSiteId
string
required
keepExistingSiteMembership
boolean
required

true
keepExistingLocationInformation
boolean
required

true
requireAuthentication
boolean
required

true
authenticationPrompt
string
required
preventActivationLock
boolean
required

true
enableDeviceBasedActivationLock
boolean
required

true
deviceEnrollmentProgramInstanceId
string
required
skipSetupItems
object

skipSetupItems object
locationInformation
object
required

locationInformation object
purchasingInformation
object
required

purchasingInformation object
anchorCertificates
array of strings
The Base64 encoded PEM Certificate


ADD string
enrollmentCustomizationId
string
language
string
region
string
autoAdvanceSetup
boolean
required

true
allowPairing
boolean
required

true
multiUser
boolean
required

true
supervised
boolean
required

true
maximumSharedAccounts
integer
required
configureDeviceBeforeSetupAssistant
boolean
required

true
names
object

names object
sendTimezone
boolean
required

true
timezone
string
required
storageQuotaSizeMegabytes
integer
required
useStorageQuotaSize
boolean
required

true
temporarySessionOnly
boolean

false
enforceTemporarySessionTimeout
boolean

false
temporarySessionTimeout
integer
enforceUserSessionTimeout
boolean

false
userSessionTimeout
integer
prestageMinimumOsTargetVersionTypeIos
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersionIos
string
length ≥ 0
prestageMinimumOsTargetVersionTypeIpad
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersionIpad
string
length ≥ 0
rtsEnabled
boolean

false
rtsConfigProfileId
string
preserveManagedApps
boolean
Controls whether managed apps are preserved during Return to Service operations.


false
installAppsDuringEnrollment
boolean
Controls whether apps are installed during the enrollment process.


true
doNotUseProfileFromBackup
boolean
If true, the device does not use the profile when it restores a backup. Default is false. Available in iOS 26 and later, and visionOS 26 and later; otherwise ignored by devices.


true
versionLock
integer
Responses

200
Success

Response body
object
displayName
string
required
mandatory
boolean
required
mdmRemovable
boolean
required
supportPhoneNumber
string
required
supportEmailAddress
string
required
department
string
required
defaultPrestage
boolean
required
enrollmentSiteId
string
required
keepExistingSiteMembership
boolean
required
keepExistingLocationInformation
boolean
required
requireAuthentication
boolean
required
authenticationPrompt
string
required
preventActivationLock
boolean
required
enableDeviceBasedActivationLock
boolean
required
deviceEnrollmentProgramInstanceId
string
required
skipSetupItems
object
Has additional fields
locationInformation
object
required
username
string
required
realname
string
required
phone
string
required
email
string
required
room
string
required
position
string
required
departmentId
string
required
buildingId
string
required
id
string
required
versionLock
integer
required
purchasingInformation
object
required
id
string
required
leased
boolean
required
purchased
boolean
required
appleCareId
string
required
poNumber
string
required
vendor
string
required
purchasePrice
string
required
lifeExpectancy
integer
required
purchasingAccount
string
required
purchasingContact
string
required
leaseDate
string
required
poDate
string
required
warrantyDate
string
required
versionLock
integer
required
anchorCertificates
array of strings
The Base64 encoded PEM Certificate

enrollmentCustomizationId
string
language
string
region
string
autoAdvanceSetup
boolean
required
allowPairing
boolean
required
multiUser
boolean
required
supervised
boolean
required
maximumSharedAccounts
integer
required
configureDeviceBeforeSetupAssistant
boolean
required
names
object
assignNamesUsing
string
prestageDeviceNames
array of objects
object
id
string
deviceName
string
used
boolean
deviceNamePrefix
string
deviceNameSuffix
string
singleDeviceName
string
manageNames
boolean
deviceNamingConfigured
boolean
sendTimezone
boolean
required
timezone
string
required
storageQuotaSizeMegabytes
integer
required
useStorageQuotaSize
boolean
required
temporarySessionOnly
boolean
enforceTemporarySessionTimeout
boolean
temporarySessionTimeout
integer
enforceUserSessionTimeout
boolean
userSessionTimeout
integer
prestageMinimumOsTargetVersionTypeIos
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersionIos
string
length ≥ 0
prestageMinimumOsTargetVersionTypeIpad
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersionIpad
string
length ≥ 0
rtsEnabled
boolean
rtsConfigProfileId
string
preserveManagedApps
boolean
Controls whether managed apps are preserved during Return to Service operations.

installAppsDuringEnrollment
boolean
Controls whether apps are installed during the enrollment process.

doNotUseProfileFromBackup
boolean
If true, the device does not use the profile when it restores a backup. Default is false. Available in iOS 26 and later, and visionOS 26 and later; otherwise ignored by devices.

id
string
profileUuid
string
siteId
string
versionLock
integer

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "mandatory": true,
  "mdmRemovable": true,
  "defaultPrestage": true,
  "keepExistingSiteMembership": true,
  "keepExistingLocationInformation": true,
  "requireAuthentication": true,
  "preventActivationLock": true,
  "enableDeviceBasedActivationLock": true,
  "purchasingInformation": {
    "leased": true,
    "purchased": true
  },
  "autoAdvanceSetup": true,
  "allowPairing": true,
  "multiUser": true,
  "supervised": true,
  "configureDeviceBeforeSetupAssistant": true,
  "sendTimezone": true,
  "useStorageQuotaSize": true,
  "displayName": "Example Mobile Prestage Name",
  "supportPhoneNumber": "5555555555",
  "supportEmailAddress": "example@example.com",
  "department": "Oxbow"
}
'
{
  "displayName": "Example Mobile Prestage Name",
  "mandatory": false,
  "mdmRemovable": true,
  "supportPhoneNumber": "5555555555",
  "supportEmailAddress": "example@example.com",
  "department": "Oxbow",
  "defaultPrestage": false,
  "enrollmentSiteId": "-1",
  "keepExistingSiteMembership": true,
  "keepExistingLocationInformation": true,
  "requireAuthentication": true,
  "authenticationPrompt": "LDAP authentication prompt",
  "preventActivationLock": true,
  "enableDeviceBasedActivationLock": true,
  "deviceEnrollmentProgramInstanceId": "5",
  "skipSetupItems": {
    "Location": true,
    "Privacy": false
  },
  "locationInformation": {
    "username": "name",
    "realname": "realName",
    "phone": "123-456-7890",
    "email": "test@jamf.com",
    "room": "room",
    "position": "postion",
    "departmentId": "1",
    "buildingId": "1",
    "id": "-1",
    "versionLock": 1
  },
  "purchasingInformation": {
    "id": "-1",
    "leased": true,
    "purchased": true,
    "appleCareId": "abcd",
    "poNumber": "53-1",
    "vendor": "Example Vendor",
    "purchasePrice": "$500",
    "lifeExpectancy": 5,
    "purchasingAccount": "admin",
    "purchasingContact": "true",
    "leaseDate": "2019-01-01",
    "poDate": "2019-01-01",
    "warrantyDate": "2019-01-01",
    "versionLock": 1
  },
  "anchorCertificates": [
    "xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
  ],
  "enrollmentCustomizationId": "2",
  "language": "en",
  "region": "US",
  "autoAdvanceSetup": true,
  "allowPairing": true,
  "multiUser": true,
  "supervised": true,
  "maximumSharedAccounts": 10,
  "configureDeviceBeforeSetupAssistant": true,
  "names": {
    "assignNamesUsing": "List of Names",
    "prestageDeviceNames": [
      {
        "id": "1",
        "deviceName": "iPad",
        "used": false
      }
    ],
    "deviceNamePrefix": "prefix",
    "deviceNameSuffix": "suffix",
    "singleDeviceName": "name",
    "manageNames": true,
    "deviceNamingConfigured": true
  },
  "sendTimezone": true,
  "timezone": "America/Chicago",
  "storageQuotaSizeMegabytes": 4096,
  "useStorageQuotaSize": true,
  "temporarySessionOnly": false,
  "enforceTemporarySessionTimeout": false,
  "temporarySessionTimeout": 30,
  "enforceUserSessionTimeout": false,
  "userSessionTimeout": 30,
  "prestageMinimumOsTargetVersionTypeIos": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIos": "17.1",
  "prestageMinimumOsTargetVersionTypeIpad": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersionIpad": "17.1",
  "rtsEnabled": false,
  "rtsConfigProfileId": "1",
  "preserveManagedApps": false,
  "installAppsDuringEnrollment": true,
  "doNotUseProfileFromBackup": true,
  "id": "1",
  "profileUuid": "29d-a8d8f-b8sdjndf-dsa9",
  "siteId": "5",
  "versionLock": 0
}
-----

Delete a Mobile Device Prestage with the supplied id
delete
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}

Deletes a Mobile Device Prestage with the supplied id

Path Params
id
string
required
Mobile Device Prestage identifier

1

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages/1
-----
Get attachments for a Mobile Device Prestage
get
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/attachments

Get attachments for a Mobile Device Prestage

Path Params
id
string
required
Mobile Device Prestage identifier

Responses

200
Success

Response body
array of objects
object
id
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages//attachments \
     --header 'accept: application/json'

[
  {
    "id": "1",
    "name": "receipt.pdf"
  }
]
-----
Add an attachment to a Mobile Device Prestage
post
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/attachments

Add an attachment to a Mobile Device prestage

Path Params
id
string
required
Identifier of the Mobile Device Prestage the attachment should be assigned to

Body Params
file
file
required
The file to upload

No file chosen
Responses

201
Success

Response body
object
id
string
name
string
fileType
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages//attachments \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "id": "1",
  "name": "receipt.pdf",
  "fileType": "pdf"
}
-----
Remove an attachment for a Mobile Device Prestage
post
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/attachments/delete-multiple

Remove an attachment for a Mobile Device Prestage

Path Params
id
string
required
Mobile Device Prestage identifier

1
Body Params
ids
array of strings

string

1

ADD string
Response

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages/1/attachments/delete-multiple \
     --header 'content-type: application/json' \
     --data '{"ids":["1"]}'
-----

Get sorted and paged Mobile Device Prestage history objects
get
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/history

Gets sorted and paged mobile device prestage history objects

Path Params
id
string
required
Mobile Device Prestage identifier

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
Defaults to date:desc
Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc


string

date:desc

ADD string
Response

200
Details of mobile device prestage history were found

Response body
object
totalCount
integer
≥ 0
results
array of objects
length ≥ 0
object
id
integer
≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages/1/history?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": 1,
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}
-----
Add Mobile Device Prestage history object notes
post
https://yourServer.jamfcloud.com/api/v3/mobile-device-prestages/{id}/history

Adds mobile device prestage history object notes

Path Params
id
string
required
Mobile Device Prestage identifier

Body Params
History notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of mobile deivce prestage history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/mobile-device-prestages//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----