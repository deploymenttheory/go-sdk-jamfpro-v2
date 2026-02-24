Get the current Jamf API Roles
get
https://yourServer.jamfcloud.com/api/v1/api-roles

Get roles with Search Criteria

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
Defaults to id:asc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in the query: id, displayName. Example: sort=displayName:desc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter app titles collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, displayName. Example: displayName=="myRole"

Response

200
A list of the current Jamf API Roles

Response body
object
totalCount
integer
required
≥ 0
results
array of objects
required
length ≥ 0
object
id
string
required
displayName
string
required
privileges
array of strings
required
length ≥ 0

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/api-roles?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "displayName": "One Role to Rule them all",
      "privileges": [
        "View License Serial Numbers"
      ]
    }
  ]
}

-----

Create a new API role
post
https://yourServer.jamfcloud.com/api/v1/api-roles

Post to create new Role

Body Params
API Integrations Role to create

displayName
string
required
privileges
array of strings
required
length ≥ 0

ADD string
Responses

200
The created Role

Response body
object
id
string
required
displayName
string
required
privileges
array of strings
required
length ≥ 0

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/api-roles \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "displayName": "One Role to Rule them all",
  "privileges": [
    "View License Serial Numbers"
  ]
}

-----

Get the specific Jamf API Role
get
https://yourServer.jamfcloud.com/api/v1/api-roles/{id}

Get specific Role

Path Params
id
string
required
instance id of API role

Responses

200
The requested Jamf API Role

Response body
object
id
string
required
displayName
string
required
privileges
array of strings
required
length ≥ 0

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/api-roles/ \
     --header 'accept: application/json'

{
  "id": "1",
  "displayName": "One Role to Rule them all",
  "privileges": [
    "View License Serial Numbers"
  ]
}

-----

Update API Integrations Role
put
https://yourServer.jamfcloud.com/api/v1/api-roles/{id}

Update specific Role

Path Params
id
string
required
instance id of API role

Body Params
API Integrations Role to update

displayName
string
required
privileges
array of strings
required
length ≥ 0

ADD string
Responses

200
Jamf API Integrations Role updated

Response body
object
id
string
required
displayName
string
required
privileges
array of strings
required
length ≥ 0

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/api-roles/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "displayName": "One Role to Rule them all",
  "privileges": [
    "View License Serial Numbers"
  ]
}

-----

Delete API Integrations Role
delete
https://yourServer.jamfcloud.com/api/v1/api-roles/{id}

Delete specific Role

Path Params
id
string
required
instance id of API role

Responses
204
Jamf API Integrations Role deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/api-roles/ \
     --header 'accept: application/json'