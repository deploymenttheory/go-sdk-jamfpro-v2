Return information about the Jamf Pro including the current version
get
https://yourServer.jamfcloud.com/api/v1/jamf-pro-version

Returns information about the Jamf Pro including the current version.

Response

200
Successful response

Response body
object
version
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-pro-version \
     --header 'accept: application/json'

{
  "version": "10.9.0"
}