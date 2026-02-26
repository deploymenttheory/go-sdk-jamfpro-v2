Search for sorted and paged iOS branding configurations
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios

Search for sorted and paged iOS branding configurations

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
brandingName
string
required
iconId
integer
headerBackgroundColorCode
string
required
menuIconColorCode
string
required
brandingNameColorCode
string
required
statusBarTextColor
string
required

curl --request GET \
     --url 'https://yourserver.jamfcloud.com/api/v1/self-service/branding/ios?page=0&page-size=100&sort=id%3Aasc' \
     --header 'accept: application/json'

{
  "totalCount": 1,
  "results": [
    {
      "id": "1",
      "brandingName": "Self Service",
      "iconId": 1,
      "headerBackgroundColorCode": "FFFFFF",
      "menuIconColorCode": "000001",
      "brandingNameColorCode": "000000",
      "statusBarTextColor": "dark"
    }
  ]
}
-----
Create a Self Service iOS branding configuration with the supplied
post
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios

Create a Self Service iOS branding configuration with the supplied details

Body Params
The iOS branding configuration to create

brandingName
string
required
Self Service
iconId
integer
1
headerBackgroundColorCode
string
required
FFFFFF
menuIconColorCode
string
required
000001
brandingNameColorCode
string
required
000000
statusBarTextColor
string
required
dark
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
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/ios \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "brandingName": "Self Service",
  "iconId": 1,
  "headerBackgroundColorCode": "FFFFFF",
  "menuIconColorCode": "000001",
  "brandingNameColorCode": "000000",
  "statusBarTextColor": "dark"
}
'

{
  "id": "1",
  "href": "https://yourJamfProUrl.jamf/api/v1/resource/1"
}
-----
Read a single Self Service iOS branding configuration indicated by the provided id
get
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios/{id}

Read a single Self Service iOS branding configuration indicated by the provided id.

Path Params
id
string
required
id of iOS branding configuration

1
Response

200
Successful response

Response body
object
id
string
brandingName
string
required
iconId
integer
headerBackgroundColorCode
string
required
menuIconColorCode
string
required
brandingNameColorCode
string
required
statusBarTextColor
string
required

curl --request GET \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/ios/1 \
     --header 'accept: application/json'

{
  "id": "1",
  "brandingName": "Self Service",
  "iconId": 1,
  "headerBackgroundColorCode": "FFFFFF",
  "menuIconColorCode": "000001",
  "brandingNameColorCode": "000000",
  "statusBarTextColor": "dark"
}
-----
Update a Self Service iOS branding configuration with the supplied details
put
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios/{id}

Update a Self Service iOS branding configuration with the supplied details

Path Params
id
string
required
id of iOS branding configuration

Body Params
The iOS branding configuration values to update

brandingName
string
required
Self Service
iconId
integer
1
headerBackgroundColorCode
string
required
FFFFFF
menuIconColorCode
string
required
000001
brandingNameColorCode
string
required
000000
statusBarTextColor
string
required
dark
Responses

200
Successful response

Response body
object
id
string
brandingName
string
required
iconId
integer
headerBackgroundColorCode
string
required
menuIconColorCode
string
required
brandingNameColorCode
string
required
statusBarTextColor
string
required

curl --request PUT \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/ios/ \
     --header 'accept: application/json' \
     --header 'content-type: application/json' \
     --data '
{
  "brandingName": "Self Service",
  "iconId": 1,
  "headerBackgroundColorCode": "FFFFFF",
  "menuIconColorCode": "000001",
  "statusBarTextColor": "dark",
  "brandingNameColorCode": "000000"
}
'

{
  "id": "1",
  "brandingName": "Self Service",
  "iconId": 1,
  "headerBackgroundColorCode": "FFFFFF",
  "menuIconColorCode": "000001",
  "brandingNameColorCode": "000000",
  "statusBarTextColor": "dark"
}
-----
Delete the Self Service iOS branding configuration indicated by the provided id
delete
https://yourServer.jamfcloud.com/api/v1/self-service/branding/ios/{id}

Delete the Self Service iOS branding configuration indicated by the provided id.

Path Params
id
string
required
id of iOS branding configuration

Responses
204
Successful response

curl --request DELETE \
     --url https://yourserver.jamfcloud.com/api/v1/self-service/branding/ios/ \
     --header 'accept: application/json'
