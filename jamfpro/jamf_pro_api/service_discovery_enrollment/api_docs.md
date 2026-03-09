Get service discovery well-known settings for all organizations
get
https://yourServer.jamfcloud.com/api/v1/service-discovery-enrollment/well-known-settings

Returns current settings for all AxM organizations.

Responses

200
Well-known settings retrieved successfully

Response body
object
wellKnownSettings
array of objects
required
Array of well-known settings for all AxM organizations

object
orgName
string
Organization display name

serverUuid
string
required
Server UUID identifier

enrollmentType
string
enum
required
Service discovery enrollment version

none mdm-byod mdm-adde

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/service-discovery-enrollment/well-known-settings \
     --header 'accept: application/json'

{
  "wellKnownSettings": [
    {
      "orgName": "Acme Corporation",
      "serverUuid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "enrollmentType": "mdm-byod"
    }
  ]
}
-----
Update service discovery well-known settings
put
https://yourServer.jamfcloud.com/api/v1/service-discovery-enrollment/well-known-settings

Accepts JSON payload to update enrollment types for AxM organizations. Requires "Update User-Initiated Enrollment" privilege.

Body Params
wellKnownSettings
array of objects
required
Array of well-known settings to update


object

orgName
string
Organization display name

Acme Corporation
serverUuid
string
required
Server UUID identifier

a1b2c3d4-e5f6-7890-abcd-ef1234567890
enrollmentType
string
enum
required
Service discovery enrollment version


mdm-adde
Allowed:

none

mdm-byod

mdm-adde

object

orgName
string
Organization display name

Acme Corporation
serverUuid
string
required
Server UUID identifier

a1b2c3d4-e5f6-7890-abcd-ef1234567890
enrollmentType
string
enum
required
Service discovery enrollment version


mdm-byod
Allowed:

none

mdm-byod

mdm-adde

ADD object
Responses
204
Well-known settings updated successfully

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/service-discovery-enrollment/well-known-settings \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "wellKnownSettings": [
    {
      "enrollmentType": "mdm-adde",
      "orgName": "Acme Corporation",
      "serverUuid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
    },
    {
      "enrollmentType": "mdm-byod",
      "orgName": "Acme Corporation",
      "serverUuid": "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
    }
  ]
}
'
