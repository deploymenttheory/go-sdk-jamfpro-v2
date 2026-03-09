Updates Activation Code
put
https://yourServer.jamfcloud.com/api/v1/activation-code

Updates Activation Code in Jamf Pro.

Body Params
activationCode
string
required
length between 32 and 39
Activation Code for Jamf Pro. Hyphens are optional.

Response
202
Update successful.

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/activation-code \
     --header 'content-type: application/json'

-----

Updates Organization Name
patch
https://yourServer.jamfcloud.com/api/v1/activation-code/organization-name


Updates Organization Name in Jamf Pro.

Body Params
organizationName
string
required
The Organization Name for Jamf Pro.

Response
202
Update successful.

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/activation-code/organization-name \
     --header 'content-type: application/json'
-----

Get Activation Code history object
get
https://yourServer.jamfcloud.com/api/v1/activation-code/history


Get Activation Code history object

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
Sorting criteria in the format: property:asc/desc. Default sort is date:desc. Fields allowed in the query: id, username, date, note, details Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,note:asc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default filter is empty query - returning all results for the requested page. Fields allowed in the query: id, username, date, note, details. This param can be combined with paging and sorting. Example: filter=username!=admin and details==disabled and date<2019-12-15

Responses

200
Details of Activation Code history was found.

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
     --url 'https://yourserver.jamfcloud.com/api/v1/activation-code/history?page=0&page-size=100&sort=date%3Adesc' \
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

Add Activation Code object note
post
https://yourServer.jamfcloud.com/api/v1/activation-code/history


Adds Activation Code object note.

Body Params
Activation Code history notes to create.

note
string
required
Responses

201
Notes of Activation Code history were added.

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
     --url https://yourserver.jamfcloud.com/api/v1/activation-code/history \
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

Export history object collection in specified format for Activation Code
post
https://yourServer.jamfcloud.com/api/v1/activation-code/history/export


Export history object collection in specified format for Activation Code

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


application/json
Allowed:

application/json

text/csv
Responses

curl --request POST \
     --url 'https://yourserver.jamfcloud.com/api/v1/activation-code/history/export?page=0&page-size=100&sort=date%3Adesc' \
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
