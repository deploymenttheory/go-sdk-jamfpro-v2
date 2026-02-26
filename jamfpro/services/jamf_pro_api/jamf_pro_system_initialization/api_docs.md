Set up fresh installed Jamf Pro Server
post
https://yourServer.jamfcloud.com/api/v1/system/initialize

Set up fresh installed Jamf Pro Server

Body Params
activationCode
string
required
length between 39 and 39
VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB
institutionName
string
required
length ≥ 1
Jamf
eulaAccepted
boolean
required

true
username
string
required
length ≥ 1
admin
password
password
required
length ≥ 1
•••••
email
string
ITBob@jamf.com
jssUrl
string
required
length ≥ 1
https://jamf.jamfcloud.com
Responses
202
Jamf Pro Server has been initialized

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/system/initialize \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "eulaAccepted": true,
  "activationCode": "VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB",
  "institutionName": "Jamf",
  "username": "admin",
  "password": "12345",
  "email": "ITBob@jamf.com",
  "jssUrl": "https://jamf.jamfcloud.com"
}
'
-----
Provide Database Password during startup
post
https://yourServer.jamfcloud.com/api/v1/system/initialize-database-connection


Provide database password during startup. Endpoint is accessible when database password was not configured and Jamf Pro server has not been initialized yet.

Body Params
password
password
required
length ≥ 1
•••••
Responses
202
OK

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/system/initialize-database-connection \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"password":"12345"}'
-----

Set up fresh installed Jamf Pro Server for Platform
post
https://yourServer.jamfcloud.com/api/v1/system/platform-initialize


Set up fresh installed Jamf Pro Server with OIDC SSO enabled and single federated user

Body Params
activationCode
string
required
length between 39 and 39
VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB
institutionName
string
required
length ≥ 1
jamf
eulaAccepted
boolean
required

true
username
string
required
length ≥ 1
Federated user OIDC username to create

admin
email
string
required
Federated user OIDC email to create

ITBob@jamf.com
jssUrl
string
required
length ≥ 1
https://jamf.jamfcloud.com
Responses
202
Jamf Pro Server has been initialized

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/system/platform-initialize \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "eulaAccepted": true,
  "activationCode": "VFAB-YDAB-DFAB-UDAB-DEAB-EFAB-ABAB-DEAB",
  "institutionName": "jamf",
  "username": "admin",
  "email": "ITBob@jamf.com",
  "jssUrl": "https://jamf.jamfcloud.com"
}
'