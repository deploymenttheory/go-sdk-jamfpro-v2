Upload an icon
post
https://yourServer.jamfcloud.com/api/v1/icon

Uploads an icon

Body Params
file
file
required
The file to upload

No file chosen
Responses

201
Icon successfully uploaded

Response body
object
url
string
id
integer

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/icon \
     --header 'accept: application/json' \
     --header 'content-type: multipart/form-data'

{
  "url": "https://stage-ics.services.jamfcloud.com/icon/hash_c315ef577b84505de1bfcb50b0c4b1c963da30b2a805f84b24ad09f282b7fad4",
  "id": 5
}

Download a self service icon
get
https://yourServer.jamfcloud.com/api/v1/icon/download/{id}

Download a self service icon

Path Params
id
string
required
id of the self service icon

Query Params
res
string
Defaults to original
request a specific resolution of original, 300, or 512; invalid options will result in original resolution

original
scale
string
Defaults to 0
request a scale; 0 results in original image, non-0 results in scaled to 300

0
Response

200
Successful response

Response body
file

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/icon/download/?res=original&scale=0' \
     --header 'accept: image/*'

-----
Get an icon
get
https://yourServer.jamfcloud.com/api/v1/icon/{id}

Get an icon

Path Params
id
string
required
id of the icon

Response

200
Successful response

Response body
object
url
string
id
integer

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/icon/ \
     --header 'accept: application/json'

{
  "url": "https://stage-ics.services.jamfcloud.com/icon/hash_c315ef577b84505de1bfcb50b0c4b1c963da30b2a805f84b24ad09f282b7fad4",
  "id": 5
}
