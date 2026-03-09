Return information about the currently supported Time Zones
get
https://yourServer.jamfcloud.com/api/v1/time-zones

Returns information about the currently supported time zones

Response

200
Successful response

Response body
array of objects
object
zoneId
string
region
string
enum
Africa America Asia Atlantic Australia Europe Indian Pacific None

displayName
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/time-zones \
     --header 'accept: application/json'

[
  {
    "zoneId": "America/Chicago",
    "region": "America",
    "displayName": "Chicago - CT (-0500)"
  }
]