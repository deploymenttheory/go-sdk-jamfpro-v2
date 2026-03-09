Return a list of groups for a device
get
https://yourServer.jamfcloud.com/api/v1/devices/{id}/groups

Returns a list of groups that the specified device belongs to.

Path Params
id
string
required
Device Platform ID

Responses

200
OK

Response body
array of objects
object
id
string
length ≥ 1
Group Platform ID

name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/devices//groups \
     --header 'accept: application/json'

[
  {
    "id": "699bc5b1-efb7-431a-b18f-b2ec5c435631",
    "name": "Test Group"
  }
]