Deploy packages using MDM
post
https://yourServer.jamfcloud.com/api/v1/deploy-package

Deploys packages to macOS devices using the InstallEnterpriseApplication MDM command.

Query Params
verbose
boolean
Defaults to false
Enables the 'verbose' response, which includes information about the commands queued as well as information about commands that failed to queue.


false
Body Params
manifest
object
required

manifest object
installAsManaged
boolean

true
devices
array of integers

ADD integer
groupId
string
1
Responses

200
Package deployment was successfully processed. See response body for further information.

Response body
object
queuedCommands
array of objects
object
device
integer
commandUuid
string
errors
array of objects
object
device
integer
group
integer
reason
string
202
Package deployment was queued up

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/deploy-package?verbose=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "manifest": {
    "hashType": "MD5"
  },
  "installAsManaged": true,
  "groupId": "1"
}
'

{
  "queuedCommands": [
    {
      "device": 1,
      "commandUuid": "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937"
    }
  ],
  "errors": [
    {
      "device": 2,
      "group": 3,
      "reason": "Device does not support the InstallEnterpriseApplication command"
    }
  ]
}
-----
Renew MDM Profile
post
https://yourServer.jamfcloud.com/api/v1/mdm/renew-profile

Renews the device's MDM Profile, including the device identity certificate within the MDM Profile.

Body Params
List of devices' UDIDs to perform MDM profile renewal

udids
array of strings

string

6E47EF55-5318-494F-A09E-70F613E0AFD1

string

YE47EF55-5318-494F-A09E-70F613E0AFD1

ADD string
Responses

202
The Renew MDM Profile action was queued with APNs for submitted devices. Any device udids returned were unknown and not submitted.

Response body
object
udidsNotProcessed
object
udids
array of strings

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/mdm/renew-profile \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "udids": [
    "6E47EF55-5318-494F-A09E-70F613E0AFD1",
    "YE47EF55-5318-494F-A09E-70F613E0AFD1"
  ]
}
'

{
  "udidsNotProcessed": {
    "udids": [
      "6E47EF55-5318-494F-A09E-70F613E0AFD1"
    ]
  }
}
-----
Send blank push notifications to a list of client management IDs.
post
https://yourServer.jamfcloud.com/api/v2/mdm/blank-push

Accepts a list of client management IDs and sends a blank push notification to each. Returns a list of UUIDs that encountered errors.

Body Params
A list of client management IDs to send push notifications to.

clientManagementIds
array of strings
required

string

fd68c371-5921-436e-b16b-8a3c1bf90ee5

ADD string
Responses

200
Push notifications sent successfully, with a list of UUIDs that encountered errors.

Response body
object
errorUuids
array of strings
required

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mdm/blank-push \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "clientManagementIds": [
    "fd68c371-5921-436e-b16b-8a3c1bf90ee5"
  ]
}
'

{
  "errorUuids": [
    "a1b2c3d4-5678-90ab-cdef-1234567890ab",
    "fd68c371-5921-436e-b16b-8a3c1bf90ee5"
  ]
}
-----
Get information about mdm commands made by Jamf Pro.
get
https://yourServer.jamfcloud.com/api/v2/mdm/commands

Get information about mdm commands made by Jamf Pro.

Query Params
page
integer
Defaults to 0
0
page-size
integer
Defaults to 100
100
sort
array of strings
Defaults to dateSent:asc
Default sort is dateSent:asc. Multiple sort criteria are supported and must be separated with a comma.


string

dateSent:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter, for a list of commands. All url must contain minimum one filter field. Fields allowed in the query: uuid, clientManagementId, command, status, clientType, dateSent, validAfter, dateCompleted, profileId, profileIdentifier, and active. This param can be combined with paging. Please note that any date filters must be used with gt, lt, ge, le Example: clientManagementId==fb511aae-c557-474f-a9c1-5dc845b90d0f;status==Pending;command==INSTALL_PROFILE;uuid==9e18f849-e689-4f2d-b616-a99d3da7db42;clientType==COMPUTER_USER;profileId==1;profileIdentifier==18cc61c2-01fc-11ed-b939-0242ac120002;dateCompleted=ge=2021-08-04T14:25:18.26Z;dateCompleted=le=2021-08-04T14:25:18.26Z;validAfter=ge=2021-08-05T14:25:18.26Z;active==true

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
uuid
string
dateSent
date-time
dateCompleted
date-time
client
object

client object
managementId
string
clientType
string
enum
MOBILE_DEVICE TV VISION_PRO WATCH COMPUTER COMPUTER_USER MOBILE_DEVICE_USER UNKNOWN

commandState
string
enum
PENDING ACKNOWLEDGED NOT_NOW ERROR

commandType
string
enum
APPLY_REDEMPTION_CODE CERTIFICATE_LIST CLEAR_PASSCODE CLEAR_RESTRICTIONS_PASSWORD DECLARATIVE_MANAGEMENT DELETE_USER DEVICE_INFORMATION DEVICE_LOCATION DEVICE_LOCK DISABLE_LOST_MODE DISABLE_REMOTE_DESKTOP ENABLE_LOST_MODE ENABLE_REMOTE_DESKTOP ERASE_DEVICE INSTALLED_APPLICATION_LIST LOG_OUT_USER MANAGED_APPLICATION_LIST MANAGED_MEDIA_LIST REFRESH_CELLULAR_PLANS PLAY_LOST_MODE_SOUND PROVISIONING_PROFILE_LIST RESTART_DEVICE REQUEST_MIRRORING SECURITY_INFO SETTINGS SET_AUTO_ADMIN_PASSWORD SET_RECOVERY_LOCK SHUT_DOWN_DEVICE STOP_MIRRORING UNLOCK_USER_ACCOUNT VALIDATE_APPLICATIONS VALIDATE_RECOVERY_LOCK

commandError
object

commandError object
errorCode
integer
errorDomain
string
errorLocalizedDescription
string
errorEnglishDescription
string
profileId
integer

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/mdm/commands?page=0&page-size=100&sort=dateSent%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "uuid": "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937",
      "dateSent": "2019-05-16T20:43:43.945Z",
      "dateCompleted": "2019-05-16T20:43:43.945Z",
      "client": {
        "managementId": "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937",
        "clientType": "MOBILE_DEVICE"
      },
      "commandState": "PENDING",
      "commandType": "ENABLE_LOST_MODE",
      "commandError": {
        "errorCode": 1234,
        "errorDomain": "An error occurred while processing the command.",
        "errorLocalizedDescription": "An error occurred while processing the command.",
        "errorEnglishDescription": "An error occurred while processing the command."
      },
      "profileId": 1
    }
  ]
}
-----
Post a command for creation and queuing
post
https://yourServer.jamfcloud.com/api/v2/mdm/commands

Provided an MDM command type and appropriate information, will create and then queue said command. A separate privilege is required for each device type and MDM command you want to view or send.

For additional details on how to use the DECLARATIVE_MANAGEMENT command type, see Apple's documentation .

Body Params
The mdm command object to create and queue

clientData
array of objects

object

managementId
string
aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937

ADD object
commandData

APPLY_REDEMPTION_CODE

CERTIFICATE_LIST

CLEAR_PASSCODE

CLEAR_RESTRICTIONS_PASSWORD

DECLARATIVE_MANAGEMENT

DELETE_USER

DEVICE_INFORMATION

DEVICE_LOCATION

DEVICE_LOCK

DISABLE_LOST_MODE

DISABLE_REMOTE_DESKTOP

ENABLE_LOST_MODE

ENABLE_REMOTE_DESKTOP

ERASE_DEVICE

INSTALLED_APPLICATION_LIST

LOG_OUT_USER

MANAGED_APPLICATION_LIST

MANAGED_MEDIA_LIST

PLAY_LOST_MODE_SOUND

PROVISIONING_PROFILE_LIST

REFRESH_CELLULAR_PLANS

RESTART_DEVICE

REQUEST_MIRRORING

SECURITY_INFO

SETTINGS

SET_AUTO_ADMIN_PASSWORD

SET_RECOVERY_LOCK
commandType
string
enum
required

SET_RECOVERY_LOCK

Show 32 enum values
newPassword
password
The new password for Recovery Lock. Set as an empty string to clear the Recovery Lock password.

•••••••••••••

SHUT_DOWN_DEVICE

STOP_MIRRORING

UNLOCK_USER_ACCOUNT

VALIDATE_APPLICATIONS

VALIDATE_RECOVERY_LOCK
Responses

201
Success

Response body
array of objects
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v2/mdm/commands \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "commandData": {
    "commandType": "SET_RECOVERY_LOCK",
    "newPassword": "newQuerty1234"
  },
  "clientData": [
    {
      "managementId": "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937"
    }
  ]
}
'

[
  {
    "id": "1",
    "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
  }
]