Get Re-enrollment object
get
https://yourServer.jamfcloud.com/api/v1/reenrollment

Gets Re-enrollment object

Responses

200
Details of Re-enrollment object were found

Response body
object
isFlushPolicyHistoryEnabled
boolean
Defaults to false
isFlushLocationInformationEnabled
boolean
Defaults to false
isFlushLocationInformationHistoryEnabled
boolean
Defaults to false
isFlushExtensionAttributesEnabled
boolean
Defaults to false
isFlushSoftwareUpdatePlansEnabled
boolean
Defaults to false
flushMDMQueue
string
enum
required
DELETE_NOTHING DELETE_ERRORS DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED DELETE_EVERYTHING

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/reenrollment \
     --header 'accept: application/json'

{
  "isFlushPolicyHistoryEnabled": false,
  "isFlushLocationInformationEnabled": false,
  "isFlushLocationInformationHistoryEnabled": false,
  "isFlushExtensionAttributesEnabled": false,
  "isFlushSoftwareUpdatePlansEnabled": false,
  "flushMDMQueue": "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"
}
-----
Update the Re-enrollment object
put
https://yourServer.jamfcloud.com/api/v1/reenrollment

Update the Re-enrollment object

Body Params
Re-enrollment object to update

isFlushPolicyHistoryEnabled
boolean
Defaults to false

false
isFlushLocationInformationEnabled
boolean
Defaults to false

false
isFlushLocationInformationHistoryEnabled
boolean
Defaults to false

false
isFlushExtensionAttributesEnabled
boolean
Defaults to false

false
isFlushSoftwareUpdatePlansEnabled
boolean
Defaults to false

false
flushMDMQueue
string
enum
required

DELETE_NOTHING
Allowed:

DELETE_NOTHING

DELETE_ERRORS

DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED

DELETE_EVERYTHING
Response

201
Re-enrollment record was updated

Response body
object
isFlushPolicyHistoryEnabled
boolean
Defaults to false
isFlushLocationInformationEnabled
boolean
Defaults to false
isFlushLocationInformationHistoryEnabled
boolean
Defaults to false
isFlushExtensionAttributesEnabled
boolean
Defaults to false
isFlushSoftwareUpdatePlansEnabled
boolean
Defaults to false
flushMDMQueue
string
enum
required
DELETE_NOTHING DELETE_ERRORS DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED DELETE_EVERYTHING

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/reenrollment \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "isFlushPolicyHistoryEnabled": false,
  "isFlushLocationInformationEnabled": false,
  "isFlushLocationInformationHistoryEnabled": false,
  "isFlushExtensionAttributesEnabled": false,
  "isFlushSoftwareUpdatePlansEnabled": false,
  "flushMDMQueue": "DELETE_NOTHING"
}
'

{
  "isFlushPolicyHistoryEnabled": false,
  "isFlushLocationInformationEnabled": false,
  "isFlushLocationInformationHistoryEnabled": false,
  "isFlushExtensionAttributesEnabled": false,
  "isFlushSoftwareUpdatePlansEnabled": false,
  "flushMDMQueue": "DELETE_EVERYTHING_EXCEPT_ACKNOWLEDGED"
}
-----
Get Re-enrollment history object
get
https://yourServer.jamfcloud.com/api/v1/reenrollment/history

Gets Re-enrollment history object

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
string
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc

date:desc
Show Deprecated
Response

200
Details of re-enrollment history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/reenrollment/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Re-enrollment history object notes
post
https://yourServer.jamfcloud.com/api/v1/reenrollment/history

Adds specified Re-enrollment history object notes

Body Params
history notes to create

note
string
required
Responses

201
Notes of re-enrollment history were added

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
     --url https://yourserver.jamfcloud.com/api/v1/reenrollment/history \
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
Export reenrollment history collection
post
https://yourServer.jamfcloud.com/api/v1/reenrollment/history/export


Export reenrollment history collection

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
     --url 'https://yourserver.jamfcloud.com/api/v1/reenrollment/history/export?page=0&page-size=100&sort=id%3Aasc' \
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
