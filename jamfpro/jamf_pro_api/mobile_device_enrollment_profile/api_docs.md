Retrieve the MDM Enrollment Profile
get
https://yourServer.jamfcloud.com/api/v1/mobile-device-enrollment-profile/{id}/download-profile

Retrieve the MDM Enrollment Profile

Path Params
id
string
required
MDM Enrollment Profile identifier

Responses

200
Success

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/mobile-device-enrollment-profile//download-profile \
     --header 'accept: application/x-apple-aspen-config'
