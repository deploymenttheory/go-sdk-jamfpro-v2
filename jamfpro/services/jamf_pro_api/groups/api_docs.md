Returns group information for all Mobile Device and Computer groups
get
https://yourServer.jamfcloud.com/api/v1/groups

Returns group information for all Mobile Device and Computer groups. The type of groups returned will be dependent upon the corresponding group type READ privileges. Results can be sorted by name, description, group type, or isSmart. Default sorting is by group name in ascending order.

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
Defaults to groupName:asc
Sorting criteria in the format: property:asc/desc. Default sort is groupName:asc. Multiple sort criteria are supported and must be separated with a comma. Fields allowed in sorting: groupName, groupDescription, groupType, isSmart. Example: sort=groupName:asc,groupType:desc


string

groupName:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter group collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: groupName, groupDescription, groupType, isSmart. This param can be combined with paging and sorting. When using groupType in the filter, the value must be either "MOBILE" or "COMPUTER" but not both. When using groupType in the filter, the value is case sensitive. When using groupType in the filter, it will exclude groups of the other type regardless of or/and conditionals. Example: filter=groupName=="Managed" and isSmart=="true" Example: filter=groupType=="COMPUTER" and groupDescription=="Admin"

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
groupPlatformId
string
groupJamfProId
string
groupName
string
groupDescription
string
groupType
string
enum
MOBILE COMPUTER

smart
boolean
membershipCount
integer
≥ 0

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/groups?page=0&page-size=100&sort=groupName%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 2,
  "results": [
    {
      "groupPlatformId": "56585788-8937-4bdf-b986-458048b1c36c",
      "groupJamfProId": "1",
      "groupName": "New Group Name",
      "groupDescription": "New Group Description",
      "groupType": "MOBILE",
      "smart": true,
      "membershipCount": 3
    }
  ]
}
-----
Returns group information for the given platform UUID
get
https://yourServer.jamfcloud.com/api/v1/groups/{id}

Returns group information for the given platform UUID. Dependent upon the returned group type the corresponding READ privilege for that group type will be needed.

Path Params
id
string
required
The platform UUID of a group

Responses

200
Successful response

Response body
object
groupPlatformId
string
groupJamfProId
string
groupName
string
groupDescription
string
groupType
string
enum
MOBILE COMPUTER

smart
boolean
membershipCount
integer
≥ 0
criteria
array of objects | null
object
name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

priority
integer
required
≥ 0
The priority order of this criterion

andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

value
string
required
The value to search for

openingParen
boolean
Whether to add an opening parenthesis before this criterion

closingParen
boolean
Whether to add a closing parenthesis after this criterion

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/groups/ \
     --header 'accept: application/json'

{
  "groupPlatformId": "20b66625-39c0-41d9-9b2c-98e57a3554c8",
  "groupJamfProId": "18",
  "groupName": "Smart Mobile Device Group",
  "groupDescription": "Devices with iOS 15 or higher",
  "groupType": "MOBILE",
  "smart": true,
  "membershipCount": 12,
  "criteria": [
    {
      "name": "Operating System Version",
      "priority": 0,
      "andOr": "and",
      "searchType": "greater than or equal",
      "value": "15.0",
      "openingParen": false,
      "closingParen": false
    },
    {
      "name": "Device Name",
      "priority": 1,
      "andOr": "and",
      "searchType": "like",
      "value": "iPhone",
      "openingParen": false,
      "closingParen": false
    }
  ]
}
-----
Delete a group by platform UUID
delete
https://yourServer.jamfcloud.com/api/v1/groups/{id}

Deletes a group by its platform UUID. Returns a 400 error if the group is being used as a dependency. Requires appropriate DELETE privileges.

Path Params
id
string
required
The platform UUID of a group

Responses
204
Group deleted successfully

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/groups/ \
     --header 'accept: application/json'

-----
Update a group by platform UUID
patch
https://yourServer.jamfcloud.com/api/v1/groups/{id}

Updates a group by its platform UUID. For both smart and static groups, groupName and groupDescription can be updated. For smart groups, criteria can also be updated. For static groups, assignments can also be updated. Requires appropriate UPDATE privileges.

Path Params
id
string
required
The platform UUID of a group

123
Body Params
groupName
string
Updated Group Name
groupDescription
string
Updated Group Description
criteria
array of objects

object

name
string
required
length ≥ 1
The field to search on (e.g., Model, OS Version, etc.)

Model
priority
integer
required
≥ 0
The priority order of this criterion

1
andOr
string
required
length ≥ 1
Whether this criterion should be ANDed or ORed with the previous criterion

and
searchType
string
required
length ≥ 1
The type of search to perform (e.g., is, is not, like, etc.)

is
value
string
required
The value to search for

iPad
openingParen
boolean
Whether to add an opening parenthesis before this criterion


true
closingParen
boolean
Whether to add a closing parenthesis after this criterion


true

ADD object
assignments
array of objects | null

object

deviceId
string
required
length between 1 and 50
56585788-8937-4bdf-b986-458048b1c36c
selected
boolean
required

true

ADD object
Responses
204
Group updated successfully

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/groups/123 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupName": "Updated Group Name",
  "groupDescription": "Updated Group Description",
  "assignments": [
    {
      "selected": true,
      "deviceId": "56585788-8937-4bdf-b986-458048b1c36c"
    }
  ],
  "criteria": [
    {
      "name": "Model",
      "priority": 1,
      "andOr": "and",
      "searchType": "is",
      "value": "iPad",
      "openingParen": true,
      "closingParen": true
    }
  ]
}
'
