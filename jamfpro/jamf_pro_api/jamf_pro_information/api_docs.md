Get basic information about the Jamf Pro Server
get
https://yourServer.jamfcloud.com/api/v2/jamf-pro-information

Get basic information about the Jamf Pro Server

Response

200
Successful response

Response body
object
vppTokenEnabled
boolean
depAccountEnabled
boolean
byodEnabled
boolean
Defaults to false
deprecated
Deprecated as of 11.25. This field always returns false.

userMigrationEnabled
boolean
cloudDeploymentsEnabled
boolean
patchEnabled
boolean
ssoSamlEnabled
boolean
smtpEnabled
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/jamf-pro-information \
     --header 'accept: application/json'

{
  "vppTokenEnabled": false,
  "depAccountEnabled": false,
  "userMigrationEnabled": false,
  "cloudDeploymentsEnabled": false,
  "patchEnabled": false,
  "ssoSamlEnabled": false,
  "smtpEnabled": false
}