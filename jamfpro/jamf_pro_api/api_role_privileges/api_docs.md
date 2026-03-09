Get the current Jamf API Role Privileges
get
https://yourServer.jamfcloud.com/api/v1/api-role-privileges

Get role privileges

Response

200
A sorted list of the current Jamf API Roles Privileges

Response body
object
privileges
array of strings
required
length ≥ 0

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/api-role-privileges \
     --header 'accept: application/json'

{
  "privileges": [
    "Flush MDM Commands"
  ]
}
-----

Search the current Jamf API Role Privileges
get
https://yourServer.jamfcloud.com/api/v1/api-role-privileges/search


Search role privileges

Query Params
name
string
required
The partial or complete privilege name we are searching for

limit
string
Defaults to 15
Limit the query results, defaults to 15

15
Response

200
A list of matches Jamf API Roles Privileges

Response body
object
privileges
array of strings
required
length ≥ 0

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/api-role-privileges/search?limit=15' \
     --header 'accept: application/json'

{
  "privileges": [
    "Flush MDM Commands"
  ]
}