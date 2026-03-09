Create Azure Cloud Identity Provider configuration
post
https://yourServer.jamfcloud.com/api/v1/cloud-azure

Create new Azure Cloud Identity Provider configuration with unique display name.

Body Params
Azure Cloud Identity Provider configuration to create

cloudIdPCommon
object
required
A Cloud Identity Provider information for request


cloudIdPCommon object
server
object
required
Azure Cloud Identity Provider configuration request


server object
Responses

201
Azure Cloud Identity Provider configuration created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-azure \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cloudIdPCommon": {
    "providerName": "GOOGLE"
  },
  "server": {
    "enabled": true,
    "transitiveMembershipEnabled": true,
    "transitiveDirectoryMembershipEnabled": true
  }
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----

Get default server configuration
get
https://yourServer.jamfcloud.com/api/v1/cloud-azure/defaults/server-configuration


This is the default set of attributes that allows you to return the data you need from Azure AD. Some fields may be empty and may be edited when creating a new configuration.

Responses

200
Default server configuration returned.

Response body
object
id
string
required
tenantId
string
required
enabled
boolean
required
deprecatedConsent
boolean
required
migrated
boolean
required
mappings
object
required
Azure Cloud Identity Provider mappings

userId
string
required
userName
string
required
realName
string
required
email
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
groupId
string
required
groupName
string
required
searchTimeout
integer
required
5 to 600
transitiveMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup with Single Sign On

transitiveMembershipUserField
string
required
Use this field to set user field mapping for transitive membership lookup with Single Sign On

transitiveDirectoryMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On

membershipCalculationOptimizationEnabled
boolean
Use this field to enable membership calculation optimization. This setting would not apply to Single Sign On

type
string
enum
Type of Entra ID connection

PUBLIC GCC_HIGH

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-azure/defaults/server-configuration \
     --header 'accept: application/json'

{
  "id": "1001",
  "tenantId": "db65d325-0350-4a17-9af9-b302d0fc386b",
  "enabled": true,
  "deprecatedConsent": true,
  "migrated": true,
  "mappings": {
    "userId": "id",
    "userName": "userPrincipalName",
    "realName": "displayName",
    "email": "mail",
    "department": "department",
    "building": "companyName",
    "room": "officeLocation",
    "phone": "mobilePhone",
    "position": "jobTitle",
    "groupId": "id",
    "groupName": "displayName"
  },
  "searchTimeout": 30,
  "transitiveMembershipEnabled": false,
  "transitiveMembershipUserField": "userPrincipalName",
  "transitiveDirectoryMembershipEnabled": false,
  "membershipCalculationOptimizationEnabled": true,
  "type": "PUBLIC"
}
-----

Get Azure Cloud Identity Provider configuration with given ID.
get
https://yourServer.jamfcloud.com/api/v1/cloud-azure/{id}

Get Azure Cloud Identity Provider configuration with given ID.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses

200
Cloud Identity Provider configuration returned.

Response body
object
cloudIdPCommon
object
required
A Cloud Identity Provider information

id
string
required
displayName
string
required
providerName
string
enum
required
GOOGLE AZURE

server
object
required
Azure Cloud Identity Provider configuration

id
string
required
tenantId
string
required
enabled
boolean
required
deprecatedConsent
boolean
required
migrated
boolean
required
mappings
object
required
Azure Cloud Identity Provider mappings


mappings object
searchTimeout
integer
required
5 to 600
transitiveMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup with Single Sign On

transitiveMembershipUserField
string
required
Use this field to set user field mapping for transitive membership lookup with Single Sign On

transitiveDirectoryMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On

membershipCalculationOptimizationEnabled
boolean
Use this field to enable membership calculation optimization. This setting would not apply to Single Sign On

type
string
enum
Type of Entra ID connection

PUBLIC GCC_HIGH

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-azure/ \
     --header 'accept: application/json'

{
  "cloudIdPCommon": {
    "id": "1001",
    "displayName": "Cloud Identity Provider",
    "providerName": "PROVIDER"
  },
  "server": {
    "id": "1001",
    "tenantId": "db65d325-0350-4a17-9af9-b302d0fc386b",
    "enabled": true,
    "deprecatedConsent": true,
    "migrated": true,
    "mappings": {
      "userId": "id",
      "userName": "userPrincipalName",
      "realName": "displayName",
      "email": "mail",
      "department": "department",
      "building": "companyName",
      "room": "officeLocation",
      "phone": "mobilePhone",
      "position": "jobTitle",
      "groupId": "id",
      "groupName": "displayName"
    },
    "searchTimeout": 30,
    "transitiveMembershipEnabled": false,
    "transitiveMembershipUserField": "userPrincipalName",
    "transitiveDirectoryMembershipEnabled": false,
    "membershipCalculationOptimizationEnabled": true,
    "type": "PUBLIC"
  }
}
-----

Update Azure Cloud Identity Provider configuration
put
https://yourServer.jamfcloud.com/api/v1/cloud-azure/{id}

Update Azure Cloud Identity Provider configuration. Cannot be used for partial updates, all content body must be sent.

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Azure Cloud Identity Provider configuration to update

cloudIdPCommon
object
required
A Cloud Identity Provider information


cloudIdPCommon object
server
object
required
Azure Cloud Identity Provider configuration update


server object
Responses

200
Cloud Identity Provider configuration updated

Response body
object
cloudIdPCommon
object
required
A Cloud Identity Provider information

id
string
required
displayName
string
required
providerName
string
enum
required
GOOGLE AZURE

server
object
required
Azure Cloud Identity Provider configuration

id
string
required
tenantId
string
required
enabled
boolean
required
deprecatedConsent
boolean
required
migrated
boolean
required
mappings
object
required
Azure Cloud Identity Provider mappings


mappings object
userId
string
required
userName
string
required
realName
string
required
email
string
required
department
string
required
building
string
required
room
string
required
phone
string
required
position
string
required
groupId
string
required
groupName
string
required
searchTimeout
integer
required
5 to 600
transitiveMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup with Single Sign On

transitiveMembershipUserField
string
required
Use this field to set user field mapping for transitive membership lookup with Single Sign On

transitiveDirectoryMembershipEnabled
boolean
required
Use this field to enable transitive membership lookup. This setting would not apply to Single Sign On

membershipCalculationOptimizationEnabled
boolean
Use this field to enable membership calculation optimization. This setting would not apply to Single Sign On

type
string
enum
Type of Entra ID connection

PUBLIC GCC_HIGH

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-azure/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "cloudIdPCommon": {
    "providerName": "GOOGLE"
  },
  "server": {
    "enabled": true,
    "transitiveMembershipEnabled": true,
    "transitiveDirectoryMembershipEnabled": true
  }
}
'
{
  "cloudIdPCommon": {
    "id": "1001",
    "displayName": "Cloud Identity Provider",
    "providerName": "PROVIDER"
  },
  "server": {
    "id": "1001",
    "tenantId": "db65d325-0350-4a17-9af9-b302d0fc386b",
    "enabled": true,
    "deprecatedConsent": true,
    "migrated": true,
    "mappings": {
      "userId": "id",
      "userName": "userPrincipalName",
      "realName": "displayName",
      "email": "mail",
      "department": "department",
      "building": "companyName",
      "room": "officeLocation",
      "phone": "mobilePhone",
      "position": "jobTitle",
      "groupId": "id",
      "groupName": "displayName"
    },
    "searchTimeout": 30,
    "transitiveMembershipEnabled": false,
    "transitiveMembershipUserField": "userPrincipalName",
    "transitiveDirectoryMembershipEnabled": false,
    "membershipCalculationOptimizationEnabled": true,
    "type": "PUBLIC"
  }
}
-----

Delete Cloud Identity Provider configuration.
delete
https://yourServer.jamfcloud.com/api/v1/cloud-azure/{id}

Delete Cloud Identity Provider configuration.

Path Params
id
string
required
Cloud Identity Provider identifier

Responses
204
Cloud Identity Provider configuration deleted.

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-azure/ \
     --header 'accept: application/json'