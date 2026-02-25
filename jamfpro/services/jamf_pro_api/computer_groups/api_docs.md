Returns the list of all computer groups
get
https://yourServer.jamfcloud.com/api/v1/computer-groups

Use it to get the list of all computer groups.

Response

200
Success

Response body
array of objects
object
id
string
length ≥ 1
name
string
description
string
smartGroup
boolean

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/computer-groups \
     --header 'accept: application/json'

[
  {
    "id": "1",
    "name": "All Managed Computers",
    "description": "A group containing all managed computers",
    "smartGroup": true
  }
]
-----
Get the membership of a Smart Computer Group
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-group-membership/{id}


Gets the membership of a Smart Computer Group

Path Params
id
string
required
id of the Smart Computer Group

Responses

200
Successful response - Smart Computer Group membership retrieved

Response body
object
members
array of integers

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-group-membership/ \
     --header 'accept: application/json'

{
  "members": [
    1,
    2,
    3
  ]
}
-----
Search for Smart Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups


Search for Smart Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter smart computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"

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
string
length ≥ 1
siteId
string
length ≥ 1
name
string
length ≥ 1
description
string
length ≥ 0
membershipCount
integer

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "siteId": "1",
      "name": "My computer group",
      "description": "My computer group description",
      "membershipCount": 231
    }
  ]
}

Create a Smart Computer Group
post
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups


Creates a Smart Computer Group

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
name
string
required
description
string
criteria
array of objects

ADD object
siteId
string | null
Response

201
Successful response - Smart Computer Group created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Update a Smart Computer Group
put
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Updates a Smart Computer Group

Path Params
id
string
required
id of target Smart Computer Group

Body Params
name
string
required
description
string
criteria
array of objects

ADD object
siteId
string | null
Responses

202
Successful response - Smart Computer Group updated

Response body
object
name
string
required
description
string
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
siteId
string | null

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "name": "New Group Name",
  "description": "New Group Description",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "siteId": "-1"
}
-----
Remove specified Smart Computer Group
delete
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Remove specified Smart Computer Group

Path Params
id
string
required
id of target Smart Computer Group

1
Responses
204
Successful response - Smart Computer Group removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/1 \
     --header 'accept: application/json'
-----
Get Smart Computer Group by Id
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/smart-groups/{id}


Get Smart Computer Group by Id

Path Params
id
string
required
instance id of smart computer group

Responses

200
Successful response

Response body
object
name
string
required
description
string
criteria
array of objects
object
name
string
required
priority
integer
andOr
string
required
searchType
string
required
value
string
required
openingParen
boolean
closingParen
boolean
siteId
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/smart-groups/ \
     --header 'accept: application/json'

{
  "name": "New Group Name",
  "description": "New Group Description",
  "criteria": [
    {
      "name": "Account",
      "priority": 0,
      "andOr": "and",
      "searchType": "is",
      "value": "test",
      "openingParen": false,
      "closingParen": false
    }
  ],
  "siteId": "-1"
}
_____
Search for Static Computer Groups
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups


Search for Static Computer Groups

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter static computer group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name, siteId. The siteId field can only be filtered by admins with full access. Any sited admin will have siteId filtered automatically. Example: name=="group"

Response

200
Successful response

Response body
object
totalCount
integer
required
≥ 0
results
array of objects
required
object
id
string
required
length ≥ 1
name
string
required
length ≥ 1
description
string | null
siteId
string | null
count
integer
≥ 0

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 2,
  "results": [
    {
      "id": "1",
      "name": "Test Static Computer Group",
      "description": "A test static computer group",
      "siteId": "1",
      "count": 5
    }
  ]
}
-----
Create membership of a static computer group.
post
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups


Create membership of a static computer group.

Query Params
platform
boolean
Defaults to false
Optional. Return platform identifiers instead of internal identifiers when set to true.


false
Body Params
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group


ADD string
Responses

201
Static computer group created successfully

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups?platform=false' \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
et Static Computer Group by Id
get
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Get Static Computer Group by Id

Path Params
id
string
required
instance id of static computer group

Responses

200
Successful response

Response body
object
id
string
required
length ≥ 1
name
string
required
length ≥ 1
description
string | null
siteId
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Test Static Computer Group",
  "description": "A test static computer group",
  "siteId": "1"
}
-----
Update membership of a static computer group.
put
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Update membership of a static computer group.

Path Params
id
string
required
instance id of a static computer group

Body Params
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group


ADD string
Responses

202
Successful response

Response body
object
id
string
name
string
required
length ≥ 1
description
string | null
siteId
string | null
assignments
array of strings
Array of computer IDs to assign to the static group

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "name": "Test Static Computer Group",
  "description": "A test static computer group",
  "siteId": "1",
  "assignments": [
    "1"
  ]
}
-----
Remove Static Computer Group by Id
delete
https://yourServer.jamfcloud.com/api/v2/computer-groups/static-groups/{id}


Remove Static Computer Group by Id

Path Params
id
string
required
instance id of static computer group

1
Responses
204
Static Computer Group successfully removed

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v2/computer-groups/static-groups/1 \
     --header 'accept: application/json'

     