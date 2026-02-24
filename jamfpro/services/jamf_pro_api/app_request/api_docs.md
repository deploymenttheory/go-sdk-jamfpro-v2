Search for Form Input Fields
get
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields


Search for form input fields

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
integer
title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": 1,
      "title": "Quantity",
      "description": "How many of these would you like?",
      "priority": 1
    }
  ]
}

-----

Replace all Form Input Fields
put
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields


Replace all form input fields. Will delete, update, and create all input fields accordingly.

Body Params
list of form input fields to replace all existing fields. Will delete, update, and create all input fields accordingly.


ADD object
Responses

200
form input fields were replaced

Response body
array of objects
object
id
integer
title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

[
  {
    "id": 1,
    "title": "Quantity",
    "description": "How many of these would you like?",
    "priority": 1
  }
]

-----

Create Form Input Field record
post
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields


Create form input field record

Body Params
form input field object to create. ids defined in this body will be ignored

title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

Responses

201
form input field record was created

Response body
object
id
integer
title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "title": "Quantity",
  "description": "How many of these would you like?",
  "priority": 1
}

-----

Get specified Form Input Field object
get
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields/{id}


Gets specified form input field object

Path Params
id
integer
required
Instance id of form input field record

Responses

200
Details of form input field were found

Response body
object
id
integer
title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields/ \
     --header 'accept: application/json'

{
  "id": 1,
  "title": "Quantity",
  "description": "How many of these would you like?",
  "priority": 1
}

-----

Update specified Form Input Field object
put
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields/{id}


Update specified form input field object

Path Params
id
integer
required
Instance id of form input field record

Body Params
form input field object to create. ids defined in this body will be ignored

title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

Responses

200
form input field update

Response body
object
id
integer
title
string
required
length ≥ 1
description
string | null
priority
integer
required
1 to 255
Highest priority is 1, lowest is 255

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "title": "Quantity",
  "description": "How many of these would you like?",
  "priority": 1
}

-----

Remove specified Form Input Field record
delete
https://yourServer.jamfcloud.com/api/v1/app-request/form-input-fields/{id}


Removes specified form input field record

Path Params
id
integer
required
Instance id of form input field record

Responses
204
form input field record was deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/form-input-fields/ \
     --header 'accept: application/json'

-----


advanced-user-content-searches

api-authentication

api-integrations

api-role-privileges

api-roles

apns-client-push-status

app-request-preview

Search for Form Input Fields
get
Replace all Form Input Fields
put
Create Form Input Field record
post
Get specified Form Input Field object
get
Update specified Form Input Field object
put
Remove specified Form Input Field record
del
Get Applicastion Request Settings
get
Update Application Request Settings
put
app-store-country-codes-preview

branding

buildings

cache-settings

categories

certificate-authority

classic-ldap

client-check-in

cloud-azure

cloud-distribution-point

cloud-idp

cloud-information

cloud-ldap

computer-extension-attributes

computer-groups

computer-inventory

computer-inventory-collection-settings

computer-prestages

computers-preview

conditional-access

csa

dashboard

declarative-device-management

departments

devices

device-communication-settings

device-enrollments

device-enrollments-devices

digicert

distribution-point

dock-items

ebooks

engage
enrollment

enrollment-customization

enrollment-customization-preview

groups

gsx-connection

health-check

icon

impact-alert-notification-settings

inventory-information

inventory-preload

jamf-cloud-distribution-service

jamf-connect

jamf-management-framework

jamf-package

jamf-pro-account-preferences

jamf-pro-information

jamf-pro-initialization

jamf-pro-notifications

jamf-pro-server-url-preview

jamf-pro-user-account-settings

jamf-pro-version

jamf-protect

jamf-remote-assist

ldap

local-admin-password

locales-preview

log-flushing

login-customization

macos-managed-software-updates

managed-software-updates

mdm

mdm-renewal

mobile-device-apps

mobile-device-enrollment-profile

mobile-device-extension-attributes

mobile-device-extension-attributes-preview

mobile-device-groups

mobile-device-prestages

mobile-devices

oidc

onboarding

packages

parent-app-preview

patch-management

patch-policies

patch-policy-logs

patch-software-title-configurations

policies-preview

re-enrollment-preview

remote-administration

return-to-service

scheduler

scripts

self-service

self-service-branding-ios

self-service-branding-macos

self-service-branding-preview

self-service-plus

service-discovery-enrollment

sites

slasa

smart-computer-groups-preview

smart-mobile-device-groups-preview

smart-user-groups-preview

smtp-server

sso-certificate

sso-failover

sso-oauth-session-tokens

sso-settings

startup-status

static-user-groups-preview

supervision-identities-preview

teacher-app

team-viewer-remote-administration

time-zones-preview

tomcat-settings-preview

user

user-session-preview

venafi-preview

volume-purchasing-locations

volume-purchasing-subscriptions

Title Editor
auth

capabilities

codesigning

components

criteria

extensionattributes

externaltitles

killapps

overrides

patches

preferences

privileges

requirements

smtpserver

softwaretitles

sources

users

valuelists

Powered by 

Get Applicastion Request Settings
get
https://yourServer.jamfcloud.com/api/v1/app-request/settings


Get app request settings

Response

200
Successful response

Response body
object
isEnabled
boolean
appStoreLocale
string
Can be any of the country codes from /v1/app-store-country-codes or "deviceLocale" to use each individual device's locale

requesterUserGroupId
integer
approverEmails
array of strings

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/settings \
     --header 'accept: application/json'

{
  "isEnabled": true,
  "appStoreLocale": "deviceLocale",
  "requesterUserGroupId": 1,
  "approverEmails": [
    "jane.doe@company.com, john.doe@company.com"
  ]
}

-----

Update Application Request Settings
put
https://yourServer.jamfcloud.com/api/v1/app-request/settings


Update app request settings

Body Params
App request settings object

isEnabled
boolean

true
appStoreLocale
string
Can be any of the country codes from /v1/app-store-country-codes or "deviceLocale" to use each individual device's locale

deviceLocale
requesterUserGroupId
integer
1
approverEmails
array of strings

string

jane.doe@company.com, john.doe@company.com

ADD string
Responses

200
App request settings updated

Response body
object
isEnabled
boolean
appStoreLocale
string
Can be any of the country codes from /v1/app-store-country-codes or "deviceLocale" to use each individual device's locale

requesterUserGroupId
integer
approverEmails
array of strings

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/app-request/settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "approverEmails": [
    "jane.doe@company.com, john.doe@company.com"
  ],
  "isEnabled": true,
  "appStoreLocale": "deviceLocale",
  "requesterUserGroupId": 1
}
'

{
  "isEnabled": true,
  "appStoreLocale": "deviceLocale",
  "requesterUserGroupId": 1,
  "approverEmails": [
    "jane.doe@company.com, john.doe@company.com"
  ]
}