Get mappings for OnPrem Ldap configuration with given id.
get
https://yourServer.jamfcloud.com/api/v1/classic-ldap/{id}

Get mappings for OnPrem Ldap configuration with given id.

Path Params
id
string
required
OnPrem Ldap identifier

Responses

200
OnPrem Ldap mappings returned.

Response body
object
userObjectMapIdTo
string
required
userObjectMapUsernameTo
string
required
userObjectMapRealNameTo
string
required
userObjectMapEmailTo
string
required
userObjectMapDepartmentTo
string
required
userObjectMapBuildingTo
string
required
userObjectMapRoomTo
string
required
userObjectMapPhoneTo
string
required
userObjectMapPositionTo
string
required
userObjectMapUuidTo
string
required
userGroupObjectMapIdTo
string
required
userGroupObjectMapGroupNameTo
string
required
userGroupObjectMapUuidTo
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/classic-ldap/ \
     --header 'accept: application/json'

{
  "userObjectMapIdTo": "mail",
  "userObjectMapUsernameTo": "uid",
  "userObjectMapRealNameTo": "displayName",
  "userObjectMapEmailTo": "mail",
  "userObjectMapDepartmentTo": "departmentNumber",
  "userObjectMapBuildingTo": "",
  "userObjectMapRoomTo": "",
  "userObjectMapPhoneTo": "",
  "userObjectMapPositionTo": "title",
  "userObjectMapUuidTo": "uid",
  "userGroupObjectMapIdTo": "name",
  "userGroupObjectMapGroupNameTo": "name",
  "userGroupObjectMapUuidTo": "uid"
}