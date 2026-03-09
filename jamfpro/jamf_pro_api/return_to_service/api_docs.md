Get all Return to Service Configurations
get
https://yourServer.jamfcloud.com/api/v1/return-to-service

Gets all return to service configurations.

Response

200
Successful response

Response body
object
totalCount
integer
≥ 0
results
array of objects
length ≥ 0
object
id
string
≥ 1
Id of the Return to Service Configuration.

displayName
string
Defaults to false
wifiProfileId
string
Id of the wifi profile that is associated with the return to service configuration.

Create a Return to Service Configuration
post
https://yourServer.jamfcloud.com/api/v1/return-to-service

Create a return to service configuration

Body Params
Return to Service Configuration to create. ids defined in this body will be ignored

displayName
string
Defaults to false
Display name of the Return to Service Configuration.

false
wifiProfileId
string
Id of the wifi profile that is associated with the return to service configuration.

1
Responses

201
Return to Service Configuration was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/return-to-service \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "false",
  "wifiProfileId": "1"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----

Retrieve a Return to Service Configuration with the supplied id
get
https://yourServer.jamfcloud.com/api/v1/return-to-service/{id}

Retrieves a Return to Service Configuration with the supplied id

Path Params
id
string
required
Return to Service Configuration identifier

Responses

200
Success

Response body
object
id
string
≥ 1
Id of the Return to Service Configuration.

displayName
string
Defaults to false
wifiProfileId
string
Id of the wifi profile that is associated with the return to service configuration.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/return-to-service/ \
     --header 'accept: application/json'

{
  "id": "1",
  "displayName": "displayName",
  "wifiProfileId": "1"
}
-----
Update a Return to Service Configuration
put
https://yourServer.jamfcloud.com/api/v1/return-to-service/{id}

Updates a Return to Service Configuration

Path Params
id
string
required
Return to Service Configuration identifier

Body Params
Return to Service Configuration to update

displayName
string
Defaults to false
Display name of the Return to Service Configuration.

false
wifiProfileId
string
Id of the wifi profile that is associated with the return to service configuration.

1
Responses

200
Success

Response body
object
id
string
≥ 1
Id of the Return to Service Configuration.

displayName
string
Defaults to false
wifiProfileId
string
Id of the wifi profile that is associated with the return to service configuration.

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/return-to-service/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "displayName": "false",
  "wifiProfileId": "1"
}
'
{
  "id": "1",
  "displayName": "displayName",
  "wifiProfileId": "1"
}
-----
Delete a Return To Service Configuration with the supplied id
delete
https://yourServer.jamfcloud.com/api/v1/return-to-service/{id}


Deletes a Return To Service Configuration with the supplied id

Path Params
id
string
required
Return To Service Configurations identifier

Responses
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/return-to-service/ \
     --header 'accept: application/json'

