Retrieve Patch Software Title Configurations
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations


Retrieves patch software title configurations

Response

200
Success

Response body
array of objects
object
displayName
string
required
categoryId
string
Defaults to -1
siteId
string
Defaults to -1
uiNotifications
boolean
Defaults to false
emailNotifications
boolean
Defaults to false
softwareTitleId
string
required
jamfOfficial
boolean
extensionAttributes
array of objects
object
accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.

eaId
string
softwareTitleName
string
softwareTitleNameId
string
softwareTitlePublisher
string
patchSourceName
string
patchSourceEnabled
boolean
id
string
packages
array of objects
object
packageId
string
version
string
displayName
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations \
     --header 'accept: application/json'

[
  {
    "displayName": "Google Chrome",
    "categoryId": "-1",
    "siteId": "-1",
    "uiNotifications": false,
    "emailNotifications": false,
    "softwareTitleId": "1",
    "jamfOfficial": false,
    "extensionAttributes": [
      {
        "accepted": false,
        "eaId": "google-chrome-ea"
      }
    ],
    "softwareTitleName": "Safari",
    "softwareTitleNameId": "AppleSafari",
    "softwareTitlePublisher": "Apple",
    "patchSourceName": "Jamf",
    "patchSourceEnabled": true,
    "id": "1",
    "packages": [
      {
        "packageId": "1",
        "version": "1",
        "displayName": "Firefox.pkg"
      }
    ]
  }
]
-----
Create Patch Software Title Configurations
post
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations


Creates Patch Software Title Configurations using sToken

Body Params
Software title configurations to create

displayName
string
required
Google Chrome
categoryId
string
Defaults to -1
-1
siteId
string
Defaults to -1
-1
uiNotifications
boolean
Defaults to false

false
emailNotifications
boolean
Defaults to false

false
softwareTitleId
string
required
extensionAttributes
array of objects

object

accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.


false
eaId
string
google-chrome-ea

object

accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.


true
eaId
string
google-chrome-ea

ADD object
Responses

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "categoryId": "-1",
  "siteId": "-1",
  "uiNotifications": false,
  "emailNotifications": false,
  "displayName": "Google Chrome",
  "extensionAttributes": [
    {
      "accepted": false,
      "eaId": "google-chrome-ea"
    },
    {
      "accepted": true,
      "eaId": "google-chrome-ea"
    }
  ]
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve Patch Software Title Configurations with the supplied id
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}


Retrieves Patch Software Title Configurations with the supplied id

Path Params
id
string
required
Patch Software Title Configurations identifier

1
Responses

200
Success

Response body
object
displayName
string
required
categoryId
string
Defaults to -1
siteId
string
Defaults to -1
uiNotifications
boolean
Defaults to false
emailNotifications
boolean
Defaults to false
softwareTitleId
string
required
jamfOfficial
boolean
extensionAttributes
array of objects
object
accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.

eaId
string
softwareTitleName
string
softwareTitleNameId
string
softwareTitlePublisher
string
patchSourceName
string
patchSourceEnabled
boolean
id
string
packages
array of objects
object
packageId
string
version
string
displayName
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1 \
     --header 'accept: application/json'

{
  "displayName": "Google Chrome",
  "categoryId": "-1",
  "siteId": "-1",
  "uiNotifications": false,
  "emailNotifications": false,
  "softwareTitleId": "1",
  "jamfOfficial": false,
  "extensionAttributes": [
    {
      "accepted": false,
      "eaId": "google-chrome-ea"
    }
  ],
  "softwareTitleName": "Safari",
  "softwareTitleNameId": "AppleSafari",
  "softwareTitlePublisher": "Apple",
  "patchSourceName": "Jamf",
  "patchSourceEnabled": true,
  "id": "1",
  "packages": [
    {
      "packageId": "1",
      "version": "1",
      "displayName": "Firefox.pkg"
    }
  ]
}
-----
Delete Patch Software Title Configurations with the supplied id
delete
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}


Deletes Patch Software Title Configurations with the supplied id

Path Params
id
string
required
Patch Software Title Configurations identifier

Responses
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/ \
     --header 'accept: application/json'
-----
Update Patch Software Title Configurations
patch
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}


Updates Patch Software Title Configurations

Path Params
id
string
required
Patch Software Title Configurations identifier

1
Body Params
Patch Software Title Configurations to update

displayName
string
Google Chrome
categoryId
string
-1
siteId
string
-1
uiNotifications
boolean

true
emailNotifications
boolean

true
softwareTitleId
string
1
packages
array of objects

object

packageId
string
1
version
string
1

ADD object
extensionAttributes
array of objects

object

accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.


false
eaId
string
google-chrome-ea

ADD object
Responses

200
Success

Response body
object
displayName
string
required
categoryId
string
Defaults to -1
siteId
string
Defaults to -1
uiNotifications
boolean
Defaults to false
emailNotifications
boolean
Defaults to false
softwareTitleId
string
required
jamfOfficial
boolean
extensionAttributes
array of objects
object
accepted
boolean
Defaults to false
Once an extension attribute is accepted, it cannot be reverted.

eaId
string
softwareTitleName
string
softwareTitleNameId
string
softwareTitlePublisher
string
patchSourceName
string
patchSourceEnabled
boolean
id
string
packages
array of objects
object
packageId
string
version
string
displayName
string

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json' \
     --data '
{
  "displayName": "Google Chrome",
  "categoryId": "-1",
  "siteId": "-1",
  "uiNotifications": true,
  "emailNotifications": true,
  "softwareTitleId": "1",
  "packages": [
    {
      "packageId": "1",
      "version": "1"
    }
  ],
  "extensionAttributes": [
    {
      "accepted": false,
      "eaId": "google-chrome-ea"
    }
  ]
}
'

{
  "displayName": "Google Chrome",
  "categoryId": "-1",
  "siteId": "-1",
  "uiNotifications": false,
  "emailNotifications": false,
  "softwareTitleId": "1",
  "jamfOfficial": false,
  "extensionAttributes": [
    {
      "accepted": false,
      "eaId": "google-chrome-ea"
    }
  ],
  "softwareTitleName": "Safari",
  "softwareTitleNameId": "AppleSafari",
  "softwareTitlePublisher": "Apple",
  "patchSourceName": "Jamf",
  "patchSourceEnabled": true,
  "id": "1",
  "packages": [
    {
      "packageId": "1",
      "version": "1",
      "displayName": "Firefox.pkg"
    }
  ]
}
-----
Return whether or not the requested software title configuration is on the dashboard
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/dashboard


Returns whether or not the requested software title configuration is on the dashboard

Path Params
id
string
required
software title configuration id

Responses

200
Whether the software title configuration is on the Dashboard.

Response body
object
onDashboard
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//dashboard \
     --header 'accept: application/json'

{
  "onDashboard": true
}
-----
Add a software title configuration to the dashboard
post
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/dashboard


Adds asoftware title configuration to the dashboard.

Path Params
id
string
required
software title configuration id

1
Response
204
OK

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1/dashboard
-----
Remove a software title configuration from the dashboard
delete
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/dashboard


Removes a software title configuration from the dashboard.

Path Params
id
string
required
software title configuration id

1
Response

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1/dashboard
-----
Retrieve Patch Software Title Definitions with the supplied id
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/definitions


Retrieves patch software title definitions with the supplied id

Path Params
id
string
required
Patch Software Title identifier

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
Defaults to absoluteOrderId:asc
Sorting criteria in the format: property:asc/desc. Default sort is absoluteOrderId:asc. Multiple sort criteria are supported and must be separated with a comma.


string

absoluteOrderId:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Software Title Definition collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, version, minimumOperatingSystem, releaseDate, reboot, standalone and absoluteOrderId. This param can be combined with paging and sorting.

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
version
string
minimumOperatingSystem
string
Defaults to -1
releaseDate
string
Defaults to -1
rebootRequired
boolean
Defaults to false
killApps
array of objects
object
appName
string
standalone
boolean
Defaults to false
absoluteOrderId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1/definitions?page=0&page-size=100&sort=absoluteOrderId%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "version": "10.37.0",
      "minimumOperatingSystem": "12.0.1",
      "releaseDate": "2010-12-10 13:36:04",
      "rebootRequired": false,
      "killApps": [
        {
          "appName": "Firefox"
        }
      ],
      "standalone": false,
      "absoluteOrderId": "1"
    }
  ]
}
-----
Retrieve list of Patch Software Title Configuration Dependencies
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/dependencies


Retrieve list of Patch Software Title Configuration Dependencies

Path Params
id
string
required
Patch Software Title Configuration Id

1
Response

200
Success

Response body
object
totalCount
integer
results
array of objects
object
smartGroupId
string
smartGroupName
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1/dependencies \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "smartGroupId": "1",
      "smartGroupName": "name"
    }
  ]
}
-----
Export Patch Reporting Data
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/export-report


Export Patch Reporting Data

Path Params
id
string
required
Patch Software Title Configurations identifier

Query Params
filter
string
Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting.

columns-to-export
array of strings
required
Defaults to computerName,deviceId,username,operatingSystemVersion,lastContactTime,buildingName,departmentName,siteName,version
List of column names to export


string

computerName

string

deviceId

string

username

string

operatingSystemVersion

string

lastContactTime

string

buildingName

string

departmentName

string

siteName

string

version

ADD string
Headers
accept
string
File

text/csv
Responses

200
Export successful

Response body

text/csv
json
-----
Retrieve Software Title Extension Attributes with the supplied id
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/extension-attributes


Retrieves software title extension attributes with the supplied id

Path Params
id
string
required
Patch Software Title identifier

Responses

200
Success

Response body
array of objects
object
eaId
string
accepted
boolean
displayName
string
scriptContents
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//extension-attributes \
     --header 'accept: application/json'

[
  {
    "eaId": "google-chrome-ea",
    "accepted": true,
    "displayName": "Google Chrome",
    "scriptContents": "ZXhhbXBsZSBvZiBhIGJhc2U2NCBlbmNvZGVkIHZhbGlkIHAxMi4ga2V5c3RvcmUgZmlsZQ=="
  }
]
-----
Get specified Patch Software Title Configuration history object
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/history


Gets specified Patch Software Title Configuration history object

Path Params
id
string
required
Patch Software Title Configuration Id

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma.


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of Patch Software Title Configuration history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations/1/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Patch Software Title Configuration history object notes
post
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/history


Adds Patch Software Title Configuration history object notes

Path Params
id
string
required
Patch Software Title Configuration Id

Body Params
History notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of Patch Software Title Configuration history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//history \
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
Retrieve Patch Software Title Configuration Patch Report
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/patch-report


Retrieve Patch Software Title Configuration Patch Report

Path Params
id
string
required
Patch Software Title Configurations identifier

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
Defaults to computerName:asc
Sorting criteria in the format: property:asc/desc. Default sort is computerName:asc. Multiple sort criteria are supported and must be separated with a comma. Supported fields: computerName, deviceId, username, operatingSystemVersion, lastContactTime, buildingName, departmentName, siteName, version


string

computerName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Report collection on version equality only. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: version. Comparators allowed in the query: ==, != This param can be combined with paging and sorting.

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
computerName
string
deviceId
string
username
string
operatingSystemVersion
string
lastContactTime
date-time
buildingName
string
departmentName
string
siteName
string
version
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//patch-report?page=0&page-size=100&sort=computerName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "computerName": "MacBook",
      "deviceId": "2",
      "username": "User",
      "operatingSystemVersion": "10.11",
      "lastContactTime": "1970-01-01T00:00:00Z",
      "buildingName": "Building",
      "departmentName": "Department",
      "siteName": "Site",
      "version": "10.1"
    }
  ]
}
-----
Return Active Patch Summary
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/patch-summary


Returns active patch summary.

Path Params
id
string
required
Patch id

Responses

200
Active Patch Summary

Response body
object
softwareTitleId
string
title
string
latestVersion
string
releaseDate
date-time
upToDate
integer
outOfDate
integer
onDashboard
boolean
softwareTitleConfigurationId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//patch-summary \
     --header 'accept: application/json'

{
  "softwareTitleId": "1",
  "title": "Patch title",
  "latestVersion": "2",
  "releaseDate": "2018-10-15T16:39:56.307Z",
  "upToDate": 3,
  "outOfDate": 6,
  "onDashboard": false,
  "softwareTitleConfigurationId": "1"
}
-----
Returns patch versions
get
https://yourServer.jamfcloud.com/api/v2/patch-software-title-configurations/{id}/patch-summary/versions


Returns patch versions

Path Params
id
string
required
Patch id

Responses

200
Patch versions

Response body
array of objects
object
absoluteOrderId
string
version
string
onVersion
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-software-title-configurations//patch-summary/versions \
     --header 'accept: application/json'

[
  {
    "absoluteOrderId": "1",
    "version": "3",
    "onVersion": 1
  }
]