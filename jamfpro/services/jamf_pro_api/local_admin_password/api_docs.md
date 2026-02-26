Get a list of the current devices and usernames with pending LAPS rotations
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/pending-rotations


Return information about all devices and usernames currently in the state of a pending LAPS rotation

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
lapsUser
object

lapsUser object
clientManagementId
string
guid
string
username
string
userSource
string
enum
MDM JMF

createdDate
date-time

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password/pending-rotations \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "lapsUser": {
        "clientManagementId": "2db90ebf-ce9c-4078-b508-034c8ee3a060",
        "guid": "602F96A6-7BC4-43CD-95F2-1DD3B8BC0AF3",
        "username": "admin",
        "userSource": "MDM"
      },
      "createdDate": "2019-05-16T20:43:43.945Z"
    }
  ]
}
-----
Get the current LAPS settings.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/settings


Return information about the current LAPS settings.

Responses

200
Success

Response body
object
autoDeployEnabled
boolean
When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically.

passwordRotationTime
integer
The amount of time in seconds that the local admin password will be rotated after viewing.

autoRotateEnabled
boolean
When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime

autoRotateExpirationTime
integer
The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed.

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password/settings \
     --header 'accept: application/json'

{
  "autoDeployEnabled": false,
  "passwordRotationTime": 3600,
  "autoRotateEnabled": false,
  "autoRotateExpirationTime": 7776000
}
-----
Update settings for LAPS.
put
https://yourServer.jamfcloud.com/api/v2/local-admin-password/settings


Update settings for LAPS.

Body Params
LAPS settings to update

autoDeployEnabled
boolean
required
When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically.


true
passwordRotationTime
integer
required
The amount of time in seconds that the local admin password will be rotated after viewing.

autoRotateEnabled
boolean
required
When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime


true
autoRotateExpirationTime
integer
required
The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed.

Responses

200
Success

Response body
object
autoDeployEnabled
boolean
When enabled, all appropriate computers will have the SetAutoAdminPassword command sent to them automatically.

passwordRotationTime
integer
The amount of time in seconds that the local admin password will be rotated after viewing.

autoRotateEnabled
boolean
When enabled, all appropriate computers will automatically have their password expired and rotated after the configured autoRotateExpirationTime

autoRotateExpirationTime
integer
The amount of time in seconds that the local admin password will be rotated automatically if it is never viewed.

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password/settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "autoDeployEnabled": true,
  "autoRotateEnabled": true
}
'

{
  "autoDeployEnabled": false,
  "passwordRotationTime": 3600,
  "autoRotateEnabled": false,
  "autoRotateExpirationTime": 7776000
}
-----
Get LAPS password viewed history.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/audit


Get the full history of all local admin passwords for a specific username on a target device. History will include password, who viewed the password and when it was viewed. Get audit history by using the client management id and username as the path parameters. If multiple accounts with the same username exist, the MDM source will be selected by default.

Path Params
clientManagementId
string
required
client management id of target device.

1
username
string
required
user name to view audit information for

1
Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
password
password
dateLastSeen
date-time | null
expirationTime
date-time | null
audits
array of objects
object
viewedBy
string | null
dateSeen
date-time | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password/1/account/1/audit \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "password": "jamf12345",
      "dateLastSeen": "2019-05-16T20:43:43.945Z",
      "expirationTime": "2019-05-16T20:43:43.945Z",
      "audits": [
        {
          "viewedBy": "admin",
          "dateSeen": "2019-05-16T20:43:43.945Z"
        }
      ]
    }
  ]
}
-----
Get LAPS historical records for target device and username.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/history


Get the full history of all for a specific username on a target device. History will include date created, date last seen, expiration time, and rotational status. Get audit history by using the client management id and username as the path parameters.

Path Params
clientManagementId
string
required
client management id of target device.

username
string
required
user name to view history for

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
createdDate
date-time | null
dateLastSeen
date-time | null
expirationTime
date-time | null
rotationStatus
string
enum
PENDING COMPLETED VIEWED ERROR INVALID

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//account//history \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "createdDate": "2019-05-16T20:43:43.945Z",
      "dateLastSeen": "2019-05-16T20:43:43.945Z",
      "expirationTime": "2019-05-16T20:43:43.945Z",
      "rotationStatus": "PENDING"
    }
  ]
}
-----

Get current LAPS password for specified username on a client.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/password


Get current LAPS password for specified client by using the client management id and username as the path parameters. Once the password is viewed it will be rotated out with a new password based on the rotation time settings. If multiple accounts with the same username exist, the MDM source will be selected by default.

Path Params
clientManagementId
string
required
client management id of target device.

username
string
required
user name for the account

Responses

200
Success

Response body
object
password
password

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//account//password \
     --header 'accept: application/json'

{
  "password": "jamf12345"
}
-----
Get LAPS password viewed history.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/audit


Get the full history of all local admin passwords for a specific user guid on a target device. History will include password, who viewed the password and when it was viewed. Get audit history by using the client management id, username, and user guid as the path parameters.

Path Params
clientManagementId
string
required
client management id of target device.

username
string
required
user name to view audit information for

guid
string
required
user guid to view audit information for

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
password
password
dateLastSeen
date-time | null
expirationTime
date-time | null
audits
array of objects
object
viewedBy
string | null
dateSeen
date-time | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//account///audit \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "password": "jamf12345",
      "dateLastSeen": "2019-05-16T20:43:43.945Z",
      "expirationTime": "2019-05-16T20:43:43.945Z",
      "audits": [
        {
          "viewedBy": "admin",
          "dateSeen": "2019-05-16T20:43:43.945Z"
        }
      ]
    }
  ]
}
-----
Get LAPS historical records for target device and user guid.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/history

Get the full history of all for a specific user guid on a target device. History will include date created, date last seen, expiration time, and rotational status. Get audit history by using the client management id, username, and user guid as the path parameters.

Path Params
clientManagementId
string
required
client management id of target device.

username
string
required
user name to view history for

guid
string
required
user guid to view history for

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
createdDate
date-time | null
dateLastSeen
date-time | null
expirationTime
date-time | null
rotationStatus
string
enum
PENDING COMPLETED VIEWED ERROR INVALID

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//account///history \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "createdDate": "2019-05-16T20:43:43.945Z",
      "dateLastSeen": "2019-05-16T20:43:43.945Z",
      "expirationTime": "2019-05-16T20:43:43.945Z",
      "rotationStatus": "PENDING"
    }
  ]
}
-----
Get current LAPS password for specified user guid on a client.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/account/{username}/{guid}/password

Get current LAPS password for specified client by using the client management id, username, and user guid as the path parameters. Once the password is viewed it will be rotated out with a new password based on the rotation time settings.

Path Params
clientManagementId
string
required
client management id of target device.

username
string
required
user name for the account

guid
string
required
user guid for the account

Responses

200
Success

Response body
object
password
password

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//account///password \
     --header 'accept: application/json'

{
  "password": "jamf12345"
}
-----
Get the LAPS capable admin accounts for a device.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/accounts

Get a full list of admin accounts that are LAPS capable. Capable accounts are returned in the AutoSetupAdminAccounts from QueryResponses.

Path Params
clientManagementId
string
required
client management id of target device.

Responses

200
Success

Response body
object
totalCount
integer
results
array of objects
object
clientManagementId
string
guid
string
username
string
userSource
string
enum
MDM JMF

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//accounts \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "clientManagementId": "2db90ebf-ce9c-4078-b508-034c8ee3a060",
      "guid": "602F96A6-7BC4-43CD-95F2-1DD3B8BC0AF3",
      "username": "admin",
      "userSource": "MDM"
    }
  ]
}
-----
Get LAPS password viewed history, and rotation history.
get
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/history

Get the full history of all local admin passwords for all accounts for a specific management ID. History will include password, who viewed the password and when it was viewed. This will include rotation history as well.

Path Params
clientManagementId
string
required
client management id of target device.

Responses

200
OK

Response body
object
totalCount
integer
results
array of objects
object
username
string
eventType
string
enum
PENDING COMPLETED VIEWED ERROR INVALID

eventTime
date-time | null
viewedBy
string | null
userSource
string
enum
MDM JMF

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//history \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "username": "username",
      "eventType": "VIEWED",
      "eventTime": "2019-05-16T20:43:43.945Z",
      "viewedBy": "admin",
      "userSource": "MDM"
    }
  ]
}
-----
Set the LAPS password for a device.
put
https://yourServer.jamfcloud.com/api/v2/local-admin-password/{clientManagementId}/set-password

Set the LAPS password for a device. This will set the password for all LAPS capable accounts.

Path Params
clientManagementId
string
required
client management id of target device.

Body Params
LAPS password to set

lapsUserPasswordList
array of objects

object

username
string
username
password
password
••••••••

object

username
string
username2
password
password
•••••

object

username
string
password
password

ADD object
Responses

200
Success

Response body
object
lapsUserPasswordList
array of objects
object
username
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/local-admin-password//set-password \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "lapsUserPasswordList": [
    {
      "username": "username",
      "password": "jamf1234"
    },
    {
      "username": "username2",
      "password": "sfsdf"
    }
  ]
}
'

{
  "lapsUserPasswordList": [
    {
      "username": "username"
    }
  ]
}
-----