Upload an image
post
https://yourServer.jamfcloud.com/api/self-service/branding/images

Uploads an image

Body Params
file
file
required
The file to upload

No file chosen
Response

201
Image successfully uploaded

Response body
object
url
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/self-service/branding/images \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "url": "https://jamfpro.jamf/image?1"
}
-----