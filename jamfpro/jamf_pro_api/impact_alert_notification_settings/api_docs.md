Get Impact Alert Notification Settings
get
https://yourServer.jamfcloud.com/api/v1/impact-alert-notification-settings


Get Impact Alert Notification Settings

Responses

200
Successful

Response body
object
scopeableObjectsAlertEnabled
boolean
required
Defaults to true
scopeableObjectsConfirmationCodeEnabled
boolean
required
Defaults to false
deployableObjectsAlertEnabled
boolean
required
Defaults to true
deployableObjectsConfirmationCodeEnabled
boolean
required
Defaults to false

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/impact-alert-notification-settings \
     --header 'accept: application/json'

{
  "scopeableObjectsAlertEnabled": true,
  "scopeableObjectsConfirmationCodeEnabled": false,
  "deployableObjectsAlertEnabled": true,
  "deployableObjectsConfirmationCodeEnabled": false
}
-----
Update Impact Alert Notification Settings
put
https://yourServer.jamfcloud.com/api/v1/impact-alert-notification-settings


Update Impact Alert Notification Settings

Body Params
Configure Access Management settings

scopeableObjectsAlertEnabled
boolean
required
Defaults to true

true
scopeableObjectsConfirmationCodeEnabled
boolean
required
Defaults to false

false
deployableObjectsAlertEnabled
boolean
required
Defaults to true

true
deployableObjectsConfirmationCodeEnabled
boolean
required
Defaults to false

false
Responses
204
Update Successful

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/impact-alert-notification-settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "scopeableObjectsAlertEnabled": true,
  "scopeableObjectsConfirmationCodeEnabled": false,
  "deployableObjectsAlertEnabled": true,
  "deployableObjectsConfirmationCodeEnabled": false
}
'