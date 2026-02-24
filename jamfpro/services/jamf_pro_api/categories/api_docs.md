Get Category objects
get
https://yourServer.jamfcloud.com/api/v1/categories

Gets Category objects.

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter categories collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, priority. This param can be combined with paging and sorting. Example: filter=name=="Apps*" and priority>=5

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
name
string
required
priority
int32
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/categories?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "The Best",
      "priority": 9
    }
  ]
}

-----

Create Category record
post
https://yourServer.jamfcloud.com/api/v1/categories

Create category record

Body Params
category object to create. IDs defined in this body will be ignored

name
string
required
priority
int32
required
Response

201
Category record was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/categories \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "The Best",
  "priority": 9
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

-----

Delete multiple Categories by their IDs
post
https://yourServer.jamfcloud.com/api/v1/categories/delete-multiple


Delete multiple Categories by their IDs

Body Params
IDs of the categories to be deleted

ids
array of strings

string

1,2

ADD string
Responses
204
All Category IDs passed in request sucessfully deleted.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/categories/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1,2"]}'

-----

Get specified Category object
get
https://yourServer.jamfcloud.com/api/v1/categories/{id}

Gets specified Category object

Path Params
id
string
required
instance id of category record

Responses

200
Details of category were found

Response body
object
id
string
length ≥ 1
name
string
required
priority
int32
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/categories/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "The Best",
  "priority": 9
}

-----

Update specified Category object
put
https://yourServer.jamfcloud.com/api/v1/categories/{id}

Update specified category object

Path Params
id
string
required
instance id of category record

1
Body Params
category object to create. id defined in this body will be ignored

name
string
required
The Best
priority
int32
required
9
Response

200
Category record was updated

Response body
object
id
string
length ≥ 1
name
string
required
priority
int32
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/categories/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "The Best",
  "priority": 9
}
'

{
  "id": "1",
  "name": "The Best",
  "priority": 9
}

-----

Remove specified Category record
delete
https://yourServer.jamfcloud.com/api/v1/categories/{id}

Removes specified category record

Path Params
id
string
required
instance id of category record

1
Response
204
Category record was deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/categories/1

-----

Get specified Category history object
get
https://yourServer.jamfcloud.com/api/v1/categories/{id}/history


Gets specified Category history object

Path Params
id
string
required
instance id of category history record

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of category history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/categories//history?page=0&page-size=100&sort=date%3Adesc' \
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
-----

Add specified Category history object notes
post
https://yourServer.jamfcloud.com/api/v1/categories/{id}/history


Adds specified Category history object notes

Path Params
id
string
required
instance id of category history record

Body Params
history notes to create

note
string
required
Responses

201
Notes of category history were added

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
     --url https://yourserver.jamfcloud.com/api/v1/categories//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": 1,
  "username": "admin",
  "date": "2019-02-04T21:09:31.661Z",
  "note": "Sso settings update",
  "details": "Is SSO Enabled false\\nSelected SSO Provider"
}