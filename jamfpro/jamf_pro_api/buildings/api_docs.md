Search for sorted and paged Buildings
get
https://yourServer.jamfcloud.com/api/v1/buildings

Search for sorted and paged buildings

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
Query in the RSQL format, allowing to filter buildings collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: name, streetAddress1, streetAddress2, city, stateProvince, zipPostalCode, country. This param can be combined with paging and sorting. Example: filter=city=="Chicago" and name=="build"

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
object
id
string
name
string
required
streetAddress1
string | null
streetAddress2
string | null
city
string | null
stateProvince
string | null
zipPostalCode
string | null
country
string | null

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/buildings?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 3,
  "results": [
    {
      "id": "1",
      "name": "Apple Park",
      "streetAddress1": "The McIntosh Tree",
      "streetAddress2": "One Apple Park Way",
      "city": "Cupertino",
      "stateProvince": "California",
      "zipPostalCode": "95014",
      "country": "The United States of America"
    }
  ]
}

-----

Create Building record
post
https://yourServer.jamfcloud.com/api/v1/buildings

Create building record

Body Params
building object to create. ids defined in this body will be ignored

name
string
required
streetAddress1
string | null
streetAddress2
string | null
city
string | null
stateProvince
string | null
zipPostalCode
string | null
country
string | null
Response

201
Building record was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/buildings \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}

-----

Delete multiple Buildings by their ids
post
https://yourServer.jamfcloud.com/api/v1/buildings/delete-multiple


multiple many Buildings by their ids

Body Params
ids of the building to be deleted

ids
array of strings

string


ADD string
Responses
204
All building ids passed in request sucessfully deleted.

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/buildings/delete-multiple \
     --header 'accept: application/json' \
     --header 'content-type: application/json'
-----

Export Buildings collection
post
https://yourServer.jamfcloud.com/api/v1/buildings/export

Export Buildings collection

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
Sorting criteria in the format: property:asc/desc. Default sort is id:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,name:asc


string

id:asc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, name. This param can be combined with paging and sorting. Example: name=="buildings"

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
     --url 'https://yourserver.jamfcloud.com/api/v1/buildings/export?page=0&page-size=100&sort=id%3Aasc' \
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

id,name
1,Marketing
2,Accounting

-----

Get specified Building object
get
https://yourServer.jamfcloud.com/api/v1/buildings/{id}

Gets specified Building object

Path Params
id
string
required
instance id of building record

Responses

200
Details of building were found

Response body
object
id
string
name
string
required
streetAddress1
string | null
streetAddress2
string | null
city
string | null
stateProvince
string | null
zipPostalCode
string | null
country
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/buildings/ \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Apple Park",
  "streetAddress1": "The McIntosh Tree",
  "streetAddress2": "One Apple Park Way",
  "city": "Cupertino",
  "stateProvince": "California",
  "zipPostalCode": "95014",
  "country": "The United States of America"
}

-----

Update specified Building object
put
https://yourServer.jamfcloud.com/api/v1/buildings/{id}

Update specified building object

Path Params
id
string
required
instance id of building record

Body Params
building object to update. ids defined in this body will be ignored

name
string
required
streetAddress1
string | null
streetAddress2
string | null
city
string | null
stateProvince
string | null
zipPostalCode
string | null
country
string | null
Response

200
Building update

Response body
object
id
string
name
string
required
streetAddress1
string | null
streetAddress2
string | null
city
string | null
stateProvince
string | null
zipPostalCode
string | null
country
string | null

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/buildings/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "name": "Apple Park",
  "streetAddress1": "The McIntosh Tree",
  "streetAddress2": "One Apple Park Way",
  "city": "Cupertino",
  "stateProvince": "California",
  "zipPostalCode": "95014",
  "country": "The United States of America"
}

-----

Remove specified Building record
delete
https://yourServer.jamfcloud.com/api/v1/buildings/{id}

Removes specified building record

Path Params
id
string
required
instance id of building record

Response

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/buildings/

-----

Get specified Building History object
get
https://yourServer.jamfcloud.com/api/v1/buildings/{id}/history

Gets specified Building history object

Path Params
id
string
required
instance id of building history record

1
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
Details of building history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/buildings/1/history?page=0&page-size=100&sort=date%3Adesc' \
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
Add specified Building history object notes
post
https://yourServer.jamfcloud.com/api/v1/buildings/{id}/history


Adds specified Building history object notes

Path Params
id
string
required
instance id of building history record

Body Params
history notes to create

note
string
required
Responses

201
Notes of building history were added

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
     --url https://yourserver.jamfcloud.com/api/v1/buildings//history \
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

Export history object collection in specified format for specified Buildings
post
https://yourServer.jamfcloud.com/api/v1/buildings/{id}/history/export


Export history object collection in specified format for specified Buildings

Path Params
id
string
required
instance id of buildings

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
Defaults to date:desc
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

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
     --url 'https://yourserver.jamfcloud.com/api/v1/buildings//history/export?page=0&page-size=100&sort=date%3Adesc' \
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

id,username,date,note,details
1,admin,2019-02-04 21:09:31,Buildings update,Some details
