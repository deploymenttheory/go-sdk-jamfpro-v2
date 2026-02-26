Get the current Jamf Parent app settings
get
https://yourServer.jamfcloud.com/api/v1/parent-app

Get the current Jamf Parent app settings

Response

200
Details of the current Jamf Parent app settings.

Response body
object
timezoneId
string
required
restrictedTimes
object
required
key
string
enum
MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY

object

newKey object
beginTime
string
endTime
string
object

newKey-1 object
object

newKey-2 object

View Additional Properties
deviceGroupId
integer
required
isEnabled
boolean
required
allowTemplates
boolean
disassociateOnWipeAndReEnroll
boolean
allowClearPasscode
boolean
safelistedApps
array of objects
object
name
string
bundleId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/parent-app \
     --header 'accept: application/json'

{
  "timezoneId": "Europe/Paris",
  "restrictedTimes": {
    "SUNDAY": {
      "beginTime": "08:30:00",
      "endTime": "15:45:00"
    }
  },
  "deviceGroupId": 1,
  "isEnabled": true,
  "allowTemplates": true,
  "disassociateOnWipeAndReEnroll": true,
  "allowClearPasscode": true,
  "safelistedApps": [
    {
      "name": "Content Filter",
      "bundleId": "com.jamf.parent"
    }
  ]
}
-----
Update Jamf Parent app settings
put
https://yourServer.jamfcloud.com/api/v1/parent-app

Update Jamf Parent app settings

Body Params
Jamf Parent app settings to save.

timezoneId
string
required
Europe/Paris
restrictedTimes
object
required

restrictedTimes object
key
string
enum

Allowed:

MONDAY

TUESDAY

WEDNESDAY

THURSDAY

FRIDAY

SATURDAY

SUNDAY

Add Field
deviceGroupId
integer
required
1
isEnabled
boolean
required

true
allowTemplates
boolean

true
disassociateOnWipeAndReEnroll
boolean

true
allowClearPasscode
boolean

true
safelistedApps
array of objects

object

name
string
Content Filter
bundleId
string
com.jamf.parent

ADD object
Responses

200
Jamf Parent app settings updated

Response body
object
timezoneId
string
required
restrictedTimes
object
required
key
string
enum
MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY SUNDAY


View Additional Properties
deviceGroupId
integer
required
isEnabled
boolean
required
allowTemplates
boolean
disassociateOnWipeAndReEnroll
boolean
allowClearPasscode
boolean
safelistedApps
array of objects
object
name
string
bundleId
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/parent-app \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "isEnabled": true,
  "timezoneId": "Europe/Paris",
  "deviceGroupId": 1,
  "allowTemplates": true,
  "disassociateOnWipeAndReEnroll": true,
  "allowClearPasscode": true,
  "safelistedApps": [
    {
      "name": "Content Filter",
      "bundleId": "com.jamf.parent"
    }
  ]
}
'
{
  "timezoneId": "Europe/Paris",
  "restrictedTimes": {
    "SUNDAY": {
      "beginTime": "08:30:00",
      "endTime": "15:45:00"
    }
  },
  "deviceGroupId": 1,
  "isEnabled": true,
  "allowTemplates": true,
  "disassociateOnWipeAndReEnroll": true,
  "allowClearPasscode": true,
  "safelistedApps": [
    {
      "name": "Content Filter",
      "bundleId": "com.jamf.parent"
    }
  ]
}
-----
Get Jamf Parent app settings history
get
https://yourServer.jamfcloud.com/api/v1/parent-app/history

Gets Jamf Parent app settings history

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
filter
string
Query in the RSQL format, allowing to filter results. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: status, updated, version This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

sort
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

date:desc
Response

200
Details of Jamf Parent app settings history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/parent-app/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Jamf Parent app settings history notes
post
https://yourServer.jamfcloud.com/api/v1/parent-app/history

Adds Jamf Parent app settings history notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes to Jamf Parent app settings history were added

Response body
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

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/parent-app/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}
-----