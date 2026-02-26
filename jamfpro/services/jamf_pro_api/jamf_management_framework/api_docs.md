Redeploy Jamf Management Framework
post
https://yourServer.jamfcloud.com/api/v1/jamf-management-framework/redeploy/{id}


Redeploys the Jamf Management Framework for enrolled device

Path Params
id
string
required
instance id of computer

1
Responses

202
Command successfully queued to redeploy the Jamf Managment Framework

Response body
object
deviceId
string
commandUuid
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/jamf-management-framework/redeploy/1 \
     --header 'accept: application/json'

{
  "deviceId": "1",
  "commandUuid": "f5965c4f-0db4-4dc4-9f37-6f1dad4e939c"
}