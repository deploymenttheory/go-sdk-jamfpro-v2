Determines if Self Service Plus feature toggle is enabled.
get
https://yourServer.jamfcloud.com/api/v1/self-service-plus/feature-toggle/enabled

This endpoint is used to determine if the Self Service Plus feature toggle is enabled.

Responses
204
Feature toggle is enabled

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/self-service-plus/feature-toggle/enabled
-----
Get Self Service Plus settings.
get
https://yourServer.jamfcloud.com/api/v1/self-service-plus/settings

Get Self Service Plus settings.

Response

200
Self Service Plus settings

Response body
object
enabled
boolean
Whether Self Service Plus is enabled

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/self-service-plus/settings \
     --header 'accept: application/json'

{
  "enabled": true
}
-----
Save Self Service Plus settings.
put
https://yourServer.jamfcloud.com/api/v1/self-service-plus/settings

Save Self Service Plus settings.

Body Params
Self Service Plus settings

enabled
boolean
Whether Self Service Plus is enabled


true
Response
204
Successful PUT

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/self-service-plus/settings \
     --header 'content-type: application/json'