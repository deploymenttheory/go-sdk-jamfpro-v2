Retrieve the Account Driven User Enrollment Session Token Settings
get
https://yourServer.jamfcloud.com/api/v1/adue-session-token-settings


Retrieve the Account Driven User Enrollment Session Token Settings

Response

200
Successful response

Response body
object
enabled
boolean
expirationIntervalDays
integer
expirationIntervalSeconds
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/adue-session-token-settings \
     --header 'accept: application/json'

{
  "enabled": false,
  "expirationIntervalDays": 1,
  "expirationIntervalSeconds": 86400
}
-----
Update Account Driven User Enrollment Session Token Settings.
put
https://yourServer.jamfcloud.com/api/v1/adue-session-token-settings


Update the Account Driven User Enrollment Session Token Settings object.

Body Params
Update Account Driven User Enrollment Session Token Settings.

enabled
boolean

true
expirationIntervalDays
integer
1
expirationIntervalSeconds
integer
86400
Responses

200
Successfully updated Account Driven User Enrollment Session Token Settings object

Response body
object
enabled
boolean
expirationIntervalDays
integer
expirationIntervalSeconds
integer

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/adue-session-token-settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "enabled": true,
  "expirationIntervalDays": 1,
  "expirationIntervalSeconds": 86400
}
'

{
  "enabled": false,
  "expirationIntervalDays": 1,
  "expirationIntervalSeconds": 86400
}
-----
Get sorted and paged Enrollment history object
get
https://yourServer.jamfcloud.com/api/v2/enrollment/history

Gets sorted and paged Enrollment history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

date:desc

ADD string
Response

200
Details of enrollment history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v2/enrollment/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Enrollment history object notes
post
https://yourServer.jamfcloud.com/api/v2/enrollment/history

Adds Enrollment history object notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of enrollment history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/enrollment/history \
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
Export enrollment history collection
post
https://yourServer.jamfcloud.com/api/v2/enrollment/history/export


Export enrollment history collection

Query Params
export-fields
array of strings
Defaults to
Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username


ADD string
export-labels
array of strings
Defaults to
Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username


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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name=="script"

Body Params
Optional. Override query parameters since they can make URI exceed 2,000 character limit.

page
integer | null
Defaults to 0
0
pageSize
integer | null
Defaults to 100
100
sort
array of strings | null
Defaults to id:desc
Sorting criteria in the format: [[:asc/desc]. Default direction when not stated is ascending.


string

id:desc

ADD string
filter
string | null
fields
array of objects | null
Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported.


ADD object
Headers
accept
string
enum
Defaults to text/csv,application/json
Generated from available response content types


text/csv,application/json
Allowed:

application/json

text/csv,application/json
Responses

200
Export successful

Response body
json

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v2/enrollment/history/export?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: text/csv,application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "page": 0,
  "pageSize": 100,
  "sort": [
    "id:desc"
  ]
}
'

Username,DATE,NOTES,Details
admin, 2022-02-04T11:56:26.343Z, Edited ,Re-enrollment Restricted true
-----
Retrieve the configured LDAP groups configured for User-Initiated Enrollment.
get
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups


Retrieves the configured LDAP groups configured for User-Initiated Enrollment.

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
Defaults to name:asc
Sorting criteria in the format: property:asc/desc. Default sort is name:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

name:asc

ADD string
all-users-option-first
boolean
Defaults to false
Return "All LDAP Users" option on the first position if it is present in the current page


false
Response

200
Found access groups matching search params.

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
Autogenerated ID

groupId
string
required
LDAP Group ID

ldapServerId
string
required
name
string
required
siteId
string
enterpriseEnrollmentEnabled
boolean
personalEnrollmentEnabled
boolean
accountDrivenUserEnrollmentEnabled
boolean
requireEula
boolean

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/enrollment/access-groups?page=0&page-size=100&sort=name%3Aasc&all-users-option-first=false' \
     --header 'accept: application/json'

{
  "totalCount": 10,
  "results": [
    {
      "id": "1",
      "groupId": "1",
      "ldapServerId": "1",
      "name": "Grade School Pupils",
      "siteId": "-1",
      "enterpriseEnrollmentEnabled": false,
      "personalEnrollmentEnabled": false,
      "accountDrivenUserEnrollmentEnabled": false,
      "requireEula": false
    }
  ]
}
-----
Add the configured LDAP group for User-Initiated Enrollment.
post
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups


Add the configured LDAP group for User-Initiated Enrollment.

Body Params
Configured LDAP group to create.

groupId
string
required
LDAP Group ID

1
ldapServerId
string
required
1
name
string
required
Grade School Pupils
siteId
string
-1
enterpriseEnrollmentEnabled
boolean

true
personalEnrollmentEnabled
boolean

false
accountDrivenUserEnrollmentEnabled
boolean

false
requireEula
boolean

false
Responses

201
Configured LDAP group record was created.

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/access-groups \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupId": "1",
  "ldapServerId": "1",
  "name": "Grade School Pupils",
  "siteId": "-1",
  "enterpriseEnrollmentEnabled": true
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve the configured LDAP groups configured for User-Initiated Enrollment
get
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups/{id}


Retrieves the configured LDAP groups configured for User-Initiated Enrollment.

Path Params
id
string
required
Autogenerated Access Group ID.

Responses

200
Successful query

Response body
object
id
string
Autogenerated ID

groupId
string
required
LDAP Group ID

ldapServerId
string
required
name
string
required
siteId
string
enterpriseEnrollmentEnabled
boolean
personalEnrollmentEnabled
boolean
accountDrivenUserEnrollmentEnabled
boolean
requireEula
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/access-groups/ \
     --header 'accept: application/json'

{
  "id": "1",
  "groupId": "1",
  "ldapServerId": "1",
  "name": "Grade School Pupils",
  "siteId": "-1",
  "enterpriseEnrollmentEnabled": false,
  "personalEnrollmentEnabled": false,
  "accountDrivenUserEnrollmentEnabled": false,
  "requireEula": false
}
-----
Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated.
put
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups/{id}


Modify the configured LDAP groups configured for User-Initiated Enrollment. Only exiting Access Groups can be updated.

Path Params
id
string
required
Autogenerated Access Group ID.

1
Body Params
groupId
string
required
LDAP Group ID

1
ldapServerId
string
required
1
name
string
required
Grade School Pupils
siteId
string
-1
enterpriseEnrollmentEnabled
boolean

true
personalEnrollmentEnabled
boolean

false
accountDrivenUserEnrollmentEnabled
boolean

false
requireEula
boolean

true
Responses

200
Successful update

Response body
object
id
string
Autogenerated ID

groupId
string
required
LDAP Group ID

ldapServerId
string
required
name
string
required
siteId
string
enterpriseEnrollmentEnabled
boolean
personalEnrollmentEnabled
boolean
accountDrivenUserEnrollmentEnabled
boolean
requireEula
boolean

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/access-groups/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupId": "1",
  "ldapServerId": "1",
  "name": "Grade School Pupils",
  "siteId": "-1",
  "requireEula": true,
  "accountDrivenUserEnrollmentEnabled": false,
  "personalEnrollmentEnabled": false,
  "enterpriseEnrollmentEnabled": true
}
'
{
  "id": "1",
  "groupId": "1",
  "ldapServerId": "1",
  "name": "Grade School Pupils",
  "siteId": "-1",
  "enterpriseEnrollmentEnabled": false,
  "personalEnrollmentEnabled": false,
  "accountDrivenUserEnrollmentEnabled": false,
  "requireEula": false
}
-----
Delete an LDAP group's access to user initiated Enrollment.
delete
https://yourServer.jamfcloud.com/api/v3/enrollment/access-groups/{id}


Deletes an LDAP group's access to user initiated enrollment. The group "All LDAP Users" cannot be deleted, but it can be modified to disallow User-Initiated Enrollment.

Path Params
id
string
required
Autogenerated Access Group ID.

Responses
204
Successful deletion

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/access-groups/ \
     --header 'accept: application/json'

-----
Retrieve the list of languages and corresponding ISO 639-1 Codes but only those not already added to Enrollment
get
https://yourServer.jamfcloud.com/api/v3/enrollment/filtered-language-codes


Retrieves the list of languages and corresponding ISO 639-1 Codes, but only those not already added to Enrollment.

Response

200
Retrieves the list of languages and corresponding ISO 639-1 Codes, but only those not already added to Enrollment.

Response body
array of objects
object
value
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/filtered-language-codes \
     --header 'accept: application/json'

[
  {
    "value": "en",
    "name": "English"
  }
]
-----
Retrieve the list of languages and corresponding ISO 639-1 Codes
get
https://yourServer.jamfcloud.com/api/v3/enrollment/language-codes


Retrieves the list of languages and corresponding ISO 639-1 Codes.

Response

200
List of languages and corresponding ISO 639-1 Codes.

Response body
array of objects
object
value
string
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/language-codes \
     --header 'accept: application/json'

[
  {
    "value": "en",
    "name": "English"
  }
]
-----
Get an array of the language codes that have Enrollment messaging
get
https://yourServer.jamfcloud.com/api/v3/enrollment/languages


Returns an array of the language codes that have enrollment messaging currently configured.

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
Defaults to languageCode:asc
Sorting criteria in the format: property:asc/desc. Default sort is languageCode:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc.


string

languageCode:asc

ADD string
Response

200
Found languages matching search params.

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
languageCode
string
name
string
title
string
loginDescription
string
username
string
password
string
loginButton
string
deviceClassDescription
string
deviceClassPersonal
string
deviceClassPersonalDescription
string
deviceClassEnterprise
string
deviceClassEnterpriseDescription
string
deviceClassButton
string
personalEula
string
enterpriseEula
string
eulaButton
string
siteDescription
string
certificateText
string
certificateButton
string
certificateProfileName
string
certificateProfileDescription
string
personalText
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalButton
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileName
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileDescription
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

userEnrollmentText
string
Defaults to Enter your Managed Apple id to install the MDM Profile.
userEnrollmentButton
string
Defaults to Continue
userEnrollmentProfileName
string
Defaults to MDM Profile
userEnrollmentProfileDescription
string
Defaults to MDM Profile for mobile device management
enterpriseText
string
enterpriseButton
string
enterpriseProfileName
string
enterpriseProfileDescription
string
enterprisePending
string
quickAddText
string
quickAddButton
string
quickAddName
string
Defaults to QuickAdd.pkg
quickAddPending
string
completeMessage
string
failedMessage
string
tryAgainButton
string
checkNowButton
string
checkEnrollmentMessage
string
logoutButton
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v3/enrollment/languages?page=0&page-size=100&sort=languageCode%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 10,
  "results": [
    {
      "languageCode": "en",
      "name": "English",
      "title": "Enroll Your Device",
      "loginDescription": "Log in to enroll your device.",
      "username": "admin",
      "password": "12345",
      "loginButton": "Log in",
      "deviceClassDescription": "Specify if this device is institutionally owned or personally owned.",
      "deviceClassPersonal": "Personally Owned",
      "deviceClassPersonalDescription": "For personally owned devices, IT administrators **can**\n\n         *   Lock the device\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n\n\n         For personally owned devices, IT administrators **cannot**\n\n         *   Wipe all data and settings from your device\n         *   Track the location of your device\n         *   Remove anything they did not install\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n",
      "deviceClassEnterprise": "Institutionally Owned",
      "deviceClassEnterpriseDescription": "For institutionally owned devices, IT administrators **can**\n\n         *   Wipe all data and settings from the device\n         *   Lock the device\n         *   Remove the passcode\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n\n         For institutionally owned devices, IT administrators **cannot**:\n\n         *   Remove anything they did not install\n",
      "deviceClassButton": "Enroll",
      "personalEula": "Personal Eula",
      "enterpriseEula": "Enterprise Eula",
      "eulaButton": "Accept",
      "siteDescription": "Select the site to use for enrolling this computer or mobile device.",
      "certificateText": "To continue with enrollment, you need to install the CA certificate for your organization.",
      "certificateButton": "Continue",
      "certificateProfileName": "CA Certificate",
      "certificateProfileDescription": "CA Certificate for mobile device management",
      "userEnrollmentText": "Enter your Managed Apple id to install the MDM Profile.",
      "userEnrollmentButton": "Continue",
      "userEnrollmentProfileName": "MDM Profile",
      "userEnrollmentProfileDescription": "MDM Profile for mobile device management",
      "enterpriseText": "To continue with enrollment, you need to install the MDM profile for your organization.",
      "enterpriseButton": "Continue",
      "enterpriseProfileName": "MDM Profile",
      "enterpriseProfileDescription": "MDM Profile for mobile device management",
      "enterprisePending": "To continue with enrollment, install the CA Certificate and MDM Profile that were downloaded to your computer.",
      "quickAddText": "Download and install this package.",
      "quickAddButton": "Download",
      "quickAddName": "QuickAdd.pkg",
      "quickAddPending": "Install the downloaded QuickAdd.pkg.",
      "completeMessage": "The enrollment process is complete.",
      "failedMessage": "The enrollment process could not be completed. Contact your IT administrator.",
      "tryAgainButton": "Try Again",
      "checkNowButton": "Proceed",
      "checkEnrollmentMessage": "Tap \"Proceed\" to view the enrollment status for this device.",
      "logoutButton": "Log Out"
    }
  ]
}
-----
Delete multiple configured languages from User-Initiated Enrollment settings
post
https://yourServer.jamfcloud.com/api/v3/enrollment/languages/delete-multiple


Delete multiple configured languages from User-Initiated Enrollment settings

Body Params
ids of each language to delete

ids
array of strings

string


ADD string
Responses
204
All languages ids passed in request sucessfully deleted.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/languages/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json'
-----
Retrieve the Enrollment messaging for a language
get
https://yourServer.jamfcloud.com/api/v3/enrollment/languages/{languageId}


Retrieves the enrollment messaging for a language.

Path Params
languageId
string
required
Two letter ISO 639-1 Language Code

Responses

200
Successful response

Response body
object
languageCode
string
name
string
title
string
loginDescription
string
username
string
password
string
loginButton
string
deviceClassDescription
string
deviceClassPersonal
string
deviceClassPersonalDescription
string
deviceClassEnterprise
string
deviceClassEnterpriseDescription
string
deviceClassButton
string
personalEula
string
enterpriseEula
string
eulaButton
string
siteDescription
string
certificateText
string
certificateButton
string
certificateProfileName
string
certificateProfileDescription
string
personalText
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalButton
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileName
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileDescription
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

userEnrollmentText
string
Defaults to Enter your Managed Apple id to install the MDM Profile.
userEnrollmentButton
string
Defaults to Continue
userEnrollmentProfileName
string
Defaults to MDM Profile
userEnrollmentProfileDescription
string
Defaults to MDM Profile for mobile device management
enterpriseText
string
enterpriseButton
string
enterpriseProfileName
string
enterpriseProfileDescription
string
enterprisePending
string
quickAddText
string
quickAddButton
string
quickAddName
string
Defaults to QuickAdd.pkg
quickAddPending
string
completeMessage
string
failedMessage
string
tryAgainButton
string
checkNowButton
string
checkEnrollmentMessage
string
logoutButton
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/languages/ \
     --header 'accept: application/json'

{
  "languageCode": "en",
  "name": "English",
  "title": "Enroll Your Device",
  "loginDescription": "Log in to enroll your device.",
  "username": "admin",
  "password": "12345",
  "loginButton": "Log in",
  "deviceClassDescription": "Specify if this device is institutionally owned or personally owned.",
  "deviceClassPersonal": "Personally Owned",
  "deviceClassPersonalDescription": "For personally owned devices, IT administrators **can**\n\n         *   Lock the device\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n\n\n         For personally owned devices, IT administrators **cannot**\n\n         *   Wipe all data and settings from your device\n         *   Track the location of your device\n         *   Remove anything they did not install\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n",
  "deviceClassEnterprise": "Institutionally Owned",
  "deviceClassEnterpriseDescription": "For institutionally owned devices, IT administrators **can**\n\n         *   Wipe all data and settings from the device\n         *   Lock the device\n         *   Remove the passcode\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n\n         For institutionally owned devices, IT administrators **cannot**:\n\n         *   Remove anything they did not install\n",
  "deviceClassButton": "Enroll",
  "personalEula": "Personal Eula",
  "enterpriseEula": "Enterprise Eula",
  "eulaButton": "Accept",
  "siteDescription": "Select the site to use for enrolling this computer or mobile device.",
  "certificateText": "To continue with enrollment, you need to install the CA certificate for your organization.",
  "certificateButton": "Continue",
  "certificateProfileName": "CA Certificate",
  "certificateProfileDescription": "CA Certificate for mobile device management",
  "userEnrollmentText": "Enter your Managed Apple id to install the MDM Profile.",
  "userEnrollmentButton": "Continue",
  "userEnrollmentProfileName": "MDM Profile",
  "userEnrollmentProfileDescription": "MDM Profile for mobile device management",
  "enterpriseText": "To continue with enrollment, you need to install the MDM profile for your organization.",
  "enterpriseButton": "Continue",
  "enterpriseProfileName": "MDM Profile",
  "enterpriseProfileDescription": "MDM Profile for mobile device management",
  "enterprisePending": "To continue with enrollment, install the CA Certificate and MDM Profile that were downloaded to your computer.",
  "quickAddText": "Download and install this package.",
  "quickAddButton": "Download",
  "quickAddName": "QuickAdd.pkg",
  "quickAddPending": "Install the downloaded QuickAdd.pkg.",
  "completeMessage": "The enrollment process is complete.",
  "failedMessage": "The enrollment process could not be completed. Contact your IT administrator.",
  "tryAgainButton": "Try Again",
  "checkNowButton": "Proceed",
  "checkEnrollmentMessage": "Tap \"Proceed\" to view the enrollment status for this device.",
  "logoutButton": "Log Out"
}
-----
Edit Enrollment messaging for a language
put
https://yourServer.jamfcloud.com/api/v3/enrollment/languages/{languageId}


Edit enrollment messaging for a language.

Path Params
languageId
string
required
Two letter ISO 639-1 Language Code

Body Params
languageCode
string
name
string
title
string
loginDescription
string
username
string
password
string
loginButton
string
deviceClassDescription
string
deviceClassPersonal
string
deviceClassPersonalDescription
string
deviceClassEnterprise
string
deviceClassEnterpriseDescription
string
deviceClassButton
string
personalEula
string
enterpriseEula
string
eulaButton
string
siteDescription
string
certificateText
string
certificateButton
string
certificateProfileName
string
certificateProfileDescription
string
userEnrollmentText
string
Defaults to Enter your Managed Apple id to install the MDM Profile.
Enter your Managed Apple id to install the MDM Profile.
userEnrollmentButton
string
Defaults to Continue
Continue
userEnrollmentProfileName
string
Defaults to MDM Profile
MDM Profile
userEnrollmentProfileDescription
string
Defaults to MDM Profile for mobile device management
MDM Profile for mobile device management
enterpriseText
string
enterpriseButton
string
enterpriseProfileName
string
enterpriseProfileDescription
string
enterprisePending
string
quickAddText
string
quickAddButton
string
quickAddName
string
Defaults to QuickAdd.pkg
QuickAdd.pkg
quickAddPending
string
completeMessage
string
failedMessage
string
tryAgainButton
string
checkNowButton
string
checkEnrollmentMessage
string
logoutButton
string
Responses

200
Successful response

Response body
object
languageCode
string
name
string
title
string
loginDescription
string
username
string
password
string
loginButton
string
deviceClassDescription
string
deviceClassPersonal
string
deviceClassPersonalDescription
string
deviceClassEnterprise
string
deviceClassEnterpriseDescription
string
deviceClassButton
string
personalEula
string
enterpriseEula
string
eulaButton
string
siteDescription
string
certificateText
string
certificateButton
string
certificateProfileName
string
certificateProfileDescription
string
personalText
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalButton
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileName
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

personalProfileDescription
string
deprecated
Deprecated as of 11.25. This field always returns empty string in GET responses and ignores any input values in PUT requests.

userEnrollmentText
string
Defaults to Enter your Managed Apple id to install the MDM Profile.
userEnrollmentButton
string
Defaults to Continue
userEnrollmentProfileName
string
Defaults to MDM Profile
userEnrollmentProfileDescription
string
Defaults to MDM Profile for mobile device management
enterpriseText
string
enterpriseButton
string
enterpriseProfileName
string
enterpriseProfileDescription
string
enterprisePending
string
quickAddText
string
quickAddButton
string
quickAddName
string
Defaults to QuickAdd.pkg
quickAddPending
string
completeMessage
string
failedMessage
string
tryAgainButton
string
checkNowButton
string
checkEnrollmentMessage
string
logoutButton
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/languages/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "userEnrollmentText": "Enter your Managed Apple id to install the MDM Profile.",
  "userEnrollmentButton": "Continue",
  "userEnrollmentProfileName": "MDM Profile",
  "userEnrollmentProfileDescription": "MDM Profile for mobile device management",
  "quickAddName": "QuickAdd.pkg"
}
'

{
  "languageCode": "en",
  "name": "English",
  "title": "Enroll Your Device",
  "loginDescription": "Log in to enroll your device.",
  "username": "admin",
  "password": "12345",
  "loginButton": "Log in",
  "deviceClassDescription": "Specify if this device is institutionally owned or personally owned.",
  "deviceClassPersonal": "Personally Owned",
  "deviceClassPersonalDescription": "For personally owned devices, IT administrators **can**\n\n         *   Lock the device\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n\n\n         For personally owned devices, IT administrators **cannot**\n\n         *   Wipe all data and settings from your device\n         *   Track the location of your device\n         *   Remove anything they did not install\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n",
  "deviceClassEnterprise": "Institutionally Owned",
  "deviceClassEnterpriseDescription": "For institutionally owned devices, IT administrators **can**\n\n         *   Wipe all data and settings from the device\n         *   Lock the device\n         *   Remove the passcode\n         *   Apply institutional settings\n         *   Install and remove institutional data\n         *   Install and remove institutional apps\n         *   Add/remove configuration profiles\n         *   Add/remove provisioning profiles\n\n         For institutionally owned devices, IT administrators **cannot**:\n\n         *   Remove anything they did not install\n",
  "deviceClassButton": "Enroll",
  "personalEula": "Personal Eula",
  "enterpriseEula": "Enterprise Eula",
  "eulaButton": "Accept",
  "siteDescription": "Select the site to use for enrolling this computer or mobile device.",
  "certificateText": "To continue with enrollment, you need to install the CA certificate for your organization.",
  "certificateButton": "Continue",
  "certificateProfileName": "CA Certificate",
  "certificateProfileDescription": "CA Certificate for mobile device management",
  "userEnrollmentText": "Enter your Managed Apple id to install the MDM Profile.",
  "userEnrollmentButton": "Continue",
  "userEnrollmentProfileName": "MDM Profile",
  "userEnrollmentProfileDescription": "MDM Profile for mobile device management",
  "enterpriseText": "To continue with enrollment, you need to install the MDM profile for your organization.",
  "enterpriseButton": "Continue",
  "enterpriseProfileName": "MDM Profile",
  "enterpriseProfileDescription": "MDM Profile for mobile device management",
  "enterprisePending": "To continue with enrollment, install the CA Certificate and MDM Profile that were downloaded to your computer.",
  "quickAddText": "Download and install this package.",
  "quickAddButton": "Download",
  "quickAddName": "QuickAdd.pkg",
  "quickAddPending": "Install the downloaded QuickAdd.pkg.",
  "completeMessage": "The enrollment process is complete.",
  "failedMessage": "The enrollment process could not be completed. Contact your IT administrator.",
  "tryAgainButton": "Try Again",
  "checkNowButton": "Proceed",
  "checkEnrollmentMessage": "Tap \"Proceed\" to view the enrollment status for this device.",
  "logoutButton": "Log Out"
}
-----
Delete the Enrollment messaging for a language
delete
https://yourServer.jamfcloud.com/api/v3/enrollment/languages/{languageId}


Delete the enrollment messaging for a language.

Path Params
languageId
string
required
Two letter ISO 639-1 Language Code

Responses

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v3/enrollment/languages/ \
     --header 'accept: application/json'

-----
Get Enrollment object and Re-enrollment settings
get
https://yourServer.jamfcloud.com/api/v4/enrollment

Gets Enrollment object and re-enrollment settings.

Response

200
Successful response

Response body
object
installSingleProfile
boolean
Defaults to false
signingMdmProfileEnabled
boolean
Defaults to false
mdmSigningCertificate
object | null
filename
string
Defaults to null
md5Sum
string
The md5 checksum of the certificate file. Intended to be used in verification the cert being used to sign QuickAdd packages.

restrictReenrollment
boolean
Defaults to false
flushLocationInformation
boolean
Defaults to false
flushLocationHistoryInformation
boolean
Defaults to false
flushPolicyHistory
boolean
Defaults to false
flushExtensionAttributes
boolean
Defaults to false
flushSoftwareUpdatePlans
boolean
Defaults to false
flushMdmCommandsOnReenroll
string
enum
Defaults to DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED
DELETE_NOTHING DELETE_ERRORS DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED DELETE_EVERYTHING

macOsEnterpriseEnrollmentEnabled
boolean
Defaults to false
managementUsername
string
required
createManagementAccount
boolean
Defaults to true
hideManagementAccount
boolean
Defaults to false
allowSshOnlyManagementAccount
boolean
Defaults to false
ensureSshRunning
boolean
Defaults to true
launchSelfService
boolean
Defaults to false
signQuickAdd
boolean
Defaults to false
developerCertificateIdentity
object | null
filename
string
Defaults to null
md5Sum
string
The md5 checksum of the certificate file. Intended to be used in verification the cert being used to sign QuickAdd packages.

developerCertificateIdentityDetails
object
subject
string
serialNumber
string
mdmSigningCertificateDetails
object
subject
string
serialNumber
string
iosEnterpriseEnrollmentEnabled
boolean
Defaults to true
iosPersonalEnrollmentEnabled
boolean
Defaults to false
personalDeviceEnrollmentType
string
Defaults to USERENROLLMENT
deprecated
Deprecated as of 11.25. This field always returns "USERENROLLMENT" in GET responses and ignores any input values in PUT requests.

accountDrivenUserEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceIosEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceMacosEnrollmentEnabled
boolean
Defaults to false
accountDrivenUserVisionosEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceVisionosEnrollmentEnabled
boolean
Defaults to false
maidUsernameMergeEnabled
boolean
Defaults to false

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v4/enrollment \
     --header 'accept: application/json'

{
  "installSingleProfile": false,
  "signingMdmProfileEnabled": false,
  "mdmSigningCertificate": {
    "filename": "null",
    "md5Sum": ""
  },
  "restrictReenrollment": false,
  "flushLocationInformation": false,
  "flushLocationHistoryInformation": false,
  "flushPolicyHistory": false,
  "flushExtensionAttributes": false,
  "flushSoftwareUpdatePlans": false,
  "flushMdmCommandsOnReenroll": "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED",
  "macOsEnterpriseEnrollmentEnabled": false,
  "managementUsername": "radmin",
  "createManagementAccount": true,
  "hideManagementAccount": false,
  "allowSshOnlyManagementAccount": false,
  "ensureSshRunning": true,
  "launchSelfService": false,
  "signQuickAdd": false,
  "developerCertificateIdentity": {
    "filename": "null",
    "md5Sum": ""
  },
  "developerCertificateIdentityDetails": {
    "subject": "",
    "serialNumber": ""
  },
  "mdmSigningCertificateDetails": {
    "subject": "",
    "serialNumber": ""
  },
  "iosEnterpriseEnrollmentEnabled": true,
  "iosPersonalEnrollmentEnabled": false,
  "accountDrivenUserEnrollmentEnabled": false,
  "accountDrivenDeviceIosEnrollmentEnabled": false,
  "accountDrivenDeviceMacosEnrollmentEnabled": false,
  "accountDrivenUserVisionosEnrollmentEnabled": false,
  "accountDrivenDeviceVisionosEnrollmentEnabled": false,
  "maidUsernameMergeEnabled": false
}
-----
Update Enrollment object
put
https://yourServer.jamfcloud.com/api/v4/enrollment

Update enrollment object. Regarding the developerCertificateIdentity, if this object is omitted, the certificate will not be deleted from Jamf Pro. The identityKeystore is the entire cert file as a base64 encoded string. The md5Sum field is not required in the PUT request, but is calculated and returned in the response.

Body Params
Update enrollment

installSingleProfile
boolean
Defaults to false

false
signingMdmProfileEnabled
boolean
Defaults to false

false
mdmSigningCertificate
object | null

mdmSigningCertificate object | null
restrictReenrollment
boolean
Defaults to false

false
flushLocationInformation
boolean
Defaults to false

false
flushLocationHistoryInformation
boolean
Defaults to false

false
flushPolicyHistory
boolean
Defaults to false

false
flushExtensionAttributes
boolean
Defaults to false

false
flushSoftwareUpdatePlans
boolean
Defaults to false

false
flushMdmCommandsOnReenroll
string
enum
Defaults to DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED

DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED
Allowed:

DELETE_NOTHING

DELETE_ERRORS

DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED

DELETE_EVERYTHING
macOsEnterpriseEnrollmentEnabled
boolean
Defaults to false

false
managementUsername
string
required
createManagementAccount
boolean
Defaults to true

true
hideManagementAccount
boolean
Defaults to false

false
allowSshOnlyManagementAccount
boolean
Defaults to false

false
ensureSshRunning
boolean
Defaults to true

true
launchSelfService
boolean
Defaults to false

false
signQuickAdd
boolean
Defaults to false

false
developerCertificateIdentity
object | null

developerCertificateIdentity object | null
developerCertificateIdentityDetails
object

developerCertificateIdentityDetails object
mdmSigningCertificateDetails
object

mdmSigningCertificateDetails object
iosEnterpriseEnrollmentEnabled
boolean
Defaults to true

true
iosPersonalEnrollmentEnabled
boolean
Defaults to false

false
accountDrivenUserEnrollmentEnabled
boolean
Defaults to false

false
accountDrivenDeviceIosEnrollmentEnabled
boolean
Defaults to false

false
accountDrivenDeviceMacosEnrollmentEnabled
boolean
Defaults to false

false
accountDrivenUserVisionosEnrollmentEnabled
boolean
Defaults to false

false
accountDrivenDeviceVisionosEnrollmentEnabled
boolean
Defaults to false

false
maidUsernameMergeEnabled
boolean
Defaults to false

false
Responses

Response body
object
installSingleProfile
boolean
Defaults to false
signingMdmProfileEnabled
boolean
Defaults to false
mdmSigningCertificate
object | null
filename
string
Defaults to null
md5Sum
string
The md5 checksum of the certificate file. Intended to be used in verification the cert being used to sign QuickAdd packages.

restrictReenrollment
boolean
Defaults to false
flushLocationInformation
boolean
Defaults to false
flushLocationHistoryInformation
boolean
Defaults to false
flushPolicyHistory
boolean
Defaults to false
flushExtensionAttributes
boolean
Defaults to false
flushSoftwareUpdatePlans
boolean
Defaults to false
flushMdmCommandsOnReenroll
string
enum
Defaults to DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED
DELETE_NOTHING DELETE_ERRORS DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED DELETE_EVERYTHING

macOsEnterpriseEnrollmentEnabled
boolean
Defaults to false
managementUsername
string
required
createManagementAccount
boolean
Defaults to true
hideManagementAccount
boolean
Defaults to false
allowSshOnlyManagementAccount
boolean
Defaults to false
ensureSshRunning
boolean
Defaults to true
launchSelfService
boolean
Defaults to false
signQuickAdd
boolean
Defaults to false
developerCertificateIdentity
object | null
filename
string
Defaults to null
md5Sum
string
The md5 checksum of the certificate file. Intended to be used in verification the cert being used to sign QuickAdd packages.

developerCertificateIdentityDetails
object
subject
string
serialNumber
string
mdmSigningCertificateDetails
object
subject
string
serialNumber
string
iosEnterpriseEnrollmentEnabled
boolean
Defaults to true
iosPersonalEnrollmentEnabled
boolean
Defaults to false
personalDeviceEnrollmentType
string
Defaults to USERENROLLMENT
deprecated
Deprecated as of 11.25. This field always returns "USERENROLLMENT" in GET responses and ignores any input values in PUT requests.

accountDrivenUserEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceIosEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceMacosEnrollmentEnabled
boolean
Defaults to false
accountDrivenUserVisionosEnrollmentEnabled
boolean
Defaults to false
accountDrivenDeviceVisionosEnrollmentEnabled
boolean
Defaults to false
maidUsernameMergeEnabled
boolean
Defaults to false

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v4/enrollment \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "installSingleProfile": false,
  "signingMdmProfileEnabled": false,
  "mdmSigningCertificate": {
    "filename": "null"
  },
  "restrictReenrollment": false,
  "flushLocationInformation": false,
  "flushLocationHistoryInformation": false,
  "flushPolicyHistory": false,
  "flushExtensionAttributes": false,
  "flushSoftwareUpdatePlans": false,
  "flushMdmCommandsOnReenroll": "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED",
  "macOsEnterpriseEnrollmentEnabled": false,
  "createManagementAccount": true,
  "hideManagementAccount": false,
  "allowSshOnlyManagementAccount": false,
  "ensureSshRunning": true,
  "launchSelfService": false,
  "signQuickAdd": false,
  "developerCertificateIdentity": {
    "filename": "null"
  },
  "iosEnterpriseEnrollmentEnabled": true,
  "iosPersonalEnrollmentEnabled": false,
  "accountDrivenUserEnrollmentEnabled": false,
  "accountDrivenDeviceIosEnrollmentEnabled": false,
  "accountDrivenDeviceMacosEnrollmentEnabled": false,
  "accountDrivenUserVisionosEnrollmentEnabled": false,
  "accountDrivenDeviceVisionosEnrollmentEnabled": false,
  "maidUsernameMergeEnabled": false
}
'
-----
Get Access Management settings
get
https://yourServer.jamfcloud.com/api/v4/enrollment/access-management


Get Access Management settings

Response

200
Successful response

Response body
object
automatedDeviceEnrollmentServerUuid
string
length ≤ 256
-----

Configure Access Management settings
post
https://yourServer.jamfcloud.com/api/v4/enrollment/access-management


Configure Access Management settings

Body Params
Configure Access Management settings

automatedDeviceEnrollmentServerUuid
string
length ≤ 256
Response

200
Successfully Configured Access Management settings

Response body
object
automatedDeviceEnrollmentServerUuid
string
length ≤ 256
Updated 13 days ago

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v4/enrollment/access-management \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "automatedDeviceEnrollmentServerUuid": "4B637BAB65D14E6DA63A74E4F6F82C4B"
}
-----