Search for sorted and paged macOS branding configurations
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos

Search for sorted and paged macOS branding configurations

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
Sorting criteria in the format: property:asc/desc. Default sort is id:asc. Multiple sort criteria are supported and must be separated with a comma. Example: sort=id:desc,brandingName:asc


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
object
id
string
applicationName
string
brandingName
string
brandingNameSecondary
string
iconId
integer
brandingHeaderImageId
integer
homeHeading
string
homeSubheading
string

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/self-service/branding/macos?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "applicationName": "Self Service",
      "brandingName": "Self Service",
      "brandingNameSecondary": "Self Service",
      "iconId": 1,
      "brandingHeaderImageId": 1,
      "homeHeading": "Welcome to Self Service",
      "homeSubheading": "We help organizations succeed with Apple"
    }
  ]
}
-----
Create a Self Service macOS branding configuration with the supplied
post
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos

Create a Self Service macOS branding configuration with the supplied details

Body Params
The macOS branding configuration to create

applicationName
string
Self Service
brandingName
string
Self Service
brandingNameSecondary
string
Self Service
iconId
integer
1
brandingHeaderImageId
integer
1
homeHeading
string
Welcome to Self Service
homeSubheading
string
We help organizations succeed with Apple
Response

201
Successful response

Response body
object
id
string
length ≥ 1
href
string

curl --request POST \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/macos \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "applicationName": "Self Service",
  "brandingName": "Self Service",
  "brandingNameSecondary": "Self Service",
  "iconId": 1,
  "brandingHeaderImageId": 1,
  "homeHeading": "Welcome to Self Service",
  "homeSubheading": "We help organizations succeed with Apple"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Read a single Self Service macOS branding configuration indicated by the provided id
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos/{id}

Read a single Self Service macOS branding configuration indicated by the provided id.

Path Params
id
string
required
id of macOS branding configuration

1
Response

200
Successful response

Response body
object
id
string
applicationName
string
brandingName
string
brandingNameSecondary
string
iconId
integer
brandingHeaderImageId
integer
homeHeading
string
homeSubheading
string

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/macos/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "applicationName": "Self Service",
  "brandingName": "Self Service",
  "brandingNameSecondary": "Self Service",
  "iconId": 1,
  "brandingHeaderImageId": 1,
  "homeHeading": "Welcome to Self Service",
  "homeSubheading": "We help organizations succeed with Apple"
}
-----
Update a Self Service macOS branding configuration with the supplied details
put
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos/{id}

Update a Self Service macOS branding configuration with the supplied details

Path Params
id
string
required
id of macOS branding configuration

1
Body Params
The macOS branding configuration values to update

applicationName
string
Self Service
brandingName
string
Self Service
brandingNameSecondary
string
Self Service
iconId
integer
1
brandingHeaderImageId
integer
1
homeHeading
string
Welcome to Self Service
homeSubheading
string
We help organizations succeed with Apple
Responses

200
Successful response

Response body
object
id
string
applicationName
string
brandingName
string
brandingNameSecondary
string
iconId
integer
brandingHeaderImageId
integer
homeHeading
string
homeSubheading
string

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/macos/1 \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "applicationName": "Self Service",
  "brandingName": "Self Service",
  "brandingNameSecondary": "Self Service",
  "iconId": 1,
  "brandingHeaderImageId": 1,
  "homeHeading": "Welcome to Self Service",
  "homeSubheading": "We help organizations succeed with Apple"
}
'

{
  "id": "1",
  "applicationName": "Self Service",
  "brandingName": "Self Service",
  "brandingNameSecondary": "Self Service",
  "iconId": 1,
  "brandingHeaderImageId": 1,
  "homeHeading": "Welcome to Self Service",
  "homeSubheading": "We help organizations succeed with Apple"
}
-----

Delete the Self Service macOS branding configuration indicated by the provided id
delete
https://yourServer.jamfcloud.com/api/v1/self-service/branding/macos/{id}

Delete the Self Service macOS branding configuration indicated by the provided id.

Path Params
id
string
required
id of macOS branding configuration

Responses
204
Successful response

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/macos/ \
     --header 'accept: application/json'