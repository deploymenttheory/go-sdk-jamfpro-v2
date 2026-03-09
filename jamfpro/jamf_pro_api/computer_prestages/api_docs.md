Get all device Scope for all Computer Prestages
get
https://yourServer.jamfcloud.com/api/v2/computer-prestages/scope


Get all device scope for all computer prestages

Response

200
Successful response

Response body
object
serialsByPrestageId
object
Has additional fields

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-prestages/scope \
     --header 'accept: application/json'

{
  "serialsByPrestageId": {
    "ABCD": 1,
    "XYZ": 12
  }
}
-----

Get device Scope for a specific Computer Prestage
get
https://yourServer.jamfcloud.com/api/v2/computer-prestages/{id}/scope


Get device scope for a specific computer prestage

Path Params
id
string
required
Computer Prestage identifier

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
     --url https://yourserver.jamfcloud.com/api/v2/computer-prestages//scope \
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

Replace device Scope for a specific Computer Prestage
put
https://yourServer.jamfcloud.com/api/v2/computer-prestages/{id}/scope


Replace device scope for a specific computer prestage

Path Params
id
string
required
Computer Prestage identifier

Body Params
Serial Numbers to scope

serialNumbers
array of strings
required

ADD string
versionLock
integer
required
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
     --url https://yourserver.jamfcloud.com/api/v2/computer-prestages//scope \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

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
Add device Scope for a specific Computer Prestage
post
https://yourServer.jamfcloud.com/api/v2/computer-prestages/{id}/scope


Add device scope for a specific computer prestage

Path Params
id
string
required
Computer Prestage identifier

Body Params
Serial Numbers to scope

serialNumbers
array of strings
required

ADD string
versionLock
integer
required
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
     --url https://yourserver.jamfcloud.com/api/v2/computer-prestages//scope \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

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
Remove device Scope for a specific Computer Prestage
post
https://yourServer.jamfcloud.com/api/v2/computer-prestages/{id}/scope/delete-multiple


Remove device scope for a specific computer prestage

Path Params
id
string
required
Computer Prestage identifier

Body Params
Serial Numbers to remove from scope

serialNumbers
array of strings
required

ADD string
versionLock
integer
required
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
     --url https://yourserver.jamfcloud.com/api/v2/computer-prestages//scope/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

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
Get sorted and paged Computer Prestages
get
https://yourServer.jamfcloud.com/api/v3/computer-prestages

Gets sorted and paged computer prestages

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
Defaults to id:desc
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

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
length ≥ 0
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

locationInformation object
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

purchasingInformation object
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
installProfilesDuringSetup
boolean
required
prestageInstalledProfileIds
array of strings
required
customPackageIds
array of strings
required
customPackageDistributionPointId
string
required
enableRecoveryLock
boolean
recoveryLockPasswordType
string
enum
MANUAL RANDOM

rotateRecoveryLockPassword
boolean
prestageMinimumOsTargetVersionType
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersion
string
length ≥ 0
pssoEnabled
boolean
Defaults to false
Indicates whether Platform SSO (PSSO) is enabled for this computer prestage, regardless of unattended or 403 workflows. When enabled, the PSSO application will be deployed to devices during the setup process to facilitate single sign-on (SSO) for users.

platformSsoAppBundleId
string
The bundle identifier for the Platform SSO (PSSO) application unattended workflow. This identifier is used to specify which PSSO app should be deployed to devices during the setup process.

profileUrl
string | null
The URL to the configuration profile for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process. Users should use either profileUrl or populate pssoConfigProfileId, but not both.

pssoConfigProfileId
string | null
The identifier for the configuration profile associated with the Platform SSO (PSSO) application 403 workflow. This ID is used to specify which configuration profile should be applied to devices during the setup process when PSSO is enabled. Users should use either pssoConfigProfileId or populate profileUrl, but not both.

manifestUrl
string | null
The URL to the manifest file for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process.

authUrl
string | null
The URL to the identity provider (IdP) for authentication for the Platform SSO (PSSO) application 403 workflow. This URL is used in conjunction with PSSO to facilitate single sign-on for users during the device setup process.

id
string
length ≥ 1
profileUuid
string
siteId
string
versionLock
integer
≥ 0
accountSettings
object

accountSettings object
id
string
id of Account Settings

payloadConfigured
boolean
Defaults to false
localAdminAccountEnabled
boolean
Defaults to false
adminUsername
string
length ≥ 0
hiddenAdminAccount
boolean
Defaults to false
localUserManaged
boolean
Defaults to false
userAccountType
string
enum
Defaults to STANDARD
ADMINISTRATOR STANDARD SKIP

versionLock
integer
≥ 0
Defaults to 0
prefillPrimaryAccountInfoFeatureEnabled
boolean
Defaults to false
prefillType
string
Defaults to CUSTOM
Values accepted are only CUSTOM and DEVICE_OWNER

prefillAccountFullName
string
length ≥ 0
prefillAccountUserName
string
length ≥ 0
preventPrefillInfoFromModification
boolean
Defaults to false

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/computer-prestages?page=0&page-size=100&sort=id%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
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
      "installProfilesDuringSetup": true,
      "prestageInstalledProfileIds": [
        "1"
      ],
      "customPackageIds": [
        "1"
      ],
      "customPackageDistributionPointId": "1",
      "enableRecoveryLock": true,
      "recoveryLockPasswordType": "MANUAL",
      "rotateRecoveryLockPassword": true,
      "prestageMinimumOsTargetVersionType": "MINIMUM_OS_LATEST_VERSION",
      "minimumOsSpecificVersion": "17.1",
      "pssoEnabled": true,
      "platformSsoAppBundleId": "com.okta.mobile",
      "profileUrl": "https://mdmserver.example.com/psso.mobileconfig",
      "pssoConfigProfileId": "1",
      "manifestUrl": "https://mdmserver.example.com/psso-app.plist",
      "authUrl": "https://idp.example.com/authenticate",
      "id": "1",
      "profileUuid": "29d-a8d8f-b8sdjndf-dsa9",
      "siteId": "5",
      "versionLock": 0,
      "accountSettings": {
        "id": "1",
        "payloadConfigured": true,
        "localAdminAccountEnabled": true,
        "adminUsername": "admin",
        "hiddenAdminAccount": false,
        "localUserManaged": true,
        "userAccountType": "STANDARD",
        "versionLock": 4,
        "prefillPrimaryAccountInfoFeatureEnabled": true,
        "prefillType": "DEVICE_OWNER",
        "prefillAccountFullName": "TestUser FullName",
        "prefillAccountUserName": "UserName",
        "preventPrefillInfoFromModification": false
      }
    }
  ]
}
-----
Create a Computer Prestage
post
https://yourServer.jamfcloud.com/api/v3/computer-prestages

Create a computer prestage

Body Params
Computer Prestage to create. ids defined in this body will be ignored

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
installProfilesDuringSetup
boolean
required

true
prestageInstalledProfileIds
array of strings
required

string

1

string

2

ADD string
customPackageIds
array of strings
required

string

1

string

2

ADD string
customPackageDistributionPointId
string
required
1
enableRecoveryLock
boolean

false
recoveryLockPasswordType
string
enum

MANUAL
Allowed:

MANUAL

RANDOM
rotateRecoveryLockPassword
boolean

true
prestageMinimumOsTargetVersionType
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersion
string
length ≥ 0
17.1
pssoEnabled
boolean
Defaults to false
Indicates whether Platform SSO (PSSO) is enabled for this computer prestage, regardless of unattended or 403 workflows. When enabled, the PSSO application will be deployed to devices during the setup process to facilitate single sign-on (SSO) for users.


false
platformSsoAppBundleId
string
The bundle identifier for the Platform SSO (PSSO) application unattended workflow. This identifier is used to specify which PSSO app should be deployed to devices during the setup process.

com.okta.mobile
profileUrl
string | null
The URL to the configuration profile for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process. Users should use either profileUrl or populate pssoConfigProfileId, but not both.

https://mdmserver.example.com/psso.mobileconfig
pssoConfigProfileId
string | null
The identifier for the configuration profile associated with the Platform SSO (PSSO) application 403 workflow. This ID is used to specify which configuration profile should be applied to devices during the setup process when PSSO is enabled. Users should use either pssoConfigProfileId or populate profileUrl, but not both.

1
manifestUrl
string | null
The URL to the manifest file for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process.

https://mdmserver.example.com/psso-app.plist
authUrl
string | null
The URL to the identity provider (IdP) for authentication for the Platform SSO (PSSO) application 403 workflow. This URL is used in conjunction with PSSO to facilitate single sign-on for users during the device setup process.

accountSettings
object

accountSettings object
id
string
id of Account Settings

1
payloadConfigured
boolean
Defaults to false

false
localAdminAccountEnabled
boolean
Defaults to false

false
adminUsername
string
length ≥ 0
admin
adminPassword
string
length ≥ 0
password
hiddenAdminAccount
boolean
Defaults to false

false
localUserManaged
boolean
Defaults to false

false
userAccountType
string
enum
Defaults to STANDARD

STANDARD
Allowed:

ADMINISTRATOR

STANDARD

SKIP
versionLock
integer
Defaults to 0
4
prefillPrimaryAccountInfoFeatureEnabled
boolean
Defaults to false

false
prefillType
string
Defaults to CUSTOM
Values accepted are only CUSTOM and DEVICE_OWNER

CUSTOM
prefillAccountFullName
string
length ≥ 0
TestUser FullName
prefillAccountUserName
string
length ≥ 0
UserName
preventPrefillInfoFromModification
boolean
Defaults to false

false
recoveryLockPassword
password
length ≥ 1
•••••••••••
Responses

201
Computer Prestage was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/computer-prestages \
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
    "newKey": false
  },
  "locationInformation": {
    "username": "name",
    "realname": "realName",
    "email": "test@jamf.com",
    "phone": "123-456-7890",
    "room": "room",
    "position": "postion",
    "departmentId": "1",
    "buildingId": "1",
    "versionLock": 1,
    "id": "-1"
  },
  "purchasingInformation": {
    "leased": true,
    "purchased": true,
    "id": "-1",
    "appleCareId": "abcd",
    "poNumber": "53-1",
    "vendor": "Example Vendor",
    "purchasePrice": "$500",
    "purchasingAccount": "admin",
    "lifeExpectancy": 5,
    "purchasingContact": "true",
    "leaseDate": "2019-01-01",
    "poDate": "2019-01-01",
    "warrantyDate": "2019-01-01",
    "versionLock": 1
  },
  "autoAdvanceSetup": true,
  "installProfilesDuringSetup": true,
  "pssoEnabled": false,
  "accountSettings": {
    "payloadConfigured": false,
    "localAdminAccountEnabled": false,
    "hiddenAdminAccount": false,
    "localUserManaged": false,
    "userAccountType": "STANDARD",
    "versionLock": 4,
    "prefillPrimaryAccountInfoFeatureEnabled": false,
    "prefillType": "CUSTOM",
    "preventPrefillInfoFromModification": false,
    "id": "1",
    "adminUsername": "admin",
    "adminPassword": "password",
    "prefillAccountFullName": "TestUser FullName",
    "prefillAccountUserName": "UserName"
  },
  "displayName": "Example Mobile Prestage Name",
  "supportPhoneNumber": "5555555555",
  "supportEmailAddress": "example@example.com",
  "department": "Oxbow",
  "enrollmentSiteId": "-1",
  "deviceEnrollmentProgramInstanceId": "5",
  "anchorCertificates": [
    "xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
    "xNE5HRgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
  ],
  "enrollmentCustomizationId": "2",
  "language": "en",
  "region": "US",
  "prestageInstalledProfileIds": [
    "1",
    "2"
  ],
  "customPackageIds": [
    "1",
    "2"
  ],
  "customPackageDistributionPointId": "1",
  "enableRecoveryLock": false,
  "recoveryLockPasswordType": "MANUAL",
  "rotateRecoveryLockPassword": true,
  "prestageMinimumOsTargetVersionType": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersion": "17.1",
  "platformSsoAppBundleId": "com.okta.mobile",
  "profileUrl": "https://mdmserver.example.com/psso.mobileconfig",
  "pssoConfigProfileId": "1",
  "manifestUrl": "https://mdmserver.example.com/psso-app.plist",
  "recoveryLockPassword": "password123"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a Computer Prestage with the supplied id
get
https://yourServer.jamfcloud.com/api/v3/computer-prestages/{id}


Retrieves a Computer Prestage with the supplied id

Path Params
id
string
required
Computer Prestage identifier

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
installProfilesDuringSetup
boolean
required
prestageInstalledProfileIds
array of strings
required
customPackageIds
array of strings
required
customPackageDistributionPointId
string
required
enableRecoveryLock
boolean
recoveryLockPasswordType
string
enum
MANUAL RANDOM

rotateRecoveryLockPassword
boolean
prestageMinimumOsTargetVersionType
string
enum
NO_ENFORCEMENT MINIMUM_OS_LATEST_VERSION MINIMUM_OS_LATEST_MAJOR_VERSION MINIMUM_OS_LATEST_MINOR_VERSION MINIMUM_OS_SPECIFIC_VERSION

minimumOsSpecificVersion
string
length ≥ 0
pssoEnabled
boolean
Defaults to false
Indicates whether Platform SSO (PSSO) is enabled for this computer prestage, regardless of unattended or 403 workflows. When enabled, the PSSO application will be deployed to devices during the setup process to facilitate single sign-on (SSO) for users.

platformSsoAppBundleId
string
The bundle identifier for the Platform SSO (PSSO) application unattended workflow. This identifier is used to specify which PSSO app should be deployed to devices during the setup process.

profileUrl
string | null
The URL to the configuration profile for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process. Users should use either profileUrl or populate pssoConfigProfileId, but not both.

pssoConfigProfileId
string | null
The identifier for the configuration profile associated with the Platform SSO (PSSO) application 403 workflow. This ID is used to specify which configuration profile should be applied to devices during the setup process when PSSO is enabled. Users should use either pssoConfigProfileId or populate profileUrl, but not both.

manifestUrl
string | null
The URL to the manifest file for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process.

authUrl
string | null
The URL to the identity provider (IdP) for authentication for the Platform SSO (PSSO) application 403 workflow. This URL is used in conjunction with PSSO to facilitate single sign-on for users during the device setup process.

id
string
length ≥ 1
profileUuid
string
siteId
string
versionLock
integer
≥ 0
accountSettings
object
id
string
id of Account Settings

payloadConfigured
boolean
Defaults to false
localAdminAccountEnabled
boolean
Defaults to false
adminUsername
string
length ≥ 0
hiddenAdminAccount
boolean
Defaults to false
localUserManaged
boolean
Defaults to false
userAccountType
string
enum
Defaults to STANDARD
ADMINISTRATOR STANDARD SKIP

versionLock
integer
≥ 0
Defaults to 0
prefillPrimaryAccountInfoFeatureEnabled
boolean
Defaults to false
prefillType
string
Defaults to CUSTOM
Values accepted are only CUSTOM and DEVICE_OWNER

prefillAccountFullName
string
length ≥ 0
prefillAccountUserName
string
length ≥ 0
preventPrefillInfoFromModification
boolean
Defaults to false

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/computer-prestages/ \
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
  "installProfilesDuringSetup": true,
  "prestageInstalledProfileIds": [
    "1"
  ],
  "customPackageIds": [
    "1"
  ],
  "customPackageDistributionPointId": "1",
  "enableRecoveryLock": true,
  "recoveryLockPasswordType": "MANUAL",
  "rotateRecoveryLockPassword": true,
  "prestageMinimumOsTargetVersionType": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersion": "17.1",
  "pssoEnabled": true,
  "platformSsoAppBundleId": "com.okta.mobile",
  "profileUrl": "https://mdmserver.example.com/psso.mobileconfig",
  "pssoConfigProfileId": "1",
  "manifestUrl": "https://mdmserver.example.com/psso-app.plist",
  "authUrl": "https://idp.example.com/authenticate",
  "id": "1",
  "profileUuid": "29d-a8d8f-b8sdjndf-dsa9",
  "siteId": "5",
  "versionLock": 0,
  "accountSettings": {
    "id": "1",
    "payloadConfigured": true,
    "localAdminAccountEnabled": true,
    "adminUsername": "admin",
    "hiddenAdminAccount": false,
    "localUserManaged": true,
    "userAccountType": "STANDARD",
    "versionLock": 4,
    "prefillPrimaryAccountInfoFeatureEnabled": true,
    "prefillType": "DEVICE_OWNER",
    "prefillAccountFullName": "TestUser FullName",
    "prefillAccountUserName": "UserName",
    "preventPrefillInfoFromModification": false
  }
}
-----
Update a Computer Prestage
put
https://yourServer.jamfcloud.com/api/v3/computer-prestages/{id}


Updates a Computer Prestage

Path Params
id
string
required
Computer Prestage identifier

Body Params
Computer Prestage to update

displayName
string
required
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
supportEmailAddress
string
required
department
string
required
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
installProfilesDuringSetup
boolean
required

true
prestageInstalledProfileIds
array of strings
required

ADD string
customPackageIds
array of strings
required

ADD string
customPackageDistributionPointId
string
required
enableRecoveryLock
boolean

true
recoveryLockPasswordType
string
enum

MANUAL
Allowed:

MANUAL

RANDOM
rotateRecoveryLockPassword
boolean

true
prestageMinimumOsTargetVersionType
string
enum

MINIMUM_OS_LATEST_VERSION
Allowed:

NO_ENFORCEMENT

MINIMUM_OS_LATEST_VERSION

MINIMUM_OS_LATEST_MAJOR_VERSION

MINIMUM_OS_LATEST_MINOR_VERSION

MINIMUM_OS_SPECIFIC_VERSION
minimumOsSpecificVersion
string
length ≥ 0
pssoEnabled
boolean
Defaults to false
Indicates whether Platform SSO (PSSO) is enabled for this computer prestage, regardless of unattended or 403 workflows. When enabled, the PSSO application will be deployed to devices during the setup process to facilitate single sign-on (SSO) for users.


false
platformSsoAppBundleId
string
The bundle identifier for the Platform SSO (PSSO) application unattended workflow. This identifier is used to specify which PSSO app should be deployed to devices during the setup process.

profileUrl
string | null
The URL to the configuration profile for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process. Users should use either profileUrl or populate pssoConfigProfileId, but not both.

pssoConfigProfileId
string | null
The identifier for the configuration profile associated with the Platform SSO (PSSO) application 403 workflow. This ID is used to specify which configuration profile should be applied to devices during the setup process when PSSO is enabled. Users should use either pssoConfigProfileId or populate profileUrl, but not both.

manifestUrl
string | null
The URL to the manifest file for the Platform SSO (PSSO) application 403 workflow. This URL is used when deploying the PSSO app to devices during the setup process.

authUrl
string | null
The URL to the identity provider (IdP) for authentication for the Platform SSO (PSSO) application 403 workflow. This URL is used in conjunction with PSSO to facilitate single sign-on for users during the device setup process.

accountSettings
object

accountSettings object
recoveryLockPassword
password
length ≥ 1
versionLock
integer
≥ 0
Responses


curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/computer-prestages/ \
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
  "installProfilesDuringSetup": true,
  "pssoEnabled": false,
  "accountSettings": {
    "payloadConfigured": false,
    "localAdminAccountEnabled": false,
    "hiddenAdminAccount": false,
    "localUserManaged": false,
    "userAccountType": "STANDARD",
    "versionLock": 0,
    "prefillPrimaryAccountInfoFeatureEnabled": false,
    "prefillType": "CUSTOM",
    "preventPrefillInfoFromModification": false
  }
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
  "installProfilesDuringSetup": true,
  "prestageInstalledProfileIds": [
    "1"
  ],
  "customPackageIds": [
    "1"
  ],
  "customPackageDistributionPointId": "1",
  "enableRecoveryLock": true,
  "recoveryLockPasswordType": "MANUAL",
  "rotateRecoveryLockPassword": true,
  "prestageMinimumOsTargetVersionType": "MINIMUM_OS_LATEST_VERSION",
  "minimumOsSpecificVersion": "17.1",
  "pssoEnabled": true,
  "platformSsoAppBundleId": "com.okta.mobile",
  "profileUrl": "https://mdmserver.example.com/psso.mobileconfig",
  "pssoConfigProfileId": "1",
  "manifestUrl": "https://mdmserver.example.com/psso-app.plist",
  "authUrl": "https://idp.example.com/authenticate",
  "id": "1",
  "profileUuid": "29d-a8d8f-b8sdjndf-dsa9",
  "siteId": "5",
  "versionLock": 0,
  "accountSettings": {
    "id": "1",
    "payloadConfigured": true,
    "localAdminAccountEnabled": true,
    "adminUsername": "admin",
    "hiddenAdminAccount": false,
    "localUserManaged": true,
    "userAccountType": "STANDARD",
    "versionLock": 4,
    "prefillPrimaryAccountInfoFeatureEnabled": true,
    "prefillType": "DEVICE_OWNER",
    "prefillAccountFullName": "TestUser FullName",
    "prefillAccountUserName": "UserName",
    "preventPrefillInfoFromModification": false
  }
}
-----

Delete a Computer Prestage with the supplied id
delete
https://yourServer.jamfcloud.com/api/v3/computer-prestages/{id}


Deletes a Computer Prestage with the supplied id

Path Params
id
string
required
Computer Prestage identifier

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v3/computer-prestages/

