Gets session history items.
get
https://yourServer.jamfcloud.com/api/v1/jamf-remote-assist/session


Returns tenants sessions history.

Response

200
Up to 100 latest session history items

Response body
array of objects
object
tenantId
string
sessionId
string
deviceId
string
sessionStartedTimestamp
date-time
sessionEndedTimestamp
date-time
sessionType
string
enum
ATTENDED UNATTENDED

statusType
string
enum
STARTED FINISHED ERROR

sessionAdminId
string
comment
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-remote-assist/session \
     --header 'accept: application/json'

[
  {
    "tenantId": "ee5d0ffae49e675d32ba5bdb2ce25182e958b979afc8d210c924d725bbf45f8b",
    "sessionId": "c4087bc8-cb48-4452-9bec-a3c50088a485",
    "deviceId": "aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61866",
    "sessionStartedTimestamp": "2023-11-07T10:25:27.49Z",
    "sessionEndedTimestamp": "2023-11-07T10:25:27.49Z",
    "sessionType": "ATTENDED",
    "statusType": "STARTED",
    "sessionAdminId": "sessionAdminId1",
    "comment": "Example comment"
  }
]

-----

Gets single session history item.
get
https://yourServer.jamfcloud.com/api/v1/jamf-remote-assist/session/{id}


Returns tenants session history for specific session.

Path Params
id
string
required
instance id of session

Responses

200
Single session history item

Response body
object
tenantId
string
sessionId
string
deviceId
string
sessionStartedTimestamp
date-time
sessionEndedTimestamp
date-time
sessionType
string
enum
ATTENDED UNATTENDED

statusType
string
enum
STARTED FINISHED ERROR

sessionAdminId
string
comment
string
details
object
fileTransferItemList
array of objects
object
filePath
string
transferTimestamp
date-time
fileTransferType
string
enum
DOWNLOAD UPLOAD

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-remote-assist/session/ \
     --header 'accept: application/json'

{
  "tenantId": "ee5d0ffae49e675d32ba5bdb2ce25182e958b979afc8d210c924d725bbf45f8b",
  "sessionId": "c4087bc8-cb48-4452-9bec-a3c50088a485",
  "deviceId": "aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61866",
  "sessionStartedTimestamp": "2023-11-07T10:25:27.49Z",
  "sessionEndedTimestamp": "2023-11-07T10:25:27.49Z",
  "sessionType": "ATTENDED",
  "statusType": "STARTED",
  "sessionAdminId": "sessionAdminId1",
  "comment": "Example comment",
  "details": {
    "fileTransferItemList": [
      {
        "filePath": "/SomeFilePath/ToFile.xml",
        "transferTimestamp": "2023-11-07T10:25:27.49Z",
        "fileTransferType": "DOWNLOAD"
      }
    ]
  }
}
-----

Gets session history items.
get
https://yourServer.jamfcloud.com/api/v2/jamf-remote-assist/session


Returns tenants sessions history.

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
Defaults to sessionId:desc
Sorting criteria in the format: property:asc/desc. Default sort is sessionId:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=sessionId:desc,deviceId:asc


string

sessionId:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter session history items collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: sessionId, deviceId, sessionAdminId. This param can be combined with paging and sorting. Example: sessionAdminId=="Andrzej"

Response

200
Up to 100 latest session history items

Response body
object
results
array of objects
required
object
tenantId
string
sessionId
string
deviceId
string
sessionStartedTimestamp
date-time
sessionEndedTimestamp
date-time
sessionType
string
enum
ATTENDED UNATTENDED

statusType
string
enum
STARTED FINISHED ERROR

sessionAdminId
string
comment
string
totalCount
integer
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/jamf-remote-assist/session?page=0&page-size=100&sort=sessionId%3Adesc' \
     --header 'accept: application/json'

{
  "results": [
    {
      "tenantId": "ee5d0ffae49e675d32ba5bdb2ce25182e958b979afc8d210c924d725bbf45f8b",
      "sessionId": "c4087bc8-cb48-4452-9bec-a3c50088a485",
      "deviceId": "aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61866",
      "sessionStartedTimestamp": "2023-11-07T10:25:27.49Z",
      "sessionEndedTimestamp": "2023-11-07T10:25:27.49Z",
      "sessionType": "ATTENDED",
      "statusType": "STARTED",
      "sessionAdminId": "sessionAdminId1",
      "comment": "Example comment"
    }
  ],
  "totalCount": 100
}
-----
Export Jamf Remote Assist sessions history
post
https://yourServer.jamfcloud.com/api/v2/jamf-remote-assist/session/export


Export Jamf Remote Assist sessions history

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
     --url https://yourserver.jamfcloud.com/api/v2/jamf-remote-assist/session/export \
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

sessionId,deviceId
c4087bc8-cb48-4452-9bec-a3c50088a485,aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61866
c4087bc8-cb48-4452-9bec-a3c50088a486,aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61867
c4087bc8-cb48-4452-9bec-a3c50088a487,aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61868
-----

Gets single session history item.
get
https://yourServer.jamfcloud.com/api/v2/jamf-remote-assist/session/{id}


Returns tenants session history for specific session.

Path Params
id
string
required
instance id of session

Responses

200
Single session history item

Response body
object
tenantId
string
sessionId
string
deviceId
string
sessionStartedTimestamp
date-time
sessionEndedTimestamp
date-time
sessionType
string
enum
ATTENDED UNATTENDED

statusType
string
enum
STARTED FINISHED ERROR

sessionAdminId
string
comment
string
details
object
fileTransferItemList
array of objects
object
filePath
string
transferTimestamp
date-time
fileTransferType
string
enum
DOWNLOAD UPLOAD

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/jamf-remote-assist/session/ \
     --header 'accept: application/json'

{
  "tenantId": "ee5d0ffae49e675d32ba5bdb2ce25182e958b979afc8d210c924d725bbf45f8b",
  "sessionId": "c4087bc8-cb48-4452-9bec-a3c50088a485",
  "deviceId": "aeafd2ff95c7e0b937a8dbf3121a8dcab7a70161cd4493af038fa9c784b61866",
  "sessionStartedTimestamp": "2023-11-07T10:25:27.49Z",
  "sessionEndedTimestamp": "2023-11-07T10:25:27.49Z",
  "sessionType": "ATTENDED",
  "statusType": "STARTED",
  "sessionAdminId": "sessionAdminId1",
  "comment": "Example comment",
  "details": {
    "fileTransferItemList": [
      {
        "filePath": "/SomeFilePath/ToFile.xml",
        "transferTimestamp": "2023-11-07T10:25:27.49Z",
        "fileTransferType": "DOWNLOAD"
      }
    ]
  }
}