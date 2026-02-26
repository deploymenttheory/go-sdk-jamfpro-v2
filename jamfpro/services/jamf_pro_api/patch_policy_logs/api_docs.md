Retrieve Patch Policy Logs
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs


Retrieves Patch Policy Logs

Path Params
id
string
required
patch policy id

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
Defaults to deviceName:asc
Sorting criteria in the format: property:asc/desc. Default sort is deviceName:asc. Multiple sort criteria are supported and must be separated with a comma.


string

deviceName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter Patch Policy Logs collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: deviceId, deviceName, statusCode, statusDate, attemptNumber, ignoredForPatchPolicyId. This param can be combined with paging and sorting.

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
patchPolicyId
string
deviceName
string
deviceId
string
statusCode
integer
statusDate
date-time
statusEnum
string
enum
UNKNOWN PENDING COMPLETED FAILED

attemptNumber
integer
ignoredForPatchPolicyId
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/patch-policies/1/logs?page=0&page-size=100&sort=deviceName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "patchPolicyId": "1",
      "deviceName": "Admins Macbook",
      "deviceId": "1",
      "statusCode": 1,
      "statusDate": "2019-02-04T21:09:31.661Z",
      "statusEnum": "COMPLETED",
      "attemptNumber": 1,
      "ignoredForPatchPolicyId": "1"
    }
  ]
}
-----
Return the count of the Patch Policy Logs for the patch policy id that are eligible for a retry attempt
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs/eligible-retry-count


return the count of the patch policy logs for the patch policy id that are eligible for a retry attempt

Path Params
id
string
required
patch policy id

Response

200
Number of patch policy logs found

Response body
object
count
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//logs/eligible-retry-count \
     --header 'accept: application/json'

{
  "count": 5
}
-----
Send retry attempts for specific devices
post
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs/retry


Send retry attempts for specific devices

Path Params
id
string
required
patch policy id

1
Body Params
deviceIds
array of strings

string

1

ADD string
Response

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies/1/logs/retry \
     --header 'content-type: application/json' \
     --data '{"deviceIds":["1"]}'
-----
Send retry attempts for all devices
post
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs/retry-all


Send retry attempts for all devices

Path Params
id
string
required
patch policy id

Responses
202
OK

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//logs/retry-all
-----
Retrieves a single Patch Policy Log
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs/{deviceId}


Retrieves a single Patch Policy Log

Path Params
id
string
required
patch policy id

deviceId
string
required
device id

Responses

200
Patch Policy Log

Response body
object
patchPolicyId
string
deviceName
string
deviceId
string
statusCode
integer
statusDate
date-time
statusEnum
string
enum
UNKNOWN PENDING COMPLETED FAILED

attemptNumber
integer
ignoredForPatchPolicyId
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//logs/ \
     --header 'accept: application/json'

{
  "patchPolicyId": "1",
  "deviceName": "Admins Macbook",
  "deviceId": "1",
  "statusCode": 1,
  "statusDate": "2019-02-04T21:09:31.661Z",
  "statusEnum": "COMPLETED",
  "attemptNumber": 1,
  "ignoredForPatchPolicyId": "1"
}
-----
eturn attempt details for a specific log
get
https://yourServer.jamfcloud.com/api/v2/patch-policies/{id}/logs/{deviceId}/details


Return attempt details for a specific log

Path Params
id
string
required
patch policy id

deviceId
string
required
device id

Responses

200
Attempt details

Response body
array of objects
object
id
string
attemptNumber
integer
deviceId
string
actions
array of objects
object
id
string
actionOrder
integer
action
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/patch-policies//logs//details \
     --header 'accept: application/json'

[
  {
    "id": "1",
    "attemptNumber": 1,
    "deviceId": "1",
    "actions": [
      {
        "id": "1",
        "actionOrder": 1,
        "action": "Installing..."
      }
    ]
  }
]
-----
