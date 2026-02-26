Finds the Jamf Pro GSX Connection information
get
https://yourServer.jamfcloud.com/api/v1/gsx-connection

Finds the Jamf Pro GSX Connection information

Response

200
Success

Response body
object
enabled
boolean
required
Defaults to false
username
string
required
Defaults to
serviceAccountNo
string
required
length ≤ 10
shipToNo
string
length ≤ 10
gsxKeystore
object
required
name
string
required
Defaults to
expirationEpoch
int64
errorMessage
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection \
     --header 'accept: application/json'

{
  "enabled": true,
  "username": "exampleEmail@example.com",
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "gsxKeystore": {
    "name": "certificate.p12",
    "expirationEpoch": 169195490000,
    "errorMessage": "Certificate error"
  }
}