Download a self service branding image
get
https://yourServer.jamfcloud.com/api/v1/branding-images/download/{id}


Download a self service branding image

Path Params
id
string
required
id of the self service branding image

Response

200
Successful response

Response body
file

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/branding-images/download/ \
     --header 'accept: image/*'

