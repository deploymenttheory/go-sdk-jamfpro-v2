Retrieve the configured access groups that contain the text in the search param
get
https://yourServer.jamfcloud.com/api/ldap/groups

Retrieves the configured access groups that contain the text in the searchParam.

Query Params
q
string
Defaults to null
Will perform a "contains" search on the names of access groups

null
Response

200
Successful search.

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
id
string
uuid
string
ldapServerId
integer
name
string
distinguishedName
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/ldap/groups?q=null' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
      "ldapServerId": 1,
      "name": "Grade School Teachers",
      "distinguishedName": "Grade School Teachers"
    }
  ]
}
-----
Retrieve all Servers including LDAP and Cloud Identity Providers.
get
https://yourServer.jamfcloud.com/api/ldap/servers

Retrieve all Servers including LDAP and Cloud Identity Providers.

Response

200
Successfully completed.

Response body
array of objects
object
id
integer
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/ldap/servers \
     --header 'accept: application/json'

[
  {
    "id": 1,
    "name": "Server name"
  }
]
-----
Retrieve the configured access groups that contain the text in the search param
get
https://yourServer.jamfcloud.com/api/v1/ldap/groups

Retrieves the configured access groups that contain the text in the searchParam.

Query Params
q
string
Defaults to null
Will perform a "contains" search on the names of access groups

null
Response

200
Successful search.

Response body
object
totalCount
integer
≥ 0
results
array of objects
object
id
string
uuid
string
ldapServerId
integer
name
string
distinguishedName
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/ldap/groups?q=null' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "uuid": "89AF33FC-123C-1231-AEFD-9C3ED123AFCC",
      "ldapServerId": 1,
      "name": "Grade School Teachers",
      "distinguishedName": "Grade School Teachers"
    }
  ]
}
-----
Retrieve all LDAP Servers.
get
https://yourServer.jamfcloud.com/api/v1/ldap/ldap-servers

Retrieves all not migrated, LDAP Servers.

Response

200
Successfully completed.

Response body
array of objects
object
id
integer
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ldap/ldap-servers \
     --header 'accept: application/json'

[
  {
    "id": 1,
    "name": "Server name"
  }
]
-----
Retrieve all Servers including LDAP and Cloud Identity Providers.
get
https://yourServer.jamfcloud.com/api/v1/ldap/servers

Retrieve all active Servers including LDAP and Cloud Identity Providers.

Response

200
Successfully completed.

Response body
array of objects
object
id
integer
name
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/ldap/servers \
     --header 'accept: application/json'

[
  {
    "id": 1,
    "name": "Server name"
  }
]