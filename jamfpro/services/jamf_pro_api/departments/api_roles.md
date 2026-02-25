Search for Departments
get
https://yourServer.jamfcloud.com/api/v1/departments

Search for Departments

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
Query in the RSQL format, allowing to filter department collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. Example: name=="department"

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
length between 1 and 225

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/departments?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "Department of Redundancy Department"
    }
  ]
}
-----
Create department record
post
https://yourServer.jamfcloud.com/api/v1/departments

Create department record

Body Params
department object to create. ids defined in this body will be ignored

name
string
required
length between 1 and 225
Department of Redundancy Department
Response

201
Department record was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/departments \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Department of Redundancy Department"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Deletes all departments by ids passed in body
post
https://yourServer.jamfcloud.com/api/v1/departments/delete-multiple


Deletes all departments by ids passed in body

Body Params
ids of departments to be deleted. pass in an array of ids

ids
array of strings

string

1

ADD string
Responses
204
All department ids passed in request sucessfully deleted.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/departments/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '{"ids":["1"]}'
-----
Get specified Department object
get
https://yourServer.jamfcloud.com/api/v1/departments/{id}

Gets specified Department object

Path Params
id
string
required
instance id of department record

Responses

200
Details of department were found

Response body
object
id
string
length ≥ 1
name
string
required
length between 1 and 225

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/departments/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Department of Redundancy Department"
}
-----
Update specified department object
put
https://yourServer.jamfcloud.com/api/v1/departments/{id}

Update specified department object

Path Params
id
string
required
instance id of department record

1
Body Params
department object to create. ids defined in this body will be ignored

name
string
required
length between 1 and 225
Department of Redundancy Department
Response

200
Department update

Response body
object
id
string
length ≥ 1
name
string
required
length between 1 and 225

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/departments/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Department of Redundancy Department"
}
'

{
  "id": "1",
  "name": "Department of Redundancy Department"
}
-----
Remove specified department record
delete
https://yourServer.jamfcloud.com/api/v1/departments/{id}

Removes specified department record

Path Params
id
string
required
instance id of department record

1
Response
204
Department record was deleted

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/departments/1

Get specified Department history object
get
https://yourServer.jamfcloud.com/api/v1/departments/{id}/history


Gets specified Department history object

Path Params
id
string
required
instance id of department history record

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
Details of department history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/departments//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Department history object notes
post
https://yourServer.jamfcloud.com/api/v1/departments/{id}/history


Adds specified Department history object notes

Path Params
id
string
required
instance id of department history record

Body Params
history notes to create

note
string
required
Responses

201
Notes of department history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/departments//history \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
