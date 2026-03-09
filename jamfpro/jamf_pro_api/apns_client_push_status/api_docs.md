Search for clients with push notifications disabled
get
https://yourServer.jamfcloud.com/api/v1/apns-client-push-status

Retrieve a paginated, sortable, and filterable list of MDM clients that have push notifications disabled. The endpoint queries the mdm_client table and returns information about when push was disabled and links to the device records.

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
Defaults to pushDisabledTime:asc
Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported. Sortable fields: pushDisabledTime, deviceType, managementId


string

pushDisabledTime:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter results. Fields allowed in the query: deviceType, disabledAt, managementId. This param can be combined with paging and sorting.

Example: filter=deviceType=="MOBILE_DEVICE" Example: filter=disabledAt>2024-11-01T00:00:00Z Example: filter=deviceType=="COMPUTER";disabledAt>2024-01-01T00:00:00Z

Responses

200
Successful operation. Returns a paginated list of clients with push notifications disabled.

Response body
object
totalCount
int64
required
Total number of records matching the query

results
array of objects
required
Array of APNS client push status records

object
deviceType
string
enum
The type of MDM client device

MOBILE_DEVICE MOBILE_DEVICE_USER COMPUTER COMPUTER_USER TV WATCH VISION_PRO UNKNOWN

clientId
string
Id of the Computer or Device record in Jamf Pro

disabledAt
date-time
Timestamp when push notifications were disabled for this client (ISO-8601 format)

managementId
string
Unique identifier for the device management record

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/apns-client-push-status?page=0&page-size=100&sort=pushDisabledTime%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 42,
  "results": [
    {
      "deviceType": "MOBILE_DEVICE",
      "clientId": "55",
      "disabledAt": "2024-11-06T14:30:00Z",
      "managementId": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
    }
  ]
}

-----

Enable push notifications for all clients
post
https://yourServer.jamfcloud.com/api/v1/apns-client-push-status/enable-all-clients


Create a request to enable push notifications for all MDM clients that currently have push disabled. This is an asynchronous operation that processes all disabled clients in the background.

Response

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/apns-client-push-status/enable-all-clients


-----

Get status of enable all clients request
get
https://yourServer.jamfcloud.com/api/v1/apns-client-push-status/enable-all-clients/status


Retrieve the status of the most recent request to enable push notifications for all clients. Returns 404 if no recent request exists.

Responses

200
Successful operation. Returns the status of the recent request.

Response body
object
requestedTime
date-time
Timestamp when the request was created (ISO-8601 format)

status
string
enum
Current status of the request

QUEUED STARTED COMPLETED

processedTime
date-time | null
Timestamp when the request was processed (ISO-8601 format), null if not yet processed

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/apns-client-push-status/enable-all-clients/status \
     --header 'accept: application/json'

{
  "requestedTime": "2024-11-06T14:30:00Z",
  "status": "QUEUED",
  "processedTime": "2024-11-06T15:30:00Z"
}

-----

Enable push notifications for a single client
post
https://yourServer.jamfcloud.com/api/v1/apns-client-push-status/enable-client


Enable push notifications for a single MDM client that previously had push disabled. This sets the pushEnabled flag to true for the specified client. managementId field is required in the request body.

Body Params
managementId
string
required
Unique identifier for the device management record to enable push for

Responses
204
Push notifications successfully enabled for the client. No content returned.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/apns-client-push-status/enable-client \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "managementId": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
'