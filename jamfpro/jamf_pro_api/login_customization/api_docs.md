Get current login disclaimer settings
get
https://yourServer.jamfcloud.com/api/v1/login-customization

Returns knob whether disclaimer is enabled and if saved, its contents.

Response

200
Get Login Customization disclaimer settings.

Response body
object
rampInstance
boolean
includeCustomDisclaimer
boolean
required
disclaimerHeading
string
disclaimerMainText
string
actionText
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/login-customization \
     --header 'accept: application/json'

{
  "rampInstance": false,
  "includeCustomDisclaimer": true,
  "disclaimerHeading": "Disclaimer header",
  "disclaimerMainText": "Login disclaimer main text",
  "actionText": "Accept"
}
-----
Update current login disclaimer settings.
put
https://yourServer.jamfcloud.com/api/v1/login-customization

Update current login disclaimer settings.

Body Params
Login disclaimer settings to save.

includeCustomDisclaimer
boolean
required

true
disclaimerHeading
string
Disclaimer header
disclaimerMainText
string
Login disclaimer main text
actionText
string
Accept
Response

200
Update login customization disclaimer settings.

Response body
object
includeCustomDisclaimer
boolean
required
disclaimerHeading
string
disclaimerMainText
string
actionText
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/login-customization \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "includeCustomDisclaimer": true,
  "disclaimerHeading": "Disclaimer header",
  "disclaimerMainText": "Login disclaimer main text",
  "actionText": "Accept"
}
'

{
  "includeCustomDisclaimer": true,
  "disclaimerHeading": "Disclaimer header",
  "disclaimerMainText": "Login disclaimer main text",
  "actionText": "Accept"
}