Retrieve the Status Items from the latest Status Report for a device
get
https://yourServer.jamfcloud.com/api/v1/ddm/{clientManagementId}/status-items


Retrieves the Status Items from the latest Status Report for a device

Path Params
clientManagementId
string
required
client management id of the target device.

1
Responses

200
Success

Response body
object
statusItems
array of objects
Defaults to
object
key
string
The StatusItem key

value
string
The StatusItem value

lastUpdateTime
string
The local server time when the StatusItem was last updated

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ddm/1/status-items \
     --header 'accept: application/json'

{
  "statusItems": [
    {
      "key": "device.identifier.udid",
      "value": "00008103-000C51E402E2001E",
      "lastUpdateTime": "2024-08-25T21:09:31"
    }
  ]
}
-----
Retrieve a Status Item from the latest Status Report for a device
get
https://yourServer.jamfcloud.com/api/v1/ddm/{clientManagementId}/status-items/{key}


Retrieves a Status Item from the latest Status Report for a device

Path Params
clientManagementId
string
required
client management id of the target device.

123
key
string
required
the status item key to retrieve.

1212
Responses

200
Success

Response body
object
key
string
The StatusItem key

value
string
The StatusItem value

lastUpdateTime
string
The local server time when the StatusItem was last updated

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ddm/123/status-items/1212 \
     --header 'accept: application/json'

{
  "key": "device.identifier.udid",
  "value": "00008103-000C51E402E2001E",
  "lastUpdateTime": "2024-08-25T21:09:31"
}
-----
Force a device DDM sync
post
https://yourServer.jamfcloud.com/api/v1/ddm/{clientManagementId}/sync


Force a device to sync by queuing a new DeclarativeManagementCommand

Path Params
clientManagementId
string
required
The client management id of the target device.

123
Responses
204
The device was successfully synced.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/ddm/123/sync \
     --header 'accept: application/json'
-----
Retrieve an existing declaration
get
https://yourServer.jamfcloud.com/api/v1/dss-declarations/{declarationId}


Retrieves a stored declaration based on the provided declaration id

Path Params
declarationId
string
required
Declaration UUID

538F90D7-9383-4A2B-B1C8-81E845A9CFD7
Responses

200
The Declaration request was sent successfully.

Response body
object
declarations
array of objects
Defaults to
object
uuid
string
payloadJson
string | null
type
string | null
group
string | null
enum
ACTIVATION ASSET CONFIGURATION MANAGEMENT

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/dss-declarations/538F90D7-9383-4A2B-B1C8-81E845A9CFD7 \
     --header 'accept: application/json'

{
  "declarations": [
    {
      "uuid": "72676372-af55-432f-acd8-12984522e472",
      "payloadJson": {},
      "type": "com.apple.configuration.management.status-subscriptions",
      "group": "activation"
    }
  ]
}