Read all sorted and paged Device Enrollment instances
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments

Search for sorted and paged device enrollment instances

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
Sorting criteria in the format: property:asc/desc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=date:desc,name:asc


string

id:asc

ADD string
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
supervisionIdentityId
string
siteId
string
serverName
string
serverUuid
string
adminId
string
orgName
string
orgEmail
string
orgPhone
string
orgAddress
string
tokenExpirationDate
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/device-enrollments?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "name": "Example Device Enrollment Instance",
      "supervisionIdentityId": "1",
      "siteId": "-1",
      "serverName": "Acme ASM",
      "serverUuid": "BASD08C11F3C455",
      "adminId": "admin1234",
      "orgName": "Acme Enterprises",
      "orgEmail": "admin@test.com",
      "orgPhone": "555-0123",
      "orgAddress": "124 Conch Street, Bikini Bottom, Pacific Ocean",
      "tokenExpirationDate": "2000-10-30"
    }
  ]
}
-----
Retrieve the Jamf Pro Device Enrollment public key
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/public-key


Retrieve the Jamf Pro device enrollment public key

Responses

200
Success

Response body
file
404
Not Found

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/public-key \
     --header 'accept: application/x-pem-file'
-----
Get all instance sync states for all Device Enrollment Instances
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/syncs


Get all instance sync states for all instances

Response

200
Successful response

Response body
array of objects
object
syncState
string
instanceId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/syncs \
     --header 'accept: application/json'

[
  {
    "syncState": "CONNECTION_ERROR",
    "instanceId": "1",
    "timestamp": "2019-04-17T14:08:06.706+0000"
  }
]
-----
Create a Device Enrollment Instance with the supplied Token
post
https://yourServer.jamfcloud.com/api/v1/device-enrollments/upload-token


Creates a device enrollment instance with the supplied token.

Body Params
The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance.

tokenFileName
string
Optional name of the token to be saved, if no name is provided one will be auto-generated

Acme MDM Token
encodedToken
string
The base 64 encoded token

VTI5dFpTQnlZVzVrYjIwZ1ltbDBJRzltSUhSbGVIUWdkRzhnZFhObElHRnVaQ0J6WldVZ2FXWWdZVzU1YjI1bElHRmpkSFZoYkd4NUlIUnlhV1Z6SUhSdklHUmxZMjlrWlNCcGRBPT0=
Responses

201
Device Enrollment Instance was created

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/upload-token \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "tokenFileName": "Acme MDM Token",
  "encodedToken": "VTI5dFpTQnlZVzVrYjIwZ1ltbDBJRzltSUhSbGVIUWdkRzhnZFhObElHRnVaQ0J6WldVZ2FXWWdZVzU1YjI1bElHRmpkSFZoYkd4NUlIUnlhV1Z6SUhSdklHUmxZMjlrWlNCcGRBPT0="
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Retrieve a Device Enrollment Instance with the supplied id
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}


Retrieves a Device Enrollment Instance with the supplied id

Path Params
id
string
required
Device Enrollment Instance identifier

1
Responses

200
Success

Response body
object
id
string
length ≥ 1
name
string
required
supervisionIdentityId
string
siteId
string
serverName
string
serverUuid
string
adminId
string
orgName
string
orgEmail
string
orgPhone
string
orgAddress
string
tokenExpirationDate
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "name": "Example Device Enrollment Instance",
  "supervisionIdentityId": "1",
  "siteId": "-1",
  "serverName": "Acme ASM",
  "serverUuid": "BASD08C11F3C455",
  "adminId": "admin1234",
  "orgName": "Acme Enterprises",
  "orgEmail": "admin@test.com",
  "orgPhone": "555-0123",
  "orgAddress": "124 Conch Street, Bikini Bottom, Pacific Ocean",
  "tokenExpirationDate": "2000-10-30"
}
-----
Update a Device Enrollment Instance with the supplied id
put
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}


Updates a Device Enrollment Instance with the supplied id

Path Params
id
string
required
Device Enrollment Instance identifier

1
Body Params
name
string
required
Example Device Enrollment Instance
supervisionIdentityId
string
1
siteId
string
-1
Responses

200
Success

Response body
object
id
string
length ≥ 1
name
string
required
supervisionIdentityId
string
siteId
string
serverName
string
serverUuid
string
adminId
string
orgName
string
orgEmail
string
orgPhone
string
orgAddress
string
tokenExpirationDate
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "name": "Example Device Enrollment Instance",
  "supervisionIdentityId": "1",
  "siteId": "-1"
}
'

{
  "id": "1",
  "name": "Example Device Enrollment Instance",
  "supervisionIdentityId": "1",
  "siteId": "-1",
  "serverName": "Acme ASM",
  "serverUuid": "BASD08C11F3C455",
  "adminId": "admin1234",
  "orgName": "Acme Enterprises",
  "orgEmail": "admin@test.com",
  "orgPhone": "555-0123",
  "orgAddress": "124 Conch Street, Bikini Bottom, Pacific Ocean",
  "tokenExpirationDate": "2000-10-30"
}

Delete a Device Enrollment Instance with the supplied id
delete
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}


Deletes a Device Enrollment Instance with the supplied id

Path Params
id
string
required
Device Enrollment Instance identifier

Responses
204
Success

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/ \
     --header 'accept: application/json'
-----
Disown devices from the given Device Enrollment Instance
post
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/disown


Disowns devices from the given device enrollment instance

Path Params
id
string
required
Device Enrollment Instance identifier

Body Params
List of device serial numbers to disown

devices
array of strings

ADD string
Responses

200
Success

Response body
object
devices
object
Has additional fields

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments//disown \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "devices": {
    "a2s3d4f5": "SUCCESS",
    "0o9i8u7y6t": "FAILED"
  }
}
-----
Get sorted and paged Device Enrollment history objects
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/history


Gets sorted and paged device enrollment history objects

Path Params
id
string
required
Device Enrollment Instance identifier

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
Sorting criteria in the format: property,asc/desc. Default sort order is descending. Multiple sort criteria are supported and must be entered on separate lines in Swagger UI. In the URI the 'sort' query param is duplicated for each sort criterion, e.g., ...&sort=name%2Casc&sort=date%2Cdesc


string

date:desc

ADD string
filter
string
Query in the RSQL format, allowing to filter history notes collection. Default search is empty query - returning all results for the requested page. Fields allowed in the query: username, date, note, details. This param can be combined with paging and sorting. Example: search=username!=admin and details==disabled and date<2019-12-15

Response

200
Details of device enrollment history were found

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
     --url 'https://yourserver.jamfcloud.com/api/v1/device-enrollments//history?page=0&page-size=100&sort=date%3Adesc' \
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
Add Device Enrollment history object notes
post
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/history


Adds device enrollment history object notes

Path Params
id
string
required
Device Enrollment Instance identifier

2
Body Params
History notes to create

note
string
required
A generic note can sometimes be useful, but generally not.
Responses

201
Notes of device enrollment history were added

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/2/history \
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
Get all instance sync states for a single Device Enrollment Instance
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/syncs


Get all instance sync states for a single instance

Path Params
id
string
required
Device Enrollment Instance identifier

1
Response

200
Successful response

Response body
array of objects
object
syncState
string
instanceId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/1/syncs \
     --header 'accept: application/json'

[
  {
    "syncState": "CONNECTION_ERROR",
    "instanceId": "1",
    "timestamp": "2019-04-17T14:08:06.706+0000"
  }
]
-----
Get the latest sync state for a single Device Enrollment Instance
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/syncs/latest


Get the latest sync state for a single instance

Path Params
id
string
required
Device Enrollment Instance identifier

1
Response

200
Successful response

Response body
object
syncState
string
instanceId
string
timestamp
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/1/syncs/latest \
     --header 'accept: application/json'

{
  "syncState": "CONNECTION_ERROR",
  "instanceId": "1",
  "timestamp": "2019-04-17T14:08:06.706+0000"
}
-----
Update a Device Enrollment Instance with the supplied Token
put
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/upload-token


Updates a device enrollment instance with the supplied token.

Path Params
id
string
required
Device Enrollment Instance identifier

Body Params
The downloaded token base 64 encoded from the MDM server to be used to create a new Device Enrollment Instance.

tokenFileName
string
Optional name of the token to be saved, if no name is provided one will be auto-generated

encodedToken
string
The base 64 encoded token

Responses

200
Success

Response body
object
id
string
length ≥ 1
name
string
required
supervisionIdentityId
string
siteId
string
serverName
string
serverUuid
string
adminId
string
orgName
string
orgEmail
string
orgPhone
string
orgAddress
string
tokenExpirationDate
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments//upload-token \
     --header 'accept: application/json' \
     --header 'content-type: application/json'

{
  "id": "1",
  "name": "Example Device Enrollment Instance",
  "supervisionIdentityId": "1",
  "siteId": "-1",
  "serverName": "Acme ASM",
  "serverUuid": "BASD08C11F3C455",
  "adminId": "admin1234",
  "orgName": "Acme Enterprises",
  "orgEmail": "admin@test.com",
  "orgPhone": "555-0123",
  "orgAddress": "124 Conch Street, Bikini Bottom, Pacific Ocean",
  "tokenExpirationDate": "2000-10-30"
}
-----
Retrieve a list of Devices assigned to the supplied id
get
https://yourServer.jamfcloud.com/api/v1/device-enrollments/{id}/devices


Retrieves a list of devices assigned to the supplied id

Path Params
id
string
required
Device Enrollment Instance identifier

1
Response

200
Success

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
deviceEnrollmentProgramInstanceId
string
prestageId
string
serialNumber
string
description
string
model
string
color
string
assetTag
string
profileStatus
string
enum
EMPTY ASSIGNED PUSHED REMOVED

syncState
object

syncState object
profileAssignTime
string
profilePushTime
string
deviceAssignedDate
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/device-enrollments/1/devices \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "deviceEnrollmentProgramInstanceId": "1",
      "prestageId": "1",
      "serialNumber": "R7QFDE2YCFN4",
      "description": "MBP 15.4",
      "model": "Macbook Pro 15 Retina",
      "color": "BLACK",
      "assetTag": "ACME-1234",
      "profileStatus": "ASSIGNED",
      "syncState": {
        "id": 1,
        "serialNumber": "R7QFDE2YCFN4",
        "profileUUID": "9164E5F7C74C2A4C4BE90BB15E549F14",
        "syncStatus": "ASSIGN_SUCCESS",
        "failureCount": 0,
        "timestamp": 1583855813080
      },
      "profileAssignTime": "2000-10-30T18:00:00-00:00",
      "profilePushTime": "2000-10-30T18:00:00-00:00",
      "deviceAssignedDate": "2000-10-30T18:00:00-00:00"
    }
  ]
}
-----