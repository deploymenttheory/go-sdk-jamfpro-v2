Get the current onboarding settings configuration.
get
https://yourServer.jamfcloud.com/api/v1/onboarding

Get the current onboarding settings configuration.

Response

200
Success

Response body
object
id
string
enabled
boolean
required
onboardingItems
array of objects
required
object
id
string | null
entityId
string
required
The id of the Jamf Pro object that should be added to the onboarding workflow for end users. Use this in conjunction with the selfServiceEntityType. For example, if the policy with id 132 should be added to onboarding, then entityId should be 132 and selfServiceEntityType should be OS_X_POLICY.

entityName
string
scopeDescription
string
siteDescription
string
selfServiceEntityType
string
enum
required
OS_X_POLICY OS_X_CONFIG_PROFILE OS_X_MAC_APP OS_X_APP_INSTALLER OS_X_EBOOK OS_X_PATCH_POLICY UNKNOWN

priority
integer
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/onboarding \
     --header 'accept: application/json'

{
  "id": "1",
  "enabled": true,
  "onboardingItems": [
    {
      "id": "3",
      "entityId": "1",
      "entityName": "Policy 1",
      "scopeDescription": "All Computers",
      "siteDescription": "None",
      "selfServiceEntityType": "OS_X_POLICY",
      "priority": 35
    }
  ]
}
-----
Update the onboarding configuration.
put
https://yourServer.jamfcloud.com/api/v1/onboarding

Update the onboarding configuration.

Body Params
Onboarding settings to save.

enabled
boolean
required

true
onboardingItems
array of objects
required

object

id
string | null
3
entityId
string
required
The id of the Jamf Pro object that should be added to the onboarding workflow for end users. Use this in conjunction with the selfServiceEntityType. For example, if the policy with id 132 should be added to onboarding, then entityId should be 132 and selfServiceEntityType should be OS_X_POLICY.

1
selfServiceEntityType
string
enum
required

OS_X_POLICY
Allowed:

OS_X_POLICY

OS_X_CONFIG_PROFILE

OS_X_MAC_APP

OS_X_APP_INSTALLER

OS_X_EBOOK

OS_X_PATCH_POLICY

UNKNOWN
priority
integer
required
35

object

id
string | null
3
entityId
string
required
The id of the Jamf Pro object that should be added to the onboarding workflow for end users. Use this in conjunction with the selfServiceEntityType. For example, if the policy with id 132 should be added to onboarding, then entityId should be 132 and selfServiceEntityType should be OS_X_POLICY.

1
selfServiceEntityType
string
enum
required

OS_X_MAC_APP
Allowed:

OS_X_POLICY

OS_X_CONFIG_PROFILE

OS_X_MAC_APP

OS_X_APP_INSTALLER

OS_X_EBOOK

OS_X_PATCH_POLICY

UNKNOWN
priority
integer
required
35

ADD object
Responses

200
Successfully updated

Response body
object
id
string
enabled
boolean
required
onboardingItems
array of objects
required
object
id
string | null
entityId
string
required
The id of the Jamf Pro object that should be added to the onboarding workflow for end users. Use this in conjunction with the selfServiceEntityType. For example, if the policy with id 132 should be added to onboarding, then entityId should be 132 and selfServiceEntityType should be OS_X_POLICY.

entityName
string
scopeDescription
string
siteDescription
string
selfServiceEntityType
string
enum
required
OS_X_POLICY OS_X_CONFIG_PROFILE OS_X_MAC_APP OS_X_APP_INSTALLER OS_X_EBOOK OS_X_PATCH_POLICY UNKNOWN

priority
integer
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/onboarding \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "enabled": true,
  "onboardingItems": [
    {
      "selfServiceEntityType": "OS_X_POLICY",
      "id": "3",
      "entityId": "1",
      "priority": 35
    },
    {
      "selfServiceEntityType": "OS_X_MAC_APP",
      "id": "3",
      "entityId": "1",
      "priority": 35
    }
  ]
}
'

{
  "id": "1",
  "enabled": true,
  "onboardingItems": [
    {
      "id": "3",
      "entityId": "1",
      "entityName": "Policy 1",
      "scopeDescription": "All Computers",
      "siteDescription": "None",
      "selfServiceEntityType": "OS_X_POLICY",
      "priority": 35
    }
  ]
}
-----
Retrieves a list of applications that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-apps

Retrieves a list of applications that are eligible to be used in an onboarding configuration

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
OK

Response body
object
totalCount
integer
results
array of objects
object
id
string
required
name
string
required
scopeDescription
string
required
siteDescription
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/onboarding/eligible-apps?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "Item 1",
      "scopeDescription": "All Computers",
      "siteDescription": "Example Site"
    }
  ]
}
-----
Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-configuration-profiles

Retrieves a list of configuration profiles that are eligible to be used in an onboarding configuration

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
OK

Response body
object
totalCount
integer
results
array of objects
object
id
string
required
name
string
required
scopeDescription
string
required
siteDescription
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/onboarding/eligible-configuration-profiles?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "Item 1",
      "scopeDescription": "All Computers",
      "siteDescription": "Example Site"
    }
  ]
}
-----
Retrieves a list of policies that are eligible to be used in an onboarding configuration
get
https://yourServer.jamfcloud.com/api/v1/onboarding/eligible-policies

Retrieves a list of policies that are eligible to be used in an onboarding configuration

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
OK

Response body
object
totalCount
integer
results
array of objects
object
id
string
required
name
string
required
scopeDescription
string
required
siteDescription
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/onboarding/eligible-policies?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "Item 1",
      "scopeDescription": "All Computers",
      "siteDescription": "Example Site"
    }
  ]
}
-----
Get Onboarding history object
get
https://yourServer.jamfcloud.com/api/v1/onboarding/history

Gets Onboarding history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and date<2019-12-15

Response

200
Details of onboarding history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/onboarding/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Onboarding history object notes
post
https://yourServer.jamfcloud.com/api/v1/onboarding/history

Adds Onboarding history object notes

Body Params
history notes to create

note
string
required
Responses

201
Notes of onboarding history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/onboarding/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Export history object collection in specified format for Onboarding
post
https://yourServer.jamfcloud.com/api/v1/onboarding/history/export

Export history object collection in specified format for Onboarding

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
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and date<2019-12-15

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
Defaults to application/json
Generated from available response content types


text/csv
Allowed:

application/json

text/csv
Responses

200
Export successful

Response body
json

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/onboarding/history/export?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: text/csv' \
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
