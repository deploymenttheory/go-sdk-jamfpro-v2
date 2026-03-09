Get information about all Cloud Identity Providers configurations.
get
https://yourServer.jamfcloud.com/api/v1/cloud-idp

Returns basic informations about all configured Cloud Identity Provider.

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
Defaults to id:desc
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:desc

ADD string
Response

200
Cloud Identity Provider configurations informations returned.

Response body
object
totalCount
integer
results
array of objects
object
id
string
≥ 1001
displayName
string
enabled
boolean
providerName
string
enum
GOOGLE AZURE

providerDescription
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-idp?page=0&page-size=100&sort=id%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1001",
      "displayName": "Cloud Identity Provider",
      "enabled": true,
      "providerName": "PROVIDER",
      "providerDescription": "Entra ID"
    }
  ]
}
-----

Export Cloud Identity Providers collection
post
https://yourServer.jamfcloud.com/api/v1/cloud-idp/export

Export Cloud Identity Providers collection

Query Params
export-fields
array of strings
Defaults to
Export fields parameter, used to change default order or ignore some of the response properties. Default is empty array, which means that all fields of the response entity will be serialized. Example: export-fields=id,username


ADD string
export-labels
array of strings
Defaults to
Export labels parameter, used to customize fieldnames/columns in the exported file. Default is empty array, which means that response properties names will be used. Number of the provided labels must match the number of export-fields Example: export-labels=identifier,name with matching: export-fields=id,username


ADD string
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
Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be seperated with a comma. Example: sort=id:desc,name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name=="department"

Body Params
Optional. Override query parameters since they can make URI exceed 2,000 character limit.

page
integer | null
Defaults to 0
0
pageSize
integer | null
Defaults to 100
100
sort
array of strings | null
Defaults to id:desc
Sorting criteria in the format: [[:asc/desc]. Default direction when not stated is ascending.


string

id:desc

ADD string
filter
string | null
fields
array of objects | null
Used to change default order or ignore some of the fields. When null or empty array, all fields will be exported.


ADD object
Headers
accept
string
enum
Defaults to application/json
Generated from available response content types


text/csv
Allowed:

application/json

text/csv
Responses

200
Export successful

Response body
json

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-idp/export?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: text/csv' \
     --header 'content-type: application/json' \
     --data '
{
  "page": 0,
  "pageSize": 100,
  "sort": [
    "id:desc"
  ]
}
'

Name,"Provider Name",Status
"Google LDAP",GOOGLE,true
"Azure LDAP",AZURE,true
-----
Get Cloud Identity Provider configuration with given ID.
get
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}

Get Cloud Identity Provider configuration with given ID.

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

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-idp/ \
     --header 'accept: application/json'

{
  "id": "1001",
  "displayName": "Cloud Identity Provider",
  "providerName": "PROVIDER"
}
-----

Get Cloud Identity Provider history
get
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/history

Gets specified Cloud Identity Provider object history

Path Params
id
string
required
Cloud Identity Provider identifier

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
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of Cloud Identity Provider history were found

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
integer
≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/cloud-idp//history?page=0&page-size=100&sort=date%3Adesc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": 1,
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}

Add Cloud Identity Provider history note
post
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/history

Adds specified Cloud Identity Provider object history notes

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
history notes to create

note
string
required
Responses

201
Notes of Cloud Identity Provider history were added

Response body
object
id
integer
≥ 1
username
string
date
string
note
string
details
string | null

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-idp//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}

-----

Get group test search
post
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/test-group

Do test search to ensure about configuration and mappings

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Search request

groupname
string
required
Responses

200
Cloud Identity Provider test search result returned.

Response body
object
totalCount
integer
results
array of objects
object
distinguishedName
string
id
string
uuid
string
serverId
string
name
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-idp/1/test-group \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "groupname": "users"
}
'

{
  "totalCount": 1,
  "results": [
    {
      "distinguishedName": "cn=users,ou=Groups,dc=jamf,dc=com",
      "id": "users",
      "uuid": "121100023",
      "serverId": "1001",
      "name": "users"
    }
  ]
}
-----

Get user test search
post
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/test-user

Do test search to ensure about configuration and mappings

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Search request

username
string
required
admin
Responses

200
Cloud Identity Provider test search result returned.

Response body
object
totalCount
integer
results
array of objects
object
distinguishedName
string
id
string
uuid
string
serverId
string
name
string
attributes
object

attributes object
fullName
string
emailAddress
string
phoneNumber
string
position
string
room
string
buildingName
string
departmentName
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-idp//test-user \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"username":"admin"}'

{
  "totalCount": 1,
  "results": [
    {
      "distinguishedName": "uid=admin,ou=Users,dc=jamf,dc=com",
      "id": "admin@jamf.com",
      "uuid": "admin",
      "serverId": "1001",
      "name": "admin",
      "attributes": {
        "fullName": "Bob",
        "emailAddress": "bob@jamf.com",
        "phoneNumber": "123456789",
        "position": "SE",
        "room": "1",
        "buildingName": "Jamf",
        "departmentName": "Engineering"
      }
    }
  ]
}

-----

Get membership test search
post
https://yourServer.jamfcloud.com/api/v1/cloud-idp/{id}/test-user-membership

Do test search to ensure about configuration and mappings

Path Params
id
string
required
Cloud Identity Provider identifier

Body Params
Search request

username
string
required
groupname
string
required
Responses

200
Cloud Identity Provider test search result returned.

Response body
object
username
string
isMember
boolean

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/cloud-idp//test-user-membership \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "username": "admin",
  "isMember": true
}