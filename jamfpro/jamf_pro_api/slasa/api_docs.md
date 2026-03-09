Get the status of SLASA
get
https://yourServer.jamfcloud.com/api/v1/slasa

Get if SLASA has been accepted or not

Response

200
Whether SLASA has been accepted or not

Response body
object
slasaAcceptanceStatus
string
enum
ACCEPTED NOT_ACCEPTED

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/slasa \
     --header 'accept: application/json'

{
  "slasaAcceptanceStatus": "ACCEPTED"
}
-----

Accept the SLASA
post
https://yourServer.jamfcloud.com/api/v1/slasa

Accept the SLASA for Jamf Pro.

Response
204
SLASA has been accepted

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/slasa