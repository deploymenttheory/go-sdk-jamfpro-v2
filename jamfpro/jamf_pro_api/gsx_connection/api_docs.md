Finds the Jamf Pro GSX Connection information
get
https://yourServer.jamfcloud.com/api/v1/gsx-connection

Finds the Jamf Pro GSX Connection information

Response

200
Success

Response body
object
enabled
boolean
required
Defaults to false
username
string
required
Defaults to
serviceAccountNo
string
required
length ≤ 10
shipToNo
string
length ≤ 10
gsxKeystore
object
required
name
string
required
Defaults to
expirationEpoch
int64
errorMessage
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection \
     --header 'accept: application/json'

{
  "enabled": true,
  "username": "exampleEmail@example.com",
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "gsxKeystore": {
    "name": "certificate.p12",
    "expirationEpoch": 169195490000,
    "errorMessage": "Certificate error"
  }
}
-----
Updates Jamf Pro GSX Connection information
put
https://yourServer.jamfcloud.com/api/v1/gsx-connection

Updates Jamf Pro GSX Connection information

Body Params
GSX Connection to update

enabled
boolean
required
Defaults to false

false
username
string
required
Defaults to
exampleEmail@example.com
serviceAccountNo
string
required
length ≤ 10
0000012345
shipToNo
string
length ≤ 10
0000012345
token
string
required
34dsg23-5dsgs-3sdg-4ffs-435sdgs
gsxKeystore
object
required

gsxKeystore object
name
string
required
Defaults to
certificate.p12
keystoreBytes
string
The base 64 encoded of the GSX Connection keystore.

WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09
keystorePassword
password
required
••••••••
Responses

200
Success

Response body
object
enabled
boolean
required
Defaults to false
username
string
required
Defaults to
serviceAccountNo
string
required
length ≤ 10
shipToNo
string
length ≤ 10
gsxKeystore
object
required
name
string
required
Defaults to
expirationEpoch
int64
errorMessage
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json' \
     --data '
{
  "enabled": false,
  "username": "exampleEmail@example.com",
  "gsxKeystore": {
    "name": "certificate.p12",
    "keystoreBytes": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
    "keystorePassword": "test1234"
  },
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "token": "34dsg23-5dsgs-3sdg-4ffs-435sdgs"
}
'

{
  "enabled": true,
  "username": "exampleEmail@example.com",
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "gsxKeystore": {
    "name": "certificate.p12",
    "expirationEpoch": 169195490000,
    "errorMessage": "Certificate error"
  }
}
-----
Updates Jamf Pro GSX Connection information
patch
https://yourServer.jamfcloud.com/api/v1/gsx-connection

Updates Jamf Pro GSX Connection information

Body Params
GSX Connection to update

enabled
boolean
Defaults to false

false
username
string
Defaults to
exampleEmail@example.com
serviceAccountNo
string
length ≤ 10
0000012345
shipToNo
string
length ≤ 10
0000012345
token
string
34dsg23-5dsgs-3sdg-4ffs-435sdgs
gsxKeystore
object

gsxKeystore object
name
string
required
Defaults to
certificate.p12
keystoreBytes
string
The base 64 encoded of the GSX Connection keystore.

WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09
keystorePassword
password
required
••••••••
Responses

200
Success

Response body
object
enabled
boolean
required
Defaults to false
username
string
required
Defaults to
serviceAccountNo
string
required
length ≤ 10
shipToNo
string
length ≤ 10
gsxKeystore
object
required
name
string
required
Defaults to
expirationEpoch
int64
errorMessage
string

curl --request PATCH \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection \
     --header 'accept: application/json' \
     --header 'content-type: application/merge-patch+json' \
     --data '
{
  "enabled": false,
  "username": "exampleEmail@example.com",
  "gsxKeystore": {
    "name": "certificate.p12",
    "keystoreBytes": "WlhoaGJYQnNaU0J2WmlCaElHSmhjMlUyTkNCbGJtTnZaR1ZrSUhaaGJHbGtJSEF4TWk0Z2EyVjVjM1J2Y21VZ1ptbHNaUT09",
    "keystorePassword": "test1234"
  },
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "token": "34dsg23-5dsgs-3sdg-4ffs-435sdgs"
}
'

{
  "enabled": true,
  "username": "exampleEmail@example.com",
  "serviceAccountNo": "0000012345",
  "shipToNo": "0000012345",
  "gsxKeystore": {
    "name": "certificate.p12",
    "expirationEpoch": 169195490000,
    "errorMessage": "Certificate error"
  }
}
-----
Get specified GSX Connection History object
get
https://yourServer.jamfcloud.com/api/v1/gsx-connection/history

Gets specified GSX Connection history object

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
Details of GSX Connection history were found

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
length ≥ 1
username
string
date
string
note
string
details
string | null

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection/history \
     --header 'accept: application/json'

     {
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "username": "admin",
      "date": "2019-02-04T21:09:31.661Z",
      "note": "Sso settings update",
      "details": "Is SSO Enabled false\\nSelected SSO Provider"
    }
  ]
}
-----
Add specified GSX Connection history object notes
post
https://yourServer.jamfcloud.com/api/v1/gsx-connection/history

Adds specified GSX Connection history object notes

Body Params
history notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of GSX Connection history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection/history \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "note": "A generic note can sometimes be useful, but generally not."
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Test functionality of an GSX Connection
post
https://yourServer.jamfcloud.com/api/v1/gsx-connection/test

Test functionality of an GSX Connection

Responses
202
Success

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/gsx-connection/test \
     --header 'accept: application/json'